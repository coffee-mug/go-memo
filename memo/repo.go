package memo

type MemoRepository interface {
	Save(m Memo) (id int, err error)
}