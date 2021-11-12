package dao

import (
	"fmt"
	"time"
	"weblogindemo/model"
	"weblogindemo/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Dao 数据对象访问
type Dao struct {
	DB *gorm.DB // db
}

func NewDao() (*Dao, error) {
	db, err := gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))
	if err != nil {
		fmt.Printf("链接数据库失败，请检查参数:", err)
		return nil, err
	}

	// 禁用默认表名的复数形式
	db.SingularTable(true)

	// 用model创建数据库自动迁移
	db.AutoMigrate(&model.User{})

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.DB().SetConnMaxLifetime(10 * time.Second)

	dao := &Dao{
		DB: db,
	}
	return dao, nil
}
