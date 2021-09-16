package main

import (
	"fmt"
	"log"
	"html/template"
	"net/http"
	"github.com/coffee-mug/go-memo/memo"
	"github.com/coffee-mug/go-memo/storage"
)


func main() {
	memoStorage := storage.NewMemoryRepository()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		m := memo.NewMemo(memoStorage, "This is a title", "This is a body")

		mo, _ := m.List(0,-1)

		tmpl, err := template.New("test").Parse("<h1>Some thoughts</h1> {{range .}} <article> <h2> {{.Title}} </h2> <div>{{.Body}} </div> <article> {{end}}</p>")
		if err != nil {
			log.Fatal(err)
		}

		tmpl.Execute(w, mo)
	})

	http.HandleFunc("/store", func (w http.ResponseWriter, r *http.Request){
		if r.Method != "POST" {
			fmt.Fprint(w, "Invalid method")
			return
		}

		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		m := memo.NewMemo(memoStorage, r.FormValue("title"), r.FormValue("memo"))
		m.Save()
	})

	http.ListenAndServe("127.0.0.1:9999", nil)

}