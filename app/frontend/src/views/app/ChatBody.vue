<script setup lang="ts">

import {Message} from "@/types/chat";

const props = defineProps({
  data: {
    type: Array<Message>,
    required: true
  }
})

</script>

<template>
  <div>
    <div v-for="(item, key) in data">
      <div v-if="item.type === 'self'"
           class="flex flex-col mt-2 w-full text-right justify-end"
           :key="key"
      >
        <div class="text-sm">{{item.username}}</div>

        <div>
          <div class="bg-blue text-white px-4 py-1 rounded-md inline-block mt-1">
            {{item.content}}
          </div>
        </div>
      </div>

      <div v-else
           class="mt-2"
      >
        <div class="text-sm">{{item.username}}</div>

        <div>
          <div class="bg-grey text-white px-4 py-1 rounded-md inline-block mt-1">
            {{item.content}}
          </div>
        </div>
      </div>
    </div>
  </div>

</template>

<style scoped>

</style>