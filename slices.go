package main
import "fmt"

func main() {
  s := make([]string, 3)
  fmt.Println("empty: ", s)

  // set values
  s[0] = "a"
  s[1] = "b"
  s[2] = "c"
  fmt.Println("set: ", s)
  fmt.Println("get: ", s[2])
  fmt.Println("length: ", len(s))

  s = append(s, "d")
  s = append(s, "e", "f")
  fmt.Println("append: ", s)

  c := make([]string, len(s))
  copy(c, s) // c has the same values as s
  fmt.Println("\ncopy: ", c)

  l := s[2:5] // assign the values of s from index number #2 till the index 5 to new slice
  fmt.Println("sl1: ", l)
  l = s[:5] // assign the values of s till the index 5 to new slice
  fmt.Println("sl2: ", l)

  t := []string{"a","v","f"}
  fmt.Println("dcl", t)

  // 2D slices
  twoD := make([][]int, 3)
  fmt.Println("\n2D slice: ", twoD)
  for i := 0; i < 3; i++ {
    innerLen := i + 1
    twoD[i] = make([]int, innerLen)
    for j := 0; j < innerLen; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2D: ", twoD)

  fmt.Println("Capacity:", cap(twoD))
  fmt.Println("Length:", len(twoD))

  b := []byte{'g','o','l','a','n','g'}
  fmt.Println("Byte:", b)
  d := b[0:2]
  d[1] = 'c'
  fmt.Println("Original array:", b)

  // array
  x := [3]string{"Лайка", "Белка", "Стрелка"}
  as := x[:] // a slice referencing the storage of x
  fmt.Println("Storage x:", as)
}
