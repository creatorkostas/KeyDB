package users

import "time"

type AccountState struct {
	Active       bool
	Tokens       int64
	Rate_reset   time.Duration
	Burst_time   time.Duration
	Burst_tokens int
}

var Admin_state AccountState = AccountState{Active: true, Tokens: -1, Rate_reset: 1 * time.Microsecond, Burst_time: 1 * time.Microsecond, Burst_tokens: 1000}
var User_state AccountState = AccountState{Active: true, Tokens: 100000, Rate_reset: 6 * time.Hour, Burst_time: 100 * time.Millisecond, Burst_tokens: 10}
var Free_user_state AccountState = AccountState{Active: true, Tokens: 1000, Rate_reset: 24 * time.Hour, Burst_time: 500 * time.Millisecond, Burst_tokens: 10}
var Default_state AccountState = AccountState{Active: false, Tokens: 0, Rate_reset: 24 * time.Hour * 30, Burst_time: 24 * time.Hour * 30, Burst_tokens: 1}

func (state *AccountState) Activate() {
	state.Active = true
}

func (state *AccountState) Diactivate() {
	state.Active = false
}
