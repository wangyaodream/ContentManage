package service

import (
	"ContentManage/api/operate"
	"ContentManage/internal/biz"
)

type AppService struct {
	operate.UnimplementedAppServer

	uc *biz.GreeterUsecase
}

func NewAppService(uc *biz.GreeterUsecase) *AppService {
	return &AppService{uc: uc}
}
