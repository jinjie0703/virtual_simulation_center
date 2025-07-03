<template>
  <section class="features-section">
    <div class="container">
      <h2>我们的成果</h2>
      <div
        v-for="(feature, index) in features"
        :key="index"
        class="feature-item"
        :class="{ 'layout-reverse': index % 2 !== 0 }"
        ref="featureElements"
      >
        <div class="feature-text">
          <h3>{{ feature.title }}</h3>
          <p>{{ feature.description }}</p>
        </div>
        <div class="feature-image">
          <img :src="feature.imageUrl" :alt="feature.title" />
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue'

//const props = defineProps({
//   features: {
//     type: Array,
//     required: true,
//   },
// })

const featureElements = ref([])
let observer = null

onMounted(() => {
  const options = {
    root: null,
    threshold: 0.2,
  }

  observer = new IntersectionObserver((entries) => {
    entries.forEach((entry) => {
      if (entry.isIntersecting) {
        entry.target.classList.add('is-visible')
        observer.unobserve(entry.target) // 动画只播放一次
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
  padding: 80px 0;
  background-color: #ffffff;
}
.features-section h2 {
  text-align: center;
  font-size: 2.5rem;
  margin-bottom: 60px;
}
.feature-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 80px;
  gap: 5%;
  opacity: 0;
  transform: translateY(50px);
  transition:
    opacity 0.6s ease-out,
    transform 0.6s ease-out;
}
.feature-item.is-visible {
  opacity: 1;
  transform: translateY(0);
}
.feature-item.layout-reverse {
  flex-direction: row-reverse;
}
.feature-text,
.feature-image {
  width: 47.5%;
}
.feature-image img {
  width: 100%;
  height: auto;
  display: block;
  border-radius: 8px;
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
}
.feature-text h3 {
  font-size: 2rem;
  margin-top: 0;
}
.feature-text p {
  font-size: 1.1rem;
  line-height: 1.6;
  color: #555;
}

@media (max-width: 768px) {
  .feature-item,
  .feature-item.layout-reverse {
    flex-direction: column;
    text-align: center;
  }
  .feature-text,
  .feature-image {
    width: 100%;
  }
  .feature-image {
    margin-top: 20px;
  }
}
</style>
