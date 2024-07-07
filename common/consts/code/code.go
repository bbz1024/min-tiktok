package code

// code
const (
	// -------------------- base --------------------

	OK          = 0
	ParamError  = 400
	ServerError = 500

	// -------------------- 用户 --------------------

	UserNotFoundCode      = 10001
	UserExistedCode       = 10002
	UserPasswordErrorCode = 10003

	// -------------------- 认证 --------------------

	AuthErrorCode = 20001

	// -------------------- video --------------------

	VideoOverSizeCode = 30001
	VideoTypeCode     = 30002
)

// msg
const (
	// -------------------- base --------------------

	OkMsg          = "success"
	ParamErrorMsg  = "参数错误"
	ServerErrorMsg = "服务器错误"
	// -------------------- 用户 --------------------

	UserNotFoundMsg      = "用户不存在"
	UserExistedMsg       = "用户已存在"
	UserPasswordErrorMsg = "密码错误"

	// -------------------- 认证 --------------------

	AuthErrorMsg = "认证失败"

	// -------------------- video --------------------

	VideoOverSizeMsg = "视频大小超过限制"
	VideoTypeMsg     = "视频类型错误"
)
