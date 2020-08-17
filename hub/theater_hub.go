package hub

import (
	"errors"
	"github.com/CastyLab/grpc.proto/proto"
	"github.com/getsentry/sentry-go"
	"github.com/gobwas/ws"
	"github.com/gorilla/websocket"
	cmap "github.com/orcaman/concurrent-map"
	"log"
	"net/http"
)

// TheaterHub holds theater rooms
type TheaterHub struct {
	upgrader     websocket.Upgrader
	userHub      *UserHub
	VideoPlayers cmap.ConcurrentMap
	cmap.ConcurrentMap
}

func (hub *TheaterHub) FindRoom(name string) (room Room, err error) {
	if r, ok := hub.Get(name); ok {
		return r.(*TheaterRoom), nil
	}
	return nil, errors.New("theater room is missing from cmp")
}

/* If room doesn't exist creates it then returns it */
func (hub *TheaterHub) GetOrCreateRoom(theater *proto.Theater) (room Room) {
	if r, ok := hub.Get(theater.Id); ok {
		return r.(*TheaterRoom)
	}
	room, _ = NewTheaterRoom(theater, hub)
	hub.Set(theater.Id, room)
	return
}

/* If room doesn't exist creates it then returns it */
func (hub *TheaterHub) GetRoom(name string) (*TheaterRoom, error) {
	if !hub.Has(name) {
		return nil, errors.New("room not found")
	}
	if r, ok := hub.Get(name); ok {
		return r.(*TheaterRoom), nil
	}
	return nil, errors.New("room is missing from cmp")
}

func (hub *TheaterHub) RemoveRoom(name string) {
	if hub.Has(name) {
		hub.Remove(name)
	}
	return
}

func (hub *TheaterHub) Close() error {
	return nil
}

/* Get ws conn. and hands it over to correct room */
func (hub *TheaterHub) Handler(w http.ResponseWriter, req *http.Request) {

	// Upgrade connection to websocket
	conn, _, _, err := ws.UpgradeHTTP(req, w)
	if err != nil {
		sentry.CaptureException(err)
		return
	}

	// Create a new client for user
	client := NewTheaterClient(req.Context(), conn)

	log.Printf("[%s] New client connected", client.Id)

	// Close connection after client disconnected
	defer client.Close()

	// Join user room if client recieved authorized
	client.OnAuthorized(func(auth Auth) (room Room) {

		var (
			err     error
			theater = new(proto.Theater)
			event   = auth.Event().(*proto.TheaterLogOnEvent)
		)

		// getting theater from grpc service
		theater, err = GetTheater(event.Room, auth.token)

		if err != nil {
			client.ctxCancel()
			_ = client.Close()
			return
		}

		room = hub.GetOrCreateRoom(theater)
		room.Join(client)
		return
	})

	// Listen on client events
	client.Listen()
	return
}

/* Constructor */
func NewTheaterHub(uhub *UserHub) *TheaterHub {
	return &TheaterHub{
		upgrader:       newUpgrader(),
		userHub:        uhub,
		VideoPlayers:   cmap.New(),
		ConcurrentMap:  cmap.New(),
	}
}
