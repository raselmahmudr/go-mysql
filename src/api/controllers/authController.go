
package controllers

import (
	"net/http"
)

func Login(res http.ResponseWriter,  req *http.Request)  {

}



func Register(res http.ResponseWriter,  req *http.Request)  {
	// db := database.DB_INIT()
	//
	// var body map[string]interface{}
	// json.NewDecoder(req.Body).Decode(&body)
	//
	// var newUser models.User
	//
	// if body["email"] != nil && body["email"] != "" {
	// 	newUser.Email = body["email"].(string)
	// }
	// if body["username"] != nil && body["username"] != "" {
	// 	newUser.Username = body["username"].(string)
	// }
	// if body["password"] != nil && body["password"] != "" {
	// 	hashedPass, err := bcrypt.GenerateFromPassword([]byte(body["password"].(string)), 8)
	// 	if err != nil {
	// 		fmt.Println(err)
 	// 	}
 	// 	newUser.Password = string(hashedPass)
	// }
	//
	//
	// ts := db.Create(&newUser)
	//
	// r, _ := ts.Rows()
	// if r.Next() {
	// 	var id string
	// 	var username string
	// 	// fmt.Println(r.Columns())
	// 	r.Scan(&id, &username)
	// 	fmt.Println(id, username)
	// }

	//
	// db.Create({ models.User{
	// 	Username:  "",
	// 	Email:     "",
	// 	Password:  ""
	// } })
}


// func GetUsers(res http.ResponseWriter,  req *http.Request) {
// 	res.Header().Set("Content-Type", "application/json")
// 	db := database.DB_INIT()
// 	var users []models.User
// 	db.Find(&users)
//
// 	json.NewEncoder(res).Encode(users)
// }