// composables/useTeams.js

import { ref, computed, watch } from 'vue'
// 注意：这里我们假设你有一个api.js文件来导出获取数据的函数
// 即使你不提供它，这里的结构也是清晰的
import * as Api from './api'

export function useTeams({
  activeTab,
  searchQuery,
  difficultyFilter,
  tagFilter,
  currentPage,
  itemsPerPage,
  showToastMessage,
}) {
  const loading = ref(true)
  const allTeams = ref([])

  // 获取数据的方法
  const fetchData = async () => {
    loading.value = true
    try {
      const data =
        activeTab.value === 'competition'
          ? await Api.fetchCompetitionTeams() // 你的API调用
          : await Api.fetchProjectTeams() // 你的API调用
      allTeams.value = data.map((team) => ({
        ...team,
        contact: {
          wechat: team.contact_wechat,
          qq: team.contact_qq,
          email: team.contact_email,
        },
        deadline: team.deadline ? new Date(team.deadline).toLocaleDateString() : null,
        createdAt: new Date(team.createdAt).toLocaleDateString(),
      }))
    } catch (error) {
      console.error('获取团队数据失败:', error)
      showToastMessage('数据加载失败，请稍后重试', 'error')
      allTeams.value = []
    } finally {
      // 延迟结束加载，以获得更好的视觉体验
      setTimeout(() => {
        loading.value = false
      })
    }
  }

  // 监听Tab变化，自动重新获取数据
  watch(activeTab, () => {
    currentPage.value = 1 // 切换时重置到第一页
    fetchData()
  })

  // 核心计算属性：根据搜索和排序条件过滤数据
  const filteredTeams = computed(() => {
    let teams = [...allTeams.value]

    // 1. 文本搜索
    if (searchQuery.value) {
      const query = searchQuery.value.toLowerCase()
      teams = teams.filter(
        (team) =>
          team.name.toLowerCase().includes(query) ||
          team.description.toLowerCase().includes(query) ||
          team.leader.toLowerCase().includes(query) ||
          (team.tags && team.tags.some((tag) => tag.toLowerCase().includes(query))),
      )
    }

    // 2. 难度过滤
    if (difficultyFilter.value && difficultyFilter.value !== '') {
      teams = teams.filter((team) => team.difficulty === difficultyFilter.value)
    }

    // 3. 标签过滤
    if (tagFilter.value && tagFilter.value !== '') {
      teams = teams.filter((team) => team.tags.includes(tagFilter.value))
    }

    return teams
  })

  // 监听筛选条件变化，自动重置到第一页
  watch([searchQuery, difficultyFilter, tagFilter], () => {
    currentPage.value = 1
  })

  // 计算当前页应显示的数据
  const paginatedTeams = computed(() => {
    const start = (currentPage.value - 1) * itemsPerPage.value
    const end = start + itemsPerPage.value
    return filteredTeams.value.slice(start, end)
  })

  // 在列表头部添加一个新创建的团队
  const addTeam = (newTeam) => {
    allTeams.value.unshift(newTeam)
    showToastMessage('团队创建成功！', 'success')
  }

  // 组件挂载时首次获取数据
  fetchData()

  const dynamicDifficultyOptions = computed(() => {
    const difficulties = [...new Set(allTeams.value.map((team) => team.difficulty).filter(Boolean))]
    return [{ value: '', label: '全部难度' }, ...difficulties.map((d) => ({ value: d, label: d }))]
  })

  const dynamicTagOptions = computed(() => {
    const tags = [...new Set(allTeams.value.flatMap((team) => team.tags).filter(Boolean))]
    return [{ value: '', label: '全部标签' }, ...tags.map((t) => ({ value: t, label: t }))]
  })

  // 返回所有组件需要用到的状态和方法
  return {
    loading,
    filteredTeams, // 用于计算总数
    paginatedTeams, // 用于列表渲染
    addTeam,
    fetchData, // 导出 fetchData 以便手动刷新
    dynamicDifficultyOptions,
    dynamicTagOptions,
  }
}
