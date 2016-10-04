package main

import "fmt"

func main() {
  // 死循环
  // for {
  //   fmt.Println("ha")
  // }
  for i := 0; i < 10; i++ {
    fmt.Println(i)
  }

  // while循环
  sum := 1
  for sum < 100 {
    sum += sum
  }
  fmt.Println(sum)

  // defer
  defer fmt.Println("world")
  fmt.Println("hello")
}
