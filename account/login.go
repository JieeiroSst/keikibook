package account

import (
	"fmt"
	"keikibook/utils"

	"github.com/jinzhu/gorm"
)

//login struct
type Login struct {
	ID       string `json:"ID" from:"ID" query:"ID"`
	Email    string `json:"Email" from:"Email" query:"Email"`
	Password string `json:"Password" from:"Password" query:"Password"`
	gorm.Model
}

func (login *Login) Validate() (map[string]interface{}, bool) {

	if login.ID == "" {
		return utils.Message(false, "Contact name should be on the payload"), false
	}

	if login.Email == "" {
		return utils.Message(false, "Phone number should be on the payload"), false
	}

	if login.Password == "" {
		return utils.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return utils.Message(true, "success"), true
}

func (login *Login) Create() map[string]interface{} {

	if resp, ok := login.Validate(); !ok {
		return resp
	}

	GetDB().Create(login)

	resp := utils.Message(true, "success")
	resp["login"] = login
	return resp
}

func GetContact(id int) *Login {

	login := &Login{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(login).Error
	if err != nil {
		return nil
	}
	return login
}

func GetContacts(user int) []*Login {

	login := make([]*Login, 0)
	err := GetDB().Table("logins").Where("user_id = ?", user).Find(&login).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return login
}
