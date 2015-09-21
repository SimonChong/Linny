package insights

import (
	"database/sql"
	"time"
)

type AdClickThroughs struct {
	db     *sql.DB
	insert *sql.Stmt
}

func (c *AdClickThroughs) Init(db *sql.DB) {
	c.db = db

	sql := `CREATE TABLE IF NOT EXISTS "AdClickThroughs" ("id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE, "created" DATETIME DEFAULT CURRENT_TIMESTAMP, "adID" VARCHAR, "refererURL" VARCHAR, "destinationURL" VARCHAR, "originIP" VARCHAR, "linkGeneratedOn" DATETIME, "linkTag" VARCHAR, "sessionID" VARCHAR);`
	c.db.Exec(sql)

	ins, err := db.Prepare(`INSERT INTO AdClickThroughs( adID, refererURL, destinationURL, originIP, linkGeneratedOn, linkTag, sessionID) values(?,?,?,?,?,?,?)`)
	checkErr(err)
	c.insert = ins

}

func (c *AdClickThroughs) Insert(adID string, refererURL string, destinationURL string, originIP string, linkGeneratedOn time.Time, linkTag string, sessionID string) (sql.Result, error) {

	res, err := c.insert.Exec(adID, refererURL, destinationURL, originIP, linkGeneratedOn, linkTag, sessionID)
	checkErr(err)

	return res, err
}
