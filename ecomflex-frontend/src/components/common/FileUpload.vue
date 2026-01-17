<template>
  <div 
    class="w-full"
    :class="{ 'opacity-60 pointer-events-none': disabled }"
  >
    <!-- Label -->
    <label v-if="label" class="form-label mb-2 block">
      {{ label }}
      <span v-if="required" class="text-red-600">*</span>
    </label>
    
    <!-- File Input Area -->
    <div
      ref="dropZone"
      class="relative border-2 border-dashed rounded-xl p-6 flex flex-col items-center justify-center transition-colors"
      :class="[
        isDragging ? 'border-brand-500 bg-brand-50' : 'border-gray-300 hover:border-brand-400',
        { 'border-red-500': error }
      ]"
      @dragenter.prevent="isDragging = true"
      @dragleave.prevent="isDragging = false"
      @dragover.prevent="isDragging = true"
      @drop.prevent="onDrop"
      @click="openFileDialog"
    >
      <!-- Default Upload State -->
      <div v-if="!hasFiles" class="text-center">
        <!-- Upload Icon -->
        <svg
          class="mx-auto h-12 w-12 text-gray-400"
          stroke="currentColor"
          fill="none"
          viewBox="0 0 48 48"
          aria-hidden="true"
        >
          <path
            d="M28 8H12a4 4 0 00-4 4v20m32-12v8m0 0v8a4 4 0 01-4 4H8m36-12h-4m4 0H20"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
        
        <!-- Instructions -->
        <div class="mt-2 flex flex-col sm:flex-row items-center justify-center text-sm text-gray-600">
          <label
            for="file-upload"
            class="relative cursor-pointer rounded-md font-medium text-brand-600 hover:text-brand-500 focus-within:outline-none"
          >
            <span>{{ buttonText }}</span>
            <input
              :id="id"
              ref="fileInput"
              type="file"
              name="file-upload"
              class="sr-only"
              :accept="accept"
              :multiple="multiple"
              @change="onFileChange"
              :disabled="disabled"
            />
          </label>
          <p class="pl-1">or drag and drop</p>
        </div>
        
        <!-- Accepted File Types -->
        <p class="text-xs text-gray-500 mt-1">
          {{ acceptedFileTypesLabel }}
        </p>
      </div>
      
      <!-- File Preview State -->
      <div v-else class="w-full">
        <ul class="divide-y divide-gray-200">
          <li v-for="(file, index) in fileList" :key="index" class="py-3 flex items-center justify-between">
            <div class="flex items-center">
              <!-- File Icon based on type -->
              <div class="flex-shrink-0 h-10 w-10 rounded-lg bg-gray-100 flex items-center justify-center">
                <svg v-if="isImage(file)" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                </svg>
                <svg v-else-if="isPDF(file)" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                </svg>
                <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
                </svg>
              </div>
              
              <!-- File Name and Size -->
              <div class="ml-4 flex-1 truncate">
                <div class="text-sm font-medium text-gray-900 truncate">{{ file.name }}</div>
                <div class="text-xs text-gray-500">{{ formatFileSize(file.size) }}</div>
              </div>
            </div>
            
            <!-- Remove Button -->
            <button
              @click.stop="removeFile(index)"
              type="button"
              class="ml-4 flex-shrink-0 p-1 rounded-full text-gray-400 hover:text-gray-500 hover:bg-gray-100"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
              </svg>
            </button>
          </li>
        </ul>
        
        <!-- Add More Button (for multiple files) -->
        <div v-if="multiple" class="mt-4 flex justify-center">
          <button
            @click.stop="openFileDialog"
            type="button"
            class="inline-flex items-center px-3 py-1.5 border border-gray-300 shadow-sm text-xs font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="-ml-0.5 mr-1 h-4 w-4" viewBox="0 0 20 20" fill="currentColor">
              <path fill-rule="evenodd" d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z" clip-rule="evenodd" />
            </svg>
            Add more files
          </button>
        </div>
      </div>
    </div>
    
    <!-- Preview for Image Files -->
    <div v-if="showPreview && hasImageFiles" class="mt-4">
      <div class="grid grid-cols-2 md:grid-cols-3 gap-2">
        <div 
          v-for="(file, index) in imageFiles" 
          :key="index" 
          class="relative group rounded-lg overflow-hidden shadow-sm border border-gray-200"
        >
          <img :src="getImageUrl(file)" alt="Preview" class="w-full h-24 object-cover" />
          <div class="absolute inset-0 bg-black bg-opacity-40 opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-center">
            <button
              @click.stop="removeFile(getFileIndex(file))"
              type="button"
              class="p-1 rounded-full text-white hover:bg-white hover:bg-opacity-20"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Error Message -->
    <div v-if="error" class="mt-1 text-sm text-red-600">
      {{ error }}
    </div>
    
    <!-- Helper Text -->
    <div v-else-if="helperText" class="mt-1 text-sm text-gray-500">
      {{ helperText }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onUnmounted } from 'vue'

interface Props {
  modelValue: File | File[] | null
  label?: string
  accept?: string
  multiple?: boolean
  required?: boolean
  disabled?: boolean
  showPreview?: boolean
  id?: string
  buttonText?: string
  error?: string
  helperText?: string
  maxFileSize?: number // in bytes
}

const props = withDefaults(defineProps<Props>(), {
  multiple: false,
  required: false,
  disabled: false,
  showPreview: true,
  id: 'file-upload',
  buttonText: 'Upload a file',
  accept: '.jpg,.jpeg,.png,.pdf,.doc,.docx',
  maxFileSize: 5 * 1024 * 1024 // 5MB default
})

const emit = defineEmits<{
  (e: 'update:modelValue', value: File | File[] | null): void
  (e: 'file-added', file: File): void
  (e: 'file-removed', file: File): void
  (e: 'error', message: string): void
}>()

// Refs
const fileInput = ref<HTMLInputElement | null>(null)
const dropZone = ref<HTMLDivElement | null>(null)
const isDragging = ref(false)
const imageUrlMap = ref<Map<File, string>>(new Map())

// Computed Properties
const fileList = computed<File[]>(() => {
  if (!props.modelValue) return []
  return Array.isArray(props.modelValue) ? props.modelValue : [props.modelValue]
})

const hasFiles = computed(() => {
  return !!props.modelValue && (Array.isArray(props.modelValue) ? props.modelValue.length > 0 : true)
})

const acceptedFileTypesLabel = computed(() => {
  if (props.accept === '*') return 'All file types accepted'
  
  const types = props.accept.split(',').map(type => {
    return type.trim().replace('.', '').toUpperCase()
  }).join(', ')
  
  return `Accepted file types: ${types}`
})

const imageFiles = computed(() => {
  return fileList.value.filter(file => isImage(file))
})

const hasImageFiles = computed(() => imageFiles.value.length > 0)

// Methods
const openFileDialog = () => {
  if (!props.disabled && fileInput.value) {
    fileInput.value.click()
  }
}

const onFileChange = (event: Event) => {
  const input = event.target as HTMLInputElement
  if (!input.files || input.files.length === 0) return
  
  processFiles(Array.from(input.files))
  
  // Clear the input to allow selecting the same file again
  if (fileInput.value) fileInput.value.value = ''
}

const onDrop = (event: DragEvent) => {
  isDragging.value = false
  
  if (!event.dataTransfer?.files || event.dataTransfer.files.length === 0) return
  
  processFiles(Array.from(event.dataTransfer.files))
}

const processFiles = (files: File[]) => {
  const validFiles = files.filter(file => {
    // Check file type
    if (props.accept !== '*' && !isFileTypeAccepted(file)) {
      emit('error', `File type not accepted: ${file.type}`)
      return false
    }
    
    // Check file size
    if (file.size > props.maxFileSize) {
      emit('error', `File too large: ${file.name} (max: ${formatFileSize(props.maxFileSize)})`)
      return false
    }
    
    return true
  })
  
  if (validFiles.length === 0) return
  
  if (props.multiple) {
    const currentFiles = Array.isArray(props.modelValue) ? [...props.modelValue] : []
    const newFiles = [...currentFiles, ...validFiles]
    emit('update:modelValue', newFiles)
    
    // Emit file-added for each valid file
    validFiles.forEach(file => emit('file-added', file))
  } else {
    // For single file upload, just take the first valid file
    emit('update:modelValue', validFiles[0])
    emit('file-added', validFiles[0])
  }
}

const removeFile = (index: number) => {
  if (props.multiple && Array.isArray(props.modelValue)) {
    const files = [...props.modelValue]
    const removedFile = files[index]
    
    // Clean up object URL if it's an image
    if (isImage(removedFile)) {
      const url = imageUrlMap.value.get(removedFile)
      if (url) {
        URL.revokeObjectURL(url)
        imageUrlMap.value.delete(removedFile)
      }
    }
    
    files.splice(index, 1)
    emit('update:modelValue', files.length ? files : null)
    emit('file-removed', removedFile)
  } else if (!Array.isArray(props.modelValue) && props.modelValue) {
    // Clean up object URL if it's an image
    if (isImage(props.modelValue)) {
      const url = imageUrlMap.value.get(props.modelValue)
      if (url) {
        URL.revokeObjectURL(url)
        imageUrlMap.value.delete(props.modelValue)
      }
    }
    
    emit('update:modelValue', null)
    emit('file-removed', props.modelValue)
  }
}

const getFileIndex = (file: File): number => {
  return fileList.value.indexOf(file)
}

const isFileTypeAccepted = (file: File): boolean => {
  if (props.accept === '*') return true
  
  const acceptedTypes = props.accept.split(',').map(type => type.trim())
  
  // Check by MIME type
  if (acceptedTypes.includes(file.type)) return true
  
  // Check by extension
  const fileName = file.name
  const fileExtension = fileName.substring(fileName.lastIndexOf('.')).toLowerCase()
  return acceptedTypes.some(type => type.toLowerCase() === fileExtension)
}

const isImage = (file: File): boolean => {
  return file.type.startsWith('image/')
}

const isPDF = (file: File): boolean => {
  return file.type === 'application/pdf'
}

const getImageUrl = (file: File): string => {
  if (!imageUrlMap.value.has(file)) {
    const url = URL.createObjectURL(file)
    imageUrlMap.value.set(file, url)
  }
  return imageUrlMap.value.get(file) || ''
}

const formatFileSize = (size: number): string => {
  if (size < 1024) return `${size} B`
  if (size < 1024 * 1024) return `${(size / 1024).toFixed(1)} KB`
  return `${(size / (1024 * 1024)).toFixed(1)} MB`
}

// Clean up object URLs when component is unmounted
onUnmounted(() => {
  imageUrlMap.value.forEach(url => {
    URL.revokeObjectURL(url)
  })
  imageUrlMap.value.clear()
})
</script>