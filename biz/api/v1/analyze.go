package v1

type AnalyzeRequest struct {
	Id          string `json:"id" binding:"required"`
	AccessToken string `json:"access_token"` // userId
	Owner       string `json:"owner" binding:"required"`
	Repo        string `json:"repo" binding:"required"`
	Path        string `json:"path" binding:"required"`
	Ref         string `json:"ref"`
}

type Resp struct {
	InterfaceNum int `json:"nums"`
}
