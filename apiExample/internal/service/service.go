package chat

import (
	"github.com/ezefalcon/goland/internal/config"
	"github.com/jmoiron/sqlx"
)

// Message ...
type Message struct {
	ID int64
	Text string
}

// Service ...
type Service interface {
	AddMessage(Message) error
	FindById(int) *Message
	FindAll() []*Message
}

type service struct {
	conf *config.Config
	db *sqlx.DB
}

func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{c, db}, nil;
}

func (s service) AddMessage(m Message) error {
	return nil
}

func (s service) FindById(ID int) *Message {
	return nil
}

func (s service) FindAll() []*Message {
	var list []*Message
	s.db.Select(&list, "SELECT * FROM messages")
	return list
}