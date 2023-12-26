package model

import (
	"log"
	"time"
)


type Transaction struct {
	Id int64
	User_id int64
	Stock_id int64
	Tipo string
	Value float32
	Quantity int
	Data_at time.Time	
}

func (t *Transaction) Adicionar() error {

	db:= getConnection()
	defer db.Close()

	_, err := db.Exec(
		"INSERT INTO transactions (user_id, stock_id, type, value, quantity, data_at) values ($1, $2, $3, $4, $5, $6)",
	&t.User_id, &t.Stock_id, &t.Tipo, &t.Value, &t.Quantity, &t.Data_at)
	if err != nil {
		//log.Fatal(err)
		log.Println("Falha ao inserir.", err)
		return err
	}
	return nil
}