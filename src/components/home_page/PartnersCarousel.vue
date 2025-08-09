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
            <img
              :src="getImageUrl(company.logo)"
              :alt="company.name"
              class="partner-logo"
              @load="onImageLoad('companies')"
            />
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
            <img
              :src="getImageUrl(school.logo)"
              :alt="school.name"
              class="partner-logo"
              @load="onImageLoad('schools')"
            />
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

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL

const getImageUrl = (logoUrl) => {
  if (!logoUrl) return ''
  return `${API_BASE_URL}${logoUrl}`
}

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

const loadedImages = ref({
  companies: 0,
  schools: 0,
})

const onImageLoad = (type) => {
  loadedImages.value[type]++
  if (
    type === 'companies' &&
    loadedImages.value.companies === props.companies.length &&
    companiesTrack.value
  ) {
    animationState.value.companies.resetPoint = companiesTrack.value.scrollWidth / 2
    if (!isPaused.value.companies) scrollCarousel('companies')
  }
  if (
    type === 'schools' &&
    loadedImages.value.schools === props.schools.length &&
    schoolsTrack.value
  ) {
    animationState.value.schools.resetPoint = schoolsTrack.value.scrollWidth / 2
    if (!isPaused.value.schools) scrollCarousel('schools')
  }
}

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
  // 现在初始化逻辑由 onImageLoad 触发，但保留 resize 监听器
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
  font-size: 3rem;
  margin-bottom: 40px;
  color: black;
  position: relative;
  display: inline-block;
  left: 50%;
  transform: translateX(-50%);
}
.section-title::after {
  content: '';
  display: block;
  width: 60%;
  height: 3px;
  background-color: rgb(176, 170, 170);
  margin: 15px auto 0;
}

.partners-section {
  margin-bottom: 40px;
}

.partner-type {
  text-align: center;
  font-size: 2rem;
  margin-bottom: 20px;
  color: rgb(51, 49, 49);
  font-weight: 500;
}

.carousel-container {
  overflow: hidden;
  position: relative;
  width: 100%;
  padding-top: 20px;
  padding-bottom: 20px;

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
  flex-direction: row;
  align-items: center;
  margin: 0 15px; /* 增加了间距 */
  width: 250px; /* 增加了宽度 */
  height: 100px; /* 设置固定高度 */
  padding: 15px;
  transition: transform 0.3s;
  background-color: white;
  border-radius: 16px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.05);
  padding: 15px;
  box-sizing: border-box;
}

.partner-item:hover {
  transform: scale(1.05);
  box-shadow: 0 6px 15px rgba(0, 0, 0, 0.1);
}

.partner-logo {
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
}
</style>
