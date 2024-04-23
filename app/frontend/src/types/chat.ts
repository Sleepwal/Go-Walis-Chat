export type Room = {
  id: string,
  name: string,
}

export type User = {
  username: string
  id: string
}

export type Message = {
  client_id: string,
  username: string,
  room_id: string,
  content: string,
  type: 'receive' | 'self'
}