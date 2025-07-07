<template>
  <div>
    <!-- 1. 轮播组件，通过 :slides 传入数据 -->
    <Carousel :slides="carouselData" />

    <main>
      <!-- 2. 成果展示组件，通过 :features 传入数据 -->
      <FeatureSection :features="featureData" />

      <!-- 3. 新闻与竞赛组件，传入对应的数据 -->
      <UpdatesSection :news-items="newsData" :competition-items="competitionData" />
    </main>
  </div>
</template>

<script setup>
import { ref } from 'vue'

// 导入子组件
import Carousel from '@/components/home_page/HomeCarousel.vue'
import FeatureSection from '@/components/home_page/FeatureSection.vue'
import UpdatesSection from '@/components/home_page/UpdatesSection.vue'

// 1. 轮播图的数据
const carouselData = ref([
  {
    title: '标题1',
    imageUrl:
      'https://images.unsplash.com/photo-1542370285-b8eb8317691c?q=80&w=1974&auto=format&fit=crop',
  },
  {
    title: '标题2',
    imageUrl:
      'https://images.unsplash.com/photo-1519010470956-6d877008eaa4?q=80&w=2070&auto=format&fit=crop',
  },
  {
    title: '标题3',
    imageUrl:
      'https://tse4.mm.bing.net/th/id/OIP.-tDxOx7ZR2xR63hBP47YuAHaEw?r=0&rs=1&pid=ImgDetMain&o=7&rm=3',
  },
])

// 2. 成果展示的数据
const featureData = ref([
  {
    title: '项目1',
    description: '介绍1',
    imageUrl:
      'https://images.unsplash.com/photo-1506748686214-e9df14d4d9d0?q=80&w=800&auto=format&fit=crop',
  },
  {
    title: '项目2',
    description: '介绍2',
    imageUrl:
      'https://images.unsplash.com/photo-1506748686214-e9df14d4d9d0?q=80&w=800&auto=format&fit=crop',
  },
  {
    title: '项目3',
    description: '介绍3',
    imageUrl:
      'https://images.unsplash.com/photo-1506748686214-e9df14d4d9d0?q=80&w=800&auto=format&fit=crop',
  },
  {
    title: '项目4',
    description: '介绍4。',
    imageUrl:
      'https://images.unsplash.com/photo-1506748686214-e9df14d4d9d0?q=80&w=800&auto=format&fit=crop',
  },
])

// 3. 新闻与竞赛的数据
const newsData = ref([
  { title: '新闻1', date: '2024-05-20', url: '#' },
  { title: '新闻2', date: '2024-05-15', url: '#' },
  { title: '新闻3', date: '2024-05-10', url: '#' },
  { title: '新闻4', date: '2024-05-05', url: '#' },
  { title: '新闻5', date: '2024-04-28', url: '#' },
])

const competitionData = ref([
  { title: '竞赛1', date: '2024-05-18', url: '#' },
  { title: '竞赛2', date: '2024-05-12', url: '#' },
  { title: '竞赛3', date: '2024-05-09', url: '#' },
  { title: '竞赛4', date: '2024-05-02', url: '#' },
  { title: '竞赛5', date: '2024-04-25', url: '#' },
])
</script>

<style>
/* 可以在这里放一些全局样式，或者在 main.css 中定义 */
body {
  margin: 0;
  font-family:
    -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  background-color: #f4f4f4;
  color: #333;
}
main {
  display: block; /* 确保 main 标签的兼容性 */
}
</style>
