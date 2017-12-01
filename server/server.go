package server

import(
	"github.com/diegoseso/parchis/models"
	"github.com/davecgh/go-spew/spew"
)

type Server struct{

}

func(S *Server)Run(){

	InitializeRoomsFromConfig()
	IniitializeBoardsOnRoom()
}

func IniitializeBoardsOnRoom(){

}

func InitializeRoomsFromConfig(){
	room := &models.Room{}
	AddBoardToRoom(room)
}

func AddBoardToRoom(Room *models.Room){
	spew.Dump(Room)
}