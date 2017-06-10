package main

import (
  "os"
  "fmt"
  "github.com/picatz/noface"
)

// Let the party begin.
func main() {
  // Get that function output.
  iface, e := noface.FirstIface()
  // Did we error out?
  if e != nil {
    // Oh noes.
    fmt.Println(e)
    os.Exit(1)
  } else {
    // Yay!
    fmt.Println(iface)
  }
}
