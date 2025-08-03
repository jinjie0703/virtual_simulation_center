<template>
  <div class="showcase-page">
    <AchievementHeader
      title="卓越成果展示"
      subtitle="探索我们如何通过前沿技术，将创新理念转化为具有影响力的教学与科研成果。"
    />
    <main class="content-container">
      <AchievementActions
        v-model="searchQuery"
        v-model:teacherFilter="teacherFilter"
        v-model:tagFilter="tagFilter"
        v-model:timeFilter="timeFilter"
        :teacher-options="teacherOptions"
        :tag-options="tagOptions"
        @open-submission="isSubmissionModalVisible = true"
      />
      <AchievementGrid :achievements="paginatedAchievements" @view-details="openDetailsModal" />
    </main>
    <AchievementDetailModal
      :achievement="selectedAchievement"
      :is-visible="isModalVisible"
      @close="closeDetailsModal"
    />
    <SubmissionModal
      :is-visible="isSubmissionModalVisible"
      @close="isSubmissionModalVisible = false"
      @submit="handleSubmission"
    />
    <AchievementPagination
      v-if="totalPages > 1"
      :current-page="currentPage"
      :total-pages="totalPages"
      @page-changed="goToPage"
    />
  </div>
</template>

<script setup>
import { ref, computed, nextTick } from 'vue'
import AchievementHeader from '@/components/achievement_showcase/AchievementHeader.vue'
import AchievementGrid from '@/components/achievement_showcase/AchievementGrid.vue'
import AchievementDetailModal from '@/components/achievement_showcase/AchievementDetailModal.vue'
import AchievementActions from '@/components/achievement_showcase/AchievementActions.vue'
import SubmissionModal from '@/components/achievement_showcase/SubmissionModal.vue'
import AchievementPagination from '@/components/achievement_showcase/AchievementPagination.vue'

// --- State ---
// const achievementTemplate = {
//   id: 1, // 这个 id 会在循环中被新的、唯一的 id 覆盖
//   title: '虚拟现实教学平台',
//   description:
//     '基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。基于先进的VR技术构建的沉浸式教学环境，为学生提供真实感强、交互性好的学习体验。该平台集成了多种学科内容，支持个性化学习路径定制。',
//   image:
//     'https://images.unsplash.com/photo-1551288049-bebda4e38f71?q=80&w=1200&auto=format&fit=crop',
//   tags: ['VR技术', '教学创新', '沉浸式体验'],
//   author: '张教授',
//   authorTitle: '虚拟现实研究中心主任',
//   authorAvatar:
//     'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?q=80&w=150&auto=format&fit=crop',
//   publishDate: new Date(),
//   contact1: '邮箱zhang@university.edu.cn',
//   contact2: '电话123456789',
//   projectUrl: 'https://github.com',

//   // 视频和图片库
//   videoUrl: 'http://vjs.zencdn.net/v/oceans.mp4',
//   gallery: [
//     'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//     'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//     'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//     'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=2070&auto=format&fit=crop',
//   ],
// }

// // 2. 准备一个空数组来存放生成的数据
// const generatedData = []

// // 3. 循环100次，批量生成数据
// for (let i = 0; i < 1000; i++) {
//   generatedData.push({
//     ...achievementTemplate, // 使用展开运算符(...)来复制模板对象的所有属性
//     id: i + 1, // 关键：为每个新对象设置一个唯一的id，这对于测试和渲染都非常重要
//   })
// }

// const achievements = ref(generatedData)

// --- 数据池 (Data Pools) ---
// 为了创建多样化的数据，我们先准备一些可供选择的内容。

const titles = [
  '基于AI的个性化学习路径推荐系统',
  '虚拟现实（VR）在医学手术模拟中的应用',
  '面向大规模数据处理的分布式计算框架',
  '智能物联网（AIoT）环境监测平台',
  '量子计算在密码学中的应用探索',
  '基于深度学习的自动驾驶感知系统',
  '新一代高能效光子芯片研发',
  '交互式全息投影技术的商业化应用',
  '脑机接口（BCI）技术与残障辅助',
  '面向金融风控的联邦学习模型',
  '数字孪生（Digital Twin）城市管理系统',
  '区块链技术在供应链溯源中的实践',
  '用于蛋白质结构预测的AI模型',
  '低功耗广域物联网（LPWAN）通信协议',
  '增强现实（AR）工业维修指导系统',
  '自然语言处理（NLP）情感分析引擎',
  '智能机器人自主导航与避障算法',
  '高精度室内定位技术研究',
  '基于生成对抗网络（GAN）的艺术创作',
  '面向开发者的云原生API网关',
  '服务于校内社团的活动管理平台',
  '沉浸式虚拟校园导览系统',
  '课堂专注度实时分析系统',
  '校友网络与资源共享平台',
  '智能图书馆导览与书籍推荐机器人',
  '校园二手交易与物品共享平台',
  '基于边缘计算的智能安防系统',
  //  '支持在线协作的Markdown编辑器',
  '电子证书区块链存证系统',
  '失物招领信息智能匹配平台',
]

const authors = [
  {
    name: '李晓明',
    title: '人工智能实验室首席科学家',
    avatar:
      'https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?q=80&w=150&auto=format&fit=crop',
  },
  {
    name: '王静',
    title: 'VR/AR研究中心主任',
    avatar:
      'https://images.unsplash.com/photo-1494790108377-be9c29b29330?q=80&w=150&auto=format&fit=crop',
  },
  {
    name: '张伟',
    title: '数据科学系教授',
    avatar:
      'https://images.unsplash.com/photo-1500648767791-00dcc994a43e?q=80&w=150&auto=format&fit=crop',
  },
  {
    name: '刘敏',
    title: '物联网技术研究员',
    avatar:
      'https://images.unsplash.com/photo-1580489944761-15a19d654956?q=80&w=150&auto=format&fit=crop',
  },
  {
    name: '陈浩',
    title: '量子计算项目负责人',
    avatar:
      'https://images.unsplash.com/photo-1568602471122-7832951cc4c5?q=80&w=150&auto=format&fit=crop',
  },
  {
    name: '赵婷',
    title: '计算机视觉博士后',
    avatar:
      'https://images.unsplash.com/photo-1544005313-94ddf0286df2?q=80&w=150&auto=format&fit=crop',
  },
]

const images = [
  'https://images.unsplash.com/photo-1551288049-bebda4e38f71?q=80&w=1200&auto=format&fit=crop',
  'https://images.unsplash.com/photo-1581092921461-eab62e97a780?q=80&w=1200&auto=format&fit=crop',
  'https://images.unsplash.com/photo-1518770660439-4636190af475?q=80&w=1200&auto=format&fit=crop',
  'https://images.unsplash.com/photo-1517048676732-d65bc937f952?q=80&w=1200&auto=format&fit=crop',
  'https://images.unsplash.com/photo-1531297484001-80022131f5a1?q=80&w=1200&auto=format&fit=crop',
  'https://images.unsplash.com/photo-1611095965934-a3c37553b6f5?q=80&w=1200&auto=format&fit=crop',
  'https://images.unsplash.com/photo-1580894732444-8ec5224a734d?q=80&w=1200&auto=format&fit=crop',
  'https://images.unsplash.com/photo-1526628953301-3e589a6a8b74?q=80&w=1200&auto=format&fit=crop',
  'https://images.unsplash.com/photo-1507146426996-32bb7d2f2d39?q=80&w=1200&auto=format&fit=crop',
  'https://images.unsplash.com/photo-1605810230464-4284e3f1341c?q=80&w=1200&auto=format&fit=crop',
]

const allTags = [
  '人工智能',
  '虚拟现实',
  '大数据',
  '物联网',
  '量子计算',
  '深度学习',
  '芯片技术',
  '人机交互',
  '脑机接口',
  '金融科技',
  '数字孪生',
  '区块链',
  '生物信息',
  '通信技术',
  '增强现实',
  '自然语言处理',
  '机器人',
  '创新应用',
  '教学改革',
  '智慧城市',
]

const videoSources = [
  'http://vjs.zencdn.net/v/oceans.mp4',
  'http://clips.vorwaerts-gmbh.de/VfE_html5.mp4',
  'https://media.w3.org/2010/05/sintel/trailer.mp4',
]

const galleryPool = [
  'https://images.unsplash.com/photo-1581093450021-4a7360e9a6b5?q=80&w=800&auto=format&fit=crop',
  'https://images.unsplash.com/photo-1581092580497-909141b44917?q=80&w=800&auto=format&fit=crop',
  'https://images.unsplash.com/photo-1581091224003-01e3d649f8a4?q=80&w=800&auto=format&fit=crop',
  'https://images.unsplash.com/photo-1558201005-296b86026a27?q=80&w=800&auto=format&fit=crop',
  'https://images.unsplash.com/photo-1554474013-09418b32060a?q=80&w=800&auto=format&fit=crop',
  'https://images.unsplash.com/photo-1518770660439-4636190af475?q=80&w=800&auto=format&fit=crop',
  'https://images.unsplash.com/photo-1532187643623-8f061554032a?q=80&w=800&auto=format&fit=crop',
  'https://images.unsplash.com/photo-1544724424-b17178a213e4?q=80&w=800&auto=format&fit=crop',
]

const teamMembersPool = [
  {
    name: '周毅',
    role: '项目经理',
    avatar:
      'https://images.unsplash.com/photo-1599566150163-29194dcaad36?q=80&w=150&auto=format&fit=crop',
  },
  {
    name: '林悦',
    role: '前端工程师',
    avatar:
      'https://images.unsplash.com/photo-1521119989659-a83eee488004?q=80&w=150&auto=format&fit=crop',
  },
  {
    name: '徐磊',
    role: '后端架构师',
    avatar:
      'https://images.unsplash.com/photo-1506794778202-cad84cf45f1d?q=80&w=150&auto=format&fit=crop',
  },
  {
    name: '黄敏',
    role: 'UI/UX设计师',
    avatar:
      'https://images.unsplash.com/photo-1554151228-14d9def656e4?q=80&w=150&auto=format&fit=crop',
  },
  {
    name: '郑凯',
    role: '测试工程师',
    avatar:
      'https://images.unsplash.com/photo-1519085360753-af0119f7cbe7?q=80&w=150&auto=format&fit=crop',
  },
  {
    name: '邱琳',
    role: '数据分析师',
    avatar:
      'https://images.unsplash.com/photo-1488426862026-3ee34a7d66df?q=80&w=150&auto=format&fit=crop',
  },
]

// --- 辅助函数 (Helper Functions) ---

// 随机从一个数组中抽取N个不重复的元素
const getRandomItems = (arr, n) => {
  const shuffled = [...arr].sort(() => 0.5 - Math.random())
  return shuffled.slice(0, n)
}

// --- 数据生成逻辑 ---

const generatedData = []

for (let i = 0; i < 30; i++) {
  const authorInfo = authors[i % authors.length] // 循环使用作者信息
  const title = titles[i] // 每条数据使用一个唯一的标题

  generatedData.push({
    id: i + 1,
    title: title,
    description: `本项目（${title}）旨在解决行业内的关键技术难题，通过创新的方法和前沿的技术，我们取得了突破性的进展。该成果不仅在学术上具有重要价值，也展现了巨大的商业应用潜力，吸引了众多行业合作伙伴的关注。`,
    image: images[i % images.length], // 循环使用主图
    tags: getRandomItems(allTags, 3), // 随机抽取3个标签
    author: authorInfo.name,
    authorTitle: authorInfo.title,
    authorAvatar: authorInfo.avatar,
    publishDate: new Date(new Date().setDate(new Date().getDate() - i * 3)), // 发布日期每天递减3天
    contact1: `邮箱 ${authorInfo.name.toLowerCase().replace(' ', '')}@university.edu`, // 生成模拟邮箱
    contact2: `电话 138${String(Math.floor(Math.random() * 100000000)).padStart(8, '0')}`, // 生成模拟电话
    projectUrl: 'https://github.com/example/project-' + (i + 1),

    // 视频和图片库
    videoUrl: videoSources[i % videoSources.length],
    gallery: getRandomItems(galleryPool, 4), // 从图片池随机选4张
    teamMembers: getRandomItems(teamMembersPool, Math.floor(Math.random() * 4) + 2), // 随机选择2-5个团队成员
  })
}

// 最终导出给Vue组件使用的数据
const achievements = ref(generatedData)

const selectedAchievement = ref(null)
const isModalVisible = ref(false)
const isSubmissionModalVisible = ref(false)
const searchQuery = ref('')
const teacherFilter = ref('')
const tagFilter = ref('')
const timeFilter = ref('')
const currentPage = ref(1)
const itemsPerPage = 15

// 动态生成筛选选项
const teacherOptions = computed(() => {
  const teachers = new Set(achievements.value.map((a) => a.author))
  const options = Array.from(teachers).map((teacher) => ({ label: teacher, value: teacher }))
  return [{ label: '按作者筛选', value: '' }, ...options]
})

const tagOptions = computed(() => {
  const tags = new Set(achievements.value.flatMap((a) => a.tags))
  const options = Array.from(tags).map((tag) => ({ label: tag, value: tag }))
  return [{ label: '按标签筛选', value: '' }, ...options]
})

// 搜索筛选函数
const filteredAchievements = computed(() => {
  let filtered = achievements.value

  // 教师筛选
  if (teacherFilter.value) {
    filtered = filtered.filter((achievement) => achievement.author === teacherFilter.value)
  }

  // 标签筛选
  if (tagFilter.value) {
    filtered = filtered.filter((achievement) => achievement.tags.includes(tagFilter.value))
  }

  // 时间筛选
  if (timeFilter.value) {
    const now = new Date()
    const filterDate = new Date()

    switch (timeFilter.value) {
      case 'week':
        filterDate.setDate(now.getDate() - 7)
        break
      case 'month':
        filterDate.setMonth(now.getMonth() - 1)
        break
      case 'three_months':
        filterDate.setMonth(now.getMonth() - 3)
        break
      case 'half_year':
        filterDate.setMonth(now.getMonth() - 6)
        break
      case 'year':
        filterDate.setFullYear(now.getFullYear() - 1)
        break
    }

    filtered = filtered.filter((achievement) => new Date(achievement.publishDate) >= filterDate)
  }

  // 搜索查询
  if (searchQuery.value) {
    const lowerCaseQuery = searchQuery.value.toLowerCase()
    filtered = filtered.filter((achievement) => {
      const titleMatch = achievement.title.toLowerCase().includes(lowerCaseQuery)
      const tagMatch = achievement.tags.some((tag) => tag.toLowerCase().includes(lowerCaseQuery))
      return titleMatch || tagMatch
    })
  }

  return filtered
})

// 分页计算属性
const paginatedAchievements = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage
  const end = start + itemsPerPage
  return filteredAchievements.value.slice(start, end)
})

const totalPages = computed(() => {
  return Math.ceil(filteredAchievements.value.length / itemsPerPage)
})

// --- Methods ---
const scrollToTop = () => {
  nextTick(() => {
    window.scrollTo({ top: 300, behavior: 'smooth' })
  })
}
const goToPage = (page) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
    scrollToTop()
  }
}
const openDetailsModal = (achievement) => {
  selectedAchievement.value = achievement
  isModalVisible.value = true
}

const closeDetailsModal = () => {
  isModalVisible.value = false
  setTimeout(() => {
    selectedAchievement.value = null
  }, 300)
}

const handleSubmission = (newAchievement) => {
  achievements.value.unshift(newAchievement)
  // In a real app, you'd likely send this to a server.
  // You might also want to show a success notification.
}
</script>

<style scoped>
.showcase-page {
  background-color: #f8f9fa;
  min-height: 100vh;
}
.content-container {
  max-width: 1600px;
  margin: 0 auto;
  padding: 0 40px 40px;
}
</style>
