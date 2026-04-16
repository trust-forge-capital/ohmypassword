package strategy

import (
	"math"
	"strings"

	"github.com/trust-forge-capital/ohmypassword/internal/generator"
	"github.com/trust-forge-capital/ohmypassword/internal/random"
)

type PassphraseStrategy struct {
	words     []string
	separator string
	rng       random.RNG
}

func NewPassphraseStrategy() *PassphraseStrategy {
	return &PassphraseStrategy{
		words:     defaultWordList,
		separator: "-",
		rng:       random.NewCryptoRNG(),
	}
}

func (s *PassphraseStrategy) Generate(opts *generator.Options) (string, error) {
	wordCount := opts.Length
	if wordCount < 4 {
		wordCount = 4
	}
	if wordCount > 10 {
		wordCount = 10
	}

	result := make([]string, 0, wordCount)
	for i := 0; i < wordCount; i++ {
		word, err := s.randomWord()
		if err != nil {
			return "", err
		}
		result = append(result, word)
	}

	hasDigit := strings.Contains(opts.Charset, "digit") || opts.Charset == "all"
	if hasDigit {
		d, _ := s.rng.Intn(100)
		result = append(result, string(rune('0'+d/10)), string(rune('0'+d%10)))
	}

	hasSymbol := strings.Contains(opts.Charset, "symbol") || opts.Charset == "all"
	if hasSymbol {
		symbols := "!@#$%^&*"
		pos, _ := s.rng.Intn(len(result))
		sym, _ := s.rng.Intn(len(symbols))
		result = insertString(result, pos, string(symbols[sym]))
	}

	return strings.Join(result, s.separator), nil
}

func (s *PassphraseStrategy) randomWord() (string, error) {
	n, err := s.rng.Intn(len(s.words))
	if err != nil {
		return "", err
	}
	return s.words[n], nil
}

func (s *PassphraseStrategy) CalculateEntropy(opts *generator.Options) float64 {
	wordCount := opts.Length
	if wordCount < 4 {
		wordCount = 4
	}
	entropy := float64(wordCount) * math.Log2(float64(len(s.words)))
	entropy += 10
	return entropy
}

func insertString(slice []string, index int, value string) []string {
	result := make([]string, len(slice)+1)
	copy(result, slice[:index])
	result[index] = value
	copy(result[index+1:], slice[index:])
	return result
}

var defaultWordList = []string{
	"acorn", "adventure", "airplane", "anchor", "ancient", "apple", "arctic", "arrow",
	"autumn", "avenue", "bamboo", "banana", "beach", "bear", "beaver", "bee", "birch",
	"bishop", "blizzard", "bluebird", "boat", "boulder", "breeze", "bridge", "brilliant",
	"bronze", "brook", "brother", "brush", "bubble", "bucket", "buddy", "buffalo",
	"butterfly", "cabin", "cable", "cactus", "cafe", "calm", "camel", "camp", "canal",
	"candle", "canyon", "captain", "caravan", "castle", "cedar", "celestial", "chain",
	"chamber", "champion", "channel", "chapter", "charm", "chase", "cherry", "chestnut",
	"chicken", "chief", "chipmunk", "cinnamon", "circle", "clap", "clarity", "clover",
	"cluster", "coach", "coast", "cobblestone", "coconut", "comet", "compass", "copper",
	"coral", "corner", "cosmos", "cottage", "couch", "count", "couple", "court", "craft",
	"crater", "crazy", "creek", "crest", "crisp", "crystal", "culture", "curtain",
	"cycle", "dagger", "dawn", "daylight", "deep", "deer", "delta", "desert", "diamond",
	"dinosaur", "distant", "dock", "dolphin", "donkey", "door", "dove", "dragon",
	"drama", "dream", "dress", "drift", "drink", "drive", "drum", "duck", "dune",
	"eagle", "earth", "eastern", "echo", "eclipse", "edge", "editor", "egg", "eight",
	"elephant", "ember", "emerald", "empire", "empty", "enchanted", "energy", "epic",
	"evergreen", "explorer", "falcon", "famous", "fantasy", "farm", "feather", "ferry",
	"festival", "fever", "field", "fiesta", "figure", "filter", "finch", "firefly",
	"fish", "flame", "flamingo", "flash", "fleet", "flicker", "flight", "flock",
	"flood", "floral", "flower", "fluffy", "flute", "fog", "foliage", "fond", "forest",
	"forger", "fork", "formal", "fort", "forum", "fossil", "fountain", "four", "fox",
	"frame", "frank", "frost", "fruit", "funny", "galaxy", "garden", "garland", "gate",
	"gauntlet", "gaze", "gear", "gem", "genius", "gentle", "ghost", "giant", "gift",
	"giraffe", "girl", "glacier", "glade", "gland", "glare", "glass", "gleam", "globe",
	"glory", "goat", "gold", "golden", "gondola", "goose", "gorge", "grace", "grade",
	"grain", "grape", "graph", "grasp", "grass", "grave", "gravel", "gravy", "gray",
	"grazing", "great", "greed", "green", "greet", "grid", "grief", "grill", "grin",
	"grip", "grit", "groan", "groove", "gross", "group", "grove", "growl", "growth",
	"guard", "guest", "guide", "guild", "guilt", "guitar", "gulf", "gull", "gum",
	"gun", "guppy", "guru", "gush", "gust", "gutter", "habitat", "hail", "hair",
	"hairy", "half", "hall", "halo", "halt", "ham", "hammer", "hand", "handle", "hang",
	"happy", "harbor", "hard", "hare", "harm", "harp", "harvest", "hash", "haste",
	"hat", "hatch", "haven", "hawk", "haze", "hazel", "head", "heal", "health", "heap",
	"hear", "heart", "heat", "heath", "heaven", "heavy", "hedge", "heel", "height",
	"heir", "helix", "hell", "hello", "helm", "helmet", "help", "hemp", "herb", "herd",
	"hero", "heron", "herring", "hertz", "hex", "hey", "hiatus", "hibiscus", "hide",
	"high", "hill", "him", "hippo", "hire", "hiss", "hit", "hive", "hoard", "hobby",
	"hoist", "hollow", "holy", "home", "honest", "honey", "hood", "hoof", "hook",
	"hoop", "hop", "hope", "horde", "horn", "hornet", "horror", "horse", "host",
	"hotel", "hound", "hour", "house", "hover", "howl", "hub", "huddle", "huge",
	"hull", "human", "humid", "humor", "hump", "humus", "hunch", "hundred", "hunger",
	"hunt", "hurdle", "hurl", "hurricane", "hush", "husk", "hut", "hybrid", "hydra",
	"hydro", "hyena", "hymn", "hyphen", "ice", "icon", "idea", "igloo", "ignore",
	"ill", "image", "imbue", "imp", "impulse", "in", "inch", "incline", "index",
	"indigo", "inept", "infant", "infer", "inflame", "ink", "inlet", "inn", "inner",
	"input", "inter", "intro", "inward", "ion", "iris", "iron", "island", "isolate",
	"ivory", "ivy", "jackal", "jade", "jaguar", "jail", "jam", "jar", "jaw", "jazz",
	"jeans", "jelly", "jet", "jewel", "jiffy", "job", "jog", "join", "joke", "jolt",
	"journal", "joy", "judge", "jug", "juice", "jump", "jungle", "junior", "junk",
	"jury", "just", "kale", "kangaroo", "kayak", "kazoo", "kebab", "keen", "keep",
	"kelp", "kennel", "kernel", "key", "kick", "kid", "kill", "kiln", "kilt", "kind",
	"king", "kiosk", "kiss", "kit", "kitchen", "kite", "kitten", "kiwi", "knee",
	"kneel", "knife", "knight", "knit", "knob", "knock", "knot", "know", "koala",
	"kraken", "label", "labor", "lace", "lack", "lady", "lake", "lamb", "lamp",
	"land", "lane", "lap", "large", "laser", "lasso", "last", "late", "laugh",
	"launch", "lava", "lawn", "lawsuit", "layer", "lead", "leaf", "leak", "lean",
	"leap", "learn", "lease", "leather", "leave", "lecture", "ledge", "left", "leg",
	"lemon", "lemur", "lend", "length", "lens", "leopard", "lesson", "letter", "level",
	"lever", "liar", "liberty", "librate", "lichen", "lick", "lie", "life", "lift",
	"light", "like", "limb", "lime", "limit", "limp", "line", "link", "lion", "list",
	"live", "liver", "lizard", "load", "loan", "lobby", "local", "lock", "locus",
	"lodge", "loft", "log", "logic", "login", "lonely", "long", "look", "loop",
	"loose", "lord", "lose", "loss", "lost", "lot", "lotus", "loud", "lounge", "love",
	"low", "loyal", "luck", "lunar", "lunch", "lure", "lush", "lust", "lynx", "lyric",
	"mace", "machine", "mad", "magic", "magma", "mail", "main", "make", "mammal",
	"man", "mandate", "mango", "manor", "maple", "marble", "march", "mare", "mark",
	"market", "mars", "marsh", "match", "mate", "material", "matrix", "maze", "meal",
	"mean", "meant", "meat", "mechanic", "medal", "medial", "medic", "meek", "meet",
	"mellow", "melon", "melt", "member", "memo", "memory", "men", "mend", "mental",
	"mentor", "menu", "mercury", "merit", "merry", "mess", "metal", "meteor", "meter",
	"method", "metro", "mice", "midst", "might", "mild", "mile", "milk", "mill",
	"mimic", "mind", "mine", "mint", "minus", "minute", "mire", "mirror", "mirth",
	"misery", "miss", "mist", "mite", "mix", "moan", "moat", "mob", "mock", "mode",
	"model", "modem", "modern", "modest", "modify", "moist", "mold", "moment", "money",
	"monitor", "monk", "monster", "month", "mood", "moon", "moor", "moose", "mop",
	"moral", "more", "morning", "mosaic", "mosquito", "most", "moth", "mother", "motion",
	"motive", "motor", "motto", "mount", "mountain", "mouse", "mouth", "move", "much",
	"mud", "muffin", "mug", "mule", "mull", "multi", "mural", "murder", "muse",
	"museum", "mushroom", "music", "must", "mustard", "mute", "mutter", "mutual",
	"muzzle", "myth", "nail", "name", "nape", "napkin", "narrow", "nasty", "nation",
	"nature", "naval", "navel", "near", "neat", "neck", "need", "needle", "neglect",
	"neighbor", "neither", "nephew", "nerve", "nest", "net", "network", "neutral",
	"new", "news", "next", "nice", "niche", "niece", "night", "nine", "noble", "noise",
	"nomad", "noon", "norm", "north", "nose", "notch", "note", "nothing", "notice",
	"notify", "notion", "noun", "novel", "nudge", "numb", "nurse", "nut", "nylon",
	"oak", "oasis", "oat", "obey", "object", "obtain", "ocean", "octopus", "odd",
	"offense", "offer", "office", "often", "oil", "old", "olive", "omen", "omit",
	"once", "one", "onion", "only", "onset", "onto", "open", "operate", "opinion",
	"opt", "optic", "option", "opus", "orange", "orbit", "orca", "orchid", "order",
	"ore", "organ", "orient", "origin", "ornament", "orphan", "ostrich", "other",
	"otter", "ought", "ounce", "our", "out", "outer", "output", "outset", "oval",
	"oven", "over", "owl", "own", "owner", "oxide", "ozone", "pace", "pack", "packet",
	"pact", "page", "pail", "pain", "paint", "pair", "pale", "palm", "pan", "panda",
	"panel", "panic", "pant", "pants", "paper", "parade", "parcel", "park", "parody",
	"parrot", "part", "pass", "past", "paste", "patch", "path", "patio", "patrol",
	"pattern", "pause", "pave", "pawn", "pay", "peace", "peach", "peak", "pearl",
	"pedal", "peel", "peer", "pelican", "pen", "penalty", "pencil", "pendulum", "penguin",
	"penny", "people", "pepper", "perch", "peril", "period", "perish", "permit",
	"person", "pest", "petal", "petite", "petrol", "phase", "phone", "photo", "piano",
	"pick", "pickle", "picnic", "picture", "pie", "piece", "pig", "pigeon", "piggy",
	"pile", "pill", "pillow", "pilot", "pin", "pine", "ping", "pink", "pint", "pipe",
	"pirate", "pistol", "pit", "pitch", "pizza", "place", "plaid", "plain", "plan",
	"plane", "planet", "plant", "plate", "play", "plaza", "plea", "please", "pledge",
	"plenty", "plot", "plow", "plug", "plum", "plump", "plunge", "plus", "pocket",
	"poem", "poet", "point", "poison", "polar", "pole", "police", "pond", "pony",
	"pool", "poor", "pop", "popcorn", "porch", "pork", "port", "pose", "position",
	"positive", "possible", "post", "pot", "potato", "potter", "pouch", "pound",
	"powder", "power", "praise", "pray", "preach", "precious", "predict", "prefer",
	"prefix", "press", "pretty", "prey", "price", "pride", "prime", "print", "prior",
	"prism", "prize", "probe", "prodigy", "produce", "product", "profit", "program",
	"project", "prom", "promise", "proof", "proper", "prophet", "prose", "protect",
	"proud", "prove", "provide", "prowl", "prune", "psalm", "pub", "pulse", "puma",
	"pump", "punch", "pupil", "puppet", "puppy", "pure", "purple", "purpose", "purse",
	"push", "put", "puzzle", "pyramid", "quack", "quail", "quake", "quality", "quantity",
	"quark", "quart", "quartz", "queen", "query", "quest", "queue", "quick", "quiet",
	"quilt", "quirk", "quota", "quote", "rabbit", "race", "rack", "radar", "radiant",
	"radio", "raft", "rage", "raid", "rail", "rain", "raise", "rally", "ranch", "range",
	"rank", "rapid", "rare", "rash", "rat", "rate", "raven", "raw", "ray", "razor",
	"reach", "react", "read", "reader", "ready", "realm", "rear", "reason", "rebel",
	"build", "recall", "receive", "recipe", "record", "recover", "recruit", "rectify",
	"recycle", "red", "reed", "reef", "reel", "refer", "reform", "refuge", "refuse",
	"regal", "regard", "regime", "region", "regret", "reject", "relate", "relax",
	"relay", "release", "relief", "rely", "remain", "remark", "remedy", "remember",
	"remind", "remote", "remove", "render", "renew", "rent", "repair", "repeat",
	"repel", "reply", "report", "rescue", "resemble", "resist", "resort", "resource",
	"respect", "respond", "rest", "result", "resume", "retail", "retain", "retire",
	"return", "reveal", "review", "revise", "revive", "reward", "rhino", "rhyme",
	"rhythm", "rib", "ribbon", "rice", "rich", "ride", "ridge", "rifle", "right",
	"rigid", "ring", "riot", "rip", "ripe", "rise", "risk", "ritual", "rival", "river",
	"road", "roam", "roar", "roast", "rob", "robot", "rock", "rocket", "rode", "role",
	"roll", "romance", "roof", "room", "root", "rope", "rose", "rotate", "rotor",
	"rotten", "rough", "round", "route", "royal", "rub", "rubber", "ruby", "rude",
	"rug", "ruin", "rule", "rumor", "run", "runway", "rural", "rush", "rust", "rut",
	"sack", "sad", "saddle", "sadness", "safe", "safety", "sail", "saint", "salad",
	"sale", "sales", "salmon", "salt", "salute", "same", "sand", "sane", "sash",
	"sat", "satin", "satire", "satisfy", "saturn", "sauce", "sauna", "save", "say",
	"scale", "scam", "scan", "scare", "scarf", "scary", "scene", "scent", "school",
	"science", "scissors", "scoop", "scope", "score", "scorn", "scout", "scrap",
	"scratch", "scream", "screen", "screw", "script", "scroll", "scrub", "sea",
	"seal", "seam", "search", "season", "seat", "second", "secret", "section", "sector",
	"secure", "see", "seed", "seek", "seem", "seize", "seldom", "select", "self",
	"sell", "semester", "senate", "send", "senior", "sense", "sensor", "sent",
	"sentence", "separate", "sequence", "serial", "series", "serious", "servant",
	"serve", "service", "set", "settle", "setup", "seven", "sever", "severe", "sew",
	"shade", "shadow", "shaft", "shake", "shallow", "shame", "shape", "share", "shark",
	"sharp", "shave", "she", "shed", "sheep", "sheet", "shelf", "shell", "shelter",
	"shepherd", "shield", "shift", "shine", "shiny", "ship", "shiver", "shock",
	"shoe", "shoot", "shop", "short", "shot", "shoulder", "shout", "shove", "show",
	"shower", "shred", "shrimp", "shrine", "shrink", "shrug", "shuffle", "shun",
	"shut", "shy", "sibling", "sick", "side", "siege", "sigh", "sight", "sign",
	"signal", "silence", "silent", "silk", "silly", "silver", "similar", "simple",
	"simulate", "sin", "since", "sing", "singer", "single", "sink", "sip", "sir",
	"siren", "sister", "sit", "site", "situation", "six", "sixth", "size", "skate",
	"sketch", "ski", "skill", "skin", "skirt", "skull", "slab", "slack", "slain",
	"slang", "slap", "slate", "slave", "sleek", "sleep", "sleet", "slice", "slide",
	"slim", "sling", "slip", "slit", "sliver", "slope", "slot", "slow", "slump",
	"small", "smart", "smash", "smell", "smile", "smoke", "snack", "snail", "snake",
	"snap", "snare", "snarl", "sneak", "snow", "snug", "soak", "soap", "soar",
	"soccer", "social", "sock", "soda", "sofa", "soft", "soil", "soldier", "sole",
	"solid", "solve", "some", "son", "sonar", "song", "soon", "sore", "sorrow",
	"sort", "soul", "sound", "soup", "sour", "source", "south", "space", "spade",
	"span", "spare", "spark", "sparrow", "speak", "speaker", "spear", "spec", "special",
	"speck", "speed", "spell", "spend", "spent", "sphere", "spice", "spider", "spike",
	"spin", "spine", "spiral", "spirit", "spit", "splash", "split", "spoil", "spoke",
	"sponge", "spoon", "sport", "spot", "spouse", "spray", "spread", "spring", "spy",
	"squad", "square", "squat", "squid", "stack", "staff", "stage", "stain", "stair",
	"stake", "stale", "stall", "stamp", "stand", "star", "stare", "stark", "start",
	"starve", "state", "static", "statue", "status", "stay", "steak", "steal", "steam",
	"steel", "steep", "steer", "stem", "step", "stereo", "stick", "stiff", "still",
	"sting", "stink", "stock", "stole", "stomach", "stone", "stool", "stop", "store",
	"storm", "story", "stove", "straight", "strain", "strange", "strap", "straw",
	"stray", "street", "stress", "stretch", "strict", "stride", "strife", "strike",
	"string", "strip", "strive", "stroke", "strong", "struck", "struggle", "strum",
	"strut", "stub", "student", "study", "stuff", "stumble", "stump", "stun", "stunt",
	"style", "subject", "submit", "subway", "success", "such", "sudden", "sue", "sugar",
	"suggest", "suit", "summer", "summit", "sun", "sung", "sunny", "sunset", "super",
	"supply", "support", "suppose", "sure", "surface", "surge", "surprise", "surround",
	"survey", "suspect", "sustain", "swallow", "swamp", "swan", "swap", "swarm",
	"sway", "swear", "sweat", "sweep", "sweet", "swell", "swept", "swift", "swim",
	"swing", "switch", "sword", "symbol", "sympathy", "system", "table", "tablet",
	"tack", "tactic", "tag", "tail", "take", "tale", "talk", "tall", "tame", "tank",
	"tap", "tape", "target", "task", "taste", "tattoo", "taxi", "tea", "teach", "team",
	"tear", "tech", "teeth", "tell", "temp", "tempo", "tend", "tennis", "tense", "tent",
	"term", "terrain", "test", "text", "than", "thank", "that", "theft", "theme", "then",
	"theory", "there", "thermal", "these", "thick", "thief", "thigh", "thing", "think",
	"third", "thirst", "this", "thorn", "those", "though", "thread", "threat", "three",
	"thrill", "thrive", "throat", "throne", "through", "throw", "thumb", "thunder",
	"thus", "ticket", "tide", "tidy", "tie", "tiger", "tight", "tile", "till", "time",
	"tin", "tiny", "tip", "tire", "tissue", "title", "toad", "toast", "today", "toe",
	"together", "toilet", "token", "told", "toll", "tomato", "tomb", "tone", "tongue",
	"tonic", "tonight", "too", "tool", "tooth", "top", "topic", "torch", "tornado",
	"tortoise", "toss", "total", "totem", "touch", "tough", "tour", "tournament", "toward",
	"tower", "town", "toy", "trace", "track", "trade", "trail", "train", "trait",
	"tram", "trap", "trash", "travel", "tray", "treat", "tree", "trek", "trend", "trial",
	"tribe", "trick", "tried", "trigger", "trim", "trip", "troop", "trophy", "tropical",
	"trouble", "truck", "true", "trumpet", "trunk", "trust", "truth", "try", "tub",
	"tube", "tulip", "tumor", "tune", "tunic", "turkey", "turn", "turtle", "tutor",
	"twelve", "twenty", "twice", "twig", "twin", "twist", "two", "type", "ugly",
	"ulcer", "ultra", "unable", "unaware", "uncle", "uncover", "under", "unfair",
	"unfold", "unhappy", "uniform", "union", "unique", "unit", "universe", "unknown",
	"unlock", "until", "unusual", "unveil", "update", "upgrade", "uphold", "upon",
	"upper", "upset", "urban", "urge", "usage", "use", "used", "useless", "user",
	"usual", "utility", "utter", "vague", "valid", "valley", "valor", "value", "valve",
	"van", "vanish", "variable", "variety", "various", "vase", "vast", "vegetable",
	"vehicle", "vein", "velvet", "vendor", "venom", "vent", "venue", "verb", "verify",
	"verse", "very", "vessel", "vest", "veteran", "vial", "vibrate", "vice", "victim",
	"victory", "video", "view", "village", "vine", "violin", "virtual", "virus", "visit",
	"visual", "vital", "vivid", "vocal", "voice", "void", "volcano", "volume", "vote",
	"vowel", "voyage", "wade", "waffle", "wage", "wagon", "waist", "wait", "wake",
	"walk", "wall", "walnut", "want", "war", "ward", "warm", "warn", "warp", "wart",
	"wary", "wash", "wasp", "waste", "watch", "water", "wave", "wax", "way", "weak",
	"wealth", "weapon", "wear", "weasel", "weather", "web", "wedding", "wedge", "weed",
	"week", "weep", "weigh", "weird", "welcome", "well", "west", "wet", "whale", "what",
	"wheat", "wheel", "when", "where", "whip", "whisper", "whistle", "white", "whole",
	"why", "wicked", "wide", "widow", "width", "wife", "wild", "will", "win", "wind",
	"window", "wine", "wing", "wink", "winner", "winter", "wire", "wisdom", "wise",
	"wish", "witch", "within", "without", "witness", "wolf", "woman", "wonder", "wood",
	"wool", "word", "work", "world", "worm", "worn", "worry", "worth", "wrap", "wreck",
	"wrist", "write", "wrong", "yard", "yarn", "year", "yell", "yellow", "yes", "yet",
	"yoga", "young", "your", "youth", "zebra", "zero", "zone", "zoo",
}