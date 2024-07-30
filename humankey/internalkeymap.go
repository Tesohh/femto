package humankey

import (
	"fmt"
	"log/slog"
	"reflect"
)

type InternalKeymapEntry struct {
	Sequence []InternalKey
	Command  string
}
type InternalKeymap map[string][]InternalKeymapEntry

func (ik InternalKeymap) GetMatches(mode string, currentSequence []InternalKey) []InternalKeymapEntry {
	entries := []InternalKeymapEntry{}

	for _, v := range ik[mode] {
		slog.Info(fmt.Sprintf("%v %v", len(v.Sequence), len(currentSequence)))
		if len(v.Sequence) < len(currentSequence) {
			continue
		}

		n := len(currentSequence)
		temp := v.Sequence[:n]
		if reflect.DeepEqual(temp, currentSequence) {
			entries = append(entries, v)
		}
	}

	return entries
}
