import request from '@/utils/request';

export type JoinRoomParams = {
  room_id: string;
  user_id: string;
  username: string;
}
export function joinRoom(params: JoinRoomParams) {

}

export type CreateRoomParams = {
  room_id: string;
  room_name: string;
}
export function  createRoom(data: CreateRoomParams) {
  return request.post('/chat/create_room', data)
}

export function getRooms() {
    return request.get('/chat/room_list')
}

export function getClientsInRoom(roomId: string) {
    return request.get('/chat/client_list',
        {params: {room_id: roomId}})
}