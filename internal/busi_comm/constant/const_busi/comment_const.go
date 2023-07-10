package const_busi

// comment definition
const (
	COMMENT_STATUS_UNREAD int32 = 0
	COMMENT_STATUS_READ   int32 = 1
)
const (
	WORK_COMMENT_STATUS_DISABLE int32 = 0
	WORK_COMMENT_STATUS_ENABLE  int32 = 1
)

// mongodb sort des asc
const (
	MONGODB_SORT_DES int32 = -1
	MONGODB_SORT_AES int32 = 1
)
const (
	INVALID_COMMENT_WORK_EXPIRE_TM = 12 * 3600
)
const (
	MAX_COMMENT_NUMS_ON_WORK        = 5
	MAX_COMMENT_NUMS_PER_DAY_PERSON = 9999
)

const (
	MAX_COMMENT_LEN = 30
)
