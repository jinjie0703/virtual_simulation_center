<!-- InfoTabs.vue -->
<template>
  <div class="tabs-container">
    <!-- 搜索框 -->
    <div class="search-box">
      <input
        :value="searchQuery"
        @input="emit('search', $event.target.value)"
        type="text"
        placeholder="按标题和概述搜索..."
        class="search-input"
      />
      <svg class="search-icon" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path
          d="M21 21L15 15M17 10C17 13.866 13.866 17 10 17C6.13401 17 3 13.866 3 10C3 6.13401 6.13401 3 10 3C13.866 3 17 6.13401 17 10Z"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
        />
      </svg>
    </div>

    <div
      class="tabs"
      ref="tabsRef"
      :style="{
        '--slider-left': sliderLeft,
        '--slider-width': sliderWidth,
      }"
    >
      <button
        v-for="tab in tabs"
        :key="tab.id"
        :class="['tab-btn', { active: activeTab === tab.id }]"
        @click="emit('update:activeTab', tab.id)"
        :ref="(el) => setTabRef(el, tab.id)"
      >
        {{ tab.name }}
      </button>
    </div>

    <!-- 时间筛选 -->
    <div class="filter-box" ref="filterBoxRef">
      <button @click="toggleDropdown" class="filter-btn">
        <svg class="filter-icon" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M3 4H21V6H3V4ZM3 11H15V13H3V11ZM3 18H9V20H3V18Z" fill="currentColor" />
        </svg>
        {{ selectedFilterLabel }}
      </button>
      <transition name="dropdown">
        <div v-if="isDropdownOpen" class="dropdown-menu">
          <button
            v-for="filter in timeFilters"
            :key="filter.value"
            @click="selectTimeFilter(filter)"
            class="dropdown-item"
          >
            {{ filter.text }}
          </button>
        </div>
      </transition>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, onMounted, nextTick, onBeforeUnmount } from 'vue'

const props = defineProps({
  tabs: Array,
  activeTab: String,
  searchQuery: String,
})
const emit = defineEmits(['update:activeTab', 'search', 'timeFilter'])

// 时间筛选
const filterBoxRef = ref(null)
const isDropdownOpen = ref(false)
const timeFilters = ref([
  { value: 'all', text: '全部时间' },
  { value: 'oneMonth', text: '最近一个月' },
  { value: 'threeMonths', text: '最近三个月' },
  { value: 'sixMonths', text: '最近半年' },
  { value: 'oneYear', text: '最近一年' },
])
const selectedFilterLabel = ref('时间筛选')

const toggleDropdown = () => {
  isDropdownOpen.value = !isDropdownOpen.value
}

const selectTimeFilter = (filter) => {
  selectedFilterLabel.value = filter.text
  emit('timeFilter', filter.value)
  isDropdownOpen.value = false
}

const handleClickOutside = (event) => {
  if (filterBoxRef.value && !filterBoxRef.value.contains(event.target)) {
    isDropdownOpen.value = false
  }
}

// 1. 创建 Ref 来存储滑块的位置和宽度
const sliderLeft = ref('0px')
const sliderWidth = ref('0px')

// 2. 创建一个对象来存储每个 tab 按钮的 DOM 引用
const tabRefs = ref({})
const setTabRef = (el, tabId) => {
  if (el) {
    tabRefs.value[tabId] = el
  }
}

// 3. 定义一个函数来更新滑块的位置
const updateSlider = () => {
  const activeEl = tabRefs.value[props.activeTab]
  if (!activeEl) return

  // activeEl.offsetLeft: 元素左边框相对于父元素左内边框的距离
  // activeEl.offsetWidth: 元素的总宽度（包括 padding 和 border）
  sliderLeft.value = `${activeEl.offsetLeft}px`
  sliderWidth.value = `${activeEl.offsetWidth}px`
}

// 4. 监听 activeTab 的变化，当它变化时更新滑块
watch(
  () => props.activeTab,
  (newTab, oldTab) => {
    updateSlider()
    // 当 tab 切换时，重置筛选
    if (newTab !== oldTab) {
      selectedFilterLabel.value = '时间筛选'
    }
  },
)

onMounted(async () => {
  await nextTick()
  updateSlider()
  await nextTick()
  const tabsEl = document.querySelector('.tabs')
  if (tabsEl) {
    tabsEl.style.setProperty('--slider-transition', 'all 0.4s cubic-bezier(0.4, 0, 0.2, 1)')
  }
  document.addEventListener('click', handleClickOutside)
})

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.tabs-container {
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 50px;
  width: 100%;
  min-height: 60px;
}

/* 搜索框样式 - 绝对左对齐 */
.search-box {
  position: absolute;
  left: 0;
  display: flex;
  align-items: center;
}

.search-input {
  width: 250px;
  padding: 12px 16px 12px 40px;
  border: 2px solid #e2e8f0;
  border-radius: 999px;
  font-size: 14px;
  background-color: white;
  transition: border-color 0.3s ease;
}

.search-input:focus {
  outline: none;
  border-color: #8fd3f4;
}

.search-icon {
  position: absolute;
  left: 12px;
  width: 20px;
  height: 20px;
  color: #a0aec0;
}

/* 筛选按钮样式 - 绝对右对齐 */
.filter-box {
  position: absolute;
  right: 0;
  display: flex;
  align-items: center;
  z-index: 10; /* 确保下拉菜单在最上层 */
}

.filter-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 20px;
  border: 2px solid #e2e8f0;
  border-radius: 999px;
  background-color: white;
  font-size: 14px;
  font-weight: 600;
  color: #2d3748;
  cursor: pointer;
  transition: all 0.3s ease;
}

.filter-btn:hover,
.filter-btn.active {
  border-color: #8fd3f4;
  color: #8fd3f4;
}

.filter-icon {
  width: 16px;
  height: 16px;
}

/* tabs保持居中 */
.tabs {
  position: relative; /* 成为伪元素定位的基准 */
  display: inline-flex;
  background-color: white;
  border-radius: 999px;
  padding: 6px;
}

/* 滑块 */
.tabs::before {
  content: '';
  position: absolute;
  top: 6px;
  left: var(--slider-left, 0px);
  width: var(--slider-width, 0px);
  height: calc(100% - 12px);
  border-radius: 999px;
  background-color: white;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  transition: var(--slider-transition, none);
}

.tab-btn {
  padding: 12px 30px;
  border: none;
  background-color: transparent;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 600;
  color: #2d3748;
  transition: color 1.5s ease;
  position: relative;
  z-index: 1;
}

.tab-btn.active {
  color: #8fd3f4;
}

/* 下拉菜单样式 */
.dropdown-menu {
  position: absolute;
  top: calc(100% + 8px);
  right: 0;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  padding: 8px;
  display: flex;
  flex-direction: column;
  gap: 4px;
  width: 150px;
}

.dropdown-item {
  width: 100%;
  padding: 10px 12px;
  border: none;
  background-color: transparent;
  text-align: left;
  font-size: 14px;
  border-radius: 6px;
  cursor: pointer;
  transition:
    background-color 0.2s ease,
    color 0.2s ease;
}

.dropdown-item:hover {
  background-color: #f0f2f5;
  color: #8fd3f4;
}

/* 下拉菜单过渡动画 */
.dropdown-enter-active,
.dropdown-leave-active {
  transition: all 0.2s ease-out;
}

.dropdown-enter-from,
.dropdown-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}
</style>
