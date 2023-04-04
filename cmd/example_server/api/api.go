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

// CreateRoom 创建房间,默认房间id为1，房间成员为1,2
func (h *WebAPI) CreateRoom(ctx *gin.Context) {
	roomInfo := RoomInfo{Room: 1, Member: []uint64{1, 2}}
	room, err := h.m.CreateRoom(roomInfo.Room, 0, roomInfo.Member, 0, "test")
	if nil != err {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	ret := fmt.Sprintf("room.ID=[%d] room.Secret=[%s] room.Time=[%d], room.Member=[%v]",
		room.ID(), room.SecretKey(), room.TimeStamp(), roomInfo.Member)
	ctx.JSON(http.StatusOK, gin.H{"room": ret})
}
