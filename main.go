package main

import (
	admincontrollers "Project_2/controllers/admin_controllers"
	usercontrollers "Project_2/controllers/user_controllers"
	"Project_2/database"
	"fmt"
	"os"

	adminmiddlewares "Project_2/middlewares/admin_middlewares"
	usermiddlewares "Project_2/middlewares/user_middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//loading env files into main
	err := godotenv.Load("globals.env")
	
	if err != nil {
		fmt.Println(err)
		panic("cannot connect to the database")
	}
	env := os.Getenv("DATABASE_ADDR")
	//connecting to database and migrating the models to tables
	database.Connect(env)
	database.Migrator()
	database.Migrateadmnm()
	database.Migratetodos()
	router := InitGin()
	router.Run(":2031")
}

//intitializing the router and setting routes
func InitGin() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")//to render html files
	router.GET("/", usermiddlewares.SecureHome(), usercontrollers.GetHome)//home for users
	router.GET("/login", usermiddlewares.Clearcache(), usercontrollers.GetLogin)//called when the user is not logged in
	router.GET("/signup", usercontrollers.GetSignup)//called when the user doesnt have an account
	router.POST("/submitsignup", usercontrollers.PostSignup)//called to create an account
	router.POST("/gethome", usermiddlewares.Validateuser())//called after pressing the login button
	router.GET("/loginadmin", admincontrollers.Getadminlogin)//calling admin login page
	router.POST("/getadminhome", adminmiddlewares.Validateadmin(), admincontrollers.Getadminhome)//called when the admin presses the login button
	router.GET("/adminhome", adminmiddlewares.SecureadminHome(), admincontrollers.Getadminhome)//home admin
	router.POST("/search", admincontrollers.Getsearches)//to get the search results about users in admin home
	router.GET("/delete/:username", admincontrollers.Deleteuser)//called for deleting a user
	router.GET("/edit/:username", admincontrollers.Edituser)//calles the edit page of the users
	router.POST("/editusers/:edituser", admincontrollers.Postedit)//posting the edit on user
	router.GET("/logout", usercontrollers.Logoutuser)
	router.GET("/logoutadmin", admincontrollers.Logoutadmin)
	router.GET("/deletetodos/:username/:name", usercontrollers.Deletetodo)//delete todo
	router.GET("/edittodos/:username/:name", usercontrollers.Editpage)// loading the edit page for editing todo
	router.POST("/editedtodos/:username/:name", usercontrollers.Edittodo)//edit todo
	router.POST("/add-todo/:username", usercontrollers.Addtodos)//add todo
	return router
}
