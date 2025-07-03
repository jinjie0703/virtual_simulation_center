<template>
  <div class="page-container">
    <div class="card-wrapper">
      <!-- 欢迎语 -->
      <div class="welcome-text">
        <h1>{{ isLogin ? '欢迎回来' : '创建新账户' }}</h1>
        <p>{{ isLogin ? '登录以继续探索' : '加入我们，开启未来之旅' }}</p>
      </div>
      <!-- 表单容器 -->
      <div class="form-container">
        <transition name="form-fade" mode="out-in">
          <component :is="isLogin ? LoginForm : RegisterForm" @submit-success="handleSubmit" />
        </transition>
      </div>
      <!-- 切换链接 -->
      <div class="switch-link">
        <p>
          {{ isLogin ? '还没有账户？' : '已经有账户了？' }}
          <a href="#" @click.prevent="isLogin = !isLogin">
            {{ isLogin ? '立即注册' : '前去登录' }}
          </a>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="js">
import { ref } from 'vue'
import LoginForm from './LoginForm.vue'
import RegisterForm from './RegisterForm.vue'

const isLogin = ref(true)

const handleSubmit = (formData) => {
  if (isLogin.value) {
    console.log('父组件收到登录成功事件，数据：', formData)
  } else {
    console.log('父组件收到注册成功事件，数据：', formData)
    isLogin.value = true
  }
}
</script>

<style scoped>
/* 1. 全屏背景容器 */
.page-container {
  min-height: calc(100vh - 64px);
  display: flex;
  align-items: center;
  justify-content: center;
  box-sizing: border-box; /* 让 padding 不会撑大容器 */
  padding: 1rem;
  padding-top: 15vh; /* 实现内容上移 */
  background: linear-gradient(135deg, #6366f1, #ec4899);
  animation: gradient-animation 15s ease infinite;
  background-size: 400% 400%;
}

@keyframes gradient-animation {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}

/* 2. 主卡片: 保持较窄宽度 */
.card-wrapper {
  width: 100%;
  max-width: 350px;
  background-color: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(12px);
  border-radius: 1.5rem;
  box-shadow: 0 20px 50px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  padding: 40px;
  text-align: center;
}

/* 3. 输入框组通用样式 */
:global(.form-common .input-group) {
  position: relative;
  margin-bottom: 1.5rem;
}

/* 4. 输入框样式: 浅灰背景，聚焦时蓝色边框 */
:global(.form-common .form-input) {
  width: 94%;
  padding: 12px 10px;
  border: 1px solid #d1d5db;
  border-radius: 0.5rem;
  background-color: #f3f4f6; /* 使用新的浅灰色背景 */
  font-size: 1rem;
  color: #000000;
  outline: none;
  transition:
    border-color 0.3s,
    box-shadow 0.3s;
}

:global(.form-common .form-input:focus) {
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.3);
}

/* 5. 提交按钮: 黑底白字，阴影常驻 */
:global(.form-common .form-button) {
  width: 100%;
  padding: 12px;
  border: none;
  border-radius: 0.5rem;
  background-color: #000000;
  color: white;
  font-size: 1rem;
  font-weight: 400;
  cursor: pointer;
  transform: translateY(0px);
  box-shadow: 0 6px 15px rgba(94, 87, 87, 0.2);
  transition: all 0.2s ease;
}

/* 悬停时按钮颜色变浅，且上浮更多一点 */
:global(.form-common .form-button:hover) {
  background-color: #374151; /* 悬停时颜色变浅 */
  transform: translateY(-2px); /* 悬停时上浮更多一点 */
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.25);
}

/* 6. 按钮禁用时的样式 */
:global(.form-common .form-button:disabled) {
  background-color: #9ca3af; /* 灰色背景 */
  cursor: not-allowed;
  transform: translateY(0); /* 移除上浮效果 */
  box-shadow: none; /* 移除阴影 */
}

.welcome-text h1 {
  font-size: 2rem;
  font-weight: 700;
  color: var(--text-dark);
  margin-bottom: 0.5rem;
}
.welcome-text p {
  color: var(--text-light);
  margin-bottom: 2rem;
}
.form-container {
  position: relative;
  min-height: 200px;
}
.form-fade-enter-active,
.form-fade-leave-active {
  transition: opacity 0.3s ease;
}
.form-fade-enter-from,
.form-fade-leave-to {
  opacity: 0;
}
.switch-link {
  margin-top: 1.5rem;
  font-size: 0.9rem;
  color: var(--text-light);
}
.switch-link a {
  color: var(--primary-color);
  font-weight: 600;
  text-decoration: none;
}
.switch-link a:hover {
  text-decoration: underline;
}
</style>
