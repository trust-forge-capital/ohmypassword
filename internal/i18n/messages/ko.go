package messages

var KO = MessageMap{
	"root_use":      "ohmypassword - 안전한 비밀번호 생성기",
	"root_long":     "암호학적으로 안전한 비밀번호를 쉽게 생성합니다.",
	"generate_use":  "새 비밀번호 생성",
	"generate_long": "사용자 정의 옵션으로 안전한 비밀번호를 생성합니다.",
	"check_use":     "비밀번호 강도 확인",
	"check_long":    "하나 이상의 비밀번호 강도를 확인합니다.",
	"version_use":   "버전 정보 표시",

	"flag_length":          "비밀번호 길이 (8-128, 기본값: 16)",
	"flag_charset":         "문자 집합 (upper/lower/digit/symbol/all, 기본값: all)",
	"flag_strategy":        "생성 전략 (simple/pronounceable/passphrase/memorable/segmented, 기본값: simple)",
	"flag_count":           "생성할 비밀번호 수 (1-100, 기본값: 1)",
	"flag_validate":        "비밀번호 강도 표시",
	"flag_quiet":           "자동 모드 (비밀번호만 출력)",
	"flag_lang":            "언어 (en/zh/zh-TW/ja/ko/es/fr)",
	"flag_output":          "출력 형식 (simple/json/csv/table)",
	"flag_exclude_similar": "유사 문자 제외 (0, O, 1, l, I, |)",

	"output_password":   "비밀번호",
	"output_entropy":    "엔트로피",
	"output_strength":   "강도",
	"output_crack_time": "추정 해독 시간",

	"strength_very_weak":   "매우 약함",
	"strength_weak":        "약함",
	"strength_reasonable":  "보통",
	"strength_strong":      "강함",
	"strength_very_strong": "매우 강함",

	"error_invalid_length":            "잘못된 길이: 8에서 128 사이여야 합니다",
	"error_invalid_passphrase_length": "잘못된 단어 수: 4에서 10 사이여야 합니다",
	"error_invalid_count":             "잘못된 수: 1에서 100 사이여야 합니다",
	"error_invalid_strategy":          "잘못된 전략",
	"error_invalid_charset":           "잘못된 문자 집합",
	"error_invalid_output":            "잘못된 출력 형식: simple, json, csv 또는 table이어야 합니다",
}
