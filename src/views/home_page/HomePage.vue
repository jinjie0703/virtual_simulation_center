<template>
  <div class="homepage">
    <Carousel :slides="carouselData" />
    <div class="main-content">
      <FeatureSection :features="featureData" />
      <UpdatesSection :newsItems="newsData" :competitionItems="competitionData" />
      <PartnersCarousel :companies="companiesData" :schools="schoolsData" :scrollSpeed="1.2" />
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

// 导入子组件
import Carousel from '@/components/home_page/HomeCarousel.vue'
import FeatureSection from '@/components/home_page/FeatureSection.vue'
import UpdatesSection from '@/components/home_page/UpdatesSection.vue'
import PartnersCarousel from '@/components/home_page/PartnersCarousel.vue'

// 1. 轮播图的数据
const carouselData = ref([
  {
    title: '虚拟仿真实验平台',
    subtitle: '构建沉浸式学习环境，开启教育新篇章',
    imageUrl:
      'https://images.unsplash.com/photo-1542370285-b8eb8317691c?q=80&w=1974&auto=format&fit=crop',
  },
  {
    title: '智能化教学系统',
    subtitle: 'AI驱动的个性化学习体验',
    imageUrl:
      'https://images.unsplash.com/photo-1519010470956-6d877008eaa4?q=80&w=2070&auto=format&fit=crop',
  },
  {
    title: '数字化实验室',
    subtitle: '打破时空限制，让实验无处不在',
    imageUrl:
      'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
  },
])

// 3. 成果展示的数据
// ...existing code...
const featureData = ref([
  {
    id: 1,
    title: '虚拟现实教学平台',
    description:
      '基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。',
    image:
      'https://images.unsplash.com/photo-1592478411213-6153e4ebc696?q=80&w=1200&auto=format&fit=crop',
    tags: ['VR技术', '教学创新', '沉浸式体验'],

    // 作者信息
    author: '张教授',
    authorTitle: '虚拟现实研究中心主任',
    authorAvatar:
      'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?q=80&w=150&auto=format&fit=crop',
    contact1: '邮箱zhang@university.edu.cn',
    contact2: '电话123-456-7890',
    projectUrl: 'https://github.com/university/vr-teaching-platform',

    // 视频和图片库
    videoUrl: 'http://vjs.zencdn.net/v/oceans.mp4',
    gallery: [
      'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
    ],
  },
  {
    id: 2,
    title: 'AI智能实验助手',
    description:
      '集成人工智能算法的智能实验辅导系统，能够实时分析学生的实验操作，提供个性化指导建议，大幅提升实验教学效果和安全性。',
    image:
      'https://images.unsplash.com/photo-1485827404703-89b55fcc595e?q=80&w=1200&auto=format&fit=crop',
    tags: ['人工智能', '智能辅导', '数据分析'],

    author: '李博士',
    authorTitle: 'AI实验室技术总监',
    authorAvatar:
      'https://images.unsplash.com/photo-1494790108755-2616b612b786?q=80&w=150&auto=format&fit=crop',
    contact: 'li@university.edu.cn',
    projectUrl: 'https://github.com/university/ai-lab-assistant',

    videoUrl: 'https://example.com/video2.mp4',
    gallery: [
      'https://images.unsplash.com/photo-1485827404703-89b55fcc595e?q=80&w=400&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1581092160562-40aa08e78837?q=80&w=400&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1518709268805-4e9042af2176?q=80&w=400&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1592478411213-6153e4ebc696?q=80&w=400&auto=format&fit=crop',
    ],
  },
  {
    id: 3,
    title: '数字孪生实验室',
    description:
      '基于数字孪生技术构建的虚拟实验室，完美复现真实实验环境的每一个细节，支持大规模并发实验，解决了传统实验室资源限制的问题。',
    image:
      'https://images.unsplash.com/photo-1518709268805-4e9042af2176?q=80&w=1200&auto=format&fit=crop',
    tags: ['数字孪生', '云计算', '大数据'],

    author: '王教授',
    authorTitle: '数字化实验室主任',
    authorAvatar:
      'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?q=80&w=150&auto=format&fit=crop',
    contact: 'wang@university.edu.cn',
    projectUrl: 'https://github.com/university/digital-twin-lab',

    videoUrl: 'https://example.com/video3.mp4',
    gallery: [
      'https://images.unsplash.com/photo-1518709268805-4e9042af2176?q=80&w=400&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=400&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1485827404703-89b55fcc595e?q=80&w=400&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1592478411213-6153e4ebc696?q=80&w=400&auto=format&fit=crop',
    ],
  },
  {
    id: 4,
    title: '智能化教学管理系统',
    description:
      '基于大数据和机器学习的智能教学管理平台，实现个性化学习路径推荐、学习效果评估和教学资源智能分配。',
    image:
      'https://images.unsplash.com/photo-1551288049-bebda4e38f71?q=80&w=1200&auto=format&fit=crop',
    tags: ['智能管理', '大数据', '个性化学习'],

    author: '陈博士',
    authorTitle: '智能教育技术研究院院长',
    authorAvatar:
      'https://images.unsplash.com/photo-1560250097-0b93528c311a?q=80&w=150&auto=format&fit=crop',
    contact: 'chen@university.edu.cn',
    projectUrl: 'https://github.com/university/smart-teaching-system',

    videoUrl: 'https://media.w3.org/2010/05/sintel/trailer.mp4',
    gallery: [
      'https://images.unsplash.com/photo-1551288049-bebda4e38f71?q=80&w=400&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1460925895917-afdab827c52f?q=80&w=400&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1519389950473-47ba0277781c?q=80&w=400&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1504868584819-f8e8b4b6d7e3?q=80&w=400&auto=format&fit=crop',
    ],
  },
])
// ...existing code...

// 4. 新闻与竞赛的数据
const newsData = ref([
  {
    title: '我校虚拟仿真中心获得国家级教学成果奖',
    date: '2024-05-20',
    url: '#',
    category: '学术荣誉',
  },
  {
    title: '新一代VR教学平台正式启用',
    date: '2024-05-15',
    url: '#',
    category: '技术更新',
  },
  {
    title: '国际虚拟现实教育大会在我校举办',
    date: '2024-05-10',
    url: '#',
    category: '学术会议',
  },
  {
    title: 'AI实验助手项目获得科技进步奖',
    date: '2024-05-05',
    url: '#',
    category: '科研成果',
  },
  {
    title: '数字孪生实验室建设完成验收',
    date: '2024-04-28',
    url: '#',
    category: '基础建设',
  },
  {
    title: '虚拟仿真教学资源库正式上线',
    date: '2024-04-20',
    url: '#',
    category: '资源共享',
  },
  {
    title: '混合现实技术应用于临床医学教学',
    date: '2024-04-15',
    url: '#',
    category: '应用拓展',
  },
])

const competitionData = ref([
  {
    title: '全国大学生虚拟现实设计大赛报名开始',
    date: '2024-05-18',
    url: '#',
    deadline: '2024-06-30',
    level: '国家级',
  },
  {
    title: '国际AI创新应用竞赛征集作品',
    date: '2024-05-12',
    url: '#',
    deadline: '2024-06-15',
    level: '国际级',
  },
  {
    title: '数字化教学创新大赛开放注册',
    date: '2024-05-09',
    url: '#',
    deadline: '2024-06-20',
    level: '省级',
  },
  {
    title: '混合现实技术挑战赛即将开幕',
    date: '2024-05-02',
    url: '#',
    deadline: '2024-05-25',
    level: '校级',
  },
  {
    title: '智能实验设计竞赛作品征集中',
    date: '2024-04-25',
    url: '#',
    deadline: '2024-06-10',
    level: '国家级',
  },
  {
    title: '虚拟仿真课程设计比赛启动',
    date: '2024-04-18',
    url: '#',
    deadline: '2024-05-30',
    level: '省级',
  },
])
const schoolsData = ref([
  {
    id: 1,
    name: '哈尔滨工业大学',
    logo: 'https://placehold.co/80x80/003366/FFFFFF?text=HIT',
    url: '#',
  },
  {
    id: 2,
    name: '华南理工大学',
    logo: 'https://placehold.co/80x80/007A5E/FFFFFF?text=SCUT',
    url: '#',
  },
  { id: 3, name: '南京大学', logo: 'https://placehold.co/80x80/5A007B/FFFFFF?text=NJU', url: '#' },
  {
    id: 4,
    name: '四川农业大学',
    logo: 'https://placehold.co/80x80/008080/FFFFFF?text=SICAU',
    url: '#',
  },
  { id: 5, name: '中南大学', logo: 'https://placehold.co/80x80/005BAC/FFFFFF?text=CSU', url: '#' },
  {
    id: 6,
    name: '中央民族大学',
    logo: 'https://placehold.co/80x80/D8232A/FFFFFF?text=MUC',
    url: '#',
  },
  {
    id: 7,
    name: '北京理工大学',
    logo: 'https://placehold.co/80x80/008000/FFFFFF?text=BIT',
    url: '#',
  },
  {
    id: 8,
    name: '大连海事大学',
    logo: 'https://placehold.co/80x80/004B8D/FFFFFF?text=DMU',
    url: '#',
  },
  { id: 9, name: '嘉应学院', logo: 'https://placehold.co/80x80/4682B4/FFFFFF?text=JYU', url: '#' },
  { id: 10, name: '吉林大学', logo: 'https://placehold.co/80x80/002D62/FFFFFF?text=JLU', url: '#' },
  {
    id: 11,
    name: '辽宁石油化工大学',
    logo: 'https://placehold.co/80x80/00537E/FFFFFF?text=LNPU',
    url: '#',
  },
  {
    id: 12,
    name: '北华大学',
    logo: 'https://placehold.co/80x80/006400/FFFFFF?text=BEIHUA',
    url: '#',
  },
])

// 模拟的合作企业数据 (如果您不需要，可以传入一个空数组 [])
const companiesData = ref([
  {
    id: 1,
    name: '科技公司 A',
    logo: 'https://placehold.co/80x80/1A1A1A/FFFFFF?text=Tech+A',
    url: '#',
  },
  {
    id: 2,
    name: '创新企业 B',
    logo: 'https://placehold.co/80x80/FF6347/FFFFFF?text=Innovate+B',
    url: '#',
  },
  {
    id: 3,
    name: '互联网巨头 C',
    logo: 'https://placehold.co/80x80/1E90FF/FFFFFF?text=Web+C',
    url: '#',
  },
  {
    id: 4,
    name: '金融集团 D',
    logo: 'https://placehold.co/80x80/FFD700/000000?text=Finance+D',
    url: '#',
  },
  {
    id: 5,
    name: '制造业 E',
    logo: 'https://placehold.co/80x80/B0B0B0/FFFFFF?text=Manu+E',
    url: '#',
  },
  {
    id: 6,
    name: '初创公司 F',
    logo: 'https://placehold.co/80x80/32CD32/FFFFFF?text=Startup+F',
    url: '#',
  },
  {
    id: 1,
    name: '科技公司 A',
    logo: 'https://placehold.co/80x80/1A1A1A/FFFFFF?text=Tech+A',
    url: '#',
  },
  {
    id: 2,
    name: '创新企业 B',
    logo: 'https://placehold.co/80x80/FF6347/FFFFFF?text=Innovate+B',
    url: '#',
  },
  {
    id: 3,
    name: '互联网巨头 C',
    logo: 'https://placehold.co/80x80/1E90FF/FFFFFF?text=Web+C',
    url: '#',
  },
  {
    id: 4,
    name: '金融集团 D',
    logo: 'https://placehold.co/80x80/FFD700/000000?text=Finance+D',
    url: '#',
  },
  {
    id: 5,
    name: '制造业 E',
    logo: 'https://placehold.co/80x80/B0B0B0/FFFFFF?text=Manu+E',
    url: '#',
  },
  {
    id: 6,
    name: '初创公司 F',
    logo: 'https://placehold.co/80x80/32CD32/FFFFFF?text=Startup+F',
    url: '#',
  },
  {
    id: 1,
    name: '科技公司 A',
    logo: 'https://placehold.co/80x80/1A1A1A/FFFFFF?text=Tech+A',
    url: '#',
  },
  {
    id: 2,
    name: '创新企业 B',
    logo: 'https://placehold.co/80x80/FF6347/FFFFFF?text=Innovate+B',
    url: '#',
  },
  {
    id: 3,
    name: '互联网巨头 C',
    logo: 'https://placehold.co/80x80/1E90FF/FFFFFF?text=Web+C',
    url: '#',
  },
  {
    id: 4,
    name: '金融集团 D',
    logo: 'https://placehold.co/80x80/FFD700/000000?text=Finance+D',
    url: '#',
  },
  {
    id: 5,
    name: '制造业 E',
    logo: 'https://placehold.co/80x80/B0B0B0/FFFFFF?text=Manu+E',
    url: '#',
  },
  {
    id: 6,
    name: '初创公司 F',
    logo: 'https://placehold.co/80x80/32CD32/FFFFFF?text=Startup+F',
    url: '#',
  },
])
</script>

<style scoped>
.homepage {
  margin: 0;
  padding: 0;
  overflow-x: hidden;
  width: 100%;
  position: relative;
  scroll-behavior: auto;
}

.main-content {
  position: relative;
  z-index: 1;
  margin-top: 0;
  padding: 0;
  display: block;
}

/* 确保轮播组件完全占满且无边距 */
:deep(.hero-carousel) {
  margin: 0;
  padding: 0;
  display: block;
}

:deep(.transition-section) {
  margin: 0;
  padding: 0;
  display: block;
}

/* 确保特性区域正确连接 */
:deep(.features-section) {
  margin-top: 0 !important;
  position: relative;
  z-index: 1;
}

/* 确保更新区域正确连接 */
:deep(.latest-updates) {
  margin-top: 0 !important;
  position: relative;
  z-index: 1;
}

/* 全局滚动优化 */
* {
  scroll-behavior: smooth;
}

/* 响应式调整 */
@media (max-width: 1024px) {
  :deep(.transition-section) {
    margin-top: -80px !important;
  }
}

@media (max-width: 768px) {
  .homepage {
    -webkit-overflow-scrolling: touch;
  }

  :deep(.transition-section) {
    margin-top: -60px !important;
  }
}

@media (max-width: 480px) {
  :deep(.transition-section) {
    margin-top: -40px !important;
  }
}

/* 加载动画 */
.homepage {
  animation: fadeIn 0.5s ease-in-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

/* 确保页面内容不被其他元素遮挡 */
.main-content > * {
  position: relative;
}
</style>
