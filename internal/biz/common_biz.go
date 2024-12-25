package biz

// PaginationRequest 分页请求
type PaginationRequest struct {
	Page     int `json:"page" form:"page" default:"1" binding:"required"`          // 页码
	PageSize int `json:"pageSize" form:"pageSize" default:"10" binding:"required"` // 每页数量
}

// PaginationResponse 分页响应
type PaginationResponse struct {
	Total int64 `json:"total"` // 总数
	List  any   `json:"list"`  // 列表
}
