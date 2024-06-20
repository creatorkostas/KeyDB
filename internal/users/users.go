package users

const (
	ADMIN     = 0
	USER      = 1
	FREE_USER = 2
)

const (
	CAN_GET = 0
	CAN_ADD = 1
)

type UserInfo struct {
	Username string
	Api_key  string
	Email    string
	Password string
}

type Account struct {
	UserInfo
	AccountState
	Tier Tier
}

func (acc Account) IsAdmin() bool {
	return acc.Tier.Rules.Admin
}

func (acc Account) CanAdd() bool {
	return acc.Tier.Rules.Add
}

func (acc Account) CanGet() bool {
	return acc.Tier.Rules.Get
}

func (acc Account) CanChangePassword() bool {
	return acc.Tier.Rules.Change_password
}

func (acc Account) CanChangeApiKey() bool {
	return acc.Tier.Rules.Change_api_key
}

func (acc Account) CanGetAnalytics() bool {
	return acc.Tier.Rules.Analytics
}

// func (acc Account) CheckPermissions(permission int) {
// 	switch permission {
// 	case 0:
// 	}
// }

var Accounts = make([]Account, 50)
