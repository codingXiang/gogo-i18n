package gogo_i18n

import (
	"encoding/json"
	"errors"
	i18n2 "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type (
	//GoGoi18nInterface GoGoi18n 的實作介面
	GoGoi18nInterface interface {
		Reload(lang language.Tag) GoGoi18nInterface
		SetFileType(t string) error
		SetUseLanguage(lang language.Tag)
		LoadTranslationFile(path string, langs ...language.Tag) error
		LoadTranslationFileArray(path string, langs []language.Tag) error
		GetMessage(messageID string, data interface{}) string
	}
	//GoGoi18n 封裝 go-i18n 的結構
	GoGoi18n struct {
		bundle   *i18n2.Bundle    //go-i18n 內的 bundle
		fileType string           //檔案類型
		local    *i18n2.Localizer //翻譯語言設定參數
	}
)

var (
	GGi18n GoGoi18nInterface //i18n 實例
)

//NewGoGoi18n 建立 GoGoi18n 實例
func NewGoGoi18n(lang language.Tag) GoGoi18nInterface {
	var instance = &GoGoi18n{
		bundle: i18n2.NewBundle(lang),
	}
	instance.local = i18n2.NewLocalizer(instance.bundle, lang.String())
	return instance
}

//Reload 重新讀取語言包
func (i *GoGoi18n) Reload(lang language.Tag) GoGoi18nInterface {
	i.bundle = i18n2.NewBundle(lang)
	i.local = i18n2.NewLocalizer(i.bundle, lang.String())
	return i
}

//SetFileType 設定翻譯檔類型
func (i *GoGoi18n) SetFileType(t string) error {
	i.fileType = t
	switch t {
	case "json":
		i.bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
		break
	case "yaml":
		i.bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
		break
	default:
		return errors.New("not support file type " + t + ", can use `json` or `yaml`.")
	}
	return nil
}

//SetUseLanguage 設定當前要翻譯的語言
func (i *GoGoi18n) SetUseLanguage(lang language.Tag) {
	i.local = i18n2.NewLocalizer(i.bundle, lang.String())
}

//LoadTranslationFile 讀取翻譯檔
func (i *GoGoi18n) LoadTranslationFile(path string, langs ...language.Tag) error {
	for _, lang := range langs {
		f := path + "/" + i.fileType + "/" + lang.String() + "." + i.fileType
		if _, err := i.bundle.LoadMessageFile(f); err != nil {
			return err
		}
	}
	return nil
}

//LoadTranslationFile 讀取翻譯檔
func (i *GoGoi18n) LoadTranslationFileArray(path string, langs []language.Tag) error {
	for _, lang := range langs {
		f := path + "/" + i.fileType + "/" + lang.String() + "." + i.fileType
		if _, err := i.bundle.LoadMessageFile(f); err != nil {
			return err
		}
	}
	return nil
}

//GetMessage 取得翻譯後的訊息
func (i *GoGoi18n) GetMessage(messageID string, data interface{}) string {
	return i.local.MustLocalize(&i18n2.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: data,
	})
}

func StoreDataToFile(fileType string, path string, datas []GoGoi18nMessageInterface) error {
	var (
		rs     = make(map[string][]map[string]map[string]string)
		method = json.Marshal
		err    error
	)

	//初始化路徑
	path = path + "/" + fileType + "/"

	for _, data := range datas {
		for key, value := range data.GetI18nJSONData() {
			rs[key] = append(rs[key], value)
		}
	}
	//判斷路徑是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		//建立路徑
		if err = os.MkdirAll(path, os.FileMode(0755)); err != nil {
			return err
		}
	}

	//選擇要轉換的檔案類型
	switch fileType {
	case "json":
		method = json.Marshal
		break;
	case "yaml":
		method = yaml.Marshal
		break;
	}

	//開始進行存檔
	for key, data := range rs {
		var (
			fileName = path + key + "." + fileType
		)
		if decoder, err := method(data); err != nil {
			return err
		} else {
			if err = ioutil.WriteFile(fileName, decoder, os.ModePerm); err != nil {
				return err
			}
		}
	}
	return err
}
