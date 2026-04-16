package messages

var JA = MessageMap{
	"root_use":      "ohmypassword - 安全なパスワードジェネレーター",
	"root_long":     "暗号学的に安全なパスワードを簡単に生成します。",
	"generate_use":  "新しいパスワードを生成",
	"generate_long": "カスタマイズ可能なオプションで安全なパスワードを生成します。",
	"check_use":     "パスワード強度を確認",
	"check_long":    "1つまたは複数のパスワードの強度を確認します。",
	"version_use":   "バージョン情報を表示",

	"flag_length":          "パスワードの長さ (8-128, デフォルト: 16)",
	"flag_charset":         "文字セット (upper/lower/digit/symbol/all, デフォルト: all)",
	"flag_strategy":        "生成戦略 (simple/pronounceable/passphrase/memorable, デフォルト: simple)",
	"flag_count":           "生成するパスワードの数 (1-100, デフォルト: 1)",
	"flag_validate":        "パスワード強度を表示",
	"flag_quiet":           "Quiet モード (パスワードのみ出力)",
	"flag_lang":            "言語 (en/zh/zh-TW/ja/ko/es/fr)",
	"flag_output":          "出力形式 (simple/json/csv/table)",
	"flag_exclude_similar": "類似文字を除外 (0, O, 1, l, I, |)",

	"output_password":   "パスワード",
	"output_entropy":    "エントロピー",
	"output_strength":   "強度",
	"output_crack_time": "推定解読時間",

	"strength_very_weak":   "非常に弱い",
	"strength_weak":        "弱い",
	"strength_reasonable":  "普通",
	"strength_strong":      "強い",
	"strength_very_strong": "非常に強い",

	"error_invalid_length":            "無効な長さ: 8 から 128 の間でなければなりません",
	"error_invalid_passphrase_length": "無効な単語数: 4 から 10 の間でなければなりません",
	"error_invalid_count":             "無効な数: 1 から 100 の間でなければなりません",
	"error_invalid_strategy":          "無効な戦略",
	"error_invalid_charset":           "無効な文字セット",
	"error_invalid_output":            "無効な出力形式: simple、json、csv、または table でなければなりません",
}
