<template>
  <div class="info-center">
    <!-- Tab 导航 -->
    <div class="tabs">
      <button
        v-for="category in categories"
        :key="category"
        :class="{ active: activeCategory === category }"
        @click="activeCategory = category"
      >
        {{ category }}
      </button>
    </div>

    <!-- 新闻列表 -->
    <TransitionGroup name="fade-list" tag="div" class="news-list">
      <NewsCard
        v-for="item in filteredNews"
        :key="item.id"
        :title="item.title"
        :summary="item.summary"
        :image-url="item.imageUrl"
        :publish-date="item.publishDate"
        :link="item.link"
      />
    </TransitionGroup>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import NewsCard from './NewCard.vue'

// 模拟从 API 获取的数据
const allNews = ref([
  {
    id: 1,
    category: '公司新闻',
    title: '我们公司发布了一款革命性的新产品',
    summary: '这款产品将彻底改变行业格局，为用户带来前所未有的体验，敬请期待！',
    imageUrl: 'https://picsum.photos/id/1/400/300',
    publishDate: '2023-10-27',
    link: '#',
  },
  {
    id: 2,
    category: '行业动态',
    title: 'AI 技术在前端开发中的应用新趋势',
    summary: '从代码生成到智能测试，AI 正在深刻影响着开发者。本文将探讨最新的AI技术趋势。',
    imageUrl: 'https://picsum.photos/id/2/400/300',
    publishDate: '2023-10-26',
    link: '#',
  },
  {
    id: 3,
    category: '公司新闻',
    title: '公司荣获“年度最佳雇主”称号',
    summary: '凭借卓越的企业文化和员工福利，我司再次被评为年度最佳雇主。',
    imageUrl: 'https://picsum.photos/id/3/400/300',
    publishDate: '2023-10-25',
    link: '#',
  },
  {
    id: 4,
    category: '媒体报道',
    title: '知名科技媒体深度报道我司创新文化',
    summary: 'TechWeekly对我司进行了专访，高度评价了我们的研发实力和市场前景。',
    imageUrl: 'https://picsum.photos/id/4/400/300',
    publishDate: '2023-10-24',
    link: '#',
  },
  {
    id: 5,
    category: '行业动态',
    title: 'Vue 3.4 版本发布，性能迎来新飞跃',
    summary: '最新的Vue版本带来了更快的编译速度和更小的打包体积，开发者社区反响热烈。',
    imageUrl: 'https://picsum.photos/id/5/400/300',
    publishDate: '2023-10-23',
    link: '#',
  },
  {
    id: 6,
    category: '公司新闻',
    title: '年度开发者大会圆满落幕',
    summary: '我们成功举办了年度开发者大会，吸引了全球上千名开发者参与。',
    imageUrl: 'https://picsum.photos/id/6/400/300',
    publishDate: '2023-10-22',
    link: '#',
  },
])

const categories = computed(() => ['公司新闻', '行业动态', '媒体报道'])
const activeCategory = ref('公司新闻')

const filteredNews = computed(() => {
  if (!activeCategory.value) return allNews.value
  return allNews.value.filter((news) => news.category === activeCategory.value)
})
</script>

<style scoped>
.info-center {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
  background-color: #fff;
  border-radius: 12px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.05);
}

.tabs {
  margin-bottom: 30px;
  border-bottom: 2px solid #e0e0e0;
  display: flex;
  justify-content: center;
  gap: 10px;
}

.tabs button {
  padding: 12px 25px;
  border: none;
  background: none;
  cursor: pointer;
  font-size: 16px;
  font-weight: 500;
  color: #666;
  position: relative;
  transition: all 0.3s ease;
  border-radius: 8px 8px 0 0;
}

.tabs button:hover {
  color: #007bff;
  background-color: #f0f8ff;
}

.tabs button.active {
  color: #007bff;
  font-weight: 600;
}

.tabs button.active::after {
  content: '';
  position: absolute;
  bottom: -2px;
  left: 0;
  width: 100%;
  height: 2px;
  background-color: #007bff;
}

.news-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 30px;
}

/* 列表过渡动画 */
.fade-list-move,
.fade-list-enter-active,
.fade-list-leave-active {
  transition: all 0.5s cubic-bezier(0.55, 0, 0.1, 1);
}
.fade-list-enter-from,
.fade-list-leave-to {
  opacity: 0;
  transform: scaleY(0.01) translate(30px, 0);
}
.fade-list-leave-active {
  position: absolute;
}
</style>
