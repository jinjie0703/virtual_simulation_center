import { ref, computed } from 'vue'
import axios from 'axios'

export function useTeamData() {
  const members = ref([])
  const isLoading = ref(true)
  const error = ref(null)

  const fetchTeamData = async () => {
    isLoading.value = true
    error.value = null
    try {
      const response = await axios.get(`${import.meta.env.VITE_API_BASE_URL}/api/our_team/team_members`)
      members.value = response.data
    } catch (err) {
      error.value = 'Failed to fetch members'
      console.error(err)
    } finally {
      isLoading.value = false
    }
  }

  const teachers = computed(() => members.value.filter((m) => m.role === 'teacher'))
  const students = computed(() => members.value.filter((m) => m.role === 'student'))

  // 标签配置
  const tabsConfig = computed(() => [
    {
      id: 'teachers',
      name: '教师团队',
    },
    {
      id: 'students',
      name: '学生团队',
    },
  ])

  // 根据类型获取成员数据
  const getMembersByType = (type) => {
    return type === 'teachers' ? teachers.value : students.value
  }

  return {
    teachers,
    students,
    tabsConfig,
    getMembersByType,
    isLoading,
    error,
    fetchTeamData,
  }
}
