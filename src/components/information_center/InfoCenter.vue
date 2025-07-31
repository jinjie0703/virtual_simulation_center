<!-- InfoCenter.vue -->
<template>
  <InfoHeader title="信息中心" subtitle="洞悉前沿动态，把握机遇脉搏" />
  <div class="main-content" ref="mainContentRef">
    <InfoTabs
      :tabs="tabs"
      :active-tab="activeTab"
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
const sortOrder = ref('desc') // 'desc' = 降序 (最新), 'asc' = 升序 (最旧)

// --- 事件处理函数 ---
// 当 InfoTabs 的搜索框输入时，这个函数会被调用
const onSearch = (keyword) => {
  searchKeyword.value = keyword
}

// 当 InfoTabs 的时间筛选按钮点击时，这个函数会被调用
const onTimeFilter = () => {
  // 切换排序顺序
  sortOrder.value = sortOrder.value === 'desc' ? 'asc' : 'desc'
}

// --- 核心逻辑：使用计算属性动态生成列表 ---
const filteredAndSortedItems = computed(() => {
  // a. 先根据 activeTab 筛选
  let items = allData.value[activeTab.value] || []

  // b. 再根据搜索关键词筛选
  if (searchKeyword.value) {
    const lowerCaseKeyword = searchKeyword.value.toLowerCase()
    items = items.filter(
      (item) =>
        item.title.toLowerCase().includes(lowerCaseKeyword) ||
        item.overview.toLowerCase().includes(lowerCaseKeyword),
    )
  }

  // c. 最后根据时间排序
  items.sort((a, b) => {
    const dateA = new Date(a.date)
    const dateB = new Date(b.date)
    return sortOrder.value === 'desc' ? dateB - dateA : dateA - dateB
  })

  return items
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
