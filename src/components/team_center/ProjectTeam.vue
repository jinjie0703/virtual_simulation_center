<template>
  <TeamList
    :teams="teams"
    :search-query="searchQuery"
    :filter-option="filterOption"
    :current-page="currentPage"
    :items-per-page="itemsPerPage"
    @join-team="joinTeam"
    @update-total-items="updateTotalItems"
    @clear-filters="clearFilters"
  >
    <template #empty-state>
      <i class="fas fa-project-diagram"></i>
      <h3>没有找到相关项目团队</h3>
      <p>请尝试其他搜索条件或查看全部团队</p>
    </template>

    <template #team-card="{ team }">
      <div class="team-card-content">
        <div class="team-header">
          <h3>{{ team.name }}</h3>
          <span class="team-difficulty" :class="team.difficulty">{{ team.difficulty }}</span>
        </div>

        <div class="team-tags">
          <span class="team-tag" v-for="(tag, index) in team.tags" :key="index">{{ tag }}</span>
        </div>

        <div class="team-info">
          <p class="team-leader"><i class="fas fa-user-tie"></i> {{ team.leader }}</p>
          <p class="team-description">{{ team.description }}</p>
          <div class="team-meta">
            <span class="team-meta-item"
              ><i class="fas fa-users"></i> {{ team.memberCount }}/{{ team.maxMembers }}</span
            >
            <span class="team-meta-item"><i class="fas fa-clock"></i> {{ team.duration }}</span>
          </div>
        </div>

        <button class="join-button" @click="joinTeam(team)">
          <i class="fas fa-sign-in-alt"></i> 申请加入
        </button>
      </div>
    </template>
  </TeamList>
</template>

<script>
import TeamList from './TeamList.vue'

export default {
  name: 'ProjectTeam',
  components: {
    TeamList,
  },
  props: {
    searchQuery: String,
    filterOption: String,
    currentPage: Number,
    itemsPerPage: Number,
  },
  data() {
    return {
      teams: [],
    }
  },
  methods: {
    joinTeam(team) {
      this.$emit('join-team', team)
    },
    updateTotalItems(total) {
      this.$emit('update-total-items', total)
    },
    clearFilters() {
      this.$emit('update:searchQuery', '')
      this.$emit('update:filterOption', 'all')
      this.$emit('clear-filters')
    },
    addTeam(team) {
      // 添加默认值
      if (!team.difficulty) team.difficulty = '中等'
      this.teams.unshift(team)
    },
    fetchTeams() {
      // 模拟API调用
      return new Promise((resolve) => {
        setTimeout(() => {
          resolve([
            {
              id: 1,
              name: '校园VR导览项目',
              leader: '李明',
              description: '开发一个校园虚拟现实导览系统，需要熟悉Unity 3D和VR开发的同学加入。',
              difficulty: '中等',
              memberCount: 2,
              maxMembers: 5,
              duration: '3个月',
              tags: ['VR/AR', 'Unity', '前端开发'],
              createdAt: '2023-10-20',
            },
            {
              id: 2,
              name: '智能家居控制系统',
              leader: '王华',
              description: '基于树莓派的智能家居控制系统开发，包括传感器集成和移动端控制App开发。',
              difficulty: '较难',
              memberCount: 3,
              maxMembers: 6,
              duration: '4个月',
              tags: ['嵌入式', 'IoT', '移动开发'],
              createdAt: '2023-10-15',
            },
            {
              id: 3,
              name: '教学资源共享平台',
              leader: '陈静',
              description:
                '开发一个面向校内师生的教学资源共享平台，实现资源上传、分类、评论等功能。',
              difficulty: '简单',
              memberCount: 2,
              maxMembers: 4,
              duration: '2个月',
              tags: ['Web开发', '全栈', '教育科技'],
              createdAt: '2023-10-25',
            },
            {
              id: 4,
              name: '校园社交App',
              leader: '张小东',
              description:
                '开发一款专为大学生设计的校园社交应用，实现兴趣小组、活动组织、匿名交流等功能。',
              difficulty: '中等',
              memberCount: 3,
              maxMembers: 5,
              duration: '5个月',
              tags: ['移动开发', 'UI设计', '后端开发'],
              createdAt: '2023-10-18',
            },
            {
              id: 5,
              name: '智能图书推荐系统',
              leader: '刘思思',
              description:
                '基于用户借阅历史和阅读偏好，开发一个智能图书推荐系统，需要有机器学习和数据分析经验。',
              difficulty: '较难',
              memberCount: 2,
              maxMembers: 4,
              duration: '4个月',
              tags: ['机器学习', '数据分析', 'Python'],
              createdAt: '2023-10-22',
            },
            {
              id: 6,
              name: '校园二手交易平台',
              leader: '吴佳',
              description:
                '开发一个便于校内学生进行二手物品交易的Web平台，包括物品发布、搜索、聊天等功能。',
              difficulty: '简单',
              memberCount: 2,
              maxMembers: 4,
              duration: '3个月',
              tags: ['Web开发', '前端', '后端'],
              createdAt: '2023-10-23',
            },
            // 更多项目数据
          ])
        }, 500)
      })
    },
  },
  async created() {
    this.teams = await this.fetchTeams()
  },
}
</script>

<style scoped>
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

.team-difficulty {
  font-size: 0.85rem;
  padding: 0.4rem 0.8rem;
  border-radius: 50px;
  font-weight: 500;
  white-space: nowrap;
  margin-left: 0.5rem;
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
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: #334155;
}

.team-leader i {
  margin-right: 0.4rem;
  color: #64748b;
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

.team-meta-item i {
  margin-right: 0.4rem;
  opacity: 0.7;
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

.join-button:hover {
  background-color: #2980b9;
  transform: translateY(-1px);
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

.join-button:active {
  transform: translateY(0);
}
</style>
