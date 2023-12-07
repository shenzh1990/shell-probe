package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/robfig/cron/v3"
	"github.com/shenzh1990/shell-probe/middleware"
	"github.com/shenzh1990/shell-probe/util"
	"log"
	"net/http"
	"sync"
)

var (
	m sync.Mutex //互斥锁
)

// 返回一个支持至 秒 级别的 cron
func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

// Conn类型表示WebSocket连接。服务器应用程序从HTTP请求处理程序调用Upgrader.Upgrade方法以获取* Conn：
var (
	upgrader = websocket.Upgrader{
		// 读取存储空间大小
		ReadBufferSize: 10,
		// 写入存储空间大小
		WriteBufferSize: 10,
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return middleware.CheckToken(r.URL.Query().Get("token"))
		},
	}
)

func WsHandler(w http.ResponseWriter, r *http.Request) {
	//   完成握手 升级为 WebSocket长连接，使用conn发送和接收消息。
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer conn.Close()

	//调用连接的WriteMessage和ReadMessage方法以一片字节发送和接收消息。实现如何回显消息：
	//p是一个[]字节，messageType是一个值为websocket.BinaryMessage或websocket.TextMessage的int。
	for {
		c := newWithSeconds()
		spec := "*/3 * * * * ?"
		c.AddFunc(spec, func() {
			hostByteArr, _ := json.Marshal(util.JsonResponse(0, util.SUCCESS_RTN, HostInfo()))
			m.Lock() //加上互斥锁
			if err := conn.WriteMessage(1, hostByteArr); err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNoStatusReceived) {
					log.Printf("error: %v", err)
				}
			}
			m.Unlock()
		})
		c.Start()
		select {}
	}
}

func GetWsHostInfo(c *gin.Context) {
	WsHandler(c.Writer, c.Request)
}
