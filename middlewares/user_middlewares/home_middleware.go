package usermiddlewares

import (
	"Project_2/database"
	"Project_2/models"
	"Project_2/token"
	"fmt"
	"net/http"

	// "time"

	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SecureHome() gin.HandlerFunc {
	return func(context *gin.Context) {
		//checking if token is alresdy present or not
		tokenjwt, err := context.Request.Cookie("jwt")
		fmt.Println("first")
		if err != nil {
			fmt.Println(err)
			context.Redirect(http.StatusPermanentRedirect, "/login")
			context.Abort()
			return
		}
		fmt.Println("second")
		if tokenjwt == nil || tokenjwt.Value == "" {
			context.Redirect(http.StatusPermanentRedirect, "/login") //here should be login but i changed
			context.Abort()
			return
		}
		fmt.Println("token is below")
		fmt.Println("this is the token ", tokenjwt.Value)
		//validating the token using its credentials and also cahecks if thr expiration time of thr user has expired
		if token.ValidateToken(tokenjwt.Value) {
			context.Next()
			return
		} else {
			context.Set("message", "Session invalid")
			context.Redirect(http.StatusPermanentRedirect, "/login") //here should be login but i changed
			context.Abort()
		}
	}
}

//for validating the user
//this was actually two seperate middlewares but i had to combine the code in the securehome() middleware
//code in this validateuser() because the cookies that i am setting tin the validateuser() cannot be accesseed
//from the securehome() for the first time user logs in , but when i try to log in for the second time it works perfectly fine
func Validateuser() gin.HandlerFunc {
	return func(context *gin.Context) {
		var user models.Users
		username := context.PostForm("username")
		password := context.PostForm("password")
		database.DB.Where("username = ?", username).First(&user)
		//checking with hashed password
		pass:=[]byte(password)
		err:=bcrypt.CompareHashAndPassword([]byte(user.Password),pass)
		if user.Username == "" || err!=nil{
			context.HTML(http.StatusOK, "login_admin_user.html", gin.H{"message": "user not found"})
			context.Abort()
			return
		}
		fmt.Println("this has occured...")
		//generating the jwt authentication tocken for the given username and password
		jwttoken := token.Generatejwt(username, password)
		fmt.Println(jwttoken)
		context.SetCookie("jwt", jwttoken, 3600, "/", "localhost", true, true)
		context.SetCookie("username", username, 3600, "/", "localhost", true, true)
		fmt.Println("second")
		if jwttoken == "" {
			context.Redirect(http.StatusPermanentRedirect, "/login")
			context.Abort()
			return
		}
		fmt.Println("token is below")
		fmt.Println("this is the token ", jwttoken)
		if token.ValidateToken(jwttoken) {
			geetttodos := gettodos(username)
			fmt.Println("todooooooooooooooooooooooooooooooooooooooo")
			context.HTML(http.StatusOK, "todolist.html", gin.H{"todos": geetttodos, "user": username})
			return
		} else {
			context.Set("message", "Session invalid")
			context.Redirect(http.StatusPermanentRedirect, "/login") 
			context.Abort()
		}

	}
}

// for clearing the cache when logging out
func Clearcache() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		context.Header("Pragma", "no-cache")
		context.Header("Expires", "0")
		context.Header("Vary", "Cookie")
		context.Next()
	}
}

//for getting the todos for that specific user who's logged in
func gettodos(username string) []models.Todos {
	var gettodos []models.Todos
	query := "select * from todos where username = ?"
	database.DB.Raw(query, username).Scan(&gettodos)
	return gettodos
}
