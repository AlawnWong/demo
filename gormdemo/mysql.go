package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("mysql", "test:test123@tcp(192.168.137.57:3306)/testdb1?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect mysql database")
	}
	defer db.Close()
	fmt.Println("success to connect mysql database")

	// 自动迁移模式
	db.AutoMigrate(&Product{})
	fmt.Println("AutoMigrate Product table")

	// 创建
	db.Create(&Product{Code: "L1212", Price: 1000})
	fmt.Println("Create Product")

	// 读取
	var product Product
	// db.First(&product, 1) // 查询id为1的product
	db.First(&product, "code = ?", "L1212") // 查询code为l1212的product
	fmt.Printf("product read product = %s \n", product)

	// 更新 - 更新product的price为2000
	db.Model(&product).Update("Price", 2000)
	fmt.Printf("Update product Price = %d \n", 2000)

	// 删除 - 删除product
	db.Delete(&product)
	fmt.Printf("Delete product id = %d \n", product.ID)
}
