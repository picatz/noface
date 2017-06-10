package main

import (
  "os"
  "fmt"
  "github.com/picatz/noface"
)

// Let the party begin.
func main() {

  result := make(chan string)
  go func() {
    // Get that function output.
    iface, e := noface.FirstIface()
    // Did we error out?
    if e != nil {
      // Oh noes.
      os.Exit(1)
    } else {
      // Yay!
      result <- iface
    }
  }()
  // Print the network interface.
  fmt.Println(<- result)

}
