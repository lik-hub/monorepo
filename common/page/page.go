package page

// Request 分页请求
type Request struct {
	Page  int64  `form:"page,default=1"`
	Size  int64  `form:"size,default=20"`
	Keyword string `form:"keyword,optional"`
}

func (r *Request) Offset() int64 {
	if r.Page < 1 {
		r.Page = 1
	}
	if r.Size < 1 || r.Size > 100 {
		r.Size = 20
	}
	return (r.Page - 1) * r.Size
}

// Response 分页响应
type Response struct {
	Page  int64       `json:"page"`
	Size  int64       `json:"size"`
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

func NewResponse(page, size, total int64, list interface{}) *Response {
	return &Response{
		Page:  page,
		Size:  size,
		Total: total,
		List:  list,
	}
}
