<template>
  <header class="achievement-header">
    <div class="header-content">
      <h1 class="title">{{ title }}</h1>
      <p class="subtitle">{{ subtitle }}</p>
    </div>
  </header>
</template>

<script setup>
defineProps({
  title: String,
  subtitle: String,
})
</script>

<style scoped>
.achievement-header {
  padding: 120px 0px 80px 0px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  text-align: center;
  color: white;
  margin-bottom: 60px;
}
.title {
  font-size: 3.5rem;
  font-weight: 800;
  margin: 0 0 15px;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
  animation: fadeInDown 1s ease-out;
}
.subtitle {
  font-size: 1.25rem;
  max-width: 800px;
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
