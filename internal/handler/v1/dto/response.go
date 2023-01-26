package dto

type ResponseDto struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func NewResponseDto(msg string, data interface{}, key string) ResponseDto {
	if data != nil {
		return ResponseDto{
			Message: msg,
			Data:    map[string]interface{}{key: data},
		}
	}

	return ResponseDto{Message: msg, Data: nil}
}
