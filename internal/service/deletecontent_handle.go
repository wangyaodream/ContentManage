package service

import (
	"ContentManage/api/operate"
	"context"
)

func (a *AppService) DeleteContent(ctx context.Context, req *operate.DeleteContentReq) (*operate.DeleteContentRsp, error) {
	uc := a.uc
	if err := uc.DeleteContent(ctx, req.GetId()); err != nil {
		return nil, err
	}
	return &operate.DeleteContentRsp{}, nil
}
