package main

import "fmt"

func main() {
  /////////////////////////////////////////

  //多次元配列
    var a [3][3]int
    b := [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    // var b [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
    fmt.Println(a)
    fmt.Println(b)
  /////////////////////////////////////////

  //多次元スライス
  matrix1 := [][]float64{
    {0, 0},
    {0, 1},
    {1, 0},
    {1, 1},
  }
  matrix2 := make([][]float64, 4)
  matrix2[0] = []float64{0, 0}
  matrix2[1] = []float64{0, 1}
  matrix2[2] = []float64{1, 0}
  matrix2[3] = []float64{1, 1}

  fmt.Println(matrix1)
  fmt.Println(matrix2)
}
