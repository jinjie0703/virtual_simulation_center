<template>
  <footer class="site-footer">
    <div class="footer-container">
      <!-- 上部分：关注我们 -->
      <div class="social-section">
        <span class="social-title">关注我们</span>
        <div class="social-icons">
          <!-- 注意：这里用文字模拟图标，实际项目中建议替换为 SVG 图标 -->
          <a href="#" aria-label="微信公众号"><span class="icon-placeholder">x</span></a>
          <a href="#" aria-label="视频号"><span class="icon-placeholder">x</span></a>
          <a href="#" aria-label="腾讯视频"><span class="icon-placeholder">x</span></a>
          <a href="#" aria-label="Bilibili"><span class="icon-placeholder">x</span></a>
        </div>
      </div>

      <hr class="divider" />

      <!-- 中间部分：链接和Logo -->
      <div class="links-section">
        <div class="links-group">
          <div class="links-column">
            <h4>xxxx</h4>
            <ul>
              <li><a href="#">xxxx</a></li>
              <li><a href="#">xxxx</a></li>
              <li><a href="#">xxxx</a></li>
            </ul>
          </div>
          <div class="links-column">
            <h4>xxxx</h4>
            <ul>
              <li><a href="#">xxxx</a></li>
              <li><a href="#">xxxx</a></li>
              <li><a href="#">xxxx</a></li>
            </ul>
          </div>
        </div>
        <div class="logo-section">
          <span class="logo-text">xxxx</span>
          <span class="logo-chinese">xxxx</span>
        </div>
      </div>

      <hr class="divider" />

      <!-- 底部：版权信息 -->
      <div class="copyright-section">
        <p>Copyright © xxxx - {{ currentYear }} xxxx. All Rights Reserved.</p>
      </div>
    </div>
  </footer>
</template>

<script setup>
import { computed } from 'vue'

const currentYear = computed(() => new Date().getFullYear())
</script>

<style scoped>
.site-footer {
  background-color: #f5f5f7;
  color: #5f6368;
  padding: 40px 20px;
  font-size: 14px;
}

.footer-container {
  max-width: 1200px;
  margin: 0 auto;
}

.divider {
  border: none;
  border-top: 1px solid #e0e0e0;
  margin: 25px 0;
}

/* 社交部分 */
.social-section {
  display: flex;
  align-items: center;
  gap: 20px;
}
.social-title {
  font-size: 16px;
  color: #202124;
}
.social-icons {
  display: flex;
  gap: 12px;
}
.social-icons a {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  background-color: #3c4043;
  border-radius: 6px;
  color: white;
  text-decoration: none;
  font-weight: bold;
  transition: opacity 0.3s;
}
.social-icons a:hover {
  opacity: 0.85;
}

/* 链接和Logo部分 */
.links-section {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 40px;
}
.links-group {
  display: flex;
  gap: 80px;
}
.links-column h4 {
  font-size: 16px;
  color: #202124;
  margin: 0 0 15px 0;
}
.links-column ul {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 12px;
}
.links-column a {
  color: #5f6368;
  text-decoration: none;
}
.links-column a:hover {
  text-decoration: underline;
}

.logo-section {
  display: flex;
  align-items: center;
}
.logo-text {
  font-size: 28px;
  font-weight: bold;
  color: #202124;
}
.logo-chinese {
  font-size: 24px;
  font-weight: bold;
  color: #5f6368;
  margin-left: 8px;
}

/* 版权部分 */
.copyright-section p {
  margin: 0;
  font-size: 12px;
  color: #9aa0a6;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .links-section {
    flex-direction: column;
  }
  .logo-section {
    margin-top: 20px;
  }
  .links-group {
    gap: 40px;
  }
  .social-section {
    flex-direction: column;
    align-items: flex-start;
  }
}
</style>
