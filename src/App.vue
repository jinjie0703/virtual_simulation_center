<template>
  <div id="app">
    <!-- 全局头部组件 -->
    <TheHeader />

    <!--
      RouterView 会渲染当前路由对应的页面组件。
      每个页面组件将负责自己的内部布局。
    -->
    <RouterView />

    <!--
      ✨ 核心修改：在这里添加页脚的条件渲染逻辑 ✨
      它会根据路由的 meta 信息来决定是否显示 SiteFooter 组件。
      这个逻辑与页面的具体布局无关，可以独立工作。
    -->
    <SiteFooter v-if="$route.meta.showFooter !== false" />
  </div>
</template>

<script setup>
// 注意：lang="js" 可以省略，setup 默认就是 JS
import { RouterView } from 'vue-router'
import TheHeader from './components/layout/TheHeader.vue'
// ✨ 导入我们之前创建的页脚组件
import SiteFooter from './components/layout/SiteFooter.vue'
</script>

<!--
  将 scoped 改为 lang="scss" 或移除，以便 .page-container 成为全局可用类。
  或者保持 scoped，但在每个需要它的页面组件中单独定义。
  这里我假设您希望它是一个全局可用的工具类，因此移除了 scoped。
-->
<style scoped>
/* 引入字体 */
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700;800&display=swap');

/* CSS Reset 和 基础样式 */
body {
  margin: 0;
  background-color: #f9fafb; /* 将背景色移到 body 上更合适 */
}

#app {
  font-family: 'Inter', sans-serif;
  color: #111827;
  min-height: 100vh;
  /*
    使用 Flex 布局，并让 RouterView (通过其子组件) 自动撑开空间，
    从而将页脚推到底部。
  */
  display: flex;
  flex-direction: column;
}

/*
  让 RouterView 渲染出的组件能够占据剩余的所有空间。
  这需要给 RouterView 的直接子组件（即你的页面组件）添加 flex-grow: 1。
  更简单的方式是直接给 <RouterView /> 的父级（除了 #app）应用 flex-grow: 1，
  但由于这里没有父级，我们可以在每个页面组件的根元素上应用这个样式。
  或者，我们可以在这里定义一个通用的规则。
*/
#app > *:not(header):not(footer) {
  flex-grow: 1;
}

/*
  这是一个可复用的全局布局容器类。
  任何需要居中、有内边距的页面都可以在其根元素上使用 <div class="page-container">
*/
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
