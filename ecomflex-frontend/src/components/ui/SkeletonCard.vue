<template>
  <div class="rounded-xl bg-white shadow-md overflow-hidden dark:bg-gray-800">
    <!-- Image Skeleton -->
    <SkeletonLoader
      variant="image"
      height="200"
      :shimmer="true"
      color="light"
    />
    
    <!-- Content Skeleton -->
    <div class="p-4 space-y-3">
      <!-- Title -->
      <SkeletonLoader
        variant="text"
        width="3/4"
        height="6"
        color="dark"
        :shimmer="true"
      />
      
      <!-- Subtitle -->
      <SkeletonLoader
        variant="text"
        width="1/2"
        height="4"
        color="default"
        :shimmer="true"
      />
      
      <!-- Description Lines -->
      <div class="space-y-2 pt-2">
        <SkeletonLoader
          v-for="i in 3"
          :key="i"
          variant="text"
          width="full"
          height="4"
          color="light"
          :shimmer="true"
        />
      </div>
      
      <!-- Footer Skeleton -->
      <div class="flex justify-between items-center pt-4">
        <SkeletonLoader
          variant="button"
          width="1/3"
          height="10"
          color="accent"
          :shimmer="true"
        />
        
        <SkeletonLoader
          variant="circle"
          width="10"
          height="10"
          color="light"
          :shimmer="true"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import SkeletonLoader from './SkeletonLoader.vue'
</script>