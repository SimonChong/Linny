package insights

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Data struct {
	connection      *sql.DB
	AdDownloads     *AdDownloads
	AdViews         *AdViews
	AdClickThroughs *AdClickThroughs
}

func (d *Data) Init() {

	fmt.Println("Data INIT")

	db, err := sql.Open("sqlite3", "./insights.db")
	checkErr(err)
	d.connection = db

	d.AdDownloads = new(AdDownloads)
	d.AdDownloads.Init(d.connection)

	d.AdViews = new(AdViews)
	d.AdViews.Init(d.connection)

	d.AdClickThroughs = new(AdClickThroughs)
	d.AdClickThroughs.Init(d.connection)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (d *Data) Close() {
	d.connection.Close()
}
