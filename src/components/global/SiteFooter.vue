<template>
  <footer class="site-footer">
    <div class="footer-container">
      <!-- 上部分：关注我们 -->
      <div class="social-section">
        <span class="social-title">关注我们</span>
        <div class="social-icons">
          <a href="https://wx.qq.com/" aria-label="微信公众号" class="social-link" target="_blank">
            <img src="@/assets/global/SiteFooter/WeChat_public.svg" class="icon-placeholder" />
          </a>
          <a href="https://www.douyin.com/" aria-label="抖音" class="social-link" target="_blank">
            <img src="@/assets/global/SiteFooter/TikTok_CN.svg" class="icon-placeholder" />
          </a>
          <a
            href="https://www.bilibili.com/"
            aria-label="bilibili"
            class="social-link"
            target="_blank"
          >
            <img src="@/assets/global/SiteFooter/Bilibili.svg" class="icon-placeholder" />
          </a>
        </div>
      </div>

      <hr class="divider" />

      <!-- 中间部分：联系我们 -->
      <div class="contact-section">
        <div class="contact-info">
          <div class="contact-column">
            <h4 class="section-title">联系方式</h4>
            <ul class="contact-list">
              <li class="contact-item">
                <span class="contact-text">电话：xxxxx</span>
              </li>
              <li class="contact-item">
                <span class="contact-text">邮箱：2594817591@qq.com</span>
              </li>
              <li class="contact-item">
                <span class="contact-text">邮编：xxxxx</span>
              </li>
            </ul>
          </div>

          <div class="contact-column">
            <h4 class="section-title">中心地址</h4>
            <ul class="contact-list">
              <li class="contact-item">
                <span class="contact-text">xxxxx</span>
              </li>
              <li class="contact-item">
                <span class="contact-text">xxxxx</span>
              </li>
              <li class="contact-item">
                <span class="contact-text">xxxxx</span>
              </li>
            </ul>
          </div>
        </div>

        <div class="logo-section">
          <div class="logo-container">
            <span class="logo-text">虚拟仿真中心</span>
            <span class="logo-english">Virtual Simulation Center Of WUST</span>
          </div>
        </div>
      </div>

      <hr class="divider" />

      <!-- 底部：版权信息 -->
      <div class="copyright-section">
        <p class="copyright-text">
          Copyright © 虚拟仿真中心 - {{ currentYear }} Virtual Simulation Center. All Rights
          Reserved.
        </p>
      </div>
    </div>
  </footer>
</template>

<script setup>
import { computed } from 'vue'

const currentYear = computed(() => new Date().getFullYear())
</script>

<style scoped>
/* ===== 主容器样式 ===== */
.site-footer {
  background: linear-gradient(135deg, #f8f9fa 0%, #e9ecef 100%);
  color: #5f6368;
  padding: 48px 20px 32px;
  font-size: 14px;
  font-family: 'Helvetica Neue', Arial, sans-serif;
  border-top: 1px solid #e1e5e9;
}

.footer-container {
  max-width: 1200px;
  margin: 0 auto;
}

/* 分割线样式 */
.divider {
  border: none;
  border-top: 1px solid #ccd2d8;
  margin: 32px 0;
  opacity: 0.8;
}

/* ===== 社交媒体部分 ===== */
.social-section {
  display: flex;
  align-items: center;
  gap: 24px;
  margin-bottom: 8px;
}

.social-title {
  font-size: 16px;
  font-weight: 600;
  color: black;
}

.social-icons {
  display: flex;
  gap: 12px;
}

.social-link {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  background: transparent;
  border-radius: 8px;
  text-decoration: none;
  font-weight: 600;
  font-size: 12px;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  /* filter: drop-shadow(0 2px 8px rgba(74, 144, 226, 0.2)); */
}

.social-link:hover {
  transform: translateY(-2px);
  filter: drop-shadow(0 4px 16px rgba(74, 144, 226, 0.8));
}

.icon-placeholder {
  width: 30px;
  height: 30px;
}

/* ===== 联系我们部分 ===== */
.contact-section {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 48px;
}

.contact-info {
  display: flex;
  gap: 80px;
  flex: 1;
}

.contact-column {
  flex: 1;
}

.section-title {
  font-size: 16px;
  font-weight: 600;
  color: black;
  margin: 0 0 16px 0;
  padding-bottom: 8px;
  border-bottom: 2px solid #4a90e2;
  display: inline-block;
}

.contact-list {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.contact-item {
  line-height: 1.6;
}

.contact-text {
  color: #6c757d;
  font-weight: 400;
}

/* ===== Logo部分 ===== */
.logo-section {
  display: flex;
  align-items: center;
  justify-content: flex-end;
}

.logo-container {
  text-align: right;
}

.logo-text {
  display: block;
  font-size: 28px;
  font-weight: 700;
  color: gray;
  margin-bottom: 4px;
  background: linear-gradient(135deg, #2c3e50, #4a90e2);
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.logo-english {
  display: block;
  font-size: 14px;
  font-weight: 500;
  color: #6c757d;
  text-transform: uppercase;
  letter-spacing: 1px;
}

/* ===== 版权信息部分 ===== */
.copyright-section {
  text-align: center;
}

.copyright-text {
  margin: 0;
  font-size: 12px;
  color: #9aa0a6;
  font-weight: 400;
}

/* ===== 响应式设计 ===== */
@media (max-width: 768px) {
  .site-footer {
    padding: 32px 16px 24px;
  }

  .contact-section {
    flex-direction: column;
    gap: 32px;
  }

  .contact-info {
    gap: 40px;
  }

  .logo-section {
    justify-content: center;
  }

  .logo-container {
    text-align: center;
  }

  .social-section {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }

  .social-icons {
    justify-content: center;
    width: 100%;
  }

  .divider {
    margin: 24px 0;
  }
}

@media (max-width: 480px) {
  .contact-info {
    flex-direction: column;
    gap: 24px;
  }

  .social-icons {
    gap: 8px;
  }

  .social-link {
    width: 32px;
    height: 32px;
    font-size: 11px;
  }
}
</style>
