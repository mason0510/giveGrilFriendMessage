package bootstrap

import (
	"encoding/json"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"giveGrilFriendMessage/handlers"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type dataFetch func() any

// 天气服务
func GetGreeting() string {
	//获取随机数
	rand.Seed(time.Now().Unix())
	r := rand.Intn(4)
	fmt.Println(r)
	//随机数为1的时候发送早安
	//获取问候表情数组
	//获取今天的天气信息
	//天气好时候发的信息
	greeting := []string{"早安~" + openwechat.Emoji.Kiss, "早上好~" + openwechat.Emoji.Heart, "早上好" + openwechat.Emoji.Rose, "早安，今天天气不错。" + openwechat.Emoji.Blowkiss, "早安。" + openwechat.Emoji.InLove}
	//天气不好时候发的信息
	badWeather := []string{"早安~" + openwechat.Emoji.Hug + openwechat.Emoji.Kiss, "早。" + openwechat.Emoji.InLove, "早上好" + openwechat.Emoji.Rose, "早安" + openwechat.Emoji.Blowkiss}
	fetch := dataFetch(getData)
	a := fetch()
	//获取到的天气信息
	//fmt.Println(a)
	//判断天气信息
	//如果天气大于10度 greeting
	//如果天气小于10度 badWeather
	if a.(float64) > 10 {
		//发送信息
		return greeting[r]
	} else {
		//发送信息
		return badWeather[r]
	}

}

func getData() any {
	type weatherData struct {
		Main struct {
			Temp     float64 `json:"temp"`
			Humidity int     `json:"humidity"`
		} `json:"main"`
	}
	apiKey := "8c1a38623e2a01368b8e78508d96aef2"
	//打印apiKey
	fmt.Println(apiKey)
	city := "Shanghai, CN"
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	//打印data
	//fmt.Println(string(data))

	var weather weatherData
	if err := json.Unmarshal(data, &weather); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Temperature in %s: %.2f°C\n", city, weather.Main.Temp-273.15)
	fmt.Printf("Humidity in %s: %d%%\n", city, weather.Main.Humidity)
	return weather.Main.Temp - 273.15
}
func Run() {
	wb, err := NewWeChatBot("zhxcmason")
	if err != nil {
		fmt.Println(err)
		return
	}
	go wb.SendMessageToGirlFriend()
	wb.bot.Block()
}

func (wb *WeChatBot) SendMessageToGirlFriend() {
	for {
		now := time.Now()
		rand.Seed(time.Now().Unix())
		//r := rand.Intn(30)+16
		//t := time.Date(now.Year(), now.Month(), now.Day(), 9, r, 0, 0, now.Location())
		//每五分钟发一次
		t := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute()+2, 0, 0, now.Location())
		timer := time.NewTimer(now.Sub(t))
		<-timer.C
		wb.gf.SendText(GetGreeting())
		break
	}
}

type WeChatBot struct {
	bot *openwechat.Bot
	gf  *openwechat.Friend
}

func NewWeChatBot(gfName string) (*WeChatBot, error) {
	bot := openwechat.DefaultBot(openwechat.Desktop) // Desktop mode, try this mode if you can't login

	// Register message handler
	bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() && msg.Content == "ping" {
			msg.ReplyText("pong")
		}
	}

	// Register login QR code callback
	bot.UUIDCallback = handlers.QrCodeCallBack

	// Login
	if err := bot.Login(); err != nil {
		return nil, err
	}

	// Get logged-in user
	self, err := bot.GetCurrentUser()
	if err != nil {
		return nil, err
	}

	// Get all friends
	friends, err := self.Friends()
	if err != nil {
		return nil, err
	}
	//打印所有的好友
	for _, friend := range friends {
		fmt.Println(friend.RemarkName)
	}

	// Search for girlfriend by remark name
	gf := friends.SearchByUserName(1, gfName)

	if gf.Count() == 0 {
		return nil, fmt.Errorf("girlfriend not found")
	}

	return &WeChatBot{bot, gf.First()}, nil
}

//import (
//	"fmt"
//	"github.com/eatmoreapple/openwechat"
//	"time"
//)

//import (
//	"fmt"
//	"github.com/eatmoreapple/openwechat"
//	"giveGrilFriendMessage/handlers"
//	"time"
//)
//
//// 定义一个发消息的接口
//type MessageSender interface {
//	SendMessage(msg string)
//}
//
//// 定义一个定时任务的结构体 结构体里包含发送方的信息
//type MessageScheduler struct {
//	Sender  MessageSender
//	Enabled bool
//}
//
//// 具体化定时任务的方法
//func (s *MessageScheduler) Schedule(text string, hour, minute int) {
//	if !s.Enabled {
//		return
//	}
//	for {
//		now := time.Now()
//		t := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())
//		timer := time.NewTimer(t.Sub(now))
//		//阻塞线程
//		<-timer.C
//		s.Sender.SendMessage(text)
//		break
//	}
//}
//
//// Bot
//type Bot struct {
//	Client         *openwechat.Bot
//	UUIDCallback   func(string)
//	MessageHandler func(*openwechat.Message)
//}
//
//func (b *Bot) Login() error {
//	//登陆
//	b.Client = openwechat.DefaultBot(openwechat.Desktop)
//	//注册二维码
//	b.Client.UUIDCallback = b.UUIDCallback
//	//注册消息处理函数
//	b.Client.MessageHandler = b.MessageHandler
//	return b.Client.Login()
//}
//
//// 封装获取当前用户的方法
//func (b *Bot) GetCurrentUser() (*openwechat.Self, error) {
//	user, err := b.Client.GetCurrentUser()
//	if err != nil {
//		return nil, err
//	}
//	return user, nil
//}
//
//// 使用依赖注入的方式，将依赖注入到bootstrap中
//func Run() {
//	//初始化定时任务
//	//var scheduler MessageScheduler
//	//初始化机器人
//	var bot Bot
//	bot.Client = openwechat.DefaultBot()
//	//bot.Client.MessageHandler = func(msg *openwechat.Message) {
//	//	if msg.IsText() && msg.Content != "ping" {
//	//		msg.ReplyText(openwechat.Emoji.Smirk + "你好，小可爱，我是葛优。")
//	//	}
//	//}
//	bot.Client.MessageHandler = handlers.Handler
//	bot.Client.UUIDCallback = handlers.QrCodeCallBack
//	//发送消息
//	//if SendWifeMessage(bot, scheduler) {
//	//	return
//	//}
//	bot.Client.Block()
//}
//
//func SendWifeMessage(bot Bot, scheduler MessageScheduler) bool {
//	// Get current user
//	user, err := bot.GetCurrentUser()
//	if err != nil {
//		fmt.Println(err)
//		return true
//	}
//
//	// Get all friends
//	friends, err := user.Friends()
//	if err != nil {
//		fmt.Println(err)
//		return true
//	}
//	girlFriend := friends.SearchByRemarkName(1, "似水流年")
//	if girlFriend.Count() > 0 {
//		//调用定时任务发送
//		scheduler.Schedule("早上好，爱你哟", 7, 0)
//	}
//	return false
//}

//两种方式对比
//struct
//初始化后的参数可以直接使用
//--new --func
