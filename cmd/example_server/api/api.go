package api

import (
	_ "embed"
	"fmt"
	"github.com/byebyebruce/lockstepserver/logic"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "net/http/pprof"
)

// WebAPI http api
type WebAPI struct {
	m *logic.RoomManager
}

// NewWebAPI 构造
func NewWebAPI(addr string, m *logic.RoomManager) *WebAPI {
	api := &WebAPI{
		m: m,
	}

	r := gin.New()
	r.GET("/create", api.CreateRoom)
	go func() {
		_ = r.Run(addr)
	}() // listen and serve on 0.0.0.0:8080
	return api
}

type RoomInfo struct {
	Room   uint64   `json:"room" form:"room"`
	Member []uint64 `json:"member" form:"member"`
}

func (h *WebAPI) CreateRoom(ctx *gin.Context) {
	var roomInfo RoomInfo
	err := ctx.Bind(&roomInfo)
	fmt.Println("roomInfo=", roomInfo, "err=", err)

	room, err := h.m.CreateRoom(roomInfo.Room, 0, roomInfo.Member, 0, "test")
	if nil != err {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	ret := fmt.Sprintf("room.ID=[%d] room.Secret=[%s] room.Time=[%d], room.Member=[%v]",
		room.ID(), room.SecretKey(), room.TimeStamp(), roomInfo.Member)
	ctx.JSON(http.StatusOK, gin.H{"room": ret})
}
