// src/mock-data.js

// 导出一份“高保真”的单篇文章模拟数据
export const mockArticle = {
  id: 'news-1',
  title: '标题',
  author: '',
  publishDate: '',
  category: '类别',
  tags: ['tag1', 'tag2', 'tag3', 'tag4'],
  content: `
## 1. 项目背景与挑战

随着信息技术的飞速发展，传统的教学模式正面临深刻变革。我校虚拟仿真中心自成立以来，始终致力于探索如何利用 **VR (虚拟现实)** 与 **AR (增强现实)** 技术，为学生打造沉浸式、高效率的学习环境。

### 1.1 面临的挑战
*   实验设备昂贵，难以普及到每位学生。
*   高危或极端环境下的实验无法真实开展。
*   部分抽象理论难以通过书本直观理解。

![虚拟现实技术](https://images.unsplash.com/photo-1554415707-6e8cfc93fe23?q=80&w=1200&auto=format&fit=crop)
*图片来源: Unsplash*

---

## 2. 我们的解决方案

我们自主研发了一套名为 **“V-Lab”** 的虚拟实验平台。学生只需佩戴VR头显，即可进入一个与真实世界1:1的虚拟实验室。

**支持的文件下载：**
*   [V-Lab 项目白皮书.pdf](/downloads/v-lab-whitepaper.pdf)
*   [项目演示视频链接](https://www.youtube.com/watch?v=dQw4w9WgXcQ)

> “这不仅仅是技术的胜利，更是教育理念的革新。” — 项目负责人

## 3. 取得的成就

平台上线至今，已成功服务超过 **10,000人次** 的学生，有效提升了学生的学习兴趣和实践能力。此次荣获国家级教学成果奖，是对我们团队工作的最大肯定。
  `,
}

// 同时也导出您在InfoCenter中使用的列表数据，方便管理
export const allMockData = {
  news: Array.from({ length: 250 }, (_, i) => ({
    id: `news-${i + 1}`,
    title: `我校虚拟仿真中心获得第${i + 1}项国家级教学成果奖`,
    summary: '中心凭借其在沉浸式教学领域的创新实践，再次荣获国家级表彰，引领了教育技术的新潮流。',
    date: `2024-05-${25 - i > 0 ? 25 - i : 1}`,
    category: '学术荣誉',
  })),
  competitions: Array.from({ length: 18 }, (_, i) => ({
    id: `comp-${i + 1}`,
    title: `第${i + 1}届全国大学生虚拟现实设计大赛`,
    summary: '本次大赛旨在激发学生的创新精神和实践能力，推动VR技术在各行业的应用与发展。',
    deadline: '2024-07-30',
    level: '国家级',
    tags: ['AI', 'Vue.js', 'Python'],
  })),
  projects: Array.from({ length: 12 }, (_, i) => ({
    id: `proj-${i + 1}`,
    title: `开发一个基于AI的个性化学习路径推荐系统 #${i + 1}`,
    summary: '需要利用机器学习算法，根据学生的学习行为和成绩，动态生成最优学习路径。',
    deadline: '2024-08-15',
    tags: ['AI', 'Vue.js', 'Python'],
  })),
}
