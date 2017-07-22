/* Variables in Go */

package main

import "fmt"

func main() {
  var s string = "Kunal"
  t := "Go language"      // shorthand notation for declaring anf initializing
  fmt.Println(s, t)

  var i, j int = 10, 20   // initializing multiple values
  fmt.Println(i, j)

  var k = true
  var l bool = false  
  
  fmt.Println(k, l)

  var x int
  fmt.Println(x)
}