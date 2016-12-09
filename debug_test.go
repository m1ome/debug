package debug

import (
  "testing"
  "os"
  "bytes"
  _ "fmt"
  "strings"
)

func TestCreation(t *testing.T) {
  buf := new (bytes.Buffer)

  os.Setenv("DEBUG", "worker:*")
  d1 := NewDebugger("worker:one")
  d1.SetOutput(buf)

  str := ""

  d1.Info("Hello %s", "friend")
  str = buf.String()

  if (!strings.Contains(str, "Hello friend")) {
    t.Fail()
  }

  if (!strings.Contains(str, "worker:one")) {
    t.Fail()
  }

  d1.Info("Hi once again!")
  str = buf.String()

  if (!strings.Contains(str, "Hi once again!")) {
    t.Fail()
  }

  if (!strings.Contains(str, "worker:one")) {
    t.Fail()
  }
}
