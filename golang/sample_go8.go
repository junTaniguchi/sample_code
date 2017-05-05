package main

import (
  "fmt"
  "net/http"
  "html"
)

type Server struct{}

func (Server) ServeHTTP(w http.ResponseWriter, r *http.Request){
  v := r.FormValue("input_value")

  h := `
        <html>
          <head>
            <title>
              HTMLのフォーム
            </titile>
          </head>
          <body>
            <form action="/" method="post">
              <input type="text" name="input_value">
              <input type="submit" name="送信"><br>
              入力値：` + html.EscapeString(v) + `
            </form>
          </body>
        </html>
  `
  //クライアント（ブラウザ）にhtmlを送信
  fmt.Fprint(w, h)
}

func main(){
  //webサーバを起動
  http.ListenAndServe(":4000", Server{})
}
