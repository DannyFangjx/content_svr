package errorcode

import (
	"content_svr/pub/errors"
	"net/http"
)

func GenBusiErr(busiErr *errors.Error, errMsg string) error {
	return &errors.Error{
		ErrMsg:       "",
		ErrDetail:    nil,
		ErrMsgParams: nil,
		UserMsg:      errMsg,
		UserErrCode:  busiErr.UserErrCode,
		//HttpCode:     0,
		Stack: "",
	}
}

var (
	OK                          = &errors.Error{UserErrCode: 1000, UserMsg: "成功[OK]"}
	UNKNOW_ERROR                = &errors.Error{UserErrCode: 100, UserMsg: "未知错误[UNKNOWN ERROR]"}
	INTERNAL_ERROR              = &errors.Error{UserErrCode: 101, UserMsg: "服务器内部错误"}
	INTERNAL_EXCEPTION          = &errors.Error{UserErrCode: 102, UserMsg: "服务器内部异常"}
	INTERNAL_DB_EXCEPTION       = &errors.Error{UserErrCode: 103, UserMsg: "数据库异常"}
	UNSUPPORTED_OPERATION       = &errors.Error{UserErrCode: 104, UserMsg: "不支持该操作"}
	INSUFFICIENT_PERMISSION     = &errors.Error{UserErrCode: 105, UserMsg: "权限不足"}
	FREQUENT_OPERATION          = &errors.Error{UserErrCode: 106, UserMsg: "操作过于频繁，请稍后再试!"}
	TRIES_LIMIT                 = &errors.Error{UserErrCode: 107, UserMsg: "次数超过限制"}
	READ_JSON_ERROR             = &errors.Error{UserErrCode: 200, UserMsg: "请求的json格式错误"}
	VALIDATE_ERROR              = &errors.Error{UserErrCode: 201, UserMsg: "校验请求参数失败"}
	VALIDATE_SIGN_ERROR         = &errors.Error{UserErrCode: 202, UserMsg: "校验签名失败"}
	VALIDATE_TIMESTAMP_ERROR    = &errors.Error{UserErrCode: 203, UserMsg: "校验时间戳失败，请确认系统时间的准确性"}
	DATA_ERROR                  = &errors.Error{UserErrCode: 204, UserMsg: "数据错误"}
	DATA_NOT_EXISTS             = &errors.Error{UserErrCode: 205, UserMsg: "数据不存在"}
	DATA_EXISTS                 = &errors.Error{UserErrCode: 206, UserMsg: "数据已存在"}
	DATA_INVALID                = &errors.Error{UserErrCode: 207, UserMsg: "数据无效"}
	ALL_GROUP_ERROR             = &errors.Error{UserErrCode: 208, UserMsg: "选择的圈子都不存在"}
	SOME_GROUP_ERROR            = &errors.Error{UserErrCode: 209, UserMsg: "选择的部分圈子不存在"}
	WITHOUT_GROUP_ERROR         = &errors.Error{UserErrCode: 210, UserMsg: "没有圈子"}
	THIRD_NET_ERROR             = &errors.Error{UserErrCode: 211, UserMsg: "请求第三方接口超时或网络出错"}
	THIRD_API_ERROR             = &errors.Error{UserErrCode: 212, UserMsg: "请求第三方接口出错"}
	PARAM_ERROR                 = &errors.Error{UserErrCode: 213, UserMsg: "请求参数错误"}
	REPEAT_PLUS_DAY             = &errors.Error{UserErrCode: 214, UserMsg: "重复操作"}
	VERSION_CODE_IS_NULL        = &errors.Error{UserErrCode: 215, UserMsg: "版本号为空"}
	APP_TYPE_IS_NULL            = &errors.Error{UserErrCode: 216, UserMsg: "客户端类型为空"}
	NOT_SUPPORT_APP_TYPE        = &errors.Error{UserErrCode: 217, UserMsg: "不支持的客户端类型"}
	NOT_AUTH                    = &errors.Error{UserErrCode: 218, UserMsg: "无效操作"}
	LOGIN_INVALID               = &errors.Error{UserErrCode: 219, UserMsg: "登录失效"}
	APP_NAME_IS_NULL            = &errors.Error{UserErrCode: 220, UserMsg: "应用名称为空"}
	NOT_SUPPORT_APP_NAME        = &errors.Error{UserErrCode: 221, UserMsg: "不支持的应用"}
	SENSITIVE_WORD              = &errors.Error{UserErrCode: 222, UserMsg: "请求参数包含非法字符"}
	REPEAT_OPERATION            = &errors.Error{UserErrCode: 223, UserMsg: "频繁操作，请稍后再试"}
	USER_ENABLED_ERROR          = &errors.Error{UserErrCode: 224, UserMsg: "账号状态异常"}
	VERIFY_CODE_INVALID         = &errors.Error{UserErrCode: 230, UserMsg: "验证码无效"}
	ACCOUNT_FROZENED            = &errors.Error{UserErrCode: 301, UserMsg: "账户被冻结"}
	ACCOUNT_CLOSED              = &errors.Error{UserErrCode: 302, UserMsg: "账户已销户"}
	ACCOUNT_STATUS_ERROR        = &errors.Error{UserErrCode: 303, UserMsg: "账户状态不正确"}
	ACCOUNT_DEAL_AMOUNT_ERROR   = &errors.Error{UserErrCode: 304, UserMsg: "账户金额变更失败"}
	ACCOUNT_RECORD_SAVE_ERROR   = &errors.Error{UserErrCode: 305, UserMsg: "账户金额变更记录保存失败"}
	SHARE_FORETELL_LIMIT        = &errors.Error{UserErrCode: 306, UserMsg: "今天分享得星币次数已到上限"}
	SHARE_FORETELL_WORK_LIMIT   = &errors.Error{UserErrCode: 307, UserMsg: "该作品已被分享过"}
	FORETELL_COMMENT_LIMIT      = &errors.Error{UserErrCode: 308, UserMsg: "今天分享得星币次数已到上限"}
	FORETELL_COMMENT_WORK_LIMIT = &errors.Error{UserErrCode: 309, UserMsg: "该作品已被分享过"}
	FORETELL_COMMENT_LIKE_LIMIT = &errors.Error{UserErrCode: 310, UserMsg: "重复点赞"}
	USER_WORKS_LOCK_INTERIM     = &errors.Error{UserErrCode: 601, UserMsg: "临时封禁作品发布权限"}
	USER_WORKS_LOCK_PERMANENT   = &errors.Error{UserErrCode: 602, UserMsg: "永久封禁作品发布权限"}
	USER_COMMENT_LOCK_INTERIM   = &errors.Error{UserErrCode: 603, UserMsg: "临时封禁评论发布权限"}
	USER_COMMENT_LOCK_PERMANENT = &errors.Error{UserErrCode: 604, UserMsg: "永久封禁评论发布权限"}
	USER_LOCK_INTERIM           = &errors.Error{UserErrCode: 605, UserMsg: "限时封禁账号"}
	USER_LOCK_PERMANENT         = &errors.Error{UserErrCode: 606, UserMsg: "永久封禁账号"}
	BIND_PHONE_REQUIRE          = &errors.Error{UserErrCode: 607, UserMsg: "未绑定手机号"}
	LockFail                    = &errors.Error{UserErrCode: 8000, UserMsg: "获取锁失败"}
	ENCRYPT_ERROR               = &errors.Error{UserErrCode: 1103, UserMsg: "加密出错"}
	DECRYPT_ERROR               = &errors.Error{UserErrCode: 1104, UserMsg: "解密出错"}
	SIGN_ERROR                  = &errors.Error{UserErrCode: 1105, UserMsg: "计算签名出错"}
	SIGN_CHECK_FAIL             = &errors.Error{UserErrCode: 1106, UserMsg: "验证签名失败"}

	WECHAT_CODE_ERROR = &errors.Error{UserErrCode: 42003, UserMsg: "微信授权错误"}
	QQ_CODE_ERROR     = &errors.Error{UserErrCode: 43001, UserMsg: "QQ授权错误"}

	BUSINESS_ERROR         = &errors.Error{UserErrCode: 300, UserMsg: "业务异常"}
	IllegalParamFailed     = &errors.Error{UserErrCode: 3108, UserMsg: "传输参数不正确"}
	PayOrderNotExists      = &errors.Error{UserErrCode: 3302, UserMsg: "支付单不存在"}
	RefundComplete         = &errors.Error{UserErrCode: 3404, UserMsg: "该笔支付已全额退款，不能再发起退款申请"}
	REFUND_AMOUNT_EXCEED   = &errors.Error{UserErrCode: 3412, UserMsg: "申请的退款金额超过了可退款的金额"}
	RefundOrderFailed      = &errors.Error{UserErrCode: 3110, UserMsg: "退款失败"}
	RefundStatusError      = &errors.Error{UserErrCode: 3406, UserMsg: "退款状态异常，不能更新退款状态"}
	RefundStatusUpdateFail = &errors.Error{UserErrCode: 3407, UserMsg: "退款单状态更新失败，请不要同时更新同一退款单"}

	INTERNAL_INTERFACE_ERROR         = &errors.Error{UserErrCode: 400, UserMsg: "内部接口异常"}
	CALL_ACCOUNTING_SERVICE_ERROR    = &errors.Error{UserErrCode: 401, UserMsg: "调用账务服务异常"}
	REFUND_SELLER_BALANCE_NOT_ENOUGH = &errors.Error{UserErrCode: 3413, UserMsg: "收款账号余额不足，请继续收款后再试"}
	Wx_REFUND_CERT_LOAD_FAIL         = &errors.Error{UserErrCode: 4120, UserMsg: "加载微信证书失败，请检查微信证书文件是否上传正确"}
	WX_REFUND_NO_CERT_FILE           = &errors.Error{UserErrCode: 4121, UserMsg: "缺少微信退款证书文件，请上传后重试"}
	WX_REFUND_SUB_MERCHANT_FORBIDDEN = &errors.Error{UserErrCode: 4122, UserMsg: "微信报没有退款权限"}
	WX_REFUND_CERT_NOT_MATCH         = &errors.Error{UserErrCode: 4123, UserMsg: "微信报退款证书文件错误，请重新上传"}
	WX_REFUND_TOO_FREQUENT           = &errors.Error{UserErrCode: 4124, UserMsg: "微信报退款请求过于频繁，请稍后再试"}
	WX_REFUND_SYSTEM_ERROR           = &errors.Error{UserErrCode: 4125, UserMsg: "微信报系统错误，请稍后再试"}
	WX_REFUND_FAIL                   = &errors.Error{UserErrCode: 4107, UserMsg: "微信退款失败"}
	WxpayResponseFailed              = &errors.Error{UserErrCode: 4102, UserMsg: "微信接口报错!"}
	WxpayResponseNullFailed          = &errors.Error{UserErrCode: 4101, UserMsg: "微信支付请求错误,响应为空!"}
	WxpayCommunicationFailed         = &errors.Error{UserErrCode: 4103, UserMsg: "微信接口调用失败"}
	WxappletAppidNull                = &errors.Error{UserErrCode: 4104, UserMsg: "微信小程序APPID为空"}
	WxAppidNull                      = &errors.Error{UserErrCode: 4105, UserMsg: "微信APPID为空"}
	WxPartneridNull                  = &errors.Error{UserErrCode: 4106, UserMsg: "微信商户号为空"}
	WxpayNotifyFailed                = &errors.Error{UserErrCode: 4131, UserMsg: "微信通知出错!"}
	WxPartnerkeyNull                 = &errors.Error{UserErrCode: 4107, UserMsg: "微信商户密钥为空"}

	QQPAY_REFUND_CERT_LOAD_FAIL         = &errors.Error{UserErrCode: 5120, UserMsg: "加载QQ钱包证书失败，请检查QQ钱包证书文件是否上传正确"}
	QQPAY_REFUND_NO_CERT_FILE           = &errors.Error{UserErrCode: 5121, UserMsg: "缺少QQ钱包退款证书文件，请上传后重试"}
	QQPAY_REFUND_SUB_MERCHANT_FORBIDDEN = &errors.Error{UserErrCode: 5122, UserMsg: "QQ钱包报没有退款权限"}
	QQPAY_REFUND_CERT_NOT_MATCH         = &errors.Error{UserErrCode: 5123, UserMsg: "QQ钱包报退款证书文件错误，请重新上传"}
	QQPAY_REFUND_TOO_FREQUENT           = &errors.Error{UserErrCode: 5124, UserMsg: "QQ钱包报退款请求过于频繁，请稍后再试"}
	QQPAY_REFUND_SYSTEM_ERROR           = &errors.Error{UserErrCode: 5125, UserMsg: "QQ钱包报系统错误，请稍后再试"}
	QQPAY_REFUND_FAIL                   = &errors.Error{UserErrCode: 5107, UserMsg: "QQ钱包退款失败"}
	QQPayResponseFailed                 = &errors.Error{UserErrCode: 5102, UserMsg: "QQ钱包接口报错!"}
	QQPayResponseNullFailed             = &errors.Error{UserErrCode: 5101, UserMsg: "QQ钱包支付请求错误,响应为空!"}
	QQPayCommunicationFailed            = &errors.Error{UserErrCode: 5103, UserMsg: "QQ钱包接口调用失败"}
	QQAppletAppidNull                   = &errors.Error{UserErrCode: 5104, UserMsg: "QQ小程序APPID为空"}
	QQAppidNull                         = &errors.Error{UserErrCode: 5105, UserMsg: "QQ公众平台APPID为空"}
	QQPartneridNull                     = &errors.Error{UserErrCode: 5106, UserMsg: "QQ平台商户号为空"}
	QQPayNotifyFailed                   = &errors.Error{UserErrCode: 5131, UserMsg: "QQ钱包通知出错!"}
	QQPartnerkeyNull                    = &errors.Error{UserErrCode: 5107, UserMsg: "QQ平台商户密钥为空"}
	QQAppkeyNull                        = &errors.Error{UserErrCode: 5108, UserMsg: "QQ平台移动应用密钥为空"}

	REALNAME_CHECK_ERROR = &errors.Error{UserErrCode: 5201, UserMsg: "QQ钱包未实名认证!"}

	ALIPAY_SIGN_CHECK_ERROR        = &errors.Error{UserErrCode: 6001, UserMsg: "支付宝签名验证失败"}
	ALIPAY_NOTIFY_DATA_CHECK_ERROR = &errors.Error{UserErrCode: 6002, UserMsg: "支付宝回调数据验证失败"}

	MIDAS_INTERFACE_ERROR        = &errors.Error{UserErrCode: 7001, UserMsg: "米大师接口异常!"}
	MIDAS_BUSY                   = &errors.Error{UserErrCode: 7002, UserMsg: "系统繁忙，请稍候再试!"}
	MIDAS_QQ_SIGN_ERROR          = &errors.Error{UserErrCode: 7003, UserMsg: "qq_sig签名错误!"}
	MIDAS_SIGN_ERROR             = &errors.Error{UserErrCode: 7003, UserMsg: "sig签名错误!"}
	MIDAS_LOGIN_ERROR            = &errors.Error{UserErrCode: 7004, UserMsg: "用户未登录或登录态已过期!"}
	MIDAS_BILL_NO_EXISTS         = &errors.Error{UserErrCode: 7005, UserMsg: "订单已存在!"}
	MIDAS_BALANCE_NOT_ENOUGH     = &errors.Error{UserErrCode: 7006, UserMsg: "余额不足!"}
	MIDAS_INTERFACE_INSUFFICIENT = &errors.Error{UserErrCode: 7007, UserMsg: "没有调用接口的权限!"}
	MIDAS_PARAMS_ERROR           = &errors.Error{UserErrCode: 7008, UserMsg: "参数错误!"}
	MIDAS_ACCESS_TOKEN_ERROR     = &errors.Error{UserErrCode: 7009, UserMsg: "access_token 校验失败，access_token需要放在url中!"}

	IAP_VERIFY_ERROR   = &errors.Error{UserErrCode: 8001, UserMsg: "验证收据出错"}
	IAP_SANDBOX_2_PROD = &errors.Error{UserErrCode: 8002, UserMsg: "receipt是Sandbox receipt，但却发送至生产系统的验证服务"}
	IAP_PROD_2_SANDBOX = &errors.Error{UserErrCode: 8003, UserMsg: "receipt是生产receipt，但却发送至Sandbox环境的验证服务"}
	IAP_INSUFFICIENT   = &errors.Error{UserErrCode: 8004, UserMsg: "无效收据"}

	ICR_VER  = &errors.Error{UserErrCode: 9010, UserMsg: "请下载最新版本注册"}
	ICR_CODE = &errors.Error{UserErrCode: 9011, UserMsg: "邀请码已失效"}

	group_expired                = &errors.Error{UserErrCode: 10001, UserMsg: "群聊已过期"}
	single_talk_limit            = &errors.Error{UserErrCode: 11001, UserMsg: "每日单聊人数已达上限"}
	VERSION_LOW                  = &errors.Error{UserErrCode: 12001, UserMsg: "版本过低，请升级到新版本"}
	cd_key_invalid               = &errors.Error{UserErrCode: 13001, UserMsg: "无效的激活码"}
	cd_key_used                  = &errors.Error{UserErrCode: 13002, UserMsg: "激活码已被使用"}
	cd_key_register_date_invalid = &errors.Error{UserErrCode: 13003, UserMsg: "注册时间不够"}
	cd_key_black_house_invalid   = &errors.Error{UserErrCode: 13004, UserMsg: "近期有违规记录"}
	cd_key_block_invalid         = &errors.Error{UserErrCode: 13005, UserMsg: "近期有封号记录"}

	Comment_op_fail               = &errors.Error{UserErrCode: 15000, UserMsg: "操作失败"}
	Comment_unlawfully            = &errors.Error{UserErrCode: 15001, UserMsg: "内容非法"}
	Comment_delete_fail           = &errors.Error{UserErrCode: 15002, UserMsg: "删除失败"}
	Comment_nums_used_up_by_user  = &errors.Error{UserErrCode: 15003, UserMsg: "今日评论已达到上限，请明日再来"}
	Comment_nums_used_up_for_work = &errors.Error{UserErrCode: 15004, UserMsg: "这条帖子评论已达到上限"}
	Comment_disable_for_work      = &errors.Error{UserErrCode: 15005, UserMsg: "关闭了帖子评论功能"}
	Comment_work_not_exist        = &errors.Error{UserErrCode: 15006, UserMsg: "删除了动态"}
	Comment_db_op_fail            = &errors.Error{UserErrCode: 15007, UserMsg: "数据库操作失败"}

	ACLFORBID    = &errors.Error{UserErrCode: 21000, UserMsg: "acl_forbid", HttpCode: http.StatusServiceUnavailable}
	PleaseUseApp = &errors.Error{UserErrCode: 21001, UserMsg: "请使用猫爪app体验~"} //小程序男性不允许发帖
)
