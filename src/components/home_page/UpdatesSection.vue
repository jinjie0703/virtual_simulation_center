<template>
  <section class="latest-updates">
    <div class="container">
      <div class="updates-grid">
        <!-- 学院新闻 -->
        <div class="update-column">
          <h3>学院新闻</h3>
          <ul class="update-list">
            <li v-for="item in newsItems" :key="item.title">
              <a :href="item.url">
                <span class="item-title">{{ item.title }}</span>
                <span class="item-date">{{ item.date }}</span>
              </a>
            </li>
          </ul>
        </div>

        <!-- 竞赛信息 -->
        <div class="update-column">
          <h3>竞赛信息</h3>
          <ul class="update-list">
            <li v-for="item in competitionItems" :key="item.title">
              <a :href="item.url">
                <span class="item-title">{{ item.title }}</span>
                <span class="item-date">{{ item.date }}</span>
              </a>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
defineProps({
  newsItems: { type: Array, required: true },
  competitionItems: { type: Array, required: true },
})
</script>

<style scoped>
.latest-updates {
  padding: 80px 0;
  background-color: #f4f4f4;
}
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}
.updates-grid {
  display: flex;
  justify-content: space-between;
  gap: 40px;
}
.update-column {
  flex: 1;
  min-width: 0;
  background-color: #fff;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
}
.update-column h3 {
  font-size: 1.8rem;
  margin-top: 0;
  margin-bottom: 20px;
  border-bottom: 2px solid #0052d9;
  padding-bottom: 10px;
}
.update-list {
  list-style: none;
  padding: 0;
  margin: 0;
}
.update-list li {
  padding: 15px 0;
  border-bottom: 1px solid #e9e9e9;
}
.update-list li:last-child {
  border-bottom: none;
}
.update-list li a {
  display: flex;
  justify-content: space-between;
  align-items: center;
  text-decoration: none;
  color: #333;
  transition: color 0.3s;
}
.update-list li a:hover .item-title {
  color: #0052d9;
}
.item-title {
  flex-grow: 1;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  margin-right: 20px;
  font-size: 1.1rem;
}
.item-date {
  flex-shrink: 0;
  font-size: 0.9rem;
  color: #888;
}
@media (max-width: 768px) {
  .updates-grid {
    flex-direction: column;
  }
}
</style>
