package users

type Tier struct {
	Type  int8
	Rules Rules
}

var Default_tier = Tier{-1, Default_rules}

func (tier *Tier) MakeAdmin() {
	tier.Type = ADMIN
	tier.Rules = Admin_rules
}

func (tier *Tier) MakeUser() {
	tier.Type = USER
	tier.Rules = User_rules
}

func (tier *Tier) MakeFreeUser() {
	tier.Type = FREE_USER
	tier.Rules = Free_user_rules
}
