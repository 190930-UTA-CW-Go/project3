package bot

// UserHi = contains the possible "Hi" related incoming messages
var UserHi = map[string]string{
	"hi":            "rng",
	"hello":         "rng",
	"hey":           "rng",
	"yo":            "rng",
	"goodmorning":   "Good Morning ",
	"goodafternoon": "Good Afternoon ",
}

// UserBye = contains the possible "Bye" related incoming messages
var UserBye = map[string]string{
	"goodbye":   "rng",
	"bye":       "rng",
	"peace":     "rng",
	"goodnight": "Good Night ",
}

// BotHi = contains the possible "Hi" related outgoing messages
var BotHi = []string{
	"Hi ",
	"Hello ",
	"Hola ",
	"Guten Tag "}

// BotBye = contains the possible "Bye" related outgoing messages
var BotBye = []string{
	"Goodbye ",
	"Bye ",
	"See You Later "}
