import {createApp} from 'vue'
import App from './App.vue'
import './style/index.css';
import router from "./router";
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

const app = createApp(App)

app.use(router)

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)
app.use(pinia)

app.mount('#app')