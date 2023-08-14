package core

import (
	_ "github.com/lib/pq"
	"github.com/pocketbase/dbx"
)

func connectDB(data string) (*dbx.DB, error) {
	dsn := "postgresql://demo:demo5360@127.0.0.1:26257/" + data + "?sslmode=require&sslrootcert=C%3A%5CUsers%5CAdministrator%5C.cockroach-demo%5Cca.crt"
	db, err := dbx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
