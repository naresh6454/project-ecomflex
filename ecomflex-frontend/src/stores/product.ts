// ecomflex-frontend/src/stores/product.ts
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import productService, { Product, ProductFilter } from '../services/product.service'

// Define response types to handle nested API responses
interface ApiResponse<T> {
  success?: boolean;
  message?: string;
  data?: T;
}

export const useProductStore = defineStore('product', () => {
  const products = ref<Product[]>([])
  const totalProducts = ref(0)
  const currentPage = ref(1)
  const pageSize = ref(10)
  const loading = ref(false)
  const error = ref<string | null>(null)
  const selectedProduct = ref<Product | null>(null)

  // Computed properties
  const totalPages = computed(() => Math.ceil(totalProducts.value / pageSize.value))
  const hasNextPage = computed(() => currentPage.value < totalPages.value)
  const hasPreviousPage = computed(() => currentPage.value > 1)

  // Actions
  const fetchProducts = async (filter?: ProductFilter) => {
    loading.value = true
    error.value = null
    
    try {
      const response = await productService.getAllProducts({
        page: currentPage.value,
        limit: pageSize.value,
        ...filter
      })
      
      console.log('Raw API response:', response)
      console.log('Response data:', response.data)
      
      // The backend consistently returns {success: true, data: {products: [], meta: {}}}
      if (response.data && response.data.success && response.data.data) {
        const responseData = response.data.data
        
        if (responseData.products && Array.isArray(responseData.products)) {
          console.log('Found products array:', responseData.products)
          products.value = responseData.products
          
          if (responseData.meta && typeof responseData.meta.total === 'number') {
            totalProducts.value = responseData.meta.total
          } else {
            totalProducts.value = responseData.products.length
          }
        } else {
          console.log('No products array found')
          products.value = []
          totalProducts.value = 0
        }
      } else {
        console.log('Unexpected response structure:', response.data)
        products.value = []
        totalProducts.value = 0
      }
    } catch (err: any) {
      console.error('Error fetching products:', err)
      error.value = err.message || 'Failed to fetch products'
      products.value = []
      totalProducts.value = 0
    } finally {
      loading.value = false
    }
  }

  const fetchProductById = async (id: string) => {
    loading.value = true
    error.value = null
    
    try {
      const response = await productService.getProductById(id)
      // Handle the response structure {success: true, data: product}
      if (response.data && response.data.success && response.data.data) {
        selectedProduct.value = response.data.data
      } else {
        selectedProduct.value = null
        console.error('Unexpected product response format:', response.data)
      }
    } catch (err: any) {
      console.error('Error fetching product details:', err)
      error.value = err.message || 'Failed to fetch product details'
      selectedProduct.value = null
    } finally {
      loading.value = false
    }
  }

  const createProduct = async (productData: any) => {
    loading.value = true
    error.value = null
    
    try {
      const response = await productService.createProduct(productData)
      console.log('Create product response:', response.data)
      
      // Check if the product was created successfully
      if (response.data && response.data.success) {
        // Refresh the product list
        await fetchProducts()
        return true
      } else {
        error.value = response.data?.message || 'Failed to create product'
        return false
      }
    } catch (err: any) {
      console.error('Error creating product:', err)
      error.value = err.message || 'Failed to create product'
      return false
    } finally {
      loading.value = false
    }
  }

  const updateProduct = async (id: string, productData: any) => {
    loading.value = true
    error.value = null
    
    try {
      await productService.updateProduct(id, productData)
      // Refresh the product list
      await fetchProducts()
      return true
    } catch (err: any) {
      console.error('Error updating product:', err)
      error.value = err.message || 'Failed to update product'
      return false
    } finally {
      loading.value = false
    }
  }

  const deleteProduct = async (id: string) => {
    loading.value = true
    error.value = null
    
    try {
      await productService.deleteProduct(id)
      // Refresh the product list
      await fetchProducts()
      return true
    } catch (err: any) {
      console.error('Error deleting product:', err)
      error.value = err.message || 'Failed to delete product'
      return false
    } finally {
      loading.value = false
    }
  }

  const toggleProductStatus = async (id: string, isActive: boolean) => {
    loading.value = true
    error.value = null
    
    try {
      await productService.toggleProductStatus(id, isActive)
      // Update the product in the local list
      const index = products.value.findIndex(p => p.id === id)
      if (index !== -1) {
        products.value[index].isActive = isActive
      }
      return true
    } catch (err: any) {
      console.error('Error toggling product status:', err)
      error.value = err.message || 'Failed to toggle product status'
      return false
    } finally {
      loading.value = false
    }
  }

  const uploadProductImage = async (id: string, imageFile: File) => {
    loading.value = true
    error.value = null
    
    try {
      const response = await productService.uploadProductImage(id, imageFile)
      
      // Handle possible wrapped response
      const responseData = response.data as ApiResponse<{imageUrl: string}> | {imageUrl: string}
      let imageUrl = '';
      
      if (typeof responseData === 'object' && 'success' in responseData && 'data' in responseData && responseData.data) {
        // It's a wrapped response
        imageUrl = responseData.data.imageUrl;
      } else if (typeof responseData === 'object' && 'imageUrl' in responseData) {
        // Direct response
        imageUrl = responseData.imageUrl;
      }
      
      // Update the product image in the local list
      const index = products.value.findIndex(p => p.id === id)
      if (index !== -1 && imageUrl) {
        products.value[index].image = imageUrl
      }
      return imageUrl
    } catch (err: any) {
      console.error('Error uploading product image:', err)
      error.value = err.message || 'Failed to upload product image'
      return null
    } finally {
      loading.value = false
    }
  }

  return {
    // State
    products,
    totalProducts,
    currentPage,
    pageSize,
    loading,
    error,
    selectedProduct,
    
    // Computed
    totalPages,
    hasNextPage,
    hasPreviousPage,
    
    // Actions
    fetchProducts,
    fetchProductById,
    createProduct,
    updateProduct,
    deleteProduct,
    toggleProductStatus,
    uploadProductImage
  }
})