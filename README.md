# Debug - Simple and flexible debug library for your library

# Usage
```go
package main

import (
  "http"

  "https://github.com/m1ome/debug"
)

func main() {
  debug := debug.NewDebugger("api")

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    debug.Info("New API request /")

    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  http.ListenAndServe(":8080", nil)
}

```

Now if you wish to see debug information just use `--DEBUG=*`

# Wildcards
The * character may be used as a wildcard. Suppose for example your library has debuggers named "api:v1", "api:v1", "api:v3", instead of listing all three with DEBUG=api:v1,api:v2,api:v2, you may simply do DEBUG=api:*, or to run everything using this module simply use DEBUG=*.
