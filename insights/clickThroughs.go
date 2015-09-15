package insights

import (
	"database/sql"
	"time"
)

type ClickThroughs struct {
	db     *sql.DB
	insert *sql.Stmt
}

func (c *ClickThroughs) Init(db *sql.DB) {
	c.db = db

	sql := `CREATE TABLE IF NOT EXISTS "clickthroughs" (
"id" INTEGER PRIMARY KEY  AUTOINCREMENT  NOT NULL  UNIQUE ,
"dateTime" DATETIME DEFAULT CURRENT_TIMESTAMP,
"userID" VARCHAR,
"refererURL" VARCHAR,
"destinationURL" VARCHAR,
"originIP" VARCHAR,
"dateTimeGen" DATETIME,
"linkTag" VARCHAR
);`
	c.db.Exec(sql)

	ins, err := db.Prepare(`INSERT INTO clickthroughs(
userID,
refererURL,
destinationURL,
originIP,
dateTimeGen,
linkTag
) values(?,?,?,?,?,?)`)
	checkErr(err)
	c.insert = ins

}

func (c *ClickThroughs) Insert(userID string, refererURL string,
	destinationURL string, originIP string, dateTimeGen time.Time,
	linkTag string) (sql.Result, error) {

	res, err := c.insert.Exec(userID, refererURL, destinationURL, originIP, dateTimeGen, linkTag)
	checkErr(err)

	return res, err
}
