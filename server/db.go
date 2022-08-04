package server

import (
	"database/sql"
	"encoding/json"
	"errors"
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
	DB_PORT     = "5431"
)

type Database struct {
	DB *sqlx.DB
}

func (d *Database) DBInit() {
	connStr := fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	d.DB = db
}

func (d *Database) Create(order model.Order) error {
	js, _ := json.Marshal(order)
	_, err := d.DB.Exec("INSERT INTO orders VALUES ($1, $2);", order.OrderUid, js)
	if err != nil {
		return errors.New("Order already exists")
	}
	return nil
}
func (d *Database) Get(order_uid string) (model.Order, error) {
	var res model.Order
	err := d.DB.QueryRow("SELECT json_data FROM orders WHERE order_uid=$1", order_uid).Scan(&res)
	if err != nil {
		return model.Order{}, err
	}
	return res, nil
}
func (d *Database) GetAll() (*sql.Rows, error) {
	res, err := d.DB.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *Database) CloseDb() {
	d.DB.Close()
}
