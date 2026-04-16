package messages

var ZHTW = MessageMap{
	"root_use":                   "ohmypassword - 安全密碼產生器",
	"root_long":                  "輕鬆產生加密安全的密碼。",
	"generate_use":               "產生新密碼",
	"generate_long":              "使用可自訂選項產生安全密碼。",
	"version_use":                "顯示版本資訊",

	"flag_length":                "密碼長度 (8-128, 預設: 16)",
	"flag_charset":               "字元集 (upper/lower/digit/symbol/all, 預設: all)",
	"flag_strategy":              "產生策略 (simple/pronounceable/passphrase, 預設: simple)",
	"flag_count":                "產生密碼數量 (1-100, 預設: 1)",
	"flag_validate":              "顯示密碼強度",
	"flag_quiet":                 "安靜模式 (僅輸出密碼)",
	"flag_lang":                  "語言 (en/zh/zh-TW/ja/ko/es/fr)",
	"flag_output":                "輸出格式 (simple/json/csv/table)",
	"flag_exclude_similar":       "排除相似字元 (0, O, 1, l, I, |)",

	"output_entropy":             "熵值",
	"output_strength":            "強度",
	"output_crack_time":          "預計破解時間",

	"strength_very_weak":         "非常弱",
	"strength_weak":              "弱",
	"strength_reasonable":        "一般",
	"strength_strong":            "強",
	"strength_very_strong":       "非常強",

	"error_invalid_length":       "無效長度: 必須在 8 到 128 之間",
	"error_invalid_count":        "無效數量: 必須在 1 到 100 之間",
	"error_invalid_strategy":     "無效策略",
	"error_invalid_charset":      "無效字元集",
}