<template>
  <transition name="modal">
    <div v-if="show" class="join-modal-overlay" @click.self="$emit('close')">
      <div class="join-modal-content">
        <div class="modal-header">
          <h3>申请加入团队</h3>
          <button class="close-button" @click="$emit('close')">×</button>
        </div>
        <div class="modal-body" v-if="team">
          <p><strong>团队名称:</strong> {{ team.name }}</p>
          <p>
            <strong>{{ team.type === 'competition' ? '队长' : '负责人' }}:</strong>
            {{ team.leader }}
          </p>
          <div class="form-group">
            <label for="join-reason">申请理由</label>
            <textarea
              id="join-reason"
              v-model="joinReason"
              rows="3"
              placeholder="请简要描述您的技能和加入团队的原因..."
              required
            ></textarea>
          </div>
          <div class="form-group">
            <label for="contact-info">联系方式</label>
            <input
              id="contact-info"
              v-model="contactInfo"
              type="text"
              placeholder="微信/QQ/手机号等联系方式"
              required
            />
          </div>
          <div class="modal-footer">
            <button class="cancel-button" @click="$emit('close')">取消</button>
            <button
              class="submit-button"
              :disabled="!joinReason || !contactInfo"
              @click="submitApplication"
            >
              提交申请
            </button>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  show: Boolean,
  team: Object,
  teamType: String
})

const emit = defineEmits(['close', 'submit'])

const joinReason = ref('')
const contactInfo = ref('')

const submitApplication = () => {
  emit('submit', {
    teamId: props.team.id,
    reason: joinReason.value,
    contact: contactInfo.value
  })
}

watch(() => props.show, (newVal) => {
  if (newVal) {
    joinReason.value = ''
    contactInfo.value = ''
  }
})
</script>

<style scoped>
.modal-enter-active,
.modal-leave-active {
  transition: all 0.3s ease;
}
.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}
.modal-enter-from .join-modal-content,
.modal-leave-to .join-modal-content {
  transform: scale(0.9);
}
.join-modal-overlay {
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
.join-modal-content {
  background: white;
  padding: 0;
  border-radius: 12px;
  width: 90%;
  max-width: 500px;
  position: relative;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.2);
  overflow: hidden;
  transition: transform 0.3s ease;
}
.modal-header {
  background-color: #f8fafc;
  padding: 1.2rem 1.5rem;
  border-bottom: 1px solid #e2e8f0;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.modal-header h3 {
  margin: 0;
  color: #334155;
  font-weight: 600;
}
.close-button {
  background: none;
  border: none;
  font-size: 1.5rem;
  cursor: pointer;
  color: #94a3b8;
  transition: color 0.2s ease;
}
.close-button:hover {
  color: #334155;
}
.modal-body {
  padding: 1.5rem;
}
.form-group {
  margin-bottom: 1.2rem;
}
.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: #334155;
}
.form-group input,
.form-group textarea {
  width: 100%;
  padding: 0.8rem;
  border: 1px solid #e2e8f0;
  border-radius: 6px;
  font-size: 1rem;
  transition: all 0.3s ease;
}
.form-group input:focus,
.form-group textarea:focus {
  outline: none;
  border-color: #3498db;
  box-shadow: 0 0 0 3px rgba(52, 152, 219, 0.15);
}
.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
  margin-top: 1.5rem;
}
.cancel-button {
  padding: 0.7rem 1.2rem;
  background-color: #f1f5f9;
  color: #64748b;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.2s ease;
}
.cancel-button:hover {
  background-color: #e2e8f0;
  color: #334155;
}
.submit-button {
  padding: 0.7rem 1.5rem;
  background-color: #3498db;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.3s ease;
}
.submit-button:hover:not(:disabled) {
  background-color: #2980b9;
}
.submit-button:disabled {
  background-color: #cbd5e1;
  cursor: not-allowed;
  opacity: 0.7;
}
</style>
