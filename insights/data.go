package insights

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" //Database driver
)

type Data struct {
	db              *sql.DB
	AdDownloads     AdDownloadsTable
	AdViews         AdViewsTable
	AdClickThroughs AdClickThroughsTable
	AdConversions   AdConversionsTable
}

func (d *Data) Init() {

	log.Println("Data INIT")

	db, err := sql.Open("sqlite3", "./insights.db")
	checkErr(err)
	d.db = db

	d.AdDownloads = &AdDownloadsSQLLite{db: db}
	d.AdDownloads.Init()

	d.AdViews = &AdViewsSQLLite{db: db}
	d.AdViews.Init()

	d.AdClickThroughs = &AdClickThroughsSQLLite{db: db}
	d.AdClickThroughs.Init()

	d.AdConversions = &AdConversionsSQLLite{db: db}
	d.AdConversions.Init()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (d *Data) Close() {
	d.db.Close()
}
