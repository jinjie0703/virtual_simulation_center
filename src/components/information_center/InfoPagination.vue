<template>
  <div class="pagination">
    <button
      @click="changePage(currentPage - 1)"
      :disabled="currentPage === 1"
      class="page-btn prev"
    >
      &lt;
    </button>
    <div class="page-numbers">
      <button
        v-for="page in visiblePages"
        :key="page"
        @click="changePage(page)"
        :class="['page-num', { active: currentPage === page, ellipsis: page === '...' }]"
        :disabled="page === '...'"
      >
        {{ page }}
      </button>
    </div>
    <button
      @click="changePage(currentPage + 1)"
      :disabled="currentPage === totalPages"
      class="page-btn next"
    >
      &gt;
    </button>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  currentPage: Number,
  totalPages: Number,
})

const emit = defineEmits(['page-changed'])

const changePage = (page) => {
  if (typeof page === 'number' && page >= 1 && page <= props.totalPages) {
    emit('page-changed', page)
  }
}

const visiblePages = computed(() => {
  const total = props.totalPages
  const current = props.currentPage
  const maxVisible = 5

  if (total <= maxVisible + 2) {
    return Array.from({ length: total }, (_, i) => i + 1)
  }

  const pages = []
  if (current < maxVisible) {
    for (let i = 1; i <= maxVisible; i++) pages.push(i)
    pages.push('...')
    pages.push(total)
  } else if (current > total - maxVisible + 1) {
    pages.push(1)
    pages.push('...')
    for (let i = total - maxVisible + 1; i <= total; i++) pages.push(i)
  } else {
    pages.push(1)
    pages.push('...')
    for (let i = current - 1; i <= current + 1; i++) pages.push(i)
    pages.push('...')
    pages.push(total)
  }
  return pages
})
</script>

<style scoped>
.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 10px;
  margin-top: 40px;
}
.page-btn,
.page-num {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  border: 1px solid #e2e8f0;
  background: #fff;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s ease;
  font-weight: 600;
  color: #4a5568;
}
.page-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}
.page-btn:not(:disabled):hover,
.page-num:not(.active):not(.ellipsis):hover {
  border-color: #4c51bf;
  color: #4c51bf;
  transform: translateY(-2px);
}
.page-numbers {
  display: flex;
  gap: 8px;
}
.page-num.active {
  background-color: #4c51bf;
  color: #fff;
  border-color: #4c51bf;
}
.page-num.ellipsis {
  border: none;
  background: transparent;
  cursor: default;
}
</style>
