<template>
  <div class="tab-container">
    <!-- 搜索框区域 -->
    <div class="search-section">
      <div class="search-box">
        <span class="search-icon">🔍</span>
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
        <span class="tab-icon">{{ tab.icon }}</span>
        <span class="tab-text">{{ tab.name }}</span>
        <span class="tab-count">({{ tab.count }})</span>
      </button>
    </div>
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
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  margin-bottom: 50px;
  gap: 30px;
  position: relative;
  padding: 0 8%;
}

.search-section {
  flex: 0 0 320px;
  display: flex;
  justify-content: flex-start;
  margin-left: -8%;
  padding-left: 8%;
}

.search-box {
  display: flex;
  align-items: center;
  background: white;
  border: 3px solid #e1e5e9;
  border-radius: 30px;
  padding: 12px 20px;
  gap: 12px;
  transition: all 0.3s ease;
  width: 100%;
}

.search-box:hover {
  border-color: #4a90e2;
  box-shadow: 0 5px 20px rgba(74, 144, 226, 0.1);
}

.search-box:focus-within {
  border-color: #4a90e2;
  box-shadow: 0 5px 20px rgba(74, 144, 226, 0.2);
}

.search-icon {
  font-size: 18px;
  color: #666;
  flex-shrink: 0;
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
  color: #999;
}

.search-input:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.tab-switcher {
  display: flex;
  justify-content: center;
  gap: 30px;
  flex-wrap: wrap;
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  width: max-content;
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

.tab-count {
  font-size: 14px;
  opacity: 0.8;
}

@media (max-width: 1400px) {
  .tab-container {
    padding: 0 6%;
  }

  .search-section {
    flex: 0 0 300px;
    margin-left: -6%;
    padding-left: 6%;
  }
}

@media (max-width: 1100px) {
  .tab-container {
    padding: 0 4%;
  }

  .search-section {
    flex: 0 0 280px;
    margin-left: -4%;
    padding-left: 4%;
  }
}

@media (max-width: 1024px) {
  .tab-container {
    flex-direction: column;
    gap: 20px;
    padding: 0 3%;
  }

  .search-section {
    flex: none;
    width: 100%;
    margin: 0;
    padding: 0;
  }

  .search-section {
    justify-content: flex-start;
  }

  .tab-switcher {
    gap: 15px;
  }
}

@media (max-width: 800px) {
  .tab-container {
    padding: 0 3%;
  }

  .search-section {
    margin-left: -3%;
    padding-left: 3%;
  }
}

@media (max-width: 768px) {
  .tab-container {
    flex-direction: column;
    gap: 20px;
    padding: 0 2%;
  }

  .search-section {
    margin: 0;
    padding: 0;
  }

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

  .search-box {
    padding: 10px 16px;
  }

  .filter-button {
    padding: 10px 16px;
    font-size: 15px;
  }
}

@media (max-width: 600px) {
  .tab-container {
    padding: 0 2%;
  }

  .search-section {
    margin-left: -2%;
    padding-left: 2%;
  }
}

@media (max-width: 480px) {
  .tab-button {
    min-width: 180px;
    padding: 10px 20px;
    font-size: 15px;
  }

  .search-box {
    padding: 8px 14px;
    font-size: 14px;
  }
}
</style>
