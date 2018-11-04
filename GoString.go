// GoString - Internationaization tool

package GoString

import (
	"encoding/json"
	"io/ioutil"
    "fmt"
    "os"
)

// Lang 建構體用來儲存語言資料。
type Lang struct {
	Original map[string]string // JSON 來源語言資料
	L10N     map[string]string // JSON 目標語言資料
}

// Setting 是 setting.json 資料的架構。
type Setting struct {
	// LangPath: 目標語言 JSON 檔案統一的存放位置 (尾端請加上 /！)
	LangPath string
	// Original: 來源語言 JSON 檔案的檔案名稱 (無後 .json 副檔名！)
	Original string
	// TranLang: 目標語言 JSON 檔案的檔案名稱 (無後 .json 副檔名！)
	TranLang string
}

// errh 是個處理 Error 的快速捷徑。
func errh(err error) {
	if err != nil {
		panic(err)
	}
}

// InitLang 是個初始化語言資料的工具。
// 使用方法請參閱 README.md。
//
// setting:  語言設定檔的檔案位址
// 回傳： Lang 建構體
func InitLang(setting string) *Lang {
	// 定義變數
	var err error
	var set = &Setting{}
	var original = make(map[string]string)
	var trans = make(map[string]string)
	var setRaw, origRaw, tranRaw []byte
	var lang = &Lang{}

	// 讀取設定檔
	setRaw, err = ioutil.ReadFile(setting)
	errh(err)
	err = json.Unmarshal(setRaw, &set)
	errh(err)

	// 開始解析設定檔
	origRaw, err = ioutil.ReadFile(set.LangPath + set.Original + ".json")
	errh(err)
	tranRaw, err = ioutil.ReadFile(set.LangPath + set.TranLang + ".json")
	errh(err)
	err = json.Unmarshal(origRaw, &original)
	errh(err)
	err = json.Unmarshal(tranRaw, &trans)
	errh(err)

	lang.Original = original
	lang.L10N = trans

	return lang
}

// Str 輸出 stri 的目標語言字串。
// 如果目標語言沒有該字串，則從來源語言尋找
// 如果來源語言亦沒有，則直接回傳 stri 的值並在
// Stderr 頻道發出 "語言字串不存在" 訊息
func (l Lang) Str(stri string) string {
  if value, ok := l.L10N[stri]; ok {
    return value
  } else if value, ok := l.Original[stri]; ok {
    return value
  } else {
    os.Stderr.WriteString(fmt.Sprintf(warning, stri))
    return stri
  }
}
