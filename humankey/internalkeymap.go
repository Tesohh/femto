package humankey

import (
	"slices"
)

type InternalKeymapEntry struct {
	Sequence []InternalKey
	Command  string
}
type InternalKeymap map[string][]InternalKeymapEntry

func (ik InternalKeymap) GetMatches(mode string, currentSequence []InternalKey) []InternalKeymapEntry {
	entries := []InternalKeymapEntry{}

	for _, v := range ik[mode] {
		if len(v.Sequence) < len(currentSequence) {
			continue
		}

		n := len(currentSequence)
		cutSeq := v.Sequence[:n]
		matches := true
		for i, currentKey := range currentSequence {
			if !cutSeq[i].MatchesInternal(currentKey) {
				matches = false
			}
		}

		if matches {
			entries = append(entries, v)
		}
	}

	return entries
}

// Merges target with other keymaps. other keymaps have priority over the target, and order matters.
func (ik InternalKeymap) MergeInternal(others ...InternalKeymap) InternalKeymap {

	ck := make(InternalKeymap)
	for k, v := range ik {
		for _, km := range v {
			ck[k] = append(ck[k], km)
		}
	}

	for _, other := range others {
		if other == nil {
			continue
		}
		for mode, keymap := range other {
			for _, entry := range keymap {
				index := slices.IndexFunc(ck[mode], func(e InternalKeymapEntry) bool {
					return slices.Equal(e.Sequence, entry.Sequence)
				})
				if index != -1 {
					ck[mode][index] = entry
				} else {
					ck[mode] = append(ck[mode], entry)
				}
			}
		}
	}

	return ck
}
