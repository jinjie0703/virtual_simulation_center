<!-- InfoCenter.vue -->
<template>
  <InfoHeader title="信息中心" subtitle="洞悉前沿动态，把握机遇脉搏" />
  <div class="main-content">
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
import { allMockData } from './TestData.js'

const activeTab = ref('news')
const currentPage = ref(1)
const itemsPerPage = 10 // 2列 * 5行

const tabs = [
  { id: 'news', name: '学院信息' },
  { id: 'competitions', name: '竞赛信息' },
  { id: 'projects', name: '项目悬赏' },
]

// --- 从外部文件导入数据 ---
const allData = ref(allMockData)

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
.main-content {
  max-width: 1400px;
  margin: 0 auto;
  padding: 0 20px 40px 20px;
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
