import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '@/views/home_page/HomePage.vue'
import InformationCenter from '@/views/information_center/InformationCenter.vue'
import AchievementShowcase from '@/views/achievement_showcase/AchievementShowcase.vue'
import OurTeam from '@/views/our_team/OurTeam.vue'
import TeamCenter from '@/views/team_center/TeamCenter.vue'
import AboutUs from '@/views/about_us/AboutUs.vue'
import NewsDetail from '@/views/information_center/NewsDetails.vue'

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
      path: '/team_center',
      name: 'team-center',
      component: TeamCenter,
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
    {
      // :id 是动态参数，它会匹配 /news/news-1, /news/any-id 等
      path: '/news/:id',
      name: 'NewsDetail',
      component: NewsDetail,
      props: true, // 允许将路由参数作为 props 传递给组件
    },
    {
      path: '/competitions/:id',
      name: 'CompetitionDetail',
      component: NewsDetail, // 复用详情页组件
      props: true,
    },
    {
      path: '/projects/:id',
      name: 'ProjectDetail',
      component: NewsDetail, // 复用详情页组件
      props: true,
    },
  ],
  scrollBehavior(_to, _from, _savedPosition) {
    // 始终滚动到顶部
    return { top: 0 }
  },
})

export default router
