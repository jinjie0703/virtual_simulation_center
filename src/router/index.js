import LoginAndRegister from '@/views/login_and_register/LoginAndRegister.vue'
import HomePage from '@/views/home_page/HomePage.vue'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home', // 建议使用小写 name
      component: HomePage,
    },
    { path: '/login_and_register', name: 'LoginAndRegister', component: LoginAndRegister }, // 添加登录/注册页面的路由
  ],
})

export default router
