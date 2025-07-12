<template>
  <div v-if="totalPages > 1" class="pagination-container">
    <button
      class="pagination-button"
      :disabled="currentPage === 1"
      @click="changePage(currentPage - 1)"
    >
      &laquo;
    </button>
    <button
      v-for="page in pages"
      :key="page"
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
      &raquo;
    </button>
  </div>
</template>

<script>
export default {
  name: 'PaginationComponent',
  props: {
    currentPage: {
      type: Number,
      required: true,
    },
    totalItems: {
      type: Number,
      required: true,
    },
    itemsPerPage: {
      type: Number,
      required: true,
    },
    maxPagesShown: {
      type: Number,
      default: 5,
    },
  },
  computed: {
    totalPages() {
      return Math.ceil(this.totalItems / this.itemsPerPage)
    },
    pages() {
      const pages = []
      const { currentPage, totalPages, maxPagesShown } = this

      if (totalPages <= maxPagesShown) {
        for (let i = 1; i <= totalPages; i++) {
          pages.push(i)
        }
      } else {
        pages.push(1)
        let startPage = Math.max(2, currentPage - Math.floor((maxPagesShown - 3) / 2))
        let endPage = Math.min(totalPages - 1, startPage + maxPagesShown - 3)

        if (currentPage > totalPages - Math.floor((maxPagesShown - 1) / 2)) {
          startPage = totalPages - maxPagesShown + 2
          endPage = totalPages - 1
        }

        if (startPage > 2) {
          pages.push('...')
        }

        for (let i = startPage; i <= endPage; i++) {
          pages.push(i)
        }

        if (endPage < totalPages - 1) {
          pages.push('...')
        }

        pages.push(totalPages)
      }
      return pages
    },
  },
  methods: {
    changePage(page) {
      if (page > 0 && page <= this.totalPages && page !== this.currentPage) {
        this.$emit('page-changed', page)
      }
    },
  },
}
</script>

<style scoped>
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
