<template>
  <header class="achievement-header">
    <div class="header-content">
      <h1 class="title">关于我们</h1>
      <p class="subtitle">每一个像素，都承载着对知识的无限热忱</p>
    </div>
  </header>
  <div class="about-us-container">
    <!-- 新增：装饰性背景元素 -->
    <div class="background-decoration">
      <div class="floating-circle circle-1"></div>
      <div class="floating-circle circle-2"></div>
      <div class="floating-circle circle-3"></div>
    </div>

    <section class="about-section story">
      <!-- 新增：图标装饰 -->
      <div class="section-icon">
        <img src="@/assets/about_us/our_story.svg" alt="story icon" />
      </div>
      <h2>我们的故事</h2>
      <p>
        我们是虚拟仿真中心的先驱者，致力于通过尖端技术打破传统教育和培训的界限。自成立以来，我们始终专注于创造沉浸式、交互式的学习体验，帮助用户在安全、可控的虚拟环境中掌握复杂技能，激发无限潜能。
      </p>
      <!-- 新增：进度条装饰 -->
      <div class="progress-bar"></div>
    </section>

    <section class="about-section mission">
      <!-- 新增：图标装饰 -->
      <div class="section-icon">
        <img src="@/assets/about_us/our_mission.svg" alt="mission icon" />
      </div>
      <h2>我们的使命</h2>
      <p>
        我们的使命是"让学习更直观，让实践无风险"。我们相信，通过高质量的虚拟仿真内容，可以革新知识传播的方式，为各行各业提供高效、安全且富有吸引力的培训解决方案，推动社会进步与创新。
      </p>
      <!-- 新增：进度条装饰 -->
      <div class="progress-bar"></div>
    </section>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'

onMounted(() => {
  // 新增：观察器，用于触发滚动时的动画
  const observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          entry.target.classList.add('animate-in')
        }
      })
    },
    { threshold: 0.1 },
  )

  // 观察所有section
  document.querySelectorAll('.about-section').forEach((section) => {
    observer.observe(section)
  })
})
</script>

<style scoped>
.achievement-header {
  padding: 120px 0px 80px 0px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  text-align: center;
  color: white;
}
.title {
  font-size: 3.5rem;
  font-weight: 800;
  margin: 0 0 15px;
  letter-spacing: -1px;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}
.subtitle {
  font-size: 1.25rem;
  max-width: 800px;
  margin: 0 auto;
  opacity: 0.9;
  line-height: 1.6;
}

.about-us-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 40px 20px;
  font-family: 'Helvetica Neue', Arial, sans-serif;
  color: #333;
  line-height: 1.8;
  position: relative;
  overflow: hidden;
}

/* 新增：背景装饰元素 */
.background-decoration {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  pointer-events: none;
  z-index: -1;
}

.floating-circle {
  position: absolute;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(74, 144, 226, 0.1), rgba(80, 227, 194, 0.1));
  animation: float 6s ease-in-out infinite;
}

.circle-1 {
  width: 120px;
  height: 120px;
  top: 10%;
  right: 10%;
  animation-delay: 0s;
}

.circle-2 {
  width: 80px;
  height: 80px;
  top: 60%;
  left: 5%;
  animation-delay: 2s;
}

.circle-3 {
  width: 60px;
  height: 60px;
  bottom: 20%;
  right: 15%;
  animation-delay: 4s;
}

@keyframes float {
  0%,
  100% {
    transform: translateY(0px) rotate(0deg);
  }
  33% {
    transform: translateY(-20px) rotate(120deg);
  }
  66% {
    transform: translateY(10px) rotate(240deg);
  }
}

.about-section {
  margin-bottom: 60px;
  padding: 40px;
  background: linear-gradient(135deg, #fdfdfd 0%, #f8f9fa 100%);
  border-radius: 16px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.06);
  border: 1px solid rgba(74, 144, 226, 0.1);
  position: relative;
  overflow: hidden;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  opacity: 0;
  transform: translateY(40px) scale(0.95);
}

/* 新增：滚动触发的动画 */
.about-section.animate-in {
  opacity: 1;
  transform: translateY(0) scale(1);
}

/* 新增：为不同section设置动画延迟 */
.about-section.mission.animate-in {
  transition-delay: 0.2s;
}

/* 增强的悬停效果 */
.about-section:hover {
  transform: translateY(-12px) scale(1.02);
  box-shadow: 0 20px 40px rgba(74, 144, 226, 0.15);
  border-color: rgba(74, 144, 226, 0.3);
}

.about-section:hover::before {
  opacity: 1;
}

/* 新增：悬停时的光晕效果 */
.about-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(74, 144, 226, 0.05), rgba(80, 227, 194, 0.05));
  opacity: 0;
  transition: opacity 0.3s ease;
  pointer-events: none;
}

/* 新增：图标样式 */
.section-icon {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 80px;
  height: 80px;
  background: linear-gradient(135deg, rgba(74, 144, 226, 0.1), rgba(80, 227, 194, 0.1));
  border-radius: 50%;
  margin-bottom: 20px;
  transition: transform 0.3s ease;
}

.about-section:hover .section-icon {
  transform: scale(1.1) rotate(5deg);
}

.about-section:last-child {
  margin-bottom: 20px;
}

.about-section h2 {
  font-size: 2.4rem;
  color: #2c3e50;
  margin-bottom: 24px;
  padding-bottom: 12px;
  border-bottom: 3px solid transparent;
  background: linear-gradient(90deg, #4a90e2, #50e3c2);
  background-size: 100% 3px;
  background-repeat: no-repeat;
  background-position: bottom;
  display: inline-block;
  font-weight: 600;
  transition: all 0.3s ease;
}

.about-section:hover h2 {
  background-size: 100% 3px;
  color: #1a252f;
}

.about-section p {
  font-size: 1.15rem;
  color: #555;
  margin-bottom: 20px;
  text-align: justify;
}

/* 新增：进度条装饰 */
.progress-bar {
  height: 4px;
  background: linear-gradient(90deg, #4a90e2, #50e3c2);
  border-radius: 2px;
  width: 0%;
  transition: width 1s ease-out 0.5s;
}

.about-section.animate-in .progress-bar {
  width: 100%;
}

/* 新增：响应式设计优化 */
@media (max-width: 768px) {
  .about-us-container {
    padding: 20px 15px;
  }

  .about-section {
    padding: 30px 20px;
    margin-bottom: 40px;
  }

  .about-section h2 {
    font-size: 2rem;
  }

  .section-icon {
    width: 60px;
    height: 60px;
  }

  .floating-circle {
    display: none;
  }
}

/* 新增：打字机效果（可选） */
@keyframes typewriter {
  from {
    width: 0;
  }
  to {
    width: 100%;
  }
}

/* 新增：渐变文字效果 */
.about-section h2 {
  background: linear-gradient(135deg, #2c3e50, #4a90e2);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}
</style>
