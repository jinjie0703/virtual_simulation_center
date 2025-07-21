// src/composables/useScroll.js

/**
 * 提供滚动相关功能的 Composable 函数
 */
export function useScroll() {
  /**
   * 将窗口滚动到页面顶部
   * @param {'smooth' | 'auto'} behavior 滚动行为, 'smooth'为平滑, 'auto'为瞬间
   */
  const scrollToTop = (behavior = 'smooth') => {
    window.scrollTo({
      top: 0,
      behavior: behavior,
    })
  }

  // 未来你还可以添加滚动到指定元素等功能
  // const scrollToElement = (elementId) => { ... };

  // 将这个功能返回，以便其他组件可以使用
  return {
    scrollToTop,
    // scrollToElement,
  }
}
