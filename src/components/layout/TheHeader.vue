<template>
  <!--
    悬停触发器保持不变，它的逻辑是正确的。
    占位符已移除，因为导航栏现在总是浮动的。
  -->
  <div
    v-show="isScrolled"
    class="top-hover-trigger"
    @mouseenter="expandHeader"
    @mouseleave="collapseHeader"
  ></div>

  <header
    ref="headerRef"
    class="header-container"
    :class="{ scrolled: isScrolled, expanded: isExpanded && isScrolled }"
    @mouseenter="expandHeader"
    @mouseleave="collapseHeader"
  >
    <div class="content-wrapper">
      <div class="logo">
        <router-link to="/">虚拟仿真中心</router-link>
      </div>

      <nav class="navigation-links">
        <router-link to="/">首页</router-link>
        <router-link to="/information_center">信息中心</router-link>
        <router-link to="/showcase">成果展示</router-link>
        <router-link to="/our_team">团队介绍</router-link>
        <router-link to="/forum">交流中心</router-link>
        <router-link to="/about_us">关于我们</router-link>
        <router-link to="/login_and_register">登录/注册</router-link>
      </nav>
    </div>
  </header>
</template>

<script setup>
// [改动] 脚本逻辑微调以适应新效果
import { ref, onMounted, onUnmounted } from 'vue'

const headerRef = ref(null)
const isScrolled = ref(false)
const isExpanded = ref(false)

// [改动] 不再需要 navbarHeight
// const navbarHeight = ref(0)

const handleScroll = () => {
  // 当滚动一小段距离后（比如 > 10px），就判定为 "scrolled" 状态
  if (window.scrollY > 10) {
    isScrolled.value = true
  } else {
    isScrolled.value = false
    isExpanded.value = false
  }
}
const expandHeader = () => {
  if (isScrolled.value) isExpanded.value = true
}
const collapseHeader = () => {
  isExpanded.value = false
}

// [改动] 移除了 onMounted 中获取高度的逻辑
onMounted(() => {
  window.addEventListener('scroll', handleScroll)
})
onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
})
</script>

<style scoped>
/* 悬停触发区 (不变) */
.top-hover-trigger {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 72px;
  z-index: 1001;
}

/*
 * =======================================================
 *               ↓↓↓ 核心修改区域 ↓↓↓
 * =======================================================
 */

/* 1. 修改 header 容器的初始状态，并整合悬停延迟逻辑 */
.header-container {
  /* [改动] 初始定位改为 absolute，使其浮动 */
  position: absolute;
  top: 0;
  left: 0;

  /* [改动] 初始背景色为透明，初始边框也为透明 */
  background-color: transparent;
  border-bottom: 1px solid transparent;

  /* [整合] 为多个属性添加过渡动画，并为 transform 设置 0s 的隐藏延迟 */
  transition:
    transform 0.4s ease-in-out 0s,
    background-color 0.3s ease-in-out,
    border-color 0.3s ease-in-out;

  /* 以下是您原有的布局样式，保持不变 */
  height: 72px;
  width: 100%;
  z-index: 1000;
}

/* 2. 定义滚动后的状态 */
.header-container.scrolled {
  /* [改动] 定位改为 fixed，使其固定在顶部 */
  position: fixed;

  /* [改动] 背景色变为白色，边框颜色恢复 */
  background-color: white;
  border-bottom-color: #e5e7eb;

  /* 以下是您之前就有的滚动/展开逻辑，保持不变 */
  box-shadow:
    0 4px 6px -1px rgb(0 0 0 / 0.1),
    0 2px 4px -2px rgb(0 0 0 / 0.1);
  transform: translateY(-100%);
}

/* 3. [整合] 定义展开状态，并在此处应用“显示延迟” */
.header-container.scrolled.expanded {
  transform: translateY(0);

  /* [整合] 只对 transform 应用 0.2s 的显示延迟，背景色等变化不受影响 */
  transition-delay: 0.2s, 0s, 0s;
}

/* 4. [新增] 根据状态动态改变文字和元素的颜色 */

/* 初始状态 (透明背景时)，文字为白色 */
.header-container:not(.scrolled) .logo a,
.header-container:not(.scrolled) .navigation-links a {
  color: white;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.5); /* 添加文字阴影以增强可读性 */
}

/* 初始状态下的搜索框 */
.search-bar .initial-state-input {
  background-color: rgba(255, 255, 255, 0.2);
  border-color: rgba(255, 255, 255, 0.5);
  color: white;
}
.search-bar .initial-state-input::placeholder {
  color: rgba(255, 255, 255, 0.8);
}

/*
 * =======================================================
 *           ↓↓↓ 您原有的样式，完全保留 ↓↓↓
 * =======================================================
 */
.content-wrapper {
  max-width: 1280px;
  margin: 0 auto;
  padding: 0 24px;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.logo a {
  font-size: 20px;
  font-weight: bold;
  color: #1f2937;
  text-decoration: none;
}
.navigation-links {
  display: flex;
  gap: 24px;
}
.navigation-links a {
  font-size: 16px;
  color: #4b5563;
  text-decoration: none;
  transition: color 0.3s;
}
.navigation-links a:hover {
  color: #4f46e5;
}
.navigation-links a.router-link-active {
  color: #4f46e5;
  font-weight: 600;
}
</style>
