
package controllers

import (
	"encoding/json"
	"fmt"
	"hello/database"
	"hello/models"
	"net/http"
	"strings"
)

func Login(res http.ResponseWriter,  req *http.Request)  {
	fmt.Println("hi")
	err := database.DB()
	if err != nil {
		fmt.Println("Connection fail")
	} else {
		fmt.Println("Connected")
	}
}



func Register(res http.ResponseWriter,  req *http.Request)  {
	
	db := database.DB()
	
	var body map[string]string
	json.NewDecoder(req.Body).Decode(&body)
	
	u := models.User{}
	g, _ := json.Marshal(u)
	userObj := map[string]interface{}{}
	
	_ = json.Unmarshal(g, &userObj)
	
	sql := "INSERT INTO "
	sql += "USERS ("
	fields := ""
	values := ""
	
	
	// make dynamic sql statement ........
	for key, _ := range userObj{
		if key != "id" {
			if body[key] == "" {
				json.NewEncoder(res).Encode(map[string]string{
					"message": key + " Must be required",
				})
				return
			} else {
				fields = fields + ", " + key
				values = values + ",'" + body[key] + "'"
			}
		}
	}
	
	sql = sql + strings.Replace(fields, ",", "", 1) + ") VALUES(" + strings.Replace(values, ",", "", 1) + " )"
	
	// sql = "INSERT INTO USERS (username, email) VALUES('rasel mahmud', 'rasel@gmail.com' )"
	
	fmt.Println(sql)
	
	rows, err := db.Query(sql)
	fmt.Println(rows, err)
	
	
}


func GetUsers(res http.ResponseWriter,  req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	db := database.DB()
	user := models.User{}
	allUsers := user.Find(db)
	
	json.NewEncoder(res).Encode(allUsers)
}