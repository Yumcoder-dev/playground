// Copyright 2018-present The Yumcoder Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 
// Author: yumcoder (omid.jn@gmail.com)
//
package mysql

import (
	"database/sql"
	"testing"
)

type User struct {
	Id                int32
	UserType          int8
	AccessHash        int64
	FirstName         string
	LastName          string
	Username          string
	Phone             string
	CountryCode       string
	Verified          int8
	About             string
	State             int32
	IsBot             int8
	AccountDaysTtl    int32
	Photos            string
	Min               int8
	Restricted        int8
	RestrictionReason string
	Deleted           int8
	DeleteReason      string
	CreatedAt         string
	UpdatedAt         string
}

type UserStorage struct{
	*sql.DB
}

func (db *UserStorage) SelectById1(id int32) (*User, error) {
	var query = `select id, user_type, access_hash, first_name, last_name, username, phone, country_code, verified, about, state, is_bot, account_days_ttl, photos, min, restricted, restriction_reason, deleted, delete_reason from user where id=?`
	row := db.QueryRow(query, id)

	result := &User{}
	err := row.Scan(&result.Id, &result.UserType, &result.AccessHash, &result.FirstName, &result.LastName, &result.Username, &result.Phone, &result.CountryCode, &result.Verified, &result.About, &result.State, &result.IsBot, &result.AccountDaysTtl, &result.Photos, &result.Min, &result.Restricted, &result.RestrictionReason, &result.Deleted, &result.DeleteReason)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return result, nil
	default:
		return nil, err
	}
}

func (db *UserStorage) SelectById2(id int32) (User, error) {
	var query = `select id, user_type, access_hash, first_name, last_name, username, phone, country_code, verified, about, state, is_bot, account_days_ttl, photos, min, restricted, restriction_reason, deleted, delete_reason from user where id=?`
	row := db.QueryRow(query, id)

	result := User{}
	err := row.Scan(&result.Id, &result.UserType, &result.AccessHash, &result.FirstName, &result.LastName, &result.Username, &result.Phone, &result.CountryCode, &result.Verified, &result.About, &result.State, &result.IsBot, &result.AccountDaysTtl, &result.Photos, &result.Min, &result.Restricted, &result.RestrictionReason, &result.Deleted, &result.DeleteReason)

	switch err {
	case sql.ErrNoRows:
		return result, nil
	case nil:
		return result, nil
	default:
		return result, err
	}
}

func Benchmark_sql(b *testing.B) {
	db, err := sql.Open("mysql", "root:@tcp(172.17.0.2:3306)/yumdcoder")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	us := UserStorage{db}
	var u *User
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		u, err = us.SelectById1(int32(i))
	}
	_ = u
}