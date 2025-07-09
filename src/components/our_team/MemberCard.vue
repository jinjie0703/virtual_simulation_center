<template>
  <div class="member-card" @mouseenter="flipCard" @mouseleave="unflipCard">
    <div class="card-inner" :class="{ flipped: isFlipped }">
      <!-- 正面 -->
      <div class="card-front">
        <div class="avatar-container">
          <img
            :src="member.avatar"
            :alt="member.name"
            class="avatar"
            @error="handleImageError"
            @load="handleImageLoad"
          />
          <div class="avatar-overlay" :class="{ visible: isFlipped }">
            <div class="flip-hint">
              <span class="hint-icon">👁️</span>
              <span class="hint-text">查看详情</span>
            </div>
          </div>
          <div class="status-indicator" :class="getStatusClass()"></div>
        </div>
        <h3 class="member-name">{{ member.name }}</h3>
        <p class="member-title">{{ member.title }}</p>
        <p class="member-description">{{ member.description }}</p>
        <div class="card-actions">
          <button class="action-btn" @click.stop="sendEmail" title="发送邮件">
            <span>✉️</span>
          </button>
          <button class="action-btn" @click.stop="flipCard" title="查看详情">
            <span>🔄</span>
          </button>
        </div>
      </div>

      <!-- 背面 -->
      <div class="card-back">
        <div class="contact-info">
          <h3 class="contact-title">
            <span class="title-icon">📋</span>
            联系信息
          </h3>
          <div class="contact-details">
            <div class="contact-item">
              <span class="contact-label">
                <span class="label-icon">📧</span>
                邮箱:
              </span>
              <a :href="`mailto:${member.email}`" class="contact-link" @click.stop>
                {{ member.email }}
              </a>
            </div>
            <div class="contact-item">
              <span class="contact-label">
                <span class="label-icon">🏢</span>
                办公室:
              </span>
              <span class="contact-text">{{ member.office }}</span>
            </div>
            <div class="contact-item">
              <span class="contact-label">
                <span class="label-icon">🔬</span>
                研究方向:
              </span>
              <span class="contact-text">{{ member.research }}</span>
            </div>
            <div class="contact-item">
              <span class="contact-label">
                <span class="label-icon">🏆</span>
                学术成果:
              </span>
              <span class="contact-text">{{ member.achievements }}</span>
            </div>
          </div>
          <div class="back-actions">
            <button class="back-btn" @click.stop="unflipCard">
              <span>↩️</span>
              <span>返回</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const props = defineProps({
  member: {
    type: Object,
    required: true,
  },
})

const isFlipped = ref(false)
const imageLoaded = ref(false)

const flipCard = () => {
  isFlipped.value = true
}

const unflipCard = () => {
  isFlipped.value = false
}

const sendEmail = () => {
  window.location.href = `mailto:${props.member.email}`
}

const handleImageError = (event) => {
  // 使用默认头像
  event.target.src =
    'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?w=200&h=200&fit=crop&crop=face'
}

const handleImageLoad = () => {
  imageLoaded.value = true
}

const getStatusClass = () => {
  // 根据成员类型返回不同的状态指示器
  if (props.member.title.includes('教授')) return 'professor'
  if (props.member.title.includes('博士')) return 'doctor'
  if (props.member.title.includes('研究生')) return 'graduate'
  return 'student'
}
</script>

<style scoped>
.member-card {
  perspective: 1000px;
  height: 380px; /* 从320px增加到380px */
  width: 100%;
  cursor: pointer;
  transition: transform 0.3s ease;
}

.member-card:hover {
  transform: translateY(-8px);
}

.card-inner {
  position: relative;
  width: 100%;
  height: 100%;
  transition: transform 0.8s cubic-bezier(0.4, 0, 0.2, 1);
  transform-style: preserve-3d;
}

.card-inner.flipped {
  transform: rotateY(180deg);
}

.card-front,
.card-back {
  position: absolute;
  width: 100%;
  height: 100%;
  backface-visibility: hidden;
  border-radius: 20px;
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.1);
  padding: 30px;
  background: white;
  border: 1px solid #f0f0f0;
  display: flex;
  flex-direction: column;
  justify-content: flex-start; /* 改为顶部对齐 */
  align-items: center;
  text-align: center;
  box-sizing: border-box;
  overflow: hidden; /* 正面保持隐藏溢出 */
}

.card-back {
  transform: rotateY(180deg);
  justify-content: flex-start; /* 背面也改为顶部对齐 */
  padding: 20px; /* 减少内边距为滚动条留出空间 */
  overflow-y: auto; /* 添加垂直滚动条 */
  overflow-x: hidden; /* 隐藏水平滚动条 */
}

/* 正面样式 */
.avatar-container {
  position: relative;
  margin-bottom: 20px;
}

.avatar {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  object-fit: cover;
  border: 4px solid #4a90e2;
  transition: all 0.3s ease;
  box-shadow: 0 8px 20px rgba(74, 144, 226, 0.3);
}

.status-indicator {
  position: absolute;
  bottom: 5px;
  right: 5px;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  border: 3px solid white;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

.status-indicator.professor {
  background: linear-gradient(135deg, #ffd700, #ffed4e);
}

.status-indicator.doctor {
  background: linear-gradient(135deg, #8b5cf6, #a78bfa);
}

.status-indicator.graduate {
  background: linear-gradient(135deg, #10b981, #34d399);
}

.status-indicator.student {
  background: linear-gradient(135deg, #3b82f6, #60a5fa);
}

.avatar-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 120px;
  height: 120px;
  border-radius: 50%;
  background: rgba(74, 144, 226, 0.95);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity 0.3s ease;
  backdrop-filter: blur(5px);
}

.member-card:hover .avatar-overlay {
  opacity: 1;
}

.flip-hint {
  color: white;
  font-size: 12px;
  font-weight: 600;
  text-align: center;
  padding: 0 10px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 4px;
}

.hint-icon {
  font-size: 16px;
}

.member-name {
  font-size: 24px;
  font-weight: 600;
  color: #2c3e50;
  margin-bottom: 8px;
  background: linear-gradient(135deg, #2c3e50, #4a90e2);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.member-title {
  font-size: 16px;
  color: #4a90e2;
  margin-bottom: 16px;
  font-weight: 500;
  padding: 4px 12px;
  background: rgba(74, 144, 226, 0.1);
  border-radius: 12px;
}

.member-description {
  font-size: 14px;
  color: #666;
  line-height: 1.6;
  flex: 1;
  display: flex;
  align-items: center;
  margin-bottom: 16px;
}

.card-actions {
  display: flex;
  gap: 12px;
  margin-top: auto;
}

.action-btn {
  width: 40px;
  height: 40px;
  border: none;
  border-radius: 50%;
  background: linear-gradient(135deg, #4a90e2, #50e3c2);
  color: white;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
  box-shadow: 0 4px 12px rgba(74, 144, 226, 0.3);
}

.action-btn:hover {
  transform: translateY(-2px) scale(1.1);
  box-shadow: 0 6px 16px rgba(74, 144, 226, 0.4);
}

/* 自定义滚动条样式 */
.card-back::-webkit-scrollbar {
  width: 6px;
}

.card-back::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

.card-back::-webkit-scrollbar-thumb {
  background: linear-gradient(135deg, #4a90e2, #50e3c2);
  border-radius: 3px;
}

.card-back::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(135deg, #357abd, #3dd1a7);
}

/* 背面样式 */
.contact-info {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 15px;
  min-height: 100%; /* 确保内容可以超出容器高度 */
}

.contact-title {
  font-size: 20px;
  font-weight: 700;
  color: #333;
  margin-bottom: 15px;
  background: linear-gradient(135deg, #4a90e2, #50e3c2);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  text-align: center;
  border-bottom: 2px solid #e1e5e9;
  padding-bottom: 10px;
  flex-shrink: 0; /* 防止标题被压缩 */
}

.title-icon {
  font-size: 18px;
}

.contact-details {
  display: flex;
  flex-direction: column;
  gap: 12px;
  flex-grow: 1; /* 让详情区域占用剩余空间 */
}

.contact-item {
  display: flex;
  align-items: flex-start;
  gap: 10px;
  padding: 8px 0;
  border-bottom: 1px solid #f5f5f5;
}

.contact-label {
  font-size: 13px;
  color: #666;
  font-weight: 600;
  display: flex;
  align-items: center;
  gap: 6px;
}

.label-icon {
  font-size: 14px;
}

.contact-link {
  color: #4a90e2;
  text-decoration: none;
  font-size: 13px;
  transition: all 0.3s ease;
  word-break: break-all;
}

.contact-link:hover {
  color: #2980b9;
  text-decoration: underline;
  transform: translateX(4px);
}

.contact-text {
  color: #333;
  font-size: 13px;
  line-height: 1.4;
}

.back-actions {
  margin-top: 15px;
  display: flex;
  gap: 10px;
  justify-content: center;
  flex-shrink: 0; /* 防止按钮被压缩 */
}

.back-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 20px;
  background: linear-gradient(135deg, #6b7280, #9ca3af);
  color: white;
  border: none;
  border-radius: 20px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
  transition: all 0.3s ease;
  box-shadow: 0 4px 12px rgba(107, 114, 128, 0.3);
}

.back-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(107, 114, 128, 0.4);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .member-card {
    height: 420px; /* 移动端适当增加高度 */
  }

  .member-name {
    font-size: 20px;
  }

  .member-title {
    font-size: 14px;
  }

  .member-description {
    font-size: 13px;
  }

  .avatar {
    width: 100px;
    height: 100px;
  }

  .avatar-overlay {
    width: 100px;
    height: 100px;
  }
}

@media (max-width: 480px) {
  .member-card {
    height: 400px;
  }

  .card-front,
  .card-back {
    padding: 20px;
  }

  .member-name {
    font-size: 18px;
  }

  .member-title {
    font-size: 13px;
  }

  .member-description {
    font-size: 12px;
  }

  .avatar {
    width: 90px;
    height: 90px;
  }

  .avatar-overlay {
    width: 90px;
    height: 90px;
  }
}
@media (max-width: 768px) {
  .card-back {
    padding: 15px; /* 移动端进一步减少内边距 */
  }

  .contact-title {
    font-size: 18px;
  }

  .contact-item {
    font-size: 14px;
  }
}

@media (max-width: 480px) {
  .card-back {
    padding: 12px;
  }

  .contact-title {
    font-size: 16px;
  }

  .contact-item {
    font-size: 13px;
  }
}
</style>
