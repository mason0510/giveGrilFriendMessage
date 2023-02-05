package handlers

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/skip2/go-qrcode"
	"giveGrilFriendMessage/config"
	"giveGrilFriendMessage/service"
	"log"
)

// MessageHandlerInterface 消息处理接口
type MessageHandlerInterface interface {
	handle(*openwechat.Message) error
	ReplyText(*openwechat.Message) error
}

type HandlerType string

const (
	GroupHandler = "group"
	UserHandler  = "user"
)

// handlers 所有消息类型类型的处理器
var handlers map[HandlerType]MessageHandlerInterface
var UserService service.UserServiceInterface

func init() {
	handlers = make(map[HandlerType]MessageHandlerInterface)
	UserService = service.NewUserService()
	//handlers[GroupHandler] = NewGroupMessageHandler()
	//handlers[UserHandler] = NewUserMessageHandler()
}

// QrCodeCallBack 登录扫码回调，
func QrCodeCallBack(uuid string) {
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
	fmt.Println(q.ToString(true))
}

// Handler 全局处理入口
func Handler(msg *openwechat.Message) {
	msgContent := msg.Content
	log.Printf("hadler Received msg : %v", msgContent)
	handlers[UserHandler].handle(msg)
	if msg.IsFriendAdd() {
		if config.LoadConfig().AutoPass {
			msg.Agree("你好,我是Alex，你可以向我提问任何问题。")
		}
	}
}
