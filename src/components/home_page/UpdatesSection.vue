<template>
  <section class="latest-updates">
    <div class="container">
      <!-- 增强的标题区域 -->
      <div class="section-header">
        <h2 class="section-title">
          <span class="title-icon">📰</span>
          最新资讯
        </h2>
        <p class="section-subtitle">掌握第一手学院动态与竞赛信息</p>
        <div class="title-decoration"></div>
      </div>

      <div class="updates-grid">
        <!-- 学院新闻 -->
        <div class="update-column news-column">
          <div class="column-header">
            <div class="header-icon">
              <img
                src="@/assets/home_page/UpdatesSection/news-icon.svg"
                alt="新闻图标"
                class="news-icon"
              />
            </div>
            <div class="header-content">
              <h3>学院新闻</h3>
              <span class="update-count">{{ newsItems.length }} 条更新</span>
            </div>
            <button class="more-btn">
              <span>查看更多</span>
              <i class="arrow">→</i>
            </button>
          </div>

          <div class="update-list-container">
            <ul class="update-list">
              <li
                v-for="(item, index) in newsItems"
                :key="item.id"
                class="update-item"
                :style="{ animationDelay: `${index * 0.1}s` }"
              >
                <router-link
                  :to="{ name: 'NewsDetail', params: { id: item.id } }"
                  class="item-link"
                >
                  <div class="item-indicator"></div>
                  <div class="item-content">
                    <span class="item-title">{{ item.title }}</span>
                    <div class="item-meta">
                      <span class="item-date">{{ formatDate(item.date) }}</span>
                      <span class="item-badge new" v-if="isNew(item.date)">新</span>
                      <span class="item-badge hot" v-if="isHot(index)">热</span>
                    </div>
                  </div>
                  <div class="item-arrow">
                    <i>➤</i>
                  </div>
                </router-link>
              </li>
            </ul>
          </div>
        </div>

        <!-- 竞赛信息 -->
        <div class="update-column competition-column">
          <div class="column-header">
            <div class="header-icon competition-icon">
              <img
                src="@/assets/home_page/UpdatesSection/competition-icon.svg"
                alt="竞赛信息"
                class="competition-icon"
              />
            </div>
            <div class="header-content">
              <h3>竞赛信息</h3>
              <span class="update-count">{{ competitionItems.length }} 条更新</span>
            </div>
            <button class="more-btn">
              <span>查看更多</span>
              <i class="arrow">→</i>
            </button>
          </div>

          <div class="update-list-container">
            <ul class="update-list">
              <li
                v-for="(item, index) in competitionItems"
                :key="item.id"
                class="update-item"
                :style="{ animationDelay: `${index * 0.1}s` }"
              >
                <router-link
                  :to="{ name: 'CompetitionDetail', params: { id: item.id } }"
                  class="item-link"
                >
                  <div class="item-indicator"></div>
                  <div class="item-content">
                    <span class="item-title">{{ item.title }}</span>
                    <div class="item-meta">
                      <span class="item-date">{{ formatDate(item.date) }}</span>
                      <span class="item-badge new" v-if="isNew(item.date)">新</span>
                      <span class="item-badge urgent" v-if="isUrgent(index)">急</span>
                    </div>
                  </div>
                  <div class="item-arrow">
                    <i>➤</i>
                  </div>
                </router-link>
              </li>
            </ul>
          </div>
        </div>
      </div>

      <!-- 底部装饰 -->
      <div class="section-footer">
        <div class="footer-decoration">
          <div class="decoration-dot"></div>
          <div class="decoration-line"></div>
          <div class="decoration-dot"></div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
//import { computed } from 'vue'

defineProps({
  newsItems: { type: Array, required: true },
  competitionItems: { type: Array, required: true },
})

// 日期格式化
const formatDate = (dateStr) => {
  const date = new Date(dateStr)
  const now = new Date()
  const diffTime = Math.abs(now - date)
  const diffDays = Math.floor(diffTime / (1000 * 60 * 60 * 24))

  if (diffDays === 0) return '今天'
  if (diffDays === 1) return '昨天'
  if (diffDays < 7) return `${diffDays}天前`

  return date.toLocaleDateString('zh-CN', {
    month: 'short',
    day: 'numeric',
  })
}

// 判断是否为新内容（3天内）
const isNew = (dateStr) => {
  const date = new Date(dateStr)
  const now = new Date()
  const diffDays = Math.floor((now - date) / (1000 * 60 * 60 * 24))
  return diffDays <= 3
}

// 判断是否为热门（前2条）
const isHot = (index) => {
  return index < 2
}

// 判断是否紧急（前1条）
const isUrgent = (index) => {
  return index < 1
}
</script>

<style scoped>
.latest-updates {
  padding: 100px 0;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  position: relative;
  overflow: hidden;
}

.latest-updates::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: url('data:image/svg+xml,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100"><circle cx="20" cy="20" r="2" fill="white" opacity="0.1"/><circle cx="80" cy="40" r="1" fill="white" opacity="0.1"/><circle cx="40" cy="80" r="1.5" fill="white" opacity="0.1"/></svg>')
    repeat;
  background-size: 100px 100px;
  animation: float 20s infinite linear;
}

@keyframes float {
  0% {
    transform: translateX(0);
  }
  100% {
    transform: translateX(-100px);
  }
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
  position: relative;
  z-index: 1;
}

/* 标题区域 */
.section-header {
  text-align: center;
  margin-bottom: 60px;
}

.section-title {
  font-size: 3rem;
  font-weight: 700;
  color: white;
  margin: 0 0 16px 0;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 16px;
}

.title-icon {
  font-size: 2.5rem;
  filter: drop-shadow(0 4px 8px rgba(0, 0, 0, 0.2));
}

.section-subtitle {
  font-size: 1.2rem;
  color: rgba(255, 255, 255, 0.9);
  margin-bottom: 20px;
}

.title-decoration {
  width: 100px;
  height: 4px;
  background: linear-gradient(90deg, transparent 0%, white 50%, transparent 100%);
  margin: 0 auto;
  border-radius: 2px;
}

/* 网格布局 */
.updates-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 40px;
  margin-bottom: 60px;
}

.update-column {
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(20px);
  border-radius: 24px;
  overflow: hidden;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.1);
  border: 1px solid rgb(255, 255, 255);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
}

.update-column::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #667eea 0%, #764ba2 100%);
}

.news-column::before {
  background: linear-gradient(90deg, #4facfe 0%, #00f2fe 100%);
}

.competition-column::before {
  background: linear-gradient(90deg, #fa709a 0%, #fee140 100%);
}

.update-column:hover {
  transform: translateY(-12px);
  box-shadow: 0 30px 80px rgba(var(--black), 0.2);
}

/* 列标题 */
.column-header {
  padding: 32px 32px 24px;
  display: flex;
  align-items: center;
  gap: 16px;
  border-bottom: 1px solid rgba(var(--black), 0.08);
  background: rgba(var(--white), 0.9);
}

.header-icon {
  width: 56px;
  height: 56px;
  border-radius: 16px;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 8px 24px rgba(var(--black), 0.2);
  flex-shrink: 0;
}

.news-icon {
  width: 56px;
  height: 56px;
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
}

.competition-icon {
  width: 56px;
  height: 56px;
  background: linear-gradient(135deg, #fa709a 0%, #fee140 100%);
}
/* 自适应宽度的弹性项目 */
.header-content {
  flex: 1;
}

.header-content h3 {
  font-size: 1.5rem;
  font-weight: 600;
  color: black;
  margin: 0 0 4px 0;
}

.update-count {
  font-size: 0.9rem;
  color: grey;
  font-weight: 500;
}

.more-btn {
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--secondary-color) 100%);
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 20px;
  font-size: 0.9rem;
  font-weight: 500;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
  transition: all 0.3s ease;
  box-shadow: 0 4px 12px rgba(var(--UpdatesSection-primary-color-rgb), 0.3);
}

.more-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(var(--UpdatesSection-primary-color-rgb), 0.5);
}

.more-btn .arrow {
  transition: transform 0.3s ease;
}

.more-btn:hover .arrow {
  transform: translateX(4px);
}

/* 列表容器 */
.update-list-container {
  max-height: 480px;
  overflow-y: auto;
  scrollbar-width: thin;
  scrollbar-color: rgba(var(--UpdatesSection-primary-color-rgb), 0.3) transparent;
}

.update-list-container::-webkit-scrollbar {
  width: 6px;
}

.update-list-container::-webkit-scrollbar-track {
  background: transparent;
}

.update-list-container::-webkit-scrollbar-thumb {
  background: rgba(var(--UpdatesSection-primary-color-rgb), 0.3);
  border-radius: 3px;
}

.update-list-container::-webkit-scrollbar-thumb:hover {
  background: rgba(var(--UpdatesSection-primary-color-rgb), 0.5);
}

/* 列表样式 */
.update-list {
  list-style: none;
  padding: 0;
  margin: 0;
  /* 删除底部滚动条 */
  overflow: hidden;
}

.update-item {
  opacity: 0;
  transform: translateY(100px);
  animation: slideInUp 1s ease forwards;
  /* 为放大和阴影效果提供垂直和水平空间 */
  /* padding: 8px; */
  /* 抵消 padding 对布局的影响 */
  /* margin: -8px; */
}

@keyframes slideInUp {
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.item-link {
  display: flex;
  align-items: center;
  padding: 20px 32px;
  text-decoration: none;
  color: inherit;
  transition: all 0.3s ease;
  border-bottom: 1px solid rgba(var(--black), 0.05);
  position: relative;
  overflow: hidden;
}

.item-link::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(
    90deg,
    transparent 0%,
    rgba(var(--UpdatesSection-primary-color-rgb), 0.05) 50%,
    transparent 100%
  );
  transition: left 1s ease;
}

.item-link:hover::before {
  left: 100%;
}

.item-link:hover {
  background: rgba(var(--UpdatesSection-primary-color-rgb), 0.2);
  transform: translateX(8px);
}

/* 首先找到 class="update-list" 的 <ul> 元素。
然后，通过 > (子代选择器) 和 :last-child 伪类，它精确地选中了作为 <ul> 最后一个直接子元素的那个 <li> (.update-item)。
最后，通过空格（后代选择器），它会选中这个特定的最后一个 <li> 内部的 <a> 元素 (.item-link)。 */
.update-list > .update-item:last-child .item-link {
  border-bottom: none;
}

/* item左边那个点 */
.item-indicator {
  width: 8px;
  height: 8px;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--secondary-color) 100%);
  border-radius: 50%;
  margin-right: 16px;
  flex-shrink: 0;
  box-shadow: 0 0 0 3px rgba(var(--UpdatesSection-primary-color-rgb), 0.2);
  transition: all 0.3s ease;
}

.item-link:hover .item-indicator {
  transform: scale(1.2);
  box-shadow: 0 0 0 6px rgba(var(---UpdatesSection-primary-color-rgb), 0.3);
}

.item-content {
  flex: 1;
  min-width: 0;
}

.item-title {
  display: block;
  font-size: 1.05rem;
  font-weight: 500;
  color: black;
  margin-bottom: 8px;
  white-space: nowrap;
  overflow: hidden;
  /* 文本溢出处理方式 */
  text-overflow: ellipsis;
  transition: color 0.3s ease;
}

.item-link:hover .item-title {
  color: var(--primary-color);
}

.item-meta {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-wrap: wrap;
}

.item-date {
  font-size: 0.85rem;
  color: var(--item-date-color);
  font-weight: 500;
}

.item-badge {
  padding: 2px 8px;
  border-radius: 10px;
  font-size: 0.7rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.item-badge.new {
  background: linear-gradient(135deg, #48bb78 0%, #38a169 100%);
  color: white;
  box-shadow: 0 2px 4px rgba(72, 187, 120, 0.3);
  animation: pulse 1s infinite;
}

.item-badge.hot {
  background: linear-gradient(135deg, #ed8936 0%, #dd6b20 100%);
  color: white;
  box-shadow: 0 2px 4px rgba(237, 137, 54, 0.3);
  animation: pulse 1s infinite;
}

.item-badge.urgent {
  background: linear-gradient(135deg, #f56565 0%, #e53e3e 100%);
  color: white;
  box-shadow: 0 2px 4px rgba(245, 101, 101, 0.3);
  animation: pulse 1s infinite;
}

@keyframes pulse {
  0%,
  100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.1);
  }
}

.item-arrow {
  margin-left: 12px;
  color: var(--gray);
  font-size: 1.1rem;
  transition: all 0.3s ease;
  opacity: 0;
  transform: translateX(-10px);
}

.item-link:hover .item-arrow {
  opacity: 1;
  transform: translateX(0);
  color: var(--primary-color);
}

/* 底部装饰 */
.section-footer {
  display: flex;
  justify-content: center;
}

.footer-decoration {
  display: flex;
  align-items: center;
  gap: 12px;
}

.decoration-dot {
  width: 8px;
  height: 8px;
  background: rgba(255, 255, 255, 0.6);
  border-radius: 50%;
}

.decoration-line {
  width: 60px;
  height: 3px;
  background: linear-gradient(
    90deg,
    transparent 0%,
    rgba(255, 255, 255, 0.6) 50%,
    transparent 100%
  );
}

/* 响应式设计 */
@media (max-width: 768px) {
  .updates-grid {
    grid-template-columns: 1fr;
    gap: 24px;
  }

  .section-title {
    font-size: 2.2rem;
  }

  .column-header {
    padding: 24px 20px 20px;
    flex-wrap: wrap;
  }

  .more-btn {
    order: 3;
    margin-top: 12px;
    width: 100%;
    justify-content: center;
  }

  .item-link {
    padding: 16px 20px;
  }

  .header-content h3 {
    font-size: 1.3rem;
  }
}

@media (max-width: 480px) {
  .latest-updates {
    padding: 60px 0;
  }

  .section-header {
    margin-bottom: 40px;
  }

  .section-title {
    font-size: 1.8rem;
    flex-direction: column;
    gap: 8px;
  }

  .update-column {
    border-radius: 16px;
  }

  .item-title {
    font-size: 1rem;
  }
}
</style>
