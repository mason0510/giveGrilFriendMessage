package bootstrap

import (
	"fmt"
	"testing"
	"time"
)

// test
// 测试时间
// test
func TestRun(t *testing.T) {
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), 6, 20, 0, 0, now.Location())
	//打印时间
	fmt.Println(date)
	//当前的时区
	fmt.Println(now.Location())
	//当前的时间
}
