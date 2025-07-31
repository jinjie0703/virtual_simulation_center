import { ref, computed } from 'vue'

export function useTeamData() {
  // 教师数据
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
  ])

  // 标签配置
  const tabsConfig = computed(() => [
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

  // 根据类型获取成员数据
  const getMembersByType = (type) => {
    return type === 'teachers' ? teachers.value : students.value
  }

  return {
    teachers,
    students,
    tabsConfig,
    getMembersByType,
  }
}
