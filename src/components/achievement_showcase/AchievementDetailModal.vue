<template>
  <transition name="modal-fade">
    <div v-if="isVisible" class="modal-backdrop" @click.self="$emit('close')">
      <div class="modal-container">
        <button class="close-btn" @click="$emit('close')">&times;</button>
        <div v-if="achievement" class="modal-content">
          <div class="media-column">
            <!-- Video Player -->
            <div class="video-wrapper">
              <video
                :ref="(el) => (videoRefs[achievement.id] = el)"
                :src="achievement.videoUrl"
                playsinline
                class="main-video"
                :poster="achievement.image"
                @play="onVideoPlay(achievement.id)"
                @pause="onVideoPause(achievement.id)"
              ></video>
              <div v-show="!isPlaying[achievement.id]" class="video-overlay">
                <button class="play-button" @click="playVideo(achievement.id)">
                  <img
                    src="@/assets/home_page/FeatureSection/video_play.svg"
                    alt="Play"
                    class="play-icon"
                  />
                </button>
              </div>
            </div>

            <!-- Image Gallery -->
            <div class="gallery">
              <h4>项目掠影</h4>
              <div class="gallery-grid">
                <div
                  v-for="(img, index) in achievement.gallery"
                  :key="index"
                  class="thumbnail"
                  @click="openImagePreview(img, achievement.gallery)"
                >
                  <img :src="img" alt="Gallery image" class="gallery-img" />
                  <div class="thumbnail-overlay">
                    <img
                      src="@/assets/home_page/FeatureSection/image_enhance.svg"
                      alt="Zoom"
                      class="zoom-icon"
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Info Column -->
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
              </div>
            </div>
            <!-- Team Members Section -->
            <div
              class="team-members-section"
              v-if="achievement.teamMembers && achievement.teamMembers.length"
            >
              <h4>团队成员</h4>
              <div class="members-grid">
                <div
                  v-for="member in achievement.teamMembers"
                  :key="member.name"
                  class="member-card"
                >
                  <img :src="member.avatar" :alt="member.name" class="member-avatar" />
                  <div class="member-details">
                    <span class="member-name">{{ member.name }}</span>
                    <span class="member-role">{{ member.role }}</span>
                  </div>
                </div>
              </div>
            </div>

            <div class="contact-section">
              <div class="contact-item" v-if="achievement.contact1">
                <span class="contact-label">联系方式:</span>
                <span class="contact-value"
                  >{{ achievement.contact1 }}<br />{{ achievement.contact2 }}</span
                >
              </div>
              <div class="contact-item">
                <span class="contact-label">项目地址:</span>
                <a :href="achievement.projectUrl || '#'" class="project-link" target="_blank">
                  {{ achievement.projectUrl }}
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </transition>

  <!-- Image Preview Modal -->
  <div v-if="showImagePreview" class="image-preview-modal" @click="closeImagePreview">
    <div class="preview-container" @click.stop>
      <img :src="previewImageUrl" :alt="'预览图片'" class="preview-image" />
      <button class="preview-close-button" @click="closeImagePreview">
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
</template>

<script setup>
import { ref, watch, onMounted, onBeforeUnmount } from 'vue'

const props = defineProps({
  achievement: Object,
  isVisible: Boolean,
})

defineEmits(['close'])

// Video state
const videoRefs = ref({})
const isPlaying = ref({})

// Image preview state
const showImagePreview = ref(false)
const previewImageUrl = ref('')
const currentGallery = ref([])
const currentImageIndex = ref(0)

// Stop video when modal closes
watch(
  () => props.isVisible,
  (newVal) => {
    if (!newVal && props.achievement && videoRefs.value[props.achievement.id]) {
      const video = videoRefs.value[props.achievement.id]
      if (!video.paused) {
        video.pause()
      }
    }
  },
)

// Video controls
const playVideo = (id) => {
  const video = videoRefs.value[id]
  if (video) {
    video.play().catch((e) => console.error('Video play failed:', e))
  }
}
const onVideoPlay = (id) => {
  isPlaying.value[id] = true
}
const onVideoPause = (id) => {
  isPlaying.value[id] = false
}

// Image preview controls
const openImagePreview = (imageUrl, gallery) => {
  currentGallery.value = gallery
  currentImageIndex.value = gallery.findIndex((img) => img === imageUrl)
  previewImageUrl.value = imageUrl
  showImagePreview.value = true
  document.body.style.overflow = 'hidden' // Prevent background scroll
}

const closeImagePreview = () => {
  showImagePreview.value = false
  document.body.style.overflow = 'auto' // Restore scroll
}

const previousImage = () => {
  currentImageIndex.value =
    currentImageIndex.value > 0 ? currentImageIndex.value - 1 : currentGallery.value.length - 1
  previewImageUrl.value = currentGallery.value[currentImageIndex.value]
}

const nextImage = () => {
  currentImageIndex.value =
    currentImageIndex.value < currentGallery.value.length - 1 ? currentImageIndex.value + 1 : 0
  previewImageUrl.value = currentGallery.value[currentImageIndex.value]
}

const handleKeyPress = (event) => {
  if (!showImagePreview.value) return
  if (event.key === 'Escape') {
    closeImagePreview()
  } else if (event.key === 'ArrowLeft') {
    previousImage()
  } else if (event.key === 'ArrowRight') {
    nextImage()
  }
}

onMounted(() => {
  document.addEventListener('keydown', handleKeyPress)
})

onBeforeUnmount(() => {
  document.removeEventListener('keydown', handleKeyPress)
})
</script>

<style scoped>
/* Modal Fade Transition */
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

/* Main Modal Styles */
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
  top: 40px;
  right: 40px;
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

/* Video Player Styles */
.video-wrapper {
  position: relative;
  width: 100%;
  margin-bottom: 30px;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
  aspect-ratio: 16 / 9; /* Fixed aspect ratio */
}
.main-video {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
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
  transition: opacity 0.3s ease;
}
.play-button {
  width: 80px;
  height: 80px;
  border: none;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.2);
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
.play-icon {
  width: 50px;
  height: 50px;
}

/* Gallery Styles */
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
.thumbnail {
  position: relative;
  border-radius: 8px;
  overflow: hidden;
  cursor: pointer;
  height: 150px;
}
.gallery-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.2s;
}
.thumbnail:hover .gallery-img {
  transform: scale(1.05);
}
.thumbnail-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.3);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity 0.3s ease;
}
.thumbnail:hover .thumbnail-overlay {
  opacity: 1;
}
.zoom-icon {
  width: 40px;
  height: 40px;
}

/* Info Column Styles */
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
  display: flex;
  align-items: center;
  margin-bottom: 20px;
}
.author-avatar {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  margin-right: 20px;
  border: 3px solid #e2e8f0;
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

/* Team Members Section */
.team-members-section {
  margin-bottom: 20px;
  padding-bottom: 20px;
  border-bottom: 1px solid #e2e8f0;
}
.team-members-section h4 {
  font-size: 1.2rem;
  color: #343a40;
  margin-bottom: 15px;
}
.members-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(120px, 1fr));
  gap: 20px;
}
.member-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
}
.member-avatar {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  margin-bottom: 10px;
  object-fit: cover;
  border: 2px solid #e2e8f0;
}
.member-details {
  line-height: 1.4;
}
.member-name {
  font-weight: 600;
  font-size: 0.9rem;
  color: #212529;
}
.member-role {
  font-size: 0.8rem;
  color: #6c757d;
  display: block;
}

/* Contact Section Styles from FeatureSection */
.contact-section {
  margin-top: 20px;
  padding-top: 20px;
  border-top: 1px solid #e2e8f0;
}
.contact-item {
  margin-bottom: 15px;
}
.contact-label {
  display: block;
  font-weight: 600;
  color: #334155;
  margin-bottom: 5px;
  font-size: 0.9rem;
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
  word-break: break-all;
}
.project-link:hover {
  color: #5a67d8;
  text-decoration: underline;
}

/* Image Preview Modal Styles */
.image-preview-modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: rgba(0, 0, 0, 0.8);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 9999;
  cursor: pointer;
}
.preview-container {
  position: relative;
  width: 80vw;
  height: 80vh;
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
  object-fit: contain;
  border-radius: 8px;
}
.preview-close-button {
  position: absolute;
  top: 0px;
  right: -40px;
  background: rgba(0, 0, 0, 0.5);
  border: none;
  color: white;
  font-size: 2rem;
  cursor: pointer;
  padding: 10px 15px;
  border-radius: 50%;
  transition: all 0.3s ease;
}
.preview-close-button:hover {
  background: rgba(0, 0, 0, 0.8);
  transform: scale(1.1);
}
.image-navigation {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  width: calc(100% + 100px);
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
}

/* Responsive */
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
