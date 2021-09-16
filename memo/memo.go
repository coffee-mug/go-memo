package memo 

import (
	"log"
	"time"
	"errors"
)

type Memo struct {
	Title string
	Body string 
	CreatedAt time.Time
	UpdatedAt time.Time
	r MemoRepository
}

func NewMemo(repo MemoRepository, title string, body string) Memo {
	createdAt := time.Now()
	updatedAt := time.Now()

	return Memo{
		Title: title,
		Body: body,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		r: repo,
	}

}

func (m Memo) Save() (id int, err error) {
	// Call to db code
	id, err = m.r.Save(m)
	return
}

func (m Memo) Get(id int) (mo *Memo, err error) {
	mo, err = m.r.Get(id)
	return
}

func (m Memo) List(offset int, count int) (ms []Memo, err error) {
	data, err := m.r.Data()

	log.Printf("%v", data)
	if err != nil {
		log.Fatal(err)
	}
	if count + offset > len(data) {
		return nil, errors.New("Invalid range")
	}

	if count == -1 {
		count = len(data)
	} 

	ms, err = m.r.List(offset, count)
	return

}
