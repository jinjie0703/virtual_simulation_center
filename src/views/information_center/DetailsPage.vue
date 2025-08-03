<!-- src/views/information_center/DetailsPage.vue -->
<template>
  <div class="article-page" v-if="item">
    <header class="article-header">
      <div class="header-content-wrapper">
        <span class="category-badge">{{ item.category || item.level }}</span>
        <h1 class="main-title">{{ item.title }}</h1>
        <div class="meta-data">
          <span v-if="item.author">作者：{{ item.author }}</span>
          <span v-if="item.publish_date">发布时间：{{ formatDate(item.publish_date) }}</span>
          <span v-if="item.deadline">截止时间：{{ formatDate(item.deadline) }}</span>
        </div>
      </div>
    </header>

    <main class="article-body-container" :class="{ 'no-sidebar': itemType === 'news' }">
      <div class="article-main-content">
        <div class="prose" v-html="formattedContent"></div>
        <button @click="goBack" class="back-link">← 返回上一页</button>
      </div>

      <aside class="article-sidebar" v-if="itemType !== 'news'">
        <div class="sidebar-widget">
          <h3>标签</h3>
          <div class="tags-cloud">
            <span v-for="tag in (item.tags || '').split(',')" :key="tag" class="tag">{{
              tag
            }}</span>
          </div>
        </div>
      </aside>
    </main>
  </div>
  <div v-else class="loading-state">
    <p>正在加载内容...</p>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { marked } from 'marked'
import DOMPurify from 'dompurify'
import axios from 'axios'

const route = useRoute()
const router = useRouter()
const item = ref(null)
const itemType = ref(route.meta.itemType || 'news')

const goBack = () => {
  router.back()
}

const formatDate = (dateString) => {
  if (!dateString) return ''
  return new Date(dateString).toISOString().split('T')[0]
}

const formattedContent = computed(() => {
  if (item.value && item.value.content) {
    const rawHtml = marked.parse(item.value.content)
    return DOMPurify.sanitize(rawHtml)
  }
  return ''
})

onMounted(async () => {
  const itemId = route.params.id
  try {
    const itemResponse = await axios.get(
      `https://localhost:8080/api/information_center/${itemType.value}/${itemId}`,
    )
    const itemData = itemResponse.data

    if (itemData.detail_url) {
      try {
        const contentResponse = await axios.get(itemData.detail_url)
        itemData.content = contentResponse.data
      } catch (contentError) {
        console.error(`Failed to fetch content from ${itemData.detail_url}:`, contentError)
        // 如果无法从 detail_url 获取内容，我们可以选择显示一条消息或使用默认内容
        itemData.content = '# 内容加载失败\n\n无法从指定的 URL 加载文章内容。'
      }
    } else {
      // 如果没有 detail_url，也设置一个默认内容
      itemData.content = '# 暂无详细内容'
    }

    item.value = itemData
  } catch (error) {
    console.error(`Failed to fetch initial ${itemType.value} data:`, error)
  }
})
</script>

<style>
/* Styles from NewsDetails.vue can be copied here */
.article-page {
  background-color: #f8f9fa;
  padding-bottom: 60px;
}

.article-header {
  background:
    linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)),
    url('https://images.unsplash.com/photo-1519389950473-47ba0277781c?q=80&w=1400&auto=format&fit=crop')
      no-repeat center center;
  background-size: cover;
  padding: 100px 20px 60px;
  color: white;
  text-align: center;
}

.header-content-wrapper {
  max-width: 800px;
  margin: 0 auto;
}

.category-badge {
  display: inline-block;
  background-color: rgba(255, 255, 255, 0.2);
  padding: 6px 14px;
  border-radius: 999px;
  font-size: 0.9rem;
  font-weight: 600;
  margin-bottom: 20px;
}

.main-title {
  font-size: 3rem;
  font-weight: 800;
  line-height: 1.2;
  margin: 0 0 20px;
}

.meta-data {
  display: flex;
  justify-content: center;
  gap: 30px;
  font-size: 0.9rem;
  opacity: 0.8;
}

.article-body-container {
  display: flex;
  max-width: 1200px;
  margin: -40px auto 0;
  gap: 40px;
  padding: 0 20px;
  align-items: flex-start;
}

.article-body-container.no-sidebar .article-main-content {
  flex: 1;
  max-width: 1080px; /* Increased width for better readability */
  margin: 0 auto;
}

.article-main-content {
  flex: 3;
  background: #ffffff;
  padding: 50px;
  border-radius: 12px;
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.08);
}

.article-sidebar {
  flex: 1;
  position: sticky;
  top: 100px;
}

.sidebar-widget {
  background: #ffffff;
  padding: 25px;
  border-radius: 12px;
  margin-bottom: 30px;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.05);
}
.sidebar-widget h3 {
  font-size: 1.2rem;
  margin: 0 0 20px;
  border-bottom: 1px solid #e2e8f0;
  padding-bottom: 15px;
}
.tags-cloud {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}
.tags-cloud .tag {
  background: #e2e8f0;
  color: #4a5568;
  padding: 6px 12px;
  border-radius: 6px;
  font-size: 0.85rem;
  cursor: pointer;
  transition: all 0.2s;
}
.tags-cloud .tag:hover {
  background-color: #4c51bf;
  color: white;
}
.back-link {
  display: inline-block;
  margin-top: 40px;
  color: #4c51bf;
  text-decoration: none;
  font-weight: 600;
  padding: 10px 20px;
  border-radius: 8px;
  background-color: #f0f2f5;
  border: none;
  font-family: inherit;
  font-size: inherit;
  cursor: pointer;
  transition: background-color 0.2s;
}

.back-link:hover {
  background-color: #e2e8f0;
}

.prose {
  line-height: 1.8;
  color: #34495e;
}
.prose h2,
.prose h3 {
  color: #2c3e50;
  margin-top: 2em;
  margin-bottom: 1em;
  padding-bottom: 0.4em;
  border-bottom: 1px solid #e2e8f0;
}
.prose h2 {
  font-size: 1.75rem;
}
.prose h3 {
  font-size: 1.4rem;
}
.prose p {
  margin-bottom: 1.25em;
}
.prose ul,
.prose ol {
  padding-left: 1.5em;
  margin-bottom: 1.25em;
}
.prose a {
  color: #4c51bf;
  font-weight: 600;
  text-decoration: none;
  transition: all 0.2s;
}
.prose a:hover {
  text-decoration: underline;
}
.prose blockquote {
  border-left: 4px solid #a0aec0;
  padding: 10px 20px;
  margin: 2em 0;
  background: #f8f9fa;
  color: #718096;
  font-style: italic;
}
.prose strong {
  font-weight: 700;
  color: #2c3e50;
}
.prose img {
  max-width: 100%;
  border-radius: 8px;
  margin: 2em 0;
}
.prose hr {
  border: none;
  border-top: 1px solid #e2e8f0;
  margin: 3em 0;
}

.loading-state {
  display: flex;
  justify-content: center;
  align-items: center;
  height: calc(100vh - 200px);
  font-size: 1.5rem;
  color: #4a5568;
}

@media (max-width: 992px) {
  .article-body-container {
    flex-direction: column;
    margin-top: -60px;
  }
  .article-sidebar {
    position: static;
    width: 100%;
  }
}
@media (max-width: 768px) {
  .main-title {
    font-size: 2.2rem;
  }
  .article-main-content {
    padding: 30px;
  }
}
</style>
