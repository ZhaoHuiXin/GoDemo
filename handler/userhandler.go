package handler

import (
	"net/http"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)
var DB *sql.DB
func LoginHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("hello !")
	query_users_sql := `insert into userinfo(username) values(?)`
	if stmt, err := DB.Prepare(query_users_sql); err != nil{
		fmt.Println("DB prepare users error:", err)
		fmt.Fprintf(w, `{"code":500, "msg":"服务器异常!"}`)
		return
	}else{
		stmt.Exec("lucyca")
	}


}