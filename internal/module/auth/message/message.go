package message

import "oms-be/internal/message"

const (
	MsgErrInvalidPassword      message.Message = "Password invalid"
	MsgErrUserNotFound         message.Message = "User not found"
	MsgErrGenerateAccessToken  message.Message = "Generate access token error"
	MsgErrGenerateRefreshToken message.Message = "Generate refresh token error"
	MsgErrStoreRefreshToken    message.Message = "Store refresh token error"
)
