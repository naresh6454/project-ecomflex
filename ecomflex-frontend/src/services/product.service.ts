import api from './api'

// ===============================
// 🧩 Product Interface
// ===============================
export interface Product {
  id: string
  name: string
  asin: string
  founder: string
  productLink: string
  details: string
  requiredBookings: number
  currentBookings: number
  image?: string
  imageUrl?: string  // Matches backend
  isActive: boolean
  createdAt: string
  updatedAt: string
  // Optional UI fields
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

// ===============================
// 🎯 Product Filter Interface
// ===============================
export interface ProductFilter {
  category?: string
  search?: string
  sortBy?: 'newest' | 'popular' | 'price'
  page?: number
  limit?: number
}

// ===============================
// 🧠 ProductService Class
// ===============================
class ProductService {
  /**
   * 🛍️ Get all active products (Admin-created, visible to influencers)
   */
  async getAllProducts(filters?: ProductFilter) {
    // Filter out undefined values to prevent "undefined" string in query params
    const cleanedFilters = filters 
      ? Object.fromEntries(
          Object.entries(filters).filter(([_, value]) => value !== undefined)
        )
      : {}
    
    const queryString = Object.keys(cleanedFilters).length > 0 
      ? `?${new URLSearchParams(cleanedFilters as any).toString()}` 
      : ''
    const response = await api.get(`/products${queryString}`)
    return response
  }

  /**
   * 🔍 Get a single product by ID
   */
  async getProductById(productId: string) {
    const response = await api.get(`/products/${productId}`)
    return response
  }

  /**
   * 🔥 Get trending products
   */
  async getTrendingProducts(limit: number = 5) {
    const response = await api.get(`/products/trending?limit=${limit}`)
    return response
  }

  /**
   * 🏷️ Get product categories
   */
  async getCategories() {
    const response = await api.get('/products/categories')
    return response
  }

  /**
   * ✨ For Admin: Create a new product
   */
  async createProduct(
    productData: Omit<Product, 'id' | 'currentBookings' | 'createdAt' | 'updatedAt'>
  ) {
    const response = await api.post('/products', productData)
    return response
  }

  /**
   * ♻️ For Admin: Update product details
   */
  async updateProduct(productId: string, productData: Partial<Product>) {
    const response = await api.put(`/products/${productId}`, productData)
    return response
  }

  /**
   * 🚦 For Admin: Toggle product activation
   */
  async toggleProductStatus(productId: string, isActive: boolean) {
    const response = await api.put(`/products/${productId}/status`, { isActive })
    return response
  }

  /**
   * ❌ For Admin: Delete product
   */
  async deleteProduct(productId: string) {
    await api.delete(`/products/${productId}`)
  }

  /**
   * 🖼️ For Admin: Upload product image
   */
  async uploadProductImage(productId: string, imageFile: File) {
    const formData = new FormData()
    formData.append('image', imageFile)
    const response = await api.post(`/products/${productId}/image`, formData)
    return response
  }
}

export default new ProductService()
