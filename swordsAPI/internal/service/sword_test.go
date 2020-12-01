package service

import (
	"fmt"
	"github.com/Ezefalcon/golang/swordsAPI/internal/config"
	"github.com/Ezefalcon/golang/swordsAPI/internal/database"
	"testing"
)

var s Service

func init() {
	//configFile := flag.String("config", "./configs/config-test.yaml", "this is the service config")
	//flag.Parse()
	conf, err := config.LoadConfig("../../configs/config-test.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	db, err := database.NewDatabase(conf)
	database.CreateSchema(db)
	if err != nil {
		fmt.Println(err.Error())
	}
	s, err = NewSwordService(conf,db)
}

func TestService_AddSword(t *testing.T) {
	sword := Sword{
		Name: "Excalibur",
		Speciality: "Fire",
		Damage: 200,
	}
	insertedId, err := s.AddSword(sword)
	foundSword := s.FindById(int(insertedId))
	if err != nil || foundSword == nil {
		t.Error("The sword could not be added")
	}
}

func TestService_FindById(t *testing.T) {
	addedSwordId, _ := s.AddSword(Sword{Name: "KingSword", Speciality: "Strength", Damage: 100})
	sword := s.FindById(int(addedSwordId))

	if sword == nil {
		t.Error("Could not find the provided id")
	}
}

func TestService_Update(t *testing.T) {

	addedSwordId, _ := s.AddSword(Sword{Name: "KingSword", Speciality: "Strength", Damage: 100})

	s.UpdateSword(Sword{Name: "KingSword", Speciality: "Strength", Damage: 150}, int(addedSwordId))
	sword := s.FindById(int(addedSwordId))

	if sword == nil {
		t.Error("The sword could not be updated")
	} else if sword.Damage != 150 {
		t.Error("The update is wrong")
	}
}

func TestService_DeleteById(t *testing.T) {
	addedId, _ := s.AddSword(Sword{Name: "KingSword", Speciality: "Strength", Damage: 100})
	s.DeleteById(int(addedId))
	sword := s.FindById(int(addedId))
	if sword != nil {
		t.Error("sword was not deleted")
	}
}