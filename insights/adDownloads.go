package insights

import "database/sql"

type AdDownloadsTable interface {
	Init()
	Insert(adID string, filePath string, refererURL string, originIP string, sessionID string) (sql.Result, error)
}

type AdDownloadsSQLLite struct {
	db     *sql.DB
	insert *sql.Stmt
}

func (c *AdDownloadsSQLLite) Init() {

	sql := `CREATE TABLE IF NOT EXISTS "AdDownloads" ("id" INTEGER PRIMARY KEY AUTOINCREMENT  NOT NULL UNIQUE, "created" DATETIME DEFAULT CURRENT_TIMESTAMP, "adID" VARCHAR, "filePath" VARCHAR, "refererURL" VARCHAR, "originIP" VARCHAR, "sessionID" VARCHAR);`
	c.db.Exec(sql)

	ins, err := c.db.Prepare(`INSERT INTO AdDownloads( adID, filePath, refererURL, originIP, sessionID ) values(?,?,?,?,?)`)
	checkErr(err)
	c.insert = ins
}

func (c *AdDownloadsSQLLite) Insert(adID string, filePath string, refererURL string, originIP string, sessionID string) (sql.Result, error) {

	res, err := c.insert.Exec(adID, filePath, refererURL, originIP, sessionID)
	checkErr(err)

	return res, err
}
