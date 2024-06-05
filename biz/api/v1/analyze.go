package v1

type AnalyzeRequest struct {
	BranchId string `json:"branch_id" binding:"required,len=19"`
	UserId   string `json:"user_id" binding:"required,len=19"`
	Owner    string `json:"owner" binding:"required"`
	Repo     string `json:"repo" binding:"required"`
	Path     string `json:"path" binding:"required"`
}

type AnalyzeResponse struct {
	InterfaceNum int `json:"nums"`
}
