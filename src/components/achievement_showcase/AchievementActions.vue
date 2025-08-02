<template>
  <div class="actions-container">
    <div class="search-wrapper">
      <img src="@/assets/achievement_showcase/search.svg" alt="Search Icon" class="search-icon" />
      <input
        type="text"
        :value="modelValue"
        @input="$emit('update:modelValue', $event.target.value)"
        placeholder="搜索项目名称或标签"
        class="search-input"
      />
    </div>

    <div class="right-actions">
      <div class="filters-wrapper">
        <CustomSelect
          :model-value="props.teacherFilter || ''"
          :options="teacherOptions"
          placeholder="按教师筛选"
          size="medium"
          min-width="140px"
          trigger-class="filter-trigger"
          options-class="filter-options"
          @update:model-value="handleTeacherChange"
        />

        <CustomSelect
          :model-value="props.tagFilter || ''"
          :options="tagOptions"
          placeholder="按标签筛选"
          size="medium"
          min-width="140px"
          trigger-class="filter-trigger"
          options-class="filter-options"
          @update:model-value="handleTagChange"
        />

        <CustomSelect
          :model-value="props.timeFilter || ''"
          :options="timeOptions"
          placeholder="发布时间"
          size="medium"
          min-width="140px"
          trigger-class="filter-trigger"
          options-class="filter-options"
          @update:model-value="handleTimeChange"
        />
      </div>

      <button class="submit-btn" @click="$emit('open-submission')">
        <img src="@/assets/achievement_showcase/submit.svg" alt="Submit Icon" class="submit-icon" />
        <span>提交项目</span>
      </button>
    </div>
  </div>
</template>

<script setup>
import CustomSelect from '@/components/common/CustomSelect.vue'

const props = defineProps({
  modelValue: String,
  teacherFilter: {
    type: String,
    default: '',
  },
  tagFilter: {
    type: String,
    default: '',
  },
  timeFilter: {
    type: String,
    default: '',
  },
  teacherOptions: {
    type: Array,
    required: true,
  },
  tagOptions: {
    type: Array,
    required: true,
  },
})

const emit = defineEmits([
  'update:modelValue',
  'update:teacherFilter',
  'update:tagFilter',
  'update:timeFilter',
  'open-submission',
])

const timeOptions = [
  { label: '所有时间', value: '' },
  { label: '一周内', value: 'week' },
  { label: '一个月内', value: 'month' },
  { label: '三个月内', value: 'three_months' },
  { label: '半年内', value: 'half_year' },
  { label: '一年内', value: 'year' },
]

// 事件处理
const handleTeacherChange = (value) => {
  emit('update:teacherFilter', value)
}

const handleTagChange = (value) => {
  emit('update:tagFilter', value)
}

const handleTimeChange = (value) => {
  emit('update:timeFilter', value)
}
</script>

<style scoped>
.actions-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 30px 0;
  margin-bottom: 40px;
  border-bottom: 1px solid #e2e8f0;
}

.search-wrapper {
  position: relative;
  width: 100%;
  max-width: 400px;
}

.search-icon {
  position: absolute;
  left: 15px;
  top: 50%;
  transform: translateY(-50%);
  width: 24px;
  height: 24px;
}

.search-input {
  width: 100%;
  padding: 12px 20px 12px 45px;
  border: 1px solid #e2e8f0;
  border-radius: 50px;
  font-size: 1rem;
  transition: all 0.3s ease;
  outline: none;
}

.search-input:focus {
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.2);
}

.right-actions {
  display: flex;
  align-items: center;
  gap: 12px;
}

.filters-wrapper {
  display: flex;
  gap: 8px;
}

/* 自定义 CustomSelect 在此页面的样式 */
:deep(.filter-trigger) {
  border-radius: 50px !important;
  border-color: #e2e8f0 !important;
  padding: 12px 16px !important;
  font-size: 0.95rem !important;
  background: white !important;
  transition: all 0.3s ease !important;
  height: 48px !important;
  box-sizing: border-box !important;
  display: flex !important;
  align-items: center !important;
  min-width: 100px !important;
}

:deep(.filter-trigger:hover) {
  border-color: #667eea !important;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1) !important;
}

:deep(.custom-select.select-open .filter-trigger) {
  border-color: #667eea !important;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.2) !important;
}

:deep(.filter-options) {
  border-radius: 12px !important;
  border-color: #e2e8f0 !important;
  box-shadow: 0 8px 25px rgba(102, 126, 234, 0.15) !important;
  margin-top: 8px !important;
}

:deep(.filter-options .select-option) {
  padding: 10px 16px !important;
  font-size: 0.95rem !important;
  border-left: none !important;
  transition: all 0.2s ease !important;
}

:deep(.filter-options .select-option:hover) {
  background-color: #f1f5f9 !important;
  color: #667eea !important;
}

:deep(.filter-options .select-option.active) {
  background-color: #eff6ff !important;
  color: #667eea !important;
  font-weight: 500 !important;
}

:deep(.custom-select) {
  min-width: 140px !important;
}

.submit-btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 12px 24px;
  border: none;
  border-radius: 50px;
  background-color: #667eea;
  color: white;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 6px rgba(102, 126, 234, 0.2);
  height: 48px;
  box-sizing: border-box;
}

.submit-btn:hover {
  background-color: #5a67d8;
  transform: translateY(-2px);
  box-shadow: 0 7px 14px rgba(102, 126, 234, 0.3);
}

.submit-icon {
  width: 20px;
  height: 20px;
}
</style>
