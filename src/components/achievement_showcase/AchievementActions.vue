<template>
  <div class="actions-container">
    <div class="search-wrapper">
      <input
        type="text"
        :value="modelValue"
        @input="$emit('update:modelValue', $event.target.value)"
        placeholder="搜索项目名称或标签..."
        class="search-input"
      />
      <svg
        class="search-icon"
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 24 24"
        fill="currentColor"
      >
        <path
          d="M15.5 14h-.79l-.28-.27A6.471 6.471 0 0 0 16 9.5 6.5 6.5 0 1 0 9.5 16c1.61 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z"
        />
      </svg>
    </div>
    <button class="submit-btn" @click="$emit('open-submission')">
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor">
        <path
          d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm5 11h-4v4h-2v-4H7v-2h4V7h2v4h4v2z"
        />
      </svg>
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

.search-icon {
  position: absolute;
  left: 15px;
  top: 50%;
  transform: translateY(-50%);
  width: 24px;
  height: 24px;
  color: #a0aec0;
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

.submit-btn svg {
  width: 20px;
  height: 20px;
}
</style>
