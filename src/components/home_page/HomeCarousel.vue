<template>
  <section class="hero-carousel">
    <div v-if="loading" class="carousel-loading">加载中...</div>
    <div v-else-if="error" class="carousel-error">{{ error }}</div>
    <div v-else-if="slides.length > 0" class="carousel-container">
      <!-- 轮播图片 -->
      <div class="carousel-inner">
        <div
          v-for="(slide, index) in slides"
          :key="slide.ID"
          class="carousel-slide"
          :class="{ active: index === currentIndex }"
        >
          <img :src="getImageUrl(slide.imageUrl)" :alt="slide.title" class="slide-image" />
          <div class="slide-overlay"></div>
        </div>
      </div>

      <!-- 左右切换按钮 - 调小尺寸 -->
      <button class="carousel-control prev" @click="prevSlide" aria-label="上一张">
        <span>‹</span>
      </button>
      <button class="carousel-control next" @click="nextSlide" aria-label="下一张">
        <span>›</span>
      </button>

      <!-- 中心内容区域 -->
      <div class="carousel-caption">
        <h1>{{ currentSlideTitle }}</h1>
        <p class="carousel-subtitle">{{ currentSlideSubtitle }}</p>
      </div>

      <!-- 指示器 - 移动到底部 -->
      <div class="carousel-indicators">
        <button
          v-for="(slide, index) in slides"
          :key="index"
          class="indicator"
          :class="{ active: index === currentIndex }"
          @click="goToSlide(index)"
          :aria-label="`切换到第${index + 1}张`"
        ></button>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, onMounted, onUnmounted, computed } from 'vue'
import axios from 'axios'

const currentIndex = ref(0)
const slides = ref([])
let autoplayTimer = null
const loading = ref(true)
const error = ref(null)

const API_BASE_URL = 'https://localhost:8080'

const currentSlideTitle = computed(() => {
  if (slides.value.length > 0 && slides.value[currentIndex.value]) {
    return slides.value[currentIndex.value].title
  }
  return '加载中...'
})

const currentSlideSubtitle = computed(() => {
  if (slides.value.length > 0 && slides.value[currentIndex.value]) {
    return slides.value[currentIndex.value].subtitle
  }
  return '请稍候...'
})

const getImageUrl = (imageName) => {
  if (!imageName) return ''
  return `${API_BASE_URL}/static/images/home_page/HomeCarousel/${imageName}`
}

// 轮播数据
// const slides = ref([
//   {
//     title: '虚拟仿真实验平台',
//     subtitle: '构建沉浸式学习环境，开启教育新篇章',
//     imageUrl:
//       'https://images.unsplash.com/photo-1542370285-b8eb8317691c?q=80&w=1974&auto=format&fit=crop',
//   },
//   {
//     title: '智能化教学系统',
//     subtitle: 'AI驱动的个性化学习体验',
//     imageUrl:
//       'https://images.unsplash.com/photo-1519010470956-6d877008eaa4?q=80&w=2070&auto=format&fit=crop',
//   },
//   {
//     title: '数字化实验室',
//     subtitle: '打破时空限制，让实验无处不在',
//     imageUrl:
//       'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//   },
// ])

const nextSlide = () => {
  currentIndex.value = (currentIndex.value + 1) % slides.value.length
}

const prevSlide = () => {
  currentIndex.value = currentIndex.value === 0 ? slides.value.length - 1 : currentIndex.value - 1
}

const goToSlide = (index) => {
  currentIndex.value = index
}
// 自动播放
const startAutoplay = () => {
  autoplayTimer = setInterval(() => {
    nextSlide()
  }, 5000)
}

const stopAutoplay = () => {
  if (autoplayTimer) {
    clearInterval(autoplayTimer)
    autoplayTimer = null
  }
}

onMounted(async () => {
  try {
    const response = await axios.get(`${API_BASE_URL}/api/home_page_carousel`)
    slides.value = response.data
    if (slides.value.length > 0) {
      startAutoplay()
    }
  } catch {
    error.value = '加载轮播数据失败，请稍后再试。'
  } finally {
    loading.value = false
  }
})

onUnmounted(() => {
  stopAutoplay()
})
</script>

<style scoped>
.hero-carousel {
  position: relative;
  width: 100%;
  height: 100vh;
  overflow: hidden;
  background: #000;
}

.carousel-container {
  position: relative;
  width: 100%;
  height: 100%;
}

.carousel-inner {
  position: relative;
  width: 100%;
  height: 100%;
}

.carousel-slide {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  opacity: 0;
  transition: opacity 1s ease-in-out;
}

.carousel-slide.active {
  opacity: 1;
}

.slide-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.slide-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(
    45deg,
    rgba(0, 0, 0, 0.6) 0%,
    rgba(0, 0, 0, 0.3) 50%,
    rgba(0, 0, 0, 0.6) 100%
  );
}

/* 中心内容区域 */
.carousel-caption {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  text-align: center;
  color: white;
  z-index: 10;
  max-width: 900px;
  padding: 0 40px;
}

.carousel-caption h1 {
  font-size: 4.5rem;
  font-weight: 700;
  text-shadow: 2px 2px 8px rgba(0, 0, 0, 0.8);
  margin: 0 0 24px 0;
  line-height: 1.2;
  letter-spacing: -0.02em;
}

.carousel-subtitle {
  font-size: 1.8rem;
  font-weight: 400;
  text-shadow: 1px 1px 4px rgba(0, 0, 0, 0.7);
  margin: 0;
  opacity: 0.95;
  line-height: 1.4;
}

/* 左右切换按钮 - 缩小尺寸 */
.carousel-control {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  background: rgba(255, 255, 255, 0.15);
  color: white;
  border: 1px solid rgba(255, 255, 255, 0.25);
  width: 45px;
  height: 45px;
  border-radius: 50%;
  cursor: pointer;
  z-index: 10;
  font-size: 18px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
  backdrop-filter: blur(10px);
}

.carousel-control:hover {
  background: rgba(255, 255, 255, 0.25);
  border-color: rgba(255, 255, 255, 0.4);
  transform: translateY(-50%) scale(1.05);
}

.carousel-control.prev {
  left: 30px;
}

.carousel-control.next {
  right: 30px;
}

.carousel-control span {
  font-size: 20px;
  font-weight: 300;
  line-height: 1;
}

/* 指示器 - 移动到底部 */
.carousel-indicators {
  position: absolute;
  bottom: 40px;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  gap: 12px;
  z-index: 10;
}

.indicator {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  border: 2px solid rgba(255, 255, 255, 0.6);
  background: transparent;
  cursor: pointer;
  transition: all 0.3s ease;
  outline: none;
  box-sizing: border-box;
  padding: 0;
  margin: 0;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
  display: flex;
  align-items: center;
  justify-content: center;
}

.indicator.active {
  background: white;
  border-color: white;
  transform: scale(1.2);
}

.indicator:hover {
  border-color: rgba(255, 255, 255, 0.8);
  transform: scale(1.1);
}

/* 响应式设计 */
@media (max-width: 1024px) {
  .carousel-caption h1 {
    font-size: 3.5rem;
  }

  .carousel-subtitle {
    font-size: 1.5rem;
  }

  .carousel-control {
    width: 40px;
    height: 40px;
    font-size: 16px;
  }

  .carousel-control.prev {
    left: 25px;
  }

  .carousel-control.next {
    right: 25px;
  }
}

@media (max-width: 768px) {
  .carousel-caption {
    padding: 0 30px;
  }

  .carousel-caption h1 {
    font-size: 2.8rem;
    margin-bottom: 20px;
  }

  .carousel-subtitle {
    font-size: 1.3rem;
  }

  .carousel-control {
    width: 36px;
    height: 36px;
    font-size: 14px;
  }

  .carousel-control span {
    font-size: 18px;
  }

  .carousel-control.prev {
    left: 20px;
  }

  .carousel-control.next {
    right: 20px;
  }

  .carousel-indicators {
    bottom: 30px;
    gap: 10px;
  }

  .indicator {
    width: 10px;
    height: 10px;
  }
}

@media (max-width: 480px) {
  .carousel-caption {
    padding: 0 20px;
  }

  .carousel-caption h1 {
    font-size: 2.2rem;
    margin-bottom: 16px;
  }

  .carousel-subtitle {
    font-size: 1.1rem;
  }

  .carousel-control {
    width: 32px;
    height: 32px;
    font-size: 12px;
  }

  .carousel-control span {
    font-size: 16px;
  }

  .carousel-control.prev {
    left: 15px;
  }

  .carousel-control.next {
    right: 15px;
  }

  .carousel-indicators {
    bottom: 25px;
    gap: 8px;
  }

  .indicator {
    width: 8px;
    height: 8px;
  }
}
</style>
