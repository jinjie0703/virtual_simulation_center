<template>
  <div class="header-section">
    <h1 class="main-title">{{ title }}</h1>
    <p class="main-subtitle">{{ subtitle }}</p>
  </div>
</template>

<script setup>
defineProps({
  title: {
    type: String,
    required: true,
  },
  subtitle: {
    type: String,
    required: true,
  },
})
</script>

<style scoped>
.header-section {
  text-align: center;
  margin-bottom: 40px;
}

.main-title {
  font-size: 3.5rem;
  font-weight: 800;
  color: #1a202c;
  margin: 0 0 10px 0;
  letter-spacing: -1px;
}

.main-subtitle {
  font-size: 1.25rem;
  color: #718096;
  margin: 0;
}
</style>
