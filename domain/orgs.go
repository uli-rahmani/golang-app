package domain

type OrgsCommentBody struct {
	Comment string `json:"comment" db:"comment"`
}

type GetMemberResult struct {
	OrgsName  string `json:"org_name" db:"org_name"`
	AvatarURL string `json:"avatar_url" db:"avatar_url"`
	Followers int    `json:"followers" db:"followers"`
	Following int    `json:"following" db:"following"`
}
