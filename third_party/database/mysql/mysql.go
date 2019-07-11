package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var stop = make(chan string)

func main() {
	// open
	db, err := sql.Open("mysql", "root:Secret123.@tcp(localhost:3306)/test_db")
	handleErr(err, "sql.Open")
	defer func(db *sql.DB) { handleErr(db.Close(), "db.Close") }(db)

	// stats
	DbStatsListen(db)

	// ping
	handleErr(db.Ping(), "db.ping")

	// exec
	_, _ = db.Exec("use test_db")

	// conn pool
	//_, _ = db.Exec("LOCK TABLES user read ")

	// insert
	execRs, err := db.Exec("insert into user (uname, passwd, locked) values (?,?,?)", "alice", "pass", 1)
	handleErr(err, "db.Exec")
	lastId, err := execRs.LastInsertId()
	handleErr(err, "rs.LastInsertId()")
	afRows, err := execRs.RowsAffected()
	handleErr(err, "rs.rowsAffected()")
	fmt.Printf("lastId:(%d), affectRows:(%d)\n", lastId, afRows)

	// sql
	rows, err := db.Query("select uid, uname from user")
	handleErr(err, "db.Query")
	handleRows(rows)

	// prepare(multiple times)
	stmt, err := db.Prepare("select uid, uname, locked from user where uid in (?)")
	handleErr(err, "db.Prepare")
	defer stmt.Close()
	rows, err = stmt.Query(3)
	handleErr(err, "db.Query")
	handleRows2(rows)

	u := new(User)
	err = stmt.QueryRow(5).Scan(&u.id, &u.name, &u.locked)
	handleErr(err, "stmt.QueryRow")
	fmt.Printf("%+v\n", *u)

	// single-row query
	//u = new(User)
	err = db.QueryRow("select uid, uname, locked from user where uid = ?;", 10).Scan(&u.id, &u.name, &u.locked)
	handleErr(err, "db.QueryRow")
	fmt.Printf("%+v\n", *u)

	// prepare insert
	stmt, err = db.Prepare("insert into user (uid, uname, passwd, locked) values (null, ?, ?, ?);")
	handleErr(err, "db.Prepare insert")
	res, err := stmt.Exec("gaga", "gapas123", 0)
	handleErr(err, "stmt.Exec")
	lastId, err = res.LastInsertId()
	handleErr(err, "lastId")
	rowCnt, err := res.RowsAffected()
	handleErr(err, "rowCnt")
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	// unknown columns
	rows, err = db.Query("SHOW PROCESSLIST")
	handleErr(err, "db.Query show processlist")
	defer rows.Close()
	cols, err := rows.Columns()
	handleErr(err, "rows.Columns")
	vals := make([]interface{}, len(cols))

	for i, _ := range vals {
		vals[i] = new(sql.RawBytes)
	}

	var rowList [][]interface{}

	for rows.Next() {
		if err := rows.Scan(vals...); err != nil {
			log.Printf("scan err %s \n", err)
		}

		for _, v := range vals {

			switch v.(type) {
			case *sql.RawBytes:
				fmt.Printf("%s", (v))
			default:
				fmt.Printf("%s", v)
			}
		}
		fmt.Printf("\n")

		rowList = append(rowList, vals)
	}

	fmt.Printf("%+v", rowList)


	// stop
	<-stop
}

type User struct {
	id     int
	name   string
	locked bool
}

// 处理行记录
func handleRows2(rows *sql.Rows) {
	defer rows.Close()

	var users []User

	for rows.Next() {
		u := new(User)
		err := rows.Scan(&u.id, &u.name, &u.locked)
		handleErr(err, "rows.Scan")
		users = append(users, *u)
	}

	if err := rows.Err(); err != nil {
		// handle the error here
	}

	fmt.Printf("%+v\n", users)
}

// 处理行记录
func handleRows(rows *sql.Rows) {
	defer rows.Close()

	var users []User

	for rows.Next() {
		u := new(User)
		err := rows.Scan(&u.id, &u.name)
		handleErr(err, "rows.Scan")
		users = append(users, *u)
	}

	fmt.Printf("%+v\n", users)
}

// 结果集打印
func PdbResult(rs interface{}) {
	fmt.Printf("%+v\n", rs)
}

// db状态
func DbStatsListen(db *sql.DB) {
	go func() {
		tch := time.Tick(10 * time.Second)
		for {
			select {
			case <-tch:
				PdbStats(db)
			}
		}
	}()
}

func PdbStats(db *sql.DB) {
	fmt.Printf("%#v\n", *db)
	fmt.Printf("%#v\n\n", db.Stats())
}

// db错误处理
func handleErr(err error, mark string) {
	if err != nil {
		log.Fatalln(mark, err)
	}
}
