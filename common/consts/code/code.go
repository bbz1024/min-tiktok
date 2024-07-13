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
	//	-------------------- favorite --------------------
	FavoriteRepeatCode   = 40001
	FavoriteNotFoundCode = 40002

	// -------------------- relation --------------------

	IsFollowCode     = 50001
	IsNotFollowCode  = 50002
	ForbidFollowSelf = 50003
	IsNotFriendCode  = 50004
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

	// -------------------- favorite --------------------

	FavoriteRepeatMsg   = "重复点赞"
	FavoriteNotFoundMsg = "未找到点赞记录"

	// -------------------- relation --------------------

	IsFollowMsg         = "你已关注不可再次关注"
	IsNotFollowMsg      = "你未关注不可取消关注"
	ForbidFollowSelfMsg = "不能关注自己"
	IsNotFriendMsg      = "你与对方不是好友关系"
)
