<template>
  <div class="min-h-screen bg-brand-50">
    <!-- Product Detail Section -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- Back Button -->
      <button 
        @click="router.back()" 
        class="mb-6 inline-flex items-center text-sm font-medium text-brand-600 hover:text-brand-700"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" viewBox="0 0 20 20" fill="currentColor">
          <path fill-rule="evenodd" d="M9.707 16.707a1 1 0 01-1.414 0l-6-6a1 1 0 010-1.414l6-6a1 1 0 011.414 1.414L5.414 9H17a1 1 0 110 2H5.414l4.293 4.293a1 1 0 010 1.414z" clip-rule="evenodd" />
        </svg>
        Back to Products
      </button>
      
      <!-- Loading State -->
      <div v-if="loading" class="flex justify-center py-12">
        <div class="animate-pulse-gold inline-block h-12 w-12 rounded-full border-4 border-t-accent border-r-transparent border-b-transparent border-l-transparent"></div>
      </div>
      
      <!-- Error State -->
      <div v-else-if="error" class="bg-red-50 border-l-4 border-red-500 p-4 rounded-md">
        <div class="flex">
          <div class="flex-shrink-0">
            <svg class="h-5 w-5 text-red-500" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
              <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
            </svg>
          </div>
          <div class="ml-3">
            <p class="text-sm text-red-700">
              {{ error }}
            </p>
          </div>
        </div>
      </div>
      
      <!-- Product Details -->
      <div v-else-if="product" class="bg-white rounded-2xl shadow-lg overflow-hidden">
        <div class="md:flex">
          <!-- Product Image -->
          <div 
            v-motion
            :initial="{ opacity: 0, x: -20 }"
            :enter="{ opacity: 1, x: 0 }"
            class="md:w-1/2 bg-gray-100"
          >
            <div class="h-64 md:h-full w-full">
              <img
                v-if="product.image"
                :src="product.image"
                :alt="product.name"
                class="w-full h-full object-cover"
              />
              <div 
                v-else 
                class="w-full h-full flex items-center justify-center bg-gradient-to-br from-brand-100 to-brand-200"
              >
                <div class="text-center">
                  <span class="text-brand-600 font-bold text-5xl">{{ product.name.charAt(0) }}</span>
                  <div class="mt-2 text-brand-500">No image available</div>
                </div>
              </div>
            </div>
          </div>
          
          <!-- Product Info -->
          <div 
            v-motion
            :initial="{ opacity: 0, x: 20 }"
            :enter="{ opacity: 1, x: 0 }"
            class="md:w-1/2 p-6 md:p-8"
          >
            <div class="flex justify-between items-start">
              <h1 class="text-2xl md:text-3xl font-bold text-gray-900">{{ product.name }}</h1>
              <Badge 
                variant="gold" 
                size="lg"
                class="animate-bounce-subtle"
              >
                100% Cashback
              </Badge>
            </div>
            
            <div class="mt-4 grid grid-cols-2 gap-4">
              <div>
                <div class="text-sm text-gray-500">ASIN</div>
                <div class="font-medium">{{ product.asin }}</div>
              </div>
              <div>
                <div class="text-sm text-gray-500">Founder</div>
                <div class="font-medium">{{ product.founder }}</div>
              </div>
            </div>
            
            <div class="mt-6">
              <div class="text-sm text-gray-500 mb-1">Bookings</div>
              <div class="flex items-center">
                <span class="font-medium">{{ product.currentBookings }}/{{ product.requiredBookings }}</span>
                <span class="text-sm text-gray-500 ml-2">
                  ({{ Math.floor((product.currentBookings / product.requiredBookings) * 100) }}% booked)
                </span>
              </div>
              
              <div class="mt-2 w-full bg-gray-200 rounded-full h-2.5">
                <div 
                  class="bg-gradient-to-r from-brand-400 to-accent h-2.5 rounded-full" 
                  :style="`width: ${Math.min((product.currentBookings / product.requiredBookings) * 100, 100)}%`"
                ></div>
              </div>
            </div>
            
            <!-- Product Details -->
            <div class="mt-6">
              <h2 class="text-lg font-semibold text-gray-900">Details</h2>
              <p class="mt-2 text-gray-600 leading-relaxed">
                {{ product.details }}
              </p>
            </div>
            
            <!-- Referral Code Input -->
            <div class="mt-6">
              <label class="form-label" for="referral-code">Referral Code (Optional)</label>
              <Input
                id="referral-code"
                v-model="referralCode"
                placeholder="Enter referral code if you have one"
                helper-text="If you were referred by an influencer, enter their code here"
              >
                <template #left-icon>
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-gray-400" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z" clip-rule="evenodd" />
                  </svg>
                </template>
              </Input>
            </div>
            
            <!-- Action Buttons -->
            <div class="mt-8 space-y-4">
              <Button 
                variant="accent" 
                size="lg" 
                block 
                @click="openMerchantLink"
                ripple
                :left-icon="true"
              >
                <template #left-icon>
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path d="M11 3a1 1 0 100 2h2.586l-6.293 6.293a1 1 0 101.414 1.414L15 6.414V9a1 1 0 102 0V4a1 1 0 00-1-1h-5z" />
                    <path d="M5 5a2 2 0 00-2 2v8a2 2 0 002 2h8a2 2 0 002-2v-3a1 1 0 10-2 0v3H5V7h3a1 1 0 000-2H5z" />
                  </svg>
                </template>
                Book Now via Merchant
              </Button>
              
              <Button 
                variant="primary" 
                size="lg" 
                block 
                @click="showProofUploadModal = true"
                :left-icon="true"
              >
                <template #left-icon>
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M3 17a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM6.293 6.707a1 1 0 010-1.414l3-3a1 1 0 011.414 0l3 3a1 1 0 01-1.414 1.414L11 5.414V13a1 1 0 11-2 0V5.414L7.707 6.707a1 1 0 01-1.414 0z" clip-rule="evenodd" />
                  </svg>
                </template>
                Upload Purchase Proof
              </Button>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Upload Proof Modal -->
    <Modal
      v-model="showProofUploadModal"
      title="Upload Purchase Proof"
      description="Please upload a screenshot or PDF of your purchase invoice/receipt."
      size="lg"
    >
      <div class="space-y-6">
        <div>
          <label class="form-label" for="upload-proof">Your Proof</label>
          <FileUpload
            id="upload-proof"
            v-model="proofFile"
            accept=".jpg,.jpeg,.png,.pdf"
            :multiple="false"
            button-text="Upload proof"
            helper-text="Upload a clear image or PDF of your purchase receipt/invoice"
          />
        </div>
        
        <div>
          <label class="form-label" for="confirmation-referral-code">Referral Code (Optional)</label>
          <Input
            id="confirmation-referral-code"
            v-model="referralCode"
            placeholder="Enter referral code if you have one"
          />
        </div>
        
        <div class="bg-amber-50 border-l-4 border-amber-500 p-4 rounded-md">
          <div class="flex">
            <div class="flex-shrink-0">
              <svg class="h-5 w-5 text-amber-500" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
              </svg>
            </div>
            <div class="ml-3">
              <p class="text-sm text-amber-700">
                Make sure your proof clearly shows the product details, order number, and purchase date.
              </p>
            </div>
          </div>
        </div>
      </div>
      
      <template #footer>
        <Button variant="secondary" @click="showProofUploadModal = false" class="mr-3">
          Cancel
        </Button>
        <Button 
          variant="primary" 
          :loading="uploadingProof"
          :disabled="!proofFile" 
          @click="uploadProof"
        >
          Submit Proof
        </Button>
      </template>
    </Modal>
    
    <!-- Success Modal -->
    <Modal
      v-model="showSuccessModal"
      title="Proof Submitted Successfully"
      size="md"
    >
      <div class="text-center py-6">
        <div class="mx-auto flex items-center justify-center h-16 w-16 rounded-full bg-green-100">
          <svg class="h-10 w-10 text-green-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
          </svg>
        </div>
        <h3 class="mt-3 text-lg font-medium text-gray-900">Thank you!</h3>
        <p class="mt-2 text-gray-500">
          Your purchase proof has been submitted and is pending verification.
          You can track the status in "My Bookings" section.
        </p>
      </div>
      
      <template #footer>
        <Button variant="secondary" @click="showSuccessModal = false" class="mr-3">
          Close
        </Button>
        <Button variant="primary" @click="goToMyBookings">
          Go to My Bookings
        </Button>
      </template>
    </Modal>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import productService, { Product } from '@/services/product.service'
import bookingService from '@/services/booking.service'
import { useBookingStore } from '@/stores/booking'
import Button from '@/components/common/Button.vue'
import Badge from '@/components/common/Badge.vue'
import Input from '@/components/common/Input.vue'
import Modal from '@/components/common/Modal.vue'
import FileUpload from '@/components/common/FileUpload.vue'

// Router and Route
const router = useRouter()
const route = useRoute()
const bookingStore = useBookingStore()

// Product data
const product = ref<Product | null>(null)
const loading = ref(true)
const error = ref<string | null>(null)

// Form state
const referralCode = ref('')
const proofFile = ref<File | null>(null)
const uploadingProof = ref(false)

// Modals
const showProofUploadModal = ref(false)
const showSuccessModal = ref(false)

// Fetch product details
const fetchProductDetails = async () => {
  loading.value = true
  error.value = null
  
  try {
    const productId = route.params.id as string
    const response = await productService.getProductById(productId)
    product.value = response.data
  } catch (err: any) {
    error.value = err.message || 'Failed to fetch product details'
    console.error('Error fetching product details:', err)
  } finally {
    loading.value = false
  }
}

// Open merchant link
const openMerchantLink = () => {
  if (!product.value) return
  
  // First create a booking record
  createBooking()
  
  // Then open the merchant link in a new tab
  window.open(product.value.productLink, '_blank')
}

// Create a booking
const createBooking = async () => {
  if (!product.value) return
  
  try {
    await bookingStore.createBooking(
      product.value.id,
      referralCode.value || undefined
    )
  } catch (err) {
    console.error('Error creating booking:', err)
  }
}

// Upload proof
const uploadProof = async () => {
  if (!product.value || !proofFile.value) return
  
  uploadingProof.value = true
  
  try {
    // First create a booking if it doesn't exist
    const booking = await bookingStore.createBooking(
      product.value.id,
      referralCode.value || undefined
    )
    
    if (booking && proofFile.value) {
      // Then upload the proof
      await bookingStore.uploadProof(booking.id, proofFile.value)
      
      // Show success modal
      showProofUploadModal.value = false
      showSuccessModal.value = true
      
      // Reset form
      proofFile.value = null
    }
  } catch (err) {
    console.error('Error uploading proof:', err)
  } finally {
    uploadingProof.value = false
  }
}

// Go to My Bookings
const goToMyBookings = () => {
  showSuccessModal.value = false
  router.push('/my-bookings')
}

// Lifecycle
onMounted(() => {
  fetchProductDetails()
})
</script>