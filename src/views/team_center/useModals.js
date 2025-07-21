// composables/useModals.js

import { ref } from 'vue'

export function useModals(showToastMessage) {
  // “创建团队”弹窗的状态
  const showCreateModal = ref(false)

  // “申请加入”弹窗的状态和数据
  const showJoinModal = ref(false)
  const selectedTeam = ref(null)
  const joinReason = ref('')
  const contactInfo = ref('')

  // 打开“申请加入”弹窗的处理器
  const handleOpenJoinModal = (team) => {
    selectedTeam.value = team
    joinReason.value = ''
    contactInfo.value = ''
    showJoinModal.value = true
  }

  // 提交“申请加入”的处理器
  const submitJoinApplication = () => {
    // 在实际项目中，这里会是调用API的地方
    console.log(`正在为团队 "${selectedTeam.value.name}" 提交申请...`)

    showJoinModal.value = false
    showToastMessage(`已成功提交加入 "${selectedTeam.value.name}" 的申请`, 'success')
  }

  return {
    showCreateModal,
    showJoinModal,
    selectedTeam,
    joinReason,
    contactInfo,
    handleOpenJoinModal,
    submitJoinApplication,
  }
}
