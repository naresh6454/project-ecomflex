<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <!-- Page Header -->
    <div 
      v-motion
      :initial="{ opacity: 0, y: -20 }"
      :enter="{ opacity: 1, y: 0 }"
      class="md:flex md:items-center md:justify-between mb-8">
      <div class="flex-1 min-w-0">
        <h1 class="text-2xl font-bold leading-7 text-gray-900 sm:text-3xl sm:truncate">
          Profile
        </h1>
        <p class="mt-1 text-sm text-gray-500">
          Manage your profile information and public details.
        </p>
      </div>
      <div class="mt-4 flex md:mt-0 md:ml-4">
        <button 
          type="button" 
          class="ml-3 inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors">
          Save Changes
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 gap-y-8 gap-x-8 md:grid-cols-3">
      <!-- Left Column: Profile Photo & Social Links -->
      <div 
        v-motion
        :initial="{ opacity: 0, x: -20 }"
        :enter="{ opacity: 1, x: 0 }">
        <div class="bg-white shadow overflow-hidden sm:rounded-lg mb-6">
          <div class="px-4 py-5 sm:p-6">
            <div class="flex flex-col items-center">
              <div class="relative mb-4">
                <img 
                  class="h-32 w-32 rounded-full object-cover"
                  src="https://images.unsplash.com/photo-1550525811-e5869dd03032?ixlib=rb-1.2.1&auto=format&fit=facearea&facepad=2.5&w=256&h=256&q=80" 
                  alt="Profile photo"
                />
                <button 
                  type="button"
                  class="absolute bottom-0 right-0 rounded-full bg-blue-600 p-2 text-white shadow-lg hover:bg-blue-700 focus:outline-none transition-colors">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9a2 2 0 012-2h.93a2 2 0 001.664-.89l.812-1.22A2 2 0 0110.07 4h3.86a2 2 0 011.664.89l.812 1.22A2 2 0 0018.07 7H19a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2V9z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 13a3 3 0 11-6 0 3 3 0 016 0z" />
                  </svg>
                </button>
              </div>
              <h3 class="mt-2 text-xl font-medium text-gray-900">{{ profile.name }}</h3>
              <p class="text-sm text-gray-500">{{ profile.email }}</p>
              
              <div class="flex items-center mt-3">
                <span 
                  class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                  Influencer
                </span>
                <span 
                  class="ml-2 inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
                  Verified
                </span>
              </div>
            </div>
            
            <div class="mt-6">
              <h4 class="text-sm font-medium text-gray-900">Bio</h4>
              <div class="mt-2">
                <textarea 
                  rows="4" 
                  class="shadow-sm focus:ring-blue-500 focus:border-blue-500 block w-full sm:text-sm border-gray-300 rounded-md"
                  placeholder="Write a short bio about yourself..."
                  v-model="profile.bio"></textarea>
              </div>
            </div>
            
            <div class="mt-6">
              <h4 class="text-sm font-medium text-gray-900">Social Media Profiles</h4>
              <div class="mt-2 space-y-3">
                <div v-for="social in socialProfiles" :key="social.id" class="flex">
                  <div class="w-8 flex-shrink-0 flex items-center justify-center">
                    <component :is="social.icon" class="h-5 w-5 text-gray-400" />
                  </div>
                  <div class="ml-2 flex-1">
                    <input 
                      type="text" 
                      :placeholder="social.placeholder"
                      v-model="social.value"
                      class="focus:ring-blue-500 focus:border-blue-500 block w-full sm:text-sm border-gray-300 rounded-md" />
                  </div>
                </div>
                
                <button 
                  type="button"
                  class="mt-3 inline-flex items-center px-3 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors">
                  <svg class="-ml-0.5 mr-2 h-4 w-4 text-gray-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                  </svg>
                  Add Platform
                </button>
              </div>
            </div>
          </div>
        </div>
        
        <div class="bg-white shadow overflow-hidden sm:rounded-lg">
          <div class="px-4 py-5 sm:p-6">
            <h4 class="text-sm font-medium text-gray-900">Account Status</h4>
            <div class="mt-2 flex items-center">
              <div 
                class="w-3 h-3 rounded-full mr-2" 
                :class="isAccountActive ? 'bg-green-500' : 'bg-red-500'"></div>
              <span class="text-sm text-gray-700">
                {{ isAccountActive ? 'Active' : 'Inactive' }}
              </span>
            </div>
            
            <div class="mt-4">
              <button 
                type="button"
                class="inline-flex items-center px-3 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors">
                <svg class="-ml-0.5 mr-2 h-4 w-4 text-gray-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 11V7a4 4 0 118 0m-4 8v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2z" />
                </svg>
                Reset Password
              </button>
            </div>
          </div>
        </div>
      </div>
      
      <!-- Right Column: Profile Details -->
      <div 
        class="md:col-span-2"
        v-motion
        :initial="{ opacity: 0, x: 20 }"
        :enter="{ opacity: 1, x: 0 }">
        <div class="bg-white shadow overflow-hidden sm:rounded-lg">
          <div class="px-4 py-5 sm:p-6">
            <h3 class="text-lg font-medium leading-6 text-gray-900">Profile Information</h3>
            <p class="mt-1 text-sm text-gray-500">Update your personal information and preferences.</p>
            
            <div class="mt-6 grid grid-cols-1 gap-y-6 gap-x-4 sm:grid-cols-6">
              <div class="sm:col-span-3">
                <label for="first-name" class="block text-sm font-medium text-gray-700">First name</label>
                <div class="mt-1">
                  <input 
                    type="text" 
                    id="first-name" 
                    v-model="profile.firstName"
                    class="shadow-sm focus:ring-blue-500 focus:border-blue-500 block w-full sm:text-sm border-gray-300 rounded-md"
                  />
                </div>
              </div>

              <div class="sm:col-span-3">
                <label for="last-name" class="block text-sm font-medium text-gray-700">Last name</label>
                <div class="mt-1">
                  <input 
                    type="text" 
                    id="last-name" 
                    v-model="profile.lastName"
                    class="shadow-sm focus:ring-blue-500 focus:border-blue-500 block w-full sm:text-sm border-gray-300 rounded-md"
                  />
                </div>
              </div>

              <div class="sm:col-span-4">
                <label for="email" class="block text-sm font-medium text-gray-700">Email address</label>
                <div class="mt-1">
                  <input 
                    type="email" 
                    id="email" 
                    v-model="profile.email"
                    class="shadow-sm focus:ring-blue-500 focus:border-blue-500 block w-full sm:text-sm border-gray-300 rounded-md"
                  />
                </div>
              </div>

              <div class="sm:col-span-3">
                <label for="phone" class="block text-sm font-medium text-gray-700">Phone number</label>
                <div class="mt-1">
                  <input 
                    type="tel" 
                    id="phone" 
                    v-model="profile.phone"
                    class="shadow-sm focus:ring-blue-500 focus:border-blue-500 block w-full sm:text-sm border-gray-300 rounded-md"
                  />
                </div>
              </div>

              <div class="sm:col-span-3">
                <label for="referral-code" class="block text-sm font-medium text-gray-700">Referral Code</label>
                <div class="mt-1">
                  <input 
                    type="text" 
                    id="referral-code" 
                    v-model="profile.referralCode"
                    class="shadow-sm focus:ring-blue-500 focus:border-blue-500 block w-full sm:text-sm border-gray-300 rounded-md"
                    readonly
                  />
                </div>
                <p class="mt-1 text-xs text-gray-500">This is your unique referral code.</p>
              </div>

              <div class="sm:col-span-6">
                <label for="categories" class="block text-sm font-medium text-gray-700">Preferred Categories</label>
                <div class="mt-1">
                  <select 
                    id="categories" 
                    multiple
                    v-model="profile.categories"
                    class="shadow-sm focus:ring-blue-500 focus:border-blue-500 block w-full sm:text-sm border-gray-300 rounded-md"
                  >
                    <option value="electronics">Electronics</option>
                    <option value="home">Home & Kitchen</option>
                    <option value="beauty">Beauty & Personal Care</option>
                    <option value="fitness">Fitness</option>
                    <option value="fashion">Fashion</option>
                    <option value="toys">Toys & Games</option>
                  </select>
                </div>
                <p class="mt-1 text-xs text-gray-500">Select categories you're interested in promoting.</p>
              </div>

              <div class="sm:col-span-6">
                <label for="audience" class="block text-sm font-medium text-gray-700">Audience Description</label>
                <div class="mt-1">
                  <textarea 
                    id="audience" 
                    rows="3" 
                    v-model="profile.audienceDescription"
                    class="shadow-sm focus:ring-blue-500 focus:border-blue-500 block w-full sm:text-sm border-gray-300 rounded-md"
                    placeholder="Describe your audience demographics, interests, etc."
                  ></textarea>
                </div>
              </div>

              <div class="sm:col-span-6">
                <div class="flex items-start">
                  <div class="flex items-center h-5">
                    <input 
                      id="newsletter" 
                      type="checkbox" 
                      v-model="profile.newsletterSubscription"
                      class="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300 rounded"
                    />
                  </div>
                  <div class="ml-3 text-sm">
                    <label for="newsletter" class="font-medium text-gray-700">Email Notifications</label>
                    <p class="text-gray-500">Receive updates about new products, promotions, and platform news.</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        
        <!-- Verification Documents Section -->
        <div class="mt-8 bg-white shadow overflow-hidden sm:rounded-lg">
          <div class="px-4 py-5 sm:p-6">
            <h3 class="text-lg font-medium leading-6 text-gray-900">Verification Documents</h3>
            <p class="mt-1 text-sm text-gray-500">Documents you've provided to verify your influencer status.</p>
            
            <div class="mt-4">
              <ul class="divide-y divide-gray-200">
                <li 
                  v-for="(document, index) in verificationDocuments" 
                  :key="document.id"
                  class="py-4"
                  v-motion
                  :initial="{ opacity: 0, y: 10 }"
                  :enter="{ opacity: 1, y: 0, transition: { delay: index * 100 } }">
                  <div class="flex items-center justify-between">
                    <div class="flex items-center">
                      <div class="flex-shrink-0">
                        <svg class="h-8 w-8 text-gray-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                        </svg>
                      </div>
                      <div class="ml-3">
                        <p class="text-sm font-medium text-gray-900">{{ document.name }}</p>
                        <p class="text-sm text-gray-500">Uploaded on {{ document.uploadDate }}</p>
                      </div>
                    </div>
                    <div class="flex">
                      <span 
                        class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                        :class="{
                          'bg-green-100 text-green-800': document.status === 'Verified',
                          'bg-yellow-100 text-yellow-800': document.status === 'Pending',
                          'bg-red-100 text-red-800': document.status === 'Rejected'
                        }">
                        {{ document.status }}
                      </span>
                      <button class="ml-4 text-sm text-blue-600 hover:text-blue-900 transition-colors">
                        View
                      </button>
                    </div>
                  </div>
                </li>
              </ul>
              
              <div class="mt-6">
                <button 
                  type="button"
                  class="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors">
                  <svg class="-ml-1 mr-2 h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
                  </svg>
                  Upload New Document
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, h } from 'vue'
import { useReferralStore } from '@/stores/referral'

// Define the social profile interface
interface SocialProfile {
  id: number;
  platform: string;
  placeholder: string;
  value: string;
  icon: (props: any) => any;
}

// Referral store
const referralStore = useReferralStore()

// Account status
const isAccountActive = ref(true)

// Profile data
const profile = ref({
  name: 'Sarah Johnson',
  firstName: 'Sarah',
  lastName: 'Johnson',
  email: 'sarah@example.com',
  phone: '+1 (555) 123-4567',
  bio: 'Tech enthusiast and lifestyle influencer with a passion for finding the best products that make life easier. I love sharing honest reviews and helping my audience discover great deals.',
  referralCode: 'SARAH25',
  categories: ['electronics', 'beauty', 'home'],
  audienceDescription: 'My audience consists primarily of young professionals (25-35) interested in tech gadgets, home decor, and personal care products. They value quality and are looking for honest reviews before making purchasing decisions.',
  newsletterSubscription: true
})

// Social media profiles with proper Vue render functions
const socialProfiles = ref<SocialProfile[]>([
  {
    id: 1,
    platform: 'Instagram',
    placeholder: 'Instagram username',
    value: '@sarahjohnson',
    icon: (props: any) => {
      return h('svg', {
        xmlns: "http://www.w3.org/2000/svg",
        fill: "currentColor",
        viewBox: "0 0 24 24",
        ...props
      }, [
        h('path', {
          d: "M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zm0-2.163c-3.259 0-3.667.014-4.947.072-4.358.2-6.78 2.618-6.98 6.98-.059 1.281-.073 1.689-.073 4.948 0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98 1.281.058 1.689.072 4.948.072 3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98-1.281-.059-1.69-.073-4.949-.073zm0 5.838c-3.403 0-6.162 2.759-6.162 6.162s2.759 6.163 6.162 6.163 6.162-2.759 6.162-6.163c0-3.403-2.759-6.162-6.162-6.162zm0 10.162c-2.209 0-4-1.79-4-4 0-2.209 1.791-4 4-4s4 1.791 4 4c0 2.21-1.791 4-4 4zm6.406-11.845c-.796 0-1.441.645-1.441 1.44s.645 1.44 1.441 1.44c.795 0 1.439-.645 1.439-1.44s-.644-1.44-1.439-1.44z"
        })
      ])
    }
  },
  {
    id: 2,
    platform: 'TikTok',
    placeholder: 'TikTok username',
    value: '@sarahj',
    icon: (props: any) => {
      return h('svg', {
        xmlns: "http://www.w3.org/2000/svg",
        fill: "currentColor",
        viewBox: "0 0 24 24",
        ...props
      }, [
        h('path', {
          d: "M19.59 6.69a4.83 4.83 0 0 1-3.77-4.25V2h-3.45v13.67a2.89 2.89 0 0 1-5.2 1.74 2.89 2.89 0 0 1 2.31-4.64 2.93 2.93 0 0 1 .88.13V9.4a6.84 6.84 0 0 0-1-.05A6.33 6.33 0 0 0 5 20.1a6.34 6.34 0 0 0 10.86-4.43v-7a8.16 8.16 0 0 0 4.77 1.52v-3.4a4.85 4.85 0 0 1-1-.1z"
        })
      ])
    }
  },
  {
    id: 3,
    platform: 'YouTube',
    placeholder: 'YouTube channel URL',
    value: 'https://youtube.com/sarahjohnson',
    icon: (props: any) => {
      return h('svg', {
        xmlns: "http://www.w3.org/2000/svg",
        fill: "currentColor",
        viewBox: "0 0 24 24",
        ...props
      }, [
        h('path', {
          d: "M23.498 6.186a3.016 3.016 0 0 0-2.122-2.136C19.505 3.545 12 3.545 12 3.545s-7.505 0-9.377.505A3.017 3.017 0 0 0 .502 6.186C0 8.07 0 12 0 12s0 3.93.502 5.814a3.016 3.016 0 0 0 2.122 2.136c1.871.505 9.376.505 9.376.505s7.505 0 9.377-.505a3.015 3.015 0 0 0 2.122-2.136C24 15.93 24 12 24 12s0-3.93-.502-5.814zM9.545 15.568V8.432L15.818 12l-6.273 3.568z"
        })
      ])
    }
  }
])

// Verification documents
const verificationDocuments = ref([
  {
    id: 1,
    name: 'Social Media Profiles Verification',
    uploadDate: 'Aug 15, 2025',
    status: 'Verified'
  },
  {
    id: 2,
    name: 'ID Verification Document',
    uploadDate: 'Aug 15, 2025',
    status: 'Verified'
  },
  {
    id: 3,
    name: 'Audience Analytics Report',
    uploadDate: 'Sept 1, 2025',
    status: 'Pending'
  }
])

// Fetch data on component mount
onMounted(async () => {
  await referralStore.fetchReferralCode()
})
</script>