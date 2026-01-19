import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { MotionPlugin } from '@vueuse/motion'
import App from './App.vue'
import router from './router'
import Toast from './components/ui/Toast.vue'
import './styles/base.css'
import './styles/animations.css'
import './styles/global.css' // Import the new global CSS
import './services/urlInterceptor';

// Note: Do NOT hardcode mock authentication data here
// Let the actual login process set the user role from the backend
// This ensures that the correct user role is maintained across page reloads
// If you need to test as admin, please login with admin credentials instead

// Import PWA initialization - DISABLED FOR DEVELOPMENT
// import { initPWA } from './pwa/register'

// Initialize PWA functionality - DISABLED FOR DEVELOPMENT
// initPWA()

// Create Vue app
const app = createApp(App)

// Use Pinia for state management
const pinia = createPinia()
app.use(pinia)

// Use router
app.use(router)

// Use MotionPlugin from @vueuse/motion
app.use(MotionPlugin)

// Register global components
app.component('Toast', Toast)

// Apply global directives
app.directive('click-outside', {
  beforeMount(el, binding) {
    el.clickOutsideEvent = (event: Event) => {
      if (!(el === event.target || el.contains(event.target as Node))) {
        binding.value(event)
      }
    }
    document.addEventListener('click', el.clickOutsideEvent)
  },
  unmounted(el) {
    document.removeEventListener('click', el.clickOutsideEvent)
  }
})

// Add ripple effect directive
app.directive('ripple', {
  mounted(el, binding) {
    el.addEventListener('click', (e: MouseEvent) => {
      const rect = el.getBoundingClientRect()
      const x = e.clientX - rect.left
      const y = e.clientY - rect.top
     
      const circle = document.createElement('span')
      const diameter = Math.max(rect.width, rect.height)
      const radius = diameter / 2
     
      // Get ripple color from binding value or default to golden
      const color = binding.value?.color || 'rgba(255, 215, 0, 0.3)'
     
      circle.style.width = circle.style.height = `${diameter}px`
      circle.style.left = `${x - radius}px`
      circle.style.top = `${y - radius}px`
      circle.style.position = 'absolute'
      circle.style.borderRadius = '50%'
      circle.style.backgroundColor = color
      circle.style.transform = 'scale(0)'
      circle.style.animation = 'ripple 0.6s linear'
     
      el.style.position = 'relative'
      el.style.overflow = 'hidden'
     
      el.appendChild(circle)
     
      setTimeout(() => {
        circle.remove()
      }, 600)
    })
  }
})

// Add offline detection
app.provide('isOnline', navigator.onLine)
window.addEventListener('online', () => {
  app.provide('isOnline', true)
})
window.addEventListener('offline', () => {
  app.provide('isOnline', false)
})

// Add a global error handler
app.config.errorHandler = (err, vm, info) => {
  // Log the error but don't crash the app
  console.error('Global error:', err)
  console.error('Component:', vm)
  console.error('Info:', info)
 
  // Here you could also report errors to a service like Sentry
}

// Mount the app
app.mount('#app')

// Add CSS animation for ripple effect to the document
const style = document.createElement('style')
style.textContent = `
  @keyframes ripple {
    to {
      transform: scale(2);
      opacity: 0;
    }
  }
`
document.head.appendChild(style)