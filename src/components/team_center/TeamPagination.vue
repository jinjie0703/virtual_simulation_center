<template>
  <div v-if="totalPages > 1" class="pagination-container">
    <button
      class="pagination-button"
      :disabled="currentPage === 1"
      @click="changePage(currentPage - 1)"
    >
      «
    </button>
    <button
      v-for="(page, index) in pages"
      :key="index"
      class="pagination-button"
      :class="{ active: currentPage === page, dots: page === '...' }"
      :disabled="page === '...'"
      @click="changePage(page)"
    >
      {{ page }}
    </button>
    <button
      class="pagination-button"
      :disabled="currentPage === totalPages"
      @click="changePage(currentPage + 1)"
    >
      »
    </button>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  currentPage: { type: Number, required: true },
  totalItems: { type: Number, required: true },
  itemsPerPage: { type: Number, required: true },
  maxPagesShown: { type: Number, default: 5 },
})

const emit = defineEmits(['page-changed'])

const totalPages = computed(() => Math.ceil(props.totalItems / props.itemsPerPage))

const pages = computed(() => {
  const pagesArray = []
  const { currentPage, maxPagesShown } = props
  const total = totalPages.value

  if (total <= maxPagesShown) {
    for (let i = 1; i <= total; i++) pagesArray.push(i)
    return pagesArray
  }

  pagesArray.push(1)
  let start = Math.max(2, currentPage - Math.floor((maxPagesShown - 3) / 2))
  let end = Math.min(total - 1, start + maxPagesShown - 3)

  if (currentPage > total - Math.floor((maxPagesShown - 1) / 2)) {
    start = total - maxPagesShown + 2
    end = total - 1
  }

  if (end - start < maxPagesShown - 3) {
    start = end - maxPagesShown + 3
  }

  if (start > 2) pagesArray.push('...')

  for (let i = start; i <= end; i++) pagesArray.push(i)

  if (end < total - 1) pagesArray.push('...')

  pagesArray.push(total)
  return pagesArray
})

const changePage = (page) => {
  if (page !== '...' && page > 0 && page <= totalPages.value && page !== props.currentPage) {
    emit('page-changed', page)
  }
}
</script>

<style scoped>
/* 样式与原文件保持一致 */
.pagination-container {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 2rem;
  user-select: none;
}
.pagination-button {
  margin: 0 0.25rem;
  padding: 0.6rem 1rem;
  border: 1px solid #dfe4ea;
  background-color: #fff;
  color: #34495e;
  cursor: pointer;
  border-radius: 6px;
  transition: all 0.3s ease;
  min-width: 40px;
  text-align: center;
}
.pagination-button:hover:not(:disabled):not(.active) {
  background-color: #f8f9fa;
  border-color: #3498db;
  color: #3498db;
}
.pagination-button:disabled:not(.dots) {
  cursor: not-allowed;
  opacity: 0.6;
}
.pagination-button.active {
  background-color: #3498db;
  color: white;
  border-color: #3498db;
  font-weight: bold;
}
.pagination-button.dots {
  cursor: default;
  border: none;
  background: transparent;
}
</style>
