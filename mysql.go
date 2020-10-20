package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mygo/main/my_object/struct"
)

var Db *sql.DB
func init(){
	var err error
	Db,err=sql.Open("mysql", "root:123456@tcp(localhost:3306)/school")
	if err!=nil{
		fmt.Println(err)
	}
}
func (d *_struct.Delu)Denglu (err error) {
	dl:= _struct.Delu{}
	err =Db.QueryRow("select num,password from delu where num=$1 and password=$2").Scan(&dl.Num,&dl.Password)
	return err
}
func (r *_struct.Delu)zhuce(err error)  {
	zc:= _struct.Delu{}
	_, err = Db.Exec("insert into delu(num,password) values (?,?)", &zc.Num, &zc.Password)
	return err
}
func retrieve(id int)(post _struct.Text,err error) {
	post = _struct.Text{}
	err =Db.QueryRow("select id,text from jishiben where id=$1").Scan(&post.ID)
	return
}
