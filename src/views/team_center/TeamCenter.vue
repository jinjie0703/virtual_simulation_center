<!-- TeamCenter.vue -->
<template>
  <TeamHeader />
  <div class="team-center-wrapper">
    <TeamActions
      v-model:activeTab="activeTab"
      v-model:searchQuery="searchQuery"
      v-model:filterOption="filterOption"
      v-model:difficultyFilter="difficultyFilter"
      v-model:tagFilter="tagFilter"
      :filter-options="filterOptions"
      :difficulty-options="difficultyOptions"
      :tag-options="tagOptions"
      @create-team="showCreateModal = true"
    />

    <!-- 团队列表 -->
    <div class="tab-content">
      <transition name="fade" mode="out-in">
        <TeamList
          :loading="loading"
          :teams="paginatedTeams"
          @join-team="handleOpenJoinModal"
          @clear-filters="clearFilters"
        >
          <template #empty-state>
            <i
              :class="activeTab === 'competition' ? 'fas fa-trophy' : 'fas fa-project-diagram'"
            ></i>
            <h3>没有找到相关{{ activeTab === 'competition' ? '竞赛' : '项目' }}团队</h3>
            <p>请尝试其他搜索条件或查看全部团队</p>
          </template>
        </TeamList>
      </transition>
    </div>

    <!-- 分页 -->
    <Pagination
      :current-page="currentPage"
      :total-items="filteredTeams.length"
      :items-per-page="itemsPerPage"
      @page-changed="handlePageChange"
    />

    <!-- 模态框和通知 -->
    <CreateTeamModal
      v-if="showCreateModal"
      :team-type="activeTab"
      @close="showCreateModal = false"
      @team-created="addTeam"
    />

    <JoinTeamModal
      :show="showJoinModal"
      :team="selectedTeam"
      :team-type="activeTab"
      @close="showJoinModal = false"
      @submit="submitJoinApplication"
    />

    <transition name="toast">
      <div v-if="showToast" class="toast-message" :class="toastType">
        <i :class="toastIcon"></i> {{ toastMessage }}
      </div>
    </transition>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import TeamHeader from '@/components/team_center/TeamHeader.vue'
import TeamActions from '@/components/team_center/TeamActions.vue'
import TeamList from '@/components/team_center/TeamList.vue'
import CreateTeamModal from '@/components/team_center/CreateTeamModal.vue'
import JoinTeamModal from '@/components/team_center/JoinTeamModal.vue'
import Pagination from '@/components/team_center/TeamPagination.vue'

// 1. 导入你的组合式函数
import { useToast } from './useToast.js'
import { useModals } from './useModals.js'
import { useTeams } from './useTeams.js'

// 2. 定义组件顶层的、驱动UI交互的状态
const activeTab = ref('competition')
const searchQuery = ref('')
const filterOption = ref('') // 默认无排序，UI显示placeholder
const currentPage = ref(1)
const itemsPerPage = ref(20)
// 新增筛选状态
const difficultyFilter = ref('')
const tagFilter = ref('')

// 修改：使用计算属性定义筛选选项
const filterOptions = computed(() => {
  const options = [
    { value: 'newest', label: '最新创建' },
    { value: 'members_desc', label: '人数最多' },
    { value: 'members_asc', label: '急需招人' },
  ]
  // 只有竞赛组队有“临近截止”选项
  if (activeTab.value === 'competition') {
    options.splice(1, 0, { value: 'deadline', label: '临近截止' })
  }
  return options
})

// 新增：难度选项
const difficultyOptions = computed(() => [
  { value: '', label: '全部难度' },
  { value: 'easy', label: '简单' },
  { value: 'medium', label: '中等' },
  { value: 'hard', label: '困难' },
  { value: 'expert', label: '专家' },
])

// 新增：标签选项
const tagOptions = computed(() => {
  // 根据activeTab返回不同的标签选项
  if (activeTab.value === 'competition') {
    return [
      { value: '', label: '全部标签' },
      { value: 'ai', label: '人工智能' },
      { value: 'algorithm', label: '算法' },
      { value: 'data', label: '数据分析' },
      { value: 'security', label: '网络安全' },
      { value: 'innovation', label: '创新设计' },
    ]
  } else {
    return [
      { value: '', label: '全部标签' },
      { value: 'web', label: '网站开发' },
      { value: 'mobile', label: '移动应用' },
      { value: 'game', label: '游戏开发' },
      { value: 'iot', label: '物联网' },
      { value: 'research', label: '研究项目' },
    ]
  }
})

// 增加对activeTab的监听，当切换标签页时重置筛选选项
watch(activeTab, (newTab, oldTab) => {
  if (newTab === oldTab) return

  // 切换标签页时重置所有筛选
  searchQuery.value = ''
  filterOption.value = ''
  difficultyFilter.value = ''
  tagFilter.value = ''
  currentPage.value = 1
})

// 3. 实例化组合式函数，建立逻辑联系
const { showToast, toastMessage, toastType, toastIcon, showToastMessage } = useToast()
const {
  showCreateModal,
  showJoinModal,
  selectedTeam,
  joinReason,
  contactInfo,
  handleOpenJoinModal,
  submitJoinApplication,
} = useModals(showToastMessage)
const { loading, filteredTeams, paginatedTeams, addTeam, fetchData } = useTeams({
  activeTab,
  searchQuery,
  filterOption,
  difficultyFilter, // 新增筛选参数
  tagFilter, // 新增筛选参数
  currentPage,
  itemsPerPage,
  showToastMessage,
})

// 强制刷新数据以解决HMR缓存问题
onMounted(() => {
  fetchData()
})

// 4. 定义仅属于此组件的、简单的事件处理器
const handlePageChange = (page) => {
  currentPage.value = page
  window.scrollTo({ top: 300, behavior: 'smooth' })
}

const clearFilters = () => {
  searchQuery.value = ''
  filterOption.value = ''
  difficultyFilter.value = ''
  tagFilter.value = ''
}
</script>

<style scoped>
.team-center-wrapper {
  width: 100%;
  padding: 2rem;
  max-width: 1400px;
  margin: 0 auto;
  position: relative;
}

.fade-enter-active,
.fade-leave-active {
  transition:
    opacity 0.3s ease,
    transform 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(10px);
}

.tab-content {
  width: 100%;
}

.toast-message {
  position: fixed;
  bottom: 2rem;
  left: 50%;
  transform: translateX(-50%);
  padding: 1rem 1.5rem;
  background-color: white;
  color: #1e293b;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  z-index: 1100;
  display: flex;
  align-items: center;
  min-width: 300px;
  max-width: 90%;
}

.toast-message i {
  margin-right: 0.8rem;
  font-size: 1.2rem;
}

.toast-message.success {
  border-left: 4px solid #10b981;
}

.toast-message.success i {
  color: #10b981;
}

.toast-message.error {
  border-left: 4px solid #ef4444;
}

.toast-message.error i {
  color: #ef4444;
}

.toast-message.info {
  border-left: 4px solid #3b82f6;
}

.toast-message.info i {
  color: #3b82f6;
}

.toast-enter-active,
.toast-leave-active {
  transition: all 0.3s cubic-bezier(0.68, -0.55, 0.27, 1.55);
}

.toast-enter-from,
.toast-leave-to {
  opacity: 0;
  transform: translate(-50%, 20px);
}

@media (max-width: 768px) {
  .team-center-wrapper {
    padding: 1.5rem;
  }
}
</style>
