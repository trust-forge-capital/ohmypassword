package validator

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
		return "Very Weak"
	case LevelWeak:
		return "Weak"
	case LevelReasonable:
		return "Reasonable"
	case LevelStrong:
		return "Strong"
	case LevelVeryStrong:
		return "Very Strong"
	default:
		return "Unknown"
	}
}