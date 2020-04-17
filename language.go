package gogo_i18n

import "golang.org/x/text/language"

type (
	LanguageHandlerInterface interface {
		GetLanguageTag(lang string) (language.Tag, error)
	}
	LanguageHandler struct {}
)

// 預設搜尋的語系，可以自行增加
var (
	ServerLanguage = []language.Tag{
		language.English,
		language.TraditionalChinese,
		language.SimplifiedChinese,
	}
	LangHandler LanguageHandlerInterface
)

//NewLanguageHandler 建立語言處理器
func NewLanguageHandler() LanguageHandlerInterface {
	return &LanguageHandler{}
}

//GetLanguageTag 傳入字串取得 Language 的 tag
func (l *LanguageHandler) GetLanguageTag(lang string) (language.Tag, error) {
	var matcher = language.NewMatcher(ServerLanguage)
	if t, _, err := language.ParseAcceptLanguage(lang); err != nil {
		return ServerLanguage[0], err
	} else {
		_, idx, _ := matcher.Match(t...)
		return ServerLanguage[idx], nil
	}
}
