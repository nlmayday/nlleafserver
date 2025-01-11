package global

import (
	"fmt"
	"os"
	"server/conf"
	"server/global/internal"
	"server/utils"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	G_DB  *gorm.DB
	G_LOG *zap.Logger
	G_VIP *viper.Viper
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func InitZap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(conf.Server.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", conf.Server.Zap.Director)
		_ = os.Mkdir(conf.Server.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if conf.Server.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// GormMysql 初始化Mysql数据库
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func GormMysql() *gorm.DB {
	m := conf.Server.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

func InitGlobal() {
	G_LOG = InitZap() // 初始化zap日志库
	zap.ReplaceGlobals(G_LOG)
	G_LOG.Log(zap.InfoLevel, "init zap success")

	G_DB = GormMysql() // gorm连接数据库
	G_LOG.Log(zap.InfoLevel, "init global success")
}
