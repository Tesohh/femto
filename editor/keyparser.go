package editor

import (
	"strings"

	"github.com/gdamore/tcell/v2"
)

func SplitHumanKeySequence(s string) []string {
	keys := []string{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '<' {
			nextAngle := strings.Index(s[i:], ">")
			keys = append(keys, s[i:nextAngle+2])
			i = nextAngle + 1
		} else {
			keys = append(keys, string(c))
		}
	}
	return keys
}

func ParseHumanKey(key string) (InternalKey, error) {
	// # Rules
	// single runes are just runes
	// things in <> are special keys OR modifier keys OR modifier special keys (eg <C-F10>)
	// use recursion to parse stuff like F10, then add modifiers to it

	isSpecial := strings.HasPrefix(key, "<") && strings.HasSuffix(key, ">")

	if len(key) == 1 {
		return InternalKey{
			Key:  tcell.KeyRune,
			Rune: rune(key[0]),
		}, nil
	} else if isSpecial && key[2] == '-' {
		key = strings.Trim(key, "<>")
		sections := strings.Split(key, "-")
		finalSection := sections[len(sections)-1]

		// parse the second part of the key through recursion
		var parsedKey InternalKey
		var err error
		if len(finalSection) == 0 {
			parsedKey, err = ParseHumanKey(finalSection)
		} else {
			parsedKey, err = ParseHumanKey("<" + finalSection + ">")
		}
		if err != nil {
			return InternalKey{}, err
		}

		// add modifiers
		for _, v := range sections[:len(sections)] {
			v = strings.ToLower(v)
			if v == "c" {
				// noop; should be handled by `else if isSpecial` case
			} else if v == "s" {

			}
		}

	} else if isSpecial {
		key = strings.Trim(key, "<>")
	}

	return InternalKey{}, ErrKeyUnparsable
}

type InternalKey struct {
	Key       tcell.Key
	Rune      rune
	Modifiers tcell.ModMask
}
