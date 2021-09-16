package storage

import (
	"errors"
	"github.com/coffee-mug/go-memo/memo"
)

type MemoryRepository struct {
	data []memo.Memo
}

func (m *MemoryRepository) Save(mo memo.Memo) (id int, err error) {
	m.data = append(m.data, mo)
	id = len(m.data)

	return id - 1, nil
}

func (m *MemoryRepository) Get(id int) (mo *memo.Memo, err error) {
	if id < len(m.data) {
		return &m.data[id], nil
	}

	return nil, errors.New("No entry for provided index")
}

func (m *MemoryRepository) Data() (ms []memo.Memo, err error) {
	return m.data, nil
}

func (m *MemoryRepository) List(offset, count int) (ms []memo.Memo, err error) {
	return m.data[offset:offset+count], nil
}

func NewMemoryRepository() *MemoryRepository {
	d := make([]memo.Memo, 0, 10)
	return &MemoryRepository{
		data: d,
	}
}
