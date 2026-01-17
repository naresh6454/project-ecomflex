<template>
  <div class="p-4 border rounded bg-gray-100 m-4">
    <h2 class="text-lg font-bold mb-4">S3 Image Direct Test</h2>
    
    <div class="space-y-4">
      <!-- Test 1: Direct URL -->
      <div class="border p-2 rounded bg-white">
        <div class="font-medium">Test 1: Direct URL from database</div>
        <div class="text-xs text-gray-500 mb-2">{{ directUrl }}</div>
        <img :src="directUrl" alt="Direct test" class="h-24 object-cover" />
      </div>
      
      <!-- Test 2: No domain URL -->
      <div class="border p-2 rounded bg-white">
        <div class="font-medium">Test 2: Path-only URL</div>
        <div class="text-xs text-gray-500 mb-2">{{ pathOnlyUrl }}</div>
        <img :src="pathOnlyUrl" alt="Path-only test" class="h-24 object-cover" />
      </div>
      
      <!-- Test 3: Hardcoded URL -->
      <div class="border p-2 rounded bg-white">
        <div class="font-medium">Test 3: Hardcoded filename</div>
        <div class="text-xs text-gray-500 mb-2">{{ hardcodedUrl }}</div>
        <img :src="hardcodedUrl" alt="Hardcoded test" class="h-24 object-cover" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

// The product ID from your database
const productId = 'a77961d3-ea5e-401a-ba6f-1f8632f04c85'

// Direct URL from your database
const directUrl = ref(`https://ecomflex-uploads-dev.s3.eu-north-1.amazonaws.com/products/${productId}-1760098163851479974.png?t=${Date.now()}`)

// Path-only URL
const pathOnlyUrl = ref(`/products/${productId}-1760098163851479974.png?t=${Date.now()}`)

// Try with a hardcoded filename
const hardcodedUrl = ref(`https://ecomflex-uploads-dev.s3.eu-north-1.amazonaws.com/uploads/${productId}-1760098163851479974.png?t=${Date.now()}`)

onMounted(async () => {
  console.log('S3 Test URLs:')
  console.log('Direct URL:', directUrl.value)
  console.log('Path-only URL:', pathOnlyUrl.value)
  console.log('Hardcoded URL:', hardcodedUrl.value)
  
  // Direct test with the browser's Fetch API
  try {
    const response = await fetch(directUrl.value, { method: 'HEAD' })
    console.log(`Fetch test result for direct URL: ${response.status} ${response.statusText}`)
  } catch (error) {
    console.error('Fetch test error:', error)
  }
})
</script>