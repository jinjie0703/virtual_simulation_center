<template>
  <div class="info-center-container">
    <InfoHeader title="信息中心" subtitle="洞悉前沿动态，把握机遇脉搏" />

    <InfoTabs :tabs="tabs" :active-tab="activeTab" @update:activeTab="changeTab" />

    <div class="content-area">
      <transition name="fade" mode="out-in">
        <InfoGrid :key="activeTab" :items="paginatedItems" :item-type="activeTab" />
      </transition>
    </div>

    <InfoPagination
      v-if="totalPages > 1"
      :current-page="currentPage"
      :total-pages="totalPages"
      @page-changed="goToPage"
    />
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import InfoHeader from './InfoHeader.vue'
import InfoTabs from './InfoTabs.vue'
import InfoGrid from './InfoGrid.vue'
import InfoPagination from './InfoPagination.vue'

const activeTab = ref('news')
const currentPage = ref(1)
const itemsPerPage = 10 // 2列 * 5行

const tabs = [
  { id: 'news', name: '学院信息' },
  { id: 'competitions', name: '竞赛信息' },
  { id: 'projects', name: '项目悬赏' },
]

// --- 模拟数据 ---
const allData = ref({
  news: Array.from({ length: 25 }, (_, i) => ({
    id: `news-${i + 1}`,
    title: `我校虚拟仿真中心获得第${i + 1}项国家级教学成果奖`,
    summary: '中心凭借其在沉浸式教学领域的创新实践，再次荣获国家级表彰，引领了教育技术的新潮流。',
    imageUrl: `https://picsum.photos/seed/news${i}/400/250`,
    date: `2024-05-${25 - i > 0 ? 25 - i : 1}`,
    category: '学术荣誉',
  })),
  competitions: Array.from({ length: 18 }, (_, i) => ({
    id: `comp-${i + 1}`,
    title: `第${i + 1}届全国大学生虚拟现实设计大赛`,
    summary: '本次大赛旨在激发学生的创新精神和实践能力，推动VR技术在各行业的应用与发展。',
    imageUrl: `https://picsum.photos/seed/comp${i}/400/250`,
    deadline: '2024-07-30',
    level: '国家级',
  })),
  projects: Array.from({ length: 12 }, (_, i) => ({
    id: `proj-${i + 1}`,
    title: `开发一个基于AI的个性化学习路径推荐系统 #${i + 1}`,
    summary: '需要利用机器学习算法，根据学生的学习行为和成绩，动态生成最优学习路径。',
    imageUrl: `https://picsum.photos/seed/proj${i}/400/250`,
    deadline: '2024-08-15',
    reward: `${5 + i}k`,
    tags: ['AI', 'Vue.js', 'Python'],
  })),
})

const currentItems = computed(() => allData.value[activeTab.value] || [])

const totalPages = computed(() => Math.ceil(currentItems.value.length / itemsPerPage))

const paginatedItems = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage
  const end = start + itemsPerPage
  return currentItems.value.slice(start, end)
})

const changeTab = (tabId) => {
  activeTab.value = tabId
  currentPage.value = 1
}

const goToPage = (page) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
  }
}
</script>

<style scoped>
.info-center-container {
  max-width: 1400px;
  margin: 0 auto;
  padding: 40px 20px;
}

.content-area {
  min-height: 600px; /* 防止切换时页面抖动 */
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
