package db

import (
	"github.com/duke-git/lancet/v2/strutil"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	iniUtils "ruoyi-gin/src/ruoyi-common/core/ini"
	"ruoyi-gin/src/ruoyi-common/core/log"
	"time"
)

var SqlDB *gorm.DB

func init() {
	var err error

	if SqlDB, err = gorm.Open(mysql.Open(iniUtils.MYSQL_DSN), &gorm.Config{
		//if SqlDB, err = gorm.Open(mysql.Open("root:!QAZ2wsx#EDC4rfv@tcp(127.0.0.1:3306)/goruoyi?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}); err != nil {
		log.SugarLogger.Errorf("数据库连接失败")
		log.SugarLogger.Errorf(err.Error())
	} else {
		log.SugarLogger.Infof("数据库连接成功")
		//打印sql
		db, _ := SqlDB.DB()
		db.SetMaxIdleConns(iniUtils.MYSQL_MAX_IDLE_CONNS)
		db.SetMaxOpenConns(iniUtils.MYSQL_MAX_OPEN_CONNS)
		db.SetConnMaxLifetime(time.Minute * time.Duration(iniUtils.MYSQL_CONN_MAX_LIFETIME))
	}
}

// s 是接受select结果的指针，args是参数
func Exec[T any](sql string, s *T, args ...interface{}) {
	if strutil.HasPrefixAny(sql, []string{"select", "Select", "SELECT"}) {
		SqlDB.Raw(sql, args).Scan(s)
	} else {
		tx := SqlDB.Exec(sql, args)
		tx.Commit()
	}
}
