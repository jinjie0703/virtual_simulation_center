// composables/useToast.js

import { ref, computed, onUnmounted } from 'vue'

export function useToast() {
  const showToast = ref(false)
  const toastMessage = ref('')
  const toastType = ref('success')
  let toastTimeout = null

  const toastIcon = computed(
    () =>
      ({
        success: 'fas fa-check-circle',
        error: 'fas fa-exclamation-circle',
        info: 'fas fa-info-circle',
      })[toastType.value],
  )

  const showToastMessage = (message, type = 'info', duration = 3000) => {
    if (toastTimeout) clearTimeout(toastTimeout)

    toastMessage.value = message
    toastType.value = type
    showToast.value = true

    toastTimeout = setTimeout(() => {
      showToast.value = false
    }, duration)
  }

  // 组件卸载时清除定时器，防止内存泄漏
  onUnmounted(() => {
    if (toastTimeout) clearTimeout(toastTimeout)
  })

  return {
    showToast,
    toastMessage,
    toastType,
    toastIcon,
    showToastMessage,
  }
}
