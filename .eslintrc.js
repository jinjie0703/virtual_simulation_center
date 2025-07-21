// .eslintrc.js (使用 ES Module 语法)

// 注意：这里不再是 module.exports =
export default {
  root: true,
  env: {
    browser: true,
    es2021: true,
    node: true,
  },
  extends: ['eslint:recommended', 'plugin:vue/vue3-essential'],
  parserOptions: {
    ecmaVersion: 'latest',
    sourceType: 'module',
  },
  rules: {
    // 将 "unused" 相关的规则从错误(error)降级为警告(warn)
    'no-unused-vars': ['warn', { argsIgnorePattern: '^_' }],
    'vue/no-unused-components': 'warn',

    // 其他自定义规则
    'no-console': 'off',
    'vue/multi-word-component-names': 'off',
  },
}
