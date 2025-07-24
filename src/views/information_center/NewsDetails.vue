<!-- src/views/NewsDetail.vue -->
<template>
  <div class="article-page" v-if="article">
    <header class="article-header">
      <div class="header-content-wrapper">
        <span class="category-badge">{{ article.category }}</span>
        <h1 class="main-title">{{ article.title }}</h1>
        <div class="meta-data">
          <span>作者：{{ article.author }}</span>
          <span>发布时间：{{ article.publishDate }}</span>
        </div>
      </div>
    </header>

    <main class="article-body-container">
      <div class="article-main-content">
        <!-- 动态渲染的Markdown内容将在这里显示 -->
        <div class="prose" v-html="formattedContent"></div>
        <button @click="goBack" class="back-link">← 返回上一页</button>
      </div>

      <aside class="article-sidebar">
        <div class="sidebar-widget">
          <h3>文章标签</h3>
          <div class="tags-cloud">
            <span v-for="tag in article.tags" :key="tag" class="tag">{{ tag }}</span>
          </div>
        </div>
        <div class="sidebar-widget">
          <h3>相关推荐</h3>
          <ul class="related-list">
            <li><a href="#">文章1</a></li>
            <li><a href="#">文章2</a></li>
          </ul>
        </div>
      </aside>
    </main>
  </div>
  <div v-else class="loading-state">
    <p>正在加载文章...</p>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { marked } from 'marked'
import DOMPurify from 'dompurify'

// 引入我们的模拟数据
import { mockArticle } from '@/components/information_center/TestData.js'

const route = useRoute()
const router = useRouter()
const article = ref(null)

const goBack = () => {
  router.back()
}

// 计算属性，用于安全地将Markdown转换为HTML
const formattedContent = computed(() => {
  if (article.value && article.value.content) {
    const rawHtml = marked.parse(article.value.content)
    return DOMPurify.sanitize(rawHtml)
  }
  return ''
})

onMounted(() => {
  const articleId = route.params.id

  // --- 模拟API请求 ---
  // 真实场景: article.value = await fetch(`/api/news/${articleId}`).then(res => res.json())
  console.log(`正在查找ID为 "${articleId}" 的文章...`)

  // 在这个示例中，我们简单地加载唯一的模拟文章
  // 在真实应用中，您会用 find 方法在文章列表中查找
  // article.value = allMockData.news.find(a => a.id === articleId)
  article.value = mockArticle
  // --- 模拟结束 ---
})
</script>

<style>
/*
  为了让 v-html 渲染出的内容能应用样式，我们在这里使用全局style。
  如果您的项目配置了scoped CSS，可以改用 <style scoped> 和 ::v-deep 或 :deep() 选择器。
*/

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
  margin: -40px auto 0; /* 负外边距让内容区和头部有重叠，更现代 */
  gap: 40px;
  padding: 0 20px;
  align-items: flex-start;
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
  top: 100px; /* 滚动时侧边栏会固定 */
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
.related-list {
  list-style: none;
  padding: 0;
  margin: 0;
}
.related-list li a {
  display: block;
  padding: 10px 0;
  text-decoration: none;
  color: #2d3748;
  border-bottom: 1px solid #f0f2f5;
  transition: all 0.2s;
}
.related-list li a:hover {
  color: #4c51bf;
  transform: translateX(5px);
}
.related-list li:last-child a {
  border-bottom: none;
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
  /* --- 新增样式以重置按钮默认外观 --- */
  border: none;
  font-family: inherit; /* 继承父元素的字体 */
  font-size: inherit; /* 继承父元素的字号 */
  cursor: pointer;
  transition: background-color 0.2s;
}

.back-link:hover {
  background-color: #e2e8f0;
}

/* 核心：为动态生成的HTML内容（文章本身）设计样式 */
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

/* 响应式设计 */
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
