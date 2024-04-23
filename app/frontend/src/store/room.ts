import {defineStore} from "pinia";
import {ref} from "vue";
import {Room} from "@/types/chat";

export const useRoomStore = defineStore('room', () => {
    const rooms = ref<Room[] | null>(null)

    const setRooms = (newRooms: Room[]) => {
        rooms.value = newRooms
    }

    const addRoom = (newRoom: Room) => {
        if (rooms.value === null) {
            rooms.value = [newRoom]
            return
        }
        rooms.value.push(newRoom)
    }

    return { rooms, setRooms, addRoom }
}, {
    persist: true,
} )