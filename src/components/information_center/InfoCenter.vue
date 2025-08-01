<!-- InfoCenter.vue -->
<template>
  <InfoHeader title="信息中心" subtitle="洞悉前沿动态，把握机遇脉搏" />
  <div class="main-content" ref="mainContentRef">
    <InfoTabs
      :tabs="tabs"
      :active-tab="activeTab"
      :search-query="searchKeyword"
      @update:activeTab="changeTab"
      @search="onSearch"
      @time-filter="onTimeFilter"
    />

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
import { ref, computed, nextTick } from 'vue'
import InfoHeader from './InfoHeader.vue'
import InfoTabs from './InfoTabs.vue'
import InfoGrid from './InfoGrid.vue'
import InfoPagination from './InfoPagination.vue'
import { allMockData } from './TestData.js'

const activeTab = ref('news')
const currentPage = ref(1)
const itemsPerPage = 10

const mainContentRef = ref(null)

const tabs = [
  { id: 'news', name: '学院信息' },
  { id: 'competitions', name: '竞赛信息' },
  { id: 'projects', name: '项目悬赏' },
]

const allData = ref(allMockData)

const scrollToTop = () => {
  nextTick(() => {
    mainContentRef.value?.scrollIntoView({ behavior: 'smooth', block: 'start' })
  })
}

const changeTab = (tabId) => {
  activeTab.value = tabId
  currentPage.value = 1
  searchKeyword.value = ''
  timeFilter.value = 'all'
  scrollToTop()
}

const goToPage = (page) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
  }
  scrollToTop()
}

// --- 搜索和筛选的状态 ---
const searchKeyword = ref('')
const timeFilter = ref('all') // 'all', 'oneMonth', 'threeMonths', 'sixMonths', 'oneYear'
const sortOrder = ref('desc') // 'desc' = 降序 (最新), 'asc' = 升序 (最旧)

// --- 事件处理函数 ---
const onSearch = (keyword) => {
  searchKeyword.value = keyword
  currentPage.value = 1
}

const onTimeFilter = (filterValue) => {
  timeFilter.value = filterValue
  currentPage.value = 1
}

// --- 核心逻辑：使用计算属性动态生成列表 ---
const filteredAndSortedItems = computed(() => {
  let items = allData.value[activeTab.value] || []

  // 1. 标准化数据，统一日期字段为 effectiveDate，并修正搜索字段
  let normalizedItems = items.map((item) => ({
    ...item,
    effectiveDate: new Date(item.date || item.deadline),
  }))

  // 2. 搜索筛选
  if (searchKeyword.value) {
    const lowerCaseKeyword = searchKeyword.value.toLowerCase()
    normalizedItems = normalizedItems.filter(
      (item) =>
        item.title.toLowerCase().includes(lowerCaseKeyword) ||
        (item.summary && item.summary.toLowerCase().includes(lowerCaseKeyword)),
    )
  }

  // 3. 时间筛选
  if (timeFilter.value !== 'all') {
    const now = new Date()
    const cutoffDate = new Date()
    let monthsToSubtract = 0

    if (timeFilter.value === 'oneMonth') monthsToSubtract = 1
    if (timeFilter.value === 'threeMonths') monthsToSubtract = 3
    if (timeFilter.value === 'sixMonths') monthsToSubtract = 6
    if (timeFilter.value === 'oneYear') monthsToSubtract = 12

    if (monthsToSubtract > 0) {
      cutoffDate.setMonth(now.getMonth() - monthsToSubtract)
      // 筛选时，确保日期在截止日期和当前时间之间
      normalizedItems = normalizedItems.filter(
        (item) => item.effectiveDate >= cutoffDate && item.effectiveDate <= now,
      )
    }
  }

  // 4. 时间排序
  normalizedItems.sort((a, b) => {
    return sortOrder.value === 'desc'
      ? b.effectiveDate - a.effectiveDate
      : a.effectiveDate - b.effectiveDate
  })

  return normalizedItems
})

const totalPages = computed(() => Math.ceil(filteredAndSortedItems.value.length / itemsPerPage))

const paginatedItems = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage
  const end = start + itemsPerPage
  return filteredAndSortedItems.value.slice(start, end)
})
</script>

<style scoped>
.main-content {
  max-width: 1400px;
  margin: 0 auto;
  padding: 0 20px 40px 20px;
  scroll-margin-top: 80px;
}

.content-area {
  min-height: 600px;
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
