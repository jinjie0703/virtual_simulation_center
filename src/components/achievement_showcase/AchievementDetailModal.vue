<template>
  <transition name="modal-fade">
    <div v-if="isVisible" class="modal-backdrop" @click.self="$emit('close')">
      <div class="modal-container">
        <button class="close-btn" @click="$emit('close')">&times;</button>
        <div v-if="achievement" class="modal-content">
          <div class="media-column">
            <div class="video-wrapper">
              <video :src="achievement.videoUrl" controls playsinline class="main-video"></video>
            </div>
            <div class="gallery">
              <h4>项目掠影</h4>
              <div class="gallery-grid">
                <img
                  v-for="(img, index) in achievement.gallery"
                  :key="index"
                  :src="img"
                  alt="Gallery image"
                  class="gallery-img"
                />
              </div>
            </div>
          </div>
          <div class="info-column">
            <div class="tags">
              <span v-for="tag in achievement.tags" :key="tag" class="tag">{{ tag }}</span>
            </div>
            <h2 class="title">{{ achievement.title }}</h2>
            <p class="description">{{ achievement.description }}</p>
            <div class="author-card">
              <img
                :src="achievement.authorAvatar"
                :alt="achievement.author"
                class="author-avatar"
              />
              <div class="author-details">
                <span class="author-name">{{ achievement.author }}</span>
                <span class="author-title">{{ achievement.authorTitle }}</span>
                <a :href="'mailto:' + achievement.contact" class="author-contact">联系我们</a>
              </div>
            </div>
            <a
              :href="achievement.projectUrl"
              target="_blank"
              rel="noopener noreferrer"
              class="project-link"
            >
              <span>查看项目源码</span>
              <span>&rarr;</span>
            </a>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script setup>
defineProps({
  achievement: Object,
  isVisible: Boolean,
})
defineEmits(['close'])
</script>

<style scoped>
.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: opacity 0.3s ease;
}
.modal-fade-enter-from,
.modal-fade-leave-to {
  opacity: 0;
}
.modal-fade-enter-active .modal-container,
.modal-fade-leave-active .modal-container {
  transition: transform 0.3s ease;
}
.modal-fade-enter-from .modal-container,
.modal-fade-leave-to .modal-container {
  transform: scale(0.95);
}
.modal-backdrop {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.6);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
  overflow-y: auto;
  padding: 20px;
}
.modal-container {
  background: #fff;
  border-radius: 16px;
  width: 90%;
  max-width: 1200px;
  max-height: 90vh;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
  position: relative;
  display: flex;
  overflow: hidden;
}
.close-btn {
  position: absolute;
  top: 15px;
  right: 15px;
  width: 40px;
  height: 40px;
  border: none;
  background: rgba(0, 0, 0, 0.1);
  color: #fff;
  border-radius: 50%;
  font-size: 24px;
  cursor: pointer;
  z-index: 10;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background-color 0.2s;
}
.close-btn:hover {
  background: rgba(0, 0, 0, 0.3);
}
.modal-content {
  display: flex;
  width: 100%;
  max-height: 90vh;
}
.media-column,
.info-column {
  padding: 40px;
  overflow-y: auto;
  height: 90vh;
}
.media-column {
  flex: 6;
  background: #f8f9fa;
}
.info-column {
  flex: 4;
}
.video-wrapper {
  width: 100%;
  margin-bottom: 30px;
  border-radius: 12px;
  overflow: hidden;
}
.main-video {
  width: 100%;
  display: block;
}
.gallery h4 {
  font-size: 1.2rem;
  color: #343a40;
  margin-bottom: 15px;
}
.gallery-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 10px;
}
.gallery-img {
  width: 100%;
  height: 150px;
  object-fit: cover;
  border-radius: 8px;
  cursor: pointer;
  transition: transform 0.2s;
}
.gallery-img:hover {
  transform: scale(1.05);
}
.tags {
  display: flex;
  gap: 8px;
  margin-bottom: 15px;
}
.tag {
  background: #e9ecef;
  color: #495057;
  padding: 5px 12px;
  border-radius: 50px;
  font-size: 0.8rem;
  font-weight: 500;
}
.title {
  font-size: 2.2rem;
  font-weight: 700;
  color: #212529;
  margin-bottom: 20px;
}
.description {
  font-size: 1rem;
  line-height: 1.7;
  color: #495057;
  margin-bottom: 30px;
}
.author-card {
  background: #f8f9fa;
  border-radius: 12px;
  padding: 20px;
  display: flex;
  align-items: center;
  margin-bottom: 30px;
}
.author-avatar {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  margin-right: 20px;
}
.author-details {
  line-height: 1.4;
}
.author-name {
  font-weight: 600;
  font-size: 1.1rem;
  color: #212529;
  display: block;
}
.author-title {
  color: #6c757d;
  font-size: 0.9rem;
}
.author-contact {
  color: #667eea;
  font-weight: 600;
  text-decoration: none;
  font-size: 0.9rem;
  margin-top: 5px;
  display: inline-block;
}
.project-link {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 18px 25px;
  background: #667eea;
  color: #fff;
  text-decoration: none;
  border-radius: 12px;
  font-weight: 600;
  transition:
    background-color 0.2s,
    transform 0.2s;
}
.project-link:hover {
  background-color: #5a67d8;
  transform: translateY(-2px);
}
@media (max-width: 992px) {
  .modal-content {
    flex-direction: column;
  }
  .media-column,
  .info-column {
    flex: 1;
    height: auto;
    max-height: 50vh;
  }
}
</style>
