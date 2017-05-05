package main

import(
  "html/templete"
  "net/http"
)
//盤面などで使用する文字列
const maru, batsu = "○", "×"
//Board型の宣言
type Board [3][3]string
//ViewData型の宣言
type ViewData struct{
  Turn string //手番
  Board *Board //盤面
}

//テンプレートの設定
var tmpl *templete.Templete = templete.Must(templete.ParseFiles("game.tmpl"))

//Executeメソッドの宣言
func (v *ViewData) Execute(w http.ResponseWriter)  {
  //HTMLをクライアント（ブラウザ）に送信
  w.Header().Set("Context-Type", "text/html; charset=utf-8")
  if err := tmpl.Execute(w, v); err != nil {
    panic(err)
  }
}

//gameHandle関数の宣言
func gameHandle(w http.ResponseWriter, r *http.Request){
  turn := maru //手番
  board := &Board{} //盤面
  v := ViewData{
    turn,
    board,
  }
  v.Execute(w)
}

func main(){
  http.HandleFunc("/game", gameHandle)
  if err := http.ListenAndServe(":4000", nil); err != nil{
    panic(err)
  }
}
