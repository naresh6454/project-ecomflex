<template>
  <div
  class="bg-white overflow-hidden shadow rounded-lg"
  v-motion="{
    initial: { opacity: 0, y: 20 },
    enter: { opacity: 1, y: 0, transition: { delay: (delay ?? 0) * 100 } }
  }"
>
  <div class="p-5">
    <div class="flex items-center">
      <div 
        class="flex-shrink-0 rounded-md p-3"
        :class="iconBackgroundClass">
        <component 
          :is="icon" 
          class="h-6 w-6" 
          :class="iconColorClass" 
          aria-hidden="true" />
      </div>
      <div class="ml-5 w-0 flex-1">
        <dl>
          <dt class="text-sm font-medium text-gray-500 truncate">{{ title }}</dt>
          <dd>
            <div class="text-lg font-medium text-gray-900">
              {{ formattedValue }}
            </div>
          </dd>
        </dl>
      </div>
    </div>
  </div>
  <!-- removed the extra </div> here -->
  <div class="bg-gray-50 px-5 py-3">
    <div class="text-sm">
      <a 
        :href="actionUrl" 
        class="font-medium text-blue-600 hover:text-blue-500 transition-colors">
        {{ actionText }}
      </a>
    </div>
  </div>
</div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  title: string
  value: number | string
  icon: object
  type?: 'default' | 'primary' | 'success' | 'warning' | 'danger'
  format?: 'none' | 'currency' | 'percent' | 'number'
  actionText?: string
  actionUrl?: string
  delay?: number
}>()


// Computed properties for dynamic styling
const iconBackgroundClass = computed(() => {
  const classes = {
    'default': 'bg-blue-100',
    'primary': 'bg-blue-500',
    'success': 'bg-green-100',
    'warning': 'bg-yellow-100',
    'danger': 'bg-red-100'
  }
return classes[props.type ?? 'default']
})

const iconColorClass = computed(() => {
  const classes = {
    'default': 'text-blue-600',
    'primary': 'text-white',
    'success': 'text-green-600',
    'warning': 'text-yellow-600',
    'danger': 'text-red-600'
  }
return classes[props.type ?? 'default']
})

const formattedValue = computed(() => {
  if (props.format === 'currency') {
    return new Intl.NumberFormat('en-US', { 
      style: 'currency', 
      currency: 'USD',
      minimumFractionDigits: 0,
      maximumFractionDigits: 0
    }).format(Number(props.value))
  } else if (props.format === 'percent') {
    return `${props.value}%`
  } else if (props.format === 'number') {
    return new Intl.NumberFormat('en-US').format(Number(props.value))
  } else {
    return props.value
  }
})
</script>