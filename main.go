package main

import (
  "net/http"
  "strconv"
  "errors"
  "github.com/gin-gonic/gin"
)

type User struct {
	  Id  int    `json:"id"`
	  Name string `json:"name"`
	  Email string `json:"email"`
	  Completed bool `json:"completed"`
}

var users = []User{
	  {Id: 1, Name: "John", Email: "John@mail.com", Completed: false},
	  {Id: 2, Name: "Smith", Email: "Smith@mail.com", Completed: false},
	  {Id: 3, Name: "Jane", Email: "Jane@mail.com", Completed: false},
}

func getUsers(context *gin.Context){
	context.IndentedJSON(http.StatusOK, users)
}

func addUser(context *gin.Context){
	var newUser User

	if err := context.ShouldBindJSON(&newUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users = append(users, newUser)
	context.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func getUserById(id int) (*User, error){
	for i, u := range users {
		if u.Id == id {
			return &users[i], nil
		}
	}
	return nil, errors.New("User not found")
}

func getUser(context *gin.Context){
	
	id := context.Param("id")
	idnumber, _ := strconv.Atoi(id)
	user, err := getUserById(idnumber)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, user)
}

func main() {
	  router := gin.Default()

	  router.GET("/users", getUsers)
	  router.POST("/users", addUser)	
	  router.GET("/users/:id", getUser)
	
  router.Run(":8080")
}