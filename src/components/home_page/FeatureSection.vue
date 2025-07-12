<template>
  <section class="features-section">
    <div class="container">
      <div class="section-header">
        <h2 class="section-title">Highligts</h2>
        <p class="section-subtitle">Explore the results of our innovative projects</p>
      </div>

      <div class="features-grid">
        <div
          v-for="(feature, index) in features"
          :key="feature.id"
          :class="['feature-item', index % 2 === 0 ? 'layout-left' : 'layout-right']"
          :data-feature-id="feature.id"
          ref="featureElements"
        >
          <!-- 项目卡片 -->
          <div class="feature-content">
            <div
              class="project-card"
              @mouseenter="handleMouseEnter(index)"
              @mouseleave="handleMouseLeave(index)"
            >
              <div class="card-inner" :class="{ 'is-flipped': flippedCards[index] }">
                <!-- 正面 -->
                <div class="card-front">
                  <div class="project-header">
                    <h3 class="project-title">{{ feature.title }}</h3>
                    <div class="project-tags">
                      <span v-for="tag in feature.tags" :key="tag" class="tag">{{ tag }}</span>
                    </div>
                  </div>
                  <div class="project-description">
                    {{ feature.description }}
                  </div>
                </div>

                <!-- 背面 -->
                <div class="card-back">
                  <div class="author-section">
                    <div class="author-info">
                      <img
                        :src="feature.authorAvatar"
                        :alt="feature.author"
                        class="author-avatar"
                      />
                      <div class="author-details">
                        <h4 class="author-name">{{ feature.author }}</h4>
                        <p class="author-title">{{ feature.authorTitle }}</p>
                      </div>
                    </div>

                    <div class="contact-section">
                      <div class="contact-item">
                        <span class="contact-label">联系方式:</span>
                        <span class="contact-value"
                          >{{ feature.contact1 }}<br />{{ feature.contact2 }}</span
                        >
                      </div>
                      <div class="contact-item">
                        <span class="contact-label">项目地址:</span>
                        <a :href="feature.projectUrl || '#'" class="project-link" target="_blank">
                          {{ feature.projectUrl }}
                        </a>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- 媒体展示区域 -->
          <div class="feature-media">
            <div class="media-container">
              <!-- 主视频显示,视频地址，视频封面图片 -->
              <div class="main-video">
                <video
                  :src="feature.videoUrl"
                  :poster="feature.image"
                  controls
                  preload="metadata"
                  muted
                  playsinline
                  class="feature-video"
                  @error="handleVideoError"
                  @loadstart="handleVideoLoad"
                >
                  <source :src="feature.videoUrl" type="video/mp4" />
                  <source :src="feature.videoUrl" type="video/webm" />
                  <source :src="feature.videoUrl" type="video/ogg" />
                  您的浏览器不支持视频播放
                </video>
                <!-- 只在视频加载失败时显示播放按钮 -->
                <div v-if="!videoLoaded[feature.id]" class="video-overlay">
                  <button class="play-button" @click="playVideo(feature.id)">
                    <span class="icon-play"></span>
                  </button>
                </div>
              </div>

              <!-- 下方四张图片 -->
              <div class="thumbnail-grid">
                <div
                  v-for="(image, imgIndex) in feature.gallery.slice(0, 4)"
                  :key="imgIndex"
                  class="thumbnail"
                  @click="openImagePreview(image, feature.gallery)"
                >
                  <img :src="image" :alt="`项目图片 ${imgIndex + 1}`" />
                  <div class="thumbnail-overlay">
                    <span class="icon-zoom"></span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 图片预览模态框 -->
    <div v-if="showImagePreview" class="image-preview-modal" @click="closeImagePreview">
      <div class="preview-container" @click.stop>
        <img :src="previewImageUrl" :alt="'预览图片'" class="preview-image" />
        <button class="close-button" @click="closeImagePreview">
          <span>&times;</span>
        </button>
        <div class="image-navigation" v-if="currentGallery.length > 1">
          <button class="nav-button prev" @click.stop="previousImage">‹</button>
          <button class="nav-button next" @click.stop="nextImage">›</button>
        </div>
        <div class="image-counter" v-if="currentGallery.length > 1">
          {{ currentImageIndex + 1 }} / {{ currentGallery.length }}
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue'
//接受外部props(数据)
defineProps({
  features: {
    type: Array,
    required: true,
  },
})
// 用于收集所有DOM元素的ref，进入视口动画
const featureElements = ref([])
// 用于控制卡片翻转状态
const flippedCards = ref({})
// 用于记录视频加载状态，控制是否显示备用按钮
const videoLoaded = ref({})

// 图片预览相关状态
const showImagePreview = ref(false)
const previewImageUrl = ref('')
const currentGallery = ref([])
const currentImageIndex = ref(0)

// 用来保存交叉观察器实例的变量
let observer = null

// 卡片翻转控制
const handleMouseEnter = (index) => {
  flippedCards.value[index] = true
}

const handleMouseLeave = (index) => {
  flippedCards.value[index] = false
}

// 视频处理函数
const handleVideoError = (event) => {
  console.error('视频加载失败:', event.target.src)
}

const handleVideoLoad = (event) => {
  const videoElement = event.target
  const featureId = videoElement.closest('.feature-item').dataset.featureId
  if (featureId) {
    videoLoaded.value[featureId] = true
  }
}

const playVideo = (featureId) => {
  const videoElement = document.querySelector(`[data-feature-id="${featureId}"] video`)
  if (videoElement) {
    videoElement.play().catch((e) => {
      console.error('视频播放失败:', e)
    })
  }
}

// 图片预览功能
const openImagePreview = (imageUrl, gallery = []) => {
  if (Array.isArray(gallery) && gallery.length > 0) {
    currentGallery.value = gallery
    currentImageIndex.value = gallery.findIndex((img) => img === imageUrl) || 0
  } else {
    currentGallery.value = [imageUrl]
    currentImageIndex.value = 0
  }

  previewImageUrl.value = imageUrl
  showImagePreview.value = true

  // 防止页面滚动
  document.body.style.overflow = 'hidden'
}

const closeImagePreview = () => {
  showImagePreview.value = false
  previewImageUrl.value = ''
  currentGallery.value = []
  currentImageIndex.value = 0

  // 恢复页面滚动
  document.body.style.overflow = 'auto'
}

const previousImage = () => {
  if (currentGallery.value.length > 1) {
    currentImageIndex.value =
      currentImageIndex.value > 0 ? currentImageIndex.value - 1 : currentGallery.value.length - 1
    previewImageUrl.value = currentGallery.value[currentImageIndex.value]
  }
}

const nextImage = () => {
  if (currentGallery.value.length > 1) {
    currentImageIndex.value =
      currentImageIndex.value < currentGallery.value.length - 1 ? currentImageIndex.value + 1 : 0
    previewImageUrl.value = currentGallery.value[currentImageIndex.value]
  }
}

// 键盘事件处理
const handleKeyPress = (event) => {
  if (!showImagePreview.value) return

  switch (event.key) {
    case 'Escape':
      closeImagePreview()
      break
    case 'ArrowLeft':
      previousImage()
      break
    case 'ArrowRight':
      nextImage()
      break
  }
}

onMounted(() => {
  // 添加键盘事件监听
  document.addEventListener('keydown', handleKeyPress)

  if (!featureElements.value || featureElements.value.length === 0) return

  const options = {
    threshold: 0.1,
    rootMargin: '0px 0px -50px 0px',
  }

  observer = new IntersectionObserver((entries) => {
    entries.forEach((entry) => {
      if (entry.isIntersecting) {
        entry.target.classList.add('animate-in')
      }
    })
  }, options)

  featureElements.value.forEach((el) => {
    if (el) observer.observe(el)
  })
})

onBeforeUnmount(() => {
  // 移除键盘事件监听
  document.removeEventListener('keydown', handleKeyPress)

  // 恢复页面滚动
  document.body.style.overflow = 'auto'

  if (observer) {
    observer.disconnect()
  }
})
</script>

<style scoped>
/* 主容器样式 */
.features-section {
  padding: 80px 0;
  background-color: white;
  position: relative;
  min-height: 100vh;
}

.container {
  max-width: 1600px;
  margin: 0 auto;
  padding: 0 40px;
  position: relative;
  z-index: 1;
}

/* 标题区域 */
.section-header {
  text-align: center;
  margin-bottom: 50px;
}

.section-title {
  font-size: 3.5rem;
  font-weight: 700;
  color: black;
  margin-bottom: 12px;
}

.section-subtitle {
  font-size: 1.3rem;
  color: gray;
  max-width: 800px;
  margin: 0 auto;
}

/* 网格布局 */
.features-grid {
  display: flex;
  flex-direction: column;
  gap: 100px;
}

.feature-item {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 50px;
  align-items: center;
  min-height: 600px;
  opacity: 0;
  transform: translateY(50px);
  transition: all 0.8s ease;
}

.feature-item.layout-left {
  grid-template-areas: 'content media';
}

.feature-item.layout-right {
  grid-template-areas: 'media content';
}

.feature-content {
  grid-area: content;
}

.feature-media {
  grid-area: media;
}

/* 项目卡片样式 */
.project-card {
  perspective: 2000px;
  height: 540px;
  width: 100%;
  max-width: 450px;
  margin: 0 auto;
}

.card-inner {
  position: relative;
  width: 100%;
  height: 100%;
  text-align: center;
  transition: transform 0.6s;
  transform-style: preserve-3d;
  cursor: pointer;
}

.card-inner.is-flipped {
  transform: rotateY(180deg);
}

.card-front,
.card-back {
  position: absolute;
  width: 100%;
  height: 100%;
  backface-visibility: hidden;
  background: white;
  border-radius: 24px;
  padding: 40px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  box-sizing: border-box;
}
/* 背面旋转180度 */
.card-back {
  transform: rotateY(180deg);
}

/* 正面内容 */
.project-header {
  margin-bottom: 20px;
  flex-shrink: 0;
}

.project-title {
  font-size: 2.4rem;
  font-weight: 600;
  color: black;
  margin-bottom: 10px;
}

.project-tags {
  display: flex;
  gap: 5px;
  flex-wrap: wrap;
  justify-content: center;
}

.tag {
  background: linear-gradient(135deg, purple 0%, white 100%);
  color: black;
  padding: 6px 14px;
  border-radius: 16px;
  font-size: 0.9rem;
  font-weight: 500;
}

.project-description {
  font-size: 1.2rem;
  line-height: 1.8;
  color: gray;
  text-align: left;
  flex: 1;
  display: flex;
  align-items: start;
}

/* 背面内容 */
.author-section {
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.author-info {
  display: flex;
  align-items: center;
  margin-bottom: 30px;
  flex-shrink: 0;
}

.author-avatar {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  object-fit: cover;
  margin-right: 20px;
  border: 4px solid #e2e8f0;
}

.author-details {
  text-align: left;
}

.author-name {
  font-size: 1.5rem;
  font-weight: 600;
  color: black;
  margin-bottom: 5px;
}

.author-title {
  font-size: 1rem;
  color: grey;
}

.contact-section {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.contact-item {
  margin-bottom: 20px;
  text-align: left;
}

.contact-label {
  display: block;
  font-weight: 600;
  color: #2d3748;
  margin-bottom: 8px;
}

.contact-value {
  color: #4a5568;
  font-size: 1rem;
  word-break: break-all;
}

.project-link {
  color: #667eea;
  text-decoration: none;
  font-weight: 500;
  transition: color 0.3s ease;
}

.project-link:hover {
  color: #764ba2;
  text-decoration: underline;
}

/* 媒体展示区域 */
.media-container {
  width: 100%;
}

.main-video {
  position: relative;
  width: 100%;
  height: 350px;
  margin-bottom: 20px;
  border-radius: 16px;
  overflow: hidden;
  box-shadow: 0 15px 40px rgba(0, 0, 0, 0.1);
}

.feature-video {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 16px;
}

.video-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(0, 0, 0, 0.3);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.main-video:hover .video-overlay {
  opacity: 1;
}

.play-button {
  width: 80px;
  height: 80px;
  border: none;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.9);
  color: #667eea;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.play-button:hover {
  transform: scale(1.1);
  background: white;
}

/* 缩略图网格 */
.thumbnail-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 15px;
}

.thumbnail {
  position: relative;
  aspect-ratio: 1;
  border-radius: 12px;
  overflow: hidden;
  cursor: pointer;
  transition: transform 0.3s ease;
}

.thumbnail:hover {
  transform: scale(1.05);
}

.thumbnail img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.thumbnail-overlay {
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

.thumbnail:hover .thumbnail-overlay {
  opacity: 1;
}

/* 图标样式 */
.icon-play::before,
.icon-zoom::before {
  font-size: 24px;
  color: white;
}

.icon-play::before {
  content: '▶';
}

.icon-zoom::before {
  content: '🔍';
}

/* 图片预览模态框样式 */
.image-preview-modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: rgba(0, 0, 0, 0.9);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 9999;
  cursor: pointer;
}

.preview-container {
  position: relative;
  width: 80vw; /* 固定宽度 */
  height: 80vh; /* 固定高度 */
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: default;
}

.preview-image {
  max-width: 100%;
  max-height: 100%;
  width: auto;
  height: auto;
  object-fit: contain; /* 保持图片比例，完整显示 */
  border-radius: 8px;
  box-shadow: 0 20px 60px rgba(255, 255, 255, 0.1);
}

.close-button {
  position: absolute;
  top: -50px;
  right: -10px;
  background: rgba(0, 0, 0, 0.5);
  border: none;
  color: white;
  font-size: 2rem;
  cursor: pointer;
  padding: 10px 15px;
  border-radius: 50%;
  transition: all 0.3s ease;
  z-index: 10001;
}

.close-button:hover {
  background: rgba(0, 0, 0, 0.8);
  transform: scale(1.1);
}

.image-navigation {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  width: calc(100% + 100px); /* 扩展导航按钮区域 */
  left: -50px;
  display: flex;
  justify-content: space-between;
  pointer-events: none;
}

.nav-button {
  background: rgba(0, 0, 0, 0.5);
  border: none;
  color: white;
  font-size: 2rem;
  padding: 15px 20px;
  cursor: pointer;
  border-radius: 50%;
  transition: all 0.3s ease;
  pointer-events: auto;
  z-index: 10001;
}

.nav-button:hover {
  background: rgba(0, 0, 0, 0.8);
  transform: scale(1.1);
}

.image-counter {
  position: absolute;
  bottom: -60px;
  left: 50%;
  transform: translateX(-50%);
  color: white;
  font-size: 1.1rem;
  background: rgba(0, 0, 0, 0.7);
  padding: 10px 20px;
  border-radius: 25px;
  font-weight: 500;
}

/* 动画效果 */
.feature-item.animate-in {
  opacity: 1;
  transform: translateY(0);
}

/* 响应式设计 */
@media (max-width: 1024px) {
  .feature-item {
    grid-template-columns: 1fr;
    gap: 40px;
  }

  .feature-item.layout-left,
  .feature-item.layout-right {
    grid-template-areas: 'content' 'media';
  }

  .project-card {
    height: 450px;
    max-width: 400px;
  }
}

@media (max-width: 768px) {
  .preview-container {
    width: 90vw;
    height: 70vh;
  }

  .nav-button {
    font-size: 1.5rem;
    padding: 12px 16px;
  }

  .close-button {
    font-size: 1.5rem;
    top: -40px;
    right: -5px;
    padding: 8px 12px;
  }

  .image-counter {
    bottom: -50px;
    font-size: 1rem;
    padding: 8px 16px;
  }
}

@media (max-width: 480px) {
  .preview-container {
    width: 95vw;
    height: 60vh;
  }

  .image-navigation {
    width: calc(100% + 60px);
    left: -30px;
  }

  .nav-button {
    font-size: 1.2rem;
    padding: 10px 14px;
  }
}
</style>
