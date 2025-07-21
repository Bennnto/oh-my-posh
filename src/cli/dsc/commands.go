package dsc

import (
	"encoding/json"

	"github.com/jandedobbeleer/oh-my-posh/src/cache"
	"github.com/jandedobbeleer/oh-my-posh/src/log"
)

const (
	cacheKey = "DSC_STATE"
)

func Get(c cache.Cache) *State {
	state := new(State)
	state.Configurations = make(Configurations, 0)
	state.Shells = make(Shells, 0)
	state.Fonts = make(Fonts, 0)

	cached, ok := c.Get(cacheKey)
	if !ok {
		return state
	}

	if err := json.Unmarshal([]byte(cached), state); err == nil {
		return state
	}

	return state
}

func Update(c cache.Cache, schema *State) {
	data, err := json.Marshal(schema)
	if err != nil {
		log.Error(err)
		return
	}

	c.Set(cacheKey, string(data), cache.INFINITE)
}

func Set(c cache.Cache, schema string) error {
	s := new(State)

	if err := json.Unmarshal([]byte(schema), s); err != nil {
		return err
	}

	if err := s.Apply(); err != nil {
		return err
	}

	Update(c, s)

	return nil
}
