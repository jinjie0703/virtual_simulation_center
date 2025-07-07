<template>
  <section class="features-section">
    <div class="container">
      <h2>我们的成果</h2>
      <div
        v-for="(feature, index) in features"
        :key="feature.id || index"
        class="feature-item"
        :class="{ 'layout-text-left': index % 2 !== 0 }"
        ref="featureElements"
      >
        <!-- 文字区域 -->
        <div class="feature-text">
          <div class="card-container">
            <div class="card-front">
              <h3>{{ feature.title }}</h3>
              <p class="project-description">{{ feature.description }}</p>
            </div>
            <div class="card-back">
              <div class="author-info">
                <p><strong>作者：</strong>{{ feature.author }}</p>
                <p><strong>联系方式：</strong>{{ feature.contact }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- 图片区域 -->
        <div class="feature-image">
          <img :src="feature.imageUrl" :alt="feature.title" />
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue'

defineProps({
  features: {
    type: Array,
    required: true,
  },
})

const featureElements = ref([])
let observer = null

onMounted(() => {
  if (!featureElements.value || featureElements.value.length === 0) return

  const options = {
    root: null,
    rootMargin: '0px',
    threshold: 0.3,
  }

  observer = new IntersectionObserver((entries) => {
    entries.forEach((entry) => {
      if (entry.isIntersecting) {
        entry.target.classList.add('is-visible')
        observer.unobserve(entry.target)
      }
    })
  }, options)

  featureElements.value.forEach((el) => {
    if (el) observer.observe(el)
  })
})

onBeforeUnmount(() => {
  if (observer) {
    observer.disconnect()
  }
})
</script>

<style scoped>
.features-section {
  padding-top: 80px;
  background-color: #fdfdfd;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

.features-section h2 {
  text-align: center;
  font-size: 2.8rem;
  margin-bottom: 80px;
  color: #333;
}

.feature-item {
  display: flex;
  align-items: center;
  justify-content: space-evenly; /* 使用 evenly 分布，两边留有空间 */

  /* 核心改动：严格设置行高为半屏 */
  height: 50vh;

  /* 默认布局：图片在左，文字在右 */
  flex-direction: row-reverse;

  /* 动画 */
  opacity: 0;
  transform: translateY(50px);
  transition:
    opacity 0.7s ease-out,
    transform 0.7s ease-out;
}

.feature-item:nth-child(even) {
  background-color: #f7faff;
}

.feature-item.is-visible {
  opacity: 1;
  transform: translateY(0);
}

/* 交替布局：文字在左，图片在右 */
.feature-item.layout-text-left {
  flex-direction: row;
}

.feature-image {
  flex-basis: 50%;
  max-width: 520px;
  /* 核心改动：让图片容器高度占满行高 */
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 3vh 0; /* 在容器内部增加垂直 padding，给图片呼吸空间 */
}

.feature-text {
  flex-basis: 40%;
  max-width: 450px;
}

.feature-image img {
  /* 核心改动：让图片自适应容器，不变形 */
  max-width: 100%;
  max-height: 100%;
  object-fit: contain; /* 关键属性！ */

  border-radius: 12px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
}

/* --- 卡片翻转效果（保持不变） --- */
.card-container {
  position: relative;
  perspective: 1200px;
  cursor: pointer;
}
.card-front,
.card-back {
  width: 100%;
  min-height: 280px;
  padding: 30px;
  border-radius: 12px;
  background-color: #ffffff;
  border: 1px solid #e8e8e8;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.07);
  transition: transform 0.7s cubic-bezier(0.4, 0.2, 0.2, 1);
  backface-visibility: hidden;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.card-back {
  position: absolute;
  top: 0;
  left: 0;
  transform: rotateY(180deg);
}
.card-container:hover .card-front {
  transform: rotateY(-180deg);
}
.card-container:hover .card-back {
  transform: rotateY(0deg);
}
.card-front h3 {
  font-size: 1.8rem;
  margin: 0 0 15px 0;
  color: #0056b3;
}
.card-front .project-description {
  color: #555;
  line-height: 1.7;
}

/* --- 响应式设计：平板和手机屏幕 --- */
@media (max-width: 991px) {
  .features-section h2 {
    margin-bottom: 60px;
  }

  .feature-item {
    /* 在移动端，取消半屏高度，让内容自适应 */
    height: auto;
    flex-direction: column !important;
    gap: 40px;
    padding: 60px 20px;
  }

  .feature-text,
  .feature-image {
    flex-basis: auto;
    width: 100%;
    max-width: 500px;
    padding: 0; /* 重置 padding */
    height: auto; /* 重置高度 */
  }

  /* 在移动端，统一为图片在上，文字在下 */
  .feature-image {
    order: -1;
  }
}
</style>
