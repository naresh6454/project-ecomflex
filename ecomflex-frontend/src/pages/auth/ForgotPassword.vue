<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-r from-brand-light to-brand p-4">
    <div 
      class="w-full max-w-md bg-white rounded-xl shadow-xl overflow-hidden transform transition-all duration-500"
      :class="{ 'translate-y-0 opacity-100': mounted, 'translate-y-4 opacity-0': !mounted }"
    >
      <!-- Logo and Header -->
      <div class="p-6 bg-gradient-to-r from-brand-light/50 to-brand/20 text-center border-b">
        <div class="flex justify-center mb-2">
          <div class="flex flex-col items-center">
            <img src="@/assets/ecomflex-logo.png" alt="Ecomflex Logo" class="h-20 w-auto mb-2" />
            <span class="font-roboto font-bold text-2xl text-accent">Ecomflex</span>
          </div>
        </div>
        <h1 class="text-2xl font-bold text-gray-800">Reset Your Password</h1>
        <p class="text-gray-600 mt-1">We'll send you a link to reset your password</p>
      </div>

      <!-- Step Indicator with improved spacing -->
      <div class="px-6 pt-6">
        <div class="flex items-center justify-between mb-4">
          <div v-for="(step, index) in steps" :key="index" class="flex flex-col items-center w-20">
            <div 
              class="flex items-center justify-center h-8 w-8 rounded-full text-sm font-medium transition-all duration-300"
              :class="[
                currentStep >= index 
                  ? 'bg-accent text-white' 
                  : 'bg-gray-200 text-gray-500'
              ]"
            >
              {{ index + 1 }}
            </div>
            <div class="text-xs mt-1 text-center whitespace-nowrap" :class="currentStep >= index ? 'text-accent' : 'text-gray-500'">
              {{ step }}
            </div>
          </div>
          <div 
            v-for="index in steps.length - 1" 
            :key="`line-${index}`" 
            class="flex-1 h-0.5 mx-2"
            :class="currentStep > index ? 'bg-accent' : 'bg-gray-200'"
          ></div>
        </div>
      </div>

      <!-- Step 1: Email Entry -->
      <form v-if="currentStep === 0" @submit.prevent="requestResetCode" class="p-6 space-y-4">
        <div>
          <label for="email" class="block text-sm font-medium text-gray-700 mb-1">Email Address</label>
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
            Sending...
          </span>
          <span v-else>Send Reset Code</span>
        </button>
      </form>

      <!-- Step 2: Verification Code Entry -->
      <form v-else-if="currentStep === 1" @submit.prevent="verifyResetCode" class="p-6 space-y-4">
        <div>
          <label for="verificationCode" class="block text-sm font-medium text-gray-700 mb-1">Verification Code</label>
          <div class="space-y-2">
            <p class="text-sm text-gray-600">
              Enter the 6-digit verification code sent to <span class="font-medium">{{ email }}</span>
            </p>
            <div 
              class="relative transition-all duration-300"
              :class="{ 'focus-within:ring-4 focus-within:ring-accent/30': true }"
            >
              <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <i class="fas fa-key text-gray-400"></i>
              </div>
              <input
                id="verificationCode"
                v-model="verificationCode"
                type="text"
                placeholder="123456"
                maxlength="6"
                pattern="[0-9]{6}"
                required
                class="w-full pl-10 pr-4 py-3 bg-gray-50 border border-gray-300 rounded-lg focus:outline-none focus:border-accent transition-colors tracking-widest"
                :class="{ 'border-red-500': errors.verificationCode }"
                @focus="errors.verificationCode = ''"
              />
            </div>
            <p v-if="errors.verificationCode" class="mt-1 text-sm text-red-600">{{ errors.verificationCode }}</p>
          </div>
        </div>

        <div class="flex items-center justify-between">
          <button 
            type="button" 
            @click="resendCode"
            :disabled="resendCooldown > 0 || isResending"
            class="text-sm text-accent hover:text-accent/80 disabled:text-gray-400 transition-colors font-medium focus:outline-none"
          >
            <span v-if="isResending">
              <i class="fas fa-spinner fa-spin mr-1"></i>
              Sending...
            </span>
            <span v-else-if="resendCooldown > 0">
              Resend in {{ resendCooldown }}s
            </span>
            <span v-else>
              Resend Code
            </span>
          </button>
          <button 
            type="button" 
            @click="currentStep = 0"
            class="text-sm text-gray-600 hover:text-gray-800 transition-colors font-medium focus:outline-none"
          >
            Change Email
          </button>
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
            Verifying...
          </span>
          <span v-else>Verify Code</span>
        </button>
      </form>

      <!-- Step 3: Set New Password -->
      <form v-else-if="currentStep === 2" @submit.prevent="resetPassword" class="p-6 space-y-4">
        <!-- New Password -->
        <div>
          <label for="newPassword" class="block text-sm font-medium text-gray-700 mb-1">New Password</label>
          <div 
            class="relative transition-all duration-300"
            :class="{ 'focus-within:ring-4 focus-within:ring-accent/30': true }"
          >
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <i class="fas fa-lock text-gray-400"></i>
            </div>
            <input
              id="newPassword"
              v-model="newPassword"
              :type="showPassword ? 'text' : 'password'"
              placeholder="••••••••"
              required
              class="w-full pl-10 pr-10 py-3 bg-gray-50 border border-gray-300 rounded-lg focus:outline-none focus:border-accent transition-colors"
              :class="{ 'border-red-500': errors.newPassword }"
              @focus="errors.newPassword = ''"
            />
            <button 
              type="button" 
              @click="showPassword = !showPassword"
              class="absolute inset-y-0 right-0 pr-3 flex items-center"
            >
              <i :class="showPassword ? 'fas fa-eye-slash' : 'fas fa-eye'" class="text-gray-400 hover:text-gray-600"></i>
            </button>
          </div>
          <p v-if="errors.newPassword" class="mt-1 text-sm text-red-600">{{ errors.newPassword }}</p>
        </div>

        <!-- Confirm New Password -->
        <div>
          <label for="confirmNewPassword" class="block text-sm font-medium text-gray-700 mb-1">Confirm New Password</label>
          <div 
            class="relative transition-all duration-300"
            :class="{ 'focus-within:ring-4 focus-within:ring-accent/30': true }"
          >
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <i class="fas fa-lock text-gray-400"></i>
            </div>
            <input
              id="confirmNewPassword"
              v-model="confirmNewPassword"
              :type="showPassword ? 'text' : 'password'"
              placeholder="••••••••"
              required
              class="w-full pl-10 pr-4 py-3 bg-gray-50 border border-gray-300 rounded-lg focus:outline-none focus:border-accent transition-colors"
              :class="{ 'border-red-500': errors.confirmNewPassword }"
              @focus="errors.confirmNewPassword = ''"
            />
          </div>
          <p v-if="errors.confirmNewPassword" class="mt-1 text-sm text-red-600">{{ errors.confirmNewPassword }}</p>
        </div>

        <!-- Password Requirements -->
        <div class="bg-blue-50 rounded-lg p-4">
          <h4 class="text-sm font-medium text-blue-700 mb-2">Password Requirements:</h4>
          <ul class="text-xs text-blue-700 space-y-1">
            <li class="flex items-center">
              <i 
                :class="newPassword.length >= 8 ? 'fas fa-check text-green-500' : 'fas fa-times text-gray-400'"
                class="mr-2"
              ></i>
              At least 8 characters
            </li>
            <li class="flex items-center">
              <i 
                :class="/[A-Z]/.test(newPassword) ? 'fas fa-check text-green-500' : 'fas fa-times text-gray-400'"
                class="mr-2"
              ></i>
              One uppercase letter
            </li>
            <li class="flex items-center">
              <i 
                :class="/[0-9]/.test(newPassword) ? 'fas fa-check text-green-500' : 'fas fa-times text-gray-400'"
                class="mr-2"
              ></i>
              One number
            </li>
            <li class="flex items-center">
              <i 
                :class="/[!@#$%^&*]/.test(newPassword) ? 'fas fa-check text-green-500' : 'fas fa-times text-gray-400'"
                class="mr-2"
              ></i>
              One special character (!@#$%^&*)
            </li>
          </ul>
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
            Updating...
          </span>
          <span v-else>Reset Password</span>
        </button>
      </form>

      <!-- Step 4: Success -->
      <div v-else-if="currentStep === 3" class="p-6 space-y-6">
        <div class="text-center">
          <div class="flex justify-center mb-4">
            <div class="h-16 w-16 rounded-full bg-green-100 flex items-center justify-center">
              <i class="fas fa-check text-2xl text-green-500"></i>
            </div>
          </div>
          <h2 class="text-xl font-bold text-gray-800 mb-2">Password Reset Successful</h2>
          <p class="text-gray-600">Your password has been updated successfully. You can now login with your new password.</p>
        </div>

        <button
          type="button"
          @click="goToLogin"
          class="w-full py-3 px-4 bg-accent text-white rounded-lg shadow-md hover:bg-brand-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-opacity-50 transition-all duration-300 transform hover:scale-[1.02] active:scale-[0.98]"
        >
          Continue to Login
        </button>
      </div>

      <!-- Back to Login Link -->
      <div v-if="currentStep < 3" class="px-6 py-4 bg-gray-50 border-t border-gray-200 text-center">
        <p class="text-gray-600">
          <router-link 
            to="/auth/login" 
            class="font-medium text-accent hover:text-accent/80 transition-colors"
          >
            Back to Login
          </router-link>
        </p>
      </div>
    </div>

    <!-- CAPTCHA container (hidden) -->
    <div id="captcha-container" class="hidden"></div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'

// Temporarily mock the auth store until it's implemented
const useAuthStore = () => ({
  requestPasswordReset: async (_: any) => ({}),
  verifyResetCode: async (_: any) => ({}),
  resetPassword: async (_: any) => ({})
})

// Steps
const steps = ['Email', 'Verification', 'New Password', 'Complete']
const currentStep = ref(0)

// Component state
const email = ref('')
const verificationCode = ref('')
const newPassword = ref('')
const confirmNewPassword = ref('')
const showPassword = ref(false)
const isLoading = ref(false)
const isResending = ref(false)
const errorMessage = ref('')
const mounted = ref(false)
const resendCooldown = ref(0)
const resendInterval = ref<number | null>(null)

// Validation errors
const errors = reactive({
  email: '',
  verificationCode: '',
  newPassword: '',
  confirmNewPassword: ''
})

// Hooks
const router = useRouter()
const authStore = useAuthStore()

// Mount animation
onMounted(() => {
  setTimeout(() => {
    mounted.value = true
  }, 100)
  
  // Initialize reCAPTCHA (this is a placeholder, you would integrate the actual service)
  initCaptcha()
})

// Clear intervals on component destroy
onUnmounted(() => {
  if (resendInterval.value) {
    clearInterval(resendInterval.value)
  }
})

// Methods
const validateEmail = (): boolean => {
  errors.email = ''
  
  if (!email.value) {
    errors.email = 'Email is required'
    return false
  } 
  
  if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email.value)) {
    errors.email = 'Please enter a valid email address'
    return false
  }
  
  return true
}

const validateVerificationCode = (): boolean => {
  errors.verificationCode = ''
  
  if (!verificationCode.value) {
    errors.verificationCode = 'Verification code is required'
    return false
  }
  
  if (!/^\d{6}$/.test(verificationCode.value)) {
    errors.verificationCode = 'Verification code must be 6 digits'
    return false
  }
  
  return true
}

const validateNewPassword = (): boolean => {
  errors.newPassword = ''
  errors.confirmNewPassword = ''
  
  if (!newPassword.value) {
    errors.newPassword = 'New password is required'
    return false
  }
  
  if (newPassword.value.length < 8) {
    errors.newPassword = 'Password must be at least 8 characters'
    return false
  }
  
  if (!/[A-Z]/.test(newPassword.value)) {
    errors.newPassword = 'Password must contain at least one uppercase letter'
    return false
  }
  
  if (!/[0-9]/.test(newPassword.value)) {
    errors.newPassword = 'Password must contain at least one number'
    return false
  }
  
  if (!/[!@#$%^&*]/.test(newPassword.value)) {
    errors.newPassword = 'Password must contain at least one special character (!@#$%^&*)'
    return false
  }
  
  if (newPassword.value !== confirmNewPassword.value) {
    errors.confirmNewPassword = 'Passwords do not match'
    return false
  }
  
  return true
}

const requestResetCode = async () => {
  if (!validateEmail()) return
  
  try {
    isLoading.value = true
    errorMessage.value = ''
    
    // Execute reCAPTCHA verification (this is a placeholder)
    const captchaToken = await executeCaptcha()
    if (!captchaToken) {
      errorMessage.value = 'CAPTCHA verification failed. Please try again.'
      return
    }
    
    // Call authentication service
    await authStore.requestPasswordReset({
      email: email.value,
      captchaToken
    })
    
    // Start resend cooldown
    startResendCooldown()
    
    // Move to next step
    currentStep.value = 1
  } catch (error: any) {
    errorMessage.value = error.message || 'Failed to send reset code. Please try again.'
  } finally {
    isLoading.value = false
  }
}

const verifyResetCode = async () => {
  if (!validateVerificationCode()) return
  
  try {
    isLoading.value = true
    errorMessage.value = ''
    
    // Call authentication service
    await authStore.verifyResetCode({
      email: email.value,
      code: verificationCode.value
    })
    
    // Move to next step
    currentStep.value = 2
  } catch (error: any) {
    errorMessage.value = error.message || 'Invalid verification code. Please try again.'
  } finally {
    isLoading.value = false
  }
}

const resetPassword = async () => {
  if (!validateNewPassword()) return
  
  try {
    isLoading.value = true
    errorMessage.value = ''
    
    // Call authentication service
    await authStore.resetPassword({
      email: email.value,
      code: verificationCode.value,
      newPassword: newPassword.value
    })
    
    // Move to next step
    currentStep.value = 3
  } catch (error: any) {
    errorMessage.value = error.message || 'Failed to reset password. Please try again.'
  } finally {
    isLoading.value = false
  }
}

const resendCode = async () => {
  if (!validateEmail()) return
  
  try {
    isResending.value = true
    errorMessage.value = ''
    
    // Execute reCAPTCHA verification (this is a placeholder)
    const captchaToken = await executeCaptcha()
    if (!captchaToken) {
      errorMessage.value = 'CAPTCHA verification failed. Please try again.'
      return
    }
    
    // Call authentication service
    await authStore.requestPasswordReset({
      email: email.value,
      captchaToken
    })
    
    // Start resend cooldown
    startResendCooldown()
  } catch (error: any) {
    errorMessage.value = error.message || 'Failed to resend code. Please try again.'
  } finally {
    isResending.value = false
  }
}

const startResendCooldown = () => {
  // Start with 60 seconds cooldown
  resendCooldown.value = 60
  
  // Clear any existing interval
  if (resendInterval.value) {
    clearInterval(resendInterval.value)
  }
  
  // Set up a new interval
  resendInterval.value = setInterval(() => {
    if (resendCooldown.value > 0) {
      resendCooldown.value--
    } else {
      // Clear the interval when cooldown reaches 0
      if (resendInterval.value) {
        clearInterval(resendInterval.value)
        resendInterval.value = null
      }
    }
  }, 1000) as unknown as number
}

const goToLogin = () => {
  router.push('/auth/login')
}

// CAPTCHA functions (placeholders - you would implement with actual service)
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