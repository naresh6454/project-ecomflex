<template>
  <div class="min-h-screen bg-brand-50 flex items-center justify-center px-4 sm:px-6 lg:px-8">
    <div 
      v-motion
      :initial="{ opacity: 0, scale: 0.9 }"
      :enter="{ opacity: 1, scale: 1, transition: { duration: 500 } }"
      class="max-w-md w-full text-center"
    >
      <!-- 404 Animation -->
      <div class="relative h-48 mb-6">
        <div class="absolute inset-0 flex items-center justify-center">
          <div class="text-9xl font-bold text-brand-200">404</div>
        </div>
        <div class="absolute inset-0 flex items-center justify-center">
          <div 
            v-motion
            :initial="{ opacity: 0, y: 20 }"
            :enter="{ opacity: 1, y: 0, transition: { delay: 300, duration: 500 } }"
            class="text-5xl font-bold text-brand-600"
          >
            404
          </div>
        </div>
        <div class="absolute w-full bottom-0 flex justify-center">
          <div 
            v-motion
            :initial="{ opacity: 0 }"
            :enter="{ opacity: 1, transition: { delay: 600, duration: 500 } }"
            class="h-1 w-32 bg-gradient-to-r from-brand-400 to-accent rounded-full"
          ></div>
        </div>
      </div>
      
      <h2 
        v-motion
        :initial="{ opacity: 0 }"
        :enter="{ opacity: 1, transition: { delay: 800, duration: 500 } }"
        class="text-3xl font-extrabold text-gray-900 tracking-tight"
      >
        Page not found
      </h2>
      <p 
        v-motion
        :initial="{ opacity: 0 }"
        :enter="{ opacity: 1, transition: { delay: 1000, duration: 500 } }"
        class="mt-2 text-lg text-gray-500"
      >
        The page you're looking for doesn't exist or has been moved.
      </p>
      <div 
        v-motion
        :initial="{ opacity: 0, y: 10 }"
        :enter="{ opacity: 1, y: 0, transition: { delay: 1200, duration: 500 } }"
        class="mt-6"
      >
        <router-link 
          to="/" 
          class="btn btn-primary ripple"
        >
          Go back home
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// 404 Not Found page
</script>