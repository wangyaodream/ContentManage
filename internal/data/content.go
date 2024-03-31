package data

import (
	"ContentManage/internal/biz"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type contentRepo struct {
	data *Data
	log  *log.Helper
}
type ContentDetail struct {
	ID             int64         `gorm:"column:id;primary_key"`  // 内容ID
	Title          string        `gorm:"column:title"`           // 内容标题
	Description    string        `gorm:"column:description"`     // 内容描述
	Author         string        `gorm:"column:author"`          // 作者
	VideoURL       string        `gorm:"column:video_url"`       // 视频播放URL
	Thumbnail      string        `gorm:"column:thumbnail"`       // 封面图URL
	Category       string        `gorm:"column:category"`        // 内容分类
	Duration       time.Duration `gorm:"column:duration"`        // 内容时长
	Resolution     string        `gorm:"column:resolution"`      // 分辨率 如720p、1080p
	FileSize       int64         `gorm:"column:fileSize"`        // 文件大小
	Format         string        `gorm:"column:format"`          // 文件格式 如MP4、AVI
	Quality        int32         `gorm:"column:quality"`         // 视频质量 1-高清 2-标清
	ApprovalStatus int32         `gorm:"column:approval_status"` // 审核状态 1-审核中 2-审核通过 3-审核不通过
	UpdatedAt      time.Time     `gorm:"column:updated_at"`      // 内容更新时间
	CreatedAt      time.Time     `gorm:"column:created_at"`      // 内容创建时间
}

func (c ContentDetail) TableName() string {
	return "cms_content.t_content_details"
}

// NewGreeterRepo .
func NewContentRepo(data *Data, logger log.Logger) biz.ContentRepo {
	return &contentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *contentRepo) Create(ctx context.Context, content *biz.Content) error {
	c.log.Infof("contntentRepo Create content = %+v", content)
	db := c.data.db
	detail := ContentDetail{
		Title:          content.Title,
		Description:    content.Description,
		Author:         content.Author,
		VideoURL:       content.VideoURL,
		Thumbnail:      content.Thumbnail,
		Category:       content.Category,
		Duration:       content.Duration,
		Resolution:     content.Resolution,
		FileSize:       content.FileSize,
		Format:         content.Format,
		Quality:        content.Quality,
		ApprovalStatus: content.ApprovalStatus,
	}
	if err := db.Create(&detail).Error; err != nil {
		c.log.WithContext(ctx).Errorf("content create error = %v", err)
	}
	return nil
}

func (c *contentRepo) Update(ctx context.Context, id int64, content *biz.Content) error {
	c.log.Infof("contntentRepo Update content = %+v", content)
	db := c.data.db
	detail := ContentDetail{
		ID:             content.ID,
		Title:          content.Title,
		Description:    content.Description,
		Author:         content.Author,
		VideoURL:       content.VideoURL,
		Thumbnail:      content.Thumbnail,
		Category:       content.Category,
		Duration:       content.Duration,
		Resolution:     content.Resolution,
		FileSize:       content.FileSize,
		Format:         content.Format,
		Quality:        content.Quality,
		ApprovalStatus: content.ApprovalStatus,
	}
	if err := db.Where("id = ?", id).Updates(&detail).Error; err != nil {
		c.log.WithContext(ctx).Errorf("content update error = %v", err)
	}
	return nil
}
