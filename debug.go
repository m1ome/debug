// Package debug provides basic functionality of debugger for your Library/Application.
package debug

import (
  "os"
  "strings"
  "fmt"
  "regexp"
  "math/rand"
  "time"
  "io"
)

type debugger struct{
  name string
  enabled bool
  colorized bool
  color int
  timing time.Time
  started bool
  stream io.Writer
}

func NewDebugger(name string) *debugger {
  debugEnv := strings.ToLower(os.Getenv("DEBUG"))
  colorized := !(strings.ToLower(os.Getenv("DEBUG_COLOR")) == "no")
  color := rand.Intn(6)

  switch debugEnv {
  case "":
    return &debugger{enabled: false}
  case "*":
    return &debugger{name: name, enabled: true, colorized: colorized, color: color, stream: os.Stdout}
  }

  names := strings.Split(debugEnv, ",")
  for _, pattern := range names {
    match, _ := regexp.MatchString(pattern, name)

    if match {
      return &debugger{name: name, enabled: true, colorized: colorized, color: color, stream: os.Stdout}
    }
  }

  return &debugger{enabled: false}
}

func (d *debugger) SetOutput(w io.Writer) {
  d.stream = w
}

func (d *debugger) Info(format string, v ...interface{}) {
  if (!d.enabled) {
    return;
  }

  m := fmt.Sprintf(format, v...)
  s := ""
  if (d.colorized) {
    since := "0"
    if (d.started) {
      since = time.Since(d.timing).String()
    }

    s = fmt.Sprintf("  \u001b[3%d;1m%s\u001b[0m \u001b[3%dm+%s\u001b[0m > %s", 1, d.name, 1, since, m)
  } else {
    s = fmt.Sprintf("%s %s", d.name, m)
  }

  d.timing = time.Now()
  d.started = true
  fmt.Fprintln(d.stream, s)
}
