package models

type Users struct {
	Id int `json:"uid"`
	Name string `json:"name"`
	Password string `json:"password"`
}

type UserInfo struct {
	IdCard string `json:"id_card"`
	Age int `json:"age"`
	Sex string 	`json:"sex"`
	Address string `json:"address"`
	Phone string `json:"phone"`
}

