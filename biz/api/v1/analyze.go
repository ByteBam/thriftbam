package v1

type AnalyzeRequest struct {
	BranchId   string `json:"branchId"`
	UpdateTime string `json:"updateTime"`
	UserId     string `json:"userId"`
}

type AnalyzeResponse struct {
	InterfaceNum int `json:"nums"`
}
