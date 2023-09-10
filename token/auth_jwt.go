package token

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var jwt_token string
var Big_secret []byte

//this init function will be the first function called in this package and it is used to get values from the env file
func init()  {
	godotenv.Load("globals.env")
	jwt_token=os.Getenv("SECRET_KEY")
	Big_secret=[]byte(jwt_token)
}

//payload for creating the jwt token
type Payload struct{
	Username string
	Password string
	jwt.StandardClaims//its a predefined struct for storing the common values for all tokens such as expiresat,issuedat etc...
}

//generate the jwt for the given username and password
func Generatejwt(username string,password string) string {
	//specifies the time for which the jwt should expire and adds it to the clams payload along with the username and passwoerd
	expiresat:=time.Now().Add(48 * time.Hour)
	claims:=&Payload{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresat.Unix(),
		},
	}
	//specifies the signing method in which the jwt should be signed 
	tokenjwt:=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//specifies the secret key using which the jwt token should be signed with
	tokenstring,_:=tokenjwt.SignedString(Big_secret)
	return tokenstring
}

//for validating the jwt token coming with each request
func ValidateToken(jwt_token string) bool {
	//this code validates a JWT token by checking its signing method is hmac or not
	//jwt.parse function takes two parameter they are the jwttoken itself and 
	//and an anonymous function which return two values they are the secret key and an error
	//if this anonymous function return the secret key, then the jwt.parse function will take this 
	//return value from the anonymous function and use it to verify the token
	token, err := jwt.Parse(jwt_token, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return Big_secret, nil
    })
	//checks if the token is expired or not
	if err == nil && token.Valid {
        return true
    } else {
        return false
    }
}