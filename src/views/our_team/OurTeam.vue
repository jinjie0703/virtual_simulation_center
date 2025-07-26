<!-- filepath: /C:/Users/25948/Desktop/virtual_simulation_center/src/views/OurTeam.vue -->
<template>
  <header class="team-header">
    <div>
      <h1 class="title">师资队伍</h1>
      <p class="subtitle">汇聚顶尖人才，打造创新团队</p>
    </div>
  </header>
  <div class="our-team-container">
    <!-- 标签切换 -->
    <TeamTabs
      :tabs="tabsData"
      :activeTab="activeTab"
      :isLoading="isLoading"
      @switchTab="switchTab"
    />

    <!-- 成员展示区域 -->
    <div class="members-section">
      <div v-if="isLoading" class="loading-state">
        <div class="loading-spinner"></div>
        <p>加载中...</p>
      </div>

      <div v-else-if="paginatedMembers.length > 0" class="members-content">
        <TransitionGroup name="fade" tag="div" class="members-grid">
          <MemberCard
            v-for="member in paginatedMembers"
            :key="`${activeTab}-${member.id}`"
            :member="member"
          />
        </TransitionGroup>

        <!-- 分页组件 -->
        <TeamPagination
          v-if="totalPages > 1"
          :currentPage="currentPage"
          :totalPages="totalPages"
          @changePage="changePage"
        />
      </div>

      <div v-else class="empty-state">
        <div class="empty-icon">👥</div>
        <h3>暂无成员信息</h3>
        <p>当前分类下还没有添加成员</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, nextTick } from 'vue'
import MemberCard from '@/components/our_team/MemberCard.vue'
import TeamPagination from '@/components/our_team/TeamPagination.vue'
import TeamTabs from '@/components/our_team/TeamTabs.vue'

// 响应式数据
const isLoading = ref(false)
const activeTab = ref('teachers')
const currentPage = ref(1)
const itemsPerPage = 10 // 改为10个：一行5个，两行

// 示例数据 - 添加更多数据以便测试分页
const teachers = ref([
  {
    id: 'teacher1',
    name: '张教授',
    title: '虚拟仿真技术专家',
    description: '致力于虚拟现实技术在教育领域的应用研究，拥有丰富的项目经验和学术成果。',
    avatar:
      'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=200&h=200&fit=crop&crop=face',
    email: 'zhang.prof@university.edu',
    office: '教三楼 301',
    research: '虚拟现实技术、教育信息化、人机交互',
    achievements: '发表SCI论文15篇，主持国家级项目3项',
  },
  {
    id: 'teacher2',
    name: '李博士',
    title: 'AI算法研究员',
    description: '专注于人工智能算法的研究与开发，在机器学习和深度学习领域有深入研究。',
    avatar:
      'https://images.unsplash.com/photo-1494790108755-2616b812c97c?w=200&h=200&fit=crop&crop=face',
    email: 'li.dr@university.edu',
    office: '教三楼 302',
    research: '机器学习、深度学习、计算机视觉',
    achievements: '发表顶级会议论文20篇，获得专利5项',
  },
  {
    id: 'teacher3',
    name: '王研究员',
    title: '数据科学专家',
    description: '在大数据分析和数据挖掘方面有着丰富的经验，致力于教育数据的智能分析。',
    avatar:
      'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?w=200&h=200&fit=crop&crop=face',
    email: 'wang.researcher@university.edu',
    office: '教三楼 303',
    research: '大数据分析、数据挖掘、教育数据科学',
    achievements: '主导多个教育大数据项目，获得省部级奖励2项',
  },
  {
    id: 'teacher4',
    name: '陈副教授',
    title: '软件工程专家',
    description: '专业从事软件工程和系统架构设计，在大型教育系统开发方面经验丰富。',
    avatar:
      'https://images.unsplash.com/photo-1438761681033-6461ffad8d80?w=200&h=200&fit=crop&crop=face',
    email: 'chen.assocprof@university.edu',
    office: '教三楼 304',
    research: '软件工程、系统架构、分布式系统',
    achievements: '参与国家重点研发计划2项，软件著作权10项',
  },
  {
    id: 'teacher5',
    name: '刘讲师',
    title: '交互设计专家',
    description: '专注于用户体验设计和人机交互研究，致力于提升教育软件的用户体验。',
    avatar:
      'https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=200&h=200&fit=crop&crop=face',
    email: 'liu.lecturer@university.edu',
    office: '教三楼 305',
    research: '用户体验设计、人机交互、界面设计',
    achievements: '设计作品获得国际设计奖项3项，发表设计类论文8篇',
  },
  {
    id: 'teacher6',
    name: '赵助教',
    title: '前端开发工程师',
    description: '负责虚拟仿真平台的前端开发和维护，在现代Web技术方面有深入研究。',
    avatar:
      'https://images.unsplash.com/photo-1507591064344-4c6ce005b128?w=200&h=200&fit=crop&crop=face',
    email: 'zhao.assistant@university.edu',
    office: '教三楼 306',
    research: '前端开发、Web技术、用户界面实现',
    achievements: '开发多个教育平台前端系统，技术博客阅读量50万+',
  },
  {
    id: 'teacher7',
    name: '黄教授',
    title: '系统架构师',
    description: '在分布式系统和云计算架构方面有深入研究，负责平台整体技术架构设计。',
    avatar:
      'https://images.unsplash.com/photo-1560250097-0b93528c311a?w=200&h=200&fit=crop&crop=face',
    email: 'huang.prof@university.edu',
    office: '教三楼 307',
    research: '分布式系统、云计算、微服务架构',
    achievements: '主导大型系统架构设计10余项，获得技术专利8项',
  },
  {
    id: 'teacher8',
    name: '马副教授',
    title: '网络安全专家',
    description: '专注于网络安全和信息安全研究，在教育平台安全防护方面有丰富经验。',
    avatar:
      'https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?w=200&h=200&fit=crop&crop=face',
    email: 'ma.assocprof@university.edu',
    office: '教三楼 308',
    research: '网络安全、信息安全、密码学',
    achievements: '发表安全领域论文20篇，获得安全专利6项',
  },
  {
    id: 'teacher9',
    name: '杨讲师',
    title: '移动开发专家',
    description: '专业从事移动应用开发，在跨平台移动开发和用户界面设计方面有深入研究。',
    avatar:
      'https://images.unsplash.com/photo-1438761681033-6461ffad8d80?w=200&h=200&fit=crop&crop=face',
    email: 'yang.lecturer@university.edu',
    office: '教三楼 309',
    research: '移动开发、跨平台技术、UI/UX设计',
    achievements: '开发移动应用15款，用户下载量超过100万',
  },
  {
    id: 'teacher10',
    name: '周助教',
    title: '数据库专家',
    description: '专注于数据库设计和优化，在大规模数据存储和查询优化方面有丰富经验。',
    avatar:
      'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=200&h=200&fit=crop&crop=face',
    email: 'zhou.assistant@university.edu',
    office: '教三楼 310',
    research: '数据库设计、查询优化、数据仓库',
    achievements: '优化数据库性能项目12项，发表数据库论文8篇',
  },
  {
    id: 'teacher11',
    name: '林教授',
    title: '云计算专家',
    description: '在云计算和分布式计算领域有深入研究，负责教育云平台的架构设计。',
    avatar:
      'https://images.unsplash.com/photo-1494790108755-2616b812c97c?w=200&h=200&fit=crop&crop=face',
    email: 'lin.prof@university.edu',
    office: '教三楼 311',
    research: '云计算、分布式计算、容器技术',
    achievements: '主导云平台建设5项，获得云计算认证3项',
  },
  {
    id: 'teacher12',
    name: '郭博士',
    title: '区块链研究员',
    description: '专注于区块链技术在教育领域的应用研究，在去中心化系统设计方面有创新贡献。',
    avatar:
      'https://images.unsplash.com/photo-1500648767791-00dcc994a43e?w=200&h=200&fit=crop&crop=face',
    email: 'guo.dr@university.edu',
    office: '教三楼 312',
    research: '区块链技术、去中心化系统、智能合约',
    achievements: '发表区块链论文10篇，参与区块链项目开发3项',
  },
])

const students = ref([
  {
    id: 'student1',
    name: '孙同学',
    title: '研究生助理',
    description: '计算机科学专业研究生，参与多个虚拟仿真项目的开发工作。',
    avatar:
      'https://images.unsplash.com/photo-1534528741775-53994a69daeb?w=200&h=200&fit=crop&crop=face',
    email: 'sun.student@university.edu',
    office: '实验室 A201',
    research: '虚拟现实应用开发、移动应用开发',
    achievements: '参与项目开发5项，获得校级奖学金2次',
  },
  {
    id: 'student2',
    name: '周同学',
    title: '本科生助理',
    description: '软件工程专业优秀本科生，在Web开发和数据库设计方面表现突出。',
    avatar:
      'https://images.unsplash.com/photo-1519345182560-3f2917c472ef?w=200&h=200&fit=crop&crop=face',
    email: 'zhou.student@university.edu',
    office: '实验室 A202',
    research: 'Web开发、数据库设计、系统测试',
    achievements: '参与系统开发3项，发表学术论文1篇',
  },
  {
    id: 'student3',
    name: '吴同学',
    title: '博士研究生',
    description: '人工智能专业博士研究生，专注于教育AI算法的研究与应用。',
    avatar:
      'https://images.unsplash.com/photo-1506794778202-cad84cf45f1d?w=200&h=200&fit=crop&crop=face',
    email: 'wu.phd@university.edu',
    office: '实验室 A203',
    research: '教育人工智能、自然语言处理、知识图谱',
    achievements: '发表SCI论文3篇，参与国家项目2项',
  },
  {
    id: 'student4',
    name: '郑同学',
    title: '硕士研究生',
    description: '数据科学专业硕士研究生，在教育数据挖掘领域有重要贡献。',
    avatar:
      'https://images.unsplash.com/photo-1517841905240-472988babdf9?w=200&h=200&fit=crop&crop=face',
    email: 'zheng.master@university.edu',
    office: '实验室 A204',
    research: '教育数据挖掘、学习分析、数据可视化',
    achievements: '完成数据分析项目8项，获得优秀论文奖1次',
  },
  {
    id: 'student5',
    name: '林同学',
    title: '硕士研究生',
    description: '计算机视觉专业硕士研究生，在图像处理和模式识别方面有突出表现。',
    avatar:
      'https://images.unsplash.com/photo-1531746020798-e6953c6e8e04?w=200&h=200&fit=crop&crop=face',
    email: 'lin.master@university.edu',
    office: '实验室 A205',
    research: '计算机视觉、图像处理、模式识别',
    achievements: '发表会议论文5篇，参与视觉项目开发3项',
  },
  {
    id: 'student6',
    name: '何同学',
    title: '本科生助理',
    description: '计算机科学专业本科生，在算法设计和数据结构方面有优秀表现。',
    avatar:
      'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=200&h=200&fit=crop&crop=face',
    email: 'he.student@university.edu',
    office: '实验室 A206',
    research: '算法设计、数据结构、编程竞赛',
    achievements: '获得编程竞赛奖项3次，参与开源项目2项',
  },
  {
    id: 'student7',
    name: '陈同学',
    title: '研究生助理',
    description: '软件工程专业研究生，专注于软件测试和质量保证方面的研究。',
    avatar:
      'https://images.unsplash.com/photo-1494790108755-2616b812c97c?w=200&h=200&fit=crop&crop=face',
    email: 'chen.graduate@university.edu',
    office: '实验室 A207',
    research: '软件测试、质量保证、自动化测试',
    achievements: '完成测试项目6项，发表测试相关论文2篇',
  },
  {
    id: 'student8',
    name: '李同学',
    title: '博士研究生',
    description: '机器学习专业博士研究生，在深度学习和神经网络方面有深入研究。',
    avatar:
      'https://images.unsplash.com/photo-1438761681033-6461ffad8d80?w=200&h=200&fit=crop&crop=face',
    email: 'li.phd@university.edu',
    office: '实验室 A208',
    research: '深度学习、神经网络、机器学习',
    achievements: '发表顶级会议论文4篇，获得学术奖学金1次',
  },
])

// 计算属性
const tabsData = computed(() => [
  {
    id: 'teachers',
    name: '教师团队',
    icon: '👨‍🏫',
    count: teachers.value.length,
  },
  {
    id: 'students',
    name: '学生团队',
    icon: '👨‍🎓',
    count: students.value.length,
  },
])

const currentMembers = computed(() => {
  return activeTab.value === 'teachers' ? teachers.value : students.value
})

const totalPages = computed(() => {
  return Math.ceil(currentMembers.value.length / itemsPerPage)
})

const paginatedMembers = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage
  const end = start + itemsPerPage
  return currentMembers.value.slice(start, end)
})

// 方法
const switchTab = async (tab) => {
  if (tab === activeTab.value) return

  isLoading.value = true
  activeTab.value = tab
  currentPage.value = 1

  // 模拟加载延迟
  await nextTick()
  setTimeout(() => {
    isLoading.value = false
  }, 300)
}

const changePage = (page) => {
  if (page >= 1 && page <= totalPages.value && page !== currentPage.value) {
    currentPage.value = page
    // 滚动到顶部
    document.querySelector('.members-section')?.scrollIntoView({
      behavior: 'smooth',
      block: 'start',
    })
  }
}

// 监听切换标签时重置页码
watch(activeTab, () => {
  currentPage.value = 1
})
</script>

<style scoped>
.team-header {
  padding: 120px 0px 80px 0px;
  text-align: center;
  position: relative;
  overflow: hidden;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
}

.title {
  font-size: 3.5rem;
  font-weight: 800;
  margin: 0 0 15px;
  animation: fadeInDown 1s ease-out;
  color: #2c3e50;
  background: linear-gradient(135deg, #2c3e50, #4a90e2);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.subtitle {
  font-size: 1.25rem;
  max-width: 800px;
  margin: 0 auto;
  opacity: 0.9;
  line-height: 1.6;
  animation: fadeInUp 1s ease-out 0.3s both;
  color: #666;
}

@keyframes fadeInDown {
  from {
    opacity: 0;
    transform: translateY(-30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
.our-team-container {
  /* 移除 max-width 或增大数值 */
  /* max-width: 1600px; */
  margin: 0 auto;
  padding: 60px 40px;
  font-family: 'Helvetica Neue', Arial, sans-serif;
  background-color: #f8f9fa;
  min-height: 100vh;
  width: 100%; /* 添加这行确保全宽 */
}

.members-section {
  min-height: 800px;
  position: relative;
  padding: 0 8%; /* 默认左右留白8% */
}

/* 成员展示区域 */
.members-grid {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  grid-template-rows: repeat(2, 1fr);
  gap: 25px;
  margin-bottom: 50px;
  /* 移除 max-width 限制 */
  /* max-width: 1400px; */
  width: 100%; /* 使用全宽 */
  margin-left: auto;
  margin-right: auto;
  min-height: 500px;
}

.loading-spinner {
  width: 40px;
  height: 40px;
  border: 3px solid #f3f3f3;
  border-top: 3px solid #4a90e2;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 20px;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

/* 空状态 */
.empty-state {
  text-align: center;
  padding: 80px 20px;
  color: #666;
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

/* 过渡动画 */
.fade-enter-active,
.fade-leave-active {
  transition: all 0.6s cubic-bezier(0.4, 0, 0.2, 1);
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(30px) scale(0.95);
}

.fade-move {
  transition: all 0.6s cubic-bezier(0.4, 0, 0.2, 1);
}

/* 响应式设计 */
@media (max-width: 1400px) {
  .members-grid {
    grid-template-columns: repeat(4, 1fr); /* 4列 */
    max-width: 1200px;
  }

  .our-team-container {
    max-width: 1400px;
  }
}

@media (max-width: 1100px) {
  .members-grid {
    grid-template-columns: repeat(3, 1fr); /* 3列 */
    max-width: 900px;
  }
}

@media (max-width: 800px) {
  .members-grid {
    grid-template-columns: repeat(2, 1fr); /* 2列 */
    max-width: 600px;
    gap: 20px;
  }

  .our-team-container {
    padding: 40px 20px;
  }

  .main-title {
    font-size: 2.8rem;
  }
}

@media (max-width: 600px) {
  .members-grid {
    grid-template-columns: 1fr; /* 1列 */
    max-width: 400px;
    gap: 20px;
    grid-template-rows: auto; /* 自动行高 */
  }

  .main-title {
    font-size: 2.4rem;
  }

  .description {
    font-size: 1.1rem;
  }

  .our-team-container {
    padding: 30px 15px;
  }
}
@media (max-width: 1400px) {
  .members-section {
    padding: 0 6%; /* 中等屏幕留白6% */
  }
}

@media (max-width: 1100px) {
  .members-section {
    padding: 0 4%; /* 较小屏幕留白4% */
  }
}

@media (max-width: 800px) {
  .members-section {
    padding: 0 3%; /* 平板留白3% */
  }

  .members-grid {
    grid-template-columns: repeat(3, 1fr); /* 3列布局 */
    gap: 25px;
  }
}

@media (max-width: 600px) {
  .members-section {
    padding: 0 2%; /* 手机留白2% */
  }

  .members-grid {
    grid-template-columns: repeat(2, 1fr); /* 2列布局 */
    gap: 20px;
  }
}
</style>
