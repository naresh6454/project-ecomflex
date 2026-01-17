<template>
  <div class="h-screen flex overflow-hidden bg-gray-100">
    <!-- Mobile Sidebar Toggle -->
    <div class="fixed inset-0 flex z-40 md:hidden" role="dialog" aria-modal="true" v-if="sidebarOpen">
      <!-- Sidebar backdrop -->
      <div 
        class="fixed inset-0 bg-gray-600 bg-opacity-75 transition-opacity ease-out duration-300" 
        aria-hidden="true"
        @click="sidebarOpen = false"
        v-motion
        :initial="{ opacity: 0 }"
        :enter="{ opacity: 1 }"
        :leave="{ opacity: 0 }">
      </div>

      <!-- Sidebar panel -->
      <div 
        class="relative flex-1 flex flex-col max-w-xs w-full bg-white"
        v-motion
        :initial="{ x: -280 }"
        :enter="{ x: 0 }"
        :leave="{ x: -280 }">
        <!-- Close button -->
        <div class="absolute top-0 right-0 -mr-12 pt-2">
          <button 
            type="button" 
            class="ml-1 flex items-center justify-center h-10 w-10 rounded-full focus:outline-none focus:ring-2 focus:ring-inset focus:ring-accent" 
            @click="sidebarOpen = false">
            <span class="sr-only">Close sidebar</span>
            <svg class="h-6 w-6 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>

        <!-- Sidebar content -->
        <div class="flex-1 h-0 pt-5 pb-4 overflow-y-auto">
          <div class="flex-shrink-0 flex items-center px-4">
            <img class="h-10 w-auto" src="/ecomflex-logo.png" alt="Ecomflex" />
          </div>
          <nav class="mt-5 px-2 space-y-1">
            <router-link 
              v-for="item in navigation" 
              :key="item.name" 
              :to="item.href" 
              :class="[
                isActive(item.href) 
                  ? 'bg-accent bg-opacity-10 text-accent' 
                  : 'text-gray-600 hover:bg-accent hover:bg-opacity-5 hover:text-accent',
                'group flex items-center px-2 py-2 text-base font-medium rounded-md transition-colors'
              ]">
              <component 
                :is="item.icon" 
                :class="[
                  isActive(item.href) 
                    ? 'text-accent' 
                    : 'text-gray-400 group-hover:text-accent',
                  'mr-4 flex-shrink-0 h-6 w-6 transition-colors'
                ]" 
                aria-hidden="true" />
              {{ item.name }}
            </router-link>
          </nav>
        </div>
        <!-- User profile -->
        <div class="flex-shrink-0 flex border-t border-gray-200 p-4">
          <div class="flex-shrink-0 group block">
            <div class="flex items-center">
              <div>
                <img class="inline-block h-10 w-10 rounded-full" src="https://images.unsplash.com/photo-1550525811-e5869dd03032?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2.5&w=256&h=256&q=80" alt="" />
              </div>
              <div class="ml-3">
                <p class="text-base font-medium text-gray-700 group-hover:text-gray-900">Sarah Johnson</p>
                <p class="text-sm font-medium text-gray-500 group-hover:text-gray-700">Influencer</p>
              </div>
            </div>
          </div>
        </div>
      </div>
      
      <div class="flex-shrink-0 w-14">
        <!-- Force sidebar to shrink to fit close icon -->
      </div>
    </div>

    <!-- Static sidebar for desktop -->
    <div class="hidden md:flex md:flex-shrink-0">
      <div class="flex flex-col w-64">
        <!-- Sidebar component -->
        <div class="flex flex-col h-0 flex-1 border-r border-gray-200 bg-white">
          <div class="flex-1 flex flex-col pt-5 pb-4 overflow-y-auto">
            <div class="flex items-center flex-shrink-0 px-4">
              <img class="h-10 w-auto" src="/ecomflex-logo.png" alt="Ecomflex" />
            </div>
            <nav class="mt-5 flex-1 px-2 bg-white space-y-1">
              <router-link 
                v-for="item in navigation" 
                :key="item.name" 
                :to="item.href" 
                :class="[
                  isActive(item.href) 
                    ? 'bg-accent bg-opacity-10 text-accent' 
                    : 'text-gray-600 hover:bg-accent hover:bg-opacity-5 hover:text-accent',
                  'group flex items-center px-2 py-2 text-sm font-medium rounded-md transition-colors'
                ]">
                <component 
                  :is="item.icon" 
                  :class="[
                    isActive(item.href) 
                      ? 'text-accent' 
                      : 'text-gray-400 group-hover:text-accent',
                    'mr-3 flex-shrink-0 h-6 w-6 transition-colors'
                  ]" 
                  aria-hidden="true" />
                {{ item.name }}
              </router-link>
            </nav>
          </div>
          <!-- User profile -->
          <div class="flex-shrink-0 flex border-t border-gray-200 p-4">
            <div class="flex-shrink-0 w-full group block">
              <div class="flex items-center">
                <div>
                  <img class="inline-block h-9 w-9 rounded-full" src="https://images.unsplash.com/photo-1550525811-e5869dd03032?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2.5&w=256&h=256&q=80" alt="" />
                </div>
                <div class="ml-3">
                  <p class="text-sm font-medium text-gray-700 group-hover:text-gray-900">Sarah Johnson</p>
                  <p class="text-xs font-medium text-gray-500 group-hover:text-gray-700">Influencer</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Main content -->
    <div class="flex flex-col w-0 flex-1 overflow-hidden">
      <!-- Top header bar with profile dropdown -->
      <div class="relative z-10 flex-shrink-0 flex h-16 bg-white shadow">
        <button 
          type="button" 
          class="px-4 border-r border-gray-200 text-gray-500 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-accent md:hidden" 
          @click="sidebarOpen = true">
          <span class="sr-only">Open sidebar</span>
          <svg class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
          </svg>
        </button>
        
        <!-- Header content -->
        <div class="flex-1 px-4 flex justify-between items-center">
          <!-- Page title or breadcrumb could go here -->
          <div class="flex-1">
            <!-- Empty space or page title -->
          </div>
          
          <!-- Profile dropdown -->
          <div class="ml-4 flex items-center md:ml-6">
            <div class="relative">
              <div>
                <button 
                  type="button" 
                  class="max-w-xs bg-white rounded-full flex items-center text-sm focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-accent hover:bg-gray-50 transition-colors duration-200" 
                  id="user-menu-button" 
                  aria-expanded="false" 
                  aria-haspopup="true"
                  @click="toggleDropdown">
                  <span class="sr-only">Open user menu</span>
                  <div class="h-8 w-8 rounded-full bg-accent flex items-center justify-center">
                    <svg class="h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                    </svg>
                  </div>
                </button>
              </div>

              <!-- Dropdown menu -->
              <transition
                enter-active-class="transition ease-out duration-100"
                enter-from-class="transform opacity-0 scale-95"
                enter-to-class="transform opacity-100 scale-100"
                leave-active-class="transition ease-in duration-75"
                leave-from-class="transform opacity-100 scale-100"
                leave-to-class="transform opacity-0 scale-95">
                <div 
                  v-if="dropdownOpen" 
                  class="origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5 focus:outline-none" 
                  role="menu" 
                  aria-orientation="vertical" 
                  aria-labelledby="user-menu-button" 
                  tabindex="-1">
                  
                  <router-link 
                    to="/influencer/profile" 
                    class="flex items-center px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 transition-colors duration-150" 
                    role="menuitem" 
                    tabindex="-1"
                    @click="closeDropdown">
                    <svg class="mr-3 h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                    </svg>
                    Profile
                  </router-link>
                  
                  <button 
                    @click="handleSignOut" 
                    class="w-full flex items-center px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 transition-colors duration-150" 
                    role="menuitem" 
                    tabindex="-1">
                    <svg class="mr-3 h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                    </svg>
                    Sign Out
                  </button>
                </div>
              </transition>
            </div>
          </div>
        </div>
      </div>

      <!-- Remove the old mobile menu button that was here -->
      <main class="flex-1 relative z-0 overflow-y-auto focus:outline-none">
        <div class="py-6 px-4 sm:px-6 md:px-8">
          <!-- This router-view is the critical part that was missing -->
          <router-view></router-view>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, h, onMounted, onUnmounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

// Navigation icons as components
const HomeIcon = {
  render() {
    return h('svg', { 
      xmlns: "http://www.w3.org/2000/svg", 
      fill: "none", 
      viewBox: "0 0 24 24", 
      stroke: "currentColor", 
      'stroke-width': "1.5" 
    }, [
      h('path', { 
        'stroke-linecap': "round", 
        'stroke-linejoin': "round", 
        d: "M2.25 12l8.954-8.955c.44-.439 1.152-.439 1.591 0L21.75 12M4.5 9.75v10.125c0 .621.504 1.125 1.125 1.125H9.75v-4.875c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21h4.125c.621 0 1.125-.504 1.125-1.125V9.75M8.25 21h8.25" 
      })
    ])
  }
}

const ChartIcon = {
  render() {
    return h('svg', { 
      xmlns: "http://www.w3.org/2000/svg", 
      fill: "none", 
      viewBox: "0 0 24 24", 
      stroke: "currentColor", 
      'stroke-width': "1.5" 
    }, [
      h('path', { 
        'stroke-linecap': "round", 
        'stroke-linejoin': "round", 
        d: "M3 13.125C3 12.504 3.504 12 4.125 12h2.25c.621 0 1.125.504 1.125 1.125v6.75C7.5 20.496 6.996 21 6.375 21h-2.25A1.125 1.125 0 013 19.875v-6.75zM9.75 8.625c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125v11.25c0 .621-.504 1.125-1.125 1.125h-2.25a1.125 1.125 0 01-1.125-1.125V8.625zM16.5 4.125c0-.621.504-1.125 1.125-1.125h2.25C20.496 3 21 3.504 21 4.125v15.75c0 .621-.504 1.125-1.125 1.125h-2.25a1.125 1.125 0 01-1.125-1.125V4.125z" 
      })
    ])
  }
}

const ProductIcon = {
  render() {
    return h('svg', { 
      xmlns: "http://www.w3.org/2000/svg", 
      fill: "none", 
      viewBox: "0 0 24 24", 
      stroke: "currentColor", 
      'stroke-width': "1.5" 
    }, [
      h('path', { 
        'stroke-linecap': "round", 
        'stroke-linejoin': "round", 
        d: "M20.25 7.5l-.625 10.632a2.25 2.25 0 01-2.247 2.118H6.622a2.25 2.25 0 01-2.247-2.118L3.75 7.5M10 11.25h4M3.375 7.5h17.25c.621 0 1.125-.504 1.125-1.125v-1.5c0-.621-.504-1.125-1.125-1.125H3.375c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125z" 
      })
    ])
  }
}

const LinkIcon = {
  render() {
    return h('svg', { 
      xmlns: "http://www.w3.org/2000/svg", 
      fill: "none", 
      viewBox: "0 0 24 24", 
      stroke: "currentColor", 
      'stroke-width': "1.5" 
    }, [
      h('path', { 
        'stroke-linecap': "round", 
        'stroke-linejoin': "round", 
        d: "M13.19 8.688a4.5 4.5 0 011.242 7.244l-4.5 4.5a4.5 4.5 0 01-6.364-6.364l1.757-1.757m13.35-.622l1.757-1.757a4.5 4.5 0 00-6.364-6.364l-4.5 4.5a4.5 4.5 0 001.242 7.244" 
      })
    ])
  }
}

const MoneyIcon = {
  render() {
    return h('svg', { 
      xmlns: "http://www.w3.org/2000/svg", 
      fill: "none", 
      viewBox: "0 0 24 24", 
      stroke: "currentColor", 
      'stroke-width': "1.5" 
    }, [
      h('path', { 
        'stroke-linecap': "round", 
        'stroke-linejoin': "round", 
        d: "M2.25 18.75a60.07 60.07 0 0115.797 2.101c.727.198 1.453-.342 1.453-1.096V18.75M3.75 4.5v.75A.75.75 0 013 6h-.75m0 0v-.375c0-.621.504-1.125 1.125-1.125H20.25M2.25 6v9m18-10.5v.75c0 .414.336.75.75.75h.75m-1.5-1.5h.375c.621 0 1.125.504 1.125 1.125v9.75c0 .621-.504 1.125-1.125 1.125h-.375m1.5-1.5H21a.75.75 0 00-.75.75v.75m0 0H3.75m0 0h-.375a1.125 1.125 0 01-1.125-1.125V15m1.5 1.5v-.75A.75.75 0 003 15h-.75M15 10.5a3 3 0 11-6 0 3 3 0 016 0zm3 0h.008v.008H18V10.5zm-12 0h.008v.008H6V10.5z" 
      })
    ])
  }
}

const UserIcon = {
  render() {
    return h('svg', { 
      xmlns: "http://www.w3.org/2000/svg", 
      fill: "none", 
      viewBox: "0 0 24 24", 
      stroke: "currentColor", 
      'stroke-width': "1.5" 
    }, [
      h('path', { 
        'stroke-linecap': "round", 
        'stroke-linejoin': "round", 
        d: "M15.75 6a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0zM4.501 20.118a7.5 7.5 0 0114.998 0A17.933 17.933 0 0112 21.75c-2.676 0-5.216-.584-7.499-1.632z" 
      })
    ])
  }
}

const SettingsIcon = {
  render() {
    return h('svg', { 
      xmlns: "http://www.w3.org/2000/svg", 
      fill: "none", 
      viewBox: "0 0 24 24", 
      stroke: "currentColor", 
      'stroke-width': "1.5" 
    }, [
      h('path', { 
        'stroke-linecap': "round", 
        'stroke-linejoin': "round", 
        d: "M9.594 3.94c.09-.542.56-.94 1.11-.94h2.593c.55 0 1.02.398 1.11.94l.213 1.281c.063.374.313.686.645.87.074.04.147.083.22.127.324.196.72.257 1.075.124l1.217-.456a1.125 1.125 0 011.37.49l1.296 2.247a1.125 1.125 0 01-.26 1.431l-1.003.827c-.293.24-.438.613-.431.992a6.759 6.759 0 010 .255c-.007.378.138.75.43.99l1.005.828c.424.35.534.954.26 1.43l-1.298 2.247a1.125 1.125 0 01-1.369.491l-1.217-.456c-.355-.133-.75-.072-1.076.124a6.57 6.57 0 01-.22.128c-.331.183-.581.495-.644.869l-.213 1.28c-.09.543-.56.941-1.11.941h-2.594c-.55 0-1.02-.398-1.11-.94l-.213-1.281c-.062-.374-.312-.686-.644-.87a6.52 6.52 0 01-.22-.127c-.325-.196-.72-.257-1.076-.124l-1.217.456a1.125 1.125 0 01-1.369-.49l-1.297-2.247a1.125 1.125 0 01.26-1.431l1.004-.827c.292-.24.437-.613.43-.992a6.932 6.932 0 010-.255c.007-.378-.138-.75-.43-.99l-1.004-.828a1.125 1.125 0 01-.26-1.43l1.297-2.247a1.125 1.125 0 011.37-.491l1.216.456c.356.133.751.072 1.076-.124.072-.044.146-.087.22-.128.332-.183.582-.495.644-.869l.214-1.281z" 
      }),
      h('path', { 
        'stroke-linecap': "round", 
        'stroke-linejoin': "round", 
        d: "M15 12a3 3 0 11-6 0 3 3 0 016 0z" 
      })
    ])
  }
}

// Navigation configuration
const navigation = [
  { name: 'Dashboard', href: '/influencer/dashboard', icon: HomeIcon },
  { name: 'Referrals', href: '/influencer/referrals', icon: LinkIcon },
  { name: 'Products', href: '/influencer/products', icon: ProductIcon },
  { name: 'Earnings', href: '/influencer/earnings', icon: MoneyIcon },
  { name: 'Profile', href: '/influencer/profile', icon: UserIcon },
  { name: 'Settings', href: '/influencer/settings', icon: SettingsIcon },
]

// State
const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()
const sidebarOpen = ref(false)
const dropdownOpen = ref(false)

// Check if a route is active
const isActive = (href: string) => {
  return route.path === href || route.path.startsWith(`${href}/`)
}

// Dropdown functions
const toggleDropdown = () => {
  dropdownOpen.value = !dropdownOpen.value
}

const closeDropdown = () => {
  dropdownOpen.value = false
}

// Handle sign out
const handleSignOut = async () => {
  try {
    // Close dropdown first
    closeDropdown()
    
    // Call auth store logout method
    await authStore.logout()
    
    // Clear any remaining localStorage items
    localStorage.removeItem('token')
    localStorage.removeItem('accessToken')
    localStorage.removeItem('refreshToken')
    localStorage.removeItem('userRole')
    localStorage.removeItem('isSuperadminMode')
    
    // Redirect to login page
    router.push('/auth/login')
  } catch (error) {
    console.error('Error during sign out:', error)
    // Even if there's an error, still redirect to login
    router.push('/auth/login')
  }
}

// Click outside handler
const handleClickOutside = (event: Event) => {
  const target = event.target as HTMLElement
  const dropdown = target.closest('[role="menu"]')
  const button = target.closest('#user-menu-button')
  
  if (!dropdown && !button) {
    closeDropdown()
  }
}

// Lifecycle hooks
onMounted(() => {
  console.log("InfluencerLayout mounted");
  // Add click outside listener
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  // Remove click outside listener
  document.removeEventListener('click', handleClickOutside)
})
</script>