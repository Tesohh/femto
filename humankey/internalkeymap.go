package humankey

import "slices"

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
	for _, other := range others {
		if other == nil {
			continue
		}
		for mode, keymap := range other {
			for _, entry := range keymap {
				index := slices.IndexFunc(ik[mode], func(e InternalKeymapEntry) bool {
					return slices.Equal(e.Sequence, entry.Sequence)
				})
				if index != -1 {
					ik[mode][index] = entry
				} else {
					ik[mode] = append(ik[mode], entry)
				}
			}
		}
	}

	return ik
}
