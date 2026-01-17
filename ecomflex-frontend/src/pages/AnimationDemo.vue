<template>
  <div class="min-h-screen bg-brand-light pt-16 pb-12">
    <Navbar />
    
    <main class="container mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div class="max-w-4xl mx-auto">
        <div class="bg-white rounded-xl shadow-lg p-6 mb-8 dark:bg-gray-800">
          <h1 class="text-2xl font-bold text-gray-900 mb-6 dark:text-white">
            Animation Showcase
          </h1>
          
          <div class="space-y-12">
            <!-- Transition Demos -->
            <section>
              <h2 class="text-lg font-semibold text-gray-800 mb-4 dark:text-gray-200">
                Element Transitions
              </h2>
              
              <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
                <!-- Fade Transition -->
                <div class="p-4 bg-white rounded-lg shadow dark:bg-gray-700">
                  <h3 class="text-sm font-medium text-gray-700 mb-3 dark:text-gray-300">Fade Transition</h3>
                  
                  <div class="flex justify-center mb-4">
                    <button 
                      @click="toggleVisibility('fade')"
                      class="px-4 py-2 bg-accent text-white rounded-lg hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200"
                      v-ripple
                    >
                      {{ isVisible.fade ? 'Hide' : 'Show' }}
                    </button>
                  </div>
                  
                  <TransitionWrapper name="fade">
                    <div 
                      v-if="isVisible.fade" 
                      class="p-4 bg-gray-100 rounded-lg dark:bg-gray-800"
                    >
                      <p class="text-gray-700 dark:text-gray-300">This element fades in and out</p>
                    </div>
                  </TransitionWrapper>
                </div>
                
                <!-- Slide Transition -->
                <div class="p-4 bg-white rounded-lg shadow dark:bg-gray-700">
                  <h3 class="text-sm font-medium text-gray-700 mb-3 dark:text-gray-300">Slide Transition</h3>
                  
                  <div class="flex justify-center space-x-2 mb-4">
                    <button 
                      @click="toggleVisibility('slide')"
                      class="px-4 py-2 bg-accent text-white rounded-lg hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200"
                      v-ripple
                    >
                      {{ isVisible.slide ? 'Hide' : 'Show' }}
                    </button>
                    
                    <select
                      v-model="slideDirection"
                      class="rounded-md border-gray-300 shadow-sm focus:border-accent focus:ring focus:ring-accent focus:ring-opacity-50 dark:bg-gray-800 dark:border-gray-700 dark:text-white"
                    >
                      <option value="up">Slide Up</option>
                      <option value="down">Slide Down</option>
                      <option value="left">Slide Left</option>
                      <option value="right">Slide Right</option>
                    </select>
                  </div>
                  
                  <TransitionWrapper :name="`slide-${slideDirection}`">
                    <div 
                      v-if="isVisible.slide" 
                      class="p-4 bg-gray-100 rounded-lg dark:bg-gray-800"
                    >
                      <p class="text-gray-700 dark:text-gray-300">This element slides {{ slideDirection }}</p>
                    </div>
                  </TransitionWrapper>
                </div>
              </div>
              
              <div class="grid grid-cols-1 sm:grid-cols-2 gap-6 mt-6">
                <!-- Scale Transition -->
                <div class="p-4 bg-white rounded-lg shadow dark:bg-gray-700">
                  <h3 class="text-sm font-medium text-gray-700 mb-3 dark:text-gray-300">Scale Transition</h3>
                  
                  <div class="flex justify-center mb-4">
                    <button 
                      @click="toggleVisibility('scale')"
                      class="px-4 py-2 bg-accent text-white rounded-lg hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200"
                      v-ripple
                    >
                      {{ isVisible.scale ? 'Hide' : 'Show' }}
                    </button>
                  </div>
                  
                  <TransitionWrapper name="scale">
                    <div 
                      v-if="isVisible.scale" 
                      class="p-4 bg-gray-100 rounded-lg dark:bg-gray-800"
                    >
                      <p class="text-gray-700 dark:text-gray-300">This element scales in and out</p>
                    </div>
                  </TransitionWrapper>
                </div>
                
                <!-- Bounce Transition -->
                <div class="p-4 bg-white rounded-lg shadow dark:bg-gray-700">
                  <h3 class="text-sm font-medium text-gray-700 mb-3 dark:text-gray-300">Bounce Transition</h3>
                  
                  <div class="flex justify-center mb-4">
                    <button 
                      @click="toggleVisibility('bounce')"
                      class="px-4 py-2 bg-accent text-white rounded-lg hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200"
                      v-ripple
                    >
                      {{ isVisible.bounce ? 'Hide' : 'Show' }}
                    </button>
                  </div>
                  
                  <TransitionWrapper name="bounce">
                    <div 
                      v-if="isVisible.bounce" 
                      class="p-4 bg-gray-100 rounded-lg dark:bg-gray-800"
                    >
                      <p class="text-gray-700 dark:text-gray-300">This element bounces in and out</p>
                    </div>
                  </TransitionWrapper>
                </div>
              </div>
            </section>
            
            <!-- List Transitions -->
            <section>
              <h2 class="text-lg font-semibold text-gray-800 mb-4 dark:text-gray-200">
                List Transitions
              </h2>
              
              <div class="p-4 bg-white rounded-lg shadow dark:bg-gray-700">
                <h3 class="text-sm font-medium text-gray-700 mb-3 dark:text-gray-300">Staggered List Items</h3>
                
                <div class="flex justify-between mb-4">
                  <button 
                    @click="addItem"
                    class="px-4 py-2 bg-accent text-white rounded-lg hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200"
                    v-ripple
                  >
                    Add Item
                  </button>
                  
                  <button 
                    @click="removeItem"
                    class="px-4 py-2 bg-white border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-300 dark:hover:bg-gray-700"
                    v-ripple="{ color: 'rgba(0, 0, 0, 0.1)' }"
                  >
                    Remove Item
                  </button>
                  
                  <button 
                    @click="shuffleItems"
                    class="px-4 py-2 bg-white border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-300 dark:hover:bg-gray-700"
                    v-ripple="{ color: 'rgba(0, 0, 0, 0.1)' }"
                  >
                    Shuffle
                  </button>
                </div>
                
                <TransitionWrapper name="list" tag="ul" wrapperClass="space-y-2" :stagger-children="true" :stagger="50">
                  <li 
                    v-for="item in listItems" 
                    :key="item.id"
                    class="p-3 bg-gray-100 rounded-lg flex justify-between items-center dark:bg-gray-800"
                  >
                    <span class="text-gray-700 dark:text-gray-300">{{ item.text }}</span>
                    <button
                      @click="removeSpecificItem(item.id)"
                      class="text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200"
                    >
                      <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
                      </svg>
                    </button>
                  </li>
                </TransitionWrapper>
              </div>
            </section>
            
            <!-- Modal Demo -->
            <section>
              <h2 class="text-lg font-semibold text-gray-800 mb-4 dark:text-gray-200">
                Modal Animations
              </h2>
              
              <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
                <button
                  @click="openModal('fade')"
                  class="px-4 py-2 bg-accent text-white rounded-lg hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200"
                  v-ripple
                >
                  Fade Modal
                </button>
                
                <button
                  @click="openModal('zoom')"
                  class="px-4 py-2 bg-accent text-white rounded-lg hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200"
                  v-ripple
                >
                  Zoom Modal
                </button>
                
                <button
                  @click="openModal('slideUp')"
                  class="px-4 py-2 bg-accent text-white rounded-lg hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200"
                  v-ripple
                >
                  Slide Up Modal
                </button>
              </div>
            </section>
            
            <!-- Skeleton Loaders -->
            <section>
              <h2 class="text-lg font-semibold text-gray-800 mb-4 dark:text-gray-200">
                Skeleton Loaders
              </h2>
              
              <div class="flex items-center justify-between mb-4">
                <h3 class="text-sm font-medium text-gray-700 dark:text-gray-300">Loading States</h3>
                
                <div class="flex items-center space-x-2">
                  <span class="text-sm text-gray-500 dark:text-gray-400">Toggle Loading:</span>
                  <button
                    @click="isLoading = !isLoading"
                    class="px-4 py-2 bg-accent text-white rounded-lg hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200"
                    v-ripple
                  >
                    {{ isLoading ? 'Show Content' : 'Show Skeletons' }}
                  </button>
                </div>
              </div>
              
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <!-- Card Skeleton -->
                <div v-if="isLoading">
                  <SkeletonCard />
                </div>
                <div v-else class="rounded-xl bg-white shadow-md overflow-hidden dark:bg-gray-800">
                  <img src="https://placehold.co/600x400/e2e8f0/1e293b?text=Product+Image" alt="Product" class="w-full h-48 object-cover" />
                  <div class="p-4 space-y-3">
                    <h3 class="text-lg font-semibold text-gray-900 dark:text-white">Premium Ecommerce Product</h3>
                    <p class="text-sm text-gray-500 dark:text-gray-400">Premium category</p>
                    <p class="text-gray-700 dark:text-gray-300">
                      This is a high-quality product available with 100% cashback. Book now through our platform.
                    </p>
                    <div class="flex justify-between items-center pt-2">
                      <button 
                        class="px-4 py-2 bg-accent text-white rounded-lg hover:bg-accent-dark"
                        v-ripple
                      >
                        Book Now
                      </button>
                      <div class="h-10 w-10 rounded-full bg-gray-200 flex items-center justify-center dark:bg-gray-700">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
                        </svg>
                      </div>
                    </div>
                  </div>
                </div>
                
                <!-- List Skeleton -->
                <div class="rounded-xl bg-white shadow-md p-4 dark:bg-gray-800">
                  <h3 class="text-md font-semibold text-gray-900 mb-4 dark:text-white">Recent Bookings</h3>
                  
                  <div v-if="isLoading" class="space-y-3">
                    <div v-for="i in 5" :key="i" class="flex items-center space-x-3">
                      <SkeletonLoader variant="avatar" width="10" height="10" />
                      <div class="flex-1 space-y-2">
                        <SkeletonLoader variant="text" width="3/4" height="4" />
                        <SkeletonLoader variant="text" width="1/2" height="3" color="light" />
                      </div>
                      <SkeletonLoader variant="text" width="20" height="6" color="accent" radius="full" />
                    </div>
                  </div>
                  
                  <ul v-else class="space-y-3">
                    <li v-for="i in 5" :key="i" class="flex items-center space-x-3 p-2 hover:bg-gray-50 rounded-lg dark:hover:bg-gray-700">
                      <div class="h-10 w-10 rounded-full bg-gray-200 flex items-center justify-center dark:bg-gray-700">
                        <span class="text-sm font-medium text-gray-600 dark:text-gray-300">
                          {{ ['AB', 'CD', 'EF', 'GH', 'IJ'][i - 1] }}
                        </span>
                      </div>
                      <div class="flex-1">
                        <p class="text-sm font-medium text-gray-900 dark:text-white">
                          {{ ['Product One', 'Product Two', 'Product Three', 'Product Four', 'Product Five'][i - 1] }}
                        </p>
                        <p class="text-xs text-gray-500 dark:text-gray-400">
                          {{ ['2 hours ago', '1 day ago', '3 days ago', '1 week ago', '2 weeks ago'][i - 1] }}
                        </p>
                      </div>
                      <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium" :class="[
                        i % 2 === 0 ? 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-300' : 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-300'
                      ]">
                        {{ i % 2 === 0 ? 'Approved' : 'Pending' }}
                      </span>
                    </li>
                  </ul>
                </div>
              </div>
            </section>
          </div>
        </div>
      </div>
    </main>
    
    <!-- Modals -->
    <Modal
      v-model="modalOpen.fade"
      title="Fade Modal"
      animation-variant="fade"
      with-golden-accent
    >
      <p class="text-gray-700 mb-4 dark:text-gray-300">
        This modal uses a simple fade transition.
      </p>
      
      <div class="p-4 bg-gray-100 rounded-lg dark:bg-gray-700">
        <p class="text-gray-700 dark:text-gray-300">Modal content can contain any elements or components.</p>
      </div>
    </Modal>
    
    <Modal
      v-model="modalOpen.zoom"
      title="Zoom Modal"
      animation-variant="zoom"
      size="lg"
      with-golden-accent
    >
      <p class="text-gray-700 mb-4 dark:text-gray-300">
        This modal uses a zoom transition with scale animation.
      </p>
      
      <div class="p-4 bg-gray-100 rounded-lg dark:bg-gray-700">
        <p class="text-gray-700 dark:text-gray-300">Modal content can contain any elements or components.</p>
      </div>
    </Modal>
    
    <Modal
      v-model="modalOpen.slideUp"
      title="Slide Up Modal"
      animation-variant="slide-up"
      position="bottom"
      size="xl"
      with-golden-accent
    >
      <p class="text-gray-700 mb-4 dark:text-gray-300">
        This modal slides up from the bottom of the screen.
      </p>
      
      <div class="p-4 bg-gray-100 rounded-lg dark:bg-gray-700">
        <p class="text-gray-700 dark:text-gray-300">Modal content can contain any elements or components.</p>
      </div>
    </Modal>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import Navbar from '@/components/layouts/Navbar.vue'
import TransitionWrapper from '@/components/ui/TransitionWrapper.vue'
import Modal from '@/components/ui/Modal.vue'
import SkeletonLoader from '@/components/ui/SkeletonLoader.vue'
import SkeletonCard from '@/components/ui/SkeletonCard.vue'
import { useNotificationStore } from '@/stores/notification'

// Store
const notificationStore = useNotificationStore()

// Element visibility state
const isVisible = reactive({
  fade: false,
  slide: false,
  scale: false,
  bounce: false
})

const slideDirection = ref('up')

// Toggle visibility for demo elements
const toggleVisibility = (key: keyof typeof isVisible) => {
  isVisible[key] = !isVisible[key]
}

// List demo
const listItems = ref<Array<{ id: number, text: string }>>([
  { id: 1, text: 'Item 1' },
  { id: 2, text: 'Item 2' },
  { id: 3, text: 'Item 3' }
])
let nextId = 4

// Add a new item to the list
const addItem = () => {
  const id = nextId++
  listItems.value.push({ id, text: `Item ${id}` })
}

// Remove the last item from the list
const removeItem = () => {
  if (listItems.value.length > 0) {
    listItems.value.pop()
  }
}

// Remove a specific item by id
const removeSpecificItem = (id: number) => {
  const index = listItems.value.findIndex(item => item.id === id)
  if (index !== -1) {
    listItems.value.splice(index, 1)
  }
}

// Shuffle the list items
const shuffleItems = () => {
  listItems.value = [...listItems.value].sort(() => Math.random() - 0.5)
}

// Modal demo
const modalOpen = reactive({
  fade: false,
  zoom: false,
  slideUp: false
})

// Open a specific modal
const openModal = (type: keyof typeof modalOpen) => {
  modalOpen[type] = true
}

// Skeleton demo
const isLoading = ref(true)

// Initialize demo
onMounted(() => {
  // Show welcome notification
  setTimeout(() => {
    notificationStore.showToast({
      type: 'info',
      title: 'Animation Demo',
      message: 'Explore the various animations and transitions available'
    })
  }, 500)
  
  // Auto toggle loading state after 3 seconds
  setTimeout(() => {
    isLoading.value = false
  }, 3000)
})
</script>