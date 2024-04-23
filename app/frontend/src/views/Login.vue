<script setup lang="ts">

import {useUserStore} from "@/store/user";
import {ref} from "vue";
import router from "@/router";

const userStore = useUserStore()
const username = ref<string>('')
const login = () => {
  userStore.setUser(username.value, Math.random().toString(36))
  router.push('/room')
}
</script>

<template>
  <div class="flex items-center justify-center min-w-full min-h-screen">
    <form class="flex flex-col min-w-80">
      <div class="text-3xl font-bold text-center">
        <span class="text-blue">Welcome</span>
      </div>

      <input v-model="username"
          placeholder="please input your username"
          class="p-3 mt-8 rounded-md border-2 border-grey focus:outline-none focus:border-blue" />
      <button
          class="p-3 mt-6 rounded-md bg-blue font-bold text-white"
          type="submit"
          @click="login">Login</button>
    </form>
  </div>
</template>

<style scoped>

</style>