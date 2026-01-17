<template>
  <header 
    class="fixed top-0 left-0 right-0 z-40 transition-all duration-300"
    :class="[scrolled ? 'bg-white/90 dark:bg-gray-900/90 backdrop-blur-sm shadow-md py-2' : 'bg-transparent py-4']"
  >
    <div class="container mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center">
        <!-- Logo -->
        <router-link 
          to="/" 
          class="flex items-center space-x-2 focus:outline-none focus-visible:ring-2 focus-visible:ring-accent rounded-md"
          aria-label="Ecomflex Home"
        >
          <img src="/logo.svg" alt="Ecomflex Logo" class="h-10 w-auto" />
          <span class="text-xl font-bold text-gray-900 dark:text-white">Ecomflex</span>
        </router-link>

        <!-- Desktop Navigation -->
        <nav class="hidden md:flex items-center space-x-6">
          <template v-if="isAuthenticated">
            <router-link 
              v-for="item in getNavigationItems()" 
              :key="item.path" 
              :to="item.path"
              class="text-gray-600 hover:text-accent font-medium text-sm transition-colors duration-200 dark:text-gray-300 dark:hover:text-accent"
              :class="{ 'text-accent font-semibold': isActiveRoute(item.path) }"
            >
              {{ item.label }}
            </router-link>
          </template>
          <template v-else>
            <router-link 
              to="/auth/login" 
              class="text-gray-600 hover:text-accent font-medium text-sm transition-colors duration-200 dark:text-gray-300 dark:hover:text-accent"
              :class="{ 'text-accent font-semibold': isActiveRoute('/auth/login') }"
            >
              Login
            </router-link>
            <router-link 
              to="/auth/register" 
              class="text-gray-600 hover:text-accent font-medium text-sm transition-colors duration-200 dark:text-gray-300 dark:hover:text-accent"
              :class="{ 'text-accent font-semibold': isActiveRoute('/auth/register') }"
            >
              Register
            </router-link>
          </template>
        </nav>

        <!-- Right-side Actions -->
        <div class="flex items-center space-x-4">
          <!-- Notification Bell -->
          <NotificationDropdown v-if="isAuthenticated" />
          
          <!-- User Menu -->
          <div v-if="isAuthenticated" class="relative" ref="userMenuContainer">
            <button
              @click="toggleUserMenu"
              class="flex items-center focus:outline-none focus-visible:ring-2 focus-visible:ring-accent rounded-full"
              id="user-menu-button"
              aria-expanded="false"
              aria-haspopup="true"
            >
              <span class="sr-only">Open user menu</span>
              <div class="h-8 w-8 rounded-full bg-gray-200 flex items-center justify-center overflow-hidden border border-gray-300 dark:bg-gray-700 dark:border-gray-600">
                <img 
                  v-if="userAvatar" 
                  :src="userAvatar" 
                  alt="User avatar" 
                  class="h-full w-full object-cover"
                />
                <span v-else class="text-sm font-medium text-gray-600 dark:text-gray-300">
                  {{ userInitials }}
                </span>
              </div>
            </button>

            <!-- User Dropdown Menu -->
            <transition
              enter-active-class="transition ease-out duration-200"
              enter-from-class="transform opacity-0 scale-95"
              enter-to-class="transform opacity-100 scale-100"
              leave-active-class="transition ease-in duration-150"
              leave-from-class="transform opacity-100 scale-100"
              leave-to-class="transform opacity-0 scale-95"
            >
              <div
                v-if="userMenuOpen"
                class="origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5 focus:outline-none dark:bg-gray-800 dark:ring-gray-700"
                role="menu"
                aria-orientation="vertical"
                aria-labelledby="user-menu-button"
                tabindex="-1"
              >
                <div class="px-4 py-3 border-b border-gray-100 dark:border-gray-700">
                  <p class="text-sm text-gray-700 dark:text-gray-300">Signed in as</p>
                  <p class="text-sm font-medium text-gray-900 truncate dark:text-white">{{ userEmail }}</p>
                </div>
                
                <router-link
                  v-for="item in userMenuItems"
                  :key="item.label"
                  :to="item.path"
                  class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 dark:text-gray-300 dark:hover:bg-gray-700"
                  @click="userMenuOpen = false"
                  role="menuitem"
                >
                  {{ item.label }}
                </router-link>
                
                <button
                  class="block w-full text-left px-4 py-2 text-sm text-red-600 hover:bg-gray-100 dark:hover:bg-gray-700"
                  @click="handleLogout"
                  role="menuitem"
                >
                  Sign out
                </button>
              </div>
            </transition>
          </div>

          <!-- Login/Register Buttons (Mobile) -->
          <div v-if="!isAuthenticated" class="md:hidden flex items-center space-x-2">
            <router-link
              to="/auth/login"
              class="inline-flex items-center px-3 py-1.5 border border-transparent text-xs font-medium rounded-md text-gray-700 bg-gray-100 hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-accent dark:bg-gray-800 dark:text-gray-300 dark:hover:bg-gray-700"
            >
              Login
            </router-link>
            
            <router-link
              to="/auth/register"
              class="inline-flex items-center px-3 py-1.5 border border-transparent text-xs font-medium rounded-md text-white bg-accent hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-accent"
            >
              Register
            </router-link>
          </div>

          <!-- Mobile Menu Button -->
          <button
            @click="toggleMobileMenu"
            type="button"
            class="md:hidden inline-flex items-center justify-center p-2 rounded-md text-gray-600 hover:text-gray-900 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-accent dark:text-gray-300 dark:hover:text-white dark:hover:bg-gray-700"
            aria-controls="mobile-menu"
            :aria-expanded="mobileMenuOpen"
          >
            <span class="sr-only">{{ mobileMenuOpen ? 'Close menu' : 'Open menu' }}</span>
            <svg
              v-if="!mobileMenuOpen"
              class="block h-6 w-6"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              aria-hidden="true"
            >
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
            <svg
              v-else
              class="block h-6 w-6"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              aria-hidden="true"
            >
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Mobile Menu -->
    <transition
      enter-active-class="transition ease-out duration-200"
      enter-from-class="opacity-0 -translate-y-4"
      enter-to-class="opacity-100 translate-y-0"
      leave-active-class="transition ease-in duration-150"
      leave-from-class="opacity-100 translate-y-0"
      leave-to-class="opacity-0 -translate-y-4"
    >
      <div v-if="mobileMenuOpen" class="md:hidden bg-white shadow-lg dark:bg-gray-900" id="mobile-menu">
        <div class="px-2 pt-2 pb-3 space-y-1">
          <template v-if="isAuthenticated">
            <router-link
              v-for="item in getNavigationItems()"
              :key="item.path"
              :to="item.path"
              class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:bg-gray-100 hover:text-accent dark:text-gray-300 dark:hover:bg-gray-700"
              :class="{ 'bg-gray-100 text-accent dark:bg-gray-700': isActiveRoute(item.path) }"
              @click="mobileMenuOpen = false"
            >
              {{ item.label }}
            </router-link>
          </template>
          <template v-else>
            <router-link
              to="/auth/login"
              class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:bg-gray-100 hover:text-accent dark:text-gray-300 dark:hover:bg-gray-700"
              :class="{ 'bg-gray-100 text-accent dark:bg-gray-700': isActiveRoute('/auth/login') }"
              @click="mobileMenuOpen = false"
            >
              Login
            </router-link>
            <router-link
              to="/auth/register"
              class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 hover:bg-gray-100 hover:text-accent dark:text-gray-300 dark:hover:bg-gray-700"
              :class="{ 'bg-gray-100 text-accent dark:bg-gray-700': isActiveRoute('/auth/register') }"
              @click="mobileMenuOpen = false"
            >
              Register
            </router-link>
          </template>
        </div>
      </div>
    </transition>
  </header>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import NotificationDropdown from '@/components/ui/NotificationDropdown.vue'
import { useNotificationStore } from '@/stores/notification'
import { useAnimation } from '@/composables/useAnimation'

// Store and router setup
const router = useRouter()
const route = useRoute()
const notificationStore = useNotificationStore()
const { buttonRipple } = useAnimation()

// State
const scrolled = ref(false)
const mobileMenuOpen = ref(false)
const userMenuOpen = ref(false)
const userMenuContainer = ref<HTMLElement | null>(null)

// Mock user data - would come from a user store in a real app
const isAuthenticated = computed(() => localStorage.getItem('token') !== null)
const userAvatar = computed(() => localStorage.getItem('userAvatar') || '')
const userName = computed(() => localStorage.getItem('userName') || 'User')
const userEmail = computed(() => localStorage.getItem('userEmail') || 'user@example.com')
const userRole = computed(() => localStorage.getItem('userRole') || 'public')
const userInitials = computed(() => {
  const name = userName.value
  if (!name) return 'U'
  return name.split(' ').map(n => n[0]).join('').toUpperCase()
})

// User menu items based on role
const userMenuItems = computed(() => {
  const items = []
  
  if (userRole.value === 'admin') {
    items.push(
      { label: 'Dashboard', path: '/admin/dashboard' },
      { label: 'Profile', path: '/admin/profile' },
      { label: 'Settings', path: '/admin/settings' }
    )
  } else if (userRole.value === 'influencer') {
    items.push(
      { label: 'Dashboard', path: '/influencer/dashboard' },
      { label: 'Profile', path: '/influencer/profile' },
      { label: 'Settings', path: '/influencer/settings' }
    )
  } else {
    items.push(
      { label: 'My Bookings', path: '/my-bookings' },
      { label: 'Profile', path: '/profile' },
      { label: 'Settings', path: '/settings' }
    )
  }
  
  return items
})

// Navigation items based on user role
const getNavigationItems = () => {
  if (userRole.value === 'admin') {
    return [
      { label: 'Dashboard', path: '/admin/dashboard' },
      { label: 'Products', path: '/admin/products' },
      { label: 'Proofs', path: '/admin/proofs' },
      { label: 'Influencers', path: '/admin/influencers' }
    ]
  } else if (userRole.value === 'influencer') {
    return [
      { label: 'Dashboard', path: '/influencer/dashboard' },
      { label: 'Referrals', path: '/influencer/referrals' },
      { label: 'Products', path: '/influencer/products' },
      { label: 'Earnings', path: '/influencer/earnings' }
    ]
  } else {
    return [
      { label: 'Home', path: '/home' },
      { label: 'My Bookings', path: '/my-bookings' }
    ]
  }
}

// Active route check
const isActiveRoute = (path: string) => {
  if (path === '/home' && route.path === '/') {
    return true
  }
  return route.path.startsWith(path)
}

// Handlers
const toggleMobileMenu = () => {
  mobileMenuOpen.value = !mobileMenuOpen.value
}

const toggleUserMenu = () => {
  userMenuOpen.value = !userMenuOpen.value
}

const handleLogout = () => {
  // Clear user data
  localStorage.removeItem('token')
  localStorage.removeItem('userRole')
  localStorage.removeItem('userEmail')
  localStorage.removeItem('userName')
  localStorage.removeItem('userAvatar')
  
  // Show success notification
  notificationStore.showToast({
    type: 'success',
    message: 'You have been logged out successfully'
  })
  
  // Redirect to login
  router.push('/auth/login')
  
  // Close menu
  userMenuOpen.value = false
}

// Handle click outside to close user menu
const handleClickOutside = (event: MouseEvent) => {
  if (userMenuContainer.value && !userMenuContainer.value.contains(event.target as Node)) {
    userMenuOpen.value = false
  }
}

// Handle scroll for navbar background
const handleScroll = () => {
  scrolled.value = window.scrollY > 10
}

// Add ripple effect to all buttons
const addRippleEffect = () => {
  const buttons = document.querySelectorAll('button:not([data-no-ripple]), .btn')
  buttons.forEach(button => {
    button.addEventListener('click', (e) => buttonRipple(e as MouseEvent))
  })
}

// Lifecycle hooks
onMounted(() => {
  document.addEventListener('click', handleClickOutside)
  window.addEventListener('scroll', handleScroll)
  handleScroll() // Check initial scroll position
  addRippleEffect()
  
  // Show welcome notification for new users (demo)
  if (isAuthenticated.value && !localStorage.getItem('welcomed')) {
    setTimeout(() => {
      notificationStore.showToast({
        type: 'info',
        title: 'Welcome back!',
        message: 'Check out the latest products with cashback offers'
      })
      localStorage.setItem('welcomed', 'true')
    }, 1000)
  }
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
  window.removeEventListener('scroll', handleScroll)
})

// Close mobile menu on route change
watch(() => route.path, () => {
  mobileMenuOpen.value = false
})
</script>

<style scoped>
/* Add any component-specific styles here */
.router-link-exact-active {
  @apply text-accent font-semibold;
}
</style>