package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Supplier struct {
	gorm.Model
	//Id             int64  //主键
	Distributor    string //供应商名
	PaymentTerm    string //付款周期
	Contacts       string //联系人
	Position       string //职务
	ContactsMobile string //联系人手机
	Email          string //联系邮箱
	Bank           string //银行账号
	AccountAddress string //账号所属地
	Address        string //公司地址
}



func main() {
	db, err := gorm.Open("mysql", "root:@/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	if err1 := db.CreateTable(&Supplier{}).Error; err1 != nil {
		fmt.Println("err : ", err1)
		return
	}
	fmt.Println("success")
}
