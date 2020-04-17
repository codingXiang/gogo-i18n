# gogo-i18n
封裝 github.com/nicksnyder/go-i18n/v2/i18n 的 Package

## 如何取得
```
go get -u github.com/codingXiang/gogo-i18n
```

## 如何使用
### 建立 GoGoi18n 物件
```
// 建立 GoGoi18n 物件，預設語言為中文
GGi18n = NewGoGoi18n(language.TraditionalChinese)
```
### 讀取多語系檔案
```
// 設定 Config 類別
GGi18n.SetFileType("yaml")
// 讀取語言檔案，路徑會自動添加檔案類別子路徑，例如路徑為 `/etc` 的話，讀取 yaml 檔案就會讀取 `/etc/yaml` 下的檔案
// 第二個參數以後可以選擇要讀取哪些語系的檔案
GGi18n.LoadTranslationFile("/Users/user/go/src/pkg/gogo-i18n/example",
	language.TraditionalChinese,
	language.English)
```

### 翻譯
```
// 訊息翻譯
msg := GGi18n.GetMessage("orm.create.success", map[string]interface{}{
	"Data": "User",
})
fmt.Println(msg)
```

### 透過程式產生檔案
```

//NewGoGoi18nMessage 說明
/// 參數1 : 語言別，使用 golang 預設的 language 包
/// 參數2 : 多語系參照的 key
/// 參數3 : 翻譯的內容，可以使用 template tag {{}} 內嵌變數

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
//StoreDataToFile 儲存資料至檔案中
/// 參數1 : 檔案類型，指令要儲存的檔案類型，目前支援 yaml 與 json
/// 參數2 : 檔案儲存路徑，路徑會自動添加參數1檔案類別子路徑，例如路徑為 `/etc` 的話，參數1為 yaml 的話就會將檔案儲存到 `/etc/yaml` 中 
/// 參數3 : 轉換的檔案，使用 GoGoi18nMessageInterface 類型的陣列
StoreDataToFile("yaml", "/Users/user/go/src/pkg/gogo-i18n/example", i18nDatas)
```