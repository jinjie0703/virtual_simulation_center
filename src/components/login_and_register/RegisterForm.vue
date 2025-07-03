<!-- RegisterForm.vue (最终版) -->
<template>
  <form class="form-common" @submit.prevent="handleRegister">
    <div class="input-group">
      <input
        type="text"
        id="reg-username"
        class="form-input"
        placeholder="设置用户名"
        required
        v-model="registerForm.username"
      />
    </div>
    <div class="input-group">
      <input
        type="email"
        id="reg-email"
        class="form-input"
        placeholder="输入邮箱"
        required
        v-model="registerForm.email"
      />
    </div>
    <div class="input-group verification-code-group">
      <input
        type="text"
        id="reg-code"
        class="form-input"
        placeholder="验证码"
        required
        v-model="registerForm.code"
      />
      <button
        type="button"
        @click="getVerificationCode"
        class="get-code-btn"
        :disabled="isCounting"
      >
        {{ countdown > 0 ? `${countdown}s` : '获取验证码' }}
      </button>
    </div>
    <div class="input-group">
      <input
        type="password"
        id="reg-password"
        class="form-input"
        placeholder="设置密码"
        required
        v-model="registerForm.password"
      />
    </div>
    <div class="input-group">
      <input
        type="password"
        id="reg-confirm-password"
        class="form-input"
        placeholder="确认密码"
        required
        v-model="registerForm.confirmPassword"
      />
    </div>

    <p v-if="errorMsg" class="error-text">{{ errorMsg }}</p>

    <!-- 新增：协议勾选区域 -->
    <div class="agreement-group">
      <input type="checkbox" id="agreement" v-model="agreedToTerms" />
      <label for="agreement">
        我已阅读并同意
        <a href="#" target="_blank">服务条款</a> 和
        <a href="#" target="_blank">隐私政策</a>
      </label>
    </div>

    <!-- 按钮的 disabled 状态与协议勾选绑定 -->
    <button type="submit" class="form-button" :disabled="!agreedToTerms || isLoading">
      {{ isLoading ? '注册中...' : '立即注册' }}
    </button>
  </form>
</template>

<!-- RegisterForm.vue (修改后) -->
<script setup lang="js">
import { reactive, ref } from 'vue'
import axios from 'axios' // 导入 axios

// 定义 API 基础 URL
const API_BASE_URL = 'http://localhost:8080/api'

const agreedToTerms = ref(false)
const emit = defineEmits(['submit-success'])
const registerForm = reactive({
  username: '',
  email: '',
  code: '', // 对应后端的 captcha
  password: '',
  confirmPassword: '',
})
const errorMsg = ref('')
const countdown = ref(0)
const isCounting = ref(false)
const isLoading = ref(false) // 新增：用于表示正在提交

const getVerificationCode = async () => {
  if (isCounting.value) return
  if (!/^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/.test(registerForm.email)) {
    errorMsg.value = '请输入有效的邮箱地址！'
    return
  }
  errorMsg.value = ''
  isCounting.value = true
  countdown.value = 60

  try {
    // 调用后端 API 发送验证码
    await axios.post(`${API_BASE_URL}/captcha`, { email: registerForm.email })
    // 后端会在控制台打印验证码，开发时查看控制台即可
    console.log('验证码请求已发送，请检查后端控制台获取验证码。')

    const timer = setInterval(() => {
      countdown.value--
      if (countdown.value <= 0) {
        clearInterval(timer)
        isCounting.value = false
      }
    }, 1000)
  } catch (error) {
    errorMsg.value = error.response?.data?.msg || '发送验证码失败，请稍后再试。'
    isCounting.value = false // 失败后重置按钮
    countdown.value = 0
  }
}

const handleRegister = async () => {
  if (registerForm.password !== registerForm.confirmPassword) {
    errorMsg.value = '两次输入的密码不一致！'
    return
  }
  if (registerForm.password.length < 6) {
    errorMsg.value = '密码长度不能少于6位！'
    return
  }
  errorMsg.value = ''
  isLoading.value = true // 开始加载

  try {
    const payload = {
      username: registerForm.username,
      email: registerForm.email,
      password: registerForm.password,
      captcha: registerForm.code, // 字段名与后端对应
    }
    // 调用后端注册 API
    const response = await axios.post(`${API_BASE_URL}/register`, payload)

    // 注册成功，通知父组件
    emit('submit-success', response.data.data)
  } catch (error) {
    // 从后端响应中获取更具体的错误信息
    errorMsg.value = error.response?.data?.msg || '注册失败，请检查您的输入。'
  } finally {
    isLoading.value = false // 结束加载
  }
}
</script>

<style scoped>
/* 新增：协议组的样式 */
.agreement-group {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  margin-bottom: 1.5rem;
  font-size: 0.8rem;
  color: var(--text-light);
}
.agreement-group input[type='checkbox'] {
  /* 美化一下复选框 */
  width: 16px;
  height: 16px;
  cursor: pointer;
}
.agreement-group a {
  color: var(--primary-color);
  text-decoration: none;
}
.agreement-group a:hover {
  text-decoration: underline;
}

/* 其他样式 */
.verification-code-group {
  display: flex;
  gap: 10px;
}
.verification-code-group .form-input {
  flex-grow: 1;
}
.get-code-btn {
  padding: 0 15px;
  border: 1px solid var(--border-color);
  background-color: var(--input-bg);
  color: var(--primary-color);
  border-radius: 0.5rem;
  cursor: pointer;
  white-space: nowrap;
  transition: background-color 0.3s;
}
.get-code-btn:hover:not(:disabled) {
  background-color: #e5e7eb;
}
.get-code-btn:disabled {
  cursor: not-allowed;
  background-color: #e9e9e9;
  color: var(--text-light);
}
.error-text {
  color: #ef4444;
  font-size: 0.875rem;
  margin-bottom: 1rem;
  text-align: left;
}
</style>
