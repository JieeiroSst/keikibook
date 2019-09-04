package account

import (
	"errors"
	"keikibook/utils"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	UserId int
	jwt.StandardClaims
}

type SignUp struct {
	ID             string `json:"ID" from:"ID" query:"ID"`
	FullName       string `json:"FullName" from:"FullName" query:"FullName"`
	Email          string `json:"Email" from:"Email" query:"Email"`
	UserName       string `json:"UserName" from:"UserName" query:"UserName"`
	Password       string `json:"Password" from:"Password" query:"Password"`
	RepeatPassword string `json:"RepeatPassword" from:"RepeatPassword" query:"RepeatPassword"`
	Token          string `json:"Token" from:"Token" query:"Token"`
	lock           sync.Mutex
	gorm.Model
}

func NewAccount(InID, InFullName, InEmail, InUserName, InPassowrd, InRepeatPassword string) SignUp {
	return SignUp{
		ID:             InID,
		FullName:       InFullName,
		Email:          InEmail,
		UserName:       InUserName,
		Password:       InPassowrd,
		RepeatPassword: InRepeatPassword,
	}
}

func (a *SignUp) checkAccont() SignUp {
	a.lock.Lock()
	defer a.lock.Unlock()
	return SignUp{
		ID:             a.ID,
		FullName:       a.FullName,
		Email:          a.Email,
		UserName:       a.UserName,
		Password:       a.Password,
		RepeatPassword: a.RepeatPassword,
	}
}

func (signUp *SignUp) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(signUp.Email, "@") {
		return utils.Message(false, "Email address is required"), false
	}

	if len(signUp.Password) < 6 {
		return utils.Message(false, "Password is required"), false
	}
	temp := &SignUp{}

	err := GetDB().Table("signUp").Where("email=?", signUp.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return utils.Message(false, "Email address already in use by another user."), false
	}
	return utils.Message(false, "Requirement passed"), true
}
func stringConvertInt(s string) int {
	myInt, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return myInt
}
func (_signUp *SignUp) createUser() map[string]interface{} {
	if resp, err := _signUp.Validate(); !err {
		return resp
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(_signUp.Password), bcrypt.DefaultCost)
	_signUp.Password = string(hashedPassword)

	GetDB().Create(_signUp)

	if stringConvertInt(_signUp.ID) <= 0 {
		return utils.Message(false, "Failed to create account, connection error.")
	}
	tk := &Token{UserId: stringConvertInt(_signUp.ID)}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	_signUp.Token = tokenString

	_signUp.Password = ""
	response := utils.Message(true, "Account has been created")
	response["signUp"] = _signUp
	return response
}

func emailFormatValid(input string) bool {
	if pass, _ := regexp.MatchString(
		`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`, input,
	); pass {
		return true
	}
	return false
}
func register(req SignUp) error {
	if len(req.ID) == 0 {
		return errors.New("length of username cannot be 0")
	}

	if len(req.Password) == 0 || len(req.RepeatPassword) == 0 {
		return errors.New("password and password reinput must be longer than 0")
	}

	if req.Password != req.RepeatPassword {
		return errors.New("password and reinput must be the same")
	}

	if emailFormatValid(req.Email) {
		return errors.New("invalid email")
	}

	createUser()
	return nil
}

func login(email, password string) map[string]interface{} {

	signUp := &SignUp{}
	err := GetDB().Table("signUps").Where("email = ?", email).First(signUp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Email address not found")
		}
		return utils.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(signUp.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return utils.Message(false, "Invalid login credentials. Please try again")
	}
	//Worked! Logged In
	signUp.Password = ""

	//Create JWT token
	tk := &Token{UserId: stringConvertInt(signUp.ID)}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	signUp.Token = tokenString //Store the token in the response

	resp := utils.Message(true, "Logged In")
	resp["account"] = signUp
	return resp
}

func GetUser(u int) *SignUp {

	acc := &SignUp{}
	GetDB().Table("SignUp").Where("id = ?", u).First(acc)
	if acc.Email == "" { //User not found!
		return nil
	}

	acc.Password = ""
	return acc
}
