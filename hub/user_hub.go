package hub

import (
	"context"
	"errors"
	"github.com/CastyLab/gateway.server/grpc"
	"github.com/CastyLab/gateway.server/hub/protocol/protobuf"
	"github.com/CastyLab/gateway.server/hub/protocol/protobuf/enums"
	"github.com/CastyLab/grpc.proto/proto"
	"github.com/getsentry/sentry-go"
	"github.com/gobwas/ws"
	pb "github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/orcaman/concurrent-map"
	"log"
	"net/http"
	"time"
)

/* Controls a bunch of rooms */
type UserHub struct {
	upgrader websocket.Upgrader
	cmap     cmap.ConcurrentMap
	ctx      context.Context
}

/* If room doesn't exist creates it then returns it */
func (h *UserHub) FindRoom(name string) (room *UserRoom, err error) {
	if r, ok := h.cmap.Get(name); ok {
		return r.(*UserRoom), nil
	}
	return nil, errors.New("user room is missing from cmp")
}

/* If room doesn't exist creates it then returns it */
func (h *UserHub) GetOrCreateRoom(name string) (room *UserRoom) {
	if r, ok := h.cmap.Get(name); ok {
		return r.(*UserRoom)
	}
	return NewUserRoom(name, h)
}

func (h *UserHub) RemoveRoom(name string) {
	h.cmap.Remove(name)
	return
}

func (h *UserHub) RollbackUsersStatesToOffline()  {
	usersIds := make([]string, 0)
	for uId := range h.cmap.Items() {
		usersIds = append(usersIds, uId)
	}
	if len(usersIds) > 0 {
		mCtx, _ := context.WithTimeout(h.ctx, 5 * time.Second)
		response, err := grpc.UserServiceClient.RollbackStates(mCtx, &proto.RollbackStatesRequest{
			UsersIds: usersIds,
		})
		if err != nil {
			sentry.CaptureException(err)
			log.Fatal(err)
		}
		if response.Code == http.StatusOK {
			log.Fatal("Rolled back online users state to Offline successfully!")
		}
	}
}

func (h *UserHub) Close() {
	h.RollbackUsersStatesToOffline()
}

/* Get ws conn. and hands it over to correct room */
func (h *UserHub) Handler(w http.ResponseWriter, req *http.Request) {

	h.ctx = req.Context()
	conn, _, _, err := ws.UpgradeHTTP(req, w)
	if err != nil {
		sentry.CaptureException(err)
		log.Println("upgrade:", err)
		return
	}

	client := NewClient(h.ctx, conn, UserRoomType)
	defer client.Close()

	client.OnAuthorized(func(e pb.Message, u *proto.User) Room {
		event := e.(*protobuf.LogOnEvent)
		room := h.GetOrCreateRoom(u.Id)
		room.AuthToken = string(event.Token)
		room.Join(client)
		return room
	})

	client.OnUnauthorized(func() {
		buffer, err := protobuf.NewMsgProtobuf(enums.EMSG_UNAUTHORIZED, nil)
		if err == nil {
			_ = client.WriteMessage(buffer.Bytes())
		}
	})

	client.OnLeave(func(room Room) {
		if room != nil {
			room.Leave(client)
		}
	})

	client.Listen()
	return
}

/* Constructor */
func NewUserHub() *UserHub {
	return &UserHub{
		cmap: cmap.New(),
		upgrader: newUpgrader(),
	}
}
