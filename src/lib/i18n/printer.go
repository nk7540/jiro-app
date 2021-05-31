package i18n

import (
	"artics-api/src/internal/domain"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type entry struct {
	tag, key string
	msg      interface{}
}

var entries = [...]entry{
	{"en", domain.RequiredMessage, domain.RequiredMessage},
	{"ja", domain.RequiredMessage, "を入力してください。"},
	{"en", domain.EqFieldMessage, domain.EqFieldMessage},
	{"ja", domain.EqFieldMessage, "が%sと一致しません。"},
	{"en", domain.MinMessage, domain.MinMessage},
	{"ja", domain.MinMessage, "は%s字以上で入力してください。"},
	{"en", domain.MaxMessage, domain.MaxMessage},
	{"ja", domain.MaxMessage, "は%s字以下で入力してください。"},
	{"en", domain.EmailMessage, domain.EmailMessage},
	{"ja", domain.EmailMessage, "の形式が正しくありません。"},
	{"en", domain.UniqueMessage, domain.UniqueMessage},
	{"ja", domain.UniqueMessage, "は既に存在します。"},
	{"en", domain.CustomUniqueMessage, domain.CustomUniqueMessage},
	{"ja", domain.CustomUniqueMessage, "は既に存在します。"},
	{"en", domain.PasswordMessage, domain.PasswordMessage},
	{"ja", domain.PasswordMessage, "の形式が正しくありません。"},
	{"en", domain.PasswordConfirmationMessage, domain.PasswordConfirmationMessage},
	{"ja", domain.PasswordConfirmationMessage, "がパスワードと一致しません。"},
	{"en", domain.RequiredMessage, domain.RequiredMessage},
	{"ja", domain.RequiredMessage, "を入力してください。"},
}

type I18nPrinter struct {
	Printer *message.Printer
}

func Init() {
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

func NewI18nPrinter(lang string) *I18nPrinter {
	p := message.NewPrinter(language.MustParse(lang))

	return &I18nPrinter{p}
}

func (p *I18nPrinter) Sprintf(key message.Reference, a ...interface{}) string {
	return p.Printer.Sprintf(key, a...)
}
