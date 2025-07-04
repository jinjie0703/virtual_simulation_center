<template>
  <header class="hero-carousel">
    <div class="carousel-inner">
      <div
        v-for="(slide, index) in slides"
        :key="index"
        class="carousel-item"
        :class="{ active: currentIndex === index }"
        :style="{ backgroundImage: `url(${slide.imageUrl})` }"
      >
        <div class="carousel-caption">
          <h1>{{ slide.title }}</h1>
        </div>
      </div>
    </div>

    <button @click="prevSlide" class="carousel-control prev">&lt;</button>
    <button @click="nextSlide" class="carousel-control next">&gt;</button>

    <ul class="carousel-indicators">
      <li
        v-for="(slide, index) in slides"
        :key="index"
        :class="{ active: currentIndex === index }"
        @click="goToSlide(index)"
      ></li>
    </ul>
  </header>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

// 通过 props 接收幻灯片数据
const props = defineProps({
  slides: {
    type: Array,
    required: true,
    default: () => [],
  },
})

const currentIndex = ref(0)
let intervalId = null

const nextSlide = () => {
  currentIndex.value = (currentIndex.value + 1) % props.slides.length
}

const prevSlide = () => {
  currentIndex.value = (currentIndex.value - 1 + props.slides.length) % props.slides.length
}

const goToSlide = (index) => {
  currentIndex.value = index
}

const startAutoplay = () => {
  intervalId = setInterval(() => {
    nextSlide()
  }, 5000) // 每5秒切换
}

const stopAutoplay = () => {
  clearInterval(intervalId)
}

// 组件挂载后开始自动播放
onMounted(() => {
  startAutoplay()
})

// 组件卸载前清除定时器，防止内存泄漏
onUnmounted(() => {
  stopAutoplay()
})
</script>

<style scoped>
/* 这里只放该组件相关的样式，保持组件独立性 */
.hero-carousel {
  position: relative;
  height: 100vh;
  min-height: 400px;
  width: 100%;
  overflow: hidden;
  /* 如果图片加载慢，显示黑色背景 */
  background-color: black;
  /* 文字颜色 */
  color: white;
}
/* 直接父容器 */
.carousel-inner {
  position: relative;
  width: 100%;
  height: 100%;
}
.carousel-item {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-size: cover;
  background-position: center;
  opacity: 0;
  transition: opacity 1s ease-in-out;
  display: flex;
  align-items: center;
  justify-content: center;
}
.carousel-item.active {
  opacity: 1;
  z-index: 1;
}
/* 暗色遮罩层，增加可读性 */
.carousel-item::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
}
.carousel-caption {
  position: relative;
  z-index: 2;
  text-align: center;
}
.carousel-caption h1 {
  font-size: 3rem;
  font-weight: 500;
  margin: 0;
  text-shadow: 2px 2px 8px rgba(0, 0, 0, 0.7);
}
.carousel-control {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  z-index: 3;
  /* 1. 设定固定的宽高，让它成为一个正方形 */
  width: 45px;
  height: 45px;
  padding: 0;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  background: rgba(0, 0, 0, 0.3);
  color: white;
  border: none;
  font-size: 24px;
  cursor: pointer;
  transition:
    background-color 0.3s,
    transform 0.3s ease;
}
.carousel-control:hover {
  background: rgba(0, 0, 0, 0.6);
}
.prev {
  left: 20px;
}
.next {
  right: 20px;
}
.carousel-indicators {
  position: absolute;
  bottom: 20px;
  left: 50%;
  transform: translateX(-50%);
  z-index: 3;
  display: flex;
  list-style: none;
  padding: 0;
  margin: 0;
}
.carousel-indicators li {
  width: 12px;
  height: 12px;
  background-color: rgba(255, 255, 255, 0.5);
  border-radius: 50%;
  margin: 0 5px;
  cursor: pointer;
  transition: background-color 0.3s;
}
.carousel-indicators li.active {
  background-color: white;
}
</style>
