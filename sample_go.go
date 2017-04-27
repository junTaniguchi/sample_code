package main

import "fmt"

func main(){
  //スライスの使い方
  s := []int{}
  // append
  s = append(s, 8, 2, 10)
  // copy
  t := make([]int, len(s))
  n := copy(t, s)
  fmt.Println(s)
  fmt.Println(t)
  fmt.Println(n)

  //map(連想配列)の使い方
  m := map[string]int{"taguchi":100, "fkoji":200}
  fmt.Println(m)
  fmt.Println(len(m))
  delete(m, "taguchi")
  fmt.Println(m)
  v, ok := m["fkoji"]
  fmt.Println(v)
  fmt.Println(ok)

  //if文の使い方
  if score := 43; score > 80{
    fmt.Printf("excellent score :%d\n", score)
  } else if score > 60{
    fmt.Printf("awesome score :%d\n", score)
  }else {
    fmt.Printf("so so score :%d\n", score)
  }

  //case文の使い方
  // signal := "blue"
  // switch signal {
  // case "red":
  //     fmt.Println("Stop")
  // case "yellow":
  //     fmt.Println("Caution")
  // case "green", "blue":
  //     fmt.Println("Go")
  // default:
  //     fmt.Println("wrong signal")
  // }
  score := 82
  switch {
  case score > 80:
      fmt.Println("Great!")
  default:
      fmt.Println("so so ...")
  }

  //for文の使い方
  // for i := 0; i < 10; i++ {
  //     // if i == 3 { break }
  //     if i == 3 { continue }
  //     fmt.Println(i)
  // }

  // i := 0
  // for i < 10 {
  //     fmt.Println(i)
  //     i++
  // }
  i := 0
  for {
      fmt.Println(i)
      i++
      if i == 3 { break }
  }

  //range文の使い方
  // s := []int{2, 3, 8}
  // for i, v := range s {
  //     fmt.Println(i, v)
  // }
  // for _, v := range s {
  //     fmt.Println(v)
  // }
  m := map[string]int{"taguchi":200, "fkoji":300}
  for k, v := range m {
      fmt.Println(k, v)
  }

  //構造体の使い方
  // u := new(user)
  // // (*u).name = "taguchi"
  // u.name = "taguchi"
  // u.score = 20

  u := user{name:"taguchi", score:50}
  // fmt.Println(u)
  u.hit()
  u.show()

  //インタフェースの使い方
  greeters := []greeter{japanese{}, american{}}
  for _, greeter := range greeters {
      greeter.greet()
  }


}
//構造体の作り方
type user struct {
    name string
    score int
}
//メソッド（値渡し）の使い方
func (u user) show() {
    fmt.Printf("name:%s, score:%d\n", u.name, u.score)
}

//メソッド（参照渡し）の使い方
func (u *user) hit() {
    u.score++
}

//インタフェースの使い方
type greeter interface {
    greet()
}

type japanese struct {}
type american struct {}

func (j japanese) greet() {
    fmt.Println("Konnnichiwa!")
}
func (a american) greet() {
    fmt.Println("Hello!")
}
