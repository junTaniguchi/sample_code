package main

import(
  "fmt"
)

func typeSwitch(i interface{}){
  //iに格納されている値の型で分岐する
  switch v := i.(type){
  case int:
    fmt.Println("int:", v)
  case string:
    fmt.Println("string:", v)
  case int32, int64:
    fmt.Println("int32/64", v)
  default:
    fmt.Println("default:", v)
  }
}
func typeSwitchNil(i interface{}){
  // iに格納されている値の型で分岐する
  switch v := i.(type){
  case nil: //型のないnilのみ該当する
    fmt.Println("nil:", v)
  case *int: // *int型のnilについてはこちらに該当する
    fmt.Println("*int:", v)
  }
}

func main(){
  typeSwitch(1)
  typeSwitch("a")
  typeSwitch(int32(1))
  typeSwitch(1.1)

  typeSwitchNil(nil)
  var n *int = nil
  typeSwitchNil(n)
}
