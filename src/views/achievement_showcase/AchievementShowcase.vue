<template>
  <div class="showcase-page">
    <AchievementHeader
      title="卓越成果展示"
      subtitle="探索我们如何通过前沿技术，将创新理念转化为具有影响力的教学与科研成果。"
    />
    <main class="content-container">
      <AchievementActions
        v-model="searchQuery"
        v-model:teacherFilter="teacherFilter"
        v-model:tagFilter="tagFilter"
        v-model:timeFilter="timeFilter"
        :teacher-options="teacherOptions"
        :tag-options="tagOptions"
        @open-submission="isSubmissionModalVisible = true"
      />
      <AchievementGrid :achievements="paginatedAchievements" @view-details="openDetailsModal" />
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
    <AchievementPagination
      v-if="totalPages > 1"
      :current-page="currentPage"
      :total-pages="totalPages"
      @page-changed="goToPage"
    />
  </div>
</template>

<script setup>
import { ref, computed, nextTick } from 'vue'
import AchievementHeader from '@/components/achievement_showcase/AchievementHeader.vue'
import AchievementGrid from '@/components/achievement_showcase/AchievementGrid.vue'
import AchievementDetailModal from '@/components/achievement_showcase/AchievementDetailModal.vue'
import AchievementActions from '@/components/achievement_showcase/AchievementActions.vue'
import SubmissionModal from '@/components/achievement_showcase/SubmissionModal.vue'
import AchievementPagination from '@/components/achievement_showcase/AchievementPagination.vue'

// --- State ---
const achievementTemplate = {
  id: 1, // 这个 id 会在循环中被新的、唯一的 id 覆盖
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
  publishDate: new Date(),
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
}

// 2. 准备一个空数组来存放生成的数据
const generatedData = []

// 3. 循环100次，批量生成数据
for (let i = 0; i < 1000; i++) {
  generatedData.push({
    ...achievementTemplate, // 使用展开运算符(...)来复制模板对象的所有属性
    id: i + 1, // 关键：为每个新对象设置一个唯一的id，这对于测试和渲染都非常重要
  })
}

const achievements = ref(generatedData)

const selectedAchievement = ref(null)
const isModalVisible = ref(false)
const isSubmissionModalVisible = ref(false)
const searchQuery = ref('')
const teacherFilter = ref('')
const tagFilter = ref('')
const timeFilter = ref('')
const currentPage = ref(1)
const itemsPerPage = 15

// 动态生成筛选选项
const teacherOptions = computed(() => {
  const teachers = new Set(achievements.value.map((a) => a.author))
  const options = Array.from(teachers).map((teacher) => ({ label: teacher, value: teacher }))
  return [{ label: '所有教师', value: '' }, ...options]
})

const tagOptions = computed(() => {
  const tags = new Set(achievements.value.flatMap((a) => a.tags))
  const options = Array.from(tags).map((tag) => ({ label: tag, value: tag }))
  return [{ label: '所有标签', value: '' }, ...options]
})

// 搜索筛选函数
const filteredAchievements = computed(() => {
  let filtered = achievements.value

  // 教师筛选
  if (teacherFilter.value) {
    filtered = filtered.filter((achievement) => achievement.author === teacherFilter.value)
  }

  // 标签筛选
  if (tagFilter.value) {
    filtered = filtered.filter((achievement) => achievement.tags.includes(tagFilter.value))
  }

  // 时间筛选
  if (timeFilter.value) {
    const now = new Date()
    const filterDate = new Date()

    switch (timeFilter.value) {
      case 'week':
        filterDate.setDate(now.getDate() - 7)
        break
      case 'month':
        filterDate.setMonth(now.getMonth() - 1)
        break
      case 'three_months':
        filterDate.setMonth(now.getMonth() - 3)
        break
      case 'half_year':
        filterDate.setMonth(now.getMonth() - 6)
        break
      case 'year':
        filterDate.setFullYear(now.getFullYear() - 1)
        break
    }

    filtered = filtered.filter((achievement) => new Date(achievement.publishDate) >= filterDate)
  }

  // 搜索查询
  if (searchQuery.value) {
    const lowerCaseQuery = searchQuery.value.toLowerCase()
    filtered = filtered.filter((achievement) => {
      const titleMatch = achievement.title.toLowerCase().includes(lowerCaseQuery)
      const tagMatch = achievement.tags.some((tag) => tag.toLowerCase().includes(lowerCaseQuery))
      const authorMatch = achievement.author.toLowerCase().includes(lowerCaseQuery)
      return titleMatch || tagMatch || authorMatch
    })
  }

  return filtered
})

// 分页计算属性
const paginatedAchievements = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage
  const end = start + itemsPerPage
  return filteredAchievements.value.slice(start, end)
})

const totalPages = computed(() => {
  return Math.ceil(filteredAchievements.value.length / itemsPerPage)
})

// --- Methods ---
const scrollToTop = () => {
  nextTick(() => {
    window.scrollTo({ top: 300, behavior: 'smooth' })
  })
}
const goToPage = (page) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
    scrollToTop()
  }
}
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
  padding: 0 40px 40px;
}
</style>
