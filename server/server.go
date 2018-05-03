package server

import(
	"net"
	"log"
	"io"
	"time"
	"net/http"
)

type Server struct{

}

func NewServer() *Server{
	return &Server{}
}

var addr string

func(S *Server)Run(){
    addr = "localhost:8081"

	hub := newHub()
	go hub.run()
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		loginHandler(w, r)
	})
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsHandler(hub, w, r)
	})
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
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
 //@TODO
}