<template>
  <div class="min-h-screen flex flex-col bg-brand-50">
    <!-- Header/Navbar -->
    <header class="bg-white shadow-sm">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <!-- Logo and Navigation -->
          <div class="flex">
            <!-- Logo -->
            <div class="flex-shrink-0 flex items-center">
              <router-link to="/" class="flex items-center">
                <div class="h-8 w-8 bg-gradient-to-r from-brand-500 to-brand-600 rounded-xl flex items-center justify-center">
                  <span class="text-white font-bold">E</span>
                </div>
                <span class="ml-2 text-xl font-bold text-brand-600">Ecomflex</span>
              </router-link>
            </div>
            
            <!-- Desktop Navigation -->
            <nav class="hidden md:ml-8 md:flex md:space-x-8">
              <router-link 
                v-for="item in navigationItems" 
                :key="item.name" 
                :to="item.to" 
                class="inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
                :class="[
                  isActive(item.to) 
                    ? 'border-brand-600 text-brand-600' 
                    : 'border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700'
                ]"
              >
                {{ item.name }}
              </router-link>
            </nav>
          </div>
          
          <!-- User Menu -->
          <div class="flex items-center">
            <div class="hidden md:ml-4 md:flex md:items-center">
              <!-- Notifications -->
              <button
                type="button"
                class="p-1 rounded-full text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500"
              >
                <span class="sr-only">View notifications</span>
                <svg
                  class="h-6 w-6"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                  aria-hidden="true"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"
                  />
                </svg>
              </button>
              
              <!-- Profile Dropdown -->
              <div class="ml-4 relative">
                <div>
                  <button
                    @click="isUserMenuOpen = !isUserMenuOpen"
                    type="button"
                    class="bg-white rounded-full flex text-sm focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500"
                    id="user-menu-button"
                    aria-expanded="false"
                    aria-haspopup="true"
                  >
                    <span class="sr-only">Open user menu</span>
                    <div v-if="user?.avatar" class="h-8 w-8 rounded-full overflow-hidden">
                      <img :src="user.avatar" alt="User avatar" class="h-full w-full object-cover" />
                    </div>
                    <div v-else class="h-8 w-8 rounded-full bg-brand-100 flex items-center justify-center">
                      <span class="text-brand-600 font-medium">{{ userInitials }}</span>
                    </div>
                  </button>
                </div>

                <!-- Dropdown Menu -->
                <transition
                  enter-active-class="transition ease-out duration-100"
                  enter-from-class="transform opacity-0 scale-95"
                  enter-to-class="transform opacity-100 scale-100"
                  leave-active-class="transition ease-in duration-75"
                  leave-from-class="transform opacity-100 scale-100"
                  leave-to-class="transform opacity-0 scale-95"
                >
                  <div
                    v-if="isUserMenuOpen"
                    class="origin-top-right absolute right-0 mt-2 w-48 rounded-xl shadow-lg bg-white ring-1 ring-black ring-opacity-5 py-1 z-50"
                    role="menu"
                    aria-orientation="vertical"
                    aria-labelledby="user-menu-button"
                  >
                    <div class="px-4 py-2 border-b border-gray-100">
                      <p class="text-sm font-medium text-gray-900">{{ user?.name || 'User' }}</p>
                      <p class="text-xs text-gray-500 truncate">{{ user?.email || 'user@example.com' }}</p>
                    </div>
                    <router-link
                      to="/profile"
                      class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                      role="menuitem"
                      @click="isUserMenuOpen = false"
                    >
                      Your Profile
                    </router-link>
                    <router-link
                      to="/settings"
                      class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                      role="menuitem"
                      @click="isUserMenuOpen = false"
                    >
                      Settings
                    </router-link>
                    <button
                      @click="logout"
                      class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                      role="menuitem"
                    >
                      Sign out
                    </button>
                  </div>
                </transition>
              </div>
            </div>
            
            <!-- Mobile menu button -->
            <div class="flex items-center md:hidden">
              <button
                @click="isMobileMenuOpen = !isMobileMenuOpen"
                type="button"
                class="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-brand-500"
                aria-controls="mobile-menu"
                :aria-expanded="isMobileMenuOpen"
              >
                <span class="sr-only">Open main menu</span>
                <!-- Icon when menu is closed -->
                <svg
                  v-if="!isMobileMenuOpen"
                  class="block h-6 w-6"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                  aria-hidden="true"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M4 6h16M4 12h16M4 18h16"
                  />
                </svg>
                <!-- Icon when menu is open -->
                <svg
                  v-else
                  class="block h-6 w-6"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                  aria-hidden="true"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M6 18L18 6M6 6l12 12"
                  />
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Mobile menu -->
      <transition
        enter-active-class="transition ease-out duration-100"
        enter-from-class="transform opacity-0 scale-95"
        enter-to-class="transform opacity-100 scale-100"
        leave-active-class="transition ease-in duration-75"
        leave-from-class="transform opacity-100 scale-100"
        leave-to-class="transform opacity-0 scale-95"
      >
        <div v-if="isMobileMenuOpen" class="md:hidden" id="mobile-menu">
          <div class="pt-2 pb-3 space-y-1">
            <router-link
              v-for="item in navigationItems"
              :key="item.name"
              :to="item.to"
              class="block pl-3 pr-4 py-2 border-l-4 text-base font-medium"
              :class="[
                isActive(item.to)
                  ? 'border-brand-600 text-brand-700 bg-brand-50'
                  : 'border-transparent text-gray-600 hover:bg-gray-50 hover:border-gray-300 hover:text-gray-800'
              ]"
              @click="isMobileMenuOpen = false"
            >
              {{ item.name }}
            </router-link>
          </div>
          <div class="pt-4 pb-3 border-t border-gray-200">
            <div class="flex items-center px-4">
              <div class="flex-shrink-0">
                <div v-if="user?.avatar" class="h-10 w-10 rounded-full overflow-hidden">
                  <img :src="user.avatar" alt="User avatar" class="h-full w-full object-cover" />
                </div>
                <div v-else class="h-10 w-10 rounded-full bg-brand-100 flex items-center justify-center">
                  <span class="text-brand-600 font-medium">{{ userInitials }}</span>
                </div>
              </div>
              <div class="ml-3">
                <div class="text-base font-medium text-gray-800">{{ user?.name || 'User' }}</div>
                <div class="text-sm font-medium text-gray-500">{{ user?.email || 'user@example.com' }}</div>
              </div>
              <button
                type="button"
                class="ml-auto flex-shrink-0 p-1 rounded-full text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500"
              >
                <span class="sr-only">View notifications</span>
                <svg
                  class="h-6 w-6"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                  aria-hidden="true"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"
                  />
                </svg>
              </button>
            </div>
            <div class="mt-3 space-y-1">
              <router-link
                to="/profile"
                class="block px-4 py-2 text-base font-medium text-gray-500 hover:text-gray-800 hover:bg-gray-100"
                @click="isMobileMenuOpen = false"
              >
                Your Profile
              </router-link>
              <router-link
                to="/settings"
                class="block px-4 py-2 text-base font-medium text-gray-500 hover:text-gray-800 hover:bg-gray-100"
                @click="isMobileMenuOpen = false"
              >
                Settings
              </router-link>
              <button
                @click="logout"
                class="block w-full text-left px-4 py-2 text-base font-medium text-gray-500 hover:text-gray-800 hover:bg-gray-100"
              >
                Sign out
              </button>
            </div>
          </div>
        </div>
      </transition>
    </header>

    <!-- Main Content -->
    <main class="flex-grow">
      <slot></slot>
    </main>

    <!-- Footer -->
    <footer class="bg-white border-t border-gray-200">
      <div class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
        <div class="flex flex-col md:flex-row justify-between items-center">
          <div class="flex items-center">
            <div class="h-6 w-6 bg-gradient-to-r from-brand-500 to-brand-600 rounded-lg flex items-center justify-center">
              <span class="text-white font-bold text-xs">E</span>
            </div>
            <span class="ml-2 text-sm font-semibold text-gray-700">Ecomflex</span>
          </div>
          <div class="mt-4 md:mt-0 text-center md:text-right">
            <p class="text-sm text-gray-500">
              &copy; {{ new Date().getFullYear() }} Ecomflex. All rights reserved.
            </p>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'

// Route and Router
const route = useRoute()
const router = useRouter()

// Navigation state
const isMobileMenuOpen = ref(false)
const isUserMenuOpen = ref(false)

// Mock user data (replace with actual auth store)
const user = ref({
  name: 'Demo User',
  email: 'user@example.com',
  avatar: null,
  role: 'public'
})

// Navigation items
const navigationItems = [
  { name: 'Products', to: '/home' },
  { name: 'My Bookings', to: '/my-bookings' }
]

// Computed
const userInitials = computed(() => {
  if (!user.value?.name) return 'U'
  
  const nameParts = user.value.name.split(' ')
  if (nameParts.length === 1) return nameParts[0].charAt(0).toUpperCase()
  
  return (nameParts[0].charAt(0) + nameParts[nameParts.length - 1].charAt(0)).toUpperCase()
})

// Methods
const isActive = (path: string): boolean => {
  if (path === '/home' && route.path === '/') return true
  return route.path.startsWith(path)
}

const logout = () => {
  // Clear auth tokens
  localStorage.removeItem('token')
  localStorage.removeItem('refreshToken')
  localStorage.removeItem('userRole')
  
  // Redirect to login
  router.push('/auth/login')
}

// Close mobile menu on route change
watch(() => route.path, () => {
  isMobileMenuOpen.value = false
  isUserMenuOpen.value = false
})

// Close menus when clicking outside
const handleClickOutside = (event: MouseEvent) => {
  const target = event.target as HTMLElement
  
  // Check if click is outside user menu
  if (isUserMenuOpen.value && !target.closest('#user-menu-button')) {
    isUserMenuOpen.value = false
  }
}

// Lifecycle hooks
onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onMounted(() => {
  return () => {
    document.removeEventListener('click', handleClickOutside)
  }
})
</script>