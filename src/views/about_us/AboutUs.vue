<template>
  <div class="about-page">
    <header class="about-header">
      <div>
        <h1 class="title">关于我们</h1>
        <p class="subtitle">每一个像素，都承载着对知识的无限热忱</p>
      </div>
    </header>
    <div class="about-container">
      <section class="about-section story" ref="storySection">
        <!-- 图标装饰 -->
        <div class="section-header">
          <div class="section-icon">
            <img src="@/assets/about_us/our_story.svg" alt="story icon" />
          </div>
          <h2>我们的故事</h2>
        </div>
        <p>
          我们是虚拟仿真中心的先驱者，致力于通过尖端技术打破传统教育和培训的界限。自成立以来，我们始终专注于创造沉浸式、交互式的学习体验，帮助用户在安全、可控的虚拟环境中掌握复杂技能，激发无限潜能。
        </p>
        <!-- 进度条装饰 -->
        <div class="progress-bar"></div>
      </section>

      <section class="about-section mission" ref="missionSection">
        <!-- 图标装饰 -->
        <div class="section-header">
          <div class="section-icon">
            <img src="@/assets/about_us/our_mission.svg" alt="mission icon" />
          </div>
          <h2>我们的使命</h2>
        </div>
        <p>
          我们的使命是"让学习更直观，让实践无风险"。我们相信，通过高质量的虚拟仿真内容，可以革新知识传播的方式，为各行各业提供高效、安全且富有吸引力的培训解决方案，推动社会进步与创新。
        </p>
        <!-- 进度条装饰 -->
        <div class="progress-bar"></div>
      </section>

      <section class="about-section vision" ref="visionSection">
        <!-- 图标装饰 -->
        <div class="section-header">
          <div class="section-icon">
            <img src="@/assets/about_us/our_vision.svg" alt="vision icon" />
          </div>
          <h2>我们的愿景</h2>
        </div>
        <p>
          我们期望成为虚拟仿真教育领域的领导者，通过不断创新和技术突破，构建一个无边界的知识传播平台。我们致力于让每个人都能享受到高质量的虚拟仿真学习体验，无论其地理位置或资源条件如何。
        </p>
        <!-- 进度条装饰 -->
        <div class="progress-bar"></div>
      </section>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'

// 创建引用以直接操作DOM元素
const storySection = ref(null)
const missionSection = ref(null)
const visionSection = ref(null)

onMounted(() => {
  // 使用Intersection Observer API监测元素是否进入视口
  const observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          entry.target.classList.add('animate-in')
        }
      })
    },
    { threshold: 0.2, rootMargin: '0px 0px -100px 0px' },
  )

  // 观察所有section引用
  const sections = [storySection.value, missionSection.value, visionSection.value]
  sections.forEach((section) => {
    if (section) observer.observe(section)
  })

  // 组件卸载时清理observer
  return () => {
    sections.forEach((section) => {
      if (section) observer.unobserve(section)
    })
  }
})
</script>

<style scoped>
.about-page {
  display: flex;
  flex-direction: column;
  width: 100%;
}

/* 标题样式 */
.about-header {
  padding: 120px 0px 80px 0px;
  background: linear-gradient(135deg, #4a90e2, #50e3c2);
  text-align: center;
  color: white;
  position: relative;
  overflow: hidden;
}

.title {
  font-size: 3.5rem;
  font-weight: 800;
  margin: 0 0 15px;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
  animation: fadeInDown 1s ease-out;
}

.subtitle {
  font-size: 1.25rem;
  max-width: 800px;
  margin: 0 auto;
  opacity: 0.9;
  line-height: 1.6;
  animation: fadeInUp 1s ease-out 0.3s both;
}

@keyframes fadeInDown {
  from {
    opacity: 0;
    transform: translateY(-30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* 关于容器样式 */
.about-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 60px 30px;
  position: relative;
}

.about-section {
  margin-bottom: 60px;
  padding: 40px;
  background: linear-gradient(135deg, #fdfdfd 0%, #f8f9fa 100%);
  border-radius: 16px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(74, 144, 226, 0.1);
  position: relative;
  overflow: hidden;
  transition:
    transform 0.4s cubic-bezier(0.4, 0, 0.2, 1),
    opacity 0.4s cubic-bezier(0.4, 0, 0.2, 1),
    box-shadow 0.4s cubic-bezier(0.4, 0, 0.2, 1),
    border-color 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  opacity: 0;
  transform: translateY(40px) scale(0.9);
}

/* 滚动触发的动画,进入视口时 */
.about-section.animate-in {
  opacity: 1;
  transform: translateY(0) scale(1);
}

/* 增强的悬停效果 */
.about-section:hover {
  transform: translateY(-12px) scale(1.05);
  box-shadow: 0 20px 40px rgba(74, 144, 226, 0.15);
  border-color: rgba(74, 144, 226, 0.3);
  z-index: 10;
}

/* 悬停时的光晕效果 */
.about-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(74, 144, 226, 0.1), rgba(80, 227, 194, 0.05));
  opacity: 0;
  transition: opacity 0.3s ease;
  pointer-events: none;
}

.about-section:hover::before {
  opacity: 1;
}

.section-header {
  display: flex;
  align-items: center;
  margin-bottom: 24px;
}

/* 图标样式 */
.section-icon {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 50px;
  height: 50px;
  border-radius: 50%;
  margin-right: 20px;
  transition: transform 0.3s ease;
}

.about-section:hover .section-icon {
  transform: scale(1.1) rotate(5deg);
}

/* 由卡片容器的margin(60px)和padding(60px)控制,重新设定让页面美观 */
.about-section:last-child {
  margin-bottom: 20px;
}

.about-section h2 {
  font-size: 2.5rem;
  border-bottom: 3px solid transparent;
  background-size: 100% 3px;
  background-repeat: no-repeat;
  background-position: bottom;
  display: inline-block;
  font-weight: 600;
  transition: all 0.3s ease;

  /* 渐变文字效果 */
  background: linear-gradient(135deg, #2c3e50, #4a90e2); /* 设置背景为渐变色 */
  -webkit-background-clip: text; /* 将背景裁剪为文字的形状（-webkit-前缀兼容旧版浏览器） */
  background-clip: text; /* 标准写法 */
  -webkit-text-fill-color: transparent; /* 将文字本身的颜色设为透明，从而透出背景的渐变色 */

  position: relative;
  padding-bottom: 8px; /* 给下划线留出一点空间，避免贴得太近 */
}

/* 使用 ::after 伪元素来创建下划线 */
.about-section h2::after {
  content: ''; /* 伪元素必须有 content */
  position: absolute; /* 绝对定位，相对于 h2 */
  left: 0;
  bottom: 0; /* 定位在 h2 的最底部 */
  width: 0%; /* 初始宽度为0，准备做动画 */
  height: 3px; /* 下划线的高度 */
  background: linear-gradient(90deg, #4a90e2, #50e3c2);
  transition: width 0.4s ease-out; /* 宽度变化时有过渡效果 */
}

/* 悬停时让下划线生长出来 */
.about-section:hover h2::after {
  width: 100%;
}

.about-section p {
  font-size: 1.2rem;
  color: #333;
  margin-bottom: 20px;
  text-align: justify;
  line-height: 1.8;
}

/* 进度条装饰 */
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

/* 打字机效果（可选） */
/* @keyframes typewriter {
  from {
    width: 0;
  }
  to {
    width: 100%;
  }
} */

/* 响应式设计优化 */
@media (max-width: 768px) {
  .about-header {
    padding: 80px 0 60px;
  }

  .title {
    font-size: 2.5rem;
  }

  .about-container {
    padding: 40px 15px;
  }

  .about-section {
    padding: 30px 20px;
    margin-bottom: 40px;
  }

  .about-section h2 {
    font-size: 2rem;
  }

  .section-icon {
    width: 35px;
    height: 35px;
  }
}

@media (max-width: 480px) {
  .title {
    font-size: 2rem;
  }

  .subtitle {
    font-size: 1rem;
  }

  .about-section {
    padding: 25px 15px;
  }

  .about-section h2 {
    font-size: 1.75rem;
  }

  .about-section p {
    font-size: 1rem;
  }
}
</style>
