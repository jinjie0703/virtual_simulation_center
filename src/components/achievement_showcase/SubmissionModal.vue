<template>
  <transition name="modal-fade">
    <div v-if="isVisible" class="modal-backdrop" @click.self="$emit('close')">
      <div class="modal-container">
        <button class="close-btn" @click="$emit('close')">&times;</button>
        <div class="modal-header">
          <h2 class="modal-title">提交新项目成果</h2>
          <p class="modal-subtitle">请填写以下信息，分享您的创新成果</p>
        </div>
        <form class="submission-form" @submit.prevent="handleSubmit">
          <div class="form-grid">
            <!-- Text Inputs -->
            <div class="form-group">
              <label for="title">项目标题</label>
              <input id="title" v-model="formData.title" type="text" required />
            </div>
            <div class="form-group">
              <label for="tags">技术标签 (逗号分隔)</label>
              <input id="tags" v-model="formData.tags" type="text" required />
            </div>
            <div class="form-group full-width">
              <label for="description">项目描述</label>
              <textarea
                id="description"
                v-model="formData.description"
                rows="3"
                required
              ></textarea>
            </div>
            <div class="form-group">
              <label for="author">作者姓名</label>
              <input id="author" v-model="formData.author" type="text" required />
            </div>
            <div class="form-group">
              <label for="authorTitle">作者职称</label>
              <input id="authorTitle" v-model="formData.authorTitle" type="text" required />
            </div>
            <div class="form-group">
              <label for="contact1">联系方式 1 (选填)</label>
              <input id="contact1" v-model="formData.contact1" type="text" />
            </div>
            <div class="form-group">
              <label for="contact2">联系方式 2 (选填)</label>
              <input id="contact2" v-model="formData.contact2" type="text" />
            </div>
            <div class="form-group">
              <label for="projectUrl">项目地址URL (选填)</label>
              <input id="projectUrl" v-model="formData.projectUrl" type="url" />
            </div>

            <!-- File Inputs -->
            <div class="form-group file-group">
              <label>作者头像</label>
              <div class="file-input-wrapper">
                <button
                  type="button"
                  @click="triggerFileInput('authorAvatarInput')"
                  class="btn-upload"
                >
                  选择图片
                </button>
                <span class="file-name">{{ fileNames.authorAvatar || '未选择文件' }}</span>
                <input
                  ref="authorAvatarInput"
                  type="file"
                  @change="handleFileChange($event, 'authorAvatar')"
                  accept="image/*"
                  hidden
                />
              </div>
              <img
                v-if="previews.authorAvatar"
                :src="previews.authorAvatar"
                class="preview-avatar"
              />
            </div>

            <div class="form-group file-group">
              <label>封面图片</label>
              <div class="file-input-wrapper">
                <button type="button" @click="triggerFileInput('imageInput')" class="btn-upload">
                  选择图片
                </button>
                <span class="file-name">{{ fileNames.image || '未选择文件' }}</span>
                <input
                  ref="imageInput"
                  type="file"
                  @change="handleFileChange($event, 'image')"
                  accept="image/*"
                  hidden
                  required
                />
              </div>
              <img v-if="previews.image" :src="previews.image" class="preview-image" />
            </div>

            <div class="form-group file-group">
              <label>封面视频</label>
              <div class="file-input-wrapper">
                <button type="button" @click="triggerFileInput('videoInput')" class="btn-upload">
                  选择视频
                </button>
                <span class="file-name">{{ fileNames.videoUrl || '未选择文件' }}</span>
                <input
                  ref="videoInput"
                  type="file"
                  @change="handleFileChange($event, 'videoUrl')"
                  accept="video/*"
                  hidden
                  required
                />
              </div>
            </div>

            <div class="form-group file-group full-width">
              <label>项目掠影 (可多选)</label>
              <div class="file-input-wrapper">
                <button type="button" @click="triggerFileInput('galleryInput')" class="btn-upload">
                  选择图片
                </button>
                <span class="file-name">{{
                  fileNames.gallery ? `${fileNames.gallery.length} 个文件已选择` : '未选择文件'
                }}</span>
                <input
                  ref="galleryInput"
                  type="file"
                  @change="handleFileChange($event, 'gallery')"
                  accept="image/*"
                  multiple
                  hidden
                  required
                />
              </div>
              <div v-if="previews.gallery && previews.gallery.length" class="gallery-preview">
                <img v-for="src in previews.gallery" :key="src" :src="src" class="preview-thumb" />
              </div>
            </div>
          </div>
          <div class="form-actions">
            <button type="button" class="btn-cancel" @click="$emit('close')">取消</button>
            <button type="submit" class="btn-submit">提交审核</button>
          </div>
        </form>
      </div>
    </div>
  </transition>
</template>

<script setup>
import { ref, watch } from 'vue'

const props = defineProps({
  isVisible: Boolean,
})
const emit = defineEmits(['close', 'submit'])

// Input Refs
const authorAvatarInput = ref(null)
const imageInput = ref(null)
const videoInput = ref(null)
const galleryInput = ref(null)

const initialFormData = {
  title: '',
  description: '',
  tags: '',
  author: '',
  authorTitle: '',
  contact1: '',
  contact2: '',
  projectUrl: '',
  authorAvatar: null,
  image: null,
  videoUrl: null,
  gallery: [],
}

const formData = ref({ ...initialFormData })
const previews = ref({ authorAvatar: '', image: '', gallery: [] })
const fileNames = ref({ authorAvatar: '', image: '', videoUrl: '', gallery: null })

watch(
  () => props.isVisible,
  (newVal) => {
    if (newVal) {
      formData.value = { ...initialFormData }
      previews.value = { authorAvatar: '', image: '', gallery: [] }
      fileNames.value = { authorAvatar: '', image: '', videoUrl: '', gallery: null }
    }
  },
)

const triggerFileInput = (refName) => {
  const inputRefs = { authorAvatarInput, imageInput, videoInput, galleryInput }
  inputRefs[refName]?.value.click()
}

const handleFileChange = (event, field) => {
  const files = event.target.files
  if (!files || files.length === 0) return

  if (field === 'gallery') {
    formData.value.gallery = Array.from(files)
    fileNames.value.gallery = Array.from(files).map((f) => f.name)
    previews.value.gallery = Array.from(files).map((f) => URL.createObjectURL(f))
  } else {
    const file = files[0]
    formData.value[field] = file
    fileNames.value[field] = file.name
    if (field !== 'videoUrl') {
      previews.value[field] = URL.createObjectURL(file)
    }
  }
}

const handleSubmit = () => {
  const finalData = {
    ...formData.value,
    id: Date.now(),
    tags: formData.value.tags.split(',').map((tag) => tag.trim()),
    // Create object URLs for the files to be used in the parent component
    authorAvatar: formData.value.authorAvatar
      ? URL.createObjectURL(formData.value.authorAvatar)
      : '',
    image: formData.value.image ? URL.createObjectURL(formData.value.image) : '',
    videoUrl: formData.value.videoUrl ? URL.createObjectURL(formData.value.videoUrl) : '',
    gallery: formData.value.gallery.map((f) => URL.createObjectURL(f)),
  }
  emit('submit', finalData)
  emit('close')
}
</script>

<style scoped>
/* Modal base styles - same as before */
.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: opacity 0.3s ease;
}
.modal-fade-enter-from,
.modal-fade-leave-to {
  opacity: 0;
}
.modal-backdrop {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.6);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}
.modal-container {
  background: #fff;
  border-radius: 16px;
  width: 90%;
  max-width: 800px;
  max-height: 90vh;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
  position: relative;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}
.close-btn {
  position: absolute;
  top: 40px;
  right: 40px;
  width: 30px;
  height: 30px;
  border: none;
  background: #f1f5f9;
  color: #64748b;
  border-radius: 50%;
  font-size: 20px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
}
.close-btn:hover {
  background: #e2e8f0;
  transform: rotate(90deg);
}
.modal-header {
  padding: 30px 40px;
  border-bottom: 1px solid #e2e8f0;
  flex-shrink: 0;
}
.modal-title {
  font-size: 1.8rem;
  font-weight: 700;
  color: #1e293b;
  margin: 0 0 5px 0;
}
.modal-subtitle {
  font-size: 1rem;
  color: #64748b;
  margin: 0;
}
.submission-form {
  padding: 30px 40px;
  overflow-y: auto;
  flex-grow: 1;
}
.form-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 25px;
}
.form-group {
  display: flex;
  flex-direction: column;
}
.form-group.full-width {
  grid-column: 1 / -1;
}
.form-group label {
  margin-bottom: 8px;
  font-weight: 600;
  font-size: 0.9rem;
  color: #334155;
}
.form-group input[type='text'],
.form-group input[type='url'],
.form-group textarea {
  width: 100%;
  padding: 10px 15px;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  font-size: 1rem;
  transition: all 0.3s ease;
  outline: none;
}
.form-group input:focus,
.form-group textarea:focus {
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.2);
}
textarea {
  resize: vertical;
}

/* File Input Styles */
.file-input-wrapper {
  display: flex;
  align-items: center;
  gap: 10px;
}
.btn-upload {
  padding: 10px 15px;
  border: 1px solid #667eea;
  border-radius: 8px;
  background-color: #f0f2ff;
  color: #667eea;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
  white-space: nowrap;
}
.btn-upload:hover {
  background-color: #667eea;
  color: white;
}
.file-name {
  font-size: 0.9rem;
  color: #64748b;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.preview-avatar {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  object-fit: cover;
  margin-top: 10px;
  border: 2px solid #e2e8f0;
}
.preview-image {
  width: 100%;
  max-height: 150px;
  border-radius: 8px;
  object-fit: cover;
  margin-top: 10px;
  border: 2px solid #e2e8f0;
}
.gallery-preview {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  margin-top: 10px;
}
.preview-thumb {
  width: 70px;
  height: 70px;
  border-radius: 8px;
  object-fit: cover;
  border: 2px solid #e2e8f0;
}

/* Form Actions */
.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 15px;
  padding: 30px 40px;
  border-top: 1px solid #e2e8f0;
  flex-shrink: 0;
}
.btn-cancel,
.btn-submit {
  padding: 12px 25px;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
}
.btn-cancel {
  background-color: #f1f5f9;
  color: #475569;
}
.btn-cancel:hover {
  background-color: #e2e8f0;
}
.btn-submit {
  background-color: #667eea;
  color: white;
  box-shadow: 0 4px 6px rgba(102, 126, 234, 0.2);
}
.btn-submit:hover {
  background-color: #5a67d8;
  transform: translateY(-2px);
  box-shadow: 0 7px 14px rgba(102, 126, 234, 0.3);
}
</style>
