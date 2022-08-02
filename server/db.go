package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/ksean42/GoWB/model"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "ksean"
	DB_PASSWORD = "pass"
	DB_NAME     = "gowb"
)

type Database struct {
	DB *sqlx.DB
}

func (d *Database) DBInit() {
	connStr := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	d.DB = db
}

func (d *Database) Create(order model.Order) {
	js, _ := json.Marshal(order)
	_, err := d.DB.Exec("INSERT INTO orders VALUES ($1, $2);", order.OrderUid, js)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}
func (d *Database) Get(order_uid string) (model.Order, error) {
	var res model.Order
	err := d.DB.QueryRow("select json_data from orders WHERE order_uid=$1", order_uid).Scan(&res)
	if err != nil {
		return model.Order{}, err
	}
	// fmt.Println(res)
	return res, nil
}
func (d *Database) GetAll() (*sql.Rows, error) {
	res, err := d.DB.Query("select * from orders")
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *Database) CloseDb() {
	d.DB.Close()
}
