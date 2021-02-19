package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	Id uint8 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}


func (U *User) Find(db *sql.DB) []User {
	fmt.Println(U)
	defer db.Close()
	res, _ := db.Query("SELECT * FROM users")
	
	var users []User
	for res.Next() {
		u := User{}
		s := reflect.ValueOf(&u).Elem()
		numCols := s.NumField()
	
		columns := make([]interface{}, numCols)
		for i:=0; i<numCols; i++ {
			columns[i] = s.Field(i).Addr().Interface()
		}
	
		res.Scan(columns...)
		users = append(users, u)
	}
	return  users
}

func (U * User) FindById (db *sql.DB)  {

}



func (U *User) Sync(db *sql.DB)  {
	// fmt.Println("Sync User Table")
	
	sql := "CREATE TABLE USERS("
	// sql := "CREATE TABLE IF NOT EXISTS USERS("
	
	
	var u User
	sql = sql + ListFields(&u) + ")"
	fmt.Println(sql)
	r, e := db.Query(sql)
	fmt.Println(r, e)
	
}
func ListFields(a interface{}) string{
	// sql := "CREATE TABLE IF NOT EXISTS USERS("
	g:=""
	v := reflect.ValueOf(a).Elem()
	for j := 0; j < v.NumField(); j++ {
		f := v.Field(j)
		n := v.Type().Field(j).Name
		t := f.Type().String()

		lower := strings.ToLower(n)
		
		g = g + lower + " " + checkType(t, lower) + ","

		
		// fmt.Printf("Name: %s  Basic Type or Kind: %s  Direct or Custom Type: %s\n", n, f.Kind(), t)
	}
	
	h := g[:strings.LastIndex(g, ",")]
	return h
}


func checkType(str string, fieldName string) string {
	var res string
	if str == "uint8" &&  fieldName == "id" {
		res = "INT PRIMARY KEY NOT NULL AUTO_INCREMENT"
	}
	if str == "string" &&  fieldName == "email" {
		res = "Varchar(50) NOT NULL"
	}
	if str == "string" &&  fieldName == "password" {
		res = "Varchar(50) NOT NULL"
	}
	if str == "string" &&  fieldName == "username" {
		res = "Varchar(50) NOT NULL"
	}
	return res
}