<!-- TeamCenter.vue -->
<template>
  <TeamHeader />
  <div class="team-center-wrapper">
    <TeamActions
      v-model:activeTab="activeTab"
      v-model:searchQuery="searchQuery"
      v-model:difficultyFilter="difficultyFilter"
      v-model:tagFilter="tagFilter"
      :difficulty-options="dynamicDifficultyOptions"
      :tag-options="dynamicTagOptions"
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
const currentPage = ref(1)
const itemsPerPage = ref(20)
// 新增筛选状态
const difficultyFilter = ref('')
const tagFilter = ref('')



// 增加对activeTab的监听，当切换标签页时重置筛选选项
watch(activeTab, (newTab, oldTab) => {
  if (newTab === oldTab) return

  // 切换标签页时重置所有筛选
  searchQuery.value = ''
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
const {
  loading,
  filteredTeams,
  paginatedTeams,
  addTeam,
  fetchData,
  dynamicDifficultyOptions,
  dynamicTagOptions,
} = useTeams({
  activeTab,
  searchQuery,
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
