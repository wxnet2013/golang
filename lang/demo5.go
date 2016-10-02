package main
import "fmt"

func add1(x int, y int) int {
  return x + y
}

func add2(x, y, z int) int {
  return x + y + z
}


// 返回多个值
func swap(x, y string) (string, string) {
  return y, x
}

// 命名返回值
func split(sum int) (x, y, z int) {
  x = sum * 4 / 9
	y = sum - x
  z = 0
	return
}

func main() {
  fmt.Println("add1", add1(1,2))

  fmt.Println("add2", add2(1,2,3))

  a, b := swap("hello", "world")
  fmt.Println("swap", a, b)

  fmt.Println(split(9))
}
