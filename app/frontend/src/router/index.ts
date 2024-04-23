import {createRouter, createWebHashHistory, createWebHistory} from "vue-router";

const constantRoutes = [
  {
    path: '/',
    name: 'Home',
    redirect:'/login',
  }, {
    path: "/login",
    name: "Login",
    component: () => import("@/views/Login.vue"),
  }, {
    path: "/room",
    name: "Room",
    component: () => import("@/views/Room.vue"),
  }, {
    path: "/app/:roomId",
    name: "App",
    component: () => import("@/views/app/index.vue"),
  },
]
const router = createRouter({
  history: createWebHashHistory("#"),
  routes: constantRoutes
})
export default router;