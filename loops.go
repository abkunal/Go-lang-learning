/* In Go, there is only one type of loop - for loop */

package main

import "fmt"

func main() {
  i := 1
  j := 2
  
  // only condition for loop
  for i < 10 && j < 10 {
    fmt.Println(i)
    i = i + 1
    j = j + 5
  }

  // classic for loop
  for i := 0; i < 10; i++ {
    fmt.Println(i)
  }

  // loop until explicitly encounter break
  for {
    fmt.Println("Will execute one time")
    break
  }

  fmt.Println("Even numbers between 1 and 20")
  for i := 1; i <= 20; i++ {
    if i % 2 != 0 {
      continue
    }
    fmt.Println(i)
  }
}