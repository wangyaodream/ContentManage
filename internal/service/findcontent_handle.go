package service

import (
	"ContentManage/api/operate"
	"ContentManage/internal/biz"
	"context"
)

func (a *AppService) FindContent(ctx context.Context, req *operate.FindContentReq) (*operate.FindContentRsp, error) {
	uc := a.uc
	results, total, err := uc.FindContent(ctx, &biz.FindParams{
		ID:       req.GetId(),
		Author:   req.GetAuthor(),
		Title:    req.GetTitle(),
		Page:     req.GetPage(),
		PageSize: req.GetPageSize(),
	})
	if err != nil {
		return nil, err
	}

	var contents []*operate.Content
	for _, r := range results {
		contents = append(contents, &operate.Content{
			Id:             r.ID,
			Title:          r.Title,
			VideoUrl:       r.VideoURL,
			Author:         r.VideoURL,
			Description:    r.Description,
			Thumbnail:      r.Thumbnail,
			Category:       r.Category,
			Duration:       r.Duration.Milliseconds(),
			Resolution:     r.Resolution,
			FileSize:       r.FileSize,
			Format:         r.Format,
			Quality:        r.Quality,
			ApprovalStatus: r.ApprovalStatus,
		})
	}
	rsp := &operate.FindContentRsp{
		Total:    total,
		Contents: contents,
	}
	return rsp, nil
}
