<script setup lang="ts">

import {WEBSOCKET_URL} from "@/utils/consts";
import {useRoute} from "vue-router";
import {useUserStore} from "@/store/user";
import {onMounted, onUnmounted, ref, toRefs} from "vue";
import ChatBody from "@/views/app/ChatBody.vue";
import {useMessageStore} from "@/store/message";

const { messages, users } = toRefs(useMessageStore())
const { setUsers, addUser, addMessage, setMessages } = useMessageStore()
const { user } = useUserStore()

// 创建websocket连接
const roomId = useRoute().params.roomId
const ws = new WebSocket(`${WEBSOCKET_URL}/chat/join_room?room_id=${roomId}&user_id=${user.id}&user_name=${user.username}`)

const handleConnect = () => {
  ws.onmessage = (message) => {
    const m = JSON.parse(message.data)
    // console.log(m)

    // 加入房间
    if (m.content.includes('joined the room')) {
      addUser({ id: user.id, username: m.username })
    }

    // 离开房间
    if (m.content.includes('left the room')) {
      const deleteUser = users.value?.filter(u => u.username !== m.username)
      if (deleteUser === undefined) return
        setUsers([...deleteUser])

      addMessage({
        client_id: user.id,
        username: m.username,
        room_id: m.room_id,
        content: m.content,
        type: 'self'
      })
      return
    }

    m.type = user.username === m.username? 'self' : 'receive'
    addMessage({
      client_id: user.id,
      username: m.username,
      room_id: m.room_id,
      content: m.content,
      type: m.type
      }
    )

    console.log(users.value)
    console.log(messages.value)
  }

  ws.onopen = () => {
    console.log('websocket connected')
  }
  ws.onclose = () => {
    console.log('websocket closed')
  }
  ws.onerror = (error) => {
    console.log('websocket error', error)
  }
}
onMounted(() => {
  handleConnect()
})
onUnmounted(() => {
  ws.close()
})


const content = ref('')
const sendMessage = () => {
  console.log(content.value)
  if (!content.value.trim()) {
    return
  }

  ws.send(content.value.trim())
  content.value = ''
}

</script>

<template>
  <div class="flex flex-col w-full">
    <n-scrollbar style="height: 70vh">
      <div class="p-4 md:mx-6 mb-14">
        <ChatBody :data="messages" />
      </div>
    </n-scrollbar>

    <div class="fixed bottom-0 mt-4 w-full">
      <div class="flex md:flex-row px-4 py-2 md:mx-4 rounded-md">
        <div class="flex w-full mr-4 rounded-md border border-blue">
          <n-input
              v-model:value="content"
              type="textarea"
              placeholder="type your message here"
              class=" p-2 rounded-md"
              :autosize="{ minRows: 1 }"
          />
        </div>

        <div class="flex items-center">
          <n-button @click="sendMessage" type="success"
                  class="p-2 rounded-md">
            Send
          </n-button>
        </div>

      </div>

    </div>

  </div>
</template>

<style scoped>

</style>