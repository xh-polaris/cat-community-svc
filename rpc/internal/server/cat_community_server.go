// Code generated by goctl. DO NOT EDIT!
// Source: cat_community.proto

package server

import (
	"context"

	"github.com/xh-polaris/cat-community-svc/rpc/internal/logic"
	"github.com/xh-polaris/cat-community-svc/rpc/internal/svc"
	"github.com/xh-polaris/cat-community-svc/rpc/pb"
)

type CatCommunityServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedCatCommunityServer
}

func NewCatCommunityServer(svcCtx *svc.ServiceContext) *CatCommunityServer {
	return &CatCommunityServer{
		svcCtx: svcCtx,
	}
}

//  根据ID查询猫咪信息
func (s *CatCommunityServer) GetDetail(ctx context.Context, in *pb.CatDetailReq) (*pb.CatDetailResp, error) {
	l := logic.NewGetDetailLogic(ctx, s.svcCtx)
	return l.GetDetail(in)
}

//  查询猫咪信息
func (s *CatCommunityServer) QueryCat(ctx context.Context, in *pb.QueryCatReq) (*pb.QueryCatResp, error) {
	l := logic.NewQueryCatLogic(ctx, s.svcCtx)
	return l.QueryCat(in)
}

//  上传或更新猫咪信息
func (s *CatCommunityServer) UploadCat(ctx context.Context, in *pb.UploadCatReq) (*pb.UploadCatResp, error) {
	l := logic.NewUploadCatLogic(ctx, s.svcCtx)
	return l.UploadCat(in)
}
