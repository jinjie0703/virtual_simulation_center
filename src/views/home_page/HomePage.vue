<template>
  <div class="homepage">
    <div v-if="loading" class="loading">加载中...</div>
    <div v-if="error" class="error">{{ error }}</div>

    <Carousel :slides="carouselData" />
    <div class="main-content">
      <FeatureSection :features="featureData" />
      <UpdatesSection
        :newsItems="updatesSectionNews"
        :competitionItems="updatesSectionCompetitions"
      />
      <PartnersCarousel :companies="companiesData" :schools="schoolsData" :scrollSpeed="1.2" />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
// import { allMockData } from '@/components/information_center/TestData.js'

// 导入子组件
import Carousel from '@/components/home_page/HomeCarousel.vue'
import FeatureSection from '@/components/home_page/FeatureSection.vue'
import UpdatesSection from '@/components/home_page/UpdatesSection.vue'
import PartnersCarousel from '@/components/home_page/PartnersCarousel.vue'

const featureData = ref([])
const companiesData = ref([])
const schoolsData = ref([])
const updatesSectionNews = ref([])
const updatesSectionCompetitions = ref([])
const loading = ref(false)
const error = ref(null)

// API 基础 URL
const API_BASE_URL = 'https://localhost:8080/api' // 根据你的服务器配置调整

// 获取合作企业数据
const fetchCompanies = async () => {
  try {
    const response = await fetch(`${API_BASE_URL}/PartnerCarousel?type=companies`)
    if (!response.ok) throw new Error('Failed to fetch companies')
    companiesData.value = await response.json()
  } catch (err) {
    console.error('Error fetching companies:', err)
    error.value = err.message
  }
}

// 获取合作学校数据
const fetchSchools = async () => {
  try {
    const response = await fetch(`${API_BASE_URL}/PartnerCarousel?type=schools`)
    if (!response.ok) throw new Error('获取数据失败')
    schoolsData.value = await response.json()
  } catch (err) {
    console.error('Error fetching schools:', err)
    error.value = err.message
  }
}

// 获取新闻动态数据
const fetchUpdatesSectionNews = async () => {
  try {
    const response = await fetch(`${API_BASE_URL}/UpdatesSectionNews`)
    if (!response.ok) throw new Error('Failed to fetch news updates')
    updatesSectionNews.value = await response.json()
  } catch (err) {
    console.error('Error fetching news updates:', err)
    error.value = err.message
  }
}

// 获取竞赛动态数据
const fetchUpdatesSectionCompetitions = async () => {
  try {
    const response = await fetch(`${API_BASE_URL}/UpdatesSectionCompetitions`)
    if (!response.ok) throw new Error('Failed to fetch competitions')
    updatesSectionCompetitions.value = await response.json()
  } catch (err) {
    console.error('Error fetching competitions:', err)
    error.value = err.message
  }
}

// 获取特性区域数据
const fetchFeatureData = async () => {
  try {
    const response = await fetch(`${API_BASE_URL}/FeatureSection`)
    if (!response.ok) throw new Error('Failed to fetch feature data')
    featureData.value = await response.json()
  } catch (err) {
    console.error('Error fetching feature data:', err)
    error.value = err.message
  }
}

// 加载所有数据
// 在 homepage.vue 中

const loadAllData = async () => {
  loading.value = true
  error.value = null

  const promises = [
    fetchFeatureData(),
    fetchCompanies(),
    fetchSchools(),
    fetchUpdatesSectionNews(),
    fetchUpdatesSectionCompetitions(),
  ]

  const results = await Promise.allSettled(promises)

  // 检查每个请求的结果
  results.forEach((result, index) => {
    if (result.status === 'rejected') {
      const promiseNames = [
        'FeatureData',
        'Companies',
        'Schools',
        'UpdatesSectionNews',
        'UpdatesSectionCompetitions',
      ]
      console.error(`Failed to load ${promiseNames[index]}:`, result.reason)
      // 可选：设置一个更通用的错误信息
      if (!error.value) {
        error.value = `部分数据加载失败，请刷新页面重试。`
      }
    }
  })

  loading.value = false
}

// 组件挂载时加载数据
onMounted(() => {
  loadAllData()
})
// const featureData = ref([
//   {
//     id: 1,
//     title: '虚拟现实教学平台',
//     description:
//       '基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。',
//     image:
//       'https://images.unsplash.com/photo-1551288049-bebda4e38f71?q=80&w=1200&auto=format&fit=crop',
//     tags: ['VR技术', '教学创新', '沉浸式体验'],
//     author: '张教授',
//     authorTitle: '虚拟现实研究中心主任',
//     authorAvatar:
//       'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?q=80&w=150&auto=format&fit=crop',
//     contact1: '邮箱zhang@university.edu.cn',
//     contact2: '电话123456789',
//     projectUrl: 'https://github.com',

//     // 视频和图片库
//     videoUrl: 'http://vjs.zencdn.net/v/oceans.mp4',
//     gallery: [
//       'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//       'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//       'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//       'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//     ],
//   },
//   {
//     id: 1,
//     title: '虚拟现实教学平台',
//     description:
//       '基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。',
//     image:
//       'https://images.unsplash.com/photo-1551288049-bebda4e38f71?q=80&w=1200&auto=format&fit=crop',
//     tags: ['VR技术', '教学创新', '沉浸式体验'],
//     author: '张教授',
//     authorTitle: '虚拟现实研究中心主任',
//     authorAvatar:
//       'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?q=80&w=150&auto=format&fit=crop',
//     contact1: '邮箱zhang@university.edu.cn',
//     contact2: '电话123456789',
//     projectUrl: 'https://github.com',

//     // 视频和图片库
//     videoUrl: 'http://vjs.zencdn.net/v/oceans.mp4',
//     gallery: [
//       'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//       'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//       'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//       'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//     ],
//   },
//   {
//     id: 1,
//     title: '虚拟现实教学平台',
//     description:
//       '基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。',
//     image:
//       'https://images.unsplash.com/photo-1551288049-bebda4e38f71?q=80&w=1200&auto=format&fit=crop',
//     tags: ['VR技术', '教学创新', '沉浸式体验'],
//     author: '张教授',
//     authorTitle: '虚拟现实研究中心主任',
//     authorAvatar:
//       'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?q=80&w=150&auto=format&fit=crop',
//     contact1: '邮箱zhang@university.edu.cn',
//     contact2: '电话123456789',
//     projectUrl: 'https://github.com',

//     // 视频和图片库
//     videoUrl: 'http://vjs.zencdn.net/v/oceans.mp4',
//     gallery: [
//       'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//       'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//       'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//       'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//     ],
//   },
//   {
//     id: 1,
//     title: '虚拟现实教学平台',
//     description:
//       '基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。',
//     image:
//       'https://images.unsplash.com/photo-1551288049-bebda4e38f71?q=80&w=1200&auto=format&fit=crop',
//     tags: ['VR技术', '教学创新', '沉浸式体验'],
//     author: '张教授',
//     authorTitle: '虚拟现实研究中心主任',
//     authorAvatar:
//       'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?q=80&w=150&auto=format&fit=crop',
//     contact1: '邮箱zhang@university.edu.cn',
//     contact2: '电话123456789',
//     projectUrl: 'https://github.com',

//     // 视频和图片库
//     videoUrl: 'http://vjs.zencdn.net/v/oceans.mp4',
//     gallery: [
//       'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//       'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//       'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//       'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//     ],
//   },
// ])
// ...existing code...
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
