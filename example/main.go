package main

import (
	"errors"
	"fmt"
	. "github.com/codingXiang/gogo-i18n"
	"golang.org/x/text/language"
)

type Test struct {
	Model string
}

func main() {
	//設定 i18n 翻譯資料
	var (
		i18nDatas = []GoGoi18nMessageInterface{
			NewGoGoi18nMessage(language.TraditionalChinese, "orm.create.success", "建立 {{.Data}} 成功"),
			NewGoGoi18nMessage(language.TraditionalChinese, "orm.create.failed", "建立 {{.Data}} 失敗，原因為 {{.Error}}"),
			NewGoGoi18nMessage(language.TraditionalChinese, "orm.update.success", "更新 {{.Data}} 成功"),
			NewGoGoi18nMessage(language.TraditionalChinese, "orm.update.failed", "更新 {{.Data}} 失敗，原因為 {{.Error}}"),
			NewGoGoi18nMessage(language.TraditionalChinese, "orm.get.success", "取得 {{.Data}} 成功"),
			NewGoGoi18nMessage(language.TraditionalChinese, "orm.get.failed", "取得 {{.Data}} 失敗，原因為 {{.Error}}"),
			NewGoGoi18nMessage(language.English, "orm.create.success", "Create {{.Data}} success"),
			NewGoGoi18nMessage(language.English, "orm.create.failed", "Create {{.Data}} failed，because {{.Error}}"),
			NewGoGoi18nMessage(language.English, "orm.update.success", "Update {{.Data}} success"),
			NewGoGoi18nMessage(language.English, "orm.update.failed", "Update {{.Data}} failed，because {{.Error}}"),
			NewGoGoi18nMessage(language.English, "orm.get.success", "Get {{.Data}} success"),
			NewGoGoi18nMessage(language.English, "orm.get.failed", "Get {{.Data}} failed，because {{.Error}}"),
		}
	)
	//儲存資料至檔案中
	StoreDataToFile("yaml", "/Users/user/go/src/pkg/gogo-i18n/example", i18nDatas)
	LangHandler = NewLanguageHandler()
	// 建立 GoGoi18n 物件，預設語言為中文
	lang, _ := LangHandler.GetLanguageTag("zh-Han")
	GGi18n = NewGoGoi18n(lang)
	// 設定 Config 類別
	GGi18n.SetFileType("yaml")
	// 讀取 語言檔案
	GGi18n.LoadTranslationFile("/Users/user/go/src/pkg/gogo-i18n/example",
		language.TraditionalChinese,
		language.English)
	// 正確訊息翻譯
	msg := GGi18n.GetMessage("orm.create.success", map[string]interface{}{
		"Data": "User",
	})
	// 錯誤訊息翻譯
	msgErr := GGi18n.GetMessage("orm.update.failed", map[string]interface{}{
		"Data":  "User",
		"Error": errors.New("錯誤"),
	})
	fmt.Println(msg)
	fmt.Println(msgErr)

	//更換語言別為英文
	lang1, _ := LangHandler.GetLanguageTag("en")
	GGi18n.SetUseLanguage(lang1)
	msg1 := GGi18n.GetMessage("orm.create.success", map[string]string{
		"Data": "User",
	})
	msg1Err := GGi18n.GetMessage("orm.get.failed", map[string]interface{}{
		"Data":  "User",
		"Error": errors.New("err"),
	})
	fmt.Println(msg1)
	fmt.Println(msg1Err)

	// 重新設定物件，變更預設語系為英文
	GGi18n.Reload(language.English)
}
