package main

import (
  "fmt"
  "net/http"
)

func helloHandle(w http.ResponseWriter, r *http.Request){

  h := `
        <html>
          <head>
            <title>
              HELLO
            </titile>
          </head>
          <body>
            hello
          </body>
        </html>
  `
  //クライアント（ブラウザ）にhtmlを送信
  fmt.Fprint(w, h)
}

func goodbyHandle(w http.ResponseWriter, r *http.Request){

  h := `
        <html>
          <head>
            <title>
              GOODBY
            </titile>
          </head>
          <body>
            goodby
          </body>
        </html>
  `
  //クライアント（ブラウザ）にhtmlを送信
  fmt.Fprint(w, h)
}

func main(){
  //URLごとに関数を登録
  http.HandleFunc("/hello", helloHandle)
  http.HandleFunc("/goodby", goodbyHandle)
  //webサーバを起動
  if err := http.ListenAndServe(":4000", nil); err != nil{
    fmt.Print(err)
  }
}
