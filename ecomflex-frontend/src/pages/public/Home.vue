<template>
  <AppLayout>
    <!-- Hero Section -->
    <div class="bg-gradient-to-r from-brand-50 to-brand-100 py-12">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center">
          <h1 class="text-4xl font-extrabold text-gray-900 sm:text-5xl sm:tracking-tight lg:text-6xl">
            Discover Amazing Products
          </h1>
          <p class="mt-5 max-w-xl mx-auto text-xl text-gray-500">
            Browse our selection of premium products with 100% cashback potential.
          </p>
          <!-- Search bar in hero section -->
          <div class="mt-8 max-w-md mx-auto">
            <div class="relative rounded-full shadow-sm">
              <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <svg class="h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                  <path fill-rule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" clip-rule="evenodd" />
                </svg>
              </div>
              <Input
                v-model="searchQuery"
                type="text"
                placeholder="Search for products..."
                class="pl-10 block w-full rounded-full border-gray-300 focus:ring-brand-500 focus:border-brand-500"
                @input="handleSearch"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Main Content -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- Filter Bar -->
      <div class="flex flex-col md:flex-row justify-between items-start md:items-center mb-8 gap-4">
        <h2 class="text-2xl font-bold text-gray-900">Available Products</h2>
        
        <div class="flex flex-wrap gap-3">
          <!-- Category Filter -->
          <select
            v-model="selectedCategory"
            class="form-input rounded-md"
            @change="applyFilters"
          >
            <option value="">All Categories</option>
            <option v-for="category in categories" :key="category" :value="category">
              {{ category }}
            </option>
          </select>
          
          <!-- Sort By -->
          <select
            v-model="sortBy"
            class="form-input rounded-md"
            @change="applyFilters"
          >
            <option value="newest">Newest First</option>
            <option value="popular">Most Popular</option>
          </select>
        </div>
      </div>
      
      <!-- Loading State -->
      <div v-if="loading" class="flex justify-center py-12">
        <div class="animate-pulse-gold inline-block h-12 w-12 rounded-full border-4 border-t-accent border-r-transparent border-b-transparent border-l-transparent"></div>
      </div>
      
      <!-- Empty State -->
      <div v-else-if="allProducts.length === 0" class="py-12 text-center">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 mx-auto text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
        </svg>
        <h3 class="mt-4 text-lg font-medium text-gray-900">No products found</h3>
        <p class="mt-1 text-gray-500">Try adjusting your search or filter criteria.</p>
        <Button 
          @click="resetFilters" 
          class="mt-4"
          variant="primary"
        >
          Reset Filters
        </Button>
      </div>
      
      <!-- Product Grid -->
      <div 
        v-else 
        class="grid grid-cols-1 gap-y-8 gap-x-6 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4"
      >
        <div 
          v-for="(product, index) in allProducts" 
          :key="product.id"
          v-motion
          :initial="{ opacity: 0, y: 20 }"
          :enter="{ opacity: 1, y: 0, transition: { delay: index * 100, duration: 500 } }"
        >
          <ProductCard 
            :product="product"
            @view="viewProductDetails"
            @book="bookProduct"
          >
            <template #actions>
              <div class="flex flex-col space-y-2">
                <div class="flex space-x-2">
                  <Button 
                    variant="secondary" 
                    class="flex-1"
                    @click="viewProductDetails(product.id)"
                  >
                    View Details
                  </Button>
                  <Button 
                    variant="primary" 
                    class="flex-1"
                    @click="bookProduct(product.id)"
                  >
                    Book Now
                  </Button>
                </div>
                <div class="flex space-x-2">
                  <Button 
                    variant="secondary" 
                    class="flex-1"
                    @click="openUploadBookingModal(product.id)"
                  >
                    Upload Booking
                  </Button>
                  <Button 
                    variant="secondary" 
                    class="flex-1"
                    @click="openUploadReviewModal(product.id)"
                  >
                    Upload Review
                  </Button>
                </div>
              </div>
            </template>
          </ProductCard>
        </div>
      </div>
      
      <!-- Pagination -->
      <div v-if="allProducts.length > 0" class="mt-12 flex justify-center">
        <nav class="flex items-center space-x-2">
          <button
            @click="prevPage"
            :disabled="currentPage === 1"
            :class="[
              'px-3 py-2 rounded-md',
              currentPage === 1 
                ? 'text-gray-400 cursor-not-allowed' 
                : 'text-gray-700 hover:bg-gray-100'
            ]"
          >
            Previous
          </button>
          
          <div v-for="page in totalPages" :key="page">
            <button
              @click="goToPage(page)"
              :class="[
                'px-3 py-2 rounded-md',
                currentPage === page
                  ? 'bg-brand-600 text-white'
                  : 'text-gray-700 hover:bg-gray-100'
              ]"
            >
              {{ page }}
            </button>
          </div>
          
          <button
            @click="nextPage"
            :disabled="currentPage === totalPages"
            :class="[
              'px-3 py-2 rounded-md',
              currentPage === totalPages
                ? 'text-gray-400 cursor-not-allowed'
                : 'text-gray-700 hover:bg-gray-100'
            ]"
          >
            Next
          </button>
        </nav>
      </div>
    </div>
    
    <!-- Product Details Modal -->
    <Modal 
      v-if="selectedProduct" 
      v-model="showModal"
    >
      <div class="p-6">
        <div class="flex flex-col md:flex-row gap-8">
          <!-- Product Image -->
          <div class="w-full md:w-2/5">
            <div class="aspect-w-1 aspect-h-1 rounded-lg bg-gray-200 overflow-hidden">
              <img 
                v-if="selectedProduct.image" 
                :src="selectedProduct.image" 
                :alt="selectedProduct.name" 
                class="w-full h-full object-center object-cover"
              />
              <div v-else class="w-full h-full flex items-center justify-center bg-gradient-to-br from-brand-100 to-brand-200">
                <span class="text-brand-600 font-bold text-4xl">{{ selectedProduct.name.charAt(0) }}</span>
              </div>
            </div>
          </div>
          
          <!-- Product Info -->
          <div class="w-full md:w-3/5">
            <h3 class="text-2xl font-bold text-gray-900">{{ selectedProduct.name }}</h3>
            
            <div class="mt-4 space-y-3">
              <div>
                <span class="font-medium text-gray-700">ASIN:</span>
                <span class="ml-2 text-gray-600">{{ selectedProduct.asin }}</span>
              </div>
              
              <div>
                <span class="font-medium text-gray-700">Founder:</span>
                <span class="ml-2 text-gray-600">{{ selectedProduct.founder }}</span>
              </div>
              
              <div v-if="selectedProduct.details">
                <span class="font-medium text-gray-700">Details:</span>
                <p class="mt-1 text-gray-600 whitespace-pre-line">{{ selectedProduct.details }}</p>
              </div>
              
              <div class="pt-4">
                <span class="font-medium text-gray-700">Booking Progress:</span>
                <div class="mt-2">
                  <div class="w-full h-2 bg-gray-200 rounded-full">
                    <div 
                      class="h-2 bg-gradient-to-r from-brand-400 to-accent rounded-full" 
                      :style="`width: ${getBookingPercentage(selectedProduct)}%`"
                    ></div>
                  </div>
                  <div class="mt-1 text-sm text-gray-500">
                    {{ selectedProduct.currentBookings }}/{{ selectedProduct.requiredBookings }} bookings
                  </div>
                </div>
              </div>
              
              <div class="flex flex-wrap gap-2 pt-2">
                <Badge variant="gold">100% Cashback</Badge>
                <Badge 
                  :variant="getBookingPercentage(selectedProduct) < 50 ? 'green' : 'yellow'"
                >
                  {{ getBookingPercentage(selectedProduct) }}% Booked
                </Badge>
              </div>
            </div>
            
            <div class="mt-8">
              <a 
                :href="selectedProduct.productLink" 
                target="_blank"
                rel="noopener noreferrer"
                class="text-brand-600 hover:text-brand-700 font-medium flex items-center"
              >
                <span>View on Amazon</span>
                <svg class="ml-1 h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd" d="M5.22 14.78a.75.75 0 001.06 0l7.22-7.22v5.69a.75.75 0 001.5 0v-7.5a.75.75 0 00-.75-.75h-7.5a.75.75 0 000 1.5h5.69l-7.22 7.22a.75.75 0 000 1.06z" clip-rule="evenodd" />
                </svg>
              </a>
            </div>
            
            <div class="mt-6 flex flex-col space-y-4">
              <Button 
                variant="primary" 
                block
                @click="bookProduct(selectedProduct.id)"
              >
                Book Now
              </Button>
              
              <div class="grid grid-cols-2 gap-4">
                <Button 
                  variant="secondary" 
                  block
                  @click="openUploadBookingModal(selectedProduct.id)"
                >
                  Upload Booking Details
                </Button>
                <Button 
                  variant="secondary" 
                  block
                  @click="openUploadReviewModal(selectedProduct.id)"
                >
                  Upload Review
                </Button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Modal>
    
    <!-- Upload Booking Details Modal -->
    <Modal 
      v-if="showUploadBookingModal" 
      v-model="showUploadBookingModal"
      title="Upload Booking Details"
    >
      <div class="p-6">
        <div class="mb-6">
          <p class="text-gray-600 mb-4">
            Please upload your booking confirmation or receipt to verify your purchase. 
            We accept PDF documents and images (JPG, PNG).
          </p>
          
          <div 
            class="border-2 border-dashed border-gray-300 rounded-lg p-6 text-center cursor-pointer hover:bg-gray-50 transition-colors"
            @click="triggerFileInput('bookingFileInput')"
          >
            <div v-if="!bookingFiles.length" class="space-y-3">
              <svg class="mx-auto h-12 w-12 text-gray-400" stroke="currentColor" fill="none" viewBox="0 0 48 48">
                <path d="M28 8H12a4 4 0 00-4 4v20m32-12v8m0 0v8a4 4 0 01-4 4h-8m-12 0v-8m0 0v-8m0 0V8a4 4 0 014-4h8m8 0h4a4 4 0 014 4v4m-4 4h-4" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
              </svg>
              <div class="flex text-sm text-gray-600 justify-center">
                <span class="relative rounded-md font-medium text-brand-600 hover:text-brand-500">
                  <span>Upload a file</span>
                </span>
                <p class="pl-1">or drag and drop</p>
              </div>
              <p class="text-xs text-gray-500">
                PDF, PNG, JPG up to 10MB
              </p>
            </div>
            
            <div v-else class="space-y-2">
              <div v-for="(file, index) in bookingFiles" :key="index" class="flex items-center justify-between bg-gray-100 p-2 rounded">
                <div class="flex items-center">
                  <svg class="h-6 w-6 text-gray-500 mr-2" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                  </svg>
                  <span class="text-sm font-medium text-gray-700 truncate max-w-xs">{{ file.name }}</span>
                </div>
                <button 
                  class="text-red-500 hover:text-red-700"
                  @click.stop="removeBookingFile(index)"
                >
                  <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </button>
              </div>
              
              <p class="text-sm text-brand-600 hover:text-brand-500 cursor-pointer">
                + Add more files
              </p>
            </div>
          </div>
          
          <input 
            id="bookingFileInput"
            ref="bookingFileInput"
            type="file" 
            class="hidden"
            accept=".pdf,.jpg,.jpeg,.png"
            multiple
            @change="handleBookingFileChange"
          />
        </div>
        
        <div class="flex justify-end space-x-3">
          <Button 
            variant="secondary" 
            @click="showUploadBookingModal = false"
          >
            Cancel
          </Button>
          <Button 
            variant="primary" 
            :disabled="bookingFiles.length === 0 || isUploading"
            @click="uploadBookingFiles"
          >
            <span v-if="isUploading">Uploading...</span>
            <span v-else>Upload</span>
          </Button>
        </div>
      </div>
    </Modal>
    
    <!-- Upload Review Modal -->
    <Modal 
      v-if="showUploadReviewModal" 
      v-model="showUploadReviewModal"
      title="Upload Product Review"
    >
      <div class="p-6">
        <div class="mb-6">
          <p class="text-gray-600 mb-4">
            Share your experience with this product. You can upload videos, images, or documents 
            to support your review.
          </p>
          
          <!-- Review Text -->
          <div class="mb-4">
            <label for="reviewText" class="block text-sm font-medium text-gray-700 mb-1">Your Review (optional)</label>
            <textarea 
              id="reviewText" 
              v-model="reviewText" 
              rows="3" 
              class="form-input w-full"
              placeholder="Write your thoughts about this product..."
            ></textarea>
          </div>
          
          <!-- Rating -->
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-1">Rating</label>
            <div class="flex space-x-1">
              <button 
                v-for="i in 5" 
                :key="i" 
                @click="rating = i"
                class="text-2xl focus:outline-none"
                :class="rating >= i ? 'text-yellow-400' : 'text-gray-300'"
              >
                ★
              </button>
            </div>
          </div>
          
          <!-- File Upload -->
          <div 
            class="border-2 border-dashed border-gray-300 rounded-lg p-6 text-center cursor-pointer hover:bg-gray-50 transition-colors mt-4"
            @click="triggerFileInput('reviewFileInput')"
          >
            <div v-if="!reviewFiles.length" class="space-y-3">
              <svg class="mx-auto h-12 w-12 text-gray-400" stroke="currentColor" fill="none" viewBox="0 0 48 48">
                <path d="M28 8H12a4 4 0 00-4 4v20m32-12v8m0 0v8a4 4 0 01-4 4h-8m-12 0v-8m0 0v-8m0 0V8a4 4 0 014-4h8m8 0h4a4 4 0 014 4v4m-4 4h-4" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
              </svg>
              <div class="flex text-sm text-gray-600 justify-center">
                <span class="relative rounded-md font-medium text-brand-600 hover:text-brand-500">
                  <span>Upload media</span>
                </span>
                <p class="pl-1">or drag and drop</p>
              </div>
              <p class="text-xs text-gray-500">
                PDF, PNG, JPG, MP4 up to 50MB
              </p>
            </div>
            
            <div v-else class="space-y-2">
              <div v-for="(file, index) in reviewFiles" :key="index" class="flex items-center justify-between bg-gray-100 p-2 rounded">
                <div class="flex items-center">
                  <svg class="h-6 w-6 text-gray-500 mr-2" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                  </svg>
                  <span class="text-sm font-medium text-gray-700 truncate max-w-xs">{{ file.name }}</span>
                </div>
                <button 
                  class="text-red-500 hover:text-red-700"
                  @click.stop="removeReviewFile(index)"
                >
                  <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </button>
              </div>
              
              <p class="text-sm text-brand-600 hover:text-brand-500 cursor-pointer">
                + Add more files
              </p>
            </div>
          </div>
          
          <input 
            id="reviewFileInput"
            ref="reviewFileInput"
            type="file" 
            class="hidden"
            accept=".pdf,.jpg,.jpeg,.png,.mp4,.mov,.avi"
            multiple
            @change="handleReviewFileChange"
          />
        </div>
        
        <div class="flex justify-end space-x-3">
          <Button 
            variant="secondary" 
            @click="showUploadReviewModal = false"
          >
            Cancel
          </Button>
          <Button 
            variant="primary" 
            :disabled="(reviewFiles.length === 0 && !reviewText) || isUploading || rating === 0"
            @click="uploadReviewFiles"
          >
            <span v-if="isUploading">Uploading...</span>
            <span v-else>Submit Review</span>
          </Button>
        </div>
      </div>
    </Modal>
  </AppLayout>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue'
import productService, { Product, ProductFilter } from '@/services/product.service'
import AppLayout from '@/components/layouts/AppLayout.vue'
import Button from '@/components/common/Button.vue'
import Input from '@/components/common/Input.vue'
import ProductCard from '@/components/common/ProductCard.vue'
import Modal from '@/components/common/Modal.vue'
import Badge from '@/components/common/Badge.vue'

// Define types for API responses
interface ApiResponse<T> {
  success?: boolean;
  message?: string;
  data?: T;
}

// State
const products = ref<Product[]>([])
const trendingProducts = ref<Product[]>([])
const categories = ref<string[]>(['Electronics', 'Books', 'Home & Kitchen', 'Toys & Games', 'Beauty']) // Default categories
const loading = ref(true)
const error = ref<string | null>(null)

// Product Details Modal
const showModal = ref(false)
const selectedProduct = ref<Product | null>(null)

// Upload Modals
const showUploadBookingModal = ref(false)
const showUploadReviewModal = ref(false)
const selectedProductId = ref<string | null>(null)

// File Upload refs
const bookingFileInput = ref<HTMLInputElement | null>(null)
const reviewFileInput = ref<HTMLInputElement | null>(null)

// File Upload state
const bookingFiles = ref<File[]>([])
const reviewFiles = ref<File[]>([])
const reviewText = ref('')
const rating = ref(0)
const isUploading = ref(false)

// Filters and Pagination
const searchQuery = ref('')
const selectedCategory = ref('')
const sortBy = ref('newest')
const currentPage = ref(1)
const limit = ref(12) // Products per page
const totalProducts = ref(0)

// Computed
const totalPages = computed(() => Math.max(1, Math.ceil(totalProducts.value / limit.value)))

// Combined products from both regular and trending products
const allProducts = computed(() => {
  // Combine both product arrays and remove duplicates
  const combined = [...products.value, ...trendingProducts.value]
  
  // Remove duplicates by ID
  const uniqueProducts = combined.reduce((acc: Product[], current) => {
    const exists = acc.some(p => p.id === current.id)
    if (!exists) {
      acc.push(current)
    }
    return acc
  }, [])
  
  return uniqueProducts
})

// Calculate booking percentage for a product
const getBookingPercentage = (product: Product): number => {
  return Math.min(Math.round((product.currentBookings / product.requiredBookings) * 100), 100)
}

// Methods
const fetchProducts = async () => {
  loading.value = true
  error.value = null
  
  try {
    // Only get active products for public view
    const filters: ProductFilter = {
      search: searchQuery.value,
      category: selectedCategory.value,
      sortBy: sortBy.value as 'newest' | 'popular',
      page: currentPage.value,
      limit: limit.value
    }
    
    const response = await productService.getAllProducts(filters)
    
    console.log('API Response:', response.data)
    
    // Extract products from the API response
    if (response.data) {
      let extractedProducts: Product[] = []
      let responseData = response.data as unknown as ApiResponse<any>
      
      // Check if it matches the API response format with success and data fields
      if (responseData && typeof responseData === 'object' && 'success' in responseData && 'data' in responseData) {
        // Handle nested data format
        const data = responseData.data
        
        if (Array.isArray(data)) {
          extractedProducts = data
        } else if (data && typeof data === 'object' && 'products' in data && Array.isArray(data.products)) {
          extractedProducts = data.products
          
          // If meta data is available for pagination
          if (data.meta && typeof data.meta === 'object' && 'total' in data.meta) {
            totalProducts.value = Number(data.meta.total)
          } else {
            totalProducts.value = extractedProducts.length * 3
          }
        } else {
          // Try to extract products from whatever format we're getting
          extractedProducts = extractProductsFromResponse(data)
          totalProducts.value = extractedProducts.length * 3
        }
      } else if (Array.isArray(response.data)) {
        // Direct array of products
        extractedProducts = response.data
        totalProducts.value = extractedProducts.length * 3
      } else {
        // If we can't determine the structure, try the generic extractor
        extractedProducts = extractProductsFromResponse(response.data)
        totalProducts.value = extractedProducts.length * 3
      }
      
      products.value = extractedProducts
    } else {
      products.value = []
      totalProducts.value = 0
    }
    
    if (products.value.length === 0 && currentPage.value > 1) {
      // If no products on current page and it's not the first page, go back to first page
      currentPage.value = 1
      await fetchProducts()
    }
  } catch (err: any) {
    error.value = err.message || 'Failed to fetch products'
    console.error('Error fetching products:', err)
    products.value = []
  } finally {
    loading.value = false
  }
}

// Helper function to extract products from different response formats
const extractProductsFromResponse = (data: any): Product[] => {
  // If it's a direct array of products
  if (Array.isArray(data)) {
    return data;
  }
  
  // If it has a products property that is an array
  if (data && typeof data === 'object' && 'products' in data && Array.isArray(data.products)) {
    return data.products;
  }
  
  // If it has a data property that is an array
  if (data && typeof data === 'object' && 'data' in data && Array.isArray(data.data)) {
    return data.data;
  }
  
  // If it's a nested structure like { data: { products: [...] } }
  if (data && typeof data === 'object' && 'data' in data && 
      typeof data.data === 'object' && data.data && 
      'products' in data.data && Array.isArray(data.data.products)) {
    return data.data.products;
  }
  
  // If the response is an object with product-like items
  if (typeof data === 'object' && data !== null) {
    // Check if the object itself is product-like (has expected properties)
    if ('id' in data && 'name' in data) {
      return [data];
    }
    
    // If it's an object with numeric keys (like an object-map of products)
    const possibleProducts = Object.values(data);
    if (possibleProducts.length > 0 && typeof possibleProducts[0] === 'object' && 
        possibleProducts[0] !== null && 'id' in (possibleProducts[0] as object)) {
      return possibleProducts as Product[];
    }
  }
  
  // If we can't extract products, return an empty array
  console.warn('Could not extract products from response:', data);
  return [];
}

// Also fetch trending products to add to the main product list
const fetchTrendingProducts = async () => {
  try {
    const response = await productService.getTrendingProducts(8) // Get trending products
    console.log('Trending API Response:', response.data)
    
    if (response.data) {
      let extractedProducts: Product[] = []
      let responseData = response.data as unknown as ApiResponse<any>
      
      // Check if it matches the API response format with success and data fields
      if (responseData && typeof responseData === 'object' && 'success' in responseData && 'data' in responseData) {
        // Access the products directly from the data property
        extractedProducts = Array.isArray(responseData.data) ? responseData.data : extractProductsFromResponse(responseData.data)
      } else {
        // Fallback to the generic extractor
        extractedProducts = extractProductsFromResponse(response.data)
      }
      
      trendingProducts.value = extractedProducts
    } else {
      trendingProducts.value = []
    }
  } catch (err) {
    console.error('Error fetching trending products:', err)
    trendingProducts.value = []
  }
}

const fetchCategories = async () => {
  // Using hardcoded fallback categories in case the API fails
  const fallbackCategories = ['Electronics', 'Books', 'Home & Kitchen', 'Toys & Games', 'Beauty']
  
  try {
    const response = await productService.getCategories()
    console.log('Categories API Response:', response.data)
    
    if (response.data) {
      let extractedCategories: string[] = []
      let responseData = response.data as unknown as ApiResponse<any>
      
      // Check if it matches the API response format with success and data fields
      if (responseData && typeof responseData === 'object' && 'success' in responseData && 'data' in responseData) {
        const data = responseData.data
        
        if (Array.isArray(data)) {
          extractedCategories = data
        } else if (data && typeof data === 'object' && 'categories' in data && Array.isArray(data.categories)) {
          extractedCategories = data.categories
        } else {
          // Try to extract categories from whatever format we're getting
          const possibleCategories = data && typeof data === 'object' 
            ? Object.values(data).filter(val => typeof val === 'string')
            : []
          
          if (possibleCategories.length > 0) {
            extractedCategories = possibleCategories as string[]
          } else {
            extractedCategories = fallbackCategories
          }
        }
      } else if (Array.isArray(response.data)) {
        extractedCategories = response.data
      } else {
        extractedCategories = fallbackCategories
      }
      
      categories.value = extractedCategories
    } else {
      categories.value = fallbackCategories
    }
  } catch (err) {
    // This will be hit since we're getting a 500 error
    console.error('Error fetching categories, using fallback:', err)
    categories.value = fallbackCategories
  }
}

// Upload modal handlers
const openUploadBookingModal = (productId: string) => {
  selectedProductId.value = productId
  bookingFiles.value = []
  showUploadBookingModal.value = true
}

const openUploadReviewModal = (productId: string) => {
  selectedProductId.value = productId
  reviewFiles.value = []
  reviewText.value = ''
  rating.value = 0
  showUploadReviewModal.value = true
}

// File input handlers
const triggerFileInput = (inputId: string) => {
  const input = document.getElementById(inputId)
  if (input) {
    input.click()
  }
}

const handleBookingFileChange = (event: Event) => {
  const input = event.target as HTMLInputElement
  if (input.files && input.files.length > 0) {
    const newFiles = Array.from(input.files)
    bookingFiles.value = [...bookingFiles.value, ...newFiles]
    // Reset input to allow selecting the same file again
    input.value = ''
  }
}

const handleReviewFileChange = (event: Event) => {
  const input = event.target as HTMLInputElement
  if (input.files && input.files.length > 0) {
    const newFiles = Array.from(input.files)
    reviewFiles.value = [...reviewFiles.value, ...newFiles]
    // Reset input to allow selecting the same file again
    input.value = ''
  }
}

const removeBookingFile = (index: number) => {
  bookingFiles.value.splice(index, 1)
}

const removeReviewFile = (index: number) => {
  reviewFiles.value.splice(index, 1)
}

// Upload functions
const uploadBookingFiles = async () => {
  if (!selectedProductId.value || bookingFiles.value.length === 0) return
  
  isUploading.value = true
  
  try {
    // Here you would normally call your API service
    // For example: await productService.uploadBookingDocuments(selectedProductId.value, bookingFiles.value)
    
    // Simulate API call for now
    await new Promise(resolve => setTimeout(resolve, 1500))
    
    // Show success message or notification
    alert('Booking documents uploaded successfully')
    
    // Close modal
    showUploadBookingModal.value = false
    bookingFiles.value = []
  } catch (err: any) {
    console.error('Error uploading booking files:', err)
    alert(`Upload failed: ${err.message || 'Unknown error'}`)
  } finally {
    isUploading.value = false
  }
}

const uploadReviewFiles = async () => {
  if (!selectedProductId.value || (reviewFiles.value.length === 0 && !reviewText.value) || rating.value === 0) return
  
  isUploading.value = true
  
  try {
    // Here you would normally call your API service
    // For example: await productService.uploadReview(selectedProductId.value, {
    //   text: reviewText.value,
    //   rating: rating.value,
    //   files: reviewFiles.value
    // })
    
    // Simulate API call for now
    await new Promise(resolve => setTimeout(resolve, 1500))
    
    // Show success message or notification
    alert('Review submitted successfully')
    
    // Close modal
    showUploadReviewModal.value = false
    reviewFiles.value = []
    reviewText.value = ''
    rating.value = 0
  } catch (err: any) {
    console.error('Error uploading review:', err)
    alert(`Upload failed: ${err.message || 'Unknown error'}`)
  } finally {
    isUploading.value = false
  }
}

const handleSearch = () => {
  // Reset to first page when searching
  currentPage.value = 1
  fetchProducts()
}

const applyFilters = () => {
  // Reset to first page when filtering
  currentPage.value = 1
  fetchProducts()
}

const resetFilters = () => {
  searchQuery.value = ''
  selectedCategory.value = ''
  sortBy.value = 'newest'
  currentPage.value = 1
  fetchProducts()
}

const prevPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--
    fetchProducts()
  }
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) {
    currentPage.value++
    fetchProducts()
  }
}

const goToPage = (page: number) => {
  currentPage.value = page
  fetchProducts()
}

// View product details - show modal instead of navigating
const viewProductDetails = (productId: string) => {
  const product = allProducts.value.find(p => p.id === productId)
  if (product) {
    selectedProduct.value = product
    showModal.value = true
  }
}

// Book a product - redirect to the product link
const bookProduct = (productId: string) => {
  const product = allProducts.value.find(p => p.id === productId)
  if (product && product.productLink) {
    // Close modal if open
    showModal.value = false
    
    // Open the product link in a new tab
    window.open(product.productLink, '_blank', 'noopener,noreferrer')
  }
}

// Watch for filter changes to automatically fetch products
watch([searchQuery], () => {
  // Add debounce for search input
  const debounceTimer = setTimeout(() => {
    if (currentPage.value === 1) {
      fetchProducts()
    } else {
      currentPage.value = 1 // This will trigger fetchProducts via the watcher
    }
  }, 300)
  
  return () => clearTimeout(debounceTimer)
}, { deep: true })

watch([selectedCategory, sortBy], () => {
  if (currentPage.value === 1) {
    fetchProducts()
  } else {
    currentPage.value = 1 // This will trigger fetchProducts via the watcher
  }
}, { deep: true })

watch(currentPage, () => {
  fetchProducts()
})

// Lifecycle
onMounted(() => {
  // Fetch products first
  fetchProducts()
  
  // Also fetch trending products to add to the main product list
  fetchTrendingProducts()
  
  // Then fetch categories
  fetchCategories()
})
</script>

<style scoped>
.animate-pulse-gold {
  animation: pulse-gold 1.5s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

@keyframes pulse-gold {
  0%, 100% {
    opacity: 1;
    transform: rotate(0deg);
  }
  50% {
    opacity: 0.5;
    transform: rotate(180deg);
  }
}

.form-input {
  display: block;
  width: 100%;
  padding-left: 0.75rem;
  padding-right: 0.75rem;
  padding-top: 0.5rem;
  padding-bottom: 0.5rem;
  border-width: 1px;
  border-color: #D1D5DB;
  border-radius: 0.375rem;
  box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
}

.form-input::placeholder {
  color: #9CA3AF;
}

.form-input:focus {
  outline: none;
  border-color: #6B7280;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.45);
}
</style>