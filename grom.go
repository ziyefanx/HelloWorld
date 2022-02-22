package main

import (
	fmt "fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Student struct {
	StudentId int    `db:"studentId"`
	Name      string `db:"name"`
	Sex       int    `db:"sex"`
	Grade     int    `db:"grade"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:yxqc20161012@tcp(127.0.0.1:3306)/RUNOOB")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func insert() {
	var name string
	var sex, grade int
	fmt.Println("请输入要插入学生的信息：名字，性别，成绩")
	fmt.Scan(&name, &sex, &grade)
	r, err := Db.Exec("insert into Student(Name, Sex ,Grade)values(?, ?, ?)", name, sex, grade)
	if err != nil {
		fmt.Println("插入失败", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("插入失败", err)
		return
	}
	fmt.Println("插入成功:", id)
}
func Select() {
	var student []Student
	var id int
	fmt.Println("请输入学生的序号")
	fmt.Scan(&id)
	err := Db.Select(&student, "select studentId, name, sex, grade from student where studentId=?", id)
	if err != nil {
		fmt.Println("查询失败, ", err)
		return
	}
	fmt.Println("查询成功", student)
}
func update() {
	var name string
	var id, sex, grade int
	fmt.Println("请输入要修改学生的信息：学号，修改后的名字")
	fmt.Scan(&id, &name)
	res, err := Db.Exec("update student set name=? where studentId=?", name, id)
	if err != nil {
		fmt.Println("修改失败", err)
		return
	}
	fmt.Println("请输入要修改学生的信息：修改后的性别")
	fmt.Scan(&sex)
	res, err = Db.Exec("update student set sex=? where studentId=?", sex, id)
	if err != nil {
		fmt.Println("修改失败", err)
		return
	}
	fmt.Println("请输入要修改学生的信息：修改后的成绩")
	fmt.Scan(&grade)
	res, err = Db.Exec("update student set grade=? where studentId=?", grade, id)
	if err != nil {
		fmt.Println("修改失败", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("获取信息出错", err)
	}
	fmt.Println("修改成功:", row)
}
func Delete() {
	var id int
	fmt.Println("请输入要删除学生的学号")
	fmt.Scan(&id)
	res, err := Db.Exec("delete from student where studentId=?", id)
	if err != nil {
		fmt.Println("删除失败, ", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("行数获取失败, ", err)
	}

	fmt.Println("删除成功: 1", row)
}

func main() {
loop:
	for true {
		var choose int
		fmt.Println("请输入你的选择：" +
			"1.增加学生信息" +
			"2.删除学生信息" +
			"3.修改学生信息" +
			"4.查询学生信息" +
			"5.退出管理系统")
		fmt.Scan(&choose)
		switch choose {
		case 1:
			insert()
		case 2:
			Delete()
		case 3:
			update()
		case 4:
			Select()
		case 5:
			break loop
		default:
			fmt.Println("请重新输入")
		}
	}
}
