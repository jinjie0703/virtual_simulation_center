<template>
  <div class="achievement-card" tabindex="0">
    <div class="image-container">
      <img :src="achievement.image" :alt="achievement.title" class="card-image" />
      <div class="overlay">
        <button class="details-btn">查看详情</button>
      </div>
    </div>
    <div class="content-container">
      <div class="tags">
        <span v-for="tag in achievement.tags" :key="tag" class="tag">{{ tag }}</span>
      </div>
      <h3 class="title">{{ achievement.title }}</h3>
      <div class="author-info">
        <img :src="achievement.authorAvatar" :alt="achievement.author" class="author-avatar" />
        <div class="author-details">
          <span class="author-name">{{ achievement.author }}</span>
          <span class="author-title">{{ achievement.authorTitle }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  achievement: Object,
})
</script>

<style scoped>
.achievement-card {
  background-color: #fff;
  border-radius: 16px;
  overflow: hidden;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.05);
  cursor: pointer;
  transition:
    transform 0.3s ease,
    box-shadow 0.3s ease;
  display: flex;
  flex-direction: column;
}
.achievement-card:hover,
.achievement-card:focus-within {
  transform: translateY(-10px);
  box-shadow: 0 12px 24px rgba(102, 126, 234, 0.2);
}
.image-container {
  position: relative;
  height: 240px;
}
.card-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.4s ease;
}
.achievement-card:hover .card-image {
  transform: scale(1.05);
}
.overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity 0.3s ease;
}
.achievement-card:hover .overlay {
  opacity: 1;
}
.details-btn {
  padding: 12px 24px;
  background: #fff;
  color: #333;
  border: none;
  border-radius: 50px;
  font-weight: 600;
  cursor: pointer;
  transform: translateY(20px);
  transition:
    transform 0.3s ease,
    background-color 0.3s ease;
}
.achievement-card:hover .details-btn {
  transform: translateY(0);
}
.details-btn:hover {
  background-color: #667eea;
  color: #fff;
}
.content-container {
  padding: 25px;
  flex-grow: 1;
  display: flex;
  flex-direction: column;
}
.tags {
  display: flex;
  gap: 8px;
  margin-bottom: 15px;
}
.tag {
  background: #edf2f7;
  color: #4a5568;
  padding: 5px 12px;
  border-radius: 50px;
  font-size: 0.75rem;
  font-weight: 500;
}
.title {
  font-size: 1.3rem;
  font-weight: 700;
  color: #2d3748;
  margin: 0 0 20px;
  line-height: 1.4;
  flex-grow: 1;
}
.author-info {
  display: flex;
  align-items: center;
  margin-top: auto;
  padding-top: 20px;
  border-top: 1px solid #e2e8f0;
}
.author-avatar {
  width: 45px;
  height: 45px;
  border-radius: 50%;
  margin-right: 15px;
  object-fit: cover;
}
.author-details {
  display: flex;
  flex-direction: column;
}
.author-name {
  font-weight: 600;
  color: #2d3748;
}
.author-title {
  font-size: 0.85rem;
  color: #718096;
}
</style>
