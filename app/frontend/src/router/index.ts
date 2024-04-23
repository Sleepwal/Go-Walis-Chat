import {createMemoryHistory, createRouter, createWebHashHistory, createWebHistory} from "vue-router";

const constantRoutes = [
  {
    path: '/',
    name: 'home',
    redirect: '/login',
    component: () => import("@/layout/index.vue"),
    children: [
      {
        path: "/login",
        name: "login",
        component: () => import("@/views/Login.vue"),
      }, {
        path: "/room",
        name: "room",
        component: () => import("@/views/Room.vue"),
      }, {
        path: "/app/:roomId",
        name: "App",
        component: () => import("@/views/app/index.vue"),
      },
    ],
  },
]
const router = createRouter({
  history: createWebHashHistory(),
  scrollBehavior: () => ({ top: 0 }),
  routes: constantRoutes
})
export default router;