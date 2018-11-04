# GoString - 如何製作設定檔？

## 設定檔架構
```
{
  "LangPath": "語言檔案的存放位置 (尾端必須加 /)",
  "Original": "來源語言檔案名稱 (不用 .json)",
  "TranLang": "目標語言檔案名稱 (不用 .json)"
}
```

- `LangPath` 是語言檔案的存放位置，尾端必須加 /<br/>
  ex. `lang`<br/>
  到時候找來源或目標語言檔案就是從 lang 這個資料夾找。

- `Original` 是來源字串的檔案名稱，尾端不用加 `.json`<br/>
  ex. `en_US`<br/>
  檔案名稱為 `en_US.json`，但不需要加 `.json` 就直接 en_US 即可。

- `TranLang` 是目標翻譯的檔案名稱，尾端不用加 `.json`<br/>
  ex. `zh_TW`<br/>
  檔案名稱為 `zh_TW.json`，但不需要加 `.json` 就直接 zh_TW 即可。

## 設定檔名稱
設定檔名稱隨意，只要在 `InitLang()` 填入正確的設定檔案名稱即可。<br/>
這部份可以參考 usage.md。
