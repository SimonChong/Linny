package insights

import "database/sql"

type AdConversions struct {
	db     *sql.DB
	insert *sql.Stmt
}

func (c *AdConversions) Init(db *sql.DB) {
	c.db = db

	sql := `CREATE TABLE IF NOT EXISTS "AdConversions" ("id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE, "created" DATETIME DEFAULT CURRENT_TIMESTAMP, "adID" VARCHAR, "refererURL" VARCHAR, "originIP" VARCHAR, "conversionTag" VARCHAR, "sessionID" VARCHAR);`
	c.db.Exec(sql)

	ins, err := db.Prepare(`INSERT INTO AdConversions(adID, refererURL, originIP, conversionTag, sessionID) values(?,?,?,?,?)`)
	checkErr(err)
	c.insert = ins
}

func (c *AdConversions) Insert(adID string, refererURL string, originIP string, conversionTag string, sessionID string) (sql.Result, error) {

	res, err := c.insert.Exec(adID, refererURL, originIP, conversionTag, sessionID)
	checkErr(err)

	return res, err
}
