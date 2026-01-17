<template>
  <div class="h-screen flex overflow-hidden bg-gray-100">
    <!-- Off-canvas menu for mobile, show/hide based on off-canvas menu state -->
    <div 
      v-show="sidebarOpen" 
      class="fixed inset-0 flex z-40 md:hidden"
      role="dialog"
      aria-modal="true"
    >
      <!-- Off-canvas menu overlay, show/hide based on off-canvas menu state -->
      <div 
        v-show="sidebarOpen" 
        class="fixed inset-0 bg-gray-600 bg-opacity-75"
        @click="sidebarOpen = false"
        aria-hidden="true"
      ></div>

      <!-- Off-canvas menu, show/hide based on off-canvas menu state -->
      <div 
        v-show="sidebarOpen" 
        class="relative flex-1 flex flex-col max-w-xs w-full pt-5 pb-4 bg-brand-800"
      >
        <!-- Close button, show/hide based on off-canvas menu state -->
        <div class="absolute top-0 right-0 -mr-12 pt-2">
          <button 
            @click="sidebarOpen = false"
            class="ml-1 flex items-center justify-center h-10 w-10 rounded-full focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white"
          >
            <span class="sr-only">Close sidebar</span>
            <!-- X Icon -->
            <svg 
              class="h-6 w-6 text-white" 
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

        <!-- Sidebar content -->
        <div class="flex-shrink-0 flex items-center px-4">
          <div class="h-8 w-8 bg-white rounded-md flex items-center justify-center">
            <span class="text-brand-600 font-bold">E</span>
          </div>
          <span class="ml-2 text-xl font-bold text-white">Ecomflex</span>
        </div>

        <div class="mt-5 flex-1 h-0 overflow-y-auto">
          <nav class="px-2 space-y-1">
            <!-- Mobile nav items -->
            <router-link 
              v-for="item in navigation" 
              :key="item.name"
              :to="item.href"
              class="group flex items-center px-2 py-2 text-base font-medium rounded-md"
              :class="[
                isActiveRoute(item.href)
                  ? 'bg-brand-700 text-white'
                  : 'text-brand-100 hover:bg-brand-700 hover:text-white'
              ]"
              @click="sidebarOpen = false"
            >
              <component 
                :is="item.icon" 
                class="mr-4 flex-shrink-0 h-6 w-6" 
                :class="[
                  isActiveRoute(item.href) 
                    ? 'text-white' 
                    : 'text-brand-300 group-hover:text-white'
                ]"
                aria-hidden="true" 
              />
              {{ item.name }}
            </router-link>
          </nav>
        </div>
      </div>

      <div class="flex-shrink-0 w-14" aria-hidden="true">
        <!-- Force sidebar to shrink to fit close icon -->
      </div>
    </div>

    <!-- Static sidebar for desktop -->
    <div class="hidden md:flex md:flex-shrink-0">
      <div class="flex flex-col w-64">
        <!-- Sidebar component -->
        <div class="flex flex-col h-0 flex-1">
          <div class="flex items-center h-16 flex-shrink-0 px-4 bg-brand-800">
            <div class="h-8 w-8 bg-white rounded-md flex items-center justify-center">
              <span class="text-brand-600 font-bold">E</span>
            </div>
            <span class="ml-2 text-xl font-bold text-white">Ecomflex</span>
          </div>
          <div class="flex-1 flex flex-col overflow-y-auto">
            <nav class="flex-1 px-2 py-4 bg-brand-800 space-y-1">
              <!-- Desktop nav items -->
              <router-link 
                v-for="item in navigation" 
                :key="item.name"
                :to="item.href"
                class="group flex items-center px-2 py-2 text-sm font-medium rounded-md"
                :class="[
                  isActiveRoute(item.href)
                    ? 'bg-brand-700 text-white'
                    : 'text-brand-100 hover:bg-brand-700 hover:text-white'
                ]"
              >
                <component 
                  :is="item.icon" 
                  class="mr-3 flex-shrink-0 h-6 w-6" 
                  :class="[
                    isActiveRoute(item.href) 
                      ? 'text-white' 
                      : 'text-brand-300 group-hover:text-white'
                  ]"
                  aria-hidden="true" 
                />
                {{ item.name }}
              </router-link>
            </nav>
          </div>
        </div>
      </div>
    </div>

    <!-- Main content -->
    <div class="flex flex-col w-0 flex-1 overflow-hidden">
      <!-- Top navigation bar -->
      <div class="relative z-10 flex-shrink-0 flex h-16 bg-white shadow">
        <button 
          @click="sidebarOpen = true"
          class="px-4 border-r border-gray-200 text-gray-500 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-brand-500 md:hidden"
        >
          <span class="sr-only">Open sidebar</span>
          <svg 
            class="h-6 w-6" 
            xmlns="http://www.w3.org/2000/svg" 
            fill="none" 
            viewBox="0 0 24 24" 
            stroke="currentColor" 
            aria-hidden="true"
          >
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h7" />
          </svg>
        </button>
        
        <div class="flex-1 px-4 flex justify-between">
          <div class="flex-1 flex items-center">
            <!-- Page title (optional) -->
            <h1 class="text-xl font-semibold text-gray-900">Super Admin</h1>
          </div>
          
          <div class="ml-4 flex items-center md:ml-6">
            <!-- Notifications dropdown -->
            <div class="ml-3 relative">
              <button
                class="bg-white p-1 rounded-full text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500"
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
              <!-- Notification badge -->
              <div class="absolute top-0 right-0 h-3 w-3 bg-red-500 rounded-full border-2 border-white"></div>
            </div>

            <!-- Profile dropdown -->
            <div class="ml-3 relative">
              <div>
                <button 
                  @click="profileMenuOpen = !profileMenuOpen"
                  class="max-w-xs bg-white flex items-center text-sm rounded-full focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500"
                  id="user-menu-button" 
                  aria-expanded="false" 
                  aria-haspopup="true"
                >
                  <span class="sr-only">Open user menu</span>
                  <div class="h-8 w-8 rounded-full bg-brand-100 flex items-center justify-center">
                    <span class="font-medium text-brand-600">SA</span>
                  </div>
                </button>
              </div>

              <!-- Profile dropdown menu -->
              <div 
                v-show="profileMenuOpen"
                class="origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5 focus:outline-none"
                role="menu" 
                aria-orientation="vertical" 
                aria-labelledby="user-menu-button" 
                tabindex="-1"
              >
                <a 
                  href="#" 
                  class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100" 
                  role="menuitem" 
                  tabindex="-1" 
                  id="user-menu-item-0"
                >
                  Your Profile
                </a>
                <a 
                  href="#" 
                  class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100" 
                  role="menuitem" 
                  tabindex="-1" 
                  id="user-menu-item-1"
                >
                  Settings
                </a>
                <a 
                  href="#" 
                  @click.prevent="logout"
                  class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100" 
                  role="menuitem" 
                  tabindex="-1" 
                  id="user-menu-item-2"
                >
                  Sign out
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Main content area -->
      <main class="flex-1 relative overflow-y-auto focus:outline-none">
        <div class="py-6">
          <router-view></router-view>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, h, onMounted, onUnmounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'

// Router
const route = useRoute()
const router = useRouter()

// State
const sidebarOpen = ref(false)
const profileMenuOpen = ref(false)

// Navigation items
const navigation = [
  { 
    name: 'Dashboard', 
    href: '/admin/dashboard', 
    icon: DashboardIcon 
  },
  { 
    name: 'Products', 
    href: '/admin/products', 
    icon: ProductsIcon 
  },
  { 
    name: 'Proofs', 
    href: '/admin/proofs', 
    icon: ProofsIcon 
  },
  { 
    name: 'Influencer Requests', 
    href: '/admin/influencers', 
    icon: InfluencersIcon 
  },
  { 
    name: 'User Management', 
    href: '/admin/users', 
    icon: UsersIcon 
  },
  { 
    name: 'Audit Logs', 
    href: '/admin/logs', 
    icon: LogsIcon 
  },
]

// Check if route is active
const isActiveRoute = (path: string) => {
  return route.path.startsWith(path)
}

// Logout
const logout = () => {
  // Clear auth tokens
  localStorage.removeItem('token')
  localStorage.removeItem('refreshToken')
  localStorage.removeItem('userRole')
  
  // Redirect to login
  router.push('/auth/login')
}

// Close dropdown when clicking outside
const handleClickOutside = (event: MouseEvent) => {
  const target = event.target as HTMLElement
  
  if (profileMenuOpen.value && !target.closest('#user-menu-button')) {
    profileMenuOpen.value = false
  }
}

// Add click event listener when component is mounted
onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

// Clean up event listener when component is unmounted
onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})

// Icon components
function DashboardIcon() {
  return h('svg', {
    xmlns: 'http://www.w3.org/2000/svg',
    class: 'h-6 w-6',
    fill: 'none',
    viewBox: '0 0 24 24',
    stroke: 'currentColor'
  }, [
    h('path', {
      'stroke-linecap': 'round',
      'stroke-linejoin': 'round',
      'stroke-width': '2',
      d: 'M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6'
    })
  ])
}

function ProductsIcon() {
  return h('svg', {
    xmlns: 'http://www.w3.org/2000/svg',
    class: 'h-6 w-6',
    fill: 'none',
    viewBox: '0 0 24 24',
    stroke: 'currentColor'
  }, [
    h('path', {
      'stroke-linecap': 'round',
      'stroke-linejoin': 'round',
      'stroke-width': '2',
      d: 'M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10'
    })
  ])
}

function ProofsIcon() {
  return h('svg', {
    xmlns: 'http://www.w3.org/2000/svg',
    class: 'h-6 w-6',
    fill: 'none',
    viewBox: '0 0 24 24',
    stroke: 'currentColor'
  }, [
    h('path', {
      'stroke-linecap': 'round',
      'stroke-linejoin': 'round',
      'stroke-width': '2',
      d: 'M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z'
    })
  ])
}

function InfluencersIcon() {
  return h('svg', {
    xmlns: 'http://www.w3.org/2000/svg',
    class: 'h-6 w-6',
    fill: 'none',
    viewBox: '0 0 24 24',
    stroke: 'currentColor'
  }, [
    h('path', {
      'stroke-linecap': 'round',
      'stroke-linejoin': 'round',
      'stroke-width': '2',
      d: 'M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z'
    })
  ])
}

function UsersIcon() {
  return h('svg', {
    xmlns: 'http://www.w3.org/2000/svg',
    class: 'h-6 w-6',
    fill: 'none',
    viewBox: '0 0 24 24',
    stroke: 'currentColor'
  }, [
    h('path', {
      'stroke-linecap': 'round',
      'stroke-linejoin': 'round',
      'stroke-width': '2',
      d: 'M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z'
    })
  ])
}

function LogsIcon() {
  return h('svg', {
    xmlns: 'http://www.w3.org/2000/svg',
    class: 'h-6 w-6',
    fill: 'none',
    viewBox: '0 0 24 24',
    stroke: 'currentColor'
  }, [
    h('path', {
      'stroke-linecap': 'round',
      'stroke-linejoin': 'round',
      'stroke-width': '2',
      d: 'M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01'
    })
  ])
}
</script>