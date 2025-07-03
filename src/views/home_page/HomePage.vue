<template>
  <div>
    <!-- 1. 轮播组件，通过 :slides 传入数据 -->
    <Carousel :slides="carouselData" />

    <main>
      <!-- 2. 成果展示组件，通过 :features 传入数据 -->
      <FeatureSection :features="featureData" />

      <!-- 3. 新闻与竞赛组件，传入对应的数据 -->
      <UpdatesSection :news-items="newsData" :competition-items="competitionData" />
    </main>
  </div>
</template>

<script setup>
import { ref } from 'vue'

// 导入子组件
import Carousel from '@/components/home_page/HomeCarousel.vue'
import FeatureSection from '@/components/home_page/FeatureSection.vue'
import UpdatesSection from '@/components/home_page/UpdatesSection.vue'

// 1. 轮播图的数据
const carouselData = ref([
  {
    title: '腾讯公布二零二五年第一季业绩',
    imageUrl:
      'https://images.unsplash.com/photo-1542370285-b8eb8317691c?q=80&w=1974&auto=format&fit=crop',
  },
  {
    title: '科技向善，创造无限可能',
    imageUrl:
      'https://images.unsplash.com/photo-1519010470956-6d877008eaa4?q=80&w=2070&auto=format&fit=crop',
  },
  {
    title: '连接你我，共创美好未来',
    imageUrl:
      'https://images.unsplash.com/photo-1611606063065-ee7946f0b343?q=80&w=1974&auto=format&fit=crop',
  },
])

// 2. 成果展示的数据
const featureData = ref([
  {
    title: '创新技术驱动',
    description:
      '我们致力于前沿科技的研发与应用，从人工智能到云计算，不断探索技术的边界，为各行各业提供强大的数字化引擎。',
    imageUrl: '',
  },
  {
    title: '连接亿万用户',
    description:
      '通过多元化的社交与内容平台，我们构建了强大的连接网络。无论是生活、工作还是娱乐，我们都让沟通变得更便捷。',
    imageUrl: '',
  },
  {
    title: '赋能合作伙伴',
    description:
      '我们打造了开放共赢的生态系统，为合作伙伴提供流量、技术、数据等全方位支持，共同成长，共创价值。',
    imageUrl: '',
  },
  {
    title: '践行社会责任',
    description:
      '科技的价值在于“向善”。我们积极投身于公益事业、环境保护和青少年成长等领域，利用科技的力量解决社会问题。',
    imageUrl: '',
  },
])

// 3. 新闻与竞赛的数据
const newsData = ref([
  { title: '我院团队在国际人工智能大会上发表重要研究成果', date: '2024-05-20', url: '#' },
  { title: '2024年度“未来学者”夏令营成功举办', date: '2024-05-15', url: '#' },
  { title: '学院与XYZ公司达成战略合作，共建联合实验室', date: '2024-05-10', url: '#' },
  { title: '知名教授李明加入我院，将领导数字媒体研究方向', date: '2024-05-05', url: '#' },
  { title: '“代码之光”编程马拉松圆满落幕', date: '2024-04-28', url: '#' },
])

const competitionData = ref([
  { title: '关于举办第十届“算法之星”程序设计竞赛的通知', date: '2024-05-18', url: '#' },
  { title: '2024年全国大学生信息安全竞赛报名启动', date: '2024-05-12', url: '#' },
  { title: '“挑战杯”大学生创业计划竞赛校内选拔赛结果公示', date: '2024-05-09', url: '#' },
  { title: '“蓝桥杯”全国软件和信息技术专业人才大赛省赛获奖喜报', date: '2024-05-02', url: '#' },
  { title: '国际大学生数学建模竞赛（ICM）赛前动员会', date: '2024-04-25', url: '#' },
])
</script>

<style>
/* 可以在这里放一些全局样式，或者在 main.css 中定义 */
body {
  margin: 0;
  font-family:
    -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  background-color: #f4f4f4;
  color: #333;
}
main {
  display: block; /* 确保 main 标签的兼容性 */
}
</style>
