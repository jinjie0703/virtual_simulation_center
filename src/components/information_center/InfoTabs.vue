<!-- InfoTabs.vue -->
<template>
  <div class="tabs-container">
    <div class="tabs">
      <button
        v-for="tab in tabs"
        :key="tab.id"
        :class="['tab-btn', { active: activeTab === tab.id }]"
        @click="$emit('update:activeTab', tab.id)"
      >
        {{ tab.name }}
      </button>
    </div>
  </div>
</template>

<script setup>
defineProps({
  tabs: Array,
  activeTab: String,
})
defineEmits(['update:activeTab'])
</script>

<style scoped>
.tabs-container {
  display: flex;
  justify-content: center;
  margin-bottom: 50px;
}
.tabs {
  display: inline-flex;
  background-color: #e2e8f0;
  border-radius: 999px;
  padding: 6px;
}
.tab-btn {
  padding: 12px 30px;
  border: none;
  border-radius: 999px;
  background-color: transparent;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 600;
  color: #4a5568;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
}
.tab-btn.active {
  background-color: #ffffff;
  color: #4c51bf;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}
</style>
