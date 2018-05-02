package models

type Users struct {
	Id int `json:"uid"`
	Name string `json:"name"`
	Password string `json:"password"`
}

type Info struct {
	Id int `json:"id"`
	IdCard string `json:"id_card"`
	Age int `json:"age"`
	Sex int	`json:"sex"`
	Address string `json:"address"`
	Phone int `json:"phone"`
}

type UserInfo struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	InfoId int `json:"info_id"`
}


