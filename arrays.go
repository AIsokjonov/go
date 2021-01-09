package main
import "fmt"

func main() {
  // array declaration with int type
  // if you set the length of the array while declaring it,
  // the array initializes all the values to 0(zero)
  var a [5]int
  fmt.Println("empty: ", a)

  a[4] = 67
  fmt.Println("set: ", a)
  fmt.Println("get: ", a[4])

  fmt.Println("len: ", len(a))

  // dynamic declaration
  b := [5]int{12,42,-3,23,0}
  fmt.Println("dcl: ", b)

  // 2-Dimentional arrays
  var twoD [2][3]int
  for i := 0; i < 2; i++ {
    for j := 0; j < 3; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2D array: ", twoD)

  // 3-Dimentional arrays
  var threeD [3][2][4]int
  for i := 0; i < 3; i++ {
    for j := 0; j < 2; j++ {
      for k := 0; k < 4; k++ {
        threeD[i][j][k] = i + j + k
      }
    }
  }
  fmt.Println("3D array: ", threeD)
}
