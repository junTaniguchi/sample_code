// インターフェース
// メソッド
// 構造体

// interface{}

package main

import "fmt"

func show(t interface{}) {
    // 型アサーション
    // _, ok := t.(japanese)
    // if ok {
    //     fmt.Println("i am japanese")
    // } else {
    //     fmt.Println("i am not japanese")
    // }

    // 型Switch
    switch t.(type) {
    case japanese:
        fmt.Println("i am japanese")
    default:
        fmt.Println("i am not japanese")
    }
}

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

func main() {
    greeters := []greeter{japanese{}, american{}}
    for _, greeter := range greeters {
        greeter.greet()
        show(greeter)
    }
}
