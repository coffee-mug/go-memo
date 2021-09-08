package memo

type MemoRepository interface {
	Save(m Memo) (id int, err error)
	Get (id int) (m *Memo, err error)
	List(offset int, count int) (ms []Memo, err error)
}