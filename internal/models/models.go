package models

type DB interface {
	Connect(db string) (DB, error)
	Write() error
	Close() error
}
