package service

import (
	"ContentManage/api/operate"
	"ContentManage/internal/biz"
)

type AppService struct {
	operate.UnimplementedAppServer

	uc *biz.ContentUsecase
}

func NewAppService(uc *biz.ContentUsecase) *AppService {
	return &AppService{uc: uc}
}
