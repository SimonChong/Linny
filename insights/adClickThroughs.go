package insights

import (
	"database/sql"
	"time"
)

type AdClickThroughsTable interface {
	Init()
	Insert(adID string, refererURL string, destinationURL string, originIP string, linkGeneratedOn time.Time, linkTag string, sessionID string) (sql.Result, error)
}

type AdClickThroughsSQLLite struct {
	db     *sql.DB
	insert *sql.Stmt
}

func (c *AdClickThroughsSQLLite) Init() {

	sql := `CREATE TABLE IF NOT EXISTS "AdClickThroughs" ("id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE, "created" DATETIME DEFAULT CURRENT_TIMESTAMP, "adID" VARCHAR, "refererURL" VARCHAR, "destinationURL" VARCHAR, "originIP" VARCHAR, "linkGeneratedOn" DATETIME, "linkTag" VARCHAR, "sessionID" VARCHAR);`
	c.db.Exec(sql)

	ins, err := c.db.Prepare(`INSERT INTO AdClickThroughs( adID, refererURL, destinationURL, originIP, linkGeneratedOn, linkTag, sessionID) values(?,?,?,?,?,?,?)`)
	checkErr(err)
	c.insert = ins

}

func (c *AdClickThroughsSQLLite) Insert(adID string, refererURL string, destinationURL string, originIP string, linkGeneratedOn time.Time, linkTag string, sessionID string) (sql.Result, error) {

	res, err := c.insert.Exec(adID, refererURL, destinationURL, originIP, linkGeneratedOn, linkTag, sessionID)
	checkErr(err)

	return res, err
}
