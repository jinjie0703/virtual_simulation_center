<template>
  <div class="showcase-page">
    <AchievementHeader
      title="卓越成果展示"
      subtitle="探索我们如何通过前沿技术，将创新理念转化为具有影响力的教学与科研成果。"
    />
    <main class="content-container">
      <AchievementActions
        v-model="searchQuery"
        @open-submission="isSubmissionModalVisible = true"
      />
      <AchievementGrid :achievements="filteredAchievements" @view-details="openDetailsModal" />
    </main>
    <AchievementDetailModal
      :achievement="selectedAchievement"
      :is-visible="isModalVisible"
      @close="closeDetailsModal"
    />
    <SubmissionModal
      :is-visible="isSubmissionModalVisible"
      @close="isSubmissionModalVisible = false"
      @submit="handleSubmission"
    />
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import AchievementHeader from '@/components/achievement_showcase/AchievementHeader.vue'
import AchievementGrid from '@/components/achievement_showcase/AchievementGrid.vue'
import AchievementDetailModal from '@/components/achievement_showcase/AchievementDetailModal.vue'
import AchievementActions from '@/components/achievement_showcase/AchievementActions.vue'
import SubmissionModal from '@/components/achievement_showcase/SubmissionModal.vue'

// --- State ---
const achievements = ref([
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

const selectedAchievement = ref(null)
const isModalVisible = ref(false)
const isSubmissionModalVisible = ref(false)
const searchQuery = ref('')

// --- Computed ---
const filteredAchievements = computed(() => {
  if (!searchQuery.value) {
    return achievements.value
  }
  const lowerCaseQuery = searchQuery.value.toLowerCase()
  return achievements.value.filter((achievement) => {
    const titleMatch = achievement.title.toLowerCase().includes(lowerCaseQuery)
    const tagMatch = achievement.tags.some((tag) => tag.toLowerCase().includes(lowerCaseQuery))
    return titleMatch || tagMatch
  })
})

// --- Methods ---
const openDetailsModal = (achievement) => {
  selectedAchievement.value = achievement
  isModalVisible.value = true
}

const closeDetailsModal = () => {
  isModalVisible.value = false
  setTimeout(() => {
    selectedAchievement.value = null
  }, 300)
}

const handleSubmission = (newAchievement) => {
  achievements.value.unshift(newAchievement)
  // In a real app, you'd likely send this to a server.
  // You might also want to show a success notification.
}
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
