package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"mygo/main/my/structural_text"
	_ "mygo/main/my/structural_text"
)
var Db *sql.DB
func Init() {
	var err error
	Db,err=sql.Open("mysql", "root:123456@tcp(localhost:3306)/school")
	if err!=nil{
		return
	}
}
func Close()  {
	err:=Db.Close()
	if err!=nil{
		return
	}
}
func  Login (login structural_text.Login) bool{
	stmt, _ :=Db.Prepare("select accout,password from login where accout=? and password=?")
	rows,err:=stmt.Query(login.Accout,login.Password)
	var num1 string
	var password1 string
	if err!=nil{
		return false
	}else{
		for rows.Next(){
			rows.Columns()
			rows.Scan(&num1,&password1)
		}
		if login.Accout==num1&&login.Password==password1{
			return true
		} else{
			return false
		}
	}
}
func Register(login structural_text.Login) bool {
	stmt1, _ :=Db.Prepare("select accout from login where accout=? ")
	_,err2:=stmt1.Query(login.Accout)
	if err2!=nil{
		return false
	}else{
		stmt, err := Db.Prepare("insert into login(accout,password)values(?,?)")
		_,err1:=stmt.Exec(login.Accout,login.Password)
		if err!=nil{
			return false
		}else{
			if err1!=nil{
				return false
			} else{
				return true
			}
		}
	}
}
func Addition(test structural_text.Text) bool {
	stmt, _ := Db.Prepare("select id from Text where id=? and user1=?")
	_, err := stmt.Query(test.Num,test.User)
	if err != nil {
		return false
	} else {
		stmt1,_:=Db.Prepare("insert into Text(id,text,user1)values(?,?,?)")
		_,err:=stmt1.Exec(test.Num,test.Text,test.User)
		if err!=nil{
			return false
		}else{
			return true
		}
	}
}
func Delete(test structural_text.Text) bool {
	stmt, _ := Db.Prepare("select id from Text where id=? and user1=?")
	res, _:= stmt.Query(test.Num,test.User)
	var id string
	for res.Next() {
		res.Scan(&id)
	}
	if id==test.Num {
		stmt,_:=Db.Prepare("delete from Text where id=? and user1=?")
		_,err:=stmt.Exec(test.Num,test.User)
		if err!=nil{
			return false
		}else {
			return true
		}
	}else{
		return false
	}
}
func Modification(test structural_text.Text) bool {
	stmt, _ := Db.Prepare("select id from Text where id=? and user1=?")
	res, _:= stmt.Query(test.Num,test.User)
	var id string
	for res.Next() {
		res.Scan(&id)
	}
	if id==test.Num {
		stmt,_:=Db.Prepare("update Text set text=? where id=? and user1=?")
		_,err:=stmt.Exec(test.Text,test.Num,test.User)
		if err!=nil{
			return false
		}else {
			return true
		}
	}else{
		return false
	}
}
func Show(test structural_text.Text)(test2[] structural_text.Text){

				var data[] structural_text.Text
				stmt,_ := Db.Prepare("select id,text from Text where user1=?")
				rows,_:=stmt.Query(test.User)
				for rows.Next()  {
					var test1 structural_text.Text
			// 扫描行并把扫描到到数据赋值
			rows.Scan(&test1.Num,&test1.Text)
			data = append(data,test1)
	}
	return data
}