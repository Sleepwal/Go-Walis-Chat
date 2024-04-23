import {defineStore} from "pinia";
import {reactive} from "vue";
import {User} from "@/types/chat";

export const useUserStore = defineStore('user', () => {
    const user = reactive<User>({
        username: '',
        id: ''
    })

    const setUser = (username: string, id: string) => {
        user.username = username
        user.id = id
    }

    return { user, setUser }
},{
    persist: true,
})