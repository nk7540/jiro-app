package repository

import (
	"artics-api/src/config"
	"artics-api/src/internal/domain/notice"
	"artics-api/src/internal/infrastructure/repository/models"
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/xerrors"
)

type noticeRepository struct {
	db *config.DatabaseConfig
}

func NewNoticeRepository(db *config.DatabaseConfig) notice.NoticeRepository {
	return &noticeRepository{db}
}

type insertable interface {
	Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

func (r *noticeRepository) Create(ctx context.Context, n *notice.Notice) error {
	// Begin transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Insert into notice
	mn := models.Notice{
		UserID: int(n.UserID),
		Type:   int(n.Type),
		IsRead: bool(n.IsRead),
	}
	if err := mn.Insert(ctx, r.db, boil.Infer()); err != nil {
		tx.Rollback()
		return err
	}

	// Insert into notice according to the type
	var insertable insertable
	switch n.Type {
	case notice.Favorite:
		insertable = &models.NoticeFavorite{
			NoticeID:            mn.ID,
			FavoriteID:          int(n.Favorite.FavoriteID),
			UserID:              int(n.Favorite.UserID),
			UserThumbnailURL:    string(n.Favorite.UserThumbnailURL),
			Header:              string(n.Favorite.Header),
			Body:                string(n.Favorite.Body),
			ContentID:           int(n.Favorite.ContentID),
			ContentThumbnailURL: string(n.Favorite.ContentThumbnailURL),
		}
	case notice.Followed:
		insertable = &models.NoticeFollowed{
			NoticeID:         mn.ID,
			UserID:           int(n.Followed.UserID),
			UserThumbnailURL: string(n.Followed.UserThumbnailURL),
			Body:             string(n.Followed.Body),
		}
	default:
		tx.Rollback()
		return xerrors.New("invalid notice type")
	}

	if err := insertable.Insert(ctx, r.db, boil.Infer()); err != nil {
		tx.Rollback()
		return err
	}

	// Finish the transaction
	return tx.Commit()
}

func (r *noticeRepository) List(ctx context.Context, qry notice.QueryArgListNotice) (*notice.QueryNotices, error) {
	ns := &notice.QueryNotices{}
	sql := `
		SELECT n.*, n_favorite.*, n_followed.*
		FROM notice AS n
		LEFT OUTER JOIN notice_favorite AS n_favorite ON n.notice_id = n_favorite.notice_id
		LEFT OUTER JOIN notice_followed AS n_followed ON n.notice_id = n_followed.notice_id
		WHERE n.user_id = ?
		ORDER BY n.created_at DESC
		LIMIT ? OFFSET ?
	`
	if err := models.NewQuery(qm.SQL(sql, qry.UserID, qry.Limit, qry.Offset)).Bind(ctx, r.db, ns); err != nil {
		return nil, err
	}
	return ns, nil
}
