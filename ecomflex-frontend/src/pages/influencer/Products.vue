<template>
  <div class="influencer-products">
    <!-- Page Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900 mb-2">Products to Promote</h1>
      <p class="text-gray-600">Browse and share products to earn referral rewards and cashback.</p>
    </div>

    <!-- Filters and Search -->
    <div class="bg-white rounded-xl shadow-lg p-6 mb-8">
      <div class="flex flex-col md:flex-row justify-between gap-4">
        <!-- Search Bar -->
        <div class="relative w-full md:w-80">
          <input 
            type="text" 
            placeholder="Search products..." 
            v-model="searchQuery"
            class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-accent/30 focus:border-accent"
          >
          <div class="absolute left-3 top-2.5 text-gray-400">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </div>
        </div>

        <!-- Filters -->
        <div class="flex flex-wrap gap-3">
          <select 
            v-model="categoryFilter"
            class="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-accent/30 focus:border-accent"
          >
            <option value="">All Categories</option>
            <option value="electronics">Electronics</option>
            <option value="home">Home & Kitchen</option>
            <option value="beauty">Beauty & Personal Care</option>
            <option value="fashion">Fashion</option>
            <option value="fitness">Fitness</option>
          </select>

          <select 
            v-model="sortBy"
            class="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-accent/30 focus:border-accent"
          >
            <option value="popularity">Most Popular</option>
            <option value="newest">Newest First</option>
            <option value="priceAsc">Price: Low to High</option>
            <option value="priceDesc">Price: High to Low</option>
            <option value="referrals">Most Referrals</option>
          </select>

          <button 
            @click="showTrendingOnly = !showTrendingOnly" 
            class="px-4 py-2 border rounded-lg flex items-center gap-2 transition-colors"
            :class="showTrendingOnly ? 'bg-accent text-white border-accent' : 'border-gray-300 text-gray-700 hover:bg-gray-50'"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
            </svg>
            Trending
          </button>
        </div>
      </div>
    </div>

    <!-- Product Stats Overview -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
      <!-- Total Products -->
      <div class="bg-white rounded-xl shadow-lg p-6 transition-all hover:shadow-xl">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-gray-600 text-sm">Products Available</p>
            <h3 class="text-3xl font-bold text-gray-900 mt-1">{{ stats.totalProducts }}</h3>
            <p class="text-sm mt-2">
              <span class="text-green-500 font-medium">+{{ stats.newProducts }}</span>
              <span class="text-gray-600 ml-1">this week</span>
            </p>
          </div>
          <div class="bg-accent/10 p-3 rounded-full">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-accent" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z" />
            </svg>
          </div>
        </div>
      </div>

      <!-- Top Product -->
      <div class="bg-white rounded-xl shadow-lg p-6 transition-all hover:shadow-xl">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-gray-600 text-sm">Top Performing Product</p>
            <h3 class="text-xl font-bold text-gray-900 mt-1 truncate">{{ stats.topProduct.name }}</h3>
            <p class="text-sm mt-2">
              <span class="text-green-500 font-medium">{{ stats.topProduct.referrals }}</span>
              <span class="text-gray-600 ml-1">referrals</span>
            </p>
          </div>
          <div class="bg-amber-100 p-3 rounded-full">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-amber-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z" />
            </svg>
          </div>
        </div>
      </div>

      <!-- Referral Conversion -->
      <div class="bg-white rounded-xl shadow-lg p-6 transition-all hover:shadow-xl">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-gray-600 text-sm">Avg. Conversion Rate</p>
            <h3 class="text-3xl font-bold text-gray-900 mt-1">{{ stats.conversionRate }}%</h3>
            <p class="text-sm mt-2">
              <span class="text-green-500 font-medium">+{{ stats.conversionTrend }}%</span>
              <span class="text-gray-600 ml-1">vs last month</span>
            </p>
          </div>
          <div class="bg-green-100 p-3 rounded-full">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-green-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- Products Grid -->
    <div class="mb-8">
      <div v-if="filteredProducts.length === 0" class="bg-white rounded-xl shadow-lg p-8 text-center">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 mx-auto text-gray-400 mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <h3 class="text-xl font-medium text-gray-700 mb-2">No products found</h3>
        <p class="text-gray-500">Try changing your search or filter criteria</p>
      </div>

      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <!-- Product Card -->
        <div 
          v-for="product in filteredProducts" 
          :key="product.id" 
          class="bg-white rounded-xl shadow-lg overflow-hidden transition-all hover:shadow-xl hover:scale-[1.02] duration-300"
        >
          <!-- Product Image with Price Badge -->
          <div class="relative">
            <img :src="product.image" :alt="product.name" class="w-full h-48 object-cover">
            <span class="absolute top-3 right-3 bg-accent text-white text-sm py-1 px-3 rounded-full font-medium">
              ${{ (product.price ?? 0).toFixed(2) }}
            </span>
            <span v-if="product.trending" class="absolute top-3 left-3 bg-gradient-to-r from-amber-500 to-red-500 text-white text-xs py-1 px-3 rounded-full font-medium flex items-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
              </svg>
              Trending
            </span>
          </div>

          <!-- Product Info -->
          <div class="p-5">
            <div class="flex items-start justify-between">
              <h3 class="text-lg font-bold text-gray-900 mb-1">{{ product.name }}</h3>
              <span class="bg-gray-100 text-gray-600 text-xs px-2 py-1 rounded">{{ product.category }}</span>
            </div>
            <p class="text-gray-600 text-sm mb-4 line-clamp-2">{{ product.description }}</p>
            
            <!-- Referral Stats -->
            <div class="flex justify-between items-center mb-4">
              <div class="flex items-center">
                <span class="text-sm text-gray-700 mr-4">
                  <span class="font-medium">{{ product.referrals }}</span> referrals
                </span>
                <span class="text-sm text-gray-700">
                  <span class="font-medium">{{ product.conversionRate }}%</span> conversion
                </span>
              </div>
              <div class="flex items-center text-sm">
                <span class="w-3 h-3 rounded-full mr-1" :class="product.inStock ? 'bg-green-500' : 'bg-red-500'"></span>
                {{ product.inStock ? 'In Stock' : 'Out of Stock' }}
              </div>
            </div>

            <!-- Action Buttons -->
            <div class="flex gap-2">
              <button 
                @click="showProductDetails(product)" 
                class="flex-1 px-3 py-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50 transition-colors duration-200 flex items-center justify-center"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                </svg>
                Details
              </button>
              <button 
                @click="shareProduct(product)" 
                class="flex-1 px-3 py-2 bg-accent text-white rounded-lg hover:bg-brand-dark transition-colors duration-200 flex items-center justify-center"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
                </svg>
                Share
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Pagination -->
    <div class="flex justify-center mb-8">
      <nav class="inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
        <button class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50">
          <span class="sr-only">Previous</span>
          <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
            <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" />
          </svg>
        </button>
        <button class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-accent text-sm font-medium text-white">
          1
        </button>
        <button class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium text-gray-700 hover:bg-gray-50">
          2
        </button>
        <button class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium text-gray-700 hover:bg-gray-50">
          3
        </button>
        <button class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50">
          <span class="sr-only">Next</span>
          <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
            <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
          </svg>
        </button>
      </nav>
    </div>

    <!-- Product Detail Modal -->
    <div v-if="selectedProduct" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
      <div class="bg-white rounded-xl shadow-2xl max-w-4xl w-full max-h-[90vh] overflow-y-auto" @click.stop>
        <div class="relative">
          <img :src="selectedProduct.image" :alt="selectedProduct.name" class="w-full h-64 object-cover rounded-t-xl">
          <button @click="selectedProduct = null" class="absolute top-4 right-4 bg-white rounded-full p-2 shadow-lg">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-gray-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
          <div class="absolute bottom-4 right-4 bg-accent text-white py-1 px-3 rounded-full font-medium">
            ${{ (selectedProduct?.price ?? 0).toFixed(2) }}
          </div>
        </div>

        <div class="p-6">
          <div class="flex flex-wrap gap-2 mb-4">
            <span class="bg-gray-100 text-gray-600 text-xs px-2 py-1 rounded">{{ selectedProduct.category }}</span>
            <span v-if="selectedProduct.trending" class="bg-gradient-to-r from-amber-500 to-red-500 text-white text-xs px-2 py-1 rounded flex items-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
              </svg>
              Trending
            </span>
            <span class="bg-accent/10 text-accent text-xs px-2 py-1 rounded flex items-center">
              <span class="font-medium">{{ selectedProduct.referrals }}</span> referrals
            </span>
            <span class="bg-green-100 text-green-600 text-xs px-2 py-1 rounded flex items-center">
              <span class="font-medium">{{ selectedProduct.conversionRate }}%</span> conversion
            </span>
          </div>

          <h2 class="text-2xl font-bold text-gray-900 mb-4">{{ selectedProduct.name }}</h2>
          
          <div class="mb-6">
            <h3 class="font-medium text-gray-700 mb-2">Description</h3>
            <p class="text-gray-600">{{ selectedProduct.fullDescription || selectedProduct.description }}</p>
          </div>

          <div class="mb-6">
            <h3 class="font-medium text-gray-700 mb-2">Referral Details</h3>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
              <div class="bg-gray-50 p-4 rounded-lg">
                <p class="text-sm text-gray-500 mb-1">Commission Rate</p>
                <p class="text-lg font-medium text-gray-800">{{ selectedProduct.commissionRate }}%</p>
              </div>
              <div class="bg-gray-50 p-4 rounded-lg">
                <p class="text-sm text-gray-500 mb-1">Cashback Amount</p>
                <p class="text-lg font-medium text-gray-800">${{ selectedProduct.cashbackAmount }}</p>
              </div>
              <div class="bg-gray-50 p-4 rounded-lg">
                <p class="text-sm text-gray-500 mb-1">Average Earnings</p>
                <p class="text-lg font-medium text-gray-800">${{ selectedProduct.avgEarnings }}</p>
              </div>
            </div>
          </div>

          <div class="border-t border-gray-200 pt-6">
            <h3 class="font-medium text-gray-700 mb-4">Share this Product</h3>
            
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">Your Referral Link</label>
              <div class="flex">
                <input 
                  type="text" 
                  :value="`${baseUrl}/product/${selectedProduct.id}?ref=${referralCode}`" 
                  readonly
                  class="flex-1 px-4 py-2 border border-gray-300 rounded-l-lg focus:outline-none focus:ring-2 focus:ring-accent/30 focus:border-accent"
                >
                <button 
                  @click="copyProductLink(selectedProduct)" 
                  class="px-4 py-2 bg-accent text-white rounded-r-lg hover:bg-brand-dark transition-colors duration-200 flex items-center"
                >
                  <span>{{ isCopied ? 'Copied!' : 'Copy' }}</span>
                  <svg v-if="!isCopied" xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 ml-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3" />
                  </svg>
                  <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 ml-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                  </svg>
                </button>
              </div>
            </div>
            
            <div class="flex flex-col sm:flex-row gap-3">
              <button 
                @click="shareProductViaWhatsApp(selectedProduct)" 
                class="flex-1 px-3 py-2 bg-green-500 text-white rounded-lg hover:bg-green-600 transition-colors duration-200 flex items-center justify-center"
              >
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="currentColor" class="mr-2" viewBox="0 0 16 16">
                  <path d="M13.601 2.326A7.854 7.854 0 0 0 7.994 0C3.627 0 .068 3.558.064 7.926c0 1.399.366 2.76 1.057 3.965L0 16l4.204-1.102a7.933 7.933 0 0 0 3.79.965h.004c4.368 0 7.926-3.558 7.93-7.93A7.898 7.898 0 0 0 13.6 2.326zM7.994 14.521a6.573 6.573 0 0 1-3.356-.92l-.24-.144-2.494.654.666-2.433-.156-.251a6.56 6.56 0 0 1-1.007-3.505c0-3.626 2.957-6.584 6.591-6.584a6.56 6.56 0 0 1 4.66 1.931 6.557 6.557 0 0 1 1.928 4.66c-.004 3.639-2.961 6.592-6.592 6.592zm3.615-4.934c-.197-.099-1.17-.578-1.353-.646-.182-.065-.315-.099-.445.099-.133.197-.513.646-.627.775-.114.133-.232.148-.43.05-.197-.1-.836-.308-1.592-.985-.59-.525-.985-1.175-1.103-1.372-.114-.198-.011-.304.088-.403.087-.088.197-.232.296-.346.1-.114.133-.198.198-.33.065-.134.034-.248-.015-.347-.05-.099-.445-1.076-.612-1.47-.16-.389-.323-.335-.445-.34-.114-.007-.247-.007-.38-.007a.729.729 0 0 0-.529.247c-.182.198-.691.677-.691 1.654 0 .977.71 1.916.81 2.049.098.133 1.394 2.132 3.383 2.992.47.205.84.326 1.129.418.475.152.904.129 1.246.08.38-.058 1.171-.48 1.338-.943.164-.464.164-.86.114-.943-.049-.084-.182-.133-.38-.232z"/>
                </svg>
                Share on WhatsApp
              </button>
              
              <button 
                @click="shareProductViaTwitter(selectedProduct)" 
                class="flex-1 px-3 py-2 bg-blue-400 text-white rounded-lg hover:bg-blue-500 transition-colors duration-200 flex items-center justify-center"
              >
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="currentColor" class="mr-2" viewBox="0 0 16 16">
                  <path d="M5.026 15c6.038 0 9.341-5.003 9.341-9.334 0-.14 0-.282-.006-.422A6.685 6.685 0 0 0 16 3.542a6.658 6.658 0 0 1-1.889.518 3.301 3.301 0 0 0 1.447-1.817 6.533 6.533 0 0 1-2.087.793A3.286 3.286 0 0 0 7.875 6.03a9.325 9.325 0 0 1-6.767-3.429 3.289 3.289 0 0 0 1.018 4.382A3.323 3.323 0 0 1 .64 6.575v.045a3.288 3.288 0 0 0 2.632 3.218 3.203 3.203 0 0 1-.865.115 3.23 3.23 0 0 1-.614-.057 3.283 3.283 0 0 0 3.067 2.277A6.588 6.588 0 0 1 .78 13.58a6.32 6.32 0 0 1-.78-.045A9.344 9.344 0 0 0 5.026 15z"/>
                </svg>
                Share on Twitter
              </button>
              
              <button 
                @click="shareProductViaEmail(selectedProduct)" 
                class="flex-1 px-3 py-2 bg-gray-600 text-white rounded-lg hover:bg-gray-700 transition-colors duration-200 flex items-center justify-center"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                </svg>
                Share via Email
              </button>
            </div>
          </div>

          <div class="mt-6 flex justify-between">
            <a 
              :href="selectedProduct.productUrl" 
              target="_blank" 
              rel="noopener noreferrer" 
              class="px-6 py-3 bg-gray-800 text-white rounded-lg hover:bg-gray-900 transition-colors duration-200 inline-flex items-center"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
              </svg>
              Visit Product Page
            </a>
            
            <button
              @click="selectedProduct = null" 
              class="px-6 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors duration-200"
            >
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
import ProductService from '@/services/product.service' // ✅ adjust this path if needed
import { useAuthStore } from '@/stores/auth'

// =======================
// 🧩 Extended Product Interface
// =======================
interface Product {
  id: string
  name: string
  asin: string
  founder: string
  productLink: string
  details: string
  requiredBookings: number
  currentBookings: number
  image?: string
  imageUrl?: string
  isActive: boolean
  createdAt: string
  updatedAt: string

  // 👇 Optional (used by UI)
  price?: number
  category?: string
  trending?: boolean
  description?: string
  fullDescription?: string
  referrals?: number
  conversionRate?: number
  inStock?: boolean
  commissionRate?: number
  cashbackAmount?: number
  avgEarnings?: number
  productUrl?: string
}

// =======================
// ⚙️ Setup & Stores
// =======================
const authStore = useAuthStore()
const referralCode = ref((authStore.user as any)?.referralCode || 'REF123')
const baseUrl = window.location.origin

// =======================
// 🧠 Reactive Variables
// =======================
const products = ref<Product[]>([])
const searchQuery = ref('')
const categoryFilter = ref('')
const sortBy = ref('newest')
const showTrendingOnly = ref(false)
const selectedProduct = ref<Product | null>(null)
const isCopied = ref(false)

// Example influencer stats (can be replaced with real backend data)
const stats = ref({
  totalProducts: 0,
  newProducts: 0,
  topProduct: { name: '', referrals: 0 },
  conversionRate: 0,
  conversionTrend: 0,
})

// =======================
// 🚀 Fetch Products
// =======================
const fetchProducts = async () => {
  try {
    const response = await ProductService.getAllProducts()
    // Handle the response structure {success: true, data: {products: [], meta: {}}}
    const productsData = response.data?.data?.products || response.data?.products || []
    products.value = productsData
    stats.value.totalProducts = productsData.length
  } catch (err) {
    console.error('Error fetching products:', err)
  }
}

// =======================
// 🔍 Filters and Sorting
// =======================
const filteredProducts = computed(() => {
  let result = products.value

  // Search
  if (searchQuery.value.trim()) {
    result = result.filter(
      (p) =>
        p.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
        p.details.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }

  // Category
  if (categoryFilter.value) {
    result = result.filter(
      (p) => p.category?.toLowerCase() === categoryFilter.value.toLowerCase()
    )
  }

  // Trending only
  if (showTrendingOnly.value) {
    result = result.filter((p) => p.trending)
  }

  // Sort
  switch (sortBy.value) {
    case 'newest':
      result = [...result].sort(
        (a, b) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime()
      )
      break
    case 'popular':
      result = [...result].sort((a, b) => (b.referrals || 0) - (a.referrals || 0))
      break
  }

  return result
})

// =======================
// 🪄 Modal + Sharing
// =======================
const showProductDetails = (product: Product) => {
  selectedProduct.value = product
}

const shareProduct = (product: Product) => {
  selectedProduct.value = product
}

const copyProductLink = (product: Product) => {
  const link = `${baseUrl}/product/${product.id}?ref=${referralCode.value}`
  navigator.clipboard.writeText(link)
  isCopied.value = true
  setTimeout(() => (isCopied.value = false), 2000)
}

const shareProductViaWhatsApp = (product: Product) => {
  const link = `${baseUrl}/product/${product.id}?ref=${referralCode.value}`
  const text = `Check out this product: ${product.name}! ${link}`
  window.open(`https://wa.me/?text=${encodeURIComponent(text)}`, '_blank')
}

const shareProductViaTwitter = (product: Product) => {
  const link = `${baseUrl}/product/${product.id}?ref=${referralCode.value}`
  const text = `Check out this product: ${product.name}!`
  window.open(
    `https://twitter.com/intent/tweet?text=${encodeURIComponent(text)}&url=${encodeURIComponent(
      link
    )}`,
    '_blank'
  )
}

const shareProductViaEmail = (product: Product) => {
  const link = `${baseUrl}/product/${product.id}?ref=${referralCode.value}`
  const subject = `Check out this product: ${product.name}`
  const body = `Hey, you might like this product:\n\n${link}`
  window.open(`mailto:?subject=${encodeURIComponent(subject)}&body=${encodeURIComponent(body)}`)
}

// =======================
// ⏱ Lifecycle
// =======================
onMounted(() => {
  fetchProducts()
})
</script>
