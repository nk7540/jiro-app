package config

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/user"
	"artics-api/src/internal/interface/handler/response"

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
	// Response
	{"ja", response.BadRequest.Message, "リクエストが不正です。"},
	{"ja", response.Unauthorized.Message, "ログインが必要です。"},
	{"ja", response.Forbidden.Message, "アクセスが拒否されています。"},
	{"ja", response.NotFound.Message, "指定のデータは存在しません。"},
	{"ja", response.AlreadyExists.Message, "指定のデータは既に存在します。"},
	{"ja", response.InternalServerError.Message, "サーバーエラーです。"},
	// Domain Validation
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
	// Notifications
	{"ja", user.FavoriteNoticeTitle, "%s がお気に入り: %s"},
	{"ja", user.FollowedNoticeBody, "%s があなたをフォローしました。"},
	// User
	{"ja", "email", "Eメール"},
	{"ja", "password", "パスワード"},
	{"ja", "passwordConfirmation", "パスワード(確認)"},
}

var matcher language.Matcher

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

	matcher = language.NewMatcher([]language.Tag{
		language.Japanese,
		language.English,
	})
}

func (c *I18nConfig) NewMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		lang := c.Cookies("lang")
		accept := c.Get("Accept-Language")
		tag, _ := language.MatchStrings(matcher, lang, accept)

		p := I18nConfig{Printer: message.NewPrinter(tag)}
		c.Locals("i18n", p)
		return c.Next()
	}
}
