package main
import (
	"github.com/gorilla/websocket"
)

type message struct {
	count int `json : "count"`
	SiteName string `json : "SiteName"`
}
type hub struct {
	conn *websocket.conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:func (r *http.Request) bool {
		return true

	}
}

func newHub (conn *websocket.Conn) *hub{
return &hub{
	conn : conn,
}
}
func (h *hub) run {
	rand.Seed(time.Now().UnixNano())
for {
	_, readMessage, err := h.receive()
	if err != nil {
	log.Println("Could not ReadMessage from connection , err:" , err)
	continue
	}
	siteName:= string(readMessage)
	for {
		time.sleep(1 * time.Second)
		message:=Message{
			count:rand.Intnlen(501),
			siteName: siteName,
		}
		 err = h.send(message)
		 if err != nil {
			log.Println("Could not send Message to connection , err:" , err)
			continue
		 }
	}
}
}
func (h *hub) receive() (messageType int , message []byte , err error) {
 messageType,message,err := h.conn.ReadMessage()
 if err != nil {
	return 0, []byte(),err
 }
 return
}
func (h *hub) send (message Message) (err error) {
 err:= h.conn.WriteJSON(message)
 return
}