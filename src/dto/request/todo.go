package request

type TodoReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type TodoReqUpdate struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}
