package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "mygo/main/my/struct"
	_struct "mygo/main/my/struct"
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
func  Login (login _struct.Login) bool{
	stmt, _ :=Db.Prepare("select num,password from delu where num=? and password=?")
	rows,err:=stmt.Query(login.Num,login.Password)
	var num1 int
	var password1 string
	if err!=nil{
		return false
	}else{
		for rows.Next(){
			rows.Columns()
			rows.Scan(&num1,&password1)

		}
		if login.Num==num1&&login.Password==password1{
			return true
		} else{
			return false
		}
	}
}
func Register(login _struct.Login) bool {
	stmt1, _ :=Db.Prepare("select num from delu where num=? ")
	_,err2:=stmt1.Query(login.Num)
	if err2!=nil{
		return false
	}else{
		stmt, err := Db.Prepare("insert into delu(num,password)values(?,?)")
		_,err1:=stmt.Exec(login.Num,login.Password)
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
func Addition(test _struct.Text) bool {
	stmt, _ := Db.Prepare("select id from Text where id=?")
	_, err := stmt.Query(test.Num)
	if err != nil {
		return false
	} else {
		stmt1,_:=Db.Prepare("insert into Text(id,text)values(?,?)")
		_,err:=stmt1.Exec(test.Num,test.Text)
		if err!=nil{
			return false
		}else{
			return true
		}
	}
}
func Delete(test _struct.Text) bool {
	stmt, _ := Db.Prepare("select id from Text where id=?")
	res, _:= stmt.Query(test.Num)
	var id int
	for res.Next() {
		res.Scan(&id)
	}
	if id==test.Num {
		stmt,_:=Db.Prepare("delete from Text where id=?")
		_,err:=stmt.Exec(test.Num)
		if err!=nil{
			return false
		}else {
			return true
		}
	}else{
		return false
	}
}
func Modification(test _struct.Text) bool {
	stmt, _ := Db.Prepare("select id from Text where id=?")
	res, _:= stmt.Query(test.Num)
	var id int
	for res.Next() {
		res.Scan(&id)
	}
	if id==test.Num {
		stmt,_:=Db.Prepare("update Text set text=? where id=?")
		_,err:=stmt.Exec(test.Text,test.Num)
		if err!=nil{
			return false
		}else {
			return true
		}
	}else{
		return false
	}
}
func Show()(test2 _struct.Text){
	var test1 _struct.Text
	rows,_ := Db.Query("select id,text from Text")
	for rows.Next()  {
		// 扫描行并把扫描到到数据赋值
		rows.Scan(&test1.Num,&test1.Text)
	}
	return test1
	//var test1 [] _struct.Text
	//rows,_ := Db.Query("select id,text from Text")
	//i:=0
	//for rows.Next()  {
	//	// 扫描行并把扫描到到数据赋值
	//	rows.Scan(&test1[i].Num,&test1[i].Text)
	//	i=i+1
	//}
	//return test1

}