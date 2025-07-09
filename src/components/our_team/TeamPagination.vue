<template>
  <div class="pagination">
    <button
      class="page-btn prev-btn"
      :disabled="currentPage === 1"
      @click="handlePageChange(currentPage - 1)"
      aria-label="上一页"
    >
      <span class="btn-icon">‹</span>
      <span class="btn-text">上一页</span>
    </button>

    <div class="page-numbers">
      <button
        v-for="page in visiblePages"
        :key="page"
        :class="['page-btn', 'page-num', { active: page === currentPage }]"
        @click="handlePageChange(page)"
        :disabled="page === '...'"
        :aria-label="page === '...' ? '省略页码' : `第${page}页`"
      >
        {{ page }}
      </button>
    </div>

    <button
      class="page-btn next-btn"
      :disabled="currentPage === totalPages"
      @click="handlePageChange(currentPage + 1)"
      aria-label="下一页"
    >
      <span class="btn-text">下一页</span>
      <span class="btn-icon">›</span>
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

const handlePageChange = (page) => {
  if (page !== '...' && page >= 1 && page <= props.totalPages && page !== props.currentPage) {
    emit('changePage', page)
  }
}

const visiblePages = computed(() => {
  const pages = []
  const total = props.totalPages
  const current = props.currentPage

  if (total <= 7) {
    for (let i = 1; i <= total; i++) {
      pages.push(i)
    }
  } else {
    if (current <= 4) {
      for (let i = 1; i <= 5; i++) pages.push(i)
      pages.push('...')
      pages.push(total)
    } else if (current >= total - 3) {
      pages.push(1)
      pages.push('...')
      for (let i = total - 4; i <= total; i++) pages.push(i)
    } else {
      pages.push(1)
      pages.push('...')
      for (let i = current - 1; i <= current + 1; i++) pages.push(i)
      pages.push('...')
      pages.push(total)
    }
  }

  return pages
})
</script>

<style scoped>
.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 12px;
  margin-top: 50px;
  flex-wrap: wrap;
}

.page-btn {
  min-height: 45px;
  border: 2px solid #e1e5e9;
  background: white;
  border-radius: 12px;
  cursor: pointer;
  font-size: 16px;
  font-weight: 600;
  color: #666;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0 15px;
  position: relative;
  overflow: hidden;
}

.page-btn::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(135deg, #4a90e2, #50e3c2);
  transition: left 0.3s ease;
  z-index: -1;
}

.page-btn:hover:not(:disabled)::before {
  left: 0;
}

.page-btn:hover:not(:disabled) {
  border-color: #4a90e2;
  color: white;
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(74, 144, 226, 0.3);
}

.page-btn.active {
  background: linear-gradient(135deg, #4a90e2, #50e3c2);
  border-color: #4a90e2;
  color: white;
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(74, 144, 226, 0.4);
}

.page-btn.active::before {
  left: 0;
}

.page-btn:disabled {
  opacity: 0.4;
  cursor: not-allowed;
  transform: none;
  box-shadow: none;
}

.page-btn:disabled::before {
  display: none;
}

.prev-btn,
.next-btn {
  min-width: 100px;
  gap: 8px;
}

.page-num {
  min-width: 45px;
}

.btn-icon {
  font-size: 18px;
  line-height: 1;
}

.btn-text {
  font-size: 14px;
}

.page-numbers {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
  justify-content: center;
}

@media (max-width: 768px) {
  .pagination {
    gap: 8px;
  }

  .page-btn {
    min-height: 40px;
    font-size: 14px;
    padding: 0 12px;
  }

  .prev-btn,
  .next-btn {
    min-width: 80px;
  }

  .page-num {
    min-width: 40px;
  }

  .btn-text {
    font-size: 12px;
  }

  .btn-icon {
    font-size: 16px;
  }
}

@media (max-width: 480px) {
  .pagination {
    flex-direction: column;
    gap: 15px;
  }

  .page-numbers {
    order: 1;
  }

  .prev-btn {
    order: 2;
  }

  .next-btn {
    order: 3;
  }
}
</style>
