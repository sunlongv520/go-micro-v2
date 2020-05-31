package Boot

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"jtthink/src/Config"
	"time"
)


//mysql相关
var mysql_db *gorm.DB

func WaitForDbReady(d time.Duration)  {
	go func() {
		err:=WaitForReady(d,func() error {
			return InitMysql()
		},"数据库初始化成功","数据库初始化失败")
		if err!=nil{
			BootErrChan<-err
		}
	}()

}
func InitMysql() error {
	var err error
	mysql_db, err = gorm.Open("mysql",
		Config.JConfig.Data.Mysql.Dsn)
	if err != nil {
		 return NewFatalError(err.Error()) //这里返回致命异常
	}
	mysql_db.SingularTable(true)
	mysql_db.DB().SetMaxIdleConns(Config.JConfig.Data.Mysql.Maxidle)
	mysql_db.DB().SetMaxOpenConns(Config.JConfig.Data.Mysql.Maxopen)
	return nil
}
