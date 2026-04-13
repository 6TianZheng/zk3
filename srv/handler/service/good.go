package service

import (
	"context"
	"errors"
	"zk3/srv/basic/config"
	__ "zk3/srv/basic/proto"
	"zk3/srv/handler/model"
)

type Server struct {
	__.UnimplementedGoodsServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Login(_ context.Context, in *__.LoginReq) (*__.LoginResp, error) {

	var user model.User

	err := user.FindUserByUsername(config.DB, in.Username)
	if err != nil {
		return nil, errors.New("查询失败")
	}

	if user.Username == "" {
		return nil, errors.New("用户不存在,请前往注册")
	}

	if user.Password != in.Password {
		return nil, errors.New("密码不正确,请重新输入")
	}

	return &__.LoginResp{
		Msg:  "登录成功",
		Code: 200,
	}, nil
}

func (s *Server) GoodsAdd(_ context.Context, in *__.GoodsAddReq) (*__.GoodsAddResp, error) {
	var goods model.Goods

	err := goods.FindGoodsByName(config.DB, in.Name)
	if err != nil {
		return nil, errors.New("查询失败")
	}

	if goods.Name != "" {
		return nil, errors.New("该商品已存在")
	}

	goods = model.Goods{
		Name:   in.Name,
		Price:  float64(in.Price),
		Images: in.Images,
		Bio:    in.Bio,
		Stock:  int(in.Stock),
		Status: 1,
	}

	err = goods.GoodsAdd(config.DB)
	if err != nil {
		return nil, errors.New("商品发布失败")
	}

	return &__.GoodsAddResp{
		Msg:  "商品发布成功",
		Code: 200,
	}, nil

}

func (s *Server) GoodsDetail(_ context.Context, in *__.GoodsDetailReq) (*__.GoodsDetailResp, error) {

	var goods model.Goods

	err := goods.FindGoodsById(config.DB, in.GoodsId)
	if err != nil {
		return nil, errors.New("商品不存在")
	}

	list, err := model.GoodsDetail(config.DB, in.GoodsId)
	if err != nil {
		return nil, errors.New("查询失败")
	}

	return &__.GoodsDetailResp{
		Msg:  "商品发布成功",
		Code: 200,
		List: list,
	}, nil
}
