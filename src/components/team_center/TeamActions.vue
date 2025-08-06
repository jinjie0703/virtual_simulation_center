<template>
  <div class="team-actions-container">
    <div class="tabs-container">
      <div class="tabs">
        <button
          :class="['tab-button', { active: activeTab === 'competition' }]"
          @click="$emit('update:activeTab', 'competition')"
        >
          <img
            src="@/assets/team_center/competition.svg"
            alt="竞赛组队图标"
            class="competition-icon"
          />
          <span>竞赛组队</span>
        </button>
        <button
          :class="['tab-button', { active: activeTab === 'project' }]"
          @click="$emit('update:activeTab', 'project')"
        >
          <img src="@/assets/team_center/project.svg" alt="项目组队图标" class="project-icon" />
          <span>项目组队</span>
        </button>
      </div>
    </div>

    <div class="search-filter-container">
      <div class="search-box">
        <img src="@/assets/team_center/search.svg" />
        <input
          :value="searchQuery"
          @input="$emit('update:searchQuery', $event.target.value)"
          type="text"
          placeholder="搜索团队..."
        />
      </div>

      <div class="filter-options">
        <CustomSelect
          :model-value="filterOption"
          @update:modelValue="$emit('update:filterOption', $event)"
          :options="filterOptions"
          placeholder="选择排序方式"
          size="medium"
          :min-width="'140px'"
          :icon-path="unfoldIcon"
        />
        <CustomSelect
          :model-value="difficultyFilter"
          @update:modelValue="$emit('update:difficultyFilter', $event)"
          :options="difficultyOptions"
          placeholder="选择难度"
          size="medium"
          :min-width="'120px'"
          :icon-path="unfoldIcon"
        />
        <CustomSelect
          :model-value="tagFilter"
          @update:modelValue="$emit('update:tagFilter', $event)"
          :options="tagOptions"
          placeholder="选择标签"
          size="medium"
          :min-width="'120px'"
          :icon-path="unfoldIcon"
        />
      </div>

      <button class="create-team-button" @click="$emit('create-team')">
        <i class="fas fa-plus"></i> 创建团队
      </button>
    </div>
  </div>
</template>

<script setup>
import CustomSelect from '@/components/common/CustomSelect.vue'
import unfoldIcon from '@/assets/team_center/unfold.svg'

defineProps({
  activeTab: String,
  searchQuery: String,
  filterOption: String,
  difficultyFilter: String,
  tagFilter: String,
  filterOptions: Array,
  difficultyOptions: Array,
  tagOptions: Array
})

defineEmits([
  'update:activeTab',
  'update:searchQuery',
  'update:filterOption',
  'update:difficultyFilter',
  'update:tagFilter',
  'create-team'
])
</script>

<style scoped>
.tabs-container {
  display: flex;
  justify-content: center;
  margin-bottom: 2rem;
}
.tabs {
  display: flex;
  background-color: #f1f5f9;
  border-radius: 12px;
  padding: 0.4rem;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
}
.tab-button {
  display: inline-flex;
  align-items: center;
  gap: 12px;
  padding: 0.9rem 1.6rem;
  border: none;
  background-color: transparent;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 500;
  color: #64748b;
  transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
  border-radius: 8px;
  min-width: 140px;
}
.tab-button.active {
  background-color: #ffffff;
  color: #3498db;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.12);
  transform: translateY(-1px);
}
.tab-button:hover:not(.active) {
  color: #3498db;
  background-color: rgba(255, 255, 255, 0.5);
}
.tab-button i {
  margin-right: 0.5rem;
}
.competition-icon,
.project-icon {
  width: 24px;
  height: 24px;
}
.search-filter-container {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
  margin-bottom: 2rem;
  align-items: center;
}
.search-box {
  flex-grow: 1;
  position: relative;
  min-width: 250px;
}
.search-box img {
  width: 20px;
  height: 20px;
  position: absolute;
  top: 50%;
  left: 10px;
  transform: translateY(-50%);
}
.search-box input {
  width: 100%;
  padding: 0.9rem 2.5rem;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  font-size: 1rem;
  transition: all 0.3s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
}
.search-box input:focus {
  outline: none;
  border-color: #3498db;
  box-shadow: 0 0 0 3px rgba(52, 152, 219, 0.15);
}
.filter-options {
  display: flex;
  align-items: center;
  gap: 10px;
  flex-wrap: wrap;
  min-height: 40px;
}
:deep(.custom-select) {
  margin-bottom: 4px;
  margin-top: 4px;
}
.create-team-button {
  padding: 0.9rem 1.6rem;
  background-color: #3498db;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.3s ease;
  white-space: nowrap;
  box-shadow: 0 2px 6px rgba(52, 152, 219, 0.4);
}
.create-team-button:hover {
  background-color: #2980b9;
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(52, 152, 219, 0.5);
}
.create-team-button:active {
  transform: translateY(0);
  box-shadow: 0 1px 3px rgba(52, 152, 219, 0.4);
}
.create-team-button i {
  margin-right: 0.5rem;
}

@media (max-width: 768px) {
  .tabs {
    width: 100%;
  }
  .tab-button {
    min-width: 0;
    flex: 1;
    padding: 0.8rem 0.6rem;
    font-size: 0.9rem;
  }
  .search-filter-container {
    flex-direction: column;
    align-items: stretch;
  }
  .search-box,
  .filter-options,
  .create-team-button {
    width: 100%;
  }
  .create-team-button {
    order: -1;
    margin-bottom: 1rem;
  }
  .filter-options {
    width: 100%;
    justify-content: space-between;
    gap: 8px;
  }
  :deep(.custom-select) {
    width: 100%;
  }
}
</style>
