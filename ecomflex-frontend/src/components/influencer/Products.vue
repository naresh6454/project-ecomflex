<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <!-- Page Header -->
    <div 
      v-motion
      :initial="{ opacity: 0, y: -20 }"
      :enter="{ opacity: 1, y: 0 }"
      class="md:flex md:items-center md:justify-between mb-8">
      <div class="flex-1 min-w-0">
        <h1 class="text-2xl font-bold leading-7 text-gray-900 sm:text-3xl sm:truncate">
          Products to Promote
        </h1>
        <p class="mt-1 text-sm text-gray-500">
          Browse all available products that you can promote to your audience.
        </p>
      </div>
      <div class="mt-4 flex md:mt-0 md:ml-4">
        <div class="relative">
          <label for="search" class="sr-only">Search products</label>
          <div class="relative rounded-md shadow-sm">
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <svg class="h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" clip-rule="evenodd" />
              </svg>
            </div>
            <input 
              id="search" 
              type="text" 
              v-model="searchQuery"
              class="focus:ring-blue-500 focus:border-blue-500 block w-full pl-10 sm:text-sm border-gray-300 rounded-md" 
              placeholder="Search products"
            />
          </div>
        </div>
      </div>
    </div>

    <!-- Categories -->
    <div class="mb-8">
      <div class="flex items-center mb-4">
        <h2 class="text-lg font-medium text-gray-900">Categories</h2>
      </div>
      <div 
        class="flex flex-wrap gap-2"
        v-motion
        :initial="{ opacity: 0, y: 20 }"
        :enter="{ opacity: 1, y: 0 }">
        <button 
          v-for="category in categories" 
          :key="category.id"
          @click="selectedCategory = category.id === selectedCategory ? 'all' : category.id"
          class="inline-flex items-center px-4 py-2 rounded-full text-sm font-medium transition-colors"
          :class="selectedCategory === category.id ? 
            'bg-blue-100 text-blue-800' : 
            'bg-gray-100 text-gray-800 hover:bg-gray-200'"
        >
          {{ category.name }}
        </button>
      </div>
    </div>

    <!-- Product Grid -->
    <div class="mb-8">
      <div class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
        <div 
          v-for="(product, index) in filteredProducts" 
          :key="product.id"
          v-motion
          :initial="{ opacity: 0, y: 20 }"
          :enter="{ opacity: 1, y: 0, transition: { delay: index * 100 } }"
          class="bg-white overflow-hidden shadow rounded-lg transition-all duration-300 hover:shadow-xl">
          <div class="relative pb-2/3">
            <img :src="product.image" :alt="product.name" class="absolute h-full w-full object-cover" />
            <div class="absolute top-0 right-0 m-2">
              <span 
                class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                :class="product.category === 'Electronics' ? 'bg-blue-100 text-blue-800' :
                       product.category === 'Home' ? 'bg-green-100 text-green-800' :
                       product.category === 'Beauty' ? 'bg-purple-100 text-purple-800' :
                       'bg-gray-100 text-gray-800'">
                {{ product.category }}
              </span>
            </div>
          </div>
          <div class="px-4 py-5 sm:p-6">
            <h3 class="text-lg font-medium text-gray-900 truncate">{{ product.name }}</h3>
            <div class="mt-2 flex items-center text-sm text-gray-500">
              <span>ASIN: {{ product.asin }}</span>
            </div>
            <div class="mt-3 flex items-center">
              <div class="flex-1">
                <span class="text-xl font-bold text-blue-600">{{ product.commission }}</span>
                <span class="text-sm text-gray-500"> commission</span>
              </div>
              <div class="text-sm text-gray-500">
                <span class="font-medium text-gray-900">{{ product.bookings }}</span> bookings
              </div>
            </div>
            <div class="mt-5 flex justify-between items-center">
              <a :href="product.merchantLink" target="_blank" class="text-sm font-medium text-blue-600 hover:text-blue-500 transition-colors">
                View on Merchant
              </a>
              <button 
                @click="generateShareLink(product)"
                class="inline-flex items-center px-3 py-1.5 border border-transparent text-xs font-medium rounded shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
                </svg>
                Share
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- No Products Found -->
    <div v-if="filteredProducts.length === 0" class="text-center py-12 bg-white shadow rounded-lg">
      <svg class="mx-auto h-12 w-12 text-gray-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.618 5.984A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016zM12 9v2m0 4h.01" />
      </svg>
      <h3 class="mt-2 text-sm font-medium text-gray-900">No products found</h3>
      <p class="mt-1 text-sm text-gray-500">
        Try changing your search or filter to find products to promote.
      </p>
    </div>

    <!-- Share Modal -->
    <div v-if="shareModalOpen" class="fixed z-10 inset-0 overflow-y-auto" aria-labelledby="modal-title" role="dialog" aria-modal="true">
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <!-- Background overlay -->
        <div 
          class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" 
          aria-hidden="true"
          v-motion
          :initial="{ opacity: 0 }"
          :enter="{ opacity: 1 }"
          :leave="{ opacity: 0 }"
          @click="shareModalOpen = false"
        ></div>

        <!-- Modal panel -->
        <div 
          class="inline-block align-bottom bg-white rounded-lg px-4 pt-5 pb-4 text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full sm:p-6"
          v-motion
          :initial="{ opacity: 0, scale: 0.95 }"
          :enter="{ opacity: 1, scale: 1 }"
          :leave="{ opacity: 0, scale: 0.95 }">
          <div>
            <div class="mt-3 text-center sm:mt-5">
              <h3 class="text-lg leading-6 font-medium text-gray-900" id="modal-title">
                Share {{ selectedProduct?.name }}
              </h3>
              <div class="mt-2">
                <p class="text-sm text-gray-500">
                  Share this product with your audience. They'll get 100% cashback, and you'll earn commission.
                </p>
              </div>
            </div>
          </div>

          <!-- Referral Link -->
          <div class="mt-5">
            <label for="referral-link" class="block text-sm font-medium text-gray-700">Your Referral Link</label>
            <div class="mt-1 flex rounded-md shadow-sm">
              <input 
                type="text" 
                id="referral-link" 
                class="focus:ring-blue-500 focus:border-blue-500 flex-1 block w-full rounded-none rounded-l-md sm:text-sm border-gray-300" 
                :value="referralLink" 
                readonly 
              />
              <button 
                type="button" 
                class="inline-flex items-center px-3 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-r-md text-gray-700 bg-gray-50 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                @click="copyReferralLink">
                {{ copied ? 'Copied' : 'Copy' }}
              </button>
            </div>
          </div>

          <!-- Share Options -->
          <div class="mt-5">
            <label class="block text-sm font-medium text-gray-700">Share On</label>
            <div class="mt-2 flex flex-wrap gap-2">
              <button 
                type="button"
                class="inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 transition-colors"
                @click="shareToWhatsApp">
                <svg class="h-4 w-4 mr-1.5" fill="currentColor" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                  <path d="M.057 24l1.687-6.163c-1.041-1.804-1.588-3.849-1.587-5.946.003-6.556 5.338-11.891 11.893-11.891 3.181.001 6.167 1.24 8.413 3.488 2.245 2.248 3.481 5.236 3.48 8.414-.003 6.557-5.338 11.892-11.893 11.892-1.99-.001-3.951-.5-5.688-1.448l-6.305 1.654zm6.597-3.807c1.676.995 3.276 1.591 5.392 1.592 5.448 0 9.886-4.434 9.889-9.885.002-5.462-4.415-9.89-9.881-9.892-5.452 0-9.887 4.434-9.889 9.884-.001 2.225.651 3.891 1.746 5.634l-.999 3.648 3.742-.981zm11.387-5.464c-.074-.124-.272-.198-.57-.347-.297-.149-1.758-.868-2.031-.967-.272-.099-.47-.149-.669.149-.198.297-.768.967-.941 1.165-.173.198-.347.223-.644.074-.297-.149-1.255-.462-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.297-.347.446-.521.151-.172.2-.296.3-.495.099-.198.05-.372-.025-.521-.075-.148-.669-1.611-.916-2.206-.242-.579-.487-.501-.669-.51l-.57-.01c-.198 0-.52.074-.792.372s-1.04 1.016-1.04 2.479 1.065 2.876 1.213 3.074c.149.198 2.095 3.2 5.076 4.487.709.306 1.263.489 1.694.626.712.226 1.36.194 1.872.118.571-.085 1.758-.719 2.006-1.413.248-.695.248-1.29.173-1.414z"/>
                </svg>
                WhatsApp
              </button>

              <button 
                type="button"
                class="inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md shadow-sm text-white bg-blue-400 hover:bg-blue-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-400 transition-colors"
                @click="shareToTwitter">
                <svg class="h-4 w-4 mr-1.5" fill="currentColor" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                  <path d="M23.953 4.57a10 10 0 01-2.825.775 4.958 4.958 0 002.163-2.723c-.951.555-2.005.959-3.127 1.184a4.92 4.92 0 00-8.384 4.482C7.69 8.095 4.067 6.13 1.64 3.162a4.822 4.822 0 00-.666 2.475c0 1.71.87 3.213 2.188 4.096a4.904 4.904 0 01-2.228-.616v.06a4.923 4.923 0 003.946 4.827 4.996 4.996 0 01-2.212.085 4.936 4.936 0 004.604 3.417 9.867 9.867 0 01-6.102 2.105c-.39 0-.779-.023-1.17-.067a13.995 13.995 0 007.557 2.209c9.053 0 13.998-7.496 13.998-13.985 0-.21 0-.42-.015-.63A9.935 9.935 0 0024 4.59z"/>
                </svg>
                Twitter
              </button>

              <button 
                type="button"
                class="inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-600 transition-colors"
                @click="shareToFacebook">
                <svg class="h-4 w-4 mr-1.5" fill="currentColor" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                  <path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/>
                </svg>
                Facebook
              </button>
            </div>
          </div>

          <!-- Custom Message -->
          <div class="mt-5">
            <label for="message" class="block text-sm font-medium text-gray-700">Custom Message (Optional)</label>
            <div class="mt-1">
              <textarea 
                id="message" 
                name="message" 
                rows="3" 
                class="shadow-sm focus:ring-blue-500 focus:border-blue-500 block w-full sm:text-sm border-gray-300 rounded-md"
                placeholder="Add a personal message to your audience..."
                v-model="shareMessage"></textarea>
            </div>
          </div>

          <div class="mt-5 sm:mt-6 sm:grid sm:grid-cols-2 sm:gap-3 sm:grid-flow-row-dense">
            <button 
              type="button" 
              class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-blue-600 text-base font-medium text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 sm:col-start-2 sm:text-sm transition-colors"
              @click="copyReferralLink">
              Copy Link
            </button>
            <button 
              type="button" 
              class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 sm:mt-0 sm:col-start-1 sm:text-sm transition-colors"
              @click="shareModalOpen = false">
              Close
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useReferralStore } from '@/stores/referral'

// Define TypeScript interfaces
interface Category {
  id: string;
  name: string;
}

interface Product {
  id: number;
  name: string;
  asin: string;
  category: string;
  commission: string;
  bookings: number;
  image: string;
  merchantLink: string;
}

// Referral store
const referralStore = useReferralStore()

// State
const searchQuery = ref('')
const selectedCategory = ref('all')
const shareModalOpen = ref(false)
const selectedProduct = ref<Product | null>(null)
const referralLink = ref('')
const copied = ref(false)
const shareMessage = ref('')

// Categories
const categories: Category[] = [
  { id: 'all', name: 'All Categories' },
  { id: 'Electronics', name: 'Electronics' },
  { id: 'Home', name: 'Home & Kitchen' },
  { id: 'Beauty', name: 'Beauty & Personal Care' },
  { id: 'Fitness', name: 'Fitness' }
]

// Mock products data
const products = ref<Product[]>([
  {
    id: 1,
    name: 'Wireless Bluetooth Earbuds',
    asin: 'B08JTFPCP7',
    category: 'Electronics',
    commission: '$25',
    bookings: 48,
    image: 'https://images.unsplash.com/photo-1606741965429-8d76ff50bb2e?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=687&q=80',
    merchantLink: 'https://example.com/product/1'
  },
  {
    id: 2,
    name: 'Smart Fitness Watch',
    asin: 'B09QPNQM2D',
    category: 'Electronics',
    commission: '$30',
    bookings: 35,
    image: 'https://images.unsplash.com/photo-1508685096489-7aacd43bd3b1?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=627&q=80',
    merchantLink: 'https://example.com/product/2'
  },
  {
    id: 3,
    name: 'Portable Bluetooth Speaker',
    asin: 'B07PBVN4GH',
    category: 'Electronics',
    commission: '$20',
    bookings: 62,
    image: 'https://images.unsplash.com/photo-1608043152269-423dbba4e7e1?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=687&q=80',
    merchantLink: 'https://example.com/product/3'
  },
  {
    id: 4,
    name: 'Smart Home Security Camera',
    asin: 'B08R59YH7W',
    category: 'Home',
    commission: '$35',
    bookings: 29,
    image: 'https://images.unsplash.com/photo-1585314062340-f1a5a7c9328d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=687&q=80',
    merchantLink: 'https://example.com/product/4'
  },
  {
    id: 5,
    name: 'LED Desk Lamp with Wireless Charger',
    asin: 'B08Y8HXTBG',
    category: 'Home',
    commission: '$18',
    bookings: 43,
    image: 'https://images.unsplash.com/photo-1618641986557-1ecd230959aa?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=687&q=80',
    merchantLink: 'https://example.com/product/5'
  },
  {
    id: 6,
    name: 'Organic Face Serum',
    asin: 'B07V2BVPGM',
    category: 'Beauty',
    commission: '$22',
    bookings: 51,
    image: 'https://images.unsplash.com/photo-1570194065650-d682c29853a6?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=687&q=80',
    merchantLink: 'https://example.com/product/6'
  }
])

// Filter products based on search and category
const filteredProducts = computed(() => {
  return products.value.filter(product => {
    // Search filter
    const searchMatch = 
      searchQuery.value === '' || 
      product.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      product.asin.toLowerCase().includes(searchQuery.value.toLowerCase())
    
    // Category filter
    const categoryMatch = 
      selectedCategory.value === 'all' || 
      product.category === selectedCategory.value
    
    return searchMatch && categoryMatch
  })
})

// Generate share link
const generateShareLink = (product: Product) => {
  selectedProduct.value = product
  shareModalOpen.value = true
  
  // Generate referral link
  const baseUrl = window.location.origin
  referralLink.value = `${baseUrl}/product/${product.id}?ref=SARAH25`
  
  // Reset copied state
  copied.value = false
  
  // Set default share message
  shareMessage.value = `Check out this amazing ${product.name}! Use my referral code and get 100% cashback.`
}

// Copy referral link
const copyReferralLink = () => {
  navigator.clipboard.writeText(referralLink.value)
  copied.value = true
  setTimeout(() => {
    copied.value = false
  }, 2000)
}

// Social media sharing functions
const shareToWhatsApp = () => {
  const message = encodeURIComponent(`${shareMessage.value} ${referralLink.value}`)
  window.open(`https://wa.me/?text=${message}`, '_blank')
}

const shareToTwitter = () => {
  const message = encodeURIComponent(`${shareMessage.value} ${referralLink.value}`)
  window.open(`https://twitter.com/intent/tweet?text=${message}`, '_blank')
}

const shareToFacebook = () => {
  window.open(`https://www.facebook.com/sharer/sharer.php?u=${encodeURIComponent(referralLink.value)}`, '_blank')
}

// Fetch data on component mount
onMounted(async () => {
  await referralStore.fetchReferralCode()
})
</script>

<style scoped>
.pb-2\/3 {
  padding-bottom: 66.666667%;
}
</style>