<!-- App.vue (最终正确版) -->
<template>
  <div id="app">
    <TheHeader />
    <!--
      这里不再有 <main> 标签包裹。
      RouterView 直接渲染，让每个页面组件完全控制自己的布局。
    -->
    <RouterView />
  </div>
</template>

<script setup lang="js">
import { RouterView } from 'vue-router'
import TheHeader from './components/layout/TheHeader.vue'
</script>

<!--
  全局样式 (无 "scoped")
  我们在这里定义基础样式和一个可复用的布局容器类。
-->
<style scoped>
/* 引入字体 */
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700;800&display=swap');

/* CSS Reset 和 基础样式 */
body {
  margin: 0;
}

#app {
  font-family: 'Inter', sans-serif;
  color: #111827;
  background-color: #f9fafb;
  min-height: 100vh;
  /* 我们把 #app 设置为 flex 布局，让内容区可以自然撑开 */
  display: flex;
  flex-direction: column;
}

/*
  这是一个新的、可复用的全局布局容器类。
  任何需要居中、有内边距的页面都可以使用它。
*/
.page-container {
  max-width: 1280px;
  width: 100%; /* 确保在小屏幕上占满宽度 */
  margin-left: auto;
  margin-right: auto;
  padding: 32px 16px; /* 上下32px，左右16px */
  box-sizing: border-box; /* 确保 padding 不会增加总宽度 */
}

/* 响应式媒体查询 */
@media (min-width: 640px) {
  .page-container {
    padding-left: 24px;
    padding-right: 24px;
  }
}

@media (min-width: 1024px) {
  .page-container {
    padding-left: 32px;
    padding-right: 32px;
  }
}
</style>
