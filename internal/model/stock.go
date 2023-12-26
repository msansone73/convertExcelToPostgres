package model

import (
	"log"
)

type Stock struct{
	Id int64
	Code string
	Name string
}

func (s *Stock) GetStockById(id int64) {
	db:= getConnection()
	defer db.Close()

	err := db.QueryRow("SELECT id, code, name FROM stocks WHERE id = $1", id).
		Scan(&s.Id, &s.Code, &s.Name)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Stock) GetStockByCode(code string) error {
	db:= getConnection()
	defer db.Close()

	err := db.QueryRow("SELECT id, code, name FROM stocks WHERE code = $1", code).
		Scan(&s.Id, &s.Code, &s.Name)
	if err != nil {
		//log.Fatal(err)
		log.Println("code não endontrado.", code)
		return err
	}
	return nil
}