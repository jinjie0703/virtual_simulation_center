<!-- TeamCard.vue -->
<template>
  <div class="team-card-content">
    <!-- 卡片头部 -->
    <div class="team-header">
      <h3>{{ team.name }}</h3>
      <!-- 动态显示竞赛状态或项目难度 -->
      <span v-if="team.status" class="team-status" :class="team.status">{{
        getStatusText(team.status)
      }}</span>
      <span v-else-if="team.difficulty" class="team-difficulty" :class="team.difficulty">{{
        team.difficulty
      }}</span>
    </div>

    <!-- 标签 -->
    <div class="team-tags">
      <span v-for="(tag, index) in team.tags" :key="index" class="team-tag">{{ tag }}</span>
    </div>

    <!-- 团队信息 -->
    <div class="team-info">
      <span class="team-leader"><img src="@/assets/team_center/leader.svg"> {{ team.leader }}</span>
      <p class="team-description">{{ team.description }}</p>
      <div class="team-meta">
        <span class="team-meta-item"
          ><img src="@/assets/team_center/member.svg"> {{ team.memberCount }}/{{ team.maxMembers }}</span
        >
        <!-- 动态显示截止日期或项目周期 -->
        <span v-if="team.deadline" class="team-meta-item"
          ><img src="@/assets/team_center/deadline.svg"> {{ team.deadline }}</span
        >
        <span v-else-if="team.duration" class="team-meta-item"
          ><img src="@/assets/team_center/duration.svg"> {{ team.duration }}</span
        >
      </div>
    </div>

    <!-- 加入按钮 -->
    <button
      class="join-button"
      @click="$emit('join-team', team)"
      :disabled="team.status === 'closed'"
    >
      <i class="fas" :class="team.status === 'closed' ? 'fa-lock' : 'fa-sign-in-alt'"></i>
      {{ team.status === 'closed' ? '已截止' : '申请加入' }}
    </button>
  </div>
</template>

<script setup>
defineProps({
  team: {
    type: Object,
    required: true,
  },
})

defineEmits(['join-team'])

const getStatusText = (status) => {
  return status === 'open' ? '招募中' : '已截止'
}
</script>

<style scoped>
/* 样式合并了 CompetitionTeam 和 ProjectTeam 的所有相关样式 */
.team-card-content {
  height: 100%;
  display: flex;
  flex-direction: column;
}
.team-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}
.team-header h3 {
  font-size: 1.2rem;
  font-weight: 600;
  color: #1e293b;
  margin: 0;
  line-height: 1.4;
  flex: 1;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
}
.team-status,
.team-difficulty {
  font-size: 0.85rem;
  padding: 0.4rem 0.8rem;
  border-radius: 50px;
  font-weight: 500;
  white-space: nowrap;
  margin-left: 0.5rem;
}
.team-status.open {
  background-color: rgba(14, 165, 233, 0.15);
  color: #0284c7;
  border: 1px solid rgba(14, 165, 233, 0.3);
}
.team-status.closed {
  background-color: rgba(148, 163, 184, 0.15);
  color: #64748b;
  border: 1px solid rgba(148, 163, 184, 0.3);
}
.team-difficulty.简单 {
  background-color: rgba(34, 197, 94, 0.15);
  color: #16a34a;
  border: 1px solid rgba(34, 197, 94, 0.3);
}
.team-difficulty.中等 {
  background-color: rgba(249, 115, 22, 0.15);
  color: #ea580c;
  border: 1px solid rgba(249, 115, 22, 0.3);
}
.team-difficulty.较难 {
  background-color: rgba(239, 68, 68, 0.15);
  color: #dc2626;
  border: 1px solid rgba(239, 68, 68, 0.3);
}
.team-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  margin-bottom: 1rem;
}
.team-tag {
  font-size: 0.8rem;
  background-color: #f1f5f9;
  color: #64748b;
  padding: 0.3rem 0.7rem;
  border-radius: 50px;
  display: inline-block;
}
.team-info {
  flex: 1;
  margin-bottom: 1rem;
}
.team-leader {
  display: inline-flex;
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: #334155;
}
.team-leader img {
  width: 20px;
  height: 20px;
}
.team-description {
  color: #475569;
  margin-bottom: 1rem;
  line-height: 1.5;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
}
.team-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
  margin-top: auto;
  color: #64748b;
}
.team-meta-item {
  display: flex;
  align-items: center;
}
.team-meta-item img {
  width: 20px;
  height: 20px;
}
.join-button {
  width: 100%;
  padding: 0.8rem 0;
  border-radius: 8px;
  border: none;
  background-color: #3498db;
  color: white;
  font-weight: 500;
  cursor: pointer;
  margin-top: auto;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
}
.join-button:hover:not(:disabled) {
  background-color: #2980b9;
  transform: translateY(-1px);
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}
.join-button:active:not(:disabled) {
  transform: translateY(0);
}
.join-button:disabled {
  background-color: #cbd5e1;
  cursor: not-allowed;
}
</style>
