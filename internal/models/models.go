package models

type DB interface {
	Connect(db string) (DB, error)
	Write(data []Pair) error
	Close() error
}

type Pair struct {
	Key   string
	Value string
}
