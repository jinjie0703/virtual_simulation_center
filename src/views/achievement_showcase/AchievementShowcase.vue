<template>
  <div class="showcase-page">
    <AchievementHeader
      title="卓越成果展示"
      subtitle="探索我们如何通过前沿技术，将创新理念转化为具有影响力的教学与科研成果。"
    />
    <main class="content-container">
      <AchievementGrid :achievements="achievements" @view-details="openDetailsModal" />
    </main>
    <AchievementDetailModal
      :achievement="selectedAchievement"
      :is-visible="isModalVisible"
      @close="closeDetailsModal"
    />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import AchievementHeader from '@/components/achievement_showcase/AchievementHeader.vue'
import AchievementGrid from '@/components/achievement_showcase/AchievementGrid.vue'
import AchievementDetailModal from '@/components/achievement_showcase/AchievementDetailModal.vue'

const isModalVisible = ref(false)
const selectedAchievement = ref(null)

const openDetailsModal = (achievement) => {
  selectedAchievement.value = achievement
  isModalVisible.value = true
}

const closeDetailsModal = () => {
  isModalVisible.value = false
  // 延迟清空数据以获得更好的关闭动画效果
  setTimeout(() => {
    selectedAchievement.value = null
  }, 300)
}

// 使用 HomePage.vue 中的 featureData 结构作为模拟数据
const achievements = ref([
  {
    id: 1,
    title: '虚拟现实教学平台',
    description:
      '基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。',
    image:
      'https://images.unsplash.com/photo-1592478411213-6153e4ebc696?q=80&w=1200&auto=format&fit=crop',
    tags: ['VR技术', '教学创新', '沉浸式体验'],
    author: '张教授',
    authorTitle: '虚拟现实研究中心主任',
    authorAvatar:
      'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?q=80&w=150&auto=format&fit=crop',
    contact: 'zhang@university.edu.cn',
    projectUrl: 'https://github.com',
    videoUrl: 'http://vjs.zencdn.net/v/oceans.mp4',
    gallery: [
      'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=800&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1519010470956-6d877008eaa4?q=80&w=800&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1542370285-b8eb8317691c?q=80&w=800&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1592478411213-6153e4ebc696?q=80&w=800&auto=format&fit=crop',
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
    projectUrl: 'https://github.com',
    videoUrl: 'https://media.w3.org/2010/05/sintel/trailer.mp4',
    gallery: [
      'https://images.unsplash.com/photo-1485827404703-89b55fcc595e?q=80&w=800&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1581092160562-40aa08e78837?q=80&w=800&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1518709268805-4e9042af2176?q=80&w=800&auto=format&fit=crop',
      'https://images.unsplash.com/photo-1551288049-bebda4e38f71?q=80&w=800&auto=format&fit=crop',
    ],
  },
  // ...可以添加更多成果数据
])
</script>

<style scoped>
.showcase-page {
  background-color: #f8f9fa;
  min-height: 100vh;
}
.content-container {
  max-width: 1600px;
  margin: 0 auto;
  padding: 0 40px 80px;
}
</style>
