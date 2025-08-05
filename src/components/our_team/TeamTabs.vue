<template>
  <div class="tab-container">
    <!-- 搜索框区域 -->
    <div class="search-section">
      <div class="search-box">
        <input
          type="text"
          :value="searchKeyword"
          @input="handleSearch"
          placeholder="搜索团队成员..."
          class="search-input"
          :disabled="isLoading"
        />
      </div>
    </div>

    <!-- 标签区域 -->
    <div class="tab-switcher">
      <button
        v-for="tab in tabs"
        :key="tab.id"
        :class="['tab-button', { active: activeTab === tab.id }]"
        @click="handleTabClick(tab.id)"
        :disabled="isLoading"
      >
        <span class="tab-text">{{ tab.name }}</span>
      </button>
    </div>

    <!-- 占位符，用于辅助居中 -->
    <div class="placeholder-section"></div>
  </div>
</template>

<script setup>
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
  searchKeyword: {
    type: String,
    default: '',
  },
})

const emit = defineEmits(['switchTab', 'search'])

const handleTabClick = (tabId) => {
  emit('switchTab', tabId)
}

const handleSearch = (event) => {
  emit('search', event.target.value)
}
</script>

<style scoped>
.tab-container {
  --primary-color: #4a90e2;
  --secondary-color: #50e3c2;
  --border-color: #e0e0e0;
  --border-radius: 30px; /* Reverted from pill to original */
  --transition-fast: 0.3s ease;

  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 50px;
  gap: 30px;
  padding: 0 8%;
}

.search-section,
.placeholder-section {
  flex: 1;
}

.search-section {
  display: flex;
  justify-content: flex-start;
}

.placeholder-section {
  display: flex;
  justify-content: flex-end;
}

.search-box {
  display: flex;
  align-items: center;
  background: white;
  border: 2px solid var(--border-color);
  border-radius: var(--border-radius);
  padding: 12px 20px;
  transition: all var(--transition-fast);
  width: 100%;
  max-width: 320px;
}

.search-box:hover,
.search-box:focus-within {
  border-color: var(--primary-color);
  box-shadow: 0 4px 12px rgba(74, 144, 226, 0.15);
}

.search-input {
  border: none;
  outline: none;
  background: transparent;
  font-size: 16px;
  color: #333;
  width: 100%;
  font-weight: 500;
}

.search-input::placeholder {
  color: #aaa;
}

.search-input:disabled,
.tab-button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.tab-switcher {
  display: flex;
  justify-content: center;
  gap: 30px;
}

.tab-button {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 14px 32px;
  border: 2px solid var(--border-color);
  background: white;
  /* border-radius: var(--border-radius); */
  border-radius: 30px;
  font-size: 18px;
  font-weight: 600;
  color: #666;
  cursor: pointer;
  transition: all var(--transition-fast);
  min-width: 160px;
  /* overflow: hidden; */
}

.tab-button:hover:not(.active):not(:disabled) {
  border-color: var(--primary-color);
  color: var(--primary-color);
  transform: translateY(-2px);
}

.tab-button.active {
  color: white;
  border-color: transparent;
  background: linear-gradient(135deg, var(--primary-color), var(--secondary-color));
  background-origin: border-box;
  box-shadow: 0 8px 20px rgba(74, 144, 226, 0.25);
  transform: translateY(-2px);
}

.tab-text {
  white-space: nowrap;
}
</style>
