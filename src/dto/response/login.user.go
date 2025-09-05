package response

type LoginUserRes struct {
	AccessToken string `json:"accessToken"`
	Type        string `json:"type"`
	Exp         int    `json:"exp"`
}
