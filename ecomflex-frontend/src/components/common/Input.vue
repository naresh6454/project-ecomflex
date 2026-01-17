<template>
  <div :class="rootClass">
    <!-- Label -->
    <label 
      v-if="label" 
      :for="id" 
      class="form-label"
      :class="{ 'text-red-600': error }"
    >
      {{ label }}
      <span v-if="required" class="text-red-600">*</span>
    </label>
    
    <div class="relative">
      <!-- Left Icon -->
      <div v-if="$slots['left-icon']" class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
        <slot name="left-icon"></slot>
      </div>
      
      <!-- Input Field -->
      <input
        :id="id"
        :type="type"
        :value="modelValue"
        :placeholder="placeholder"
        :disabled="disabled"
        :required="required"
        :name="name"
        :autocomplete="autocomplete"
        @input="handleInput"
        @focus="handleFocus"
        @blur="handleBlur"
        class="form-input"
        :class="[
          { 'pl-10': $slots['left-icon'] },
          { 'pr-10': $slots['right-icon'] || clearable },
          { 'border-red-500 focus:ring-red-500 focus:border-red-500': error },
          { 'opacity-60 cursor-not-allowed': disabled },
          inputClass
        ]"
      />
      
      <!-- Clear Button -->
      <button
        v-if="clearable && modelValue"
        type="button"
        @click="clearInput"
        class="absolute inset-y-0 right-0 pr-3 flex items-center cursor-pointer text-gray-400 hover:text-gray-600"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
          <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
        </svg>
      </button>
      
      <!-- Right Icon -->
      <div v-else-if="$slots['right-icon']" class="absolute inset-y-0 right-0 pr-3 flex items-center pointer-events-none">
        <slot name="right-icon"></slot>
      </div>
    </div>
    
    <!-- Helper Text -->
    <div v-if="helperText && !error" class="mt-1 text-sm text-gray-500">
      {{ helperText }}
    </div>
    
    <!-- Error Message -->
    <div v-if="error" class="form-error">
      {{ error }}
    </div>
  </div>
</template>

<script setup lang="ts">
interface Props {
  modelValue: string
  label?: string
  id?: string
  name?: string
  type?: string
  placeholder?: string
  helperText?: string
  error?: string
  disabled?: boolean
  required?: boolean
  clearable?: boolean
  autocomplete?: string
  rootClass?: string
  inputClass?: string
}

withDefaults(defineProps<Props>(), {
  type: 'text',
  placeholder: '',
  disabled: false,
  required: false,
  clearable: false,
  id: '',
  name: '',
  autocomplete: 'off',
  rootClass: '',
  inputClass: ''
})

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void
  (e: 'focus', event: FocusEvent): void
  (e: 'blur', event: FocusEvent): void
  (e: 'clear'): void
}>()

// Handle input
const handleInput = (event: Event) => {
  const target = event.target as HTMLInputElement
  emit('update:modelValue', target.value)
}

const handleFocus = (event: FocusEvent) => emit('focus', event)
const handleBlur = (event: FocusEvent) => emit('blur', event)
const clearInput = () => {
  emit('update:modelValue', '')
  emit('clear')
}
</script>