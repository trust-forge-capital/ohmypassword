package messages

type MessageMap map[string]string

var EN = MessageMap{
	"root_use":      "ohmypassword - A secure password generator",
	"root_long":     "Generate cryptographically secure passwords with ease.",
	"generate_use":  "Generate a new password",
	"generate_long": "Generate a secure password with customizable options.",
	"version_use":   "Show version information",

	"flag_length":          "Password length (8-128, default: 16)",
	"flag_charset":         "Character set (upper/lower/digit/symbol/all, default: all)",
	"flag_strategy":        "Generation strategy (simple/pronounceable/passphrase/memorable, default: simple)",
	"flag_count":           "Number of passwords to generate (1-100, default: 1)",
	"flag_validate":        "Show password strength",
	"flag_quiet":           "Quiet mode (output password only)",
	"flag_lang":            "Language (en/zh/zh-TW/ja/ko/es/fr)",
	"flag_output":          "Output format (simple/json/csv/table)",
	"flag_exclude_similar": "Exclude similar characters (0, O, 1, l, I, |)",

	"output_password":   "PASSWORD",
	"output_entropy":    "Entropy",
	"output_strength":   "Strength",
	"output_crack_time": "Estimated crack time",

	"strength_very_weak":   "Very Weak",
	"strength_weak":        "Weak",
	"strength_reasonable":  "Reasonable",
	"strength_strong":      "Strong",
	"strength_very_strong": "Very Strong",

	"error_invalid_length":            "Invalid length: must be between 8 and 128",
	"error_invalid_passphrase_length": "Invalid word count: must be between 4 and 10",
	"error_invalid_count":             "Invalid count: must be between 1 and 100",
	"error_invalid_strategy":          "Invalid strategy",
	"error_invalid_charset":           "Invalid charset",
	"error_invalid_output":            "Invalid output format: must be simple, json, csv, or table",
}
