package messages

var ES = MessageMap{
	"root_use":      "ohmypassword - Generador de contraseñas seguras",
	"root_long":     "Genera contraseñas criptográficamente seguras con facilidad.",
	"generate_use":  "Generar nueva contraseña",
	"generate_long": "Genera una contraseña segura con opciones personalizables.",
	"version_use":   "Mostrar información de versión",

	"flag_length":          "Longitud de contraseña (8-128, predeterminado: 16)",
	"flag_charset":         "Conjunto de caracteres (upper/lower/digit/symbol/all, predeterminado: all)",
	"flag_strategy":        "Estrategia de generación (simple/pronounceable/passphrase, predeterminado: simple)",
	"flag_count":           "Número de contraseñas a generar (1-100, predeterminado: 1)",
	"flag_validate":        "Mostrar fuerza de contraseña",
	"flag_quiet":           "Modo silencioso (solo mostrar contraseña)",
	"flag_lang":            "Idioma (en/zh/zh-TW/ja/ko/es/fr)",
	"flag_output":          "Formato de salida (simple/json/csv/table)",
	"flag_exclude_similar": "Excluir caracteres similares (0, O, 1, l, I, |)",

	"output_password":   "CONTRASEÑA",
	"output_entropy":    "Entropía",
	"output_strength":   "Fortaleza",
	"output_crack_time": "Tiempo estimado de descifrado",

	"strength_very_weak":   "Muy débil",
	"strength_weak":        "Débil",
	"strength_reasonable":  "Razonable",
	"strength_strong":      "Fuerte",
	"strength_very_strong": "Muy fuerte",

	"error_invalid_length":            "Longitud inválida: debe estar entre 8 y 128",
	"error_invalid_passphrase_length": "Cantidad de palabras inválida: debe estar entre 4 y 10",
	"error_invalid_count":             "Cantidad inválida: debe estar entre 1 y 100",
	"error_invalid_strategy":          "Estrategia inválida",
	"error_invalid_charset":           "Conjunto de caracteres inválido",
	"error_invalid_output":            "Formato de salida inválido: debe ser simple, json, csv o table",
}
