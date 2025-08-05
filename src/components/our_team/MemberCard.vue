<template>
  <div class="member-card">
    <div class="avatar-container">
      <img :src="member.avatar" :alt="member.name" class="avatar" @error="handleImageError" />
      <div class="avatar-overlay">
        <span class="view-details-hint">查看详情</span>
      </div>
    </div>
    <h3 class="member-name">{{ member.name }}</h3>
    <p class="member-title">{{ member.title }}</p>
    <p class="member-description">{{ member.description }}</p>
  </div>
</template>

<script setup>
defineProps({
  member: {
    type: Object,
    required: true,
  },
})

const handleImageError = (event) => {
  // 头像加载错误时显示一个1px的空白图片
  event.target.src = 'data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7';
};
</script>

<style scoped>
.member-card {
  --avatar-size: 120px;
  --primary-color: #4a90e2;
  --card-padding: 30px;

  height: 380px;
  width: 100%;
  cursor: pointer;
  transition:
    transform 0.3s ease,
    box-shadow 0.3s ease;
  border-radius: 20px;
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.1);
  padding: var(--card-padding);
  background: white;
  border: 1px solid #f0f0f0;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  box-sizing: border-box;
  overflow: hidden;
}

.member-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 40px rgba(74, 144, 226, 0.2);
}

.member-card:hover .avatar-overlay {
  opacity: 1;
}

.avatar-container {
  position: relative;
  width: var(--avatar-size);
  height: var(--avatar-size);
  margin-bottom: 20px;
}

.avatar {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  object-fit: cover;
  border: 4px solid var(--primary-color);
  transition: all 0.3s ease;
  box-shadow: 0 8px 20px rgba(74, 144, 226, 0.3);
}

.avatar-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  border-radius: 50%;
  background: rgba(74, 144, 226, 0.8);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.view-details-hint {
  color: white;
  font-size: 14px;
  font-weight: 600;
}

.member-name {
  font-size: 24px;
  font-weight: 600;
  color: #2c3e50;
  margin-bottom: 8px;
  background: linear-gradient(135deg, #2c3e50, var(--primary-color));
  -webkit-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
}

.member-title {
  font-size: 16px;
  color: var(--primary-color);
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
  line-clamp:3;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
}

/* 响应式设计已被禁用 */
</style>
