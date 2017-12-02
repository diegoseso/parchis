package server

import(
	"github.com/diegoseso/parchis/models"
	"github.com/davecgh/go-spew/spew"
	"net"
	"log"
	"io"
	"time"
)

type Server struct{

}

func NewServer() *Server{
	return &Server{}
}

func(S *Server)Run(){
    S.listen()
	S.initializeRoomsFromConfig()
	S.iniitializeBoardsOnRoom()
}

func(S *Server)listen(){
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal()
	}
	for{
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn( c net.Conn){
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Millisecond)
	}
}

func(S *Server)Stop(){
}

func(S *Server)iniitializeBoardsOnRoom(){

}

func(S *Server)initializeRoomsFromConfig(){
	room := &models.Room{}
	AddBoardToRoom(room)
}

func AddBoardToRoom(Room *models.Room){
	spew.Dump(Room)
}