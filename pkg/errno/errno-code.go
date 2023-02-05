package errno

const SuccessCode = 0

const (
	ServiceErrCode = 10000 + iota

	ParamErrCode
	ParamBindingErrCode

	AuthorizationFailedErrCode

	UserNotExistsErrCode
	UserNameAlreadyExistErrCode

	NotSupportFileTypeErrCode
	NotSupportActionErrCode
)

var (
	Success = New(SuccessCode, "请求成功")

	ServiceErr      = New(ServiceErrCode, "服务内部错误")
	ParamErr        = New(ParamErrCode, "参数错误")
	ParamBindingErr = New(ParamBindingErrCode, "参数绑定错误")

	AuthorizationFailedErr = New(AuthorizationFailedErrCode, "认证失败")

	UserNotExistsErr        = New(UserNotExistsErrCode, "用户不存在")
	UsernameAlreadyExistErr = New(UserNameAlreadyExistErrCode, "用户名已被注册")

	NotSupportFileTypeErr = New(NotSupportFileTypeErrCode, "不支持的文件类型")
	NotSupportActionErr   = New(NotSupportActionErrCode, "不支持的操作")
)
