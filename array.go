/* Arrays in Go */

package main

import "fmt"

func main() {
  // simple array declaration
  var arr [5]int

  fmt.Println("my arr:", arr)
  arr[2] = 30
  fmt.Println(arr[2])

  // array declaration + initialization
  array := [5]int{1,2,3,4,5}
  fmt.Println("array: ", array)

  fmt.Println("length of arr: ", len(arr))

  // 2D arrays
  var td [3][3]int

  // nested loops
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      td[i][j] = i * j
    }
  }

  fmt.Println("My 2D array: ", td)
}