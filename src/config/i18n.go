package config

import (
	"artics-api/src/internal/domain"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type entry struct {
	tag, key string
	msg      interface{}
}

var entries = [...]entry{
	{"ja", domain.RequiredMessage, "を入力してください。"},
	{"ja", domain.EqFieldMessage, "が%sと一致しません。"},
	{"ja", domain.MinMessage, "は%s字以上で入力してください。"},
	{"ja", domain.MaxMessage, "は%s字以下で入力してください。"},
	{"ja", domain.EmailMessage, "の形式が正しくありません。"},
	{"ja", domain.UniqueMessage, "は既に存在します。"},
	{"ja", domain.CustomUniqueMessage, "は既に存在します。"},
	{"ja", domain.PasswordMessage, "の形式が正しくありません。"},
	{"ja", domain.PasswordConfirmationMessage, "がパスワードと一致しません。"},
	{"ja", domain.RequiredMessage, "を入力してください。"},
}

type I18nConfig struct {
	*message.Printer
}

func (c *I18nConfig) Setup() {
	for _, e := range entries {
		tag := language.MustParse(e.tag)
		switch msg := e.msg.(type) {
		case string:
			message.SetString(tag, e.key, msg)
		case catalog.Message:
			message.Set(tag, e.key, msg)
		case []catalog.Message:
			message.Set(tag, e.key, msg...)
		}
	}
}

func (c *I18nConfig) NewMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		lang := c.Get("Accept-Language")
		p := message.NewPrinter(language.MustParse(lang))
		c.Locals("i18n", p)
		return c.Next()
	}
}
