package handler

import (
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"log"
	"html/template"
	"fmt"
	"GoDemo/models"
)
var DB *sql.DB
var User models.Users
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("static/html/login.html")
		log.Println(t.Execute(w, nil))
		return
	}
	if r.Method == "POST" {

		username := r.FormValue("username")
		fmt.Println(username)
		fmt.Printf("%T\n",username)
		password := r.Form.Get("password")

		query_users_sql := `select * from users where name = ?`
		if stmt, err := DB.Prepare(query_users_sql); err != nil {
			defer stmt.Close()
			fmt.Println("DB prepare users error:", err)
			fmt.Fprintf(w, `{"code":500, "msg":"服务器异常!"}`)
		}else{
			defer stmt.Close()
			rows, err := stmt.Query(username)
			fmt.Println(rows)
			defer rows.Close()
			fmt.Println("helllllllllllll")
			for rows.Next() {
				err = rows.Scan(&User.Id, &User.Name, &User.Password )
				if err != nil {
					fmt.Println(111)
					fmt.Println(err.Error())
					continue
				}
				fmt.Println(222, User.Password)
				if User.Password == password {
					t, _ := template.ParseFiles("static/html/loginsuccess.html")
					t.Execute(w, nil)

				}else{
					fmt.Println(444)
					t, _ := template.ParseFiles("static/html/loginerr.html")
					log.Println(t.Execute(w, nil))
					return
				}
				}
			}

		}
	}

