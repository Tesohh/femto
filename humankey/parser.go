package humankey

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/gdamore/tcell/v2"
)

func Parse(s string) (InternalKey, error) {
	var key InternalKey
	var err error

	sections := strings.Split(s, "+")

	// special ctrl character case
	var isCtrl, isInCtrlKeys bool
	var trimmedRunicPart string
	var ctrlKey tcell.Key
	if len(sections) == 2 {
		isCtrl = sections[0] == "ctrl"
		trimmedRunicPart = strings.Trim(sections[1], " ")
		ctrlKey, isInCtrlKeys = ctrlKeysMap[trimmedRunicPart] // Case sensitive as ctrl+shift should be handeld with modmask
	}

	if isCtrl && isInCtrlKeys {
		var r rune
		if trimmedRunicPart == "space" {
			r = ' '
		} else {
			r = rune(trimmedRunicPart[0])
		}
		key = InternalKey{Key: ctrlKey, Rune: r, ModMask: tcell.ModCtrl}
	} else { // regular modmask case
		// add modmasks
		for _, section := range sections[:len(sections)-1] {
			switch section {
			case "ctrl":
				key.ModMask |= tcell.ModCtrl
			case "alt", "opt":
				key.ModMask |= tcell.ModAlt
			case "meta", "gui", "win", "cmd":
				key.ModMask |= tcell.ModMeta
			case "shift":
				key.ModMask |= tcell.ModShift
			}
		}

		// add rune/key checking if it exists in the special key map.
		lastSection := strings.Trim(sections[len(sections)-1], " ")
		if specialKey, ok := specialKeysMap[lastSection]; ok {
			key.Key = specialKey
		} else if len(lastSection) == 1 {
			key.Rune = unicode.ToLower(rune(lastSection[0]))
			key.Key = tcell.KeyRune

			if unicode.IsUpper(rune(lastSection[0])) && key.ModMask&tcell.ModShift == 0 { // check for shift
				key.ModMask |= tcell.ModShift
			}
		} else if lastSection == "space" {
			key.Rune = ' '
			key.Key = tcell.KeyRune
		} else {
			return InternalKey{}, fmt.Errorf("unknown key `%s`", lastSection)
		}

	}
	return key, err
}

func ParseSequence(s string) ([]InternalKey, error) {
	keys := strings.Split(s, " ")
	internalKeys := []InternalKey{}
	for _, v := range keys {
		key, err := Parse(v)
		if err != nil {
			return nil, err
		}
		internalKeys = append(internalKeys, key)
	}
	return internalKeys, nil
}
