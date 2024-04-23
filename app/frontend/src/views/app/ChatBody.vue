<script setup lang="ts">

import {Message} from "@/types/chat";
import {PropType} from "vue";

const props = defineProps({
  data: {
    type: Array as PropType<Array<Message> | null >,
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
          <n-tag type="success" class="p-4 rounded-md">
            <n-gradient-text gradient="linear-gradient(to right, #12c2e9, #c471ed, #f64f59">
              {{item.content}}
            </n-gradient-text>
          </n-tag>
        </div>
      </div>

      <div v-else
           class="mt-2"
      >
        <div class="text-sm">{{item.username}}</div>

        <div>
          <n-tag class="p-4 rounded-md">
            <n-gradient-text type="info">
              {{item.content}}
            </n-gradient-text>
          </n-tag>
        </div>
      </div>
    </div>
  </div>

</template>

<style scoped>

</style>