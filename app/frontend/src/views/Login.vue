<script setup lang="ts">

import {useUserStore} from "@/store/user";
import {ref} from "vue";
import {useRouter} from "vue-router";

const userStore = useUserStore()
const username = ref<string>('')

const router = useRouter()
const login = async () => {
  userStore.setUser(username.value, Math.random().toString(36))
  await router.push({name: 'room'})
}
</script>

<template>
  <div class="flex items-center justify-center h-96">
    <div class="flex flex-col min-w-80">
      <div class="text-3xl font-bold text-center">
        <span class="text-blue">Welcome</span>
      </div>

      <n-input
          v-model:value="username"
          @keydown.enter="login"
          placeholder="please input your username"
          class="p-3 mt-8 rounded-md" />
      <n-button
          class="p-3 mt-6 rounded-md font-bold-white"
          @click="login">Login</n-button>
    </div>
  </div>
</template>

<style scoped>

</style>