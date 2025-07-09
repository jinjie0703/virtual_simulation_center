import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '@/views/home_page/HomePage.vue'
import LoginAndRegister from '@/views/login_and_register/LoginAndRegister.vue'
import InformationCenter from '@/views/information_center/InformationCenter.vue'
import AboutUs from '@/views/about_us/AboutUs.vue'
import OurTeam from '@/views/our_team/OurTeam.vue'

// 创建并直接导出 router 实例
const router = createRouter({
  // 使用 Vite 的标准方式获取基础 URL
  history: createWebHistory(),

  // 路由规则定义
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomePage,
      // 此页面会默认显示页脚
    },
    {
      path: '/login_and_register',
      name: 'login-and-register', // 建议将 name 改为小写 kebab-case 风格
      component: LoginAndRegister,
      // 在此明确指出此页面不显示页脚
      meta: { showFooter: false },
    },
    {
      path: '/information_center',
      name: 'information-center',
      component: InformationCenter,
      // 此页面会默认显示页脚
    },
    {
      path: '/about_us',
      name: 'about-us',
      component: AboutUs,
      // 此页面会默认显示页脚
    },
    {
      path: '/our_team',
      name: 'our-team',
      component: OurTeam,
      // 此页面会默认显示页脚
    },
  ],
})

export default router
