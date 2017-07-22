/* Switch statements in Go */

package main

import "fmt"
import "time"

func main() {
  val := 10

  // simple switch statement
  switch val {
    case 1: 
      fmt.Println("value is 1")
    case 5:
      fmt.Println("value is 5")
    case 10:
      fmt.Println("value is 10")
    default:
      fmt.Println("value is not 1, 5 and 10")
  }

  switch time.Now().Weekday() {
    case time.Saturday:
      fmt.Println("Today is Saturday")
    case time.Sunday, time.Monday:
      fmt.Println("Today is Sunday or Monday")
    default:
      fmt.Println("Today is not Mon, Sat or Sun")
  }

  t := time.Now()
  // switch like if/else statement
  switch {
    case t.Hour() < 12:
      fmt.Println("It's before noon")
    default:
      fmt.Println("It's after noon")
  }

  // switch inside function, check data type of argument passed
  WhichDataType := func(h interface{}) {
    switch x := h.(type) {
      case string:
        fmt.Println("h is a string")
      case int:
        fmt.Println("h is an integer")
      case bool:
        fmt.Println("h is a boolean")
      default:
        fmt.Println("don't know type", x)
    }  
  }
  WhichDataType("hello")
}