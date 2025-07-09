<template>
  <div class="page-wrapper">
    <div class="page-content">
      <InformationCenter />
    </div>
  </div>
</template>

<script setup>
import InformationCenter from '@/components/information_center/InformationCenter1.vue'
</script>

<style scoped>
.page-wrapper {
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
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

@media (max-width: 480px) {
  .page-wrapper {
    padding-top: 60px;
  }
}
</style>
