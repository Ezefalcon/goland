package service

import (
	"fmt"
	"github.com/Ezefalcon/golang/swordsAPI/internal/config"
	"github.com/jmoiron/sqlx"
)

type Sword struct {
	ID int64
	Name string
	Speciality string
	Damage int32
}

// All the structs that implements this methods will be of "type" Service
type SwordService interface {
	AddSword(Sword) (int64, error)
	FindById(int) *Sword
	FindAll() []*Sword
	UpdateSword(Sword, int) (int64, error)
	DeleteById(int)
}

// The implementation of the service with the necessary attributes to do it
type service struct {
	conf *config.Config
	db *sqlx.DB
}

// New instance of service
func NewSwordService(c *config.Config, db *sqlx.DB) (SwordService, error) {
	return service{c, db}, nil
}

func (s service) AddSword(sword Sword) (int64, error) {
	exec := s.db.MustExec("INSERT INTO sword(name, speciality, damage) VALUES ($1, $2, $3)", sword.Name, sword.Speciality, sword.Damage)
	return exec.LastInsertId()
}

func (s service) UpdateSword(sword Sword, id int) (int64, error) {
	exec := s.db.MustExec("UPDATE sword SET name = $1, speciality = $2, damage = $3 WHERE ID = $4", sword.Name, sword.Speciality, sword.Damage, id)
	return exec.LastInsertId()
}

func (s service) FindById(id int) *Sword {
	var sword Sword
	err := s.db.Get(&sword, "SELECT * FROM sword WHERE id=$1", int64(id))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &sword
}

func (s service) FindAll() []*Sword {
	var swords []*Sword
	err := s.db.Select(&swords, "SELECT * FROM sword")
	if err != nil {
		fmt.Println(err)
	}
	return swords
}

func (s service) DeleteById(id int) {
	s.db.MustExec("DELETE FROM sword where id = $1", id)
}