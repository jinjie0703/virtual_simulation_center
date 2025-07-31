import { ref, nextTick } from 'vue'

export function useTabSwitcher(initialTab = 'teachers') {
  const activeTab = ref(initialTab)
  const isLoading = ref(false)

  const switchTab = async (tab) => {
    if (tab === activeTab.value) return

    isLoading.value = true
    activeTab.value = tab

    // 模拟加载延迟
    await nextTick()
    setTimeout(() => {
      isLoading.value = false
    }, 300)
  }

  return {
    activeTab,
    isLoading,
    switchTab,
  }
}
