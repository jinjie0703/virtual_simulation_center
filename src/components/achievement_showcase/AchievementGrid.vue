<template>
  <div class="achievement-grid">
    <AchievementCard
      v-for="achievement in achievements"
      :key="achievement.id"
      :achievement="achievement"
      @click="$emit('view-details', achievement)"
    />
  </div>
</template>

<script setup>
import AchievementCard from './AchievementCard.vue'
defineProps({
  achievements: Array,
})
defineEmits(['view-details'])
</script>

<style scoped>
.achievement-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(380px, 1fr));
  gap: 40px;
}
</style>
