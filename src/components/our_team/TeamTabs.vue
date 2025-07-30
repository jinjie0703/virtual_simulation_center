<template>
  <div class="tab-switcher">
    <button
      v-for="tab in tabs"
      :key="tab.id"
      :class="['tab-button', { active: activeTab === tab.id }]"
      @click="handleTabClick(tab.id)"
      :disabled="isLoading"
    >
      <span class="tab-icon">{{ tab.icon }}</span>
      <span class="tab-text">{{ tab.name }}</span>
    </button>
  </div>
</template>

<script setup>
import {} from 'vue'

defineProps({
  tabs: {
    type: Array,
    required: true,
  },
  activeTab: {
    type: String,
    required: true,
  },
  isLoading: {
    type: Boolean,
    default: false,
  },
})

const emit = defineEmits(['switchTab'])

const handleTabClick = (tabId) => {
  emit('switchTab', tabId)
}
</script>

<style scoped>
.tab-switcher {
  display: flex;
  justify-content: center;
  margin-bottom: 50px;
  gap: 30px;
  flex-wrap: wrap;
}

.tab-button {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 16px 32px;
  border: 3px solid #e1e5e9;
  background: white;
  border-radius: 30px;
  font-size: 18px;
  font-weight: 600;
  color: #666;
  cursor: pointer;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
  min-width: 160px;
  justify-content: center;
}

.tab-button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.tab-button::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(135deg, #4a90e2, #50e3c2);
  transition: left 0.4s ease;
  z-index: -1;
}

.tab-button:hover:not(:disabled) {
  border-color: #4a90e2;
  transform: translateY(-3px);
  box-shadow: 0 10px 30px rgba(74, 144, 226, 0.2);
}

.tab-button:hover:not(:disabled)::before {
  left: 0;
}

.tab-button:hover:not(:disabled) {
  color: white;
}

.tab-button.active {
  background: linear-gradient(135deg, #4a90e2, #50e3c2);
  border-color: #4a90e2;
  color: white;
  transform: translateY(-3px);
  box-shadow: 0 15px 40px rgba(74, 144, 226, 0.4);
}

.tab-button.active::before {
  left: 0;
}

.tab-icon {
  font-size: 22px;
  flex-shrink: 0;
}

.tab-text {
  white-space: nowrap;
}

@media (max-width: 768px) {
  .tab-switcher {
    gap: 15px;
    flex-direction: column;
    align-items: center;
  }

  .tab-button {
    padding: 12px 24px;
    font-size: 16px;
    min-width: 200px;
  }

  .tab-icon {
    font-size: 20px;
  }
}

@media (max-width: 480px) {
  .tab-button {
    min-width: 180px;
    padding: 10px 20px;
    font-size: 15px;
  }
}
</style>
