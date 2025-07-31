<template>
  <div class="empty-state">
    <div class="empty-icon">👥</div>
    <h3>暂无成员信息</h3>
    <p>当前分类下还没有添加成员</p>
  </div>
</template>

<style scoped>
.empty-state {
  text-align: center;
  padding: 80px 20px;
  color: #666;
  min-height: 400px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.empty-icon {
  font-size: 4rem;
  margin-bottom: 20px;
  opacity: 0.5;
}

.empty-state h3 {
  font-size: 1.5rem;
  margin-bottom: 10px;
  color: #333;
}

.empty-state p {
  font-size: 1rem;
  opacity: 0.8;
}
</style>
