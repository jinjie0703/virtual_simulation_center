<template>
  <div class="info-center">
    <!-- 页面标题区域 -->
    <div class="header-section">
      <h1 class="main-title">信息中心</h1>
      <p class="main-subtitle">掌握最新动态，了解前沿资讯</p>
    </div>

    <!-- 统计卡片 -->
    <div class="stats-grid">
      <div class="stat-card">
        <div class="stat-icon">📰</div>
        <div class="stat-content">
          <h3>{{ allNews.length }}</h3>
          <p>总新闻数</p>
        </div>
      </div>
      <div class="stat-card">
        <div class="stat-icon">🔥</div>
        <div class="stat-content">
          <h3>{{ todayNews }}</h3>
          <p>今日更新</p>
        </div>
      </div>
      <div class="stat-card">
        <div class="stat-icon">👀</div>
        <div class="stat-content">
          <h3>{{ totalViews }}</h3>
          <p>总浏览量</p>
        </div>
      </div>
      <div class="stat-card">
        <div class="stat-icon">⭐</div>
        <div class="stat-content">
          <h3>{{ categories.length - 1 }}</h3>
          <p>新闻分类</p>
        </div>
      </div>
    </div>

    <!-- 搜索和筛选区域 -->
    <div class="filter-section">
      <div class="search-container">
        <div class="search-box">
          <input
            v-model="searchQuery"
            type="text"
            placeholder="搜索新闻标题或内容..."
            class="search-input"
          />
          <button class="search-btn">
            <i class="search-icon">🔍</i>
          </button>
        </div>
      </div>

      <div class="filter-controls">
        <select v-model="sortBy" class="sort-select">
          <option value="date">按日期排序</option>
          <option value="views">按浏览量排序</option>
          <option value="title">按标题排序</option>
        </select>

        <select v-model="sortOrder" class="sort-select">
          <option value="desc">降序</option>
          <option value="asc">升序</option>
        </select>
      </div>
    </div>

    <!-- 分类标签 -->
    <div class="tabs">
      <button
        v-for="category in categories"
        :key="category.id"
        :class="['tab-btn', { active: activeCategory === category.id }]"
        @click="setActiveCategory(category.id)"
      >
        <span class="tab-icon">{{ category.icon }}</span>
        <span class="tab-text">{{ category.name }}</span>
        <span class="tab-count">({{ getCategoryCount(category.id) }})</span>
      </button>
    </div>

    <!-- 新闻列表 -->
    <div class="news-section">
      <div v-if="filteredNews.length === 0" class="empty-state">
        <div class="empty-icon">📭</div>
        <h3>暂无相关新闻</h3>
        <p>请尝试其他搜索条件或分类</p>
      </div>

      <div v-else class="news-grid">
        <NewsCard
          v-for="news in paginatedNews"
          :key="news.id"
          :title="news.title"
          :summary="news.summary"
          :image-url="news.imageUrl"
          :publish-date="formatDate(news.publishDate)"
          :link="news.link"
          :category="news.category"
          :views="news.views"
          :is-hot="news.isHot"
          :is-new="isNew(news.publishDate)"
        />
      </div>

      <!-- 分页组件 -->
      <div v-if="totalPages > 1" class="pagination">
        <button :disabled="currentPage === 1" @click="goToPage(currentPage - 1)" class="page-btn">
          ‹ 上一页
        </button>

        <div class="page-numbers">
          <button
            v-for="page in visiblePages"
            :key="page"
            :class="['page-num', { active: page === currentPage }]"
            @click="goToPage(page)"
          >
            {{ page }}
          </button>
        </div>

        <button
          :disabled="currentPage === totalPages"
          @click="goToPage(currentPage + 1)"
          class="page-btn"
        >
          下一页 ›
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import NewsCard from '@/components/information_center/NewsCard.vue'

// 响应式数据
const activeCategory = ref('')
const searchQuery = ref('')
const sortBy = ref('date')
const sortOrder = ref('desc')
const currentPage = ref(1)
const itemsPerPage = 12

// 分类数据
const categories = ref([
  { id: '', name: '全部', icon: '📋' },
  { id: 'company', name: '公司动态', icon: '🏢' },
  { id: 'industry', name: '行业资讯', icon: '🌐' },
  { id: 'technology', name: '技术前沿', icon: '💡' },
  { id: 'education', name: '教育培训', icon: '🎓' },
  { id: 'research', name: '科研成果', icon: '🔬' },
  { id: 'competition', name: '竞赛活动', icon: '🏆' },
])

// 模拟新闻数据
const allNews = ref([
  {
    id: 1,
    title: '虚拟仿真技术在医学教育中的突破性应用',
    summary:
      '我校医学院成功开发了基于VR技术的外科手术模拟系统，为医学生提供安全、高效的实践环境，获得教育部高度认可。',
    imageUrl: 'https://images.unsplash.com/photo-1559757148-5c350d0d3c56?w=400&h=250&fit=crop',
    publishDate: '2024-05-20',
    category: 'education',
    views: 1250,
    isHot: true,
    link: '#',
  },
  {
    id: 2,
    title: 'AI人工智能实验室正式投入使用',
    summary:
      '历时两年建设的AI人工智能实验室正式启用，配备了最先进的深度学习硬件设备和软件平台，为师生提供一流的研究环境。',
    imageUrl: 'https://images.unsplash.com/photo-1485827404703-89b55fcc595e?w=400&h=250&fit=crop',
    publishDate: '2024-05-18',
    category: 'company',
    views: 890,
    isHot: true,
    link: '#',
  },
  {
    id: 3,
    title: '全国虚拟现实教育应用大赛获奖名单公布',
    summary:
      '我校学生团队在全国虚拟现实教育应用大赛中荣获一等奖，作品"沉浸式历史文化体验系统"获得专家一致好评。',
    imageUrl: 'https://images.unsplash.com/photo-1552664730-d307ca884978?w=400&h=250&fit=crop',
    publishDate: '2024-05-15',
    category: 'competition',
    views: 567,
    isHot: false,
    link: '#',
  },
  {
    id: 4,
    title: '数字孪生技术在智能制造领域的最新进展',
    summary:
      '数字孪生技术正在revolutionizing智能制造行业，通过虚实结合的方式大幅提升生产效率和产品质量。',
    imageUrl: 'https://images.unsplash.com/photo-1518709268805-4e9042af2176?w=400&h=250&fit=crop',
    publishDate: '2024-05-12',
    category: 'technology',
    views: 743,
    isHot: false,
    link: '#',
  },
  {
    id: 5,
    title: '虚拟仿真中心与知名企业达成战略合作',
    summary:
      '我校虚拟仿真中心与多家知名科技企业签署战略合作协议，共同推进产学研一体化发展，培养高素质应用型人才。',
    imageUrl: 'https://images.unsplash.com/photo-1542744094-3a31f272c490?w=400&h=250&fit=crop',
    publishDate: '2024-05-10',
    category: 'company',
    views: 623,
    isHot: false,
    link: '#',
  },
  {
    id: 6,
    title: '混合现实技术助力远程教育新模式',
    summary:
      '基于MR技术的远程教育平台正式上线，实现了真正意义上的"身临其境"学习体验，为疫情期间的教育提供了创新解决方案。',
    imageUrl: 'https://images.unsplash.com/photo-1593508512255-86ab42a8e620?w=400&h=250&fit=crop',
    publishDate: '2024-05-08',
    category: 'education',
    views: 456,
    isHot: false,
    link: '#',
  },
  {
    id: 7,
    title: '科研团队在顶级期刊发表VR技术论文',
    summary:
      '我校虚拟现实研究团队在国际顶级期刊《Nature》上发表关于沉浸式学习效果的重要研究成果，引起学术界广泛关注。',
    imageUrl: 'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=400&h=250&fit=crop',
    publishDate: '2024-05-05',
    category: 'research',
    views: 892,
    isHot: true,
    link: '#',
  },
  {
    id: 8,
    title: '虚拟实验室建设标准化指南发布',
    summary:
      '教育部发布虚拟实验室建设标准化指南，为全国高校建设高质量虚拟实验室提供权威参考和技术规范。',
    imageUrl: 'https://images.unsplash.com/photo-1581092921461-eab62e97a780?w=400&h=250&fit=crop',
    publishDate: '2024-05-03',
    category: 'industry',
    views: 334,
    isHot: false,
    link: '#',
  },
])

// 计算属性
const todayNews = computed(() => {
  const today = new Date().toISOString().split('T')[0]
  return allNews.value.filter((news) => news.publishDate === today).length
})

const totalViews = computed(() => {
  return allNews.value.reduce((sum, news) => sum + news.views, 0).toLocaleString()
})

const filteredNews = computed(() => {
  let filtered = allNews.value

  // 分类筛选
  if (activeCategory.value) {
    filtered = filtered.filter((news) => news.category === activeCategory.value)
  }

  // 搜索筛选
  if (searchQuery.value.trim()) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(
      (news) =>
        news.title.toLowerCase().includes(query) || news.summary.toLowerCase().includes(query),
    )
  }

  // 排序
  filtered.sort((a, b) => {
    let comparison = 0
    switch (sortBy.value) {
      case 'date':
        comparison = new Date(b.publishDate) - new Date(a.publishDate)
        break
      case 'views':
        comparison = b.views - a.views
        break
      case 'title':
        comparison = a.title.localeCompare(b.title)
        break
    }
    return sortOrder.value === 'desc' ? comparison : -comparison
  })

  return filtered
})

const totalPages = computed(() => {
  return Math.ceil(filteredNews.value.length / itemsPerPage)
})

const paginatedNews = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage
  const end = start + itemsPerPage
  return filteredNews.value.slice(start, end)
})

const visiblePages = computed(() => {
  const pages = []
  const total = totalPages.value
  const current = currentPage.value

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

// 方法
const setActiveCategory = (categoryId) => {
  activeCategory.value = categoryId
  currentPage.value = 1
}

const getCategoryCount = (categoryId) => {
  if (!categoryId) return allNews.value.length
  return allNews.value.filter((news) => news.category === categoryId).length
}

const goToPage = (page) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
  }
}

const formatDate = (dateStr) => {
  const date = new Date(dateStr)
  const now = new Date()
  const diffTime = now - date
  const diffDays = Math.floor(diffTime / (1000 * 60 * 60 * 24))

  if (diffDays === 0) return '今天'
  if (diffDays === 1) return '昨天'
  if (diffDays < 7) return `${diffDays}天前`

  return date.toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
  })
}

const isNew = (dateStr) => {
  const date = new Date(dateStr)
  const now = new Date()
  const diffDays = Math.floor((now - date) / (1000 * 60 * 60 * 24))
  return diffDays <= 3
}

// 监听器
watch([searchQuery, activeCategory, sortBy, sortOrder], () => {
  currentPage.value = 1
})
</script>

<style scoped>
.info-center {
  max-width: 1200px;
  margin: 0 auto;
  padding: 30px 20px;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  min-height: 100vh;
}

/* 页面标题 */
.header-section {
  text-align: center;
  margin-bottom: 40px;
}

.main-title {
  font-size: 3rem;
  font-weight: 700;
  color: #2c3e50;
  margin: 0 0 16px 0;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.main-subtitle {
  font-size: 1.2rem;
  color: #64748b;
  margin: 0;
}

/* 统计卡片 */
.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
  margin-bottom: 40px;
}

.stat-card {
  background: white;
  padding: 25px;
  border-radius: 16px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.08);
  display: flex;
  align-items: center;
  gap: 20px;
  transition: all 0.3s ease;
}

.stat-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.12);
}

.stat-icon {
  font-size: 2.5rem;
  width: 60px;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 12px;
  color: white;
  box-shadow: 0 8px 20px rgba(102, 126, 234, 0.3);
}

.stat-content h3 {
  font-size: 2rem;
  font-weight: 700;
  color: #2c3e50;
  margin: 0 0 5px 0;
}

.stat-content p {
  font-size: 0.9rem;
  color: #64748b;
  margin: 0;
}

/* 搜索和筛选 */
.filter-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30px;
  gap: 20px;
  flex-wrap: wrap;
}

.search-container {
  flex: 1;
  min-width: 300px;
}

.search-box {
  position: relative;
  display: flex;
  background: white;
  border-radius: 25px;
  overflow: hidden;
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.08);
}

.search-input {
  flex: 1;
  padding: 15px 20px;
  border: none;
  outline: none;
  font-size: 1rem;
  background: transparent;
}

.search-btn {
  padding: 15px 20px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  cursor: pointer;
  transition: all 0.3s ease;
}

.search-btn:hover {
  background: linear-gradient(135deg, #5a6fd8 0%, #6a4190 100%);
}

.filter-controls {
  display: flex;
  gap: 15px;
}

.sort-select {
  padding: 10px 15px;
  border: 2px solid #e2e8f0;
  border-radius: 8px;
  background: white;
  font-size: 0.9rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

.sort-select:focus {
  outline: none;
  border-color: #667eea;
}

/* 分类标签 */
.tabs {
  display: flex;
  gap: 12px;
  margin-bottom: 40px;
  flex-wrap: wrap;
  justify-content: center;
}

.tab-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 20px;
  border: 2px solid #e2e8f0;
  background: white;
  border-radius: 25px;
  cursor: pointer;
  transition: all 0.3s ease;
  font-size: 0.9rem;
  font-weight: 500;
  color: #64748b;
}

.tab-btn:hover {
  border-color: #667eea;
  color: #667eea;
  transform: translateY(-2px);
}

.tab-btn.active {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-color: transparent;
  color: white;
  box-shadow: 0 8px 20px rgba(102, 126, 234, 0.3);
}

.tab-icon {
  font-size: 1.1rem;
}

.tab-count {
  font-size: 0.8rem;
  opacity: 0.8;
}

/* 新闻网格 */
.news-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 25px;
  margin-bottom: 40px;
}

/* 空状态 */
.empty-state {
  text-align: center;
  padding: 60px 20px;
  color: #64748b;
}

.empty-icon {
  font-size: 4rem;
  margin-bottom: 20px;
}

.empty-state h3 {
  font-size: 1.5rem;
  margin-bottom: 10px;
  color: #2c3e50;
}

/* 分页 */
.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 10px;
  margin-top: 40px;
}

.page-btn {
  padding: 10px 15px;
  border: 2px solid #e2e8f0;
  background: white;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  font-weight: 500;
}

.page-btn:hover:not(:disabled) {
  border-color: #667eea;
  color: #667eea;
}

.page-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.page-numbers {
  display: flex;
  gap: 5px;
}

.page-num {
  width: 40px;
  height: 40px;
  border: 2px solid #e2e8f0;
  background: white;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 500;
}

.page-num:hover {
  border-color: #667eea;
  color: #667eea;
}

.page-num.active {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-color: transparent;
  color: white;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .info-center {
    padding: 20px 15px;
  }

  .main-title {
    font-size: 2.2rem;
  }

  .stats-grid {
    grid-template-columns: 1fr;
    gap: 15px;
  }

  .filter-section {
    flex-direction: column;
    align-items: stretch;
  }

  .search-container {
    min-width: auto;
  }

  .filter-controls {
    justify-content: center;
  }

  .tabs {
    justify-content: flex-start;
    overflow-x: auto;
    padding-bottom: 10px;
  }

  .news-grid {
    grid-template-columns: 1fr;
    gap: 20px;
  }

  .pagination {
    flex-wrap: wrap;
    gap: 8px;
  }
}

@media (max-width: 480px) {
  .stat-card {
    padding: 20px;
    gap: 15px;
  }

  .stat-icon {
    font-size: 2rem;
    width: 50px;
    height: 50px;
  }

  .stat-content h3 {
    font-size: 1.5rem;
  }

  .tab-btn {
    padding: 10px 15px;
    font-size: 0.8rem;
  }
}
</style>
