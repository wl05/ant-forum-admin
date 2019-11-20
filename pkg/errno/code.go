package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "网络错误"}
	ErrBind             = &Errno{Code: 10002, Message: "参数错误"}

	ErrValidation = &Errno{Code: 20001, Message: "参数错误"}
	ErrDatabase   = &Errno{Code: 20002, Message: "数据库错误"}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "没有找到该用户"}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "无效token"}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "密码不正确"}

	// tags errors
	ErrTagNotFound = &Errno{Code: 20105, Message: "标签不存在"}

	// categories errors
	ErrCategoryNotFound = &Errno{Code: 20106, Message: "分类不存在"}

	// article errors
	ErrArticleNotFound = &Errno{Code: 20107, Message: "文章不存在"}

	// permission errors
	ErrNotPermission = &Errno{Code: 20108, Message: "没有权限"}

	// casbin errors
	ErrCreateCasbin = &Errno{Code: 20109, Message: "创建策略失败"}

	// menu errors
	ErrMenuGet = &Errno{Code: 20120, Message: "菜单不存在"}

	// role errors
	ErrRoleGet = &Errno{Code: 20121, Message: "角色不存在"}
)
