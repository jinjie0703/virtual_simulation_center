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

// 3. 成果展示的数据
// ...existing code...
const featureData = ref([
  {
    id: 1,
    title: '虚拟现实教学平台',
    description:
      '基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。',
    image:
      'https://images.unsplash.com/photo-1551288049-bebda4e38f71?q=80&w=1200&auto=format&fit=crop',
    tags: ['VR技术', '教学创新', '沉浸式体验'],
    author: '张教授',
    authorTitle: '虚拟现实研究中心主任',
    authorAvatar:
      'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?q=80&w=150&auto=format&fit=crop',
    contact1: '邮箱zhang@university.edu.cn',
    contact2: '电话123456789',
    projectUrl: 'https://github.com',

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
    id: 1,
    title: '虚拟现实教学平台',
    description:
      '基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。',
    image:
      'https://images.unsplash.com/photo-1551288049-bebda4e38f71?q=80&w=1200&auto=format&fit=crop',
    tags: ['VR技术', '教学创新', '沉浸式体验'],
    author: '张教授',
    authorTitle: '虚拟现实研究中心主任',
    authorAvatar:
      'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?q=80&w=150&auto=format&fit=crop',
    contact1: '邮箱zhang@university.edu.cn',
    contact2: '电话123456789',
    projectUrl: 'https://github.com',

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
    id: 1,
    title: '虚拟现实教学平台',
    description:
      '基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。',
    image:
      'https://images.unsplash.com/photo-1551288049-bebda4e38f71?q=80&w=1200&auto=format&fit=crop',
    tags: ['VR技术', '教学创新', '沉浸式体验'],
    author: '张教授',
    authorTitle: '虚拟现实研究中心主任',
    authorAvatar:
      'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?q=80&w=150&auto=format&fit=crop',
    contact1: '邮箱zhang@university.edu.cn',
    contact2: '电话123456789',
    projectUrl: 'https://github.com',

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
    id: 1,
    title: '虚拟现实教学平台',
    description:
      '基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。',
    image:
      'https://images.unsplash.com/photo-1551288049-bebda4e38f71?q=80&w=1200&auto=format&fit=crop',
    tags: ['VR技术', '教学创新', '沉浸式体验'],
    author: '张教授',
    authorTitle: '虚拟现实研究中心主任',
    authorAvatar:
      'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?q=80&w=150&auto=format&fit=crop',
    contact1: '邮箱zhang@university.edu.cn',
    contact2: '电话123456789',
    projectUrl: 'https://github.com',

    // 视频和图片库
    videoUrl: 'http://vjs.zencdn.net/v/oceans.mp4',
    gallery: [
      'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
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
    name: 'google',
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/2f/Google_2015_logo.svg/1920px-Google_2015_logo.svg.png',
    url: '#',
  },
  {
    id: 1,
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/2f/Google_2015_logo.svg/1920px-Google_2015_logo.svg.png',
    url: '#',
  },
  {
    id: 1,
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/2f/Google_2015_logo.svg/1920px-Google_2015_logo.svg.png',
    url: '#',
  },
  {
    id: 4,
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Microsoft_logo_%282012%29.svg/1920px-Microsoft_logo_%282012%29.svg.png',
    url: '#',
  },
  {
    id: 1,
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/2f/Google_2015_logo.svg/1920px-Google_2015_logo.svg.png',
    url: '#',
  },
  {
    id: 1,
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/2f/Google_2015_logo.svg/1920px-Google_2015_logo.svg.png',
    url: '#',
  },
  {
    id: 3,
    name: '互联网巨头 C',
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/ca/Sony_logo.svg/2560px-Sony_logo.svg.png',
    url: '#',
  },
  {
    id: 4,
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Microsoft_logo_%282012%29.svg/1920px-Microsoft_logo_%282012%29.svg.png',
    url: '#',
  },
  {
    id: 3,
    name: '互联网巨头 C',
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/ca/Sony_logo.svg/2560px-Sony_logo.svg.png',
    url: '#',
  },
  {
    id: 5,
    name: '制造业 E',
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/51/IBM_logo.svg/2560px-IBM_logo.svg.png',
    url: '#',
  },
  {
    id: 4,
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Microsoft_logo_%282012%29.svg/1920px-Microsoft_logo_%282012%29.svg.png',
    url: '#',
  },
  {
    id: 5,
    name: '制造业 E',
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/51/IBM_logo.svg/2560px-IBM_logo.svg.png',
    url: '#',
  },
  {
    id: 5,
    name: '制造业 E',
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/51/IBM_logo.svg/2560px-IBM_logo.svg.png',
    url: '#',
  },
  {
    id: 4,
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Microsoft_logo_%282012%29.svg/1920px-Microsoft_logo_%282012%29.svg.png',
    url: '#',
  },
  {
    id: 5,
    name: '制造业 E',
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/51/IBM_logo.svg/2560px-IBM_logo.svg.png',
    url: '#',
  },
  {
    id: 4,
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Microsoft_logo_%282012%29.svg/1920px-Microsoft_logo_%282012%29.svg.png',
    url: '#',
  },
  {
    id: 5,
    name: '制造业 E',
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/51/IBM_logo.svg/2560px-IBM_logo.svg.png',
    url: '#',
  },
  {
    id: 4,
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Microsoft_logo_%282012%29.svg/1920px-Microsoft_logo_%282012%29.svg.png',
    url: '#',
  },
])

// 模拟的合作企业数据,名字虽然不用，但是仍需要存在数据库里面
const companiesData = ref([
  {
    id: 1,
    name: 'google',
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/2f/Google_2015_logo.svg/1920px-Google_2015_logo.svg.png',
    url: '#',
  },
  {
    id: 1,
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/2f/Google_2015_logo.svg/1920px-Google_2015_logo.svg.png',
    url: '#',
  },
  {
    id: 1,
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/2f/Google_2015_logo.svg/1920px-Google_2015_logo.svg.png',
    url: '#',
  },
  {
    id: 4,
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Microsoft_logo_%282012%29.svg/1920px-Microsoft_logo_%282012%29.svg.png',
    url: '#',
  },
  {
    id: 1,
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/2f/Google_2015_logo.svg/1920px-Google_2015_logo.svg.png',
    url: '#',
  },
  {
    id: 1,
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/2/2f/Google_2015_logo.svg/1920px-Google_2015_logo.svg.png',
    url: '#',
  },
  {
    id: 3,
    name: '互联网巨头 C',
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/ca/Sony_logo.svg/2560px-Sony_logo.svg.png',
    url: '#',
  },
  {
    id: 4,
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Microsoft_logo_%282012%29.svg/1920px-Microsoft_logo_%282012%29.svg.png',
    url: '#',
  },
  {
    id: 3,
    name: '互联网巨头 C',
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/ca/Sony_logo.svg/2560px-Sony_logo.svg.png',
    url: '#',
  },
  {
    id: 5,
    name: '制造业 E',
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/51/IBM_logo.svg/2560px-IBM_logo.svg.png',
    url: '#',
  },
  {
    id: 4,
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Microsoft_logo_%282012%29.svg/1920px-Microsoft_logo_%282012%29.svg.png',
    url: '#',
  },
  {
    id: 5,
    name: '制造业 E',
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/51/IBM_logo.svg/2560px-IBM_logo.svg.png',
    url: '#',
  },
  {
    id: 5,
    name: '制造业 E',
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/51/IBM_logo.svg/2560px-IBM_logo.svg.png',
    url: '#',
  },
  {
    id: 4,
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Microsoft_logo_%282012%29.svg/1920px-Microsoft_logo_%282012%29.svg.png',
    url: '#',
  },
  {
    id: 5,
    name: '制造业 E',
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/51/IBM_logo.svg/2560px-IBM_logo.svg.png',
    url: '#',
  },
  {
    id: 4,
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Microsoft_logo_%282012%29.svg/1920px-Microsoft_logo_%282012%29.svg.png',
    url: '#',
  },
  {
    id: 5,
    name: '制造业 E',
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/51/IBM_logo.svg/2560px-IBM_logo.svg.png',
    url: '#',
  },
  {
    id: 4,
    logo: 'https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Microsoft_logo_%282012%29.svg/1920px-Microsoft_logo_%282012%29.svg.png',
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
