package initialize

import (
	"fmt"
	"time"

	"github.com/go-ecommerce/global"
	po "github.com/go-ecommerce/internal/persistent_object"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	fmt.Println("InitMysql")
	m := global.Config.Mysql

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	fmt.Println("Connect mysql:", s)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{})
	checkErrorPanic(err, "Init mysql error!")
	global.Mdb = db

	SetPool()
	migrateTables()
}

func SetPool() {
	m := global.Config.Mysql

	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("Mysql error: %s ::", err)
	}
	// Số kết nối nhàn rỗi tối đa
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
	)
	if err != nil {
		fmt.Println("Migrating tables error:", err)
	}
}
