
package common

import (
	"math/rand"
	"time"
	"fmt"
	"github.com/SlyMarbo/gmail"

	"github.com/dgrijalva/jwt-go"
	"../constants"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// A helper function to generate random string
func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Keep this two config private, it should not expose to open source
const NBSecretPassword = "A String Very Very Very Strong!!@##$!@#$"
const NBRandomPassword = "A String Very Very Very Niubilty!!@##$!@#4"

// A Util function to generate jwt_token which can be used in the request header
func GenToken(email string) string {
	jwt_token := jwt.New(jwt.GetSigningMethod("HS256"))
	// Set some claims
	jwt_token.Claims = jwt.MapClaims{
		"id":  email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	// Sign and get the complete encoded token as a string
	token, _ := jwt_token.SignedString([]byte(NBSecretPassword))
	return token
}

// My own Error type that will help return my customized Error info
//  {"database": {"hello":"no such table", error: "not_exists"}}
type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

// To handle the error returned by c.Bind in gin framework
// https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
func NewValidatorError(err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	// errs := err.(validator.ValidationErrors)
	// for _, v := range errs {
	// 	// can translate each error one at a time.
	// 	//fmt.Println("gg",v.NameNamespace)
	// 	if v.Param != "" {
	// 		res.Errors[v.Field] = fmt.Sprintf("{%v: %v}", v.Tag, v.Param)
	// 	} else {
	// 		res.Errors[v.Field] = fmt.Sprintf("{key: %v}", v.Tag)
	// 	}

	// }
	return res
}

// Warp the error info in a object
func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

// Changed the c.MustBindWith() ->  c.ShouldBindWith().
// I don't want to auto return 400 when error happened.
// origin function is here: https://github.com/gin-gonic/gin/blob/master/context.go
func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}



func Email(recemail string, emailsubject string, emailbody string) {
	
	email := gmail.Compose(emailsubject, emailbody)
	email.From = constants.FromEmail
	email.Password = constants.EmailPassword
	fmt.Println(recemail)

	// Defaults to "text/plain; charset=utf-8" if unset.
	email.ContentType = "text/html; charset=utf-8"

	email.AddRecipient(recemail)
	//email.AddRecipients("another@example.com", "more@example.com")

	err := email.Send()
	if err != nil {
		fmt.Println(err)
		// handle error.
	}
}

/**
This function will return user access token - permission as number


called from @useraccesstokenform.go		as @uat.GetUserAllowPermission
*/
func GetUserAllowPermission(permission string) int {
	if permission == "yes" {
		return constants.UAT_PERMISSION_YES
	} else if permission == "revoked" {
		return constants.UAT_PERMISSION_REVOKED
	} else {
		return constants.UAT_PERMISSION_NO
	}
}
