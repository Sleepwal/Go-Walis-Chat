import {defineStore} from "pinia";
import {ref} from "vue";
import {Message, User} from "@/types/chat";

export const useMessageStore = defineStore('message', () => {
  const messages = ref<Message[] | null>(null)
  function setMessages(newMessages: Message[]) {
    messages.value = newMessages
  }
  const addMessage = (newMessage: Message) => {
    if (messages.value === null) {
      setMessages([newMessage])
      return
    }
    messages.value.push(newMessage)
  }

  const users = ref<User[] | null>(null)
  const setUsers = (newUsers: User[]) => {
    users.value = newUsers
  }
  const addUser = (newUser: User) => {
    if (users.value === null) {
      setUsers([newUser])
      return
    }
    users.value.push(newUser)
  }

  return { messages, setMessages, users, setUsers, addMessage, addUser }
})