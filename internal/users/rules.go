package users

type Rules struct {
	Admin           bool
	Add             bool
	Get             bool
	Change_password bool
	Change_api_key  bool
	Analytics       bool
}

var Admin_rules Rules = Rules{Admin: true, Add: true, Get: true, Change_password: true, Change_api_key: true, Analytics: true}
var User_rules Rules = Rules{Admin: false, Add: true, Get: true, Change_password: true, Change_api_key: true, Analytics: false}
var Free_user_rules Rules = Rules{Admin: false, Add: true, Get: true, Change_password: true, Change_api_key: false, Analytics: false}
var Default_rules Rules = Rules{Admin: false, Add: false, Get: false, Change_password: false, Change_api_key: false, Analytics: false}

func (rule *Rules) IsAdmin() bool {
	return rule.Admin
}
