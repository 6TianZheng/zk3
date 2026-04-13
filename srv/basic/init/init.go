package init

import (
	"fmt"
	"time"
	"zk3/srv/basic/config"
	"zk3/srv/handler/model"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	ViperInit()
	MysqlInit()
	RedisInit()
}

func ViperInit() {
	viper.SetConfigFile("../../../config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic("配置读取失败: " + err.Error())
	}

	err = viper.Unmarshal(&config.GlobalConfig)
	if err != nil {
		panic("配置解析失败: " + err.Error())
	}
	fmt.Println(config.GlobalConfig)
}

func MysqlInit() {
	var err error
	data := config.GlobalConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		data.User,
		data.Password,
		data.Host,
		data.Port,
		data.Database)
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}
	fmt.Println("数据库连接成功")

	sqlDB, err := config.DB.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了可以重新使用连接的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = config.DB.AutoMigrate(&model.User{}, &model.Goods{}, &model.Member{}, &model.Order{}, &model.OrderItem{})
	if err != nil {
		panic("数据库迁移失败: " + err.Error())
	}
	fmt.Println("数据库迁移成功")

}

func RedisInit() {
	data := config.GlobalConfig.Redis
	Addr := fmt.Sprintf("%s:%d", data.Host, data.Port)
	config.Rdb = redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: data.Password, // no password set
		DB:       data.Database, // use default DB
	})

	err := config.Rdb.Ping(config.Ctx).Err()
	if err != nil {
		panic("redis连接失败: " + err.Error())
	}
	fmt.Println("redis连接成功")
}
