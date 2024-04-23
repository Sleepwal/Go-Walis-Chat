<script setup lang="ts">
import {darkTheme, NIcon} from 'naive-ui'
import {Component, h, ref} from 'vue'
import {BuiltInGlobalTheme} from "naive-ui/es/themes/interface";
import {BookOutline as BookIcon} from '@vicons/ionicons5'

function renderIcon(icon: Component) {
  return () => h(NIcon, null, {default: () => h(icon)})
}

const theme = ref<BuiltInGlobalTheme | null>(darkTheme)

const menuOptions = [
  {
    label: '且听风吟',
    key: 'hear-the-wind-sing',
    icon: renderIcon(BookIcon)
  },
]
</script>

<template>
  <n-config-provider :theme="theme">
    <n-card class="h-dvh">
      <n-layout sider-placement="right" class="">
        <n-layout>
          <n-scrollbar>
            <n-card>
              <n-space>
                <n-button @click="theme = darkTheme">深色</n-button>
                <n-button @click="theme = null">浅色</n-button>
              </n-space>
            </n-card>

            <router-view/>
          </n-scrollbar>

        </n-layout>
<!--        <n-layout-sider-->
<!--            bordered-->
<!--            show-trigger-->
<!--            collapse-mode="width"-->
<!--            :collapsed-width="64"-->
<!--            :width="240"-->
<!--            :native-scrollbar="false"-->
<!--        >-->
<!--          <n-menu-->
<!--              :collapsed-width="64"-->
<!--              :collapsed-icon-size="22"-->
<!--              :options="menuOptions"-->
<!--          />-->
<!--        </n-layout-sider>-->
      </n-layout>
    </n-card>
  </n-config-provider>
</template>

<style scoped>

</style>