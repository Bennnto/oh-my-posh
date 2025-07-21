package dsc

type Shells map[string]*Shell

type Shell struct {
	Command string `json:"command,omitempty"`
}

func (s *Shells) Add(name, command string) {
	(*s)[name] = &Shell{
		Command: command,
	}
}
