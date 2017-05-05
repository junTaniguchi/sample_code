package main

import (
  "fmt"
  "net/http"
)

// ServeHTTPメソッド用の構造体型
type Server struct{}

//httpリクエストを受け取るハンドラ
func (Server) ServeHTTP(w http.ResponseWriter, r *http.Request){
  h := `
        <html>
          <head>
            <titile>
              Hello
            </titile>
          </head>
          <body>
            Hello
          </body>
        </html>
      `
  fmt.Fprint(w, h)
}

func main(){
  //Webサーバを起動
  http.ListenAndServe(":4000", Server{})
}
