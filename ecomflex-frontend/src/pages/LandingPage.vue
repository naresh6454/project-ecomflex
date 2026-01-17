<template>
  <div class="min-h-screen">
    <!-- Navigation Bar -->
    <nav class="bg-white shadow-sm">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex items-center">
            <div class="flex-shrink-0 flex items-center">
              <!-- Logo -->
              <img src="@/assets/ecomflex-logo.png" alt="Ecomflex Logo" class="h-10 w-auto" />
              <span class="ml-2 text-xl font-bold text-gray-800">Ecomflex</span>
            </div>
          </div>
          <div class="flex items-center space-x-4">
            <router-link to="/auth/login" class="btn-secondary">
              Login
            </router-link>
            <router-link to="/auth/register" class="btn-primary">
              Register
            </router-link>
          </div>
        </div>
      </div>
    </nav>

    <!-- Hero Section -->
    <div class="relative overflow-hidden bg-white">
      <div class="max-w-7xl mx-auto">
        <div class="relative z-10 pb-8 sm:pb-16 md:pb-20 lg:w-full lg:pb-28 xl:pb-32">
          <main class="mt-10 mx-auto max-w-7xl px-4 sm:mt-12 sm:px-6 md:mt-16 lg:mt-20 lg:px-8 xl:mt-28">
            <div class="text-center lg:text-left">
              <h1 
                v-motion
                :initial="{ opacity: 0, y: 20 }"
                :enter="{ opacity: 1, y: 0, transition: { delay: 200, duration: 800 } }"
                class="text-4xl tracking-tight font-extrabold text-gray-900 sm:text-5xl md:text-6xl"
              >
                <span class="block">Get</span>
                <span class="block text-brand-600">100% Cashback</span>
                <span class="block">on Products!</span>
              </h1>
              <p 
                v-motion
                :initial="{ opacity: 0 }"
                :enter="{ opacity: 1, transition: { delay: 800, duration: 800 } }"
                class="mt-3 text-base text-gray-500 sm:mt-5 sm:text-lg sm:max-w-xl sm:mx-auto md:mt-5 md:text-xl lg:mx-0"
              >
                Join our product promotion platform where you can try products from top brands and get 100% cashback after verification.
              </p>
              <div 
                v-motion
                :initial="{ opacity: 0, scale: 0.9 }"
                :enter="{ opacity: 1, scale: 1, transition: { delay: 1000, duration: 800 } }"
                class="mt-5 sm:mt-8 sm:flex sm:justify-center lg:justify-start"
              >
                <div class="rounded-md shadow">
                  <router-link to="/auth/register" class="w-full flex items-center justify-center px-8 py-3 border border-transparent text-base font-medium rounded-xl text-white bg-brand-600 hover:bg-brand-700 md:py-4 md:text-lg md:px-10 transition-all duration-300 ripple">
                    Get Started
                  </router-link>
                </div>
                <div class="mt-3 sm:mt-0 sm:ml-3">
                  <a href="#features" class="w-full flex items-center justify-center px-8 py-3 border border-transparent text-base font-medium rounded-xl text-brand-600 bg-white hover:bg-gray-50 border-brand-100 md:py-4 md:text-lg md:px-10 transition-all duration-300 ripple">
                    Learn More
                  </a>
                </div>
              </div>
            </div>
          </main>
        </div>
      </div>
      <div class="lg:absolute lg:inset-y-0 lg:right-0 lg:w-1/2">
        <div 
          v-motion
          :initial="{ opacity: 0, x: 20 }"
          :enter="{ opacity: 1, x: 0, transition: { delay: 1200, duration: 800 } }"
          class="h-56 w-full bg-brand-50 sm:h-72 md:h-96 lg:w-full lg:h-full p-8 flex items-center justify-center"
        >
          <div class="relative h-full w-full rounded-2xl overflow-hidden shadow-lg">
            <div class="absolute inset-0 bg-gradient-to-br from-brand-400 to-brand-600 opacity-90"></div>
            <!-- Decorative elements -->
            <div class="absolute top-1/4 left-1/4 w-20 h-20 rounded-full bg-accent opacity-50"></div>
            <div class="absolute bottom-1/3 right-1/3 w-32 h-32 rounded-full bg-accent opacity-30"></div>
            <div class="absolute inset-0 flex flex-col items-center justify-center text-white">
              <div class="text-2xl md:text-4xl font-bold mb-4">Amazing Products</div>
              <div class="flex space-x-4">
                <div class="w-16 h-16 bg-white rounded-xl shadow-lg flex items-center justify-center animate-bounce-subtle">
                  <div class="w-12 h-12 bg-accent rounded-lg"></div>
                </div>
                <div class="w-16 h-16 bg-white rounded-xl shadow-lg flex items-center justify-center animate-bounce-subtle delay-200">
                  <div class="w-12 h-12 bg-brand-200 rounded-lg"></div>
                </div>
                <div class="w-16 h-16 bg-white rounded-xl shadow-lg flex items-center justify-center animate-bounce-subtle delay-400">
                  <div class="w-12 h-12 bg-brand-300 rounded-lg"></div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Features Section -->
    <section id="features" class="py-12 bg-brand-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center">
          <h2 class="text-3xl font-extrabold text-gray-900 sm:text-4xl">
            How It Works
          </h2>
          <p class="mt-4 max-w-2xl text-xl text-gray-500 mx-auto">
            Simple process to get products with 100% cashback
          </p>
        </div>

        <div class="mt-12">
          <div class="grid grid-cols-1 gap-8 sm:grid-cols-2 lg:grid-cols-3">
            <!-- Step 1 -->
            <div 
              v-motion
              :initial="{ opacity: 0, y: 20 }"
              :enter="{ opacity: 1, y: 0, transition: { delay: 200, duration: 600 } }"
              class="card"
            >
              <div class="h-12 w-12 rounded-md bg-brand-500 text-white flex items-center justify-center mb-5">
                <span class="text-lg font-bold">1</span>
              </div>
              <h3 class="text-lg font-medium text-gray-900">Browse Products</h3>
              <p class="mt-2 text-base text-gray-500">
                Explore our curated selection of products from top brands.
              </p>
            </div>

            <!-- Step 2 -->
            <div 
              v-motion
              :initial="{ opacity: 0, y: 20 }"
              :enter="{ opacity: 1, y: 0, transition: { delay: 400, duration: 600 } }"
              class="card"
            >
              <div class="h-12 w-12 rounded-md bg-brand-500 text-white flex items-center justify-center mb-5">
                <span class="text-lg font-bold">2</span>
              </div>
              <h3 class="text-lg font-medium text-gray-900">Book & Purchase</h3>
              <p class="mt-2 text-base text-gray-500">
                Purchase the product through our merchant link and upload your proof.
              </p>
            </div>

            <!-- Step 3 -->
            <div 
              v-motion
              :initial="{ opacity: 0, y: 20 }"
              :enter="{ opacity: 1, y: 0, transition: { delay: 600, duration: 600 } }"
              class="card"
            >
              <div class="h-12 w-12 rounded-md bg-brand-500 text-white flex items-center justify-center mb-5">
                <span class="text-lg font-bold">3</span>
              </div>
              <h3 class="text-lg font-medium text-gray-900">Get Cashback</h3>
              <p class="mt-2 text-base text-gray-500">
                After verification, receive 100% cashback for your purchase.
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- User Types Section -->
    <section class="py-12 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center">
          <h2 class="text-3xl font-extrabold text-gray-900 sm:text-4xl">
            Join as a
          </h2>
        </div>

        <div class="mt-12">
          <div class="grid grid-cols-1 gap-8 md:grid-cols-2">
            <!-- Public User -->
            <div 
              v-motion
              :initial="{ opacity: 0, x: -20 }"
              :enter="{ opacity: 1, x: 0, transition: { delay: 200, duration: 600 } }"
              class="card border-2 border-brand-100 hover:border-brand-200"
            >
              <div class="flex justify-between items-start">
                <div class="h-12 w-12 rounded-full bg-brand-100 text-brand-600 flex items-center justify-center mb-5">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                  </svg>
                </div>
                <span class="badge-blue">Popular</span>
              </div>
              <h3 class="text-xl font-bold text-gray-900">Public User</h3>
              <ul class="mt-4 space-y-3">
                <li class="flex items-start">
                  <svg class="h-6 w-6 text-green-500 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                  </svg>
                  <span>Browse and book products</span>
                </li>
                <li class="flex items-start">
                  <svg class="h-6 w-6 text-green-500 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                  </svg>
                  <span>Upload purchase proofs</span>
                </li>
                <li class="flex items-start">
                  <svg class="h-6 w-6 text-green-500 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                  </svg>
                  <span>Get 100% cashback after verification</span>
                </li>
              </ul>
              <div class="mt-8">
                <router-link to="/auth/register" class="btn-primary w-full justify-center">
                  Register as Public User
                </router-link>
              </div>
            </div>

            <!-- Influencer -->
            <div 
              v-motion
              :initial="{ opacity: 0, x: 20 }"
              :enter="{ opacity: 1, x: 0, transition: { delay: 400, duration: 600 } }"
              class="card border-2 border-accent/30 hover:border-accent"
            >
              <div class="flex justify-between items-start">
                <div class="h-12 w-12 rounded-full bg-amber-100 text-amber-600 flex items-center justify-center mb-5">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z" />
                  </svg>
                </div>
                <span class="badge-gold">Premium</span>
              </div>
              <h3 class="text-xl font-bold text-gray-900">Influencer</h3>
              <ul class="mt-4 space-y-3">
                <li class="flex items-start">
                  <svg class="h-6 w-6 text-green-500 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                  </svg>
                  <span>All Public User benefits</span>
                </li>
                <li class="flex items-start">
                  <svg class="h-6 w-6 text-green-500 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                  </svg>
                  <span>Create unique referral codes</span>
                </li>
                <li class="flex items-start">
                  <svg class="h-6 w-6 text-green-500 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                  </svg>
                  <span>Track referrals & earn bonus rewards</span>
                </li>
              </ul>
              <div class="mt-8">
                <router-link to="/auth/register?type=influencer" class="btn-accent w-full justify-center">
                  Apply as Influencer
                </router-link>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Footer -->
    <footer class="bg-brand-800 text-white">
      <div class="max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div>
            <div class="flex items-center">
              <!-- Footer Logo -->
              <img src="@/assets/ecomflex-logo.png" alt="Ecomflex Logo" class="h-10 w-auto brightness-150" />
              <span class="ml-2 text-xl font-bold text-white">Ecomflex</span>
            </div>
            <p class="mt-4 text-brand-200">
              Product promotion platform with 100% cashback guarantee.
            </p>
          </div>
          <div>
            <h3 class="text-lg font-semibold text-brand-100">Quick Links</h3>
            <ul class="mt-4 space-y-2">
              <li><a href="#" class="text-brand-300 hover:text-white transition-colors duration-200">Home</a></li>
              <li><a href="#" class="text-brand-300 hover:text-white transition-colors duration-200">About</a></li>
              <li><a href="#" class="text-brand-300 hover:text-white transition-colors duration-200">Products</a></li>
              <li><a href="#" class="text-brand-300 hover:text-white transition-colors duration-200">Contact</a></li>
            </ul>
          </div>
          <div>
            <h3 class="text-lg font-semibold text-brand-100">Contact</h3>
            <ul class="mt-4 space-y-2">
              <li class="flex items-center text-brand-300">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                </svg>
                <span>support@ecomflex.com</span>
              </li>
              <li class="flex items-center text-brand-300">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z" />
                </svg>
                <span>+1 (555) 123-4567</span>
              </li>
            </ul>
          </div>
        </div>
        <div class="mt-8 pt-8 border-t border-brand-700">
          <p class="text-brand-400 text-center text-sm">
            &copy; {{ new Date().getFullYear() }} Ecomflex. All rights reserved.
          </p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
// Landing page component
</script>

<style>
/* Add or update custom Tailwind CSS variables for the theme */
:root {
  --accent: #D4AF37;
  --brand-50: #FFF8E1;
  --brand-100: #FFECB3;
  --brand-200: #FFE082;
  --brand-300: #FFD54F;
  --brand-400: #FFCA28;
  --brand-500: #FFC107;
  --brand-600: #FFB300;
  --brand-700: #FFA000;
  --brand-800: #FF8F00;
  --brand-900: #FF6F00;
}

/* Define animation for the product boxes */
.animate-bounce-subtle {
  animation: bounce-subtle 3s infinite;
}
.delay-200 {
  animation-delay: 0.2s;
}
.delay-400 {
  animation-delay: 0.4s;
}

@keyframes bounce-subtle {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-10px);
  }
}

/* Common button styles */
.btn-primary {
  display: inline-flex;
  padding: 0.5rem 1rem;
  font-weight: 500;
  color: white;
  background-color: var(--brand-600);
  border-radius: 0.5rem;
  transition-property: all;
  transition-duration: 300ms;
  transform: translateZ(0);
}
.btn-primary:hover {
  background-color: var(--brand-700);
  transform: scale(1.02);
}
.btn-primary:active {
  transform: scale(0.98);
}
.btn-primary:focus {
  outline: none;
  box-shadow: 0 0 0 2px rgba(var(--brand-600), 0.5);
}

.btn-secondary {
  display: inline-flex;
  padding: 0.5rem 1rem;
  font-weight: 500;
  color: var(--brand-600);
  background-color: white;
  border: 1px solid var(--brand-200);
  border-radius: 0.5rem;
  transition-property: all;
  transition-duration: 300ms;
  transform: translateZ(0);
}
.btn-secondary:hover {
  background-color: #f9fafb;
  transform: scale(1.02);
}
.btn-secondary:active {
  transform: scale(0.98);
}
.btn-secondary:focus {
  outline: none;
  box-shadow: 0 0 0 2px rgba(var(--brand-600), 0.5);
}

.btn-accent {
  display: inline-flex;
  padding: 0.5rem 1rem;
  font-weight: 500;
  color: white;
  background-color: var(--accent);
  border-radius: 0.5rem;
  transition-property: all;
  transition-duration: 300ms;
  transform: translateZ(0);
}
.btn-accent:hover {
  background-color: #b8860b;
  transform: scale(1.02);
}
.btn-accent:active {
  transform: scale(0.98);
}
.btn-accent:focus {
  outline: none;
  box-shadow: 0 0 0 2px rgba(var(--accent), 0.5);
}

/* Card styles */
.card {
  background-color: white;
  padding: 1.5rem;
  border-radius: 0.75rem;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
  transition-property: all;
  transition-duration: 300ms;
}
.card:hover {
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
}

/* Badge styles */
.badge-blue {
  display: inline-block;
  padding: 0.25rem 0.5rem;
  font-size: 0.75rem;
  font-weight: 500;
  background-color: #dbeafe;
  color: #2563eb;
  border-radius: 9999px;
}
.badge-gold {
  display: inline-block;
  padding: 0.25rem 0.5rem;
  font-size: 0.75rem;
  font-weight: 500;
  background-color: #fef3c7;
  color: #d97706;
  border-radius: 9999px;
}

/* Ripple effect for buttons */
.ripple {
  position: relative;
  overflow: hidden;
}
.ripple:after {
  content: "";
  background: rgba(255, 255, 255, 0.3);
  display: block;
  position: absolute;
  border-radius: 50%;
  padding-top: 240%;
  padding-left: 240%;
  margin-top: -120%;
  margin-left: -120%;
  opacity: 0;
  transition: all 0.8s;
}
.ripple:active:after {
  padding-top: 0;
  padding-left: 0;
  margin-top: 0;
  margin-left: 0;
  opacity: 1;
  transition: 0s;
}
</style>