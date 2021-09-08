package memo 

import (
	"time"
)

type Memo struct {
	Title []byte
	Body []byte
	CreatedAt time.Time
	UpdatedAt time.Time
	r MemoRepository
}

func NewMemo(repo MemoRepository, title []byte, body []byte) Memo {
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
	ms, err = m.r.List(offset, count)
	return

}
