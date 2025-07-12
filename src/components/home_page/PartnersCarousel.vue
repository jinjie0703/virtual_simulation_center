<template>
  <div class="partners-carousel">
    <!-- 我将标题改为了和您图片一致的“部分合作院校”，您可以按需修改 -->
    <h2 class="section-title">部分合作伙伴</h2>

    <!-- 如果您只想展示院校，可以把这个 "合作企业" section 注释或删除 -->
    <div class="partners-section" v-if="companies.length > 0">
      <h3 class="partner-type">合作企业</h3>
      <div
        class="carousel-container"
        @mouseenter="pauseScroll('companies')"
        @mouseleave="resumeScroll('companies')"
      >
        <div
          ref="companiesTrack"
          class="carousel-track"
          :style="{ transform: `translateX(${companiesOffset}px)` }"
        >
          <a
            v-for="company in companiesDoubled"
            :key="company.uniqueId"
            :href="company.url"
            target="_blank"
            rel="noopener noreferrer"
            class="partner-item"
          >
            <img :src="company.logo" :alt="company.name" class="partner-logo" />
            <span class="partner-name">{{ company.name }}</span>
          </a>
        </div>
      </div>
    </div>

    <div class="partners-section" v-if="schools.length > 0">
      <h3 class="partner-type">合作学校</h3>
      <div
        class="carousel-container"
        @mouseenter="pauseScroll('schools')"
        @mouseleave="resumeScroll('schools')"
      >
        <div
          ref="schoolsTrack"
          class="carousel-track"
          :style="{ transform: `translateX(${schoolsOffset}px)` }"
        >
          <a
            v-for="school in schoolsDoubled"
            :key="school.uniqueId"
            :href="school.url"
            target="_blank"
            rel="noopener noreferrer"
            class="partner-item"
          >
            <img :src="school.logo" :alt="school.name" class="partner-logo" />
            <span class="partner-name">{{ school.name }}</span>
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'

const props = defineProps({
  companies: {
    type: Array,
    default: () => [],
  },
  schools: {
    type: Array,
    default: () => [],
  },
  scrollSpeed: {
    type: Number,
    default: 0.5, // 调慢了速度，效果更好
  },
})

// 为 v-for 提供更稳定的 key
const companiesDoubled = computed(() =>
  [...props.companies, ...props.companies].map((item, index) => ({
    ...item,
    uniqueId: `company-${index}`,
  })),
)
const schoolsDoubled = computed(() =>
  [...props.schools, ...props.schools].map((item, index) => ({
    ...item,
    uniqueId: `school-${index}`,
  })),
)

const companiesOffset = ref(0)
const schoolsOffset = ref(0)

const companiesTrack = ref(null)
const schoolsTrack = ref(null)

const animationState = ref({
  companies: { id: null, resetPoint: 0 },
  schools: { id: null, resetPoint: 0 },
})

const isPaused = ref({
  companies: false,
  schools: false,
})

const scrollCarousel = (type) => {
  if (isPaused.value[type] || !animationState.value[type].resetPoint) return

  const state = animationState.value[type]
  if (type === 'companies') {
    companiesOffset.value -= props.scrollSpeed
    if (Math.abs(companiesOffset.value) >= state.resetPoint) {
      companiesOffset.value = 0
    }
    state.id = requestAnimationFrame(() => scrollCarousel('companies'))
  } else {
    schoolsOffset.value -= props.scrollSpeed
    if (Math.abs(schoolsOffset.value) >= state.resetPoint) {
      schoolsOffset.value = 0
    }
    state.id = requestAnimationFrame(() => scrollCarousel('schools'))
  }
}

const pauseScroll = (type) => {
  isPaused.value[type] = true
  if (animationState.value[type].id) {
    cancelAnimationFrame(animationState.value[type].id)
    animationState.value[type].id = null
  }
}

const resumeScroll = (type) => {
  isPaused.value[type] = false
  if (!animationState.value[type].id) {
    scrollCarousel(type)
  }
}

const initCarousel = () => {
  // 等待DOM更新后计算宽度
  if (props.companies.length > 0 && companiesTrack.value) {
    animationState.value.companies.resetPoint = companiesTrack.value.scrollWidth / 2
    if (!isPaused.value.companies) scrollCarousel('companies')
  }
  if (props.schools.length > 0 && schoolsTrack.value) {
    animationState.value.schools.resetPoint = schoolsTrack.value.scrollWidth / 2
    if (!isPaused.value.schools) scrollCarousel('schools')
  }
}

onMounted(() => {
  // 使用 nextTick 确保DOM渲染完成
  // 为了应对图片加载可能导致的宽度变化，可以加一个延时
  setTimeout(initCarousel, 200)

  // 如果窗口大小变化，重新计算宽度
  window.addEventListener('resize', initCarousel)
})

onUnmounted(() => {
  if (animationState.value.companies.id) {
    cancelAnimationFrame(animationState.value.companies.id)
  }
  if (animationState.value.schools.id) {
    cancelAnimationFrame(animationState.value.schools.id)
  }
  window.removeEventListener('resize', initCarousel)
})
</script>

<style scoped>
.partners-carousel {
  padding: 80px 0; /* 增加上下内边距，使其呼吸感更强 */
  background-color: #f9fafb; /* 使用柔和的淡灰色背景，与白色内容区和页脚区分 */
  overflow: hidden;
  font-family:
    -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
}

.section-title {
  text-align: center;
  font-size: 28px;
  margin-bottom: 40px;
  color: #333;
  position: relative;
  display: inline-block;
  left: 50%;
  transform: translateX(-50%);
}
.section-title::after {
  content: '';
  display: block;
  width: 60%;
  height: 2px;
  background-color: #ccc;
  margin: 10px auto 0;
}

.partners-section {
  margin-bottom: 30px;
}

.partner-type {
  text-align: center;
  font-size: 22px;
  margin-bottom: 20px;
  color: #555;
  font-weight: normal;
}

.carousel-container {
  overflow: hidden;
  position: relative;
  width: 100%;

  /* ✨ 新增：添加蒙版实现两端淡出效果 */
  -webkit-mask-image: linear-gradient(to right, transparent, black 15%, black 85%, transparent);
  mask-image: linear-gradient(to right, transparent, black 15%, black 85%, transparent);
}

.carousel-track {
  display: flex;
  width: max-content;
  white-space: nowrap;
}

.partner-item {
  display: inline-flex;
  flex-direction: column;
  align-items: center;
  justify-content: center; /* 垂直居中 */
  margin: 0 25px; /* 增加了间距 */
  width: 150px; /* 增加了宽度 */
  height: 150px; /* 设置固定高度 */
  text-decoration: none;
  color: #333;
  transition: transform 0.3s;
  background-color: white;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.05);
  padding: 15px;
  box-sizing: border-box;
}

.partner-item:hover {
  transform: scale(1.05);
  box-shadow: 0 6px 25px rgba(0, 0, 0, 0.1);
}

.partner-logo {
  width: 80px;
  height: 80px;
  object-fit: contain;
  margin-bottom: 10px;
}

.partner-name {
  font-size: 14px;
  text-align: center;
  max-width: 100%;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: normal;
  line-height: 1.3;
  height: 36px;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  color: #666; /* 名字颜色调暗一点 */
}

@media (max-width: 768px) {
  .partner-item {
    width: 120px;
    height: 120px;
    margin: 0 15px;
  }
  .partner-logo {
    width: 60px;
    height: 60px;
  }
  .partner-name {
    font-size: 12px;
    height: 30px;
  }
}
</style>
