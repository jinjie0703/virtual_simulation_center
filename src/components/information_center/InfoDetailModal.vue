<!-- InfoDetailModal.vue -->
<template>
  <transition name="modal-fade">
    <div v-if="item" class="modal-backdrop" @click.self="close">
      <div class="modal-container">
        <div class="modal-header">
          <h2 class="modal-title">{{ item.title }}</h2>
          <button @click="close" class="close-button">×</button>
        </div>
        <div class="modal-body">
          <div class="meta-tags">
            <span class="tag category-tag">{{ item.category }}</span>
            <span class="tag level-tag">{{ item.level }}</span>
            <span v-for="tag in item.tags" :key="tag" class="tag">{{ tag }}</span>
          </div>
          <p class="summary">{{ item.summary }}</p>
          <div class="details-grid">
            <div>
              <strong>发布日期</strong><span>📅 {{ item.date }}</span>
            </div>
            <div>
              <strong>截止日期</strong><span>⏳ {{ item.deadline }}</span>
            </div>
            <div>
              <strong>任务奖励</strong><span class="reward">💰 {{ item.reward }}</span>
            </div>
          </div>
          <div class="full-content">
            <h3>详细说明</h3>
            <p>{{ item.fullContent || '暂无更多详细说明。' }}</p>
          </div>
        </div>
        <div class="modal-footer">
          <button @click="close" class="btn-primary">关闭</button>
        </div>
      </div>
    </div>
  </transition>
</template>

<script setup>
defineProps({
  item: {
    type: Object,
    default: null,
  },
})

const emit = defineEmits(['close'])

const close = () => {
  emit('close')
}
</script>

<style scoped>
/* 样式... (省略了详细样式，下面会统一提供) */
.modal-backdrop {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.6);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}
.modal-container {
  background-color: white;
  border-radius: 12px;
  box-shadow: 0 5px 25px rgba(0, 0, 0, 0.2);
  width: 90%;
  max-width: 600px;
  max-height: 90vh;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}
/* ... 更多样式 ... */
.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: opacity 0.3s ease;
}
.modal-fade-enter-from,
.modal-fade-leave-to {
  opacity: 0;
}
.close-button {
  background: none;
  border: none;
  font-size: 24px;
  cursor: pointer;
}
/* ... */
</style>
