<template>
  <div id="app">
    <TheHeader />
    <RouterView />
    <!-- 页脚渲染逻辑 -->
    <SiteFooter v-if="$route.meta.showFooter !== false" />
  </div>
</template>

<script setup>
// 注意：lang="js" 可以省略，setup 默认就是 JS
import { RouterView } from 'vue-router'
import TheHeader from './components/global/TheHeader.vue'
// ✨ 导入我们之前创建的页脚组件
import SiteFooter from './components/global/SiteFooter.vue'
</script>

<style scoped>
/* 引入字体 */
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700;800&display=swap');

body {
  margin: 0;
  background-color: #f9fafb;
}

#app {
  font-family: 'Inter', sans-serif;
  color: #111827;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

#app > *:not(header):not(footer) {
  flex-grow: 1;
}

.page-container {
  max-width: 1280px;
  width: 100%;
  margin-left: auto;
  margin-right: auto;
  padding: 32px 16px;
  box-sizing: border-box;
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
