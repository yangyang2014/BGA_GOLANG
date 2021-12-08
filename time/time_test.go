package time_test

import (
	"fmt"
	"testing"
	"time"
)

/*
获取当前时间的格式化
*/
//2006-01-02 15:04:05据说是golang的诞生时间，固定写法
func Test_format_time(t *testing.T) {
	//固定写法
	fmt.Println(time.Now().Format("2006-01-02-15_04_05.00000")) //2021-03-04-17_56_45.48101
	fmt.Println(time.Now().Format("06-01-02-15_04_05.00000"))   //21-03-04-17_56_45.48101
}

/*
获取当前年份
*/
func Test_get_time(t *testing.T) {
	year := time.Now().Year()
	fmt.Println(year)
	// month := time.Now().Month()
	// day := time.Now().Day()
	// minute := time.Now().Second()

}

/*
时间作差
*/
func Test_timeSub(t *testing.T) {
	begin := time.Now()
	time.Sleep(time.Second)
	end := time.Now()
	fmt.Println(end.Sub(begin).String())
}

/*
解析日期格式的字符串
*/
func Test_parseTimeStr(t *testing.T) {
	timeStr := "2021-10-12_17:18:19.012"
	//ParseInLoacation 以指定的时区来解析时间
	//2016-01-02 15:04:05 为固定格式，golang团队使用123456便于记忆，如果把小时的15改为03就只支持解析12h时间，直接15可以解析24制的时间
	t1, e := time.ParseInLocation("2006-01-02_15:04:05", timeStr, time.Local)
	if e != nil {
		fmt.Println(e)
		return
	}else{
		fmt.Println(t1)
	}

}

//TODO 可以研究这个文章常用的时间处理
//http://www.zzvips.com/article/68854.html


/*
# 时间的格式化知识：
定义了一些格式，用于将时间进行展示：
| 01/Jan | 02 | 03/15 | 04 | 05 | 06 | -07[00][:00] | PM   | Mon  |
| 月     | 日  |  时    |分 | 秒   |年 | 时差          | 上下午| 星期几|

特别解释
*易错点* 03/15  03表示12小时制 15表示24小时制
 */

/*
# 时间的控制转换层
 */

/*
# 时间的核心层
 */