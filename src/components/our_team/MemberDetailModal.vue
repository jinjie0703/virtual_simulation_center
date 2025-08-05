<template>
  <transition name="modal-fade">
    <div v-if="show" class="modal-overlay" @click.self="close">
      <div class="modal-container" :class="{ 'is-visible': show }">
        <button class="close-button" @click="close">&times;</button>
        <div v-if="member" class="modal-content">
          <div class="modal-header">
            <img :src="member.avatar" :alt="member.name" class="modal-avatar" />
            <div class="header-info">
              <h2 class="member-name">{{ member.name }}</h2>
              <p class="member-title">{{ member.title }}</p>
            </div>
          </div>
          <div class="modal-body">
            <div v-for="detail in details" :key="detail.label" class="detail-item">
              <span class="label">{{ detail.label }}:</span>
              <span class="value">{{ detail.value }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  show: Boolean,
  member: Object,
})

const emit = defineEmits(['close'])

const close = () => {
  emit('close')
}

const details = computed(() => {
  if (!props.member) return []
  return [
    { label: '联系方式', value: props.member.email },
    { label: '办公室', value: props.member.office },
    { label: '研究方向', value: props.member.research },
    { label: '学术成果', value: props.member.achievements },
  ].filter((detail) => detail.value)
})
</script>

<style scoped>
.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: opacity 0.3s ease;
}
.modal-fade-enter-from,
.modal-fade-leave-to {
  opacity: 0;
}

.modal-overlay {
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
  background: #ffffff;
  border-radius: 16px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
  width: 90%;
  max-width: 600px;
  max-height: 90vh;
  overflow-y: auto;
  transform: scale(0.95);
  opacity: 0;
  transition:
    transform 0.3s ease,
    opacity 0.3s ease;
}

.modal-container.is-visible {
  transform: scale(1);
  opacity: 1;
}

.close-button {
  position: absolute;
  top: 15px;
  right: 15px;
  background: none;
  border: none;
  font-size: 2rem;
  color: #aaa;
  cursor: pointer;
  transition: color 0.2s;
}
.close-button:hover {
  color: #333;
}

.modal-content {
  padding: 40px;
}

.modal-header {
  display: flex;
  align-items: center;
  border-bottom: 1px solid #eee;
  padding-bottom: 25px;
  margin-bottom: 25px;
}

.modal-avatar {
  width: 100px;
  height: 100px;
  border-radius: 50%;
  object-fit: cover;
  margin-right: 25px;
  border: 4px solid #4a90e2;
}

.header-info {
  flex: 1;
}

.member-name {
  font-size: 2rem;
  font-weight: 700;
  color: #2c3e50;
  margin: 0 0 5px;
}

.member-title {
  font-size: 1.1rem;
  color: #4a90e2;
  font-weight: 500;
}

.modal-body {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.detail-item {
  display: flex;
  flex-direction: column;
  gap: 5px;
}

.label {
  font-size: 0.9rem;
  color: #718096;
  font-weight: 600;
  display: flex;
  align-items: center;
  gap: 8px;
}

.value {
  font-size: 1rem;
  color: #2d3748;
  line-height: 1.6;
}

.value-link {
  font-size: 1rem;
  color: #4a90e2;
  text-decoration: none;
  transition: color 0.2s;
}
.value-link:hover {
  color: #2980b9;
}
</style>
