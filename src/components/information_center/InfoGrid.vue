<template>
  <div v-if="items.length > 0" class="info-grid">
    <InfoCard v-for="item in items" :key="item.id" :item="item" :type="itemType" />
  </div>
  <div v-else class="empty-state">
    <h3>暂无内容</h3>
    <p>此分类下暂时没有信息，请稍后再来查看。</p>
  </div>
</template>

<script setup>
import InfoCard from './InfoCard.vue'

defineProps({
  items: Array,
  itemType: String,
})
</script>

<style scoped>
.info-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 30px;
}

.empty-state {
  text-align: center;
  padding: 80px 20px;
  color: #a0aec0;
}
.empty-state h3 {
  font-size: 1.5rem;
  color: #4a5568;
}

@media (max-width: 992px) {
  .info-grid {
    grid-template-columns: 1fr;
    gap: 25px;
  }
}
</style>
