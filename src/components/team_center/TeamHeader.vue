<template>
  <header class="team-header">
    <div>
      <h1 class="title">团队中心</h1>
      <p class="subtitle">寻找志同道合的伙伴，一起组队参加竞赛或开展项目</p>
    </div>
  </header>
</template>

<script setup>
// This component is purely presentational.
</script>

<style scoped>
/* 标题样式 */
.team-header {
  padding: 120px 0px 80px 0px;
  text-align: center;
  position: relative;
  overflow: hidden;
  background: linear-gradient(135deg, #a1d3ff 0%, #2d92ff 100%);
}

.title {
  font-size: 3.5rem;
  font-weight: 800;
  margin: 0 0 15px;
  color: #2c3e50;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
  animation: fadeInDown 1s ease-out;
}

.subtitle {
  font-size: 1.25rem;
  max-width: 800px;
  color: gray;
  margin: 0 auto;
  opacity: 0.9;
  line-height: 1.6;
  animation: fadeInUp 1s ease-out 0.3s both;
}

@keyframes fadeInDown {
  from {
    opacity: 0;
    transform: translateY(-30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
