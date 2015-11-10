package insights

import (
	"database/sql"
	"time"
)

type AdViewsTable interface {
	Init()
	Insert(adID string, refererURL string, originIP string, contentGeneratedOn time.Time, sessionID string) (sql.Result, error)
}

type AdViewsSQLLite struct {
	db     *sql.DB
	insert *sql.Stmt
}

func (c *AdViewsSQLLite) Init() {

	sql := `CREATE TABLE IF NOT EXISTS "AdViews" ("id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE, "created" DATETIME DEFAULT CURRENT_TIMESTAMP, "adID" VARCHAR, "refererURL" VARCHAR, "originIP" VARCHAR, "contentGeneratedOn" DATETIME, "sessionID" VARCHAR);`
	c.db.Exec(sql)

	ins, err := c.db.Prepare(`INSERT INTO AdViews(adID, refererURL, originIP, contentGeneratedOn, sessionID) values(?,?,?,?,?)`)
	checkErr(err)
	c.insert = ins

}

func (i *AdViewsSQLLite) Insert(adID string, refererURL string, originIP string, contentGeneratedOn time.Time, sessionID string) (sql.Result, error) {
	res, err := i.insert.Exec(adID, refererURL, originIP, contentGeneratedOn, sessionID)
	checkErr(err)

	return res, err
}
