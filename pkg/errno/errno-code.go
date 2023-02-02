package errno

const SuccessCode = 0

const (
	ServiceErrCode = 10000 + iota

	ParamErrCode
	ParamBindingErrCode

	AuthorizationFailedErrCode

	UserNameAlreadyExistErrCode
)

var (
	Success = New(SuccessCode, "请求成功")

	ServiceErr      = New(ServiceErrCode, "服务内部错误")
	ParamErr        = New(ParamErrCode, "参数错误")
	ParamBindingErr = New(ParamBindingErrCode, "参数绑定错误")

	AuthorizationFailedErr = New(AuthorizationFailedErrCode, "认证失败")

	UsernameAlreadyExistErr = New(UserNameAlreadyExistErrCode, "用户名已被注册")
)
