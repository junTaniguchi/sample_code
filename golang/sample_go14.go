package main
import(
  "fmt"
)

type myInt int
type myStruct struct{
  x, y int
}

// myInt型のメソッドの宣言
func (m myInt) add(n int) myInt{
  return m + myInt(n)
}

// myStruct型のメソッドの宣言
func (m myStruct) add(n int) myStruct{
  return myStruct{
    m.x + n,
    m.y + n,
  }
}

func main(){
  //myInt型のメソッド呼び出し
  var v myInt = 1
  fmt.Println(v.add(2))

  //myStruct型のメソッドの宣言
  var s  = myStruct{4, 5}
  fmt.Println(s.add(6))
}
