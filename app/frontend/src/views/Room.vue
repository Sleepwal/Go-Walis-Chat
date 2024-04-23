<script setup lang="ts">

import {useRoomStore} from "@/store/room";
import {ref, toRefs} from "vue";
import {createRoom, getRooms} from "@/api/chat";
import router from "@/router";
import {useUserStore} from "@/store/user";

const { user } = useUserStore()
const { rooms } = toRefs(useRoomStore())
const { setRooms, addRoom } = useRoomStore()

const roomName = ref('')

// 创建新房间
const createNewRoom = async () => {
  if (roomName.value.trim() === '') {
    alert('Please enter a room name')
    return
  }
  await createRoom({
    room_id: Math.random().toString(36),
    room_name: roomName.value.trim()
  })

 await updateRooms()
}

// 更新房间列表
const updateRooms = async () => {
  const res = await getRooms()
  setRooms(res.data.rooms)
  console.log(res.data.rooms)
}
updateRooms()

// 加入房间
const joinRoom = (roomId: string) => {
  router.push('/app/' + roomId)
}

</script>

<template>
  <div class="my-8 px-4 md:mx-32 w-full h-full">
    <div>{{ user }}</div>
    <div class="flex justify-center mt-3 p-5">
      <input v-model="roomName" type="text"
             placeholder="Enter Room Name"
             class="rounded-md p-3 border border-grey focus:outline-none focus:border-blue" />
      <button @click="createNewRoom"
              class="bg-blue border text-white rounded-md p-2 md:ml-4">
        Create Room
      </button>
    </div>

    <div class="mt-6">
      <div class="font-bold">Available Rooms</div>

      <div class="grid grid-cols-1 md:grid-cols-5 gap-4 mt-6">
        <div v-for="room in rooms" :key="room.id">
          <div class="border border-blue p-4 flex items-center rounded-md w-full">
            <div class="w-full">
              <div class="text-sm">Room</div>
              <div class="text-blue font-bold text-lg">{{ room.name }}</div>

              <button @click="joinRoom(room.id)"
                      class="px-4 text-white bg-blue rounded-md">
                Join
              </button>
            </div>
          </div>
        </div>
      </div>

    </div>
  </div>
</template>

<style scoped>

</style>