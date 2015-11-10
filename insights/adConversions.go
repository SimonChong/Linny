package insights

import (
	"database/sql"
	"time"
)

type AdConversionsTable interface {
	Init()
	Insert(adID string, refererURL string, originIP string, jsGeneratedOn time.Time, conversionTag string, sessionID string) (sql.Result, error)
}

type AdConversionsSQLLite struct {
	db     *sql.DB
	insert *sql.Stmt
}

func (c *AdConversionsSQLLite) Init() {

	sql := `CREATE TABLE IF NOT EXISTS "AdConversions" ("id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE, "created" DATETIME DEFAULT CURRENT_TIMESTAMP, "adID" VARCHAR, "refererURL" VARCHAR, "originIP" VARCHAR, "jsGeneratedOn" DATETIME, "conversionTag" VARCHAR, "sessionID" VARCHAR);`
	c.db.Exec(sql)

	ins, err := c.db.Prepare(`INSERT INTO AdConversions(adID, refererURL, originIP, jsGeneratedOn, conversionTag, sessionID) values(?,?,?,?,?,?)`)
	checkErr(err)
	c.insert = ins
}

func (c *AdConversionsSQLLite) Insert(adID string, refererURL string, originIP string, jsGeneratedOn time.Time, conversionTag string, sessionID string) (sql.Result, error) {

	res, err := c.insert.Exec(adID, refererURL, originIP, conversionTag, sessionID)
	checkErr(err)

	return res, err
}
