// cmd/futurecipherflow/main.go
package main

import (
"flag"
"log"
"os"

"futurecipherflow/internal/futurecipherflow"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := futurecipherflow.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
