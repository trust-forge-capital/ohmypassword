package messages

var FR = MessageMap{
	"root_use":      "ohmypassword - Générateur de mots de passe sécurisés",
	"root_long":     "Générez facilement des mots de passe cryptographiquement sécurisés.",
	"generate_use":  "Générer un nouveau mot de passe",
	"generate_long": "Générez un mot de passe sécurisé avec des options personnalisables.",
	"version_use":   "Afficher les informations de version",

	"flag_length":          "Longueur du mot de passe (8-128, par défaut: 16)",
	"flag_charset":         "Jeu de caractères (upper/lower/digit/symbol/all, par défaut: all)",
	"flag_strategy":        "Stratégie de génération (simple/pronounceable/passphrase/memorable, par défaut: simple)",
	"flag_count":           "Nombre de mots de passe à générer (1-100, par défaut: 1)",
	"flag_validate":        "Afficher laforce du mot de passe",
	"flag_quiet":           "Mode silencieux (sortir uniquement le mot de passe)",
	"flag_lang":            "Langue (en/zh/zh-TW/ja/ko/es/fr)",
	"flag_output":          "Format de sortie (simple/json/csv/table)",
	"flag_exclude_similar": "Exclure les caractères similaires (0, O, 1, l, I, |)",

	"output_password":   "MOT DE PASSE",
	"output_entropy":    "Entropie",
	"output_strength":   "Force",
	"output_crack_time": "Temps de craquage estimé",

	"strength_very_weak":   "Très faible",
	"strength_weak":        "Faible",
	"strength_reasonable":  "Raisonnable",
	"strength_strong":      "Fort",
	"strength_very_strong": "Très fort",

	"error_invalid_length":            "Longueur invalide: doit être entre 8 et 128",
	"error_invalid_passphrase_length": "Nombre de mots invalide: doit être entre 4 et 10",
	"error_invalid_count":             "Nombre invalide: doit être entre 1 et 100",
	"error_invalid_strategy":          "Stratégie invalide",
	"error_invalid_charset":           "Jeu de caractères invalide",
	"error_invalid_output":            "Format de sortie invalide: doit être simple, json, csv ou table",
}
