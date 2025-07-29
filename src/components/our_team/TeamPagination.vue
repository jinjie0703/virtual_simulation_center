<template>
  <div class="pagination-container">
    <button
      @click="changePage(currentPage - 1)"
      :disabled="currentPage === 1"
      class="pagination-arrow"
    >
      &lt;
    </button>
    <div class="page-numbers">
      <button
        v-for="page in pages"
        :key="page"
        @click="changePage(page)"
        :class="{ active: currentPage === page, dots: page === '...' }"
        :disabled="page === '...'"
        class="page-number"
      >
        {{ page }}
      </button>
    </div>
    <button
      @click="changePage(currentPage + 1)"
      :disabled="currentPage === totalPages"
      class="pagination-arrow"
    >
      &gt;
    </button>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  currentPage: {
    type: Number,
    required: true,
  },
  totalPages: {
    type: Number,
    required: true,
  },
})

const emit = defineEmits(['changePage'])

const pages = computed(() => {
  const range = []
  const delta = 2
  const left = props.currentPage - delta
  const right = props.currentPage + delta + 1
  let l

  for (let i = 1; i <= props.totalPages; i++) {
    if (i === 1 || i === props.totalPages || (i >= left && i < right)) {
      range.push(i)
    }
  }

  const withDots = []
  for (const i of range) {
    if (l) {
      if (i - l === 2) {
        withDots.push(l + 1)
      } else if (i - l !== 1) {
        withDots.push('...')
      }
    }
    withDots.push(i)
    l = i
  }

  return withDots
})

const changePage = (page) => {
  if (page > 0 && page <= props.totalPages) {
    emit('changePage', page)
  }
}
</script>

<style scoped>
.pagination-container {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20px 0;
  margin-top: 30px;
}

.page-numbers {
  display: flex;
  gap: 8px;
  margin: 0 15px;
}

.page-number,
.pagination-arrow {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  border: 1px solid #ddd;
  background-color: #fff;
  color: #333;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

.pagination-arrow:disabled {
  cursor: not-allowed;
  opacity: 0.5;
}

.page-number:hover,
.pagination-arrow:hover:not(:disabled) {
  background-color: #4a90e2;
  color: #fff;
  border-color: #4a90e2;
  transform: translateY(-2px);
}

.page-number.active {
  background-color: #4a90e2;
  color: #fff;
  border-color: #4a90e2;
  font-weight: bold;
}

.page-number.dots {
  background-color: transparent;
  border: none;
  cursor: default;
}
.page-number.dots:hover {
  background-color: transparent;
  transform: none;
}
</style>
