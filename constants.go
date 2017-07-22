/* Constants in Go */

package main

import "fmt"
import "math"

func main() {
  const pi = 3.14             // declare constants
  fmt.Println("Pi = ", pi)              

  const e = 2.172
  fmt.Println(e)
  fmt.Println(e / 34.323)

  fmt.Println(int64(5000000000 / 2323)) // convert int to 64 bit int

  fmt.Println(math.Sin(10000000))       // sine of given number
  fmt.Println(int32(3452342/34564))     // convert int to 32 bit int
}