package humankey

type HumanKeymap map[string]map[string]string

func (h HumanKeymap) ToInternal() (InternalKeymap, error) {
	imap := InternalKeymap{}
	for mode, modemap := range h {
		for hkey, cmd := range modemap {
			ikey, err := ParseSequence(hkey)
			if err != nil {
				return nil, err
			}
			imap[mode] = append(imap[mode], InternalKeymapEntry{
				Sequence: ikey,
				Command:  cmd,
			})
		}
	}
	return imap, nil
}
