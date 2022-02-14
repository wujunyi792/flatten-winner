package dto

// JsonResponse Json 返回体结构
type JsonResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Count   int         `json:"count"`
	Data    interface{} `json:"data,omitempty"`
}

func (res *JsonResponse) Clear() *JsonResponse {
	res.Message = "success"
	res.Code = 20000
	res.Count = 0
	res.Data = nil
	return res
}

type JsonDataResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Count   int         `json:"total"`
	Data    interface{} `json:"rows,omitempty"`
}

func (res *JsonDataResponse) Clear() *JsonDataResponse {
	res.Message = "success"
	res.Code = 20000
	res.Count = 0
	res.Data = nil
	return res
}

type SetPathConfig struct {
	Path string `json:"path" binding:"required"`
	Key  string `json:"key" binding:"required"`
}
