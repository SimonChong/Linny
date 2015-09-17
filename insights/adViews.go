package insights

import (
	"database/sql"
)

type AdViews struct {
	db     *sql.DB
	insert *sql.Stmt
}

func (i *AdViews) Init(db *sql.DB) {
	i.db = db

	sql := `CREATE TABLE IF NOT EXISTS "AdViews" ( "id" INTEGER PRIMARY KEY  AUTOINCREMENT  NOT NULL  UNIQUE , "created" DATETIME DEFAULT CURRENT_TIME, "userId" VARCHAR, "originURL" VARCHAR, "ip" VARCHAR
);`
	i.db.Exec(sql)

	insertImp, err := db.Prepare("INSERT INTO AdViews(userId, originURL, ip) values(?,?,?)")
	checkErr(err)
	i.insert = insertImp

}

func (i *AdViews) Insert(userId string, originURL string, ip string) (sql.Result, error) {
	res, err := i.insert.Exec(userId, originURL, ip)
	checkErr(err)

	return res, err
}
