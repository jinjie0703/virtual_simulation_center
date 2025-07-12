<template>
  <div class="team-center-wrapper">
    <div class="page-header">
      <h1>团队中心</h1>
      <p class="subtitle">寻找志同道合的伙伴，一起组队参加竞赛或开展项目</p>
    </div>

    <div class="tabs-container">
      <div class="tabs">
        <button
          :class="['tab-button', { active: activeTab === 'competition' }]"
          @click="changeTab('competition')"
        >
          <i class="fas fa-trophy"></i> 竞赛组队
        </button>
        <button
          :class="['tab-button', { active: activeTab === 'project' }]"
          @click="changeTab('project')"
        >
          <i class="fas fa-project-diagram"></i> 项目组队
        </button>
      </div>
    </div>

    <div class="search-filter-container">
      <div class="search-box">
        <i class="fas fa-search"></i>
        <input
          type="text"
          v-model="searchQuery"
          placeholder="搜索团队..."
          @input="debounceSearch"
        />
      </div>
      <div class="filter-options">
        <select v-model="filterOption" @change="handleFilterChange">
          <option value="all">全部</option>
          <option value="recent">最近发布</option>
          <option value="popular">最受欢迎</option>
        </select>
      </div>
      <button class="create-team-button" @click="showCreateModal = true">
        <i class="fas fa-plus"></i> 创建团队
      </button>
    </div>

    <div class="tab-content">
      <transition name="fade" mode="out-in">
        <keep-alive>
          <component
            :is="activeComponent"
            ref="activeTeamListComponent"
            :search-query="searchQuery"
            :filter-option="filterOption"
            :current-page="currentPage"
            :items-per-page="itemsPerPage"
            @join-team="handleJoinTeam"
            @update-total-items="updateTotalItems"
            @clear-filters="clearFilters"
          ></component>
        </keep-alive>
      </transition>
    </div>

    <Pagination
      :current-page="currentPage"
      :total-items="totalItems"
      :items-per-page="itemsPerPage"
      @page-changed="handlePageChange"
    />

    <!-- 申请加入弹窗 -->
    <transition name="modal">
      <div v-if="showJoinModal" class="join-modal-overlay" @click.self="showJoinModal = false">
        <div class="join-modal-content">
          <div class="modal-header">
            <h3>申请加入团队</h3>
            <button class="close-button" @click="showJoinModal = false">&times;</button>
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

    <!-- 创建团队弹窗 -->
    <CreateTeamModal
      v-if="showCreateModal"
      :team-type="activeTab"
      @close="showCreateModal = false"
      @team-created="handleTeamCreated"
    />

    <!-- 通知提示 -->
    <transition name="toast">
      <div v-if="showToast" class="toast-message" :class="toastType">
        <i :class="toastIcon"></i>
        {{ toastMessage }}
      </div>
    </transition>
  </div>
</template>

<script>
import CompetitionTeam from '@/components/team_center/CompetitionTeam.vue'
import ProjectTeam from '@/components/team_center/ProjectTeam.vue'
import CreateTeamModal from '@/components/team_center/CreateTeamModal.vue'
import Pagination from '@/components/team_center/Pagination.vue'
import _ from 'lodash'

export default {
  name: 'TeamCenter',
  components: {
    CompetitionTeam,
    ProjectTeam,
    CreateTeamModal,
    Pagination,
  },
  data() {
    return {
      activeTab: 'competition',
      searchQuery: '',
      filterOption: 'all',
      currentPage: 1,
      itemsPerPage: 15, // 3列 x 5行
      totalItems: 0,
      showJoinModal: false,
      showCreateModal: false,
      selectedTeam: null,
      joinReason: '',
      contactInfo: '',
      showToast: false,
      toastMessage: '',
      toastType: 'success',
      toastTimeout: null,
    }
  },
  computed: {
    activeComponent() {
      return this.activeTab === 'competition' ? 'CompetitionTeam' : 'ProjectTeam'
    },
    toastIcon() {
      return {
        success: 'fas fa-check-circle',
        error: 'fas fa-exclamation-circle',
        info: 'fas fa-info-circle',
      }[this.toastType]
    },
  },
  methods: {
    changeTab(tab) {
      if (this.activeTab === tab) return
      this.activeTab = tab
      this.clearFiltersAndResetPage()
    },
    // 使用 lodash 的 debounce 函数防止搜索过于频繁
    debounceSearch: _.debounce(function () {
      this.currentPage = 1
    }, 300),
    handleFilterChange() {
      this.currentPage = 1
    },
    handlePageChange(page) {
      this.currentPage = page
      window.scrollTo({ top: 0, behavior: 'smooth' }) // 平滑滚动到顶部
    },
    updateTotalItems(total) {
      this.totalItems = total
    },
    handleJoinTeam(team) {
      this.selectedTeam = team
      this.joinReason = ''
      this.contactInfo = ''
      this.showJoinModal = true
    },
    clearFilters() {
      this.searchQuery = ''
      this.filterOption = 'all'
      this.currentPage = 1
    },
    clearFiltersAndResetPage() {
      this.searchQuery = ''
      this.filterOption = 'all'
      this.currentPage = 1
    },
    handleTeamCreated(newTeam) {
      // 确保组件已经挂载
      this.$nextTick(() => {
        if (this.$refs.activeTeamListComponent) {
          this.$refs.activeTeamListComponent.addTeam(newTeam)
          this.showToastMessage('团队创建成功！', 'success')
        }
      })
    },
    submitJoinApplication() {
      // 这里可以添加实际的提交逻辑
      this.showJoinModal = false
      this.showToastMessage(
        `已成功提交加入 "${this.selectedTeam.name}" 的申请，请等待审核`,
        'success',
      )
      this.joinReason = ''
      this.contactInfo = ''
    },
    showToastMessage(message, type = 'info') {
      if (this.toastTimeout) {
        clearTimeout(this.toastTimeout)
      }

      this.toastMessage = message
      this.toastType = type
      this.showToast = true

      this.toastTimeout = setTimeout(() => {
        this.showToast = false
      }, 3000)
    },
  },
  beforeUnmount() {
    if (this.toastTimeout) {
      clearTimeout(this.toastTimeout)
    }
  },
}
</script>

<style scoped>
.team-center-wrapper {
  padding: 2rem;
  max-width: 1200px;
  margin: 0 auto;
  position: relative;
}

.page-header {
  text-align: center;
  margin-bottom: 2.5rem;
  animation: fadeIn 0.6s ease-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

h1 {
  font-size: 2.5rem;
  font-weight: 700;
  color: #2c3e50;
  margin-bottom: 0.5rem;
}

.subtitle {
  font-size: 1.1rem;
  color: #7f8c8d;
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

.clear-search {
  position: absolute;
  right: 1rem;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  color: #94a3b8;
  cursor: pointer;
  padding: 5px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s ease;
}

.clear-search:hover {
  background-color: #f1f5f9;
  color: #64748b;
}

.filter-options select {
  padding: 0.9rem 2.5rem 0.9rem 1.2rem;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  font-size: 1rem;
  background-color: white;
  cursor: pointer;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
  appearance: none;
  background-image: url("data:image/svg+xml;charset=UTF-8,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='none' stroke='%2364748b' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'%3e%3cpolyline points='6 9 12 15 18 9'%3e%3c/polyline%3e%3c/svg%3e");
  background-repeat: no-repeat;
  background-position: right 1rem center;
  background-size: 1em;
  transition: all 0.3s ease;
}

.filter-options select:focus {
  outline: none;
  border-color: #3498db;
  box-shadow: 0 0 0 3px rgba(52, 152, 219, 0.15);
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

/* 淡入淡出动画 */
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

/* 模态框动画 */
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

/* 申请加入模态框 */
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

/* Toast 消息通知 */
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

/* 响应式设计 */
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

  .filter-options select {
    width: 100%;
  }
}
</style>
