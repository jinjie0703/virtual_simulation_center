<!-- TeamList.vue -->
<template>
  <div>
    <!-- 加载状态 -->
    <div v-if="loading" class="loading-state">
      <div class="loading-spinner"><div class="spinner"></div></div>
      <p>加载中，请稍候...</p>
    </div>

    <!-- 空状态 -->
    <div v-else-if="teams.length === 0" class="empty-state">
      <!-- 使用插槽来自定义空状态信息 -->
      <slot name="empty-state">
        <i class="fas fa-search"></i>
        <h3>没有找到相关团队</h3>
        <p>请尝试其他搜索条件</p>
      </slot>
      <button @click="$emit('clear-filters')" class="reset-button">
        <i class="fas fa-sync-alt"></i> 查看全部团队
      </button>
    </div>

    <!-- 团队列表 -->
    <transition-group v-else class="team-list-container" name="team-list" tag="div">
      <div v-for="team in teams" :key="team.id" class="team-item-card">
        <!-- 渲染 TeamCard 并透传 join-team 事件 -->
        <TeamCard :team="team" @join-team="$emit('join-team', team)" />
      </div>
    </transition-group>
  </div>
</template>

<script setup>
import TeamCard from './TeamCard.vue'

defineProps({
  teams: {
    // 只接收最终要渲染的团队数组
    type: Array,
    required: true,
  },
  loading: {
    type: Boolean,
    default: false,
  },
})

defineEmits(['clear-filters', 'join-team'])
</script>

<style scoped>
/* 样式与原文件保持一致 */
.loading-state,
.empty-state {
  text-align: center;
  padding: 3rem 1rem;
  color: #64748b;
  min-height: 400px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  border-radius: 12px;
  background-color: #f8fafc;
  border: 1px dashed #e2e8f0;
}
.loading-spinner {
  margin-bottom: 1rem;
}
.spinner {
  width: 48px;
  height: 48px;
  margin: 0 auto;
  border: 4px solid rgba(52, 152, 219, 0.2);
  border-radius: 50%;
  border-top-color: #3498db;
  animation: spin 1s ease-in-out infinite;
  box-shadow: 0 0 10px rgba(52, 152, 219, 0.2);
}
@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}
.loading-state p {
  animation: pulse 1.5s infinite;
  font-weight: 500;
}
@keyframes pulse {
  0% {
    opacity: 0.6;
  }
  50% {
    opacity: 1;
  }
  100% {
    opacity: 0.6;
  }
}
.empty-state i {
  font-size: 3.5rem;
  margin-bottom: 1.5rem;
  color: #cbd5e1;
  opacity: 0.8;
}
.empty-state h3 {
  font-weight: 600;
  margin-bottom: 0.5rem;
  color: #334155;
  font-size: 1.2rem;
}
.empty-state p {
  margin-bottom: 1.5rem;
  color: #64748b;
}
.reset-button {
  padding: 0.7rem 1.5rem;
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
.reset-button:hover {
  background-color: #2980b9;
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(52, 152, 219, 0.3);
}
.reset-button:active {
  transform: translateY(0);
}
.team-list-container {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 1.5rem;
  min-height: 400px;
}
.team-item-card {
  position: relative;
  background-color: #ffffff;
  border-radius: 12px;
  overflow: hidden;
  padding: 1.5rem;
  box-shadow:
    0 4px 6px rgba(0, 0, 0, 0.05),
    0 1px 3px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  border: 1px solid #e2e8f0;
  transition: all 0.3s ease;
}
.team-item-card:hover {
  box-shadow:
    0 10px 15px rgba(0, 0, 0, 0.07),
    0 4px 6px rgba(0, 0, 0, 0.05);
  transform: translateY(-5px);
}
.team-list-enter-active,
.team-list-leave-active {
  transition: all 0.5s ease;
}
.team-list-enter-from,
.team-list-leave-to {
  opacity: 0;
  transform: translateY(30px);
}
.team-list-move {
  transition: transform 0.5s ease;
}
@media (max-width: 992px) {
  .team-list-container {
    grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  }
}
@media (max-width: 768px) {
  .team-list-container {
    grid-template-columns: 1fr;
    gap: 1rem;
  }
}
</style>
