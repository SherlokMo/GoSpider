package memorydb

type Ikey interface {
	int | string
}

type IMemoryDB[key Ikey, value any] interface {
	Store(k key, v value)
	Exist(k key) bool
}
