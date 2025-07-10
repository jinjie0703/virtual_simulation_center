<template>
  <article class="info-card">
    <div class="card-image-wrapper">
      <img :src="item.imageUrl" :alt="item.title" class="card-image" loading="lazy" />
    </div>
    <div class="card-content">
      <div class="card-header">
        <span v-if="type === 'news'" class="card-category">{{ item.category }}</span>
        <span v-if="type === 'competitions'" class="card-category level">{{ item.level }}</span>
        <div v-if="type === 'projects'" class="card-tags">
          <span v-for="tag in item.tags" :key="tag" class="tag">{{ tag }}</span>
        </div>
      </div>
      <h3 class="card-title">{{ item.title }}</h3>
      <p class="card-summary">{{ item.summary }}</p>
      <div class="card-footer">
        <div class="meta-info">
          <span v-if="item.date" class="date">📅 {{ item.date }}</span>
          <span v-if="item.deadline" class="deadline">⏳ 截止：{{ item.deadline }}</span>
        </div>
        <div v-if="type === 'projects'" class="reward"><span>💰</span> {{ item.reward }}</div>
        <a href="#" class="details-link">查看详情 →</a>
      </div>
    </div>
  </article>
</template>

<script setup>
defineProps({
  item: Object,
  type: String,
})
</script>

<style scoped>
.info-card {
  background: #ffffff;
  border-radius: 16px;
  overflow: hidden;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease-in-out;
  display: flex;
  border: 1px solid #e2e8f0;
}
.info-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 15px 30px rgba(0, 0, 0, 0.1);
}
.card-image-wrapper {
  width: 200px;
  flex-shrink: 0;
}
.card-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
.card-content {
  padding: 25px;
  display: flex;
  flex-direction: column;
  flex-grow: 1;
}
.card-header {
  margin-bottom: 12px;
}
.card-category {
  display: inline-block;
  padding: 4px 12px;
  border-radius: 999px;
  font-size: 0.75rem;
  font-weight: 700;
  background-color: #e9d8fd;
  color: #805ad5;
}
.card-category.level {
  background-color: #ccebfb;
  color: #3182ce;
}
.card-tags {
  display: flex;
  gap: 8px;
}
.tag {
  background: #e2e8f0;
  color: #4a5568;
  padding: 4px 10px;
  border-radius: 6px;
  font-size: 0.75rem;
  font-weight: 500;
}
.card-title {
  font-size: 1.2rem;
  font-weight: 700;
  color: #2d3748;
  margin: 0 0 10px;
  line-height: 1.4;
}
.card-summary {
  font-size: 0.9rem;
  color: #718096;
  line-height: 1.6;
  margin: 0 0 20px;
  flex-grow: 1;
}
.card-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: auto;
  border-top: 1px solid #edf2f7;
  padding-top: 15px;
}
.meta-info .date,
.meta-info .deadline {
  font-size: 0.85rem;
  color: #a0aec0;
  font-weight: 500;
}
.reward {
  font-size: 1.1rem;
  font-weight: 700;
  color: #dd6b20;
}
.reward span {
  font-size: 1.2rem;
}
.details-link {
  text-decoration: none;
  color: #4c51bf;
  font-weight: 600;
  font-size: 0.9rem;
  transition: color 0.2s;
}
.details-link:hover {
  color: #667eea;
}

@media (max-width: 768px) {
  .info-card {
    flex-direction: column;
  }
  .card-image-wrapper {
    width: 100%;
    height: 200px;
  }
}
</style>
