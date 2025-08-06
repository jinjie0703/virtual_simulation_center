<template>
  <transition name="modal">
    <div v-if="show" class="join-modal-overlay" @click.self="$emit('close')">
      <div class="join-modal-content">
        <div class="modal-header">
          <h3>加入团队</h3>
          <button class="close-button" @click="$emit('close')">×</button>
        </div>
        <div class="modal-body" v-if="team">
          <div class="team-info-box">
            <p><strong>团队名称:</strong> {{ team.name }}</p>
            <p><strong>{{ team.type === 'competition' ? '队长' : '负责人' }}:</strong> {{ team.leader }}</p>
          </div>

          <div class="team-info-box">
            <h4>联系方式</h4>
            <div v-if="team.contact && (team.contact.wechat || team.contact.qq || team.contact.email)">
              <p v-if="team.contact.wechat"><strong>微信:</strong> {{ team.contact.wechat }}</p>
              <p v-if="team.contact.qq"><strong>QQ:</strong> {{ team.contact.qq }}</p>
              <p v-if="team.contact.email"><strong>邮箱:</strong> {{ team.contact.email }}</p>
            </div>
            <p v-else>该团队暂未提供公开联系方式。</p>
          </div>

          <div class="notice-box">
            <p><strong>如何加入:</strong> 由于网站尚在开发初期，暂不支持在线申请，请通过上方联系方式主动联系负责人，如果选择发送邮件，请作个人介绍并且说明自己的技术特长和目标。</p>
          </div>

          <div class="modal-footer">
            <button class="close-button-footer" @click="$emit('close')">关闭</button>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script setup>
defineProps({
  show: Boolean,
  team: Object,
  teamType: String
})

defineEmits(['close'])
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
  display: grid;
  grid-template-columns: 1fr auto 1fr;
  background-color: #f8fafc;
  padding: 1.2rem 1.5rem;
  border-bottom: 1px solid #e2e8f0;
  justify-content: space-between;
  align-items: center;
}
.modal-header h3 {
  grid-column: 2;
  flex-grow: 1;
  text-align: center;
  font-size: 1.5rem;
  margin: 0;
  color: #334155;
  font-weight: 600;
}
.close-button {
  grid-column: 3; /* 3. 按钮放在第三列 */
    justify-self: end;
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
.team-info-box {
  background-color: #ffffff;
  border-radius: 8px;
  margin-bottom: 1.5rem;
  border: 1px solid #e2e8f0;
  padding: 1.2rem;
}
.team-info-box p {
  margin: 0 0 0.8rem 0;
  color: #334155;
  line-height: 1.6;
  word-break: break-all;
}
.team-info-box p:last-child {
  margin-bottom: 0;
}
.team-info-box h4 {
  margin-top: 0;
  margin-bottom: 1rem;
  font-weight: 600;
  color: #334155;
  border-bottom: 1px solid #e2e8f0;
  padding-bottom: 0.8rem;
}
.notice-box {
  padding: 1rem;
  background-color: #f8fafc;
  border-radius: 8px;
  border: 1px solid #e2e8f0;
  color: #475569;
  line-height: 1.6;
}
.notice-box p {
  margin: 0;
}
.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
  margin-top: 1.5rem;
}
.close-button-footer {
  padding: 0.7rem 1.5rem;
  background-color: #3498db;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.3s ease;
}
.close-button-footer:hover {
  background-color: #2980b9;
}
</style>
