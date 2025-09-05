package request

type LoginUserReq struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}
