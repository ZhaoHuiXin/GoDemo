package handler

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"os/user"
	"fmt"
	"os"
	"strings"
	"log"
	"database/sql"
	"net/http"
	"net/url"
)

var registerInfo string
func init(){
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
	splitArr := strings.Split(cnfStr, "\n")
	usernameItem := splitArr[1]
	passwordItem := splitArr[2]
	hostItem := splitArr[3]
	username := strings.Split(usernameItem,"=")[1]
	password := strings.Split(passwordItem,"=")[1]
	host := strings.Split(hostItem,"=")[1]
	registerInfo = username + ":" + password + "@tcp(" + host +
		":3306)/test?charset=utf8"
}

func Check(err error){
	if err != nil{
		log.Fatal(err)
	}
}

func Test_md5Str(t *testing.T){
	newpwd := md5Str("test123456")
	if newpwd == "test123456" || newpwd == ""{
		t.Fatal("md5 failed !")
	}
}

func Test_IndexHandler(t *testing.T){
	resp, err := http.Get("http://localhost:8080/mydemo/index")
	if err != nil{
		t.Fatal(err)
	}
	defer resp.Body.Close()
}

func Test_LoginHandler_1(t *testing.T){
	resp, err := http.Get("http://localhost:8080/mydemo/login")
	if err != nil{
		t.Fatal(err)
	}
	defer resp.Body.Close()
}

func Test_Register_1(t *testing.T){
	db, err := sql.Open("mysql",registerInfo)
	query_sql := `SELECT * FROM users where name="test"`
	rows, err := db.Query(query_sql)
	if err !=nil{
		t.Fatal(err)
	}
	if rows != nil{
		for rows.Next(){
			err := rows.Scan(&User.Id,&User.Name,&User.Password)
			Check(err)
		}
		query_sql := `select * from user_info where userid=?`
		stmt, err :=db.Prepare(query_sql)
		Check(err)
		//t.Log("USERID:",User.Id)

		rows, err = stmt.Query(User.Id)
		Check(err)
		for rows.Next(){
			err = rows.Scan(&UserInfo.Id,&UserInfo.UserId,&UserInfo.InfoId)
			Check(err)
			del_sql := `delete from info where id=?`
			stmt, err =db.Prepare(del_sql)
			Check(err)
			t.Log("InfoId:",UserInfo.InfoId)
			res, err :=stmt.Exec(UserInfo.InfoId)
			t.Log("err1:",err)

			del_sql = `delete from users where id=?`
			stmt, err = db.Prepare(del_sql)
			t.Log("UserId:",UserInfo.UserId)

			res, err = stmt.Exec(UserInfo.UserId)
			t.Log("err2:",err)

			del_sql = `delete from user_info where id=?`
			stmt, err = db.Prepare(del_sql)
			t.Log("Id:",UserInfo.Id)

			res,err = stmt.Exec(UserInfo.Id)
			t.Log("err3:",err)

			affect, err := res.RowsAffected()
			Check(err)

			t.Log("affect:",affect)
			db.Close()
		}
	}
	resp, err := http.PostForm("http://localhost:8080/mydemo/register",
		url.Values{"username": {"test"}, "password": {"123456"},
					"idcard":{"222222222222222222"},"age":{"18"},
					"sex":{"男"},"address":{"北京朝阳区"},
					"phone":{"12345678901"}})
	Check(err)
	defer resp.Body.Close()
}

func Test_LoginHandler_2(t *testing.T){
	resp, err := http.PostForm("http://localhost:8080/mydemo/register",
		url.Values{"username": {"test"}, "password": {"123456"}})
	Check(err)
	defer resp.Body.Close()
}

