// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 
// Author: yumcoder (omid.jn@gmail.com)
//
package mysql

import (
	"database/sql"
	"encoding/json"
	"log"
	"testing"

	// mysql pkg
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbhostsip  = "127.0.0.1:4000"
	dbusername = "root"
	dbpassword = ""
	dbname     = "samp_db"
	dbcharset  = "utf8"
)

func getConnString() string {
	return dbusername + ":" + dbpassword + "@tcp(" + dbhostsip + ")/" + dbname + "?charset=" + dbcharset
}

func getJSON(sqlString string) (string, error) {
	sqlConnString := getConnString()
	db, err := sql.Open("mysql", sqlConnString)
	if err != nil {
		return "", err
	}

	defer db.Close()
	rows, err := db.Query(sqlString)
	if err != nil {
		return "", err
	}

	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func Test_Mysql_toJson(t *testing.T) {
	sql := "SHOW INDEX from person;"
	rows, err := getJSON(sql)
	log.Println(rows)
	log.Println(err)
}
