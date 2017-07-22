/* if/else statements in Go */

package main

import "fmt"

func main() {
  start := true

  // traditional if else statement
  if start == true {
    fmt.Println("Start")
  } else {
    fmt.Println("Stop")
  }

  // single if statement
  if true != false {
    fmt.Println("true not equals false!")
  }

  // declare variables in conditions
  if val := 5; val == 0 {
    fmt.Println("val = 0")
  } else if val == 5 {
    fmt.Println("val = 5")
  } else {
    fmt.Println("val is not equal to 0 and 5")
  }

}