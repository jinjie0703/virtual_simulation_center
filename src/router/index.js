import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '@/views/home_page/HomePage.vue'
import InformationCenter from '@/views/information_center/InformationCenter.vue'
import AchievementShowcase from '@/views/achievement_showcase/AchievementShowcase.vue'
import OurTeam from '@/views/our_team/OurTeam.vue'
import AboutUs from '@/views/about_us/AboutUs.vue'

// 创建并直接导出 router 实例
const router = createRouter({
  // 使用 Vite 的标准方式获取基础 URL
  history: createWebHistory(),

  // 路由规则定义
  routes: [
    {
      path: '/',
      redirect: '/home_page', // 默认重定向到首页
    },
    {
      path: '/home_page',
      name: 'home-page',
      component: HomePage,
      // 此页面会默认显示页脚
    },
    {
      path: '/information_center',
      name: 'information-center',
      component: InformationCenter,
      // 此页面会默认显示页脚
    },
    {
      path: '/achievement_showcase',
      name: 'achievement-showcase',
      component: AchievementShowcase,
      // 此页面会默认显示页脚
    },
    {
      path: '/our_team',
      name: 'our-team',
      component: OurTeam,
      // 此页面会默认显示页脚
    },
    {
      path: '/about_us',
      name: 'about-us',
      component: AboutUs,
      // 此页面会默认显示页脚
    },
  ],
})

export default router
