import { ref } from 'vue'

export function useMemberModal() {
  const selectedMember = ref(null)

  const showMemberDetails = (member) => {
    console.log('显示成员详情:', member)
    selectedMember.value = member
  }

  const closeMemberDetails = () => {
    console.log('关闭成员详情')
    selectedMember.value = null
  }

  return {
    selectedMember,
    showMemberDetails,
    closeMemberDetails,
  }
}
