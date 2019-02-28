// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Author: yumcoder (omid.jn@gmail.com)
//
package mysql

/*
https://linode.com/docs/databases/mysql/how-to-optimize-mysql-performance-using-mysqltuner/
https://proxysql.com/

*/

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

/*

+--------------+---------+------+-----+---------+-------+
| Field        | RuleType    | Null | Key | Default | Extra |
+--------------+---------+------+-----+---------+-------+
| number       | int(11) | NO   | PRI | NULL    |       |
| squareNumber | int(11) | NO   |     | NULL    |       |
+--------------+---------+------+-----+---------+-------+

CREATE TABLE squareNum2
(
   number int(11) PRIMARY KEY,
   squareNumber int(11)
) engine=NDBCLUSTER;


CREATE TABLE k1 (
    id INT NOT NULL PRIMARY KEY,
    name VARCHAR(20)
)
engine=NDBCLUSTER
PARTITION BY KEY()
PARTITIONS 2;

*/

func Test_Mysql(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(10.20.30.200:6033)/mydb")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO squareNum2 VALUES( ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT squareNumber FROM squareNum2 WHERE number = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	// Insert square numbers for 0-24 in the database
	for i := 1; i < 25; i++ {
		_, err = stmtIns.Exec(i, i*i) // Insert tuples (i, i^2)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	var squareNum int // we "scan" the result in here

	// Query the square-number of 13
	err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 13 is: %d", squareNum)

	// Query another number.. 1 maybe?
	err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 1 is: %d", squareNum)
}

type SquareNum struct {
	number int32
	squareNumber int32
}

func Test_Mysql_injection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(172.17.0.2:3306)/yumdcoder")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	id := "1; delete from squareNum2 where number=1;"

	var query = `SELECT squareNumber FROM squareNum2 WHERE number = ?`
	rows, err := db.Query(query, id)
	if err != nil {
		t.Fatal(err)
	}

	defer rows.Close()

	if rows.Next() {
		result := &SquareNum{}
		if err = rows.Scan(&result.squareNumber); err != nil {
			t.Fatal(err)
		}
		t.Log(result)
		return
	}

	if err = rows.Err(); err != nil {
		t.Fatal(err)
	}

	t.Log("nil...")
}

func Test_Mysql_SingleRow(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(172.17.0.2:3306)/yumd")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	id := "9891256212030"

	var query = `select 1 from user where phone = ?`
	row := db.QueryRow(query, id)
	result := int8(0)
	err = row.Scan(&result)
	switch err {
	case sql.ErrNoRows:
		t.Log(false)
	case nil:
		t.Log(result == 1)
	default:
		t.Log(err)
	}

	t.Log("nil...")
}