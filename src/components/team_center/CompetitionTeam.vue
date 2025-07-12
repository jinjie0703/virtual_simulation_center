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
      <i class="fas fa-trophy"></i>
      <h3>没有找到相关竞赛团队</h3>
      <p>请尝试其他搜索条件或查看全部团队</p>
    </template>

    <template #team-card="{ team }">
      <div class="team-card-content">
        <div class="team-header">
          <h3>{{ team.name }}</h3>
          <span class="team-status" :class="team.status">{{ getStatusText(team.status) }}</span>
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
            <span class="team-meta-item"
              ><i class="fas fa-calendar-alt"></i> {{ team.deadline }}</span
            >
          </div>
        </div>

        <button class="join-button" @click="joinTeam(team)" :disabled="team.status === 'closed'">
          <i class="fas" :class="team.status === 'closed' ? 'fa-lock' : 'fa-sign-in-alt'"></i>
          {{ team.status === 'closed' ? '已截止' : '申请加入' }}
        </button>
      </div>
    </template>
  </TeamList>
</template>

<script>
import TeamList from './TeamList.vue'

export default {
  name: 'CompetitionTeam',
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
    getStatusText(status) {
      return status === 'open' ? '招募中' : '已截止'
    },
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
      if (!team.status) team.status = 'open'
      this.teams.unshift(team)
    },
    fetchTeams() {
      // 模拟API调用
      return new Promise((resolve) => {
        setTimeout(() => {
          resolve([
            {
              id: 1,
              name: '电子设计竞赛队伍 A',
              leader: '张明',
              description:
                '我们正在寻找对嵌入式系统和电路设计充满热情的成员，将参加全国大学生电子设计竞赛。',
              status: 'open',
              memberCount: 3,
              maxMembers: 5,
              deadline: '2023-12-30',
              tags: ['电子设计', '嵌入式系统'],
              createdAt: '2023-10-15',
            },
            {
              id: 2,
              name: 'ACM 编程竞赛队',
              leader: '李华',
              description:
                '寻找算法高手，一起备战 ACM-ICPC 国际大学生程序设计竞赛。需要有良好的数据结构和算法基础。',
              status: 'open',
              memberCount: 2,
              maxMembers: 3,
              deadline: '2023-12-15',
              tags: ['算法', '编程', 'ACM'],
              createdAt: '2023-10-18',
            },
            {
              id: 3,
              name: '机器人竞赛小组',
              leader: '王强',
              description:
                '准备参加全国机器人大赛，需要对机器视觉、ROS系统和机械控制有经验的同学。',
              status: 'open',
              memberCount: 4,
              maxMembers: 6,
              deadline: '2023-11-30',
              tags: ['机器人', '机械控制', '计算机视觉'],
              createdAt: '2023-10-12',
            },
            {
              id: 4,
              name: '互联网+创新创业大赛',
              leader: '陈雪',
              description:
                '寻找有创业想法和热情的同学，组建团队参加今年的"互联网+"大学生创新创业大赛。',
              status: 'open',
              memberCount: 3,
              maxMembers: 5,
              deadline: '2023-11-15',
              tags: ['创新创业', '商业计划', '产品设计'],
              createdAt: '2023-10-20',
            },
            {
              id: 5,
              name: '数学建模竞赛团队',
              leader: '赵宇',
              description:
                '寻找擅长数学建模和编程的同学，参加全国大学生数学建模竞赛，需要有MATLAB或Python经验。',
              status: 'open',
              memberCount: 2,
              maxMembers: 3,
              deadline: '2023-12-05',
              tags: ['数学建模', 'MATLAB', 'Python'],
              createdAt: '2023-10-22',
            },
            {
              id: 6,
              name: '大数据应用创新大赛',
              leader: '林小明',
              description:
                '寻找熟悉大数据处理技术和数据可视化的队友，一起参加全国大数据应用创新大赛。',
              status: 'closed',
              memberCount: 3,
              maxMembers: 4,
              deadline: '2023-10-10',
              tags: ['大数据', '数据分析', '可视化'],
              createdAt: '2023-10-01',
            },
            // 更多竞赛数据
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

.team-status {
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
