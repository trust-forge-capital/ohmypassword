package i18n

import (
	"fmt"
	"os"
	"sync"

	"github.com/trust-forge-capital/ohmypassword/internal/i18n/messages"
)

var (
	currentLang  string
	mu           sync.RWMutex
	translations map[string]map[string]string
)

func init() {
	translations = map[string]map[string]string{
		"en": messages.EN,
		"zh": messages.ZH,
		"zh-TW": messages.ZHTW,
		"ja": messages.JA,
		"ko": messages.KO,
		"es": messages.ES,
		"fr": messages.FR,
	}
}

func SetLanguage(lang string) {
	mu.Lock()
	defer mu.Unlock()

	if lang == "" {
		lang = detectSystemLanguage()
	}

	if _, ok := translations[lang]; !ok {
		lang = "en"
	}

	currentLang = lang
}

func detectSystemLanguage() string {
	lang := os.Getenv("LANG")
	if lang == "" {
		return "en"
	}

	lang = lang[:2]
	if lang == "zh" {
		lang = "zh"
	}

	return lang
}

func T(key string) string {
	mu.RLock()
	defer mu.RUnlock()

	lang := currentLang
	if lang == "" {
		lang = "en"
	}

	if trans, ok := translations[lang]; ok {
		if msg, ok := trans[key]; ok {
			return msg
		}
	}

	if trans, ok := translations["en"]; ok {
		if msg, ok := trans[key]; ok {
			return msg
		}
	}

	return key
}

func TFormat(key string, args ...interface{}) string {
	return fmt.Sprintf(T(key), args...)
}

func GetCurrentLanguage() string {
	mu.RLock()
	defer mu.RUnlock()
	return currentLang
}

func GetSupportedLanguages() []string {
	mu.RLock()
	defer mu.RUnlock()

	languages := make([]string, 0, len(translations))
	for lang := range translations {
		languages = append(languages, lang)
	}
	return languages
}