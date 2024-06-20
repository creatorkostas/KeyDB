package users

import "time"

type AccountState struct {
	Active       bool
	Tokens       int64
	Rate_reset   time.Duration
	Burst_time   time.Duration
	Burst_tokens int
}

func (state *AccountState) Activate() {
	state.Active = true
}

func (state *AccountState) Diactivate() {
	state.Active = false
}
