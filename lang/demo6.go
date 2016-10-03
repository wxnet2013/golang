package main

import (
  "fmt"
  "math/cmplx"
)

// 1. bool
// 2. string
// 3. int  int8  int16  int32(rune)  int64
// 4. uint uint8(byte) uint16 uint32 uint64 uintptr
// 5. float32 float64
// 6. complex64 complex128

var a, b, c int = 1, 2, 3

var (
  isLogin bool      = false
  title   string    = "hahaha"
  z       complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
  const format = "%T(%v)\n"

  d := 4
  var e int
  // 布尔值的默认值为false，零值
  var f bool

  var g float32 = 1.3

  var h string = "hahaha"

  // 类型推导
  i := 111

  fmt.Println(a, b, c, d, e, f, g, h, z, isLogin, title)
  fmt.Printf(format, title, title)
  fmt.Printf(format, i, i)
}
