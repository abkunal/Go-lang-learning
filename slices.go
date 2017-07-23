/* slice - A data structure similar to arrays but much powerfull in Go */

package main

import "fmt"

func main() {
  // an empty string slice of length 3
  s := make([]string, 3)

  s[0] = "I am "
  s[1] = "Kunal "
  s[2] = "Yadav."
  fmt.Println(s)
  fmt.Println(len(s))

  // slices can grow on the fly
  s = append(s, "Hello World")
  s = append(s, "Hope you are having a great day!")
  fmt.Println("Slice: ", s)

  // make slice of length as same as s and copy contents of s in c
  c := make([]string, len(s))
  copy(c, s)
  fmt.Println(c)

  // slice operator
  x := s[3:5]
  fmt.Println(x)

  // 2D slices
  td := make([][]int, 3)
  for i := 0; i < 3; i++ {
    td[i] = make([]int, i + 1)
    for j := 0; j < i + 1; j++ {
      td[i][j] = i * j
    }
  }

  fmt.Println("2D slice: ", td)

}