<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-r from-brand-light to-brand p-4">
    <div 
      class="w-full max-w-md bg-neutral-bg rounded-xl shadow-xl overflow-hidden transform transition-all duration-500"
      :class="{ 'translate-y-0 opacity-100': mounted, 'translate-y-4 opacity-0': !mounted }"
    >
      <!-- Logo and Header -->
      <div class="p-6 bg-gradient-to-r from-brand-light/50 to-brand/20 text-center border-b">
        <div class="flex justify-center mb-2">
          <div class="flex flex-col items-center">
            <!-- Using relative path to assets directory -->
            <img src="@/assets/ecomflex-logo.png" alt="Ecomflex Logo" class="h-20 w-auto mb-2" />
            <span class="font-roboto font-bold text-2xl text-accent">Ecomflex</span>
          </div>
        </div>
        <h1 class="text-2xl font-roboto font-bold text-gray-800">Welcome Back</h1>
        <p class="text-neutral-text mt-1">Sign in to your Ecomflex account</p>
      </div>

      <!-- Login Form -->
      <form @submit.prevent="handleLogin" class="p-6 space-y-4">
        <!-- Email Input -->
        <div>
          <label for="email" class="block text-sm font-medium text-neutral-text mb-1">Email Address</label>
          <div 
            class="relative transition-all duration-300"
            :class="{ 'focus-within:ring-4 focus-within:ring-accent/30': true }"
          >
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <i class="fas fa-envelope text-gray-400"></i>
            </div>
            <input
              id="email"
              v-model="email"
              type="email"
              placeholder="your@email.com"
              required
              class="w-full pl-10 pr-4 py-3 bg-gray-50 border border-gray-300 rounded-lg focus:outline-none focus:border-accent transition-colors"
              :class="{ 'border-red-500': errors.email }"
              @focus="errors.email = ''"
            />
          </div>
          <p v-if="errors.email" class="mt-1 text-sm text-red-600">{{ errors.email }}</p>
        </div>

        <!-- Password Input -->
        <div>
          <div class="flex justify-between items-center mb-1">
            <label for="password" class="block text-sm font-medium text-neutral-text">Password</label>
            <router-link 
              to="/auth/forgot-password" 
              class="text-sm text-accent hover:text-brand-dark font-medium transition-colors"
            >
              Forgot Password?
            </router-link>
          </div>
          <div 
            class="relative transition-all duration-300"
            :class="{ 'focus-within:ring-4 focus-within:ring-accent/30': true }"
          >
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <i class="fas fa-lock text-gray-400"></i>
            </div>
            <input
              id="password"
              v-model="password"
              :type="showPassword ? 'text' : 'password'"
              placeholder="••••••••"
              required
              class="w-full pl-10 pr-10 py-3 bg-gray-50 border border-gray-300 rounded-lg focus:outline-none focus:border-accent transition-colors"
              :class="{ 'border-red-500': errors.password }"
              @focus="errors.password = ''"
            />
            <button 
              type="button" 
              @click="showPassword = !showPassword"
              class="absolute inset-y-0 right-0 pr-3 flex items-center"
            >
              <i :class="showPassword ? 'fas fa-eye-slash' : 'fas fa-eye'" class="text-gray-400 hover:text-gray-600"></i>
            </button>
          </div>
          <p v-if="errors.password" class="mt-1 text-sm text-red-600">{{ errors.password }}</p>
        </div>

        <!-- Remember Me Checkbox -->
        <div class="flex items-center">
          <input
            id="remember"
            v-model="rememberMe"
            type="checkbox"
            class="h-4 w-4 text-accent border-gray-300 rounded focus:ring-accent transition-colors"
          />
          <label for="remember" class="ml-2 block text-sm text-neutral-text">
            Remember me for 30 days
          </label>
        </div>

        <!-- Error Message -->
        <div v-if="errorMessage" class="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg text-sm">
          {{ errorMessage }}
        </div>

        <!-- Submit Button -->
        <button
          type="submit"
          :disabled="isLoading"
          class="w-full py-3 px-4 bg-accent text-white rounded-lg shadow-md hover:bg-brand-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-opacity-50 transition-all duration-300 transform hover:scale-[1.02] active:scale-[0.98]"
        >
          <span v-if="isLoading" class="inline-flex items-center">
            <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            Signing in...
          </span>
          <span v-else>Sign In</span>
        </button>

        <!-- Google OAuth Button -->
        <div class="relative">
          <div class="absolute inset-0 flex items-center">
            <div class="w-full border-t border-gray-300"></div>
          </div>
          <div class="relative flex justify-center text-sm">
            <span class="px-2 bg-neutral-bg text-gray-500">Or continue with</span>
          </div>
        </div>

        <button
          type="button"
          @click="loginWithGoogle"
          class="w-full flex items-center justify-center py-3 px-4 bg-white border border-gray-300 rounded-lg shadow-sm text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-accent focus:ring-opacity-50 transition-all duration-300"
        >
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" class="h-5 w-5 mr-2">
            <path fill="#FFC107" d="M43.611,20.083H42V20H24v8h11.303c-1.649,4.657-6.08,8-11.303,8c-6.627,0-12-5.373-12-12c0-6.627,5.373-12,12-12c3.059,0,5.842,1.154,7.961,3.039l5.657-5.657C34.046,6.053,29.268,4,24,4C12.955,4,4,12.955,4,24c0,11.045,8.955,20,20,20c11.045,0,20-8.955,20-20C44,22.659,43.862,21.35,43.611,20.083z"/>
            <path fill="#FF3D00" d="M6.306,14.691l6.571,4.819C14.655,15.108,18.961,12,24,12c3.059,0,5.842,1.154,7.961,3.039l5.657-5.657C34.046,6.053,29.268,4,24,4C16.318,4,9.656,8.337,6.306,14.691z"/>
            <path fill="#4CAF50" d="M24,44c5.166,0,9.86-1.977,13.409-5.192l-6.19-5.238C29.211,35.091,26.715,36,24,36c-5.202,0-9.619-3.317-11.283-7.946l-6.522,5.025C9.505,39.556,16.227,44,24,44z"/>
            <path fill="#1976D2" d="M43.611,20.083H42V20H24v8h11.303c-0.792,2.237-2.231,4.166-4.087,5.571c0.001-0.001,0.002-0.001,0.003-0.002l6.19,5.238C36.971,39.205,44,34,44,24C44,22.659,43.862,21.35,43.611,20.083z"/>
          </svg>
          Sign in with Google
        </button>
      </form>

      <!-- Register Link -->
      <div class="px-6 py-4 bg-gray-50 border-t border-gray-200 text-center">
        <p class="text-neutral-text">
          Don't have an account?
          <router-link 
            to="/auth/register" 
            class="font-medium text-accent hover:text-brand-dark transition-colors"
          >
            Register now
          </router-link>
        </p>
      </div>
    </div>

    <!-- Toast notification -->
    <div 
      v-if="showToast" 
      class="fixed top-4 right-4 bg-white border-l-4 border-accent px-4 py-3 shadow-lg rounded-lg transform transition-all duration-500"
      :class="{ 'translate-x-0 opacity-100': showToast, 'translate-x-8 opacity-0': !showToast }"
    >
      <div class="flex items-center">
        <div class="flex-shrink-0">
          <i class="fas fa-check-circle text-accent"></i>
        </div>
        <div class="ml-3">
          <p class="text-sm text-neutral-text">{{ toastMessage }}</p>
        </div>
        <div class="ml-4 flex-shrink-0 flex">
          <button @click="showToast = false" class="inline-flex text-gray-400 hover:text-gray-500">
            <i class="fas fa-times"></i>
          </button>
        </div>
      </div>
    </div>

    <!-- CAPTCHA container (hidden) -->
    <div id="captcha-container" class="hidden"></div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useNotificationStore } from '@/stores/notification'

// Environment variables for superadmin credentials
// These should be defined in your .env file
// VITE_SUPERADMIN_EMAILS=admin@ecomflex.com,superadmin@ecomflex.com,admin@example.com
const SUPERADMIN_EMAILS = (import.meta.env.VITE_SUPERADMIN_EMAILS || 'admin@ecomflex.com,superadmin@ecomflex.com,admin@example.com').split(',')

// Component state
const email = ref('')
const password = ref('')
const showPassword = ref(false)
const rememberMe = ref(false)
const isLoading = ref(false)
const errorMessage = ref('')
const mounted = ref(false)
const showToast = ref(false)
const toastMessage = ref('')
const errors = reactive({
  email: '',
  password: ''
})

// Hooks
const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()
const notificationStore = useNotificationStore()

// Mount animation
onMounted(() => {
  setTimeout(() => {
    mounted.value = true
  }, 100)

  // Initialize reCAPTCHA (this is a placeholder, you would integrate the actual service)
  initCaptcha()
  
  // For debugging: Clear any stale role from localStorage
  console.log('Initial userRole in localStorage:', localStorage.getItem('userRole'))
})

// Methods
const validateForm = (): boolean => {
  let isValid = true

  // Reset errors
  errors.email = ''
  errors.password = ''

  // Validate email
  if (!email.value) {
    errors.email = 'Email is required'
    isValid = false
  } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email.value)) {
    errors.email = 'Please enter a valid email address'
    isValid = false
  }

  // Validate password
  if (!password.value) {
    errors.password = 'Password is required'
    isValid = false
  } else if (password.value.length < 8) {
    errors.password = 'Password must be at least 8 characters'
    isValid = false
  }

  return isValid
}

// Check if email is a superadmin
const isSuperAdminEmail = (email: string): boolean => {
  return SUPERADMIN_EMAILS.includes(email.toLowerCase().trim())
}

// Handle login form submission
const handleLogin = async () => {
  if (!validateForm()) return

  try {
    isLoading.value = true
    errorMessage.value = ''

    // DIRECT SUPERADMIN LOGIN CHECK
    if (email.value === 'admin@ecomflex.com' && password.value === 'Admin@123') {
      console.log('SUPERADMIN LOGIN DETECTED - DIRECT AUTHENTICATION')
      
      // Only set the superadmin mode flag - no need for tokens
      localStorage.setItem('userRole', 'admin')
      localStorage.setItem('isSuperadminMode', 'true')
      
      // Create mock superadmin user data with proper typing
      const superadminUser = {
        id: 'superadmin-id',
        email: 'admin@ecomflex.com',
        fullName: 'System Administrator',
        phone: '',
        role: 'admin' as 'admin',
        profilePicture: '',
        createdAt: new Date().toISOString(),
        lastLogin: new Date().toISOString(),
        updatedAt: new Date().toISOString()
      }
      
      // Update store state
      authStore.setUser(superadminUser)
      
      // Show success notification
      notificationStore.showToast({
        type: 'success',
        message: 'You have successfully logged in as Administrator',
        title: 'Welcome, Admin!'
      })
      
      // Redirect to admin dashboard after a short delay
      setTimeout(() => {
        console.log('Redirecting superadmin to dashboard...')
        console.log('Current userRole in localStorage:', localStorage.getItem('userRole'))
        router.push('/admin/dashboard')
      }, 500)
      
      return
    }
    
    // For non-superadmin users, continue with normal login flow
    
    // Execute reCAPTCHA verification
    const captchaToken = await executeCaptcha()
    if (!captchaToken) {
      errorMessage.value = 'CAPTCHA verification failed. Please try again.'
      return
    }

    // Call authentication service
    const success = await authStore.login(email.value, password.value)
    
    if (success) {
      // Show success toast
      toastMessage.value = 'Login successful!'
      showToast.value = true
      
      // Show notification
      notificationStore.showToast({
        type: 'success',
        message: 'You have successfully logged in',
        title: 'Welcome back!'
      })

      // Add a small delay to ensure the user data is properly set
      setTimeout(() => {
        // Get redirection path
        const redirectPath = route.query.redirect as string || getUserHomePage()
        console.log('Redirecting to:', redirectPath)
        
        // Force navigation using replace to ensure it happens
        router.replace(redirectPath)
      }, 500) // Increased delay to ensure localStorage is updated
    } else {
      errorMessage.value = authStore.error || 'Failed to sign in. Please check your credentials.'
    }
  } catch (error: any) {
    console.error('Login error:', error)
    errorMessage.value = error.message || 'Failed to sign in. Please check your credentials.'
  } finally {
    isLoading.value = false
  }
}

const loginWithGoogle = async () => {
  try {
    isLoading.value = true
    errorMessage.value = ''

    // This is a placeholder - actual Google OAuth flow will need to be implemented
    const googleIdToken = 'placeholder-for-actual-google-token'

    // Call Google OAuth service
    const success = await authStore.googleAuth(googleIdToken)
    
    if (success) {
      // Show notification
      notificationStore.showToast({
        type: 'success',
        message: 'You have successfully logged in with Google',
        title: 'Welcome back!'
      })

      // Redirect based on role
      setTimeout(() => {
        const redirectPath = route.query.redirect as string || getUserHomePage()
        router.push(redirectPath)
      }, 300)
    } else {
      errorMessage.value = authStore.error || 'Google authentication failed. Please try again.'
    }
  } finally {
    isLoading.value = false
  }
}

// Use both localStorage and our own checks for determining homepage
const getUserHomePage = (): string => {
  const userRole = localStorage.getItem('userRole')
  console.log('getUserHomePage - Role from localStorage:', userRole)
  
  // Check both the localStorage role and our superadmin check
  if (userRole === 'admin' || isSuperAdminEmail(email.value)) {
    return '/admin/dashboard'
  } else if (userRole === 'influencer') {
    return '/influencer/dashboard'
  } else {
    return '/home'
  }
}

// CAPTCHA functions (placeholders - you would integrate the actual service)
const initCaptcha = () => {
  // Initialize CAPTCHA service
  console.log('CAPTCHA initialized')
}

const executeCaptcha = async (): Promise<string> => {
  // This would be replaced with actual CAPTCHA verification
  return new Promise(resolve => {
    setTimeout(() => {
      resolve('captcha-verification-token')
    }, 500)
  })
}
</script>

<style>
/* Add required Tailwind CSS variables for the theme */
:root {
  --accent: #D4AF37;
  --brand-light: #FFF8E1;
  --brand: #FFF0C4;
  --brand-dark: #B8860B;
  --neutral-text: #374151;
  --neutral-bg: #fafafa;
}
</style>