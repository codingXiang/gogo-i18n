package gogo_i18n

import "golang.org/x/text/language"

type (
	//GoGoi18nMessageInterface
	GoGoi18nMessageInterface interface {
		TableName() string
		GetI18nJSONData() map[string]map[string]map[string]string
	}
	//GoGoi18nMessage 資料庫欄位
	GoGoi18nMessage struct {
		ID       uint64 `json:"id" yaml:"-" gorm:"type:bigint(20) unsigned auto_increment;not null;primary_key"`
		Language string `json:"language" gorm:"unique_index:idx1;comment:'語言別'"`
		Key      string `json:"key" gorm:"unique_index:idx1;comment:'多語的key'"`
		Value    string `json:"value" gorm:"unique_index:idx1;comment:'多語的值'"`
	}
)

func NewGoGoi18nMessage(language language.Tag, key string, value string) GoGoi18nMessageInterface {
	return &GoGoi18nMessage{
		Language: language.String(),
		Key:      key,
		Value:    value,
	}
}

//TableName 修改資料庫 Table 名稱
func (g *GoGoi18nMessage) TableName() string {
	return "i18n"
}

//GetI18nJSONData 取得 i18n 轉換後的資料
func (g *GoGoi18nMessage) GetI18nJSONData() map[string]map[string]map[string]string {
	return map[string]map[string]map[string]string{
		g.Language: {
			g.Key: {
				"other": g.Value,
			},
		},
	}
}
