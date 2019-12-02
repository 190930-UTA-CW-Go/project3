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

// BotHi = contains the possible "Hi" related outgoing messages
var BotHi = []string{
	"Hi ",
	"Hello ",
	"Hola ",
	"Guten Tag "}

// UserBye = contains the possible "Bye" related incoming messages
var UserBye = map[string]string{
	"goodbye":   "rng",
	"bye":       "rng",
	"peace":     "rng",
	"goodnight": "Good Night ",
}

// BotBye = contains the possible "Bye" related outgoing messages
var BotBye = []string{
	"Goodbye ",
	"Bye ",
	"See You Later "}

// Portfolio = contains the possible "Portfolio" related incoming messages
var Portfolio = map[string]string{
	"startportfolio":    "rng",
	"createportfolio":   "rng",
	"startaportfolio":   "rng",
	"createaportfolio":  "rng",
	"startmyportfolio":  "rng",
	"createmyportfolio": "rng",
}

// BotPortfolio = contains the possible "Portfolio" related outgoing messages
var BotPortfolio = []string{
	"Done, Good Luck!",
	"I got you :)"}

// Status = contains the possible "Status" related incoming messages
var Status = map[string]string{
	"portfoliostatus":     "rng",
	"statusofmyportfolio": "rng",
	"howsmyportfolio":     "rng",
	"statusportfolio":     "rng",
}

// BotStatus = contains the possible "Status" related outgoing messages
var BotStatus = []string{
	"Done, Good Luck!",
	"I got you :)"}
