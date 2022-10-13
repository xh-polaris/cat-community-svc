// Code generated by goctl. DO NOT EDIT!
// Source: cat_community.proto

package catcommunity

import (
	"context"

	"github.com/xh-polaris/cat-community-svc/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Cat           = pb.Cat
	CatDetailReq  = pb.CatDetailReq
	CatDetailResp = pb.CatDetailResp
	QueryCatReq   = pb.QueryCatReq
	QueryCatResp  = pb.QueryCatResp
	UploadCatReq  = pb.UploadCatReq
	UploadCatResp = pb.UploadCatResp

	CatCommunity interface {
		//  根据ID查询猫咪信息
		GetDetail(ctx context.Context, in *CatDetailReq, opts ...grpc.CallOption) (*CatDetailResp, error)
		//  查询猫咪信息
		QueryCat(ctx context.Context, in *QueryCatReq, opts ...grpc.CallOption) (*QueryCatResp, error)
		//  上传或更新猫咪信息
		UploadCat(ctx context.Context, in *UploadCatReq, opts ...grpc.CallOption) (*UploadCatResp, error)
	}

	defaultCatCommunity struct {
		cli zrpc.Client
	}
)

func NewCatCommunity(cli zrpc.Client) CatCommunity {
	return &defaultCatCommunity{
		cli: cli,
	}
}

//  根据ID查询猫咪信息
func (m *defaultCatCommunity) GetDetail(ctx context.Context, in *CatDetailReq, opts ...grpc.CallOption) (*CatDetailResp, error) {
	client := pb.NewCatCommunityClient(m.cli.Conn())
	return client.GetDetail(ctx, in, opts...)
}

//  查询猫咪信息
func (m *defaultCatCommunity) QueryCat(ctx context.Context, in *QueryCatReq, opts ...grpc.CallOption) (*QueryCatResp, error) {
	client := pb.NewCatCommunityClient(m.cli.Conn())
	return client.QueryCat(ctx, in, opts...)
}

//  上传或更新猫咪信息
func (m *defaultCatCommunity) UploadCat(ctx context.Context, in *UploadCatReq, opts ...grpc.CallOption) (*UploadCatResp, error) {
	client := pb.NewCatCommunityClient(m.cli.Conn())
	return client.UploadCat(ctx, in, opts...)
}
