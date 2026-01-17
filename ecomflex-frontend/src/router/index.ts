import { createRouter, createWebHistory, RouteRecordRaw, RouteLocationNormalized, NavigationGuardNext } from 'vue-router'

// Define global types
declare global {
  interface Window {
    gtag: (command: string, id: string, config: any) => void;
  }
}

// Route definitions
const routes: Array<RouteRecordRaw> = [
  // Landing and Auth routes
  {
    path: '/',
    name: 'landing',
    component: () => import('../pages/LandingPage.vue'),
    meta: { title: 'Ecomflex - Product Promotion Platform', requiresAuth: false, cacheable: true }
  },
  {
    path: '/auth/login',
    name: 'login',
    component: () => import('../pages/auth/Login.vue'),
    meta: { title: 'Login - Ecomflex', requiresAuth: false, cacheable: true }
  },
  {
    path: '/auth/register',
    name: 'register',
    component: () => import('../pages/auth/Register.vue'),
    meta: { title: 'Register - Ecomflex', requiresAuth: false, cacheable: true }
  },
  {
    path: '/auth/forgot-password',
    name: 'forgot-password',
    component: () => import('../pages/auth/ForgotPassword.vue'),
    meta: { title: 'Forgot Password - Ecomflex', requiresAuth: false, cacheable: true }
  },
  
  // Public User routes (Phase 2)
  {
    path: '/home',
    name: 'home',
    component: () => import('../pages/public/Home.vue'),
    meta: { title: 'Products - Ecomflex', requiresAuth: true, cacheable: true }
  },
  {
    path: '/product/:id',
    name: 'product-detail',
    component: () => import('../pages/public/ProductDetail.vue'),
    meta: { title: 'Product Details - Ecomflex', requiresAuth: true, cacheable: true }
  },
  {
    path: '/my-bookings',
    name: 'my-bookings',
    component: () => import('../pages/public/MyBookings.vue'),
    meta: { title: 'My Bookings - Ecomflex', requiresAuth: true, cacheable: true }
  },
  
  // Influencer routes (Phase 3)
  {
    path: '/influencer',
    component: () => import('../components/layouts/InfluencerLayout.vue'),
    meta: { requiresAuth: true, role: 'influencer', cacheable: true },
    children: [
      {
        path: 'dashboard',
        name: 'influencer-dashboard',
        component: () => import('../pages/influencer/Dashboard.vue'),
        meta: { title: 'Influencer Dashboard - Ecomflex', cacheable: true }
      },
      {
        path: 'referrals',
        name: 'influencer-referrals',
        component: () => import('../pages/influencer/Referrals.vue'),
        meta: { title: 'Referral Tracking - Ecomflex', cacheable: true }
      },
      {
        path: 'products',
        name: 'influencer-products',
        component: () => import('../pages/influencer/Products.vue'),
        meta: { title: 'Products to Promote - Ecomflex', cacheable: true }
      },
      {
        path: 'earnings',
        name: 'influencer-earnings',
        component: () => import('../pages/influencer/Earnings.vue'),
        meta: { title: 'Earnings & Payouts - Ecomflex', cacheable: true }
      },
      {
        path: 'profile',
        name: 'influencer-profile',
        component: () => import('../pages/influencer/Profile.vue'),
        meta: { title: 'Influencer Profile - Ecomflex', cacheable: true }
      },
      {
        path: 'settings',
        name: 'influencer-settings',
        component: () => import('../pages/influencer/Settings.vue'),
        meta: { title: 'Influencer Settings - Ecomflex', cacheable: false }
      }
    ]
  },
  
  // Super Admin routes (Phase 4)
  {
    path: '/admin',
    component: () => import('../components/layouts/AdminLayout.vue'),
    meta: { requiresAuth: true, role: 'admin', cacheable: true },
    children: [
      {
        path: 'dashboard',
        name: 'admin-dashboard',
        component: () => import('../pages/superadmin/Dashboard.vue'),
        meta: { title: 'Admin Dashboard - Ecomflex', cacheable: true }
      },
      {
        path: 'products',
        name: 'admin-products',
        component: () => import('../pages/superadmin/Products.vue'),
        meta: { title: 'Manage Products - Ecomflex', cacheable: true }
      },
      {
        path: 'proofs',
        name: 'admin-proofs',
        component: () => import('../pages/superadmin/Proofs.vue'),
        meta: { title: 'Verify Proofs - Ecomflex', cacheable: false }
      },
      {
        path: 'influencers',
        name: 'admin-influencers',
        component: () => import('../pages/superadmin/InfluencerRequests.vue'),
        meta: { title: 'Influencer Requests - Ecomflex', cacheable: false }
      }
    ]
  },
  
  // Redirect influencer root to dashboard
  {
    path: '/influencer',
    redirect: '/influencer/dashboard'
  },
  
  // Redirect admin root to dashboard
  {
    path: '/admin',
    redirect: '/admin/dashboard'
  },
  
  // Offline fallback route
  {
    path: '/offline',
    name: 'offline',
    component: () => import('../pages/Offline.vue'),
    meta: { title: 'Offline - Ecomflex', requiresAuth: false, cacheable: true }
  },
  
  // 404 Page
  {
    path: '/:pathMatch(.*)*',
    name: 'not-found',
    component: () => import('../pages/NotFound.vue'),
    meta: { title: 'Page Not Found - Ecomflex', cacheable: true }
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
  // Smooth scroll to top on route change
  scrollBehavior(_to: RouteLocationNormalized, _from: RouteLocationNormalized, savedPosition: { left: number, top: number } | null) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { top: 0, behavior: 'smooth' as ScrollBehavior }
    }
  }
})

// Navigation Guards
router.beforeEach((to: RouteLocationNormalized, _from: RouteLocationNormalized, next: NavigationGuardNext) => {
  // Update document title
  document.title = to.meta.title as string || 'Ecomflex'
  
  console.log('Router: Navigating to:', to.path)
  console.log('Router: Token exists:', !!localStorage.getItem('token'))
  console.log('Router: UserRole:', localStorage.getItem('userRole'))
  
  // Check if we're offline
  if (!navigator.onLine && to.name !== 'offline') {
    // Only redirect to offline page for non-cacheable routes
    if (to.meta.cacheable !== true) {
      return next({ name: 'offline', query: { redirect: to.fullPath } })
    }
  }
  
  // Check if route requires authentication
  const requiresAuth = to.matched.some((record: RouteRecordRaw) => record.meta?.requiresAuth)
  const isAuthenticated = localStorage.getItem('token') !== null
  
  console.log('Router: RequiresAuth:', requiresAuth, 'IsAuthenticated:', isAuthenticated)
  
  if (requiresAuth && !isAuthenticated) {
    // Redirect to login page if not authenticated
    console.log('Router: Redirecting to login - not authenticated')
    next({ name: 'login', query: { redirect: to.fullPath } })
  } else if (to.matched.some((record: RouteRecordRaw) => record.meta?.role)) {
    // Check user role for role-protected routes
    const userRole = localStorage.getItem('userRole')
    const routeWithRole = to.matched.find((record: RouteRecordRaw) => record.meta?.role)
    const requiredRole = routeWithRole?.meta?.role as string | undefined
    
    console.log('Router: UserRole:', userRole, 'RequiredRole:', requiredRole)
    
    if (userRole === requiredRole) {
      console.log('Router: Role matches, proceeding to:', to.path)
      next()
    } else {
      // Redirect to appropriate dashboard based on role
      console.log('Router: Role mismatch, redirecting based on role')
      if (userRole === 'admin') {
        next({ name: 'admin-dashboard' })
      } else if (userRole === 'influencer') {
        next({ name: 'influencer-dashboard' })
      } else {
        next({ name: 'home' })
      }
    }
  } else {
    // Proceed as normal
    console.log('Router: No special requirements, proceeding to:', to.path)
    next()
  }
})

// After-navigation hook to track page views and cache routes
router.afterEach((to: RouteLocationNormalized) => {
  // Track page view for analytics (if online)
  if (navigator.onLine && window.gtag) {
    window.gtag('config', 'G-XXXXXXXXXX', {
      page_path: to.fullPath,
    })
  }
  
  // Cache current route if it's cacheable (useful for PWA)
  if (to.meta.cacheable && typeof caches !== 'undefined') {
    // Add route to runtime cache if needed
    try {
      const cacheName = 'ecomflex-pages'
      const request = new Request(location.href)
      
      caches.open(cacheName).then((cache) => {
        cache.add(request).catch((error) => {
          console.error('Failed to cache route:', error)
        })
      })
    } catch (error) {
      console.error('Error caching route:', error)
    }
  }
})

// Add types for global objects
// This needs to be placed before any usage of these types
declare global {
  interface Window {
    gtag: (command: string, id: string, config: any) => void;
  }
}

// Global error handling for navigation
router.onError((error: Error) => {
  console.error('Navigation error:', error)
  
  // Redirect to error page for chunk loading failures (common in PWAs with cached code)
  if (error.message.includes('Failed to fetch dynamically imported module') || 
      error.message.includes('Chunkload failed')) {
    // This likely means we have an outdated service worker or cached module
    console.log('Chunk loading error detected, but NOT reloading to prevent reload loops')
    console.log('Error details:', error)
    
    // Disabled automatic reload to prevent reload loops
    // If you see this error consistently, it means there's a module loading issue
  }
})

export default router