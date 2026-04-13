package model

import (
	__ "zk3/srv/basic/proto"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);comment:'用户名'"`
	Password string `gorm:"type:varchar(32);comment:'密码'"`
}

func (u *User) FindUserByUsername(db *gorm.DB, username string) error {
	return db.Debug().Where("username = ?", username).Find(&u).Error
}

type Goods struct {
	gorm.Model
	Name   string  `gorm:"type:varchar(20);comment:'商品名称'"`
	Price  float64 `gorm:"type:decimal(10,2);comment:'商品价格'"`
	Images string  `gorm:"type:varchar(100);comment:'商品图片'"`
	Bio    string  `gorm:"type:varchar(100);comment:'商品描述'"`
	Stock  int     `gorm:"type:int(11);comment:'商品库存'"`
	Status int     `gorm:"type:tinyint(1);comment:'商品状态'"`
}

func (g *Goods) FindGoodsByName(db *gorm.DB, name string) error {
	return db.Debug().Where("name = ?", name).Find(&g).Error
}

func (g *Goods) GoodsAdd(db *gorm.DB) error {
	return db.Debug().Create(&g).Error
}

func (g *Goods) FindGoodsById(db *gorm.DB, id int64) error {
	return db.Debug().Where("id = ?", id).First(&g).Error
}

type Member struct {
	gorm.Model
	UserId int `gorm:"type:int;comment:'用户id'"`
}

type Order struct {
	gorm.Model
	UserId  int     `gorm:"type:int;comment:'用户id'"`
	OrderSn string  `gorm:"type:varchar(30);comment:'订单号'"`
	PayType int     `gorm:"type:tinyint(1);comment:'支付方式'"`
	Total   float64 `gorm:"type:decimal(10,2);comment:'总价'"`
	Status  int     `gorm:"type:tinyint(1);comment:'订单状态'"`
}

type OrderItem struct {
	gorm.Model
	OrderId     int     `gorm:"type:int;comment:'订单号'"`
	GoodsName   string  `gorm:"type:varchar(20);comment:'商品名称'"`
	GoodsPrice  float64 `gorm:"type:decimal(10,2);comment:'商品价格'"`
	GoodsImages string  `gorm:"type:varchar(50);comment:'商品图片'"`
	GoodsBio    string  `gorm:"type:varchar(50);comment:'商品描述'"`
}

func GoodsDetail(db *gorm.DB, id int64) ([]*__.GoodsDetail, error) {
	var list []*__.GoodsDetail

	err := db.Debug().Model(&Goods{}).Where("id = ?", id).Select("goods.*").Find(&list).Error

	return list, err

}
