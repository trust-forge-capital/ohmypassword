package validator

import "github.com/trust-forge-capital/ohmypassword/internal/i18n"

type StrengthLevel string

const (
	LevelVeryWeak   StrengthLevel = "very_weak"
	LevelWeak       StrengthLevel = "weak"
	LevelReasonable StrengthLevel = "reasonable"
	LevelStrong     StrengthLevel = "strong"
	LevelVeryStrong StrengthLevel = "very_strong"
)

func GetMinEntropyForLevel(level StrengthLevel) int {
	switch level {
	case LevelVeryWeak:
		return 0
	case LevelWeak:
		return 28
	case LevelReasonable:
		return 36
	case LevelStrong:
		return 60
	case LevelVeryStrong:
		return 80
	default:
		return 0
	}
}

func GetDisplayName(level StrengthLevel) string {
	switch level {
	case LevelVeryWeak:
		return i18n.T("strength_very_weak")
	case LevelWeak:
		return i18n.T("strength_weak")
	case LevelReasonable:
		return i18n.T("strength_reasonable")
	case LevelStrong:
		return i18n.T("strength_strong")
	case LevelVeryStrong:
		return i18n.T("strength_very_strong")
	default:
		return "Unknown"
	}
}
