package messages

var ZH = MessageMap{
	"root_use":                   "ohmypassword - 安全密码生成器",
	"root_long":                  "轻松生成加密安全的密码。",
	"generate_use":               "生成新密码",
	"generate_long":              "使用可自定义选项生成安全密码。",
	"version_use":                "显示版本信息",

	"flag_length":                "密码长度 (8-128, 默认: 16)",
	"flag_charset":               "字符集 (upper/lower/digit/symbol/all, 默认: all)",
	"flag_strategy":              "生成策略 (simple/pronounceable/passphrase, 默认: simple)",
	"flag_count":                 "生成密码数量 (1-100, 默认: 1)",
	"flag_validate":              "显示密码强度",
	"flag_quiet":                 "静默模式 (仅输出密码)",
	"flag_lang":                  "语言 (en/zh/zh-TW/ja/ko/es/fr)",
	"flag_output":                "输出格式 (simple/json/csv/table)",
	"flag_exclude_similar":       "排除相似字符 (0, O, 1, l, I, |)",

	"output_entropy":             "熵值",
	"output_strength":            "强度",
	"output_crack_time":          "预计破解时间",

	"strength_very_weak":         "非常弱",
	"strength_weak":              "弱",
	"strength_reasonable":        "一般",
	"strength_strong":            "强",
	"strength_very_strong":       "非常强",

	"error_invalid_length":       "无效长度: 必须在 8 到 128 之间",
	"error_invalid_count":        "无效数量: 必须在 1 到 100 之间",
	"error_invalid_strategy":     "无效策略",
	"error_invalid_charset":      "无效字符集",
}