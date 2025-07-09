<template>
  <article class="news-card">
    <a :href="link" class="card-link">
      <div class="card-image">
        <img :src="imageUrl" :alt="title" loading="lazy" />
        <div class="image-overlay">
          <div class="overlay-badges">
            <span v-if="isNew" class="badge new">新</span>
            <span v-if="isHot" class="badge hot">热门</span>
          </div>
          <div class="category-badge">{{ getCategoryName(category) }}</div>
        </div>
      </div>

      <div class="card-content">
        <h3 class="card-title">{{ title }}</h3>
        <p class="card-summary">{{ summary }}</p>

        <div class="card-footer">
          <div class="card-meta">
            <span class="card-date">
              <i class="icon">📅</i>
              {{ publishDate }}
            </span>
            <span class="card-views">
              <i class="icon">👁</i>
              {{ formatViews(views) }}
            </span>
          </div>

          <div class="read-more">
            <span>阅读更多</span>
            <i class="arrow">→</i>
          </div>
        </div>
      </div>
    </a>
  </article>
</template>

<script setup>
defineProps({
  title: String,
  summary: String,
  imageUrl: String,
  publishDate: String,
  link: String,
  category: String,
  views: {
    type: Number,
    default: 0,
  },
  isHot: {
    type: Boolean,
    default: false,
  },
  isNew: {
    type: Boolean,
    default: false,
  },
})

// 获取分类名称
const getCategoryName = (categoryId) => {
  const categoryMap = {
    company: '公司动态',
    industry: '行业资讯',
    technology: '技术前沿',
    education: '教育培训',
    research: '科研成果',
    competition: '竞赛活动',
  }
  return categoryMap[categoryId] || '其他'
}

// 格式化浏览量
const formatViews = (views) => {
  if (views >= 1000) {
    return (views / 1000).toFixed(1) + 'k'
  }
  return views.toString()
}
</script>

<style scoped>
.news-card {
  background: white;
  border-radius: 16px;
  overflow: hidden;
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.08);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  height: 100%;
  display: flex;
  flex-direction: column;
}

.news-card:hover {
  transform: translateY(-12px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15);
}

.card-link {
  text-decoration: none;
  color: inherit;
  display: flex;
  flex-direction: column;
  height: 100%;
}

.card-image {
  position: relative;
  width: 100%;
  height: 220px;
  overflow: hidden;
}

.card-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.6s ease;
}

.news-card:hover .card-image img {
  transform: scale(1.08);
}

.image-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(0, 0, 0, 0.1) 0%, rgba(0, 0, 0, 0.3) 100%);
  opacity: 0;
  transition: opacity 0.3s ease;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding: 15px;
}

.news-card:hover .image-overlay {
  opacity: 1;
}

.overlay-badges {
  display: flex;
  gap: 8px;
  align-self: flex-start;
}

.badge {
  padding: 4px 10px;
  border-radius: 12px;
  font-size: 0.75rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
}

.badge.new {
  background: linear-gradient(135deg, #48bb78 0%, #38a169 100%);
  color: white;
}

.badge.hot {
  background: linear-gradient(135deg, #ed8936 0%, #dd6b20 100%);
  color: white;
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0%,
  100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.05);
  }
}

.category-badge {
  align-self: flex-end;
  padding: 6px 12px;
  background: rgba(255, 255, 255, 0.95);
  color: #667eea;
  border-radius: 15px;
  font-size: 0.8rem;
  font-weight: 600;
  backdrop-filter: blur(10px);
}

.card-content {
  padding: 25px;
  flex-grow: 1;
  display: flex;
  flex-direction: column;
}

.card-title {
  font-size: 1.25rem;
  font-weight: 600;
  margin: 0 0 15px 0;
  line-height: 1.4;
  color: #2c3e50;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 2;
  overflow: hidden;
  text-overflow: ellipsis;
  min-height: 60px;
  transition: color 0.3s ease;
}

.news-card:hover .card-title {
  color: #667eea;
}

.card-summary {
  font-size: 0.95rem;
  color: #64748b;
  margin: 0 0 20px 0;
  line-height: 1.6;
  flex-grow: 1;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 3;
  overflow: hidden;
  text-overflow: ellipsis;
}

.card-footer {
  margin-top: auto;
}

.card-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
  padding-bottom: 15px;
  border-bottom: 1px solid #e2e8f0;
}

.card-date,
.card-views {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 0.85rem;
  color: #94a3b8;
  font-weight: 500;
}

.icon {
  font-size: 0.9rem;
}

.read-more {
  display: flex;
  align-items: center;
  justify-content: space-between;
  color: #667eea;
  font-weight: 600;
  font-size: 0.9rem;
  transition: all 0.3s ease;
}

.read-more .arrow {
  transition: transform 0.3s ease;
  font-size: 1.1rem;
}

.news-card:hover .read-more .arrow {
  transform: translateX(5px);
}

.news-card:hover .read-more {
  color: #5a6fd8;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .card-content {
    padding: 20px;
  }

  .card-title {
    font-size: 1.1rem;
    min-height: 50px;
  }

  .card-summary {
    font-size: 0.9rem;
    -webkit-line-clamp: 2;
  }

  .card-image {
    height: 200px;
  }
}

@media (max-width: 480px) {
  .card-meta {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }

  .card-date,
  .card-views {
    font-size: 0.8rem;
  }

  .overlay-badges {
    gap: 6px;
  }

  .badge {
    padding: 3px 8px;
    font-size: 0.7rem;
  }
}
</style>
