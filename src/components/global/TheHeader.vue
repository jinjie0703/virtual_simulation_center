<template>
  <!--悬停触发器（顶部隐形）-->
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
        <router-link to="/home_page">首页</router-link>
        <router-link to="/information_center">信息中心</router-link>
        <router-link to="/achievement_showcase">成果展示</router-link>
        <router-link to="/team_center">组队中心</router-link>
        <router-link to="/our_team">师资队伍</router-link>
        <router-link to="/about_us">关于我们</router-link>
      </nav>
    </div>
  </header>
</template>

<script setup>
// css脚本
import { ref, onMounted, onUnmounted } from 'vue'

const headerRef = ref(null)
const isScrolled = ref(false)
const isExpanded = ref(false)

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
// 悬停触发关键：当鼠标悬停在 header 上时，触发展开
const collapseHeader = () => {
  isExpanded.value = false
}

// 监听滚动事件
onMounted(() => {
  window.addEventListener('scroll', handleScroll)
})
// 移除滚动事件监听器
onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
})
</script>

<style scoped>
/* 悬停触发区*/
.top-hover-trigger {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 72px;
  z-index: 1001;
}

/* 1. 修改 header 容器的初始状态，并整合悬停延迟逻辑 */
.header-container {
  /* 初始定位改为 absolute，使其浮动 */
  position: absolute;
  top: 0;
  left: 0;

  /* 初始背景色为透明，初始边框也为透明 */
  background-color: transparent;
  border-bottom: 1px solid transparent;

  /* 为多个属性添加过渡动画，并为 transform 设置 0s 的隐藏延迟 */
  transition:
    transform 0.3s ease-in-out 0s,
    background-color 0.5s ease-in-out,
    border-color 0.5s ease-in-out;

  height: 72px;
  width: 100%;
  z-index: 1000;
}

/* 2. 定义滚动后的状态 */
.header-container.scrolled {
  /* 定位改为 fixed，使其固定在顶部 */
  position: fixed;

  /* 背景色变为白色，边框颜色恢复 */
  background-color: white;
  border-bottom-color: white;

  /* 两个阴影增加真实感 */
  box-shadow:
    0 4px 6px -1px rgb(0 0 0 / 0.1),
    0 2px 4px -2px rgb(0 0 0 / 0.1);
  /* 滚动展开逻辑 */
  transform: translateY(-100%);
}

/* 3. 界面滚动且鼠标悬停时 */
.header-container.scrolled.expanded {
  /* 展开导航栏 */
  transform: translateY(0);

  /* 只对 transform 应用 0.1s 的显示延迟，背景色等变化不受影响 */
  transition-delay: 0.1s, 0s, 0s;
  /* 动态提高导航栏的层级 */
  z-index: 1002;
}

/* 4. [新增] 根据状态动态改变文字和元素的颜色 */

/* 初始状态，文字为白色 */
.header-container:not(.scrolled) .logo a,
.header-container:not(.scrolled) .navigation-links a {
  color: white;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.5); /* 添加文字阴影以增强可读性 */
}

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
  gap: 48px;
}
.navigation-links a {
  position: relative;
  padding-bottom: 8px;
  font-size: 16px;
  color: #4b5563;
  text-decoration: none;
  transition: color 0.3s;
}
.navigation-links a::after {
  content: ''; /* 伪元素必须有 content 属性 */
  position: absolute; /* 绝对定位，相对于 a 标签 */
  bottom: 0; /* 定位到 a 标签的底部 (padding-bottom 的下方) */
  left: 50%; /* [动画关键] 从中间开始 */

  width: 0; /* [动画关键] 初始宽度为 0 */
  height: 2px; /* 下划线的粗细 */
  background-color: white; /* 下划线的颜色 (与悬停/激活颜色一致) */

  /* [动画关键] 定义过渡效果 */
  transition: all 0.3s ease-in-out;
}
.navigation-links a:hover::after,
.navigation-links a.router-link-active::after {
  width: 100%; /* [动画关键] 宽度变为 100% */
  left: 0; /* [动画关键] 位置回到左侧 0 点 */
}

/* .header-container:not(.scrolled) .navigation-links a.router-link-active {
  color: white; 确保激活链接文字也是白色
} */

/* .header-container:not(.scrolled) .navigation-links a.router-link-active::after {
  background-color: white; 确保下划线颜色为白色
} */

.navigation-links a:hover {
  color: #4f46e5;
}
.navigation-links a.router-link-active {
  color: #000000;
  font-weight: 600;
}
</style>
