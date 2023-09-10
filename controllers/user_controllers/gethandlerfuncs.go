package usercontrollers

import (
	"Project_2/database"
	// usermiddlewares "Project_2/middlewares/user_middlewares"
	"Project_2/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLogin(context *gin.Context) {
	context.HTML(http.StatusOK, "login_admin_user.html", gin.H{})
}

func GetSignup(context *gin.Context) {
	context.HTML(http.StatusOK, "signup_users.html", gin.H{})
}

func Logoutuser(context *gin.Context) {
	context.SetCookie("jwt", "", -1, "/", "localhost", true, true)
	context.SetCookie("username", "", -1, "/", "localhost", true, true)
	fmt.Println("saodgtdysufgsfiuasssfjksghjgfjuyefgjdtyufvwejumg")
	context.Redirect(http.StatusPermanentRedirect, "/login")
}

func Deletetodo(context *gin.Context) {
	username := context.Param("username")
	todo := context.Param("name")
	fmt.Print(username, todo)
	query := "delete from todos where username = ? AND name = ?"
	database.DB.Exec(query, username, todo)
	GetHome(context)
}

func gettodos(username string) []models.Todos {
	var gettodos []models.Todos
	query := "select * from todos where username = ?"
	database.DB.Raw(query, username).Scan(&gettodos)
	return gettodos
}

func Editpage(context *gin.Context) {
	var singletodo models.Todos
	username := context.Param("username")
	name := context.Param("name")
	query := "select * from todos where username = ? AND name = ?"
	database.DB.Raw(query, username, name).Scan(&singletodo)
	context.HTML(http.StatusOK, "edittodos.html", gin.H{"todo": singletodo})
}

func Edittodo(context *gin.Context) {
	username := context.Param("username")
	name := context.Param("name")
	newtodo := context.PostForm("editedtodo")
	newdate := context.PostForm("editeddate")
	fmt.Print(username, name, newtodo, newdate)
	query := "update todos set name = ?, date = ? where username = ? AND name = ?"
	res := database.DB.Exec(query, newtodo, newdate, username, name)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	GetHome(context)
}

func Addtodos(context *gin.Context) {
	username := context.Param("username")
	addtodo := context.PostForm("todoName")
	adddate := context.PostForm("todoTime")
	fmt.Print(username, adddate, addtodo)
	query := "insert into todos (username,name,date) values(?,?,?)"
	database.DB.Exec(query, username, addtodo, adddate)
	GetHome(context)
}

func GetHome(context *gin.Context) {
	//recover fuunction to do whenever an error occures during the loading of the home page
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("this defer called")
			context.HTML(http.StatusOK, "login_admin_user.html", gin.H{"message": "an error has occured please login again"})
			return
		}
	}()
	username, _ := context.Request.Cookie("username")
	oguser := username.Value
	geetttodos := gettodos(oguser)
	fmt.Println("home called")
	context.HTML(http.StatusOK, "todolist.html", gin.H{"todos": geetttodos, "user": oguser})
}
