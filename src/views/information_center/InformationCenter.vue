<template>
  <header>
    <h1>我的信息中心</h1>
    <p>在这里，您可以找到最新的公司动态和行业资讯。</p>
  </header>
  <main>
    <InformationCenter />
  </main>
</template>

<script setup>
import InformationCenter from '@/components/information_center/InformationCenter1.vue'
</script>

<style scoped>
/* 全局样式 */
body {
  margin: 0;
  font-family:
    -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  background-color: #f4f7f9;
  color: #333;
}

#app {
  padding: 20px;
}

header {
  text-align: center;
  margin-bottom: 40px;
}

header h1 {
  font-size: 2.5rem;
  color: #2c3e50;
}

header p {
  font-size: 1.1rem;
  color: #555;
}
</style>
