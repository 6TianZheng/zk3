package service

import (
	"context"
	"reflect"
	"testing"
	__ "zk3/srv/basic/proto"
)

func TestServer_GoodsAdd(t *testing.T) {
	type fields struct {
		UnimplementedGoodsServer __.UnimplementedGoodsServer
	}
	type args struct {
		in0 context.Context
		in  *__.GoodsAddReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *__.GoodsAddResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGoodsServer: tt.fields.UnimplementedGoodsServer,
			}
			got, err := s.GoodsAdd(tt.args.in0, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GoodsAdd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GoodsAdd() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GoodsDetail(t *testing.T) {
	type fields struct {
		UnimplementedGoodsServer __.UnimplementedGoodsServer
	}
	type args struct {
		in0 context.Context
		in  *__.GoodsDetailReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *__.GoodsDetailResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGoodsServer: tt.fields.UnimplementedGoodsServer,
			}
			got, err := s.GoodsDetail(tt.args.in0, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GoodsDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GoodsDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Login(t *testing.T) {
	type fields struct {
		UnimplementedGoodsServer __.UnimplementedGoodsServer
	}
	type args struct {
		in0 context.Context
		in  *__.LoginReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *__.LoginResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedGoodsServer: tt.fields.UnimplementedGoodsServer,
			}
			got, err := s.Login(tt.args.in0, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}
