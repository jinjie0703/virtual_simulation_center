<template>
  <div class="member-card">
    <div class="avatar-container">
      <img :src="member.avatar" :alt="member.name" class="avatar" @error="handleImageError" />
      <div class="avatar-overlay">
        <div class="view-details-hint">
          <span class="hint-text">查看详情</span>
        </div>
      </div>
    </div>
    <h3 class="member-name">{{ member.name }}</h3>
    <p class="member-title">{{ member.title }}</p>
    <p class="member-description">{{ member.description }}</p>
  </div>
</template>

<script setup>
// const props = defineProps({
//   member: {
//     type: Object,
//     required: true,
//   },
// })

const handleImageError = (event) => {
  // 使用默认头像
  event.target.src =
    'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?w=200&h=200&fit=crop&crop=face'
}
</script>

<style scoped>
.member-card {
  height: 380px;
  width: 100%;
  cursor: pointer;
  transition:
    transform 0.3s ease,
    box-shadow 0.3s ease;
  border-radius: 20px;
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.1);
  padding: 30px;
  background: white;
  border: 1px solid #f0f0f0;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: center;
  text-align: center;
  box-sizing: border-box;
  overflow: hidden;
}

.member-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 40px rgba(74, 144, 226, 0.2);
}

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
  background: rgba(74, 144, 226, 0.8);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.member-card:hover .avatar-overlay {
  opacity: 1;
}

.view-details-hint {
  color: white;
  font-size: 14px;
  font-weight: 600;
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
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
  flex-grow: 1;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .member-card {
    height: auto;
    min-height: 360px;
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
</style>
