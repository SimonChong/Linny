package insights

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Data struct {
	connection    *sql.DB
	Impressions   *Impressions
	ClickThroughs *ClickThroughs
}

func (d *Data) Init() {

	fmt.Println("Data INIT")

	db, err := sql.Open("sqlite3", "./insights.db")
	checkErr(err)
	d.connection = db

	d.Impressions = new(Impressions)
	d.Impressions.Init(d.connection)

	d.ClickThroughs = new(ClickThroughs)
	d.ClickThroughs.Init(d.connection)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (d *Data) Close() {
	d.connection.Close()
}
