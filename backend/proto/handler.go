package proto

import (
	"context"
	"errors"
	"log"
	"time"
)

type User struct {
	id string
	totalNum int32
	roomId string
}

type Room struct {
	id string
	isActive bool
	users []User
}

type Server struct {
	Rooms []Room
}

// room_idとuser_idから今の合計値に+1する
func (s *Server) CountNum(ctx context.Context, p *CountNumparams) (*TotalNum, error) {
	log.Printf("count num")
	roomIdx, err := searchRooms(s.Rooms, p.RoomId)
	if err != nil {
		return nil, err
	}
	room := s.Rooms[roomIdx]
	userIdx, err := searchUsers(room.users, p.UserId)
	if err != nil {
		return nil, err
	}
	s.Rooms[roomIdx].users[userIdx].totalNum += 1
	return &TotalNum{UserId: s.Rooms[roomIdx].users[userIdx].id, RoomId: room.id, Total: s.Rooms[roomIdx].users[userIdx].totalNum}, nil
}

func (s *Server) GetRoomTotalNum(p *GetRoomTotalNumParams, stream AddNumService_GetRoomTotalNumServer) error {
	log.Printf("getRoomTotalNum")
	roomIdx, err := searchRooms(s.Rooms, p.RoomId)
	if err != nil {
		return err
	}
	for {
		totalNums := buildTotalNums(s.Rooms[roomIdx].users)
		if err := stream.Send(&RoomTotalNums{RoomId: s.Rooms[roomIdx].id, TotalNums: totalNums}); err != nil {
			return nil
		}
		time.Sleep(5 * time.Millisecond)
	}
}

func (s *Server) AddRoom(ctx context.Context, p *AddRoomParams) (*RoomInfo, error) {
	log.Printf("Add room")
	s.Rooms = append(s.Rooms, Room{id: p.RoomId, users: []User{}, isActive: true})
	roomIdx, err := searchRooms(s.Rooms, p.RoomId)
	if err != nil {
		return nil, err
	}
	room := s.Rooms[roomIdx]
	return &RoomInfo{RoomId: room.id, IsActive: room.isActive}, nil
}

func (s *Server) JoinRoom(ctx context.Context, p *JoinRoomParams) (*JoinResult, error) {
	log.Printf("join room")
	roomIdx, err := searchRooms(s.Rooms, p.RoomId)
	if err != nil {
		return nil, err
	}
	s.Rooms[roomIdx].users = append(s.Rooms[roomIdx].users, User{id: p.UserId, roomId: p.RoomId, totalNum: 0})
	if (len(s.Rooms[roomIdx].users) > 1) {
		s.Rooms[roomIdx].isActive = false
	}
	return &JoinResult{IsSuccess: true}, nil
}

func (s *Server) GetRooms(ctx context.Context, p *Null) (*RoomList, error) {
	log.Printf("get rooms")
	return &RoomList{Rooms: buildRoomInfo(s.Rooms)}, nil
}

func searchRooms(r []Room, id string) (int, error) {
    for i, v := range(r) {
        if v.id == id {
            return i, nil
        }
    }
    return -1, errors.New("Not found")
}

func searchUsers(r []User, id string) (int, error) {
    for i, v := range(r) {
        if v.id == id {
            return i, nil
        }
    }
    return -1, errors.New("Not found")
}

func buildRoomInfo(rooms []Room) ([]*RoomInfo) {
    r := make([]*RoomInfo, 0)
    for _, v := range(rooms) {
        r = append(r, &RoomInfo{RoomId: v.id, IsActive: v.isActive})
    }
    return r
}

func buildTotalNums(users []User) ([]*TotalNum) {
	u := make([]*TotalNum, 0)
	for _, v := range(users) {
		u = append(u, &TotalNum{UserId: v.id, Total: v.totalNum, RoomId: v.roomId})
	}
	return u
}
