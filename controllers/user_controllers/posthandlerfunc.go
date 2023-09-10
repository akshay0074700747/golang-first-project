package usercontrollers

import (
	"Project_2/database"
	"Project_2/models"
	"fmt"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
)

//inserting user credentials to the database
func PostSignup(context *gin.Context) {
	var user models.Users
	username := context.PostForm("username")
	email := context.PostForm("email")
	password := context.PostForm("password")
	confirmpassword := context.PostForm("confirmPassword")
	fmt.Println(password)
	fmt.Println(confirmpassword)
	if password != confirmpassword {
		context.HTML(http.StatusOK, "signup_users.html", gin.H{"message": "both the passwords should be the same"})
		return
	}
	database.DB.Where("username = ?", username).First(&user)
	if user.Username != "" {
		context.HTML(http.StatusOK, "signup_users.html", gin.H{"message": "user already exists"})
		return
	}
	//hashing the password
	pass:=[]byte(password)
	hashedpass,err:=bcrypt.GenerateFromPassword(pass,bcrypt.DefaultCost)
	if err!=nil {
		context.HTML(http.StatusOK, "signup_users.html", gin.H{"message": "error... try again later"})
		return
	}

	//creating the specific row to be inserted to the data base
	user = models.Users{
		Username: username,
		Password: string(hashedpass),
		Email:    email,
	}
	//inserting the user to the data base && checks for any errors occured during the transaction
	record := database.DB.Create(&user)
	if record.Error != nil {
		fmt.Println("Error occured while inseerting ton the data base")
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.HTML(http.StatusOK, "login_admin_user.html", gin.H{})
}
