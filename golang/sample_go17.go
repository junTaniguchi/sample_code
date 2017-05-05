package main
import(
  "fmt"
  "net/http"
)

// ServerHTMLの型宣言
type serverHTML struct{
  title string
  body string
}
// ServerHTTPメソッド(ハンドラー関数)の宣言
func (s serverHTML) ServeHTTP(w http.ResponseWriter, r *http.Request){
  h := `
    <html>
      <head>
        <titile>` + s.title + `</title>
      </head>
      <body>` + s.body + `</body>
    </html>
  `
  fmt.Fprint(w, h)
}

func main(){
  //URLごとに値を登録(各ページごとの構造体を生成)
  helloHTML  := &serverHTML{"Hello Title", "Hello Body"}
  goodbyHTML := &serverHTML{"Goodby Title", "Goodby Body"}
  //HandlerインタフェースとHandle関数を元にhttp.Handle関数を登録
  //  Handle関数は対象となるURLのリクエストを受け取ると、
  //  登録した値のServeHTTPメソッドを呼び出す
  http.Handle("/hello", helloHTML)
  http.Handle("/goodby", goodbyHTML)

  //webサーバを起動
  if err := http.ListenAndServe(":4000", nil);err != nil{
    fmt.Println(err)
  }
}
