<!-- InfoTabs.vue -->
<template>
  <div class="tabs-container">
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
        @click="$emit('update:activeTab', tab.id)"
        :ref="(el) => setTabRef(el, tab.id)"
      >
        {{ tab.name }}
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, onMounted, nextTick } from 'vue'

const props = defineProps({
  tabs: Array,
  activeTab: String,
})
defineEmits(['update:activeTab'])

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
  () => {
    updateSlider()
  },
)

// 5. 在组件挂载后，也需要立即设置一次滑块的初始位置
onMounted(async () => {
  // 使用 nextTick 确保 v-for 渲染完成，tabRefs 已经有值
  await nextTick()
  updateSlider()

  // 立即设置滑块位置后，再添加过渡效果，避免初始加载时有动画
  await nextTick()
  const tabsEl = document.querySelector('.tabs')
  if (tabsEl) {
    tabsEl.style.setProperty('--slider-transition', 'all 0.4s cubic-bezier(0.4, 0, 0.2, 1)')
  }
})
</script>

<style scoped>
.tabs-container {
  display: flex;
  justify-content: center;
  margin-bottom: 50px;
}

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
</style>
