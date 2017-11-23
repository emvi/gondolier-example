package main

import (
	"database/sql"
	"github.com/emvicom/gondolier"
	_ "github.com/lib/pq"
)

type Customer struct {
	Id   uint64 `gondolier:"type:bigint;id"`
	Name string `gondolier:"type:varchar(255);notnull"`
	Age  int    `gondolier:"type:integer;notnull"`
}

type Order struct {
	Id    uint64 `gondolier:"type:bigint;id"`
	Buyer uint64 `gondolier:"type:bigint;fk:customer.id;notnull"`
}

type OrderPosition struct {
	Id       uint64 `gondolier:"type:bigint;id"`
	Order    uint64 `gondolier:"type:bigint;fk:order.id;notnull"`
	Quantity int    `gondolier:"type:integer;notnull"`
	Cost     int    `gondolier:"type:integer;notnull"`
}

type Obsolete struct{}

func main() {
	db, _ := sql.Open("postgres", dbString())
	defer db.Close()

	postgres := &gondolier.Postgres{Schema: "public",
		DropColumns: true,
		Log:         true}
	gondolier.Use(db, postgres)
	gondolier.Model(Customer{}, Order{}, OrderPosition{})
	gondolier.Drop(Obsolete{})
	gondolier.Migrate()
}

func dbString() string {
	return "host=localhost" +
		" port=5432" +
		" user=postgres" +
		" password=postgres" +
		" dbname=gondolier" +
		" sslmode=disable"
}
