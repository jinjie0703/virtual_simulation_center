<!-- LoginForm.vue -->
<template>
  <form class="form-common" @submit.prevent="handleLogin">
    <div class="input-group">
      <input
        type="text"
        id="login-username"
        class="form-input"
        placeholder="用户名 / 邮箱"
        required
        v-model="loginForm.username"
      />
    </div>
    <div class="input-group">
      <input
        type="password"
        id="login-password"
        class="form-input"
        placeholder="密码"
        required
        v-model="loginForm.password"
      />
    </div>
    <p v-if="errorMsg" class="error-text">{{ errorMsg }}</p>
    <button type="submit" class="form-button" :disabled="isLoading">
      {{ isLoading ? '登录中...' : '登 录' }}
    </button>
  </form>
</template>

<!-- LoginForm.vue (修改后) -->
<script setup lang="js">
import { reactive, ref } from 'vue'
import axios from 'axios'

const API_BASE_URL = 'http://localhost:8080/api'

const emit = defineEmits(['submit-success'])
const loginForm = reactive({ username: '', password: '' })
const isLoading = ref(false) // 加载状态
const errorMsg = ref('') // 错误信息

const handleLogin = async () => {
  errorMsg.value = ''
  isLoading.value = true

  try {
    // 调用后端登录 API
    const response = await axios.post(`${API_BASE_URL}/login`, {
      username: loginForm.username,
      password: loginForm.password,
    })

    // 登录成功后，将 JWT Token 存入 localStorage
    localStorage.setItem('authToken', response.data.data.token)

    // 通知父组件登录成功
    emit('submit-success', response.data.data.user)
  } catch (error) {
    errorMsg.value = error.response?.data?.msg || '用户名或密码错误。'
  } finally {
    isLoading.value = false
  }
}
</script>
