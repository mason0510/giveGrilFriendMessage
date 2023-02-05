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
	//now := time.Now()
	//date := time.Date(now.Year(), now.Month(), now.Day(), 6, 20, 0, 0, now.Location())
	////打印时间
	//fmt.Println(date)
	////当前的时区
	//fmt.Println(now.Location())
	//当前的时间
	//dir, _ := os.Getwd()
	//截掉dir中的config
	//dir = dir[:len(dir)-6]
	//fmt.Println(dir)

	//获取天气
	//var dataFetcher DataFetcher
	////他会自动创建依赖关系绑定到dataFetcher上 那么我们就可以调用dataFetcher的方法了
	//wire.Build(WireSet)
	//dataFetcher.Fetch()
	//GetGreeting()
	//rand.Seed(time.Now().Unix())
	//for i := 0; i < 10; i++ {
	//	greeting := GetGreeting()
	//	fmt.Println(greeting)
	//}
	//查看当前的时区
	fmt.Println(time.Now())
	t.Log(time.Now())
	t.Log(time.Now().Location())
	fmt.Println(time.Now().Location())
	//t.Log(time.Now().Location())
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), 9, 1, 0, 0, now.Location())
	t.Log(date)
}
