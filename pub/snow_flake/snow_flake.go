package snow_flake

import (
	"time"
)

/*
参考 https://www.cnblogs.com/efish/p/snow-arithmetic.html
*/

/*
雪花算法组成部分:
共57bit
0(1位，且始终为0)|时间戳(41位)|工作机器id(5位)|序列号(10位)
//共64bit
//0(1位，且始终为0)|时间戳(41位)|工作机器id(10位)|序列号(12位)
*/
var (
	Epoch           int64 = 1552893276000 //2020年8月11号0:00 时刻的毫秒级时间戳
	machineID       int64                 //机器id
	machineIdBitLen int   = 5             //machineID 占位长度
	sn              int64                 //序列号
	snBitLen        int   = 10            //序列号站位长度
	lastTimeStamp   int64                 //记录上次的时间戳(毫秒级)
)

func init() {
	lastTimeStamp = time.Now().UnixNano()/1e6 - Epoch
}

func SetMachineID(mid int64) {
	machineID = mid << snBitLen
}
func GetSnowflakeID() int64 {
	// 单位为毫秒
	curTimeStamp := time.Now().UnixNano()/1e6 - Epoch
	if curTimeStamp == lastTimeStamp {
		sn++
		//序列号为10位， 2^10 = 1024个
		if sn > 1023 {
			//序列号超出，则重置序列号。这也意味着每毫秒最多能生成1024个id值
			time.Sleep(time.Millisecond)
			curTimeStamp = time.Now().UnixNano()/1e6 - Epoch
			lastTimeStamp = curTimeStamp //顺便更新下上次的时间戳
			sn = 0
		}
		//与运算 对应位全为1时，则为1.否则为0
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		rightBinValue <<= (machineIdBitLen + snBitLen)

		//或运算 对应位全为0时，则为0。否则为1
		id := rightBinValue | machineID | sn
		return id
	} else if curTimeStamp > lastTimeStamp {
		sn = 0
		lastTimeStamp = curTimeStamp
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		rightBinValue <<= (machineIdBitLen + snBitLen)
		return rightBinValue | machineID | sn
	}
	return 0

}

//
//func content_svr() {
//	//timesStr := ""
//	//timesInt, err := strconv.Atoi(timesStr)
//	//print(timesInt, err)
//
//	SetMachineID(3)
//	{
//		id := GetSnowflakeID()
//		fmt.Printf("id : %d\n", id)
//		fmt.Printf("时间戳: %d\n", id>>(machineIdBitLen+snBitLen)+Epoch)              //时间戳
//		fmt.Printf("机器码: %d\n", id&(-1^(-1<<machineIdBitLen)<<snBitLen)>>snBitLen) //机器码
//		fmt.Printf("序列号: %d\n", id&(-1^(-1<<snBitLen)))
//		fmt.Println("机器码:", id&(2<<(machineIdBitLen+snBitLen-1)-1)>>snBitLen)
//		fmt.Println("序列号:", id&(2<<snBitLen-1))
//	}
//
//	{
//		id := GetSnowflakeID()
//		fmt.Printf("id : %d\n", id)
//		fmt.Printf("时间戳: %d\n", id>>(machineIdBitLen+snBitLen)+Epoch)              //时间戳
//		fmt.Printf("机器码: %d\n", id&(-1^(-1<<machineIdBitLen)<<snBitLen)>>snBitLen) //机器码
//		fmt.Printf("序列号: %d\n", id&(-1^(-1<<snBitLen)))
//		fmt.Println("机器码:", id&(2<<(machineIdBitLen+snBitLen-1)-1)>>snBitLen)
//		fmt.Println("序列号:", id&(2<<snBitLen-1))
//	}
//
//	{
//		id := GetSnowflakeID()
//		fmt.Printf("id : %d\n", id)
//		fmt.Printf("时间戳: %d\n", id>>(machineIdBitLen+snBitLen)+Epoch)              //时间戳
//		fmt.Printf("机器码: %d\n", id&(-1^(-1<<machineIdBitLen)<<snBitLen)>>snBitLen) //机器码
//		fmt.Printf("序列号: %d\n", id&(-1^(-1<<snBitLen)))
//		fmt.Println("机器码:", id&(2<<(machineIdBitLen+snBitLen-1)-1)>>snBitLen)
//		fmt.Println("序列号:", id&(2<<snBitLen-1))
//	}
//	{
//		id := GetSnowflakeID()
//		fmt.Printf("id : %d\n", id)
//		fmt.Printf("时间戳: %d\n", id>>(machineIdBitLen+snBitLen)+Epoch)              //时间戳
//		fmt.Printf("机器码: %d\n", id&(-1^(-1<<machineIdBitLen)<<snBitLen)>>snBitLen) //机器码
//		fmt.Printf("序列号: %d\n", id&(-1^(-1<<snBitLen)))
//		fmt.Println("机器码:", id&(2<<(machineIdBitLen+snBitLen-1)-1)>>snBitLen)
//		fmt.Println("序列号:", id&(2<<snBitLen-1))
//	}
//
//}
