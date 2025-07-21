<!-- InfoHeader.vue -->
<template>
  <div class="header">
    <h1 class="title">{{ title }}</h1>
    <p class="subtitle">{{ subtitle }}</p>
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
.header {
  padding: 120px 0px 80px 0px;
  margin-bottom: 40px;
  background: linear-gradient(135deg, #84fab0, #8fd3f4);
  color: #2d3748;
  text-align: center;
  position: relative;
  overflow: hidden;
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
  opacity: 0.75;
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
