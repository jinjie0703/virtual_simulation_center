// eslint.config.js

import { defineConfig, globalIgnores } from 'eslint/config'
import globals from 'globals'
import js from '@eslint/js'
import pluginVue from 'eslint-plugin-vue'
import skipFormatting from '@vue/eslint-config-prettier/skip-formatting'

export default defineConfig([
  {
    name: 'app/files-to-lint',
    files: ['**/*.{js,mjs,jsx,vue}'],
  },

  globalIgnores(['**/dist/**', '**/dist-ssr/**', '**/coverage/**']),

  {
    languageOptions: {
      globals: {
        ...globals.browser,
      },
    },
  },

  js.configs.recommended,
  ...pluginVue.configs['flat/essential'],

  // ===================================================================
  // 将所有自定义规则都放在这个对象里
  // ===================================================================
  {
    rules: {
      // 1. 解决 defineProps/defineEmits 导入问题的规则
      'vue/no-import-compiler-macros': 'error',

      // 2. 您在第一个文件中自己添加的规则
      'no-unused-vars': ['warn', { argsIgnorePattern: '^_' }],
      'vue/no-unused-components': 'warn',
      'no-console': 'off',
      'vue/multi-word-component-names': 'off',
    }
  },

  skipFormatting,
])