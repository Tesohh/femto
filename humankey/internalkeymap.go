package humankey

type InternalKeymapEntry struct {
	Sequence []InternalKey
	Command  string
}
type InternalKeymap map[string][]InternalKeymapEntry

func (ik InternalKeymap) GetMatches(mode string, currentSequence []InternalKey) []InternalKeymapEntry {
	entries := []InternalKeymapEntry{}

	for _, v := range ik[mode] {
		// slog.Info(fmt.Sprintf("%v %v", len(v.Sequence), len(currentSequence)))
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
