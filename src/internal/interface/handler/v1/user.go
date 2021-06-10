package v1

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"artics-api/src/internal/application"
	"artics-api/src/internal/application/query"
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/user"
	"artics-api/src/internal/usecase/request"
	"artics-api/src/internal/usecase/response"
	"artics-api/src/pkg"
)

// V1UserHandler - v1 user handler
type V1UserHandler interface {
	Create(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Followings(c *fiber.Ctx) error
	Followers(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Suspend(c *fiber.Ctx) error
	Follow(c *fiber.Ctx) error
	Unfollow(c *fiber.Ctx) error
}

type v1UserHandler struct {
	app application.UserApplication
}

// NewV1UserHandler - setups v1 user handler
func NewV1UserHandler(app application.UserApplication) V1UserHandler {
	return &v1UserHandler{app}
}

func (h *v1UserHandler) Create(c *fiber.Ctx) error {
	req := &request.CreateUser{}
	if err := c.BodyParser(req); err != nil {
		return domain.UnableParseJSON.New(err)
	}

	if err := h.app.Commands.CreateUser.Handle(pkg.Context{Ctx: c}, user.CommandCreateUser{
		Email:                req.Email,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}); err != nil {
		return err
	}

	return c.JSON(nil)
}

func (h *v1UserHandler) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	ctx := pkg.Context{Ctx: c}
	u, err := h.app.Queries.GetUser.Handle(ctx, id)
	if err != nil {
		return err
	}

	favoriteContents, err := h.app.Queries.GetFavoriteContents.Handle(ctx, query.GetFavoriteContents{
		UserID: id,
		Limit:  3,
	})

	resFavoriteContents := make([]*response.Content, len(favoriteContents))
	for i, c := range favoriteContents {
		resFavoriteContents[i] = &response.Content{
			ID:    c.ID,
			Title: c.Title,
		}
	}

	res := &response.ShowUser{
		ID:               u.ID,
		Nickname:         u.Nickname,
		Followingcount:   u.FollowingCount,
		Followercount:    u.FollowerCount,
		FavoriteContents: resFavoriteContents,
	}

	return c.JSON(res)
}

func (h *v1UserHandler) Followings(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	us, err := h.app.Queries.Followings.Handle(pkg.Context{Ctx: c}, id)
	if err != nil {
		return err
	}

	resUsers := make([]*response.User, len(us.Users))
	for i, u := range us.Users {
		resUsers[i] = &response.User{
			ID:           u.ID,
			Nickname:     u.Nickname,
			ThumbnailURL: u.ThumbnailURL,
		}
	}
	res := &response.Users{Users: resUsers}

	return c.JSON(res)
}

func (h *v1UserHandler) Followers(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	us, err := h.app.Queries.Followers.Handle(pkg.Context{Ctx: c}, id)
	if err != nil {
		return err
	}

	resUsers := make([]*response.User, len(us.Users))
	for i, u := range us.Users {
		resUsers[i] = &response.User{
			ID:           u.ID,
			Nickname:     u.Nickname,
			ThumbnailURL: u.ThumbnailURL,
		}
	}
	res := &response.Users{Users: resUsers}

	return c.JSON(res)
}

func (h *v1UserHandler) Update(c *fiber.Ctx) error {
	ctx := pkg.Context{Ctx: c}
	req := &request.UpdateUser{}
	if err := c.BodyParser(req); err != nil {
		return domain.UnableParseFormData.New(err)
	}

	thumbnail, err := req.Thumbnail.Open()
	if err != nil {
		return domain.UnableParseFile.New(err)
	}
	thumbnailURL, err := h.app.Commands.UpdateThumbnail.Handle(ctx, thumbnail)

	if err := h.app.Commands.UpdateUser.Handle(ctx, user.CommandUpdateUser{
		Nickname:     user.Nickname(req.Nickname),
		ThumbnailURL: thumbnailURL,
	}); err != nil {
		return err
	}

	res := &response.UpdateUser{
		Nickname:     req.Nickname,
		ThumbnailURL: string(thumbnailURL),
	}

	return c.JSON(res)
}

func (h *v1UserHandler) Suspend(c *fiber.Ctx) error {
	ctx := pkg.Context{Ctx: c}
	u, _ := ctx.CurrentUser()
	if err := h.app.Commands.SuspendUser.Handle(ctx, u); err != nil {
		return err
	}

	return c.JSON(nil)
}

func (h *v1UserHandler) Follow(c *fiber.Ctx) error {
	followerID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	ctx := pkg.Context{Ctx: c}
	u, err := ctx.CurrentUser()
	if err != nil {
		return err
	}

	if err := h.app.Commands.Follow.Handle(ctx, user.CommandFollow{
		FollowingID: user.FollowingID(u.ID),
		FollowerID:  user.FollowerID(followerID),
	}); err != nil {
		return err
	}

	return c.JSON(nil)
}

func (h *v1UserHandler) Unfollow(c *fiber.Ctx) error {
	followerID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	ctx := pkg.Context{Ctx: c}
	u, err := ctx.CurrentUser()
	if err != nil {
		return err
	}

	if err := h.app.Commands.Unfollow.Handle(ctx, user.CommandUnfollow{
		FollowingID: user.FollowingID(u.ID),
		FollowerID:  user.FollowerID(followerID),
	}); err != nil {
		return err
	}

	return c.JSON(nil)
}
