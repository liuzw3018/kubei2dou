package public

/**
 * @Author: liu zw
 * @Date: 2022/10/20 17:55
 * @File: public.go
 * @Description: //TODO $
 * @Version:
 */

const (
	ValidatorKey  = "ValidatorKey"
	TranslatorKey = "TranslatorKey"
)

type ResponseCode int

// 1000以下为通用码，1000以上为用户自定义码
const (
	SuccessCode ResponseCode = iota
	UndefErrorCode
	ValidErrorCode
	InternalErrorCode

	InvalidRequestErrorCode ResponseCode = 401
	CustomizeCode           ResponseCode = 1000

	GroupAllSaveFlowError ResponseCode = 2001
	OtherErrorCode        ResponseCode = -1
)
