<!-- TeamCenter.vue -->
<template>
  <header class="team-header">
    <div>
      <h1 class="title">团队中心</h1>
      <p class="subtitle">寻找志同道合的伙伴，一起组队参加竞赛或开展项目</p>
    </div>
  </header>
  <div class="team-center-wrapper">
    <div class="tabs-container">
      <div class="tabs">
        <button
          :class="['tab-button', { active: activeTab === 'competition' }]"
          @click="activeTab = 'competition'"
        >
          <img
            src="@/assets/team_center/competition.svg"
            alt="竞赛组队图标"
            class="competition-icon"
          />
          <span>竞赛组队</span>
        </button>
        <button
          :class="['tab-button', { active: activeTab === 'project' }]"
          @click="activeTab = 'project'"
        >
          <img src="@/assets/team_center/project.svg" alt="项目组队图标" class="project-icon" />
          <span>项目组队</span>
        </button>
      </div>
    </div>

    <div class="search-filter-container">
      <div class="search-box">
        <i class="fas fa-search"></i>
        <input v-model="searchQuery" type="text" placeholder="搜索团队..." />
      </div>

      <!-- 修改筛选组件区域，确保正确渲染 -->
      <div class="filter-options">
        <!-- 排序下拉框 -->
        <CustomSelect
          v-model="filterOption"
          :options="filterOptions"
          placeholder="选择排序方式"
          size="medium"
          :min-width="'140px'"
        />

        <!-- 难度下拉框 -->
        <CustomSelect
          v-model="difficultyFilter"
          :options="difficultyOptions"
          placeholder="选择难度"
          size="medium"
          :min-width="'120px'"
        />

        <!-- 标签下拉框 -->
        <CustomSelect
          v-model="tagFilter"
          :options="tagOptions"
          placeholder="选择标签"
          size="medium"
          :min-width="'120px'"
        />
      </div>

      <button class="create-team-button" @click="showCreateModal = true">
        <i class="fas fa-plus"></i> 创建团队
      </button>
    </div>

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

    <!-- 模态框和通知 (这部分HTML结构不变) -->
    <CreateTeamModal
      v-if="showCreateModal"
      :team-type="activeTab"
      @close="showCreateModal = false"
      @team-created="addTeam"
    />

    <transition name="modal">
      <div v-if="showJoinModal" class="join-modal-overlay" @click.self="showJoinModal = false">
        <div class="join-modal-content">
          <!-- ... 申请加入弹窗的内部结构 ... -->
          <div class="modal-header">
            <h3>申请加入团队</h3>
            <button class="close-button" @click="showJoinModal = false">×</button>
          </div>
          <div class="modal-body" v-if="selectedTeam">
            <p><strong>团队名称:</strong> {{ selectedTeam.name }}</p>
            <p>
              <strong>{{ activeTab === 'competition' ? '队长' : '负责人' }}:</strong>
              {{ selectedTeam.leader }}
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
              <button class="cancel-button" @click="showJoinModal = false">取消</button>
              <button
                class="submit-button"
                :disabled="!joinReason || !contactInfo"
                @click="submitJoinApplication"
              >
                提交申请
              </button>
            </div>
          </div>
        </div>
      </div>
    </transition>

    <transition name="toast">
      <div v-if="showToast" class="toast-message" :class="toastType">
        <i :class="toastIcon"></i> {{ toastMessage }}
      </div>
    </transition>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import TeamList from '@/components/team_center/TeamList.vue'
import CreateTeamModal from '@/components/team_center/CreateTeamModal.vue'
import Pagination from '@/components/team_center/TeamPagination.vue'
// 确保路径正确 - 可能需要根据实际目录结构调整
import CustomSelect from '@/components/common/CustomSelect.vue'

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
const { loading, filteredTeams, paginatedTeams, addTeam } = useTeams({
  activeTab,
  searchQuery,
  filterOption,
  difficultyFilter, // 新增筛选参数
  tagFilter, // 新增筛选参数
  currentPage,
  itemsPerPage,
  showToastMessage,
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
/* 标题样式 */
.team-header {
  padding: 120px 0px 80px 0px;
  text-align: center;
  position: relative;
  overflow: hidden;
  background: linear-gradient(135deg, #a1d3ff 0%, #2d92ff 100%);
}

.title {
  font-size: 3.5rem;
  font-weight: 800;
  margin: 0 0 15px;
  color: #2c3e50;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
  animation: fadeInDown 1s ease-out;
}

.subtitle {
  font-size: 1.25rem;
  max-width: 800px;
  color: gray;
  margin: 0 auto;
  opacity: 0.9;
  line-height: 1.6;
  animation: fadeInUp 1s ease-out 0.3s both;
}

@keyframes fadeInDown {
  from {
    opacity: 0;
    transform: translateY(-30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.team-center-wrapper {
  width: 100%;
  padding: 2rem;
  max-width: 1400px;
  margin: 0 auto;
  position: relative;
}

.tabs-container {
  display: flex;
  justify-content: center;
  margin-bottom: 2rem;
}
.tabs {
  display: flex;
  background-color: #f1f5f9;
  border-radius: 12px;
  padding: 0.4rem;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
}
.tab-button {
  display: inline-flex;
  align-items: center;
  gap: 12px;
  padding: 0.9rem 1.6rem;
  border: none;
  background-color: transparent;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 500;
  color: #64748b;
  transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
  border-radius: 8px;
  min-width: 140px;
}
.tab-button.active {
  background-color: #ffffff;
  color: #3498db;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.12);
  transform: translateY(-1px);
}
.tab-button:hover:not(.active) {
  color: #3498db;
  background-color: rgba(255, 255, 255, 0.5);
}
.tab-button i {
  margin-right: 0.5rem;
}
.competition-icon {
  width: 24px;
  height: 24px;
}
.project-icon {
  width: 24px;
  height: 24px;
}
.search-filter-container {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
  margin-bottom: 2rem;
  align-items: center;
}
.search-box {
  flex-grow: 1;
  position: relative;
  min-width: 250px;
}
.search-box i {
  position: absolute;
  left: 1rem;
  top: 50%;
  transform: translateY(-50%);
  color: #94a3b8;
}
.search-box input {
  width: 100%;
  padding: 0.9rem 2.5rem;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  font-size: 1rem;
  transition: all 0.3s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
}
.search-box input:focus {
  outline: none;
  border-color: #3498db;
  box-shadow: 0 0 0 3px rgba(52, 152, 219, 0.15);
}

/* 修改筛选选项样式，确保正确显示 */
.filter-options {
  display: flex;
  align-items: center;
  gap: 10px;
  flex-wrap: wrap;
  min-height: 40px; /* 确保有足够的高度 */
}

/* 添加一些选择器样式确保渲染正确 */
:deep(.custom-select) {
  margin-bottom: 4px;
  margin-top: 4px;
}

.create-team-button {
  padding: 0.9rem 1.6rem;
  background-color: #3498db;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.3s ease;
  white-space: nowrap;
  box-shadow: 0 2px 6px rgba(52, 152, 219, 0.4);
}
.create-team-button:hover {
  background-color: #2980b9;
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(52, 152, 219, 0.5);
}
.create-team-button:active {
  transform: translateY(0);
  box-shadow: 0 1px 3px rgba(52, 152, 219, 0.4);
}
.create-team-button i {
  margin-right: 0.5rem;
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
  h1 {
    font-size: 2rem;
  }
  .tabs {
    width: 100%;
  }
  .tab-button {
    min-width: 0;
    flex: 1;
    padding: 0.8rem 0.6rem;
    font-size: 0.9rem;
  }
  .search-filter-container {
    flex-direction: column;
    align-items: stretch;
  }
  .search-box,
  .filter-options,
  .create-team-button {
    width: 100%;
  }
  .create-team-button {
    order: -1;
    margin-bottom: 1rem;
  }
  .filter-options {
    width: 100%;
    justify-content: space-between;
    gap: 8px;
  }

  /* 在小屏幕上让每个选择器占满宽度 */
  :deep(.custom-select) {
    width: 100%;
  }
}
</style>
