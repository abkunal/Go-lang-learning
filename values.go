/* Data types in Go, string, integers, float, boolean */

package main

import "fmt"

func main() {
  fmt.Println("This is Go " + "language!")

  fmt.Println("2 * 10 = ", 2 * 10)
  fmt.Println(7/3)
  fmt.Println(7.0/3)
  fmt.Println(7/3.0)

  fmt.Println(true && true)
  fmt.Println(true || true)
  fmt.Println(!true)
}