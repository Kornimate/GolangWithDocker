package main

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func main() {
	DatabaseConnection()

	router := gin.Default()

	router.GET("/users", GetUsersAPI)
	router.GET("/user/:id")
	router.POST("/adduser", PostUserAPI)

	router.Run("localhost:8080")

	// users, err := GetUsers()

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(users)
	// }

	// user, err := GetUserById(3)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(user)
	// }

	// err = AddUser(UserData{
	// 	name: fmt.Sprintf("Kornidesz %d", rand.Int31n(10000)),
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// users, err = GetUsers()

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(users)
	// }
}

func GetUsersAPI(context *gin.Context) {

	users, _ := GetUsers()

	context.IndentedJSON(http.StatusOK, users)
}

func PostUserAPI(context *gin.Context) {
	var usr UserData

	if err := context.BindJSON(&usr); err != nil {
		return
	}

	AddUser(usr)

	context.IndentedJSON(http.StatusCreated, usr)
}
