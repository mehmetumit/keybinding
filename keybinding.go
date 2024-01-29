package keybinding

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/awesome-gocui/gocui"
)

var translate = map[string]string{
	"/":        "Slash",
	"\\":       "Backslash",
	"[":        "LsqBracket",
	"]":        "RsqBracket",
	"_":        "Underscore",
	"escape":   "Esc",
	"~":        "Tilde",
	"pageup":   "Pgup",
	"pagedown": "Pgdn",
	"pgup":     "Pgup",
	"pgdown":   "Pgdn",
	"up":       "ArrowUp",
	"down":     "ArrowDown",
	"right":    "ArrowRight",
	"left":     "ArrowLeft",
	"ctl":      "Ctrl",
}

var display = map[string]string{
	"Slash":      "/",
	"Backslash":  "\\",
	"LsqBracket": "[",
	"RsqBracket": "]",
	"Underscore": "_",
	"Tilde":      "~",
	"Ctrl":       "^",
}

var supportedKeybindings = map[string]gocui.Key{
	"KeyF1":             gocui.KeyF1,
	"KeyF2":             gocui.KeyF2,
	"KeyF3":             gocui.KeyF3,
	"KeyF4":             gocui.KeyF4,
	"KeyF5":             gocui.KeyF5,
	"KeyF6":             gocui.KeyF6,
	"KeyF7":             gocui.KeyF7,
	"KeyF8":             gocui.KeyF8,
	"KeyF9":             gocui.KeyF9,
	"KeyF10":            gocui.KeyF10,
	"KeyF11":            gocui.KeyF11,
	"KeyF12":            gocui.KeyF12,
	"KeyInsert":         gocui.KeyInsert,
	"KeyDelete":         gocui.KeyDelete,
	"KeyHome":           gocui.KeyHome,
	"KeyEnd":            gocui.KeyEnd,
	"KeyPgup":           gocui.KeyPgup,
	"KeyPgdn":           gocui.KeyPgdn,
	"KeyArrowUp":        gocui.KeyArrowUp,
	"KeyArrowDown":      gocui.KeyArrowDown,
	"KeyArrowLeft":      gocui.KeyArrowLeft,
	"KeyArrowRight":     gocui.KeyArrowRight,
	"KeyCtrlTilde":      gocui.KeyCtrlTilde,
	"KeyCtrl2":          gocui.KeyCtrl2,
	"KeyCtrlSpace":      gocui.KeyCtrlSpace,
	"KeyCtrlA":          gocui.KeyCtrlA,
	"KeyCtrlB":          gocui.KeyCtrlB,
	"KeyCtrlC":          gocui.KeyCtrlC,
	"KeyCtrlD":          gocui.KeyCtrlD,
	"KeyCtrlE":          gocui.KeyCtrlE,
	"KeyCtrlF":          gocui.KeyCtrlF,
	"KeyCtrlG":          gocui.KeyCtrlG,
	"KeyBackspace":      gocui.KeyBackspace,
	"KeyCtrlH":          gocui.KeyCtrlH,
	"KeyTab":            gocui.KeyTab,
	"KeyCtrlI":          gocui.KeyCtrlI,
	"KeyCtrlJ":          gocui.KeyCtrlJ,
	"KeyCtrlK":          gocui.KeyCtrlK,
	"KeyCtrlL":          gocui.KeyCtrlL,
	"KeyEnter":          gocui.KeyEnter,
	"KeyCtrlM":          gocui.KeyCtrlM,
	"KeyCtrlN":          gocui.KeyCtrlN,
	"KeyCtrlO":          gocui.KeyCtrlO,
	"KeyCtrlP":          gocui.KeyCtrlP,
	"KeyCtrlQ":          gocui.KeyCtrlQ,
	"KeyCtrlR":          gocui.KeyCtrlR,
	"KeyCtrlS":          gocui.KeyCtrlS,
	"KeyCtrlT":          gocui.KeyCtrlT,
	"KeyCtrlU":          gocui.KeyCtrlU,
	"KeyCtrlV":          gocui.KeyCtrlV,
	"KeyCtrlW":          gocui.KeyCtrlW,
	"KeyCtrlX":          gocui.KeyCtrlX,
	"KeyCtrlY":          gocui.KeyCtrlY,
	"KeyCtrlZ":          gocui.KeyCtrlZ,
	"KeyEsc":            gocui.KeyEsc,
	"KeyCtrlLsqBracket": gocui.KeyCtrlLsqBracket,
	"KeyCtrl3":          gocui.KeyCtrl3,
	"KeyCtrl4":          gocui.KeyCtrl4,
	"KeyCtrlBackslash":  gocui.KeyCtrlBackslash,
	"KeyCtrl5":          gocui.KeyCtrl5,
	"KeyCtrlRsqBracket": gocui.KeyCtrlRsqBracket,
	"KeyCtrl6":          gocui.KeyCtrl6,
	"KeyCtrl7":          gocui.KeyCtrl7,
	"KeyCtrlSlash":      gocui.KeyCtrlSlash,
	"KeyCtrlUnderscore": gocui.KeyCtrlUnderscore,
	"KeySpace":          gocui.KeySpace,
	"KeyBackspace2":     gocui.KeyBackspace2,
	"KeyCtrl8":          gocui.KeyCtrl8,
	"MouseLeft":         gocui.MouseLeft,
	"MouseMiddle":       gocui.MouseMiddle,
	"MouseRight":        gocui.MouseRight,
	"MouseRelease":      gocui.MouseRelease,
	"MouseWheelUp":      gocui.MouseWheelUp,
	"MouseWheelDown":    gocui.MouseWheelDown,
}

// Key contains all relevant information about the key
type Key struct {
	Value    interface{}
	Modifier gocui.Modifier
	Tokens   []string
}

// Parse turns the input string into an actual Key.
// Returns an error when something goes wrong
func Parse(input string) (Key, error) {
	f := func(c rune) bool { return unicode.IsSpace(c) || c == '+' }
	tokens := strings.FieldsFunc(input, f)

	if len(tokens) == 1 && len([]rune(tokens[0])) == 1 {
		single := []rune(tokens[0])[0]
		if unicode.IsPrint(single) {
			return Key{single, gocui.ModNone, tokens}, nil
		}
	}

	var normalizedTokens []string
	var modifier = gocui.ModNone

	for _, token := range tokens {
		normalized := token

		if value, exists := translate[normalized]; exists {
			normalized = value
		} else {
			normalized = strings.Title(normalized)
		}

		if normalized == "Alt" {
			modifier = gocui.ModAlt
			continue
		}

		if len(normalized) == 1 {
			normalizedTokens = append(normalizedTokens, strings.ToUpper(normalized))
			continue
		}

		normalizedTokens = append(normalizedTokens, normalized)
	}

	lookup := strings.Join(normalizedTokens, "")
	if !strings.Contains(lookup, "Mouse") {
		lookup = "Key" + strings.Join(normalizedTokens, "")
	}

	if key, exists := supportedKeybindings[lookup]; exists {
		return Key{key, modifier, normalizedTokens}, nil
	}

	if modifier != gocui.ModNone {
		return Key{0, modifier, normalizedTokens}, fmt.Errorf("unsupported keybinding: %s (+%+v)", lookup, modifier)
	}

	return Key{0, modifier, normalizedTokens}, fmt.Errorf("unsupported keybinding: %s", lookup)
}

// ParseAll parses all strings to a Key.
// Returns an error when something goes wrong.
func ParseAll(input string) ([]Key, error) {
	ret := make([]Key, 0)
	for _, value := range strings.Split(input, ",") {
		key, err := Parse(value)
		if err != nil {
			return nil, fmt.Errorf("could not parse keybinding '%s' from request '%s': %+v", value, input, err)
		}
		ret = append(ret, key)
	}
	if len(ret) == 0 {
		return nil, fmt.Errorf("must have at least one keybinding")
	}
	return ret, nil
}

// MustParse parses the input string to a key but instead an
// error, it panics when things go wrong.
// This forces the caller to react to an error
func MustParse(input string) Key {
	if key, err := Parse(input); err != nil {
		panic(err)
	} else {
		return key
	}
}

// MustParseAll parses all the input strings to a key but instead an
// error, it panics when things go wrong.
// This forces the caller to react to an error
func MustParseAll(input string) []Key {
	if key, err := ParseAll(input); err != nil {
		panic(err)
	} else {
		return key
	}
}

// String returns the Key in string form
func (key Key) String() string {
	displayTokens := make([]string, 0)
	prefix := ""
	for _, token := range key.Tokens {
		if token == "Ctrl" {
			prefix = "^"
			continue
		}
		if value, exists := display[token]; exists {
			token = value
		}
		displayTokens = append(displayTokens, token)
	}
	return prefix + strings.Join(displayTokens, "+")
}
