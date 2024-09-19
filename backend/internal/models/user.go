package models

type UserSearchResult struct {
	Items []User `json:"items"`
}

type User struct {
	Login   string `json:"login"`
	HTMLURL string `json:"html_url"`
}

type UserProfile struct {
	Login          string              `json:"login"`
	Name           string              `json:"name"`
	Location       string              `json:"location"`
	AvatarURL      string              `json:"avatar_url"`
	Blog           string              `json:"blog"`
	GithubURL      string              `json:"html_url"`
	Twitter        string              `json:"twitter_username"`
	Bio            string              `json:"bio"`
	SocialAccounts []map[string]string `json:"social_accounts"`
}
