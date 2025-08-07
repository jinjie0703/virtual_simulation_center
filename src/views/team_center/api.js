// src/views/team_center/api.js
import axios from 'axios'

// 创建一个axios实例，用于后续的API请求
const apiClient = axios.create({
  baseURL: 'https://localhost:8080/api', // 你的Go后端地址
  headers: {
    'Content-Type': 'application/json',
  },
})

/**
 * 从后端获取竞赛团队列表
 * @returns {Promise<Array>}
 */
export const fetchCompetitionTeams = async () => {
  try {
    const response = await apiClient.get('/team_center/competition_teams')
    return response.data
  } catch (error) {
    console.error('获取竞赛团队数据失败:', error)
    // 抛出错误，让调用者（useTeams.js）可以捕获
    throw error
  }
}

/**
 * 从后端获取项目团队列表
 * @returns {Promise<Array>}
 */
export const fetchProjectTeams = async () => {
  try {
    const response = await apiClient.get('/team_center/project_teams')
    return response.data
  } catch (error) {
    console.error('获取项目团队数据失败:', error)
    // 抛出错误，让调用者可以捕获
    throw error
  }
}
