package insights

import (
	"database/sql"
)

type Impressions struct {
	db     *sql.DB
	insert *sql.Stmt
}

func (i *Impressions) Init(db *sql.DB) {
	i.db = db

	sql := `CREATE TABLE IF NOT EXISTS "impressions" (
"id" INTEGER PRIMARY KEY  AUTOINCREMENT  NOT NULL  UNIQUE , 
"dateTime" DATETIME DEFAULT CURRENT_TIME,
"userId" VARCHAR,
"originURL" VARCHAR, 
"ip" VARCHAR
);`
	i.db.Exec(sql)

	insertImp, err := db.Prepare("INSERT INTO impressions(userId, originURL, ip) values(?,?,?)")
	checkErr(err)
	i.insert = insertImp

}

func (i *Impressions) Insert(userId string, originURL string, ip string) (sql.Result, error) {
	res, err := i.insert.Exec(userId, originURL, ip)
	checkErr(err)

	return res, err
}
