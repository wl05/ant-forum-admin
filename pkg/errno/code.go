package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation = &Errno{Code: 20001, Message: "参数错误"}
	ErrDatabase   = &Errno{Code: 20002, Message: "数据库错误"}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}

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
)
