package const_level

import (
	"fmt"
	"strings"
)

var (
	UserLevel0 = 0
	UserLevel1 = 1
	UserLevel2 = 2
	UserLevel3 = 3
	UserLevel4 = 4
)

type ULevelInfo struct {
	Level                 int32
	Title                 string
	NoVipVisitCountLimit  int32    // 普通用户-每个帖子的曝光上限度
	VipVisitCountLimit    int32    // vip-每个帖子的曝光上限度
	DailyReplyLimit       int32    // 每天可以回复的用户帖子数限制
	FullSignTimes         int32    //当前级别需要签到数量
	FullChitchatCodeTimes int32    //当前级别满级激活码数量
	LevelUpDesc           string   // 升级后的文案。显示在等级卡片上，展示下一级权益。
	LevelUpDescArr        []string // 升级后弹窗文案.
}

var UserLeveTitleDict = map[int32]*ULevelInfo{
	0: {0, "小猫崽", 0, 0, 2, 0, 0, "", nil},
	1: {1, "小奶猫", 49, 49, 10, 7, 1, "", nil},
	2: {2, "初级猫", 99, 199, 20, 3, 0, "", nil},
	3: {3, "中级猫", 199, 399, 30, 7, 1, "", nil},
	4: {4, "高级猫", 699, 999, 40, 30, 6, "", nil},
}

func GetLevelCfg(level int32) *ULevelInfo {
	item := UserLeveTitleDict[level]
	if item == nil {
		return nil
	}

	//
	arrs := make([]string, 0)
	// 可发送自定义表情
	if level >= 3 { // 3以及以上可以发布自定义表情
		arrs = append(arrs, "可发送自定义表情")
	}

	// 可发布帖子，帖子曝光上限至%v /   可发布帖子，帖子曝光上限至299（VIP用户上限达399）
	if level == 1 {
		arrs = append(arrs, fmt.Sprintf("可发布动态，动态曝光上限至%v", item.NoVipVisitCountLimit))
	} else if level >= 2 {
		arrs = append(arrs, fmt.Sprintf("可发布动态，动态曝光上限至%v（VIP用户上限达%v)", item.NoVipVisitCountLimit, item.VipVisitCountLimit))
	}

	// 激活码生成系统开启
	if level >= 3 {
		arrs = append(arrs, "激活码生成系统开启")
	}
	// 每日可主动回复8位猫友
	arrs = append(arrs, fmt.Sprintf("每日可主动回复%v位猫友", item.DailyReplyLimit))
	item.LevelUpDescArr = arrs

	if len(arrs) > 0 {
		item.LevelUpDesc = fmt.Sprintf("%v将拥有权益:%v。", item.Title, strings.Join(arrs, ","))
	}

	newItem := *item
	return &newItem
}
