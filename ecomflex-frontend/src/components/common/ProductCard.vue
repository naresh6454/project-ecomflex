<template>
  <div 
    class="group h-full"
    :class="{ 'opacity-70': !product.isActive }"
  >
    <div class="card hover:shadow-xl transition-all duration-300 h-full flex flex-col overflow-hidden relative">
      <!-- Inactive Overlay -->
      <div v-if="!product.isActive" class="absolute inset-0 bg-gray-100 bg-opacity-50 z-10 flex items-center justify-center">
        <span class="badge badge-red px-3 py-1">Not Available</span>
      </div>

      <!-- Image - Fallback immediately to letter if no image URL -->
      <div class="aspect-w-4 aspect-h-3 bg-gray-200 rounded-t-xl overflow-hidden">
        <div v-if="hasProductImage && !imageError" class="relative w-full h-full">
          <img 
            :src="fixedImageUrl"
            :alt="product.name"
            class="w-full h-full object-cover transform group-hover:scale-105 transition-transform duration-500"
            @error="handleImageError" 
            ref="imageRef"
            crossorigin="anonymous"
          />
          <div v-if="isLoading" class="absolute inset-0 flex items-center justify-center bg-gray-100 bg-opacity-60">
            <div class="animate-pulse w-8 h-8 rounded-full border-2 border-brand-500 border-t-transparent"></div>
          </div>
        </div>
        <div v-else class="w-full h-full flex items-center justify-center bg-gradient-to-br from-brand-100 to-brand-200">
          <span class="text-brand-600 font-bold text-xl">{{ product.name.charAt(0) }}</span>
        </div>
      </div>
      
      <!-- Content -->
      <div class="flex-1 p-4 flex flex-col">
        <h3 class="text-lg font-medium text-gray-900 group-hover:text-brand-600 transition-colors duration-200">
          {{ product.name }}
        </h3>
        
        <div class="mt-2 flex items-center text-sm text-gray-500">
          <span class="font-medium">ASIN:</span>
          <span class="ml-1">{{ product.asin }}</span>
        </div>
        
        <div class="mt-1 flex items-center text-sm text-gray-500">
          <span class="font-medium">Founder:</span>
          <span class="ml-1">{{ product.founder }}</span>
        </div>
        
        <div class="mt-2 flex items-center flex-wrap gap-2">
          <Badge variant="gold">100% Cashback</Badge>
          <Badge 
            :variant="bookingPercentage < 50 ? 'green' : 'yellow'"
          >
            {{ product.currentBookings }}/{{ product.requiredBookings }} Booked
          </Badge>
        </div>
        
        <div class="mt-4 flex-grow flex flex-col justify-end">
          <div class="w-full bg-gray-200 rounded-full h-2">
            <div 
              class="bg-gradient-to-r from-brand-400 to-accent h-2 rounded-full" 
              :style="`width: ${bookingPercentage}%`"
            ></div>
          </div>
          
          <div class="mt-4 flex space-x-3">
            <slot name="actions">
              <Button 
                variant="primary" 
                block 
                @click="$emit('view', product.id)"
              >
                View Details
              </Button>
            </slot>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, onMounted } from 'vue'
import { Product } from '@/services/product.service'
import Button from './Button.vue'
import Badge from './Badge.vue'

// Refs
const imageRef = ref<HTMLImageElement | null>(null)
const imageError = ref(false)
const isLoading = ref(true)

interface Props {
  product: Product
}

const props = defineProps<Props>()

const emit = defineEmits<{
  (e: 'view', productId: string): void
  (e: 'book', productId: string): void
}>()

// Check for image in either field
const hasProductImage = computed(() => {
  const hasImage = !!(props.product.image || props.product.imageUrl);
  // If there's no image, no need to show loading state
  if (!hasImage) {
    isLoading.value = false;
  }
  return hasImage;
})

// Fixed URL calculation
const fixedImageUrl = computed(() => {
  // Get the original image source
  let originalUrl = props.product.imageUrl || props.product.image || '';
  
  // Skip processing if the URL is empty
  if (!originalUrl) {
    return '';
  }
  
  // If URL contains placeholder or incorrect bucket name, fix it
  if (originalUrl.includes('your-s3-bucket') || originalUrl.includes('placeholder')) {
    // Extract the filename from the URL
    const parts = originalUrl.split('/');
    const filename = parts[parts.length - 1];
    
    // Create a proper URL with the real bucket name
    return `https://ecomflex-uploads-dev.s3.eu-north-1.amazonaws.com/products/${filename}`;
  }
  
  // If it's a relative path without http/https, prefix with the S3 bucket URL
  if (!originalUrl.startsWith('http')) {
    // Handle relative paths with or without leading slash
    const path = originalUrl.startsWith('/') ? originalUrl.substring(1) : originalUrl;
    return `https://ecomflex-uploads-dev.s3.eu-north-1.amazonaws.com/${path}`;
  }
  
  // If URL looks good, use it as is
  return originalUrl;
})

// Computed
const bookingPercentage = computed(() => {
  const percentage = (props.product.currentBookings / props.product.requiredBookings) * 100
  return Math.min(Math.round(percentage), 100)
})

// Methods
const handleImageError = (_e: Event) => {  // Fixed: Added underscore prefix to indicate intentionally unused parameter
  console.error(`Image failed to load: ${props.product.imageUrl || props.product.image}`);
  console.error(`Attempted fixed URL: ${fixedImageUrl.value}`);
  
  // Set error state to trigger fallback
  imageError.value = true;
  isLoading.value = false;
}

onMounted(() => {
  // If image already loaded from cache, the error event might not fire
  if (imageRef.value) {
    if (imageRef.value.complete) {
      isLoading.value = false;
      if (imageRef.value.naturalWidth === 0) {
        handleImageError({ target: imageRef.value } as unknown as Event);
      }
    } else {
      // Add load event listener
      imageRef.value.addEventListener('load', () => {
        isLoading.value = false;
      });
    }
  }
})
</script>

<style scoped>
.animate-pulse {
  animation: pulse 1.5s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

@keyframes pulse {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.5;
  }
}
</style>