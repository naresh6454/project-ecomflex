// ecomflex-frontend/src/pages/superadmin/Products.vue
<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
    <!-- Page Header -->
    <div class="pb-5 border-b border-gray-200 sm:flex sm:items-center sm:justify-between">
      <h1 class="text-2xl font-bold leading-6 text-gray-900">Products Management</h1>
      <div class="mt-3 sm:mt-0 sm:ml-4">
        <button
          type="button"
          class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-brand-600 hover:bg-brand-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500"
          @click="openAddProductModal"
        >
          <svg class="-ml-1 mr-2 h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z" clip-rule="evenodd" />
          </svg>
          Add Product
        </button>
      </div>
    </div>

    <!-- Search and Filters -->
    <div class="mt-6 pb-5 border-b border-gray-200">
      <div class="flex flex-col md:flex-row md:items-center md:space-x-4 space-y-4 md:space-y-0">
        <!-- Search -->
        <div class="flex-1">
          <div class="relative rounded-md shadow-sm">
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <svg class="h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" clip-rule="evenodd" />
              </svg>
            </div>
            <input
              type="text"
              v-model="searchQuery"
              @input="debounceSearch"
              class="focus:ring-brand-500 focus:border-brand-500 block w-full pl-10 sm:text-sm border-gray-300 rounded-md"
              placeholder="Search products..."
            />
          </div>
        </div>

        <!-- Filters -->
        <div class="flex space-x-2">
          <select 
            v-model="statusFilter" 
            @change="applyFilters"
            class="block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-brand-500 focus:border-brand-500 sm:text-sm rounded-md"
          >
            <option value="">All Status</option>
            <option value="active">Active</option>
            <option value="inactive">Inactive</option>
          </select>
          <select 
            v-model="sortBy" 
            @change="applyFilters"
            class="block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-brand-500 focus:border-brand-500 sm:text-sm rounded-md"
          >
            <option value="newest">Sort by: Newest</option>
            <option value="popular">Sort by: Bookings</option>
            <option value="name">Sort by: Name</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="productStore.loading" class="flex justify-center my-12">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-brand-700"></div>
    </div>

    <!-- Error State -->
    <div v-else-if="productStore.error" class="my-6 p-4 border rounded-md bg-red-50 border-red-200 text-red-600">
      <p>{{ productStore.error }}</p>
      <button 
        @click="loadProducts" 
        class="mt-2 text-sm underline focus:outline-none"
      >
        Try again
      </button>
    </div>

    <!-- Empty State -->
    <div v-else-if="productStore.products.length === 0" class="my-12 text-center">
      <svg class="mx-auto h-12 w-12 text-gray-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
      </svg>
      <h3 class="mt-2 text-sm font-medium text-gray-900">No products</h3>
      <p class="mt-1 text-sm text-gray-500">Get started by creating a new product.</p>
      <div class="mt-6">
        <button
          type="button"
          @click="openAddProductModal"
          class="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-brand-600 hover:bg-brand-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500"
        >
          <svg class="-ml-1 mr-2 h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z" clip-rule="evenodd" />
          </svg>
          New Product
        </button>
      </div>
    </div>

    <!-- Product Table -->
    <div v-else class="mt-6 overflow-hidden shadow ring-1 ring-black ring-opacity-5 md:rounded-lg">
      <table class="min-w-full divide-y divide-gray-300">
        <thead class="bg-gray-50">
          <tr>
            <th scope="col" class="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-6">Product</th>
            <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">ASIN</th>
            <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Founder</th>
            <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Bookings</th>
            <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Status</th>
            <th scope="col" class="relative py-3.5 pl-3 pr-4 sm:pr-6">
              <span class="sr-only">Actions</span>
            </th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200 bg-white">
          <tr v-for="product in productStore.products" :key="product.id" class="transition-colors hover:bg-gray-50">
            <td class="whitespace-nowrap py-4 pl-4 pr-3 text-sm sm:pl-6">
              <div class="flex items-center">
                <div class="h-10 w-10 flex-shrink-0">
                  <div v-if="product.image" class="h-10 w-10 rounded-lg overflow-hidden">
                    <img :src="product.image" :alt="product.name" class="h-10 w-10 object-cover" />
                  </div>
                  <div v-else class="h-10 w-10 rounded-lg bg-brand-100 flex items-center justify-center">
                    <span class="text-brand-600 font-medium">{{ product.name.charAt(0) }}</span>
                  </div>
                </div>
                <div class="ml-4">
                  <div class="font-medium text-gray-900">{{ product.name }}</div>
                  <div class="text-gray-500 truncate max-w-xs">{{ product.details.substring(0, 30) }}...</div>
                </div>
              </div>
            </td>
            <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
              {{ product.asin }}
            </td>
            <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
              {{ product.founder }}
            </td>
            <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
              {{ product.currentBookings }}/{{ product.requiredBookings }}
            </td>
            <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
              <span 
                class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full"
                :class="product.isActive ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'"
              >
                {{ product.isActive ? 'Active' : 'Inactive' }}
              </span>
            </td>
            <td class="relative whitespace-nowrap py-4 pl-3 pr-4 text-right text-sm font-medium sm:pr-6">
              <button 
                @click="editProduct(product)" 
                class="text-brand-600 hover:text-brand-900 transition-colors"
              >
                Edit
              </button>
              <span class="mx-1 text-gray-300">|</span>
              <button 
                @click="confirmDelete(product)" 
                class="text-red-600 hover:text-red-900 transition-colors"
              >
                Delete
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Pagination -->
    <div v-if="productStore.products.length > 0" class="mt-6 flex items-center justify-between">
      <div class="flex-1 flex justify-between sm:hidden">
        <button 
          @click="prevPage" 
          :disabled="!productStore.hasPreviousPage" 
          class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50"
          :class="{ 'opacity-50 cursor-not-allowed': !productStore.hasPreviousPage }"
        >
          Previous
        </button>
        <button 
          @click="nextPage" 
          :disabled="!productStore.hasNextPage" 
          class="ml-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50"
          :class="{ 'opacity-50 cursor-not-allowed': !productStore.hasNextPage }"
        >
          Next
        </button>
      </div>
      <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
        <div>
          <p class="text-sm text-gray-700">
            Showing <span class="font-medium">{{ paginationStart }}</span> to
            <span class="font-medium">{{ paginationEnd }}</span> of
            <span class="font-medium">{{ productStore.totalProducts }}</span> results
          </p>
        </div>
        <div>
          <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
            <button
              @click="prevPage"
              :disabled="!productStore.hasPreviousPage"
              class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
              :class="{ 'opacity-50 cursor-not-allowed': !productStore.hasPreviousPage }"
            >
              <span class="sr-only">Previous</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" />
              </svg>
            </button>
            <!-- Page Numbers -->
            <template v-for="page in displayedPages" :key="page">
              <button
                v-if="page !== '...'"
                @click="goToPage(Number(page))"
                class="relative inline-flex items-center px-4 py-2 border text-sm font-medium"
                :class="page === productStore.currentPage
                  ? 'z-10 bg-brand-50 border-brand-500 text-brand-600'
                  : 'bg-white border-gray-300 text-gray-700 hover:bg-gray-50'"
              >
                {{ page }}
              </button>
              <span
                v-else
                class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium text-gray-700"
              >
                ...
              </span>
            </template>
            <button
              @click="nextPage"
              :disabled="!productStore.hasNextPage"
              class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
              :class="{ 'opacity-50 cursor-not-allowed': !productStore.hasNextPage }"
            >
              <span class="sr-only">Next</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
              </svg>
            </button>
          </nav>
        </div>
      </div>
    </div>

    <!-- Product Form Modal -->
    <div v-if="showProductModal" class="fixed z-10 inset-0 overflow-y-auto" aria-labelledby="modal-title" role="dialog" aria-modal="true">
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <!-- Background overlay -->
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" @click="closeProductModal"></div>

        <!-- Modal panel -->
        <div class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full">
          <form @submit.prevent="submitProductForm">
            <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
              <div class="sm:flex sm:items-start">
                <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left w-full">
                  <h3 class="text-lg leading-6 font-medium text-gray-900" id="modal-title">
                    {{ isEditing ? 'Edit Product' : 'Add New Product' }}
                  </h3>
                  <div class="mt-4 space-y-4">
                    <div>
                      <label for="product-name" class="block text-sm font-medium text-gray-700">Product Name</label>
                      <input 
                        type="text" 
                        id="product-name" 
                        v-model="productForm.name"
                        class="mt-1 focus:ring-brand-500 focus:border-brand-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md" 
                        required
                      />
                    </div>
                    <div>
                      <label for="product-asin" class="block text-sm font-medium text-gray-700">ASIN</label>
                      <input 
                        type="text" 
                        id="product-asin" 
                        v-model="productForm.asin"
                        class="mt-1 focus:ring-brand-500 focus:border-brand-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md" 
                        required
                      />
                    </div>
                    <div>
                      <label for="product-founder" class="block text-sm font-medium text-gray-700">Founder</label>
                      <input 
                        type="text" 
                        id="product-founder" 
                        v-model="productForm.founder"
                        class="mt-1 focus:ring-brand-500 focus:border-brand-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md" 
                        required
                      />
                    </div>
                    <div>
                      <label for="product-link" class="block text-sm font-medium text-gray-700">Product Link</label>
                      <input 
                        type="url" 
                        id="product-link" 
                        v-model="productForm.productLink"
                        class="mt-1 focus:ring-brand-500 focus:border-brand-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md" 
                        required
                      />
                    </div>
                    <div>
                      <label for="product-details" class="block text-sm font-medium text-gray-700">Details</label>
                      <textarea 
                        id="product-details" 
                        v-model="productForm.details"
                        rows="3" 
                        class="mt-1 focus:ring-brand-500 focus:border-brand-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                      ></textarea>
                    </div>
                    <div>
                      <label for="product-bookings" class="block text-sm font-medium text-gray-700">Required Bookings</label>
                      <input 
                        type="number" 
                        id="product-bookings" 
                        v-model="productForm.requiredBookings"
                        min="1"
                        class="mt-1 focus:ring-brand-500 focus:border-brand-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md" 
                        required
                      />
                    </div>
                    <div class="flex items-center">
                      <input 
                        type="checkbox" 
                        id="product-active" 
                        v-model="productForm.isActive"
                        class="focus:ring-brand-500 h-4 w-4 text-brand-600 border-gray-300 rounded" 
                      />
                      <label for="product-active" class="ml-2 block text-sm text-gray-700">Active</label>
                    </div>
                    <div v-if="isEditing">
                      <label class="block text-sm font-medium text-gray-700">Product Image</label>
                      <div class="mt-1 flex items-center">
                        <div v-if="productForm.image || imagePreview" class="h-16 w-16 rounded-lg overflow-hidden bg-gray-100 mr-4">
                          <img :src="imagePreview || productForm.image" alt="Product image" class="h-16 w-16 object-cover" />
                        </div>
                        <div v-else class="h-16 w-16 rounded-lg bg-brand-100 flex items-center justify-center mr-4">
                          <svg class="h-6 w-6 text-brand-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                          </svg>
                        </div>
                        <label for="file-upload" class="relative cursor-pointer bg-white rounded-md font-medium text-brand-600 hover:text-brand-500 focus-within:outline-none">
                          <span>Upload a file</span>
                          <input id="file-upload" name="file-upload" type="file" class="sr-only" accept="image/*" @change="onFileChange" />
                        </label>
                      </div>
                    </div>
                    
                    <div v-if="formError" class="mt-2 text-sm text-red-600">
                      {{ formError }}
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
              <button 
                type="submit"
                :disabled="formSubmitting"
                class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-brand-600 text-base font-medium text-white hover:bg-brand-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500 sm:ml-3 sm:w-auto sm:text-sm"
                :class="{ 'opacity-75 cursor-not-allowed': formSubmitting }"
              >
                <svg v-if="formSubmitting" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                {{ formSubmitting ? 'Saving...' : 'Save' }}
              </button>
              <button 
                type="button" 
                @click="closeProductModal" 
                class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm"
              >
                Cancel
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <div v-if="showDeleteConfirm" class="fixed z-10 inset-0 overflow-y-auto" aria-labelledby="modal-title" role="dialog" aria-modal="true">
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <!-- Background overlay -->
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" @click="showDeleteConfirm = false"></div>

        <!-- Modal panel -->
        <div class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full">
          <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <div class="sm:flex sm:items-start">
              <div class="mx-auto flex-shrink-0 flex items-center justify-center h-12 w-12 rounded-full bg-red-100 sm:mx-0 sm:h-10 sm:w-10">
                <!-- Heroicon name: outline/exclamation -->
                <svg class="h-6 w-6 text-red-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                </svg>
              </div>
              <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                <h3 class="text-lg leading-6 font-medium text-gray-900" id="modal-title">
                  Delete Product
                </h3>
                <div class="mt-2">
                  <p class="text-sm text-gray-500">
                    Are you sure you want to delete this product? This action cannot be undone.
                  </p>
                  <p class="mt-2 font-medium">{{ productToDelete?.name }}</p>
                </div>
              </div>
            </div>
          </div>
          <div class="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
            <button 
              type="button" 
              @click="deleteProduct"
              :disabled="deleteSubmitting" 
              class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-red-600 text-base font-medium text-white hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 sm:ml-3 sm:w-auto sm:text-sm"
              :class="{ 'opacity-75 cursor-not-allowed': deleteSubmitting }"
            >
              <svg v-if="deleteSubmitting" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              {{ deleteSubmitting ? 'Deleting...' : 'Delete' }}
            </button>
            <button 
              type="button" 
              @click="showDeleteConfirm = false" 
              class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm"
            >
              Cancel
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useProductStore } from '@/stores/product'
import { Product } from '@/services/product.service'

// Store
const productStore = useProductStore()

// Search and filters
const searchQuery = ref('')
const statusFilter = ref('')
const sortBy = ref('newest')

// Modals
const showProductModal = ref(false)
const showDeleteConfirm = ref(false)

// Form
const isEditing = ref(false)
const productForm = ref({
  id: '',
  name: '',
  asin: '',
  founder: '',
  productLink: '',
  details: '',
  requiredBookings: 0,
  isActive: true,
  image: ''
})
const productToDelete = ref<Product | null>(null)
const formSubmitting = ref(false)
const deleteSubmitting = ref(false)
const formError = ref('')
const selectedFile = ref<File | null>(null)
const imagePreview = ref('')

// Pagination
const paginationStart = computed(() => {
  const start = (productStore.currentPage - 1) * productStore.pageSize + 1
  return Math.min(start, productStore.totalProducts)
})

const paginationEnd = computed(() => {
  const end = productStore.currentPage * productStore.pageSize
  return Math.min(end, productStore.totalProducts)
})

const displayedPages = computed(() => {
  const totalPages = productStore.totalPages
  const currentPage = productStore.currentPage
  
  // Always show first page, last page, current page, and pages +/- 1 from current
  const pages = new Set<number | string>([1, totalPages, currentPage])
  
  if (currentPage > 1) {
    pages.add(currentPage - 1)
  }
  
  if (currentPage < totalPages) {
    pages.add(currentPage + 1)
  }
  
  // Sort pages and insert ellipsis
  const result: (number | string)[] = Array.from(pages).sort((a, b) => Number(a) - Number(b))
  
  // Add ellipsis where needed
  const withEllipsis: (number | string)[] = []
  for (let i = 0; i < result.length; i++) {
    withEllipsis.push(result[i])
    if (i < result.length - 1 && Number(result[i+1]) - Number(result[i]) > 1) {
      withEllipsis.push('...')
    }
  }
  
  return withEllipsis
})

// Initial loading
onMounted(() => {
  loadProducts()
})

// Load products with current filters
const loadProducts = () => {
  const filter: any = {
    search: searchQuery.value || undefined,
    sortBy: sortBy.value || undefined
  }
  
  if (statusFilter.value) {
    filter.isActive = statusFilter.value === 'active'
  }
  
  productStore.fetchProducts(filter)
}

// Debounce search to avoid excessive API calls
let searchTimeout: number | null = null
const debounceSearch = () => {
  if (searchTimeout) {
    clearTimeout(searchTimeout)
  }
  
  searchTimeout = setTimeout(() => {
    productStore.currentPage = 1 // Reset to first page on search
    loadProducts()
  }, 300) as unknown as number
}

// Apply filters
const applyFilters = () => {
  productStore.currentPage = 1 // Reset to first page when filters change
  loadProducts()
}

// Pagination
const nextPage = () => {
  if (productStore.hasNextPage) {
    productStore.currentPage++
    loadProducts()
  }
}

const prevPage = () => {
  if (productStore.hasPreviousPage) {
    productStore.currentPage--
    loadProducts()
  }
}

const goToPage = (page: number) => {
  productStore.currentPage = page
  loadProducts()
}

// Modal functions
const openAddProductModal = () => {
  isEditing.value = false
  resetProductForm()
  showProductModal.value = true
}

const editProduct = (product: Product) => {
  isEditing.value = true
  productForm.value = {
    id: product.id,
    name: product.name,
    asin: product.asin,
    founder: product.founder,
    productLink: product.productLink,
    details: product.details,
    requiredBookings: product.requiredBookings,
    isActive: product.isActive,
    image: product.image || ''
  }
  showProductModal.value = true
}

const closeProductModal = () => {
  showProductModal.value = false
  resetProductForm()
}

const resetProductForm = () => {
  productForm.value = {
    id: '',
    name: '',
    asin: '',
    founder: '',
    productLink: '',
    details: '',
    requiredBookings: 0,
    isActive: true,
    image: ''
  }
  formError.value = ''
  selectedFile.value = null
  imagePreview.value = ''
}

// Handle file upload
const onFileChange = (event: Event) => {
  const input = event.target as HTMLInputElement
  if (input.files && input.files.length > 0) {
    selectedFile.value = input.files[0]
    
    // Create a preview
    const reader = new FileReader()
    reader.onload = (e) => {
      imagePreview.value = e.target?.result as string
    }
    reader.readAsDataURL(selectedFile.value)
  }
}

// Form submission
const submitProductForm = async () => {
  formSubmitting.value = true
  formError.value = ''
  
  try {
    let success = false
    
    if (isEditing.value) {
      // Update existing product
      success = await productStore.updateProduct(productForm.value.id, {
        name: productForm.value.name,
        asin: productForm.value.asin,
        founder: productForm.value.founder,
        productLink: productForm.value.productLink,
        details: productForm.value.details,
        requiredBookings: productForm.value.requiredBookings,
        isActive: productForm.value.isActive
      })
      
      // Handle image upload if new image is selected
      if (success && selectedFile.value) {
        const imageUrl = await productStore.uploadProductImage(productForm.value.id, selectedFile.value)
        if (!imageUrl) {
          formError.value = 'Product was updated but image upload failed'
        }
      }
    } else {
      // Create new product
      success = await productStore.createProduct({
        name: productForm.value.name,
        asin: productForm.value.asin,
        founder: productForm.value.founder,
        productLink: productForm.value.productLink,
        details: productForm.value.details,
        requiredBookings: productForm.value.requiredBookings,
        isActive: productForm.value.isActive
      })
    }
    
    if (success) {
      closeProductModal()
    } else {
      formError.value = productStore.error || 'An error occurred while saving the product'
    }
  } catch (err: any) {
    formError.value = err.message || 'An unexpected error occurred'
  } finally {
    formSubmitting.value = false
  }
}

// Delete product
const confirmDelete = (product: Product) => {
  productToDelete.value = product
  showDeleteConfirm.value = true
}

const deleteProduct = async () => {
  if (!productToDelete.value) return
  
  deleteSubmitting.value = true
  
  try {
    const success = await productStore.deleteProduct(productToDelete.value.id)
    if (success) {
      showDeleteConfirm.value = false
      productToDelete.value = null
    }
  } catch (err) {
    console.error('Error deleting product:', err)
  } finally {
    deleteSubmitting.value = false
  }
}
</script>