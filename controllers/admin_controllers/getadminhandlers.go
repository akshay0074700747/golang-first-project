package admincontrollers

import (
	"Project_2/database"
	"Project_2/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Getadminlogin(context *gin.Context) {
	context.HTML(http.StatusOK, "login_admin.html", gin.H{})
}

var users []models.Users

func Getadminhome(context *gin.Context) {
	getusers := getusers()
	fmt.Println("starting...")
	for _, v := range users {
		fmt.Println(v.Username)
	}
	context.HTML(http.StatusOK, "home_admin.html", gin.H{"users": getusers})
}

var searchusers []models.Users

func Getsearches(context *gin.Context) {
	getusers := getusers()
	search := context.PostForm("query")
	if search == "" {
		context.HTML(http.StatusOK, "home_admin.html", gin.H{"users": users, "search": searchusers})
		return
	}
	query := "SELECT * FROM users WHERE username LIKE ?"
	database.DB.Raw(query, "%"+search+"%").Scan(&searchusers)
	context.HTML(http.StatusOK, "home_admin.html", gin.H{"users": getusers, "search": searchusers})
}

func Deleteuser(context *gin.Context) {
	username := context.Param("username")
	fmt.Println(username)
	query := "DELETE FROM users WHERE username = ?"
	database.DB.Exec(query, username)
	getuser := getusers()
	context.HTML(http.StatusOK, "home_admin.html", gin.H{"users": getuser})
}

func Edituser(context *gin.Context) {
	var sigleuser models.Users
	username := context.Param("username")
	query := "select * from users where username = ?"
	database.DB.Raw(query, username).Scan(&sigleuser)
	context.HTML(http.StatusOK, "editusers.html", gin.H{"user": sigleuser})
}

func Postedit(context *gin.Context) {
	editusers := context.Param("edituser")
	username := context.PostForm("username")
	email := context.PostForm("email")
	query := "UPDATE users SET username = ?, email = ? WHERE username = ?"
	res := database.DB.Exec(query, username, email, editusers)
	if res.Error != nil {
		fmt.Print(res.Error)
	}
	getuser := getusers() 
	context.HTML(http.StatusOK, "home_admin.html", gin.H{"users": getuser})
}

func getusers() []models.Users {
	var getusers []models.Users
	getquery := "select * from users"
	database.DB.Raw(getquery).Scan(&getusers)
	return getusers
}

func Logoutadmin(context *gin.Context)  {
	context.SetCookie("adminjwt", "", -1, "/", "localhost", true, true)
	context.Redirect(http.StatusPermanentRedirect,"/login")
}
