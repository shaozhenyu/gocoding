package main

import (
	"fmt"
	//"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	//gorm.Model
	ID       int `sql:"AUTO_INCREMENT"`
	BIRTHDAY time.Time
	AGE      int
	NAME     string //`gorm:"size:255"` // Default size for string is 255, reset it with this tag
	//Num      int    `gorm:"AUTO_INCREMENT"`
}

func main() {
	dbhost := "127.0.0.1:3306"

	//dbhost := os.Getenv("GORM_DBADDRESS")
	if dbhost != "" {
		dbhost = fmt.Sprintf("tcp(%v)", dbhost)
	}
	fmt.Println(dbhost)
	//dbname := "test"
	dbname := ""
	//db, err := gorm.Open("mysql", fmt.Sprintf("root:@/%v?charset=utf8&parseTime=True&loc=Local", dbname))
	db, err := gorm.Open("mysql", fmt.Sprintf("root:@%v/%v?charset=utf8&parseTime=True&loc=Local", dbhost, dbname))
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	if db.HasTable("users") {
		fmt.Println("users has exists")
	}

	// err1 := db.CreateTable(&User{})
	// if err1.Error != nil {
	// 	fmt.Println("error : ", err1.Error)
	// 	return
	// }
	// err1 := db.DropTableIfExists(&User{})
	// if err1.Error != nil {
	// 	fmt.Println("error: ", err1.Error)
	// 	return
	// }

	// err1 := db.Model(&User{}).ModifyColumn("name", "varchar(255)")
	// if err1.Error != nil {
	// 	fmt.Println("error : ", err1.Error)
	// 	return
	// }

	// user := User{NAME: "sc", AGE: 18, BIRTHDAY: time.Now()}
	// err1 := db.Create(&user)
	// if err1.Error != nil {
	// 	fmt.Println("error : ", err1.Error)
	// 	return
	// }

	// user1 := []User{}
	// err1 := db.Order("age", true).Find(&user1)
	// if err1.Error != nil {
	// 	fmt.Println("error : ", err1.Error)
	// 	return
	// }
	user1 := []User{}
	age := []int{}
	err1 := db.Find(&user1).Pluck("age", &age)
	if err1.Error != nil {
		fmt.Println("error : ", err1.Error)
		return
	}
	fmt.Println(age)
	fmt.Println(user1)

	fmt.Println("success")
}
