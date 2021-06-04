package v1

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"artics-api/src/internal/domain"
	"artics-api/src/internal/usecase"
	"artics-api/src/internal/usecase/request"
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
}

type v1UserHandler struct {
	u usecase.UserUsecase
}

// NewV1UserHandler - setups v1 user handler
func NewV1UserHandler(u usecase.UserUsecase) V1UserHandler {
	return &v1UserHandler{u}
}

func (h *v1UserHandler) Create(c *fiber.Ctx) error {
	req := &request.CreateUser{}
	if err := c.BodyParser(req); err != nil {
		return domain.UnableParseJSON.New(err)
	}

	if err := h.u.Create(pkg.Context{Ctx: c}, req); err != nil {
		return err
	}

	return c.JSON(nil)
}

func (h *v1UserHandler) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	res, err := h.u.Show(pkg.Context{Ctx: c}, id)
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (h *v1UserHandler) Followings(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	res, err := h.u.Followings(pkg.Context{Ctx: c}, id)
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (h *v1UserHandler) Followers(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return domain.UnableParseJSON.New(err)
	}

	res, err := h.u.Followers(pkg.Context{Ctx: c}, id)
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (h *v1UserHandler) Update(c *fiber.Ctx) error {
	req := &request.UpdateUser{}
	if err := c.BodyParser(req); err != nil {
		return domain.UnableParseFormData.New(err)
	}

	res, err := h.u.Update(pkg.Context{Ctx: c}, req)
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (h *v1UserHandler) Suspend(c *fiber.Ctx) error {
	if err := h.u.Suspend(pkg.Context{Ctx: c}); err != nil {
		return err
	}

	return c.JSON(nil)
}
