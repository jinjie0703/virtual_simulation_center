<template>
  <header class="ourTeam-header">
    <div>
      <h1 class="title">师资队伍</h1>
      <p class="subtitle">汇聚顶尖人才，打造创新团队</p>
    </div>
  </header>
  <div class="our-team-container">
    <!-- 标签切换 -->
    <TeamTabs
      v-if="tabsConfig && tabsConfig.length > 0"
      :tabs="tabsConfig"
      :activeTab="activeTab"
      :isLoading="isLoading"
      :searchKeyword="searchKeyword"
      :selectedFilter="selectedFilter"
      @switchTab="switchTab"
      @search="handleSearch"
      @filter="handleFilter"
    />
    <!-- 成员展示区域 -->
    <div class="members-section">
      <LoadingState v-if="isLoading" />

      <div v-else-if="paginatedMembers && paginatedMembers.length > 0" class="members-content">
        <MemberGrid
          :members="paginatedMembers"
          :activeTab="activeTab"
          @memberClick="showMemberDetails"
        />

        <!-- 分页组件 -->
        <TeamPagination
          v-if="totalPages > 1"
          :currentPage="currentPage"
          :totalPages="totalPages"
          @changePage="changePage"
        />
      </div>

      <EmptyState v-else />
    </div>

    <MemberDetailModal
      :member="selectedMember"
      :show="!!selectedMember"
      @close="closeMemberDetails"
    />
  </div>
</template>

<script setup>
import { computed, watch, nextTick, onMounted } from 'vue'
import TeamTabs from '@/components/our_team/TeamTabs.vue'
import TeamPagination from '@/components/our_team/TeamPagination.vue'
import MemberGrid from '@/components/our_team/MemberGrid.vue'
import MemberDetailModal from '@/components/our_team/MemberDetailModal.vue'
import LoadingState from '@/components/our_team/LoadingState.vue'
import EmptyState from '@/components/our_team/EmptyState.vue'
import { useTeamData } from '@/views/our_team/useTeamData.js'
import { useTabSwitcher } from '@/views/our_team/useTabSwitcher.js'
import { usePagination } from '@/views/our_team/usePagination.js'
import { ref } from 'vue'

console.log('=== 开始加载 OurTeam 组件 ===')

// 使用组合式函数
const { tabsConfig, getMembersByType } = useTeamData()
console.log('useTeamData 结果:', { tabsConfig: tabsConfig.value, getMembersByType })

const { activeTab, isLoading, switchTab } = useTabSwitcher()
console.log('useTabSwitcher 结果:', { activeTab: activeTab.value, isLoading: isLoading.value })

// 成员详情模态框状态
const selectedMember = ref(null)

// 搜索和筛选状态
const searchKeyword = ref('')
const selectedFilter = ref('all')

// 显示成员详情
const showMemberDetails = (member) => {
  console.log('显示成员详情:', member)
  selectedMember.value = member
}

// 关闭成员详情
const closeMemberDetails = () => {
  console.log('关闭成员详情')
  selectedMember.value = null
}

// 搜索处理函数
const handleSearch = (keyword) => {
  searchKeyword.value = keyword
  resetPage()
}

// 筛选处理函数
const handleFilter = (filterType) => {
  selectedFilter.value = filterType
  resetPage()
}

// 计算当前成员列表
const currentMembers = computed(() => {
  try {
    let members = getMembersByType(activeTab.value)

    // 应用搜索筛选
    if (searchKeyword.value.trim()) {
      const keyword = searchKeyword.value.toLowerCase().trim()
      members = members.filter(
        (member) =>
          member.name.toLowerCase().includes(keyword) ||
          member.position.toLowerCase().includes(keyword) ||
          member.department.toLowerCase().includes(keyword),
      )
    }

    // 应用排序筛选
    if (selectedFilter.value === 'name') {
      members = [...members].sort((a, b) => a.name.localeCompare(b.name))
    } else if (selectedFilter.value === 'position') {
      members = [...members].sort((a, b) => a.position.localeCompare(b.position))
    } else if (selectedFilter.value === 'department') {
      members = [...members].sort((a, b) => a.department.localeCompare(b.department))
    }

    console.log('计算当前成员:', { activeTab: activeTab.value, members })
    return members
  } catch (error) {
    console.error('计算当前成员出错:', error)
    return []
  }
})

// 分页功能
const {
  currentPage,
  totalPages,
  paginatedItems: paginatedMembers,
  changePage: pageChange,
  resetPage,
} = usePagination(currentMembers, 12)

// 页面切换功能
const changePage = (page) => {
  pageChange(page)
  nextTick(() => {
    window.scrollTo({ top: 300, behavior: 'smooth' })
  })
}

// 监听切换标签时重置页码和搜索
watch(activeTab, () => {
  resetPage()
  searchKeyword.value = ''
  selectedFilter.value = 'all'
})

// 组件挂载时的调试
onMounted(() => {
  console.log('=== OurTeam 组件调试信息 ===')
  console.log('组件已挂载')
  console.log('标签配置:', tabsConfig.value)
  console.log('当前活动标签:', activeTab.value)
  console.log('当前成员:', currentMembers.value)
  console.log('分页成员:', paginatedMembers.value)
  console.log('总页数:', totalPages.value)
  console.log('========================')
})
</script>

<style scoped>
.ourTeam-header {
  padding: 120px 0px 80px 0px;
  text-align: center;
  position: relative;
  overflow: hidden;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
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
.our-team-container {
  margin: 0 auto;
  padding: 60px 40px;
  font-family: 'Helvetica Neue', Arial, sans-serif;
  background-color: #f8f9fa;
  min-height: 100vh;
  width: 100%;
}

.members-section {
  min-height: 800px;
  position: relative;
  padding: 0 8%;
}
/* 响应式设计 */
@media (max-width: 800px) {
  .our-team-container {
    padding: 40px 20px;
  }

  .page-header {
    padding: 60px 20px 40px 20px;
  }

  .title {
    font-size: 2.8rem;
  }

  .title-divider {
    width: 60px;
    margin-bottom: 20px;
  }
}

@media (max-width: 600px) {
  .our-team-container {
    padding: 30px 15px;
  }

  .title {
    font-size: 2.4rem;
  }

  .subtitle {
    font-size: 1.1rem;
  }
}

@media (max-width: 1400px) {
  .members-section {
    padding: 0 6%;
  }
}

@media (max-width: 1100px) {
  .members-section {
    padding: 0 4%;
  }
}

@media (max-width: 800px) {
  .members-section {
    padding: 0 3%;
  }
}

@media (max-width: 600px) {
  .members-section {
    padding: 0 2%;
  }
}
</style>
