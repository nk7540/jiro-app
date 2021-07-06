package query

import (
	"artics-api/src/internal/domain/notice"
	"artics-api/src/pkg"
)

type NoticesHandler struct {
	nr notice.NoticeRepository
}

func NewNoticesHandler(nr notice.NoticeRepository) NoticesHandler {
	return NoticesHandler{nr}
}

func (h NoticesHandler) Handle(ctx pkg.Context, qry notice.QueryArgNotices) (*notice.QueryNotices, error) {
	return h.nr.List(ctx, notice.QueryArgListNotice{
		UserID: qry.UserID,
		Limit:  qry.Per,
		Offset: (qry.Page - 1) * qry.Per,
	})
}
