package main

import (
	"github.com/coffee-mug/go-memo/memo"
	"github.com/coffee-mug/go-memo/storage"
	"log"
)


func main() {
	r := storage.NewMemoryRepository()

	m := memo.NewMemo(r, []byte("This is a title"), []byte("This is a body"))

	// TODO: Move below in relevant test cases
	// Test save
	id, _ := m.Save()
	log.Println(id)

	m.Title = []byte("Index 1")

	id, _ = m.Save()
	log.Println(id)

	// Test GET
	mo, _ := m.Get(1)
	log.Println(string(mo.Title))

	// Test LIST
	ms, _ := m.List(0, 2)
	log.Printf("Total memos count: %d", len(ms))

}