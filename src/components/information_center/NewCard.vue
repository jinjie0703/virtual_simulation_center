<template>
  <a :href="link" class="news-card">
    <div class="card-image">
      <img :src="imageUrl" :alt="title" loading="lazy" />
    </div>
    <div class="card-content">
      <h3 class="card-title">{{ title }}</h3>
      <p class="card-summary">{{ summary }}</p>
      <span class="card-date">{{ publishDate }}</span>
    </div>
  </a>
</template>

<script setup>
defineProps({
  title: String,
  summary: String,
  imageUrl: String,
  publishDate: String,
  link: String,
})
</script>

<style scoped>
.news-card {
  display: flex;
  flex-direction: column;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  transition:
    transform 0.3s ease,
    box-shadow 0.3s ease;
  text-decoration: none;
  color: inherit;
  background-color: #ffffff;
}

.news-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.12);
}

.card-image {
  width: 100%;
  height: 200px; /* 固定高度 */
  overflow: hidden;
}

.card-image img {
  width: 100%;
  height: 100%;
  object-fit: cover; /* 保证图片不变形且填满容器 */
  transition: transform 0.4s ease;
}

.news-card:hover .card-image img {
  transform: scale(1.05); /* 悬停时图片轻微放大 */
}

.card-content {
  padding: 20px;
  flex-grow: 1;
  display: flex;
  flex-direction: column;
}

.card-title {
  font-size: 1.15rem;
  font-weight: 600;
  margin: 0 0 10px;
  line-height: 1.4;
  /* 多行文字超出显示省略号 */
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 2; /* 最多显示2行 */
  overflow: hidden;
  text-overflow: ellipsis;
  min-height: 45px; /* 保证标题区域高度基本一致 */
}

.card-summary {
  font-size: 0.9rem;
  color: #666;
  margin: 0 0 15px;
  line-height: 1.6;
  flex-grow: 1; /* 让摘要占据剩余空间 */
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 3; /* 最多显示3行 */
  overflow: hidden;
  text-overflow: ellipsis;
}

.card-date {
  font-size: 0.8rem;
  color: #999;
  align-self: flex-start; /* 日期总是在底部 */
}
</style>
