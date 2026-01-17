<template>
  <div class="min-h-screen bg-brand-50">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <h1 
        v-motion
        :initial="{ opacity: 0, y: -20 }"
        :enter="{ opacity: 1, y: 0 }"
        class="text-2xl font-bold text-gray-900 mb-6"
      >
        My Bookings
      </h1>
      
      <!-- Tabs -->
      <div 
        v-motion
        :initial="{ opacity: 0 }"
        :enter="{ opacity: 1, transition: { delay: 200 } }"
        class="border-b border-gray-200 mb-6"
      >
        <div class="flex -mb-px space-x-8">
          <button
            v-for="tab in tabs"
            :key="tab.value"
            @click="currentTab = tab.value"
            class="py-4 px-1 border-b-2 font-medium text-sm"
            :class="[
              currentTab === tab.value
                ? 'border-brand-600 text-brand-600'
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            ]"
          >
            {{ tab.label }}
            <span 
              v-if="tab.count" 
              class="ml-2 py-0.5 px-2 text-xs rounded-full"
              :class="[
                currentTab === tab.value
                  ? 'bg-brand-100 text-brand-800'
                  : 'bg-gray-100 text-gray-600'
              ]"
            >
              {{ tab.count }}
            </span>
          </button>
        </div>
      </div>
      
      <!-- Loading State -->
      <div v-if="loading" class="flex justify-center py-12">
        <div class="animate-pulse-gold inline-block h-12 w-12 rounded-full border-4 border-t-accent border-r-transparent border-b-transparent border-l-transparent"></div>
      </div>
      
      <!-- Empty State -->
      <div 
        v-else-if="filteredBookings.length === 0" 
        class="bg-white rounded-2xl shadow-sm p-8 text-center"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 mx-auto text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
        </svg>
        <h3 class="mt-4 text-lg font-medium text-gray-900">
          {{ getEmptyStateMessage() }}
        </h3>
        <p class="mt-2 text-gray-500">{{ getEmptyStateDescription() }}</p>
        <Button variant="primary" class="mt-6" @click="router.push('/')">
          Browse Products
        </Button>
      </div>
      
      <!-- Booking List -->
      <div v-else class="space-y-6">
        <div 
          v-for="(booking, index) in filteredBookings" 
          :key="booking.id"
          v-motion
          :initial="{ opacity: 0, y: 20 }"
          :enter="{ opacity: 1, y: 0, transition: { delay: index * 100 } }"
          class="bg-white rounded-2xl shadow-sm overflow-hidden"
        >
          <div class="md:flex">
            <!-- Status Bar (Left Side) -->
            <div class="md:w-1/6 p-4 md:p-6 md:border-r border-gray-100 flex flex-col items-center justify-center">
              <Badge 
                :variant="getStatusVariant(booking.status)" 
                size="lg"
                class="mb-2"
              >
                {{ formatStatus(booking.status) }}
              </Badge>
              <p class="text-xs text-center text-gray-500 mt-1">
                {{ formatDate(booking.createdAt) }}
              </p>
            </div>
            
            <!-- Booking Details (Middle) -->
            <div class="md:w-3/6 p-4 md:p-6 md:border-r border-gray-100">
              <div class="flex items-start">
                <!-- Product Image -->
                <div class="flex-shrink-0 h-16 w-16 bg-gray-100 rounded-lg mr-4">
                  <div class="h-full w-full flex items-center justify-center bg-brand-50">
                    <span class="text-brand-500 font-bold">{{ getProductInitial(booking) }}</span>
                  </div>
                </div>
                
                <!-- Product Info -->
                <div>
                  <h3 class="font-medium text-gray-900">
                    {{ getProductName(booking) }}
                  </h3>
                  <div class="mt-1 text-sm text-gray-500">
                    Booking ID: {{ booking.id.substring(0, 8) }}...
                  </div>
                  
                  <!-- Referral Code (if any) -->
                  <div v-if="booking.referralCode" class="mt-2 text-sm">
                    <span class="text-gray-500">Referral Code:</span>
                    <span class="font-medium text-brand-600 ml-1">{{ booking.referralCode }}</span>
                  </div>
                </div>
              </div>
            </div>
            
            <!-- Actions and Status (Right Side) -->
            <div class="md:w-2/6 p-4 md:p-6 bg-gray-50">
              <div class="h-full flex flex-col justify-between">
                <!-- Cashback Status -->
                <div>
                  <div class="text-sm text-gray-500 mb-1">Cashback Status</div>
                  <Badge 
                    :variant="booking.cashbackStatus === 'paid' ? 'green' : 'yellow'"
                    size="md"
                  >
                    {{ booking.cashbackStatus === 'paid' ? 'Paid' : 'Pending' }}
                  </Badge>
                </div>
                
                <!-- Action Buttons -->
                <div class="mt-4 space-y-2">
                  <!-- Upload Proof Button (for initiated bookings) -->
                  <Button 
                    v-if="booking.status === 'initiated'"
                    variant="primary" 
                    size="sm" 
                    block 
                    @click="showUploadProofModal(booking)"
                  >
                    Upload Proof
                  </Button>
                  
                  <!-- View Proof Button (if proof exists) -->
                  <Button 
                    v-if="booking.proofUrl"
                    variant="secondary" 
                    size="sm" 
                    block 
                    @click="viewProof(booking)"
                  >
                    View Proof
                  </Button>
                  
                  <!-- Reupload Proof Button (for rejected bookings) -->
                  <Button 
                    v-if="booking.status === 'rejected'"
                    variant="primary" 
                    size="sm" 
                    block 
                    @click="showUploadProofModal(booking)"
                  >
                    Re-upload Proof
                  </Button>
                  
                  <!-- Cancel Button (for initiated bookings) -->
                  <Button 
                    v-if="booking.status === 'initiated'"
                    variant="danger" 
                    size="sm" 
                    block 
                    @click="confirmCancel(booking)"
                  >
                    Cancel Booking
                  </Button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Upload Proof Modal -->
    <Modal
      v-if="selectedBooking"
      v-model="showProofModal"
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
        <Button variant="secondary" @click="showProofModal = false" class="mr-3">
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
    
    <!-- Confirm Cancel Modal -->
    <Modal
      v-if="selectedBooking"
      v-model="showCancelModal"
      title="Cancel Booking"
      description="Are you sure you want to cancel this booking? This action cannot be undone."
      size="md"
    >
      <template #footer>
        <Button variant="secondary" @click="showCancelModal = false" class="mr-3">
          No, Keep It
        </Button>
        <Button 
          variant="danger" 
          :loading="cancellingBooking"
          @click="cancelBooking"
        >
          Yes, Cancel Booking
        </Button>
      </template>
    </Modal>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useBookingStore, Booking } from '@/stores/booking'
import Button from '@/components/common/Button.vue'
import Badge from '@/components/common/Badge.vue'
import Modal from '@/components/common/Modal.vue'
import FileUpload from '@/components/common/FileUpload.vue'

// Router
const router = useRouter()

// Store
const bookingStore = useBookingStore()

// State
const loading = ref(true)
const currentTab = ref('all')
const proofFile = ref<File | null>(null)
const uploadingProof = ref(false)
const cancellingBooking = ref(false)

// Modals
const showProofModal = ref(false)
const showCancelModal = ref(false)
const selectedBooking = ref<Booking | null>(null)

// Tabs configuration
const tabs = computed(() => [
  { label: 'All Bookings', value: 'all', count: bookingStore.bookings.length },
  { label: 'Initiated', value: 'initiated', count: bookingStore.bookings.filter(b => b.status === 'initiated').length },
  { label: 'Pending', value: 'pending', count: bookingStore.bookings.filter(b => b.status === 'pending').length },
  { label: 'Approved', value: 'approved', count: bookingStore.bookings.filter(b => b.status === 'approved').length },
  { label: 'Rejected', value: 'rejected', count: bookingStore.bookings.filter(b => b.status === 'rejected').length }
])

// Filtered bookings based on current tab
const filteredBookings = computed(() => {
  if (currentTab.value === 'all') {
    return bookingStore.bookings
  }
  
  return bookingStore.bookings.filter(booking => booking.status === currentTab.value)
})

// Methods
const fetchBookings = async () => {
  loading.value = true
  
  try {
    await bookingStore.fetchUserBookings()
  } catch (error) {
    console.error('Error fetching bookings:', error)
  } finally {
    loading.value = false
  }
}

// Format status
const formatStatus = (status: Booking['status']): string => {
  switch (status) {
    case 'initiated':
      return 'Initiated'
    case 'pending':
      return 'Pending Verification'
    case 'approved':
      return 'Approved'
    case 'rejected':
      return 'Rejected'
    default:
      return 'Unknown'
  }
}

// Get status variant
const getStatusVariant = (status: Booking['status']): "blue" | "green" | "yellow" | "red" | "gold" | "gray" => {
  switch (status) {
    case 'initiated':
      return 'blue'
    case 'pending':
      return 'yellow'
    case 'approved':
      return 'green'
    case 'rejected':
      return 'red'
    default:
      return 'gray'
  }
}

// Format date
const formatDate = (dateString: string): string => {
  const date = new Date(dateString)
  return new Intl.DateTimeFormat('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  }).format(date)
}

// Get product initial
const getProductInitial = (booking: Booking): string => {
  // In a real app, you would have product details in the booking or fetch them
  // For demo purposes, just return a placeholder
  return 'P'
}

// Get product name
const getProductName = (booking: Booking): string => {
  // In a real app, you would have product details in the booking or fetch them
  // For demo purposes, just return a placeholder
  return `Product ${booking.productId.substring(0, 5)}`
}

// Get empty state message
const getEmptyStateMessage = (): string => {
  switch (currentTab.value) {
    case 'all':
      return 'You have no bookings yet'
    case 'initiated':
      return 'No initiated bookings'
    case 'pending':
      return 'No pending verifications'
    case 'approved':
      return 'No approved bookings yet'
    case 'rejected':
      return 'No rejected bookings'
    default:
      return 'No bookings found'
  }
}

// Get empty state description
const getEmptyStateDescription = (): string => {
  switch (currentTab.value) {
    case 'all':
      return 'Browse products and make your first booking'
    case 'initiated':
      return 'Bookings without uploaded proof will appear here'
    case 'pending':
      return 'Bookings awaiting verification will appear here'
    case 'approved':
      return 'Your approved bookings will appear here'
    case 'rejected':
      return 'Any rejected bookings will appear here'
    default:
      return 'Try adjusting your filters or make a new booking'
  }
}

// Show upload proof modal
const showUploadProofModal = (booking: Booking) => {
  selectedBooking.value = booking
  proofFile.value = null
  showProofModal.value = true
}

// Upload proof
const uploadProof = async () => {
  if (!selectedBooking.value || !proofFile.value) return
  
  uploadingProof.value = true
  
  try {
    await bookingStore.uploadProof(selectedBooking.value.id, proofFile.value)
    showProofModal.value = false
    proofFile.value = null
  } catch (error) {
    console.error('Error uploading proof:', error)
  } finally {
    uploadingProof.value = false
  }
}

// View proof
const viewProof = (booking: Booking) => {
  if (booking.proofUrl) {
    window.open(booking.proofUrl, '_blank')
  }
}

// Confirm cancel
const confirmCancel = (booking: Booking) => {
  selectedBooking.value = booking
  showCancelModal.value = true
}

// Cancel booking
const cancelBooking = async () => {
  if (!selectedBooking.value) return
  
  cancellingBooking.value = true
  
  try {
    await bookingStore.cancelBooking(selectedBooking.value.id)
    showCancelModal.value = false
  } catch (error) {
    console.error('Error cancelling booking:', error)
  } finally {
    cancellingBooking.value = false
  }
}

// Lifecycle
onMounted(() => {
  fetchBookings()
})

// Watch for tab changes to update the tab parameter in the URL
watch(currentTab, (newTab) => {
  router.replace({ query: { ...router.currentRoute.value.query, tab: newTab } })
})

// Initialize tab from URL query parameter
onMounted(() => {
  const tabParam = router.currentRoute.value.query.tab as string
  if (tabParam && tabs.value.some(tab => tab.value === tabParam)) {
    currentTab.value = tabParam
  }
})
</script>