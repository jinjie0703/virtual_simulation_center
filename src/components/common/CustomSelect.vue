<template>
  <div
    class="custom-select"
    @click="toggleDropdown"
    ref="selectRef"
    :class="[sizeClass, { 'select-open': isOpen, 'select-disabled': disabled }]"
  >
    <div class="select-trigger" :class="triggerClass">
      <span class="select-value">{{ selectedOption?.label || placeholder }}</span>
      <i :class="['fas', iconClass, { rotated: isOpen }]" :style="{ fontSize: iconSize }"></i>
    </div>
    <transition :name="transitionName">
      <ul v-if="isOpen" class="select-options" :class="optionsClass">
        <li
          v-for="option in options"
          :key="option.value"
          @click.stop="selectOption(option)"
          :class="{ active: option.value === modelValue, disabled: option.disabled }"
          class="select-option"
        >
          <slot name="option" :option="option">
            {{ option.label }}
          </slot>
        </li>
        <li v-if="!options.length" class="select-option select-empty">
          {{ emptyText }}
        </li>
      </ul>
    </transition>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted, watch } from 'vue'

// 定义 props
const props = defineProps({
  modelValue: {
    type: [String, Number],
    default: '',
  },
  options: {
    type: Array,
    required: true,
    default: () => [],
  },
  placeholder: {
    type: String,
    default: '请选择',
  },
  disabled: {
    type: Boolean,
    default: false,
  },
  clearable: {
    type: Boolean,
    default: false,
  },
  size: {
    type: String,
    default: 'medium', // small, medium, large
    validator: (value) => ['small', 'medium', 'large'].includes(value),
  },
  minWidth: {
    type: String,
    default: '140px',
  },
  maxWidth: {
    type: String,
    default: 'auto',
  },
  iconClass: {
    type: String,
    default: 'fa-chevron-down',
  },
  iconSize: {
    type: String,
    default: '0.8rem',
  },
  transitionName: {
    type: String,
    default: 'dropdown',
  },
  emptyText: {
    type: String,
    default: '暂无数据',
  },
  triggerClass: {
    type: String,
    default: '',
  },
  optionsClass: {
    type: String,
    default: '',
  },
})

// 定义 emits
const emit = defineEmits(['update:modelValue', 'change', 'visible-change'])

// 响应式数据
const isOpen = ref(false)
const selectRef = ref(null)

// 计算属性
const selectedOption = computed(() => {
  return props.options.find((option) => option.value === props.modelValue)
})

const sizeClass = computed(() => `select-${props.size}`)

// 方法
const toggleDropdown = () => {
  if (props.disabled) return
  isOpen.value = !isOpen.value
  emit('visible-change', isOpen.value)
}

const selectOption = (option) => {
  if (option.disabled) return
  emit('update:modelValue', option.value)
  emit('change', option.value, option)
  isOpen.value = false
  emit('visible-change', false)
}

const handleClickOutside = (event) => {
  if (selectRef.value && !selectRef.value.contains(event.target)) {
    if (isOpen.value) {
      isOpen.value = false
      emit('visible-change', false)
    }
  }
}

// 生命周期
onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})

// 监听 disabled 状态变化
watch(
  () => props.disabled,
  (newVal) => {
    if (newVal && isOpen.value) {
      isOpen.value = false
      emit('visible-change', false)
    }
  },
)
</script>

<style scoped>
.custom-select {
  position: relative;
  user-select: none;
  min-width: v-bind(minWidth);
  max-width: v-bind(maxWidth);
  z-index: 1; /* Default z-index */
}

.custom-select.select-open {
  z-index: 1001; /* Higher z-index when open */
}

.custom-select.select-disabled {
  opacity: 0.6;
  pointer-events: none;
}

.custom-select.select-open .select-trigger {
  border-color: #3498db;
  box-shadow: 0 0 0 3px rgba(52, 152, 219, 0.15);
}

.select-trigger {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: white;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  cursor: pointer;
  color: #334155;
  transition: all 0.3s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
  gap: 1rem;
}

.select-trigger:hover:not(.disabled) {
  border-color: #3498db;
  box-shadow: 0 2px 8px rgba(52, 152, 219, 0.15);
}

.select-trigger.disabled {
  background-color: #f8fafc;
  color: #94a3b8;
  cursor: not-allowed;
}

/* 尺寸变体 */
.custom-select.select-small .select-trigger {
  padding: 0.6rem 0.8rem;
  font-size: 0.875rem;
}

.custom-select.select-medium .select-trigger {
  padding: 0.9rem 1.2rem;
  font-size: 1rem;
}

.custom-select.select-large .select-trigger {
  padding: 1.1rem 1.4rem;
  font-size: 1.125rem;
}

.select-value {
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.select-trigger i {
  color: #64748b;
  transition: transform 0.3s ease;
  flex-shrink: 0;
}

.select-trigger i.rotated {
  transform: rotate(180deg);
}

.select-options {
  position: absolute;
  top: 100%;
  left: 0;
  right: 0;
  background: white;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
  z-index: 1000; /* Ensure dropdown is above other content */
  margin: 0;
  padding: 0.5rem 0;
  list-style: none;
  margin-top: 4px;
  overflow: hidden;
  max-height: 200px;
  overflow-y: auto;
}

.select-option {
  padding: 0.8rem 1.2rem;
  cursor: pointer;
  transition: all 0.2s ease;
  color: #334155;
  font-size: 1rem;
  border-left: 3px solid transparent;
}

.select-option:hover:not(.select-empty):not(.disabled) {
  background-color: #f8fafc;
  color: #3498db;
  border-left-color: #3498db;
}

.select-option.active {
  background-color: #eff6ff;
  color: #3498db;
  border-left-color: #3498db;
  font-weight: 500;
}

.select-option.disabled {
  color: #94a3b8;
  cursor: not-allowed;
}

.select-option.select-empty {
  color: #94a3b8;
  cursor: default;
  text-align: center;
  font-style: italic;
}

/* 小尺寸选项 */
.custom-select.select-small .select-option {
  padding: 0.6rem 1rem;
  font-size: 0.875rem;
}

/* 大尺寸选项 */
.custom-select.select-large .select-option {
  padding: 1rem 1.4rem;
  font-size: 1.125rem;
}

/* 下拉动画 */
.dropdown-enter-active {
  transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
}

.dropdown-leave-active {
  transition: all 0.2s ease-in;
}

.dropdown-enter-from {
  opacity: 0;
  transform: translateY(-10px) scale(0.95);
}

.dropdown-leave-to {
  opacity: 0;
  transform: translateY(-5px);
}

/* 滑动动画变体 */
.slide-enter-active,
.slide-leave-active {
  transition: all 0.3s ease;
}

.slide-enter-from {
  opacity: 0;
  transform: translateY(-10px);
}

.slide-leave-to {
  opacity: 0;
  transform: translateY(-5px);
}

/* 淡入动画变体 */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
