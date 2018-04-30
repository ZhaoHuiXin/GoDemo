package main

import (
	"os/user"
	"fmt"
	"os"
	"strings"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	_ "database/sql"
	"net/http"
	. "GoDemo/handler"
	"log"
)

func init() {
	currentUser, err := user.Current()
	if err != nil{
		fmt.Println(err)
	}
	userHome := currentUser.HomeDir
	configFile := userHome + "/.testdemo.cnf"
	openConfig, _ := os.Open(configFile)
	defer openConfig.Close()
	buf := make([]byte, 1024)
	openConfig.Read(buf)
	cnfStr := string(buf)
	fmt.Println(cnfStr)
	splitArr := strings.Split(cnfStr, "\n")
	usernameItem := splitArr[1]
	passwordItem := splitArr[2]
	hostItem := splitArr[3]
	username := strings.Split(usernameItem,"=")[1]
	password := strings.Split(passwordItem,"=")[1]
	host := strings.Split(hostItem,"=")[1]
	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(host)
	registerInfo := username + ":" + password + "@tcp(" + host +
		":3306)/test?charset=utf8"
	DB, err = sql.Open("mysql",registerInfo)

 }

func main() {
	defer DB.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("/mydemo/login", LoginHandler)
	mux.HandleFunc("/mydemo/register", RegisterHandler)
	//mux.Handle("/mydemo/login", http.HandlerFunc(LoginHandler))
	fmt.Println("over")
	err := http.ListenAndServe(":8080", mux)
	if err != nil{
		log.Fatal(err)
	}
}
