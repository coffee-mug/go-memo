package main

import (
	"github.com/coffee-mug/go-memo/memo"
	"log"
)

type MemoryRepository struct {
	data []memo.Memo
}

func (m *MemoryRepository) Save(mo memo.Memo) (id int, err error) {
	m.data = append(m.data, mo)
	id = len(m.data)

	return id - 1, nil
}

func NewMemoryRepository() *MemoryRepository {
	d := make([]memo.Memo, 0, 10)
	return &MemoryRepository{
		data: d,
	}
}


func main() {
	r := NewMemoryRepository()

	m := memo.NewMemo(r, []byte("This is a title"), []byte("This is a body"))

	id, _ := m.Save()
	log.Println(id)

	id, _ = m.Save()
	log.Println(id)
}