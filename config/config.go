package config

// Info describes $HOME/.op/config
type Info struct {
	LatestSignin string `json:"latest_signin"`
	Accounts     []struct {
		Shorthand  string `json:"shorthand"`
		URL        string `json:"url"`
		Email      string `json:"email"`
		AccountKey string `json:"accountKey"`
		UserUUID   string `json:"userUUID"`
	} `json:"accounts"`
}
