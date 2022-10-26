package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	Name        string
	Id          string
	Tel         string
	Bir         string
	Position    string
	Work_number int
}

var db *gorm.DB

type try interface {
	serch_name()
	serch_id()
	updata()
	serch_bigger_than_150()
}

func (u User) serch_bigger_than_150() {
	var codes []User
	db.Where("work_number >= ?", 150).Find(&codes)
	for i, u := range codes {
		fmt.Println(i, u)
	}
}
func (u User) serch_id() {
	var id string
	fmt.Println("please input the id which you want to serch")
	fmt.Scan(&id)
	db.Where("Id = ?", id).Find(&u)
	fmt.Println(u)
}
func (u User) serch_name() {
	var na string
	fmt.Println("please input the name which you want to serch")
	fmt.Scan(&na)
	db.Where("name = ?", na).Find(&u)
	fmt.Println(u)
}
func (u User) updata() {
	var na string
	fmt.Println("please input the name which you want to updata")
	fmt.Scan(&na)
	db.Where("name = ?", na).Find(&u)
	db.Model(&u).Update("name", "immhd")
}
func Initmysql() (err error) {
	db, err = gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/my_db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return
}
func input() {
	fmt.Println("how many persons do you want to join?")
	var number int
	fmt.Scan(&number)
	u := User{}
	for i := 0; i < number; i++ {
		fmt.Println("please input the name")
		fmt.Scan(&u.Name)
		fmt.Println("please input the id")
		fmt.Scan(&u.Id)
		fmt.Println("please input the tel")
		fmt.Scan(&u.Tel)
		fmt.Println("please input the bir")
		fmt.Scan(&u.Bir)
		fmt.Println("please input the position")
		fmt.Scan(&u.Position)
		fmt.Println("please input work_number")
		fmt.Scan(&u.Work_number)
		db.Create(&u)
	}
	return
}
func main() {
	var u User
	Initmysql()
	defer db.Close()
	db.AutoMigrate(&User{})
	//链接数据库
	input()
	//完成读入操作
	u.serch_bigger_than_150()
	//统计做题数大于150的同学并输出其完整信息
	u.serch_name()
	//按姓名查找
	u.serch_id()
	//按id查找
	u.updata()
	//实现name修改功能
}
