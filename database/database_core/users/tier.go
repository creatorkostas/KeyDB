package users

type Tier struct {
	Type  int8
	Rules Rules
}

var Default_tier = Tier{-1, Default_rules}
var Admin_tier = Tier{0, Admin_rules}
var User_tier = Tier{1, User_rules}
var Guest_user_tier = Tier{2, Guest_user_rules}

// func (tier *Tier) MakeAdmin() {
// 	tier.Type = ADMIN
// 	tier.Rules = Admin_rules
// }

// func (tier *Tier) MakeUser() {
// 	tier.Type = USER
// 	tier.Rules = User_rules
// }

// func (tier *Tier) MakeFreeUser() {
// 	tier.Type = FREE_USER
// 	tier.Rules = Free_user_rules
// }
