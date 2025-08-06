<!-- CreateTeamModal.vue -->
<template>
  <div class="modal-overlay" @click.self="emit('close')">
    <div class="modal-content">
      <div class="modal-header">
        <h2>创建{{ teamType === 'competition' ? '竞赛' : '项目' }}团队</h2>
        <button class="close-button" @click="emit('close')">×</button>
      </div>

      <form @submit.prevent="submitForm">
        <!-- 表单内容保持不变 -->
        <div class="form-group">
          <label for="team-name">团队名称</label>
          <input
            id="team-name"
            v-model="team.name"
            type="text"
            placeholder="请输入团队名称"
            required
            :maxlength="50"
          />
          <div class="form-help">简洁明了的团队名称可以吸引更多关注 (最多50字符)</div>
        </div>
        <div class="form-group">
          <label for="team-description">团队简介</label>
          <textarea
            id="team-description"
            v-model="team.description"
            rows="3"
            placeholder="请简要描述团队目标、需要的技能和成员期望"
            required
            :maxlength="200"
          ></textarea>
          <div class="input-counter">{{ team.description.length }}/200</div>
        </div>
        <div class="form-group">
          <label for="team-tags">技术标签</label>
          <div class="tags-input-container">
            <input
              id="team-tags"
              v-model="currentTag"
              @keydown.enter.prevent="addTag"
              @keydown.tab.prevent="addTag"
              @keydown.,.prevent="addTag"
              type="text"
              placeholder="输入标签后按回车、Tab或逗号添加"
            />
          </div>
          <div class="tags-display">
            <span v-for="(tag, index) in tags" :key="index" class="tag-item">
              {{ tag }}
              <button type="button" @click="removeTag(index)" class="remove-tag">×</button>
            </span>
            <span v-if="tags.length === 0" class="no-tags">请添加相关技术或主题标签 (最多5个)</span>
          </div>
        </div>
        <div class="form-row">
          <div class="form-group">
            <label for="max-members">招募人数</label>
            <input
              id="max-members"
              v-model.number="team.maxMembers"
              type="number"
              min="2"
              max="10"
              required
            />
            <div class="form-help">包括队长在内，最多10人</div>
          </div>
          <template v-if="teamType === 'competition'">
            <div class="form-group">
              <label for="deadline">报名截止日期</label>
              <input id="deadline" v-model="team.deadline" type="date" :min="minDate" required />
              <div class="form-help">设置合理的截止日期</div>
            </div>
          </template>
          <template v-if="teamType === 'project'">
            <div class="form-group">
              <label for="difficulty">项目难度</label>
              <select id="difficulty" v-model="team.difficulty" required>
                <option>简单</option>
                <option>中等</option>
                <option>较难</option>
              </select>
              <div class="form-help">根据所需技能和工作量评估</div>
            </div>
            <div class="form-group">
              <label for="duration">预计周期</label>
              <input
                id="duration"
                v-model="team.duration"
                type="text"
                placeholder="例如: 3个月"
                required
              />
              <div class="form-help">项目预计完成时间</div>
            </div>
          </template>
        </div>

        <div class="form-group contact-section">
          <label>联系方式 (至少填写一项)</label>
          <div class="form-row">
            <div class="form-group">
              <label for="contact-wechat">微信</label>
              <input id="contact-wechat" v-model="team.contact.wechat" type="text" placeholder="选填" />
            </div>
            <div class="form-group">
              <label for="contact-qq">QQ</label>
              <input id="contact-qq" v-model="team.contact.qq" type="text" placeholder="选填" />
            </div>
            <div class="form-group">
              <label for="contact-email">邮箱</label>
              <input id="contact-email" v-model="team.contact.email" type="email" placeholder="选填" />
            </div>
          </div>
        </div>
        <div class="form-actions">
          <button type="button" class="cancel-button" @click="emit('close')">取消</button>
          <button type="submit" class="submit-button" :disabled="isFormInvalid">
            <i class="fas fa-plus-circle"></i> 确认创建
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, nextTick } from 'vue'

const props = defineProps({
  teamType: { type: String, required: true },
})

const emit = defineEmits(['close', 'team-created'])

// 响应式表单数据
const team = reactive({
  name: '',
  description: '',
  maxMembers: 5,
  deadline: '',
  difficulty: '中等',
  duration: '',
  contact: {
    wechat: '',
    qq: '',
    email: '',
  },
})
const currentTag = ref('')
const tags = ref([])

// 计算属性
const minDate = computed(() => new Date().toISOString().split('T')[0])
const isFormInvalid = computed(() => {
  const isContactEmpty = !team.contact.wechat && !team.contact.qq && !team.contact.email
  return !team.name || !team.description || tags.value.length === 0 || isContactEmpty
})

// 方法
const addTag = () => {
  const tag = currentTag.value.trim()
  if (tag && !tags.value.includes(tag) && tags.value.length < 5) {
    tags.value.push(tag)
    currentTag.value = ''
  }
}

const removeTag = (index) => {
  tags.value.splice(index, 1)
}

const submitForm = () => {
  if (isFormInvalid.value) return

  const newTeam = {
    id: Date.now(),
    leader: '当前用户', // 实际应从用户状态管理获取
    memberCount: 1,
    createdAt: new Date().toISOString(),
    tags: tags.value,
    ...team,
  }

  if (props.teamType === 'competition') {
    newTeam.status = 'open'
  }

  emit('team-created', newTeam)
  emit('close')
}

// 生命周期钩子
onMounted(() => {
  nextTick(() => {
    document.getElementById('team-name')?.focus()
  })
})
</script>

<style scoped>
/* 样式与原文件保持一致 */
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
  animation: fadeIn 0.2s ease-out;
}
@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
.modal-content {
  background: white;
  padding: 0;
  border-radius: 12px;
  width: 90%;
  max-width: 550px;
  position: relative;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.2);
  animation: slideIn 0.3s cubic-bezier(0.68, -0.55, 0.27, 1.55);
  max-height: 90vh;
  overflow-y: auto;
}
@keyframes slideIn {
  from {
    transform: translateY(-30px);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}
.modal-header {
  position: sticky;
  top: 0;
  z-index: 10;
  padding: 1.5rem;
  border-bottom: 1px solid #e2e8f0;
  background-color: #f8fafc;
}
h2 {
  margin: 0;
  font-size: 1.5rem;
  color: #1e293b;
  text-align: center;
}
.close-button {
  position: absolute;
  top: 1.8rem;
  right: 1rem;
  background: none;
  border: none;
  font-size: 1.5rem;
  cursor: pointer;
  color: #94a3b8;
  transition: color 0.2s ease;
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
}
.close-button:hover {
  color: #334155;
  background-color: rgba(226, 232, 240, 0.5);
}
form {
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
.form-help {
  margin-top: 0.3rem;
  font-size: 0.8rem;
  color: #64748b;
}
.form-group input,
.form-group textarea,
.form-group select {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  font-size: 1rem;
  transition: all 0.2s ease;
}
.form-group input:focus,
.form-group textarea:focus,
.form-group select:focus {
  outline: none;
  border-color: #3498db;
  box-shadow: 0 0 0 3px rgba(52, 152, 219, 0.15);
}
.form-group textarea {
  resize: vertical;
  min-height: 80px;
}
.form-group input::placeholder,
.form-group textarea::placeholder {
  color: #94a3b8;
}
.input-counter {
  text-align: right;
  margin-top: 0.3rem;
  font-size: 0.8rem;
  color: #64748b;
}
.form-row {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 1rem;
}
.tags-input-container {
  position: relative;
}
.tags-display {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  margin-top: 0.75rem;
  min-height: 30px;
}
.tag-item {
  background-color: #edf2f7;
  border-radius: 4px;
  padding: 0.3rem 0.6rem;
  font-size: 0.9rem;
  color: #4a5568;
  display: flex;
  align-items: center;
  gap: 0.3rem;
}
.remove-tag {
  background: none;
  border: none;
  cursor: pointer;
  color: #718096;
  font-size: 1.1rem;
  padding: 0;
  line-height: 1;
  transition: color 0.2s;
}
.remove-tag:hover {
  color: #e53e3e;
}
.no-tags {
  color: #a0aec0;
  font-style: italic;
  font-size: 0.9rem;
}
.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
  margin-top: 2rem;
}
.cancel-button {
  padding: 0.75rem 1.5rem;
  background-color: #f1f5f9;
  color: #64748b;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.2s ease;
}
.cancel-button:hover {
  background-color: #e2e8f0;
  color: #334155;
}
.submit-button {
  padding: 0.75rem 1.5rem;
  background-color: #3498db;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}
.submit-button:hover:not(:disabled) {
  background-color: #2980b9;
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(52, 152, 219, 0.3);
}
.submit-button:active:not(:disabled) {
  transform: translateY(0);
}
.submit-button:disabled {
  background-color: #cbd5e1;
  cursor: not-allowed;
  opacity: 0.7;
}
.submit-button i {
  font-size: 1.1rem;
}
.contact-section {
  margin-top: 1.5rem;
  padding-top: 1.5rem;
  border-top: 1px solid #e2e8f0;
}
.contact-section > label {
  font-size: 1.1rem;
  font-weight: 600;
  margin-bottom: 1rem;
}
.contact-section .form-group {
  margin-bottom: 0; /* 移除内联表单组的下边距 */
}
.contact-section .form-group label {
  font-size: 0.9rem;
  font-weight: 400;
}
@media (max-width: 640px) {
  .modal-content {
    width: 95%;
    max-height: 95vh;
  }
  .form-row {
    grid-template-columns: 1fr;
  }
  .form-actions {
    flex-direction: column;
  }
  .cancel-button,
  .submit-button {
    width: 100%;
  }
  .cancel-button {
    order: 2;
  }
  .submit-button {
    order: 1;
    margin-bottom: 0.5rem;
  }
}
</style>
