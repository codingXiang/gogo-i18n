package gogo_i18n

import "golang.org/x/text/language"

type (
	LanguageHandlerInterface interface {
		GetLanguageTag(lang string) (language.Tag, error)
	}
	LanguageHandler struct {
		defaultLanguage language.Tag
	}
)

// 預設搜尋的語系，可以自行增加
var (
	ServerLanguage = []language.Tag{
		language.TraditionalChinese,
		language.SimplifiedChinese,
		language.English,
		language.Japanese,
	}
	LangHandler LanguageHandlerInterface
)

//NewLanguageHandler 建立語言處理器
func NewLanguageHandler(defaultLanguage language.Tag) LanguageHandlerInterface {
	return &LanguageHandler{
		defaultLanguage: defaultLanguage,
	}
}

//GetLanguageTag 傳入字串取得 Language 的 tag
func (l *LanguageHandler) GetLanguageTag(lang string) (language.Tag, error) {
	var matcher = language.NewMatcher(ServerLanguage)
	if t, _, err := language.ParseAcceptLanguage(lang); err != nil {
		return l.defaultLanguage, err
	} else {
		_, idx, _ := matcher.Match(t...)
		return ServerLanguage[idx], nil
	}
}
