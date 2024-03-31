package biz

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// 业务实体
type Content struct {
	ID             int64         `json:"id"`
	Title          string        `json:"title" binding:"required"`     // 内容标题
	VideoURL       string        `json:"video_url" binding:"required"` // 视频播放URL
	Author         string        `json:"author" binding:"required"`    // 作者
	Description    string        `json:"description"`                  // 内容描述
	Thumbnail      string        `json:"thumbnail"`                    // 封面图URL
	Category       string        `json:"category"`                     // 内容分类
	Duration       time.Duration `json:"duration"`                     // 内容时长
	Resolution     string        `json:"resolution"`                   // 分辨率 如720p、1080p
	FileSize       int64         `json:"fileSize"`                     // 文件大小
	Format         string        `json:"format"`                       // 文件格式 如MP4、AVI
	Quality        int32         `json:"quality"`                      // 视频质量 1-高清 2-标清
	ApprovalStatus int32         `json:"approval_status"`              // 审核状态 1-审核中 2-审核通过 3-审核不通过
	UpdatedAt      time.Time     `json:"updated_at"`                   // 内容更新时间
	CreatedAt      time.Time     `json:"created_at"`                   // 内容创建时间
}

type FindParams struct {
	ID       int64
	Author   string
	Title    string
	Page     int64
	PageSize int64
}

// 内容实体对应的数据仓
type ContentRepo interface {
	Create(ctx context.Context, c *Content) error
	Update(ctx context.Context, id int64, c *Content) error
	IsExists(ctx context.Context, contentID int64) (bool, error)
	Delete(ctx context.Context, id int64) error
	Find(ctx context.Context, param *FindParams) ([]*Content, int64, error)
}

// GreeterUsecase is a Greeter usecase.
type ContentUsecase struct {
	repo ContentRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewContentUsecase(repo ContentRepo, logger log.Logger) *ContentUsecase {
	return &ContentUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *ContentUsecase) CreateContent(ctx context.Context, c *Content) error {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %+v", c)
	return uc.repo.Create(ctx, c)
}

func (uc *ContentUsecase) UpdateContent(ctx context.Context, c *Content) error {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %+v", c)
	return uc.repo.Update(ctx, c.ID, c)
}

func (uc *ContentUsecase) DeleteContent(ctx context.Context, id int64) error {
	// 首先检查条目是否存在
	repo := uc.repo
	isExists, err := repo.IsExists(ctx, id)
	if err != nil {
		return err
	}
	if !isExists {
		return errors.New("内容不存在")
	}
	// 内容存在就执行删除
	return repo.Delete(ctx, id)
}

func (uc *ContentUsecase) FindContent(ctx context.Context, params *FindParams) ([]*Content, int64, error) {
	repo := uc.repo
	contents, total, err := repo.Find(ctx, params)
	if err != nil {
		return nil, 0, err
	}
	return contents, total, nil

}
