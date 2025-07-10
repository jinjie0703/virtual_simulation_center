<template>
  <div class="page-wrapper">
    <div class="page-content">
      <InfoCenter />
    </div>
  </div>
</template>

<script setup>
import InfoCenter from '@/components/information_center/InfoCenter.vue'
</script>

<style scoped>
.page-wrapper {
  min-height: 100vh;
  background: #f0f2f5;
  padding-top: 80px; /* 为导航栏留出空间 */
}

.page-content {
  max-width: 100%;
  margin: 0;
  padding: 0;
}

/* 响应式调整 */
@media (max-width: 768px) {
  .page-wrapper {
    padding-top: 70px;
  }
}
</style>
