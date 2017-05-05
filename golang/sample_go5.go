// goroutine : 並行処理
// channel

package main

import (
    "fmt"
    "time"
)

func task1(result chan string) {
    time.Sleep(time.Second * 2)
    // fmt.Println("task1 finished!")
    result<- "task1 result"
}

func task2() {
    fmt.Println("task2 finished!")
}

func main() {
    result := make(chan string)

    go task1(result)
    go task2()

    fmt.Println(<-result)

    time.Sleep(time.Second * 3)
}
