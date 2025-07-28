<template>
  <div class="actions-container">
    <div class="search-wrapper">
      <img src="@/assets/achievement_showcase/search.svg" alt="Search Icon" class="search-icon" />
      <input
        type="text"
        :value="modelValue"
        @input="$emit('update:modelValue', $event.target.value)"
        placeholder="搜索项目名称或标签..."
        class="search-input"
      />
    </div>
    <button class="submit-btn" @click="$emit('open-submission')">
      <img src="@/assets/achievement_showcase/submit.svg" alt="Submit Icon" class="submit-icon" />
      <span>提交新项目</span>
    </button>
  </div>
</template>

<script setup>
defineProps({
  modelValue: String,
})
defineEmits(['update:modelValue', 'open-submission'])
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
