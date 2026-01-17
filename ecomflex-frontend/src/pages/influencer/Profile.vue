<template>
  <div class="profile-page">
    <!-- Page Header -->
    <header class="mb-8">
      <h1 class="text-2xl md:text-3xl font-bold text-gray-900 mb-2">Influencer Profile</h1>
      <p class="text-gray-600">Manage your account and referral settings</p>
    </header>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- Left Column - Profile Info -->
      <div 
        class="lg:col-span-1"
        v-motion
        :initial="{ opacity: 0, x: -20 }"
        :enter="{ opacity: 1, x: 0, transition: { duration: 500 } }">
        <!-- Profile Card -->
        <div class="bg-white rounded-xl shadow-lg overflow-hidden mb-6">
          <div class="relative">
            <!-- Cover Photo -->
            <div class="h-32 bg-gradient-to-r from-accent/30 to-accent/10"></div>
            
            <!-- Profile Photo -->
            <div class="absolute bottom-0 left-1/2 transform -translate-x-1/2 translate-y-1/2">
              <div class="relative group">
                <div class="w-24 h-24 rounded-full border-4 border-white overflow-hidden bg-white">
                  <img 
                    v-if="profile.avatar" 
                    :src="profile.avatar" 
                    alt="Profile" 
                    class="w-full h-full object-cover" />
                  <div v-else class="w-full h-full flex items-center justify-center bg-accent/20 text-accent font-bold text-2xl">
                    {{ getInitials(profile.fullName) }}
                  </div>
                </div>
                <div 
                  class="absolute inset-0 rounded-full bg-black bg-opacity-50 flex items-center justify-center opacity-0 group-hover:opacity-100 cursor-pointer transition-opacity duration-300"
                  @click="triggerFileInput">
                  <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9a2 2 0 012-2h.93a2 2 0 001.664-.89l.812-1.22A2 2 0 0110.07 4h3.86a2 2 0 011.664.89l.812 1.22A2 2 0 0018.07 7H19a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2V9z"></path>
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 13a3 3 0 11-6 0 3 3 0 016 0z"></path>
                  </svg>
                </div>
                <input
                  type="file"
                  ref="fileInput"
                  accept="image/*"
                  class="hidden"
                  @change="handleFileUpload"
                />
              </div>
            </div>
          </div>
          
          <div class="pt-16 pb-6 px-6 text-center">
            <h2 class="text-xl font-bold text-gray-900 mb-1">{{ profile.fullName }}</h2>
            <p class="text-gray-500 text-sm mb-4">{{ profile.email }}</p>
            
            <div class="flex justify-center space-x-2 mb-4">
              <a 
                v-if="profile.social.instagram" 
                :href="'https://instagram.com/' + profile.social.instagram"
                target="_blank"
                class="w-8 h-8 rounded-full bg-gray-100 flex items-center justify-center text-gray-600 hover:bg-accent hover:text-white transition-colors duration-300">
                <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zm0-2.163c-3.259 0-3.667.014-4.947.072-4.358.2-6.78 2.618-6.98 6.98-.059 1.281-.073 1.689-.073 4.948 0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98 1.281.058 1.689.072 4.948.072 3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98-1.281-.059-1.69-.073-4.949-.073zm0 5.838c-3.403 0-6.162 2.759-6.162 6.162s2.759 6.163 6.162 6.163 6.162-2.759 6.162-6.163c0-3.403-2.759-6.162-6.162-6.162zm0 10.162c-2.209 0-4-1.79-4-4 0-2.209 1.791-4 4-4s4 1.791 4 4c0 2.21-1.791 4-4 4zm6.406-11.845c-.796 0-1.441.645-1.441 1.44s.645 1.44 1.441 1.44c.795 0 1.439-.645 1.439-1.44s-.644-1.44-1.439-1.44z" />
                </svg>
              </a>
              <a 
                v-if="profile.social.twitter" 
                :href="'https://twitter.com/' + profile.social.twitter"
                target="_blank"
                class="w-8 h-8 rounded-full bg-gray-100 flex items-center justify-center text-gray-600 hover:bg-accent hover:text-white transition-colors duration-300">
                <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path d="M23.953 4.57a10 10 0 01-2.825.775 4.958 4.958 0 002.163-2.723c-.951.555-2.005.959-3.127 1.184a4.92 4.92 0 00-8.384 4.482C7.69 8.095 4.067 6.13 1.64 3.162a4.822 4.822 0 00-.666 2.475c0 1.71.87 3.213 2.188 4.096a4.904 4.904 0 01-2.228-.616v.06a4.923 4.923 0 003.946 4.827 4.996 4.996 0 01-2.212.085 4.936 4.936 0 004.604 3.417 9.867 9.867 0 01-6.102 2.105c-.39 0-.779-.023-1.17-.067a13.995 13.995 0 007.557 2.209c9.053 0 13.998-7.496 13.998-13.985 0-.21 0-.42-.015-.63A9.935 9.935 0 0024 4.59z" />
                </svg>
              </a>
              <a 
                v-if="profile.social.tiktok" 
                :href="'https://tiktok.com/@' + profile.social.tiktok"
                target="_blank"
                class="w-8 h-8 rounded-full bg-gray-100 flex items-center justify-center text-gray-600 hover:bg-accent hover:text-white transition-colors duration-300">
                <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path d="M12.525.02c1.31-.02 2.61-.01 3.91-.02.08 1.53.63 3.09 1.75 4.17 1.12 1.11 2.7 1.62 4.24 1.79v4.03c-1.44-.05-2.89-.35-4.2-.97-.57-.26-1.1-.59-1.62-.93-.01 2.92.01 5.84-.02 8.75-.08 1.4-.54 2.79-1.35 3.94-1.31 1.92-3.58 3.17-5.91 3.21-1.43.08-2.86-.31-4.08-1.03-2.02-1.19-3.44-3.37-3.65-5.71-.02-.5-.03-1-.01-1.49.18-1.9 1.12-3.72 2.58-4.96 1.66-1.44 3.98-2.13 6.15-1.72.02 1.48-.04 2.96-.04 4.44-.99-.32-2.15-.23-3.02.37-.63.41-1.11 1.04-1.36 1.75-.21.51-.15 1.07-.14 1.61.24 1.64 1.82 3.02 3.5 2.87 1.12-.01 2.19-.66 2.77-1.61.19-.33.4-.67.41-1.06.1-1.79.06-3.57.07-5.36.01-4.03-.01-8.05.02-12.07z" />
                </svg>
              </a>
              <a 
                v-if="profile.social.youtube" 
                :href="'https://youtube.com/c/' + profile.social.youtube"
                target="_blank"
                class="w-8 h-8 rounded-full bg-gray-100 flex items-center justify-center text-gray-600 hover:bg-accent hover:text-white transition-colors duration-300">
                <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path d="M23.498 6.186a3.016 3.016 0 0 0-2.122-2.136C19.505 3.545 12 3.545 12 3.545s-7.505 0-9.377.505A3.017 3.017 0 0 0 .502 6.186C0 8.07 0 12 0 12s0 3.93.502 5.814a3.016 3.016 0 0 0 2.122 2.136c1.871.505 9.376.505 9.376.505s7.505 0 9.377-.505a3.015 3.015 0 0 0 2.122-2.136C24 15.93 24 12 24 12s0-3.93-.502-5.814zM9.545 15.568V8.432L15.818 12l-6.273 3.568z" />
                </svg>
              </a>
            </div>
            
            <div class="pt-4 border-t border-gray-100">
              <div class="flex items-center justify-between mb-2">
                <span class="text-sm text-gray-500">Member Since</span>
                <span class="text-sm font-medium">{{ formatDate(profile.joinDate) }}</span>
              </div>
              <div class="flex items-center justify-between">
                <span class="text-sm text-gray-500">Status</span>
                <span class="px-2 py-1 text-xs rounded-full bg-green-100 text-green-800 font-medium">
                  {{ profile.status }}
                </span>
              </div>
            </div>
          </div>
        </div>
        
        <!-- Referral Code Card -->
        <div 
          class="bg-white rounded-xl shadow-lg p-6 mb-6"
          v-motion
          :initial="{ opacity: 0 }"
          :enter="{ opacity: 1, transition: { delay: 200, duration: 500 } }">
          <h3 class="text-lg font-medium text-gray-900 mb-4">Your Referral Code</h3>
          
          <div class="p-4 bg-gray-50 rounded-lg border border-gray-200 mb-4">
            <div class="flex items-center justify-between">
              <div class="font-mono text-lg font-semibold text-accent">{{ profile.referralCode }}</div>
              <button 
                @click="copyReferralCode"
                class="text-gray-500 hover:text-accent transition-colors duration-300 focus:outline-none">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path>
                </svg>
              </button>
            </div>
          </div>
          
          <div class="flex items-center mb-4">
            <button 
              @click="showEditCodeModal = true"
              class="text-sm text-accent hover:text-accent/80 focus:outline-none transition-colors duration-300">
              Customize Referral Code
            </button>
            <span class="text-gray-300 mx-2">|</span>
            <button 
              @click="regenerateReferralCode"
              class="text-sm text-gray-600 hover:text-gray-900 focus:outline-none transition-colors duration-300">
              Generate New Code
            </button>
          </div>
          
          <div class="flex space-x-2">
            <a 
              :href="`https://wa.me/?text=Use my referral code ${profile.referralCode} on Ecomflex to get 100% cashback!`"
              target="_blank"
              class="flex-1 py-2 px-3 bg-[#25D366] text-white rounded-lg text-sm font-medium flex items-center justify-center hover:bg-opacity-90 transition-colors duration-300 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#25D366]">
              <svg class="w-4 h-4 mr-2" fill="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347m-5.421 7.403h-.004a9.87 9.87 0 01-5.031-1.378l-.361-.214-3.741.982.998-3.648-.235-.374a9.86 9.86 0 01-1.51-5.26c.001-5.45 4.436-9.884 9.888-9.884 2.64 0 5.122 1.03 6.988 2.898a9.825 9.825 0 012.893 6.994c-.003 5.45-4.437 9.884-9.885 9.884m8.413-18.297A11.815 11.815 0 0012.05 0C5.495 0 .16 5.335.157 11.892c0 2.096.547 4.142 1.588 5.945L.057 24l6.305-1.654a11.882 11.882 0 005.683 1.448h.005c6.554 0 11.89-5.335 11.893-11.893a11.821 11.821 0 00-3.48-8.413z" />
              </svg>
              WhatsApp
            </a>
            <a
              :href="`https://twitter.com/intent/tweet?text=Use my referral code ${profile.referralCode} on Ecomflex to get 100% cashback!`"
              target="_blank" 
              class="flex-1 py-2 px-3 bg-[#1DA1F2] text-white rounded-lg text-sm font-medium flex items-center justify-center hover:bg-opacity-90 transition-colors duration-300 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-[#1DA1F2]">
              <svg class="w-4 h-4 mr-2" fill="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <path d="M23.953 4.57a10 10 0 01-2.825.775 4.958 4.958 0 002.163-2.723c-.951.555-2.005.959-3.127 1.184a4.92 4.92 0 00-8.384 4.482C7.69 8.095 4.067 6.13 1.64 3.162a4.822 4.822 0 00-.666 2.475c0 1.71.87 3.213 2.188 4.096a4.904 4.904 0 01-2.228-.616v.06a4.923 4.923 0 003.946 4.827 4.996 4.996 0 01-2.212.085 4.936 4.936 0 004.604 3.417 9.867 9.867 0 01-6.102 2.105c-.39 0-.779-.023-1.17-.067a13.995 13.995 0 007.557 2.209c9.053 0 13.998-7.496 13.998-13.985 0-.21 0-.42-.015-.63A9.935 9.935 0 0024 4.59z" />
              </svg>
              Twitter
            </a>
          </div>
        </div>
        
        <!-- Account Statistics Card -->
        <div 
          class="bg-white rounded-xl shadow-lg p-6"
          v-motion
          :initial="{ opacity: 0 }"
          :enter="{ opacity: 1, transition: { delay: 300, duration: 500 } }">
          <h3 class="text-lg font-medium text-gray-900 mb-4">Account Statistics</h3>
          
          <div class="space-y-4">
            <div>
              <div class="flex justify-between mb-1">
                <span class="text-sm font-medium text-gray-700">Total Referrals</span>
                <span class="text-sm font-medium text-gray-900">{{ profile.stats.totalReferrals }}</span>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div 
                  class="bg-accent h-2 rounded-full"
                  :style="`width: ${Math.min(100, (profile.stats.totalReferrals / 100) * 100)}%`"></div>
              </div>
            </div>
            
            <div>
              <div class="flex justify-between mb-1">
                <span class="text-sm font-medium text-gray-700">Successful Conversions</span>
                <span class="text-sm font-medium text-gray-900">{{ profile.stats.conversions }} ({{ Math.round((profile.stats.conversions / profile.stats.totalReferrals) * 100) }}%)</span>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div 
                  class="bg-green-500 h-2 rounded-full"
                  :style="`width: ${(profile.stats.conversions / profile.stats.totalReferrals) * 100}%`"></div>
              </div>
            </div>
            
            <div>
              <div class="flex justify-between mb-1">
                <span class="text-sm font-medium text-gray-700">Pending Verifications</span>
                <span class="text-sm font-medium text-gray-900">{{ profile.stats.pendingVerifications }}</span>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div 
                  class="bg-amber-500 h-2 rounded-full"
                  :style="`width: ${(profile.stats.pendingVerifications / profile.stats.totalReferrals) * 100}%`"></div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Right Column - Edit Profile Tabs -->
      <div 
        class="lg:col-span-2"
        v-motion
        :initial="{ opacity: 0, x: 20 }"
        :enter="{ opacity: 1, x: 0, transition: { duration: 500 } }">
        <div class="bg-white rounded-xl shadow-lg overflow-hidden">
          <!-- Tabs Navigation -->
          <div class="border-b border-gray-200">
            <nav class="flex -mb-px">
              <button 
                v-for="tab in ['Basic Information', 'Social Media', 'Password & Security', 'Notifications']"
                :key="tab"
                @click="activeTab = tab"
                :class="[
                  'py-4 px-6 text-sm font-medium border-b-2 focus:outline-none transition-colors duration-300',
                  activeTab === tab 
                    ? 'border-accent text-accent' 
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                ]">
                {{ tab }}
              </button>
            </nav>
          </div>
          
          <!-- Basic Information Tab -->
          <div v-if="activeTab === 'Basic Information'" class="p-6">
            <form @submit.prevent="saveBasicInfo" class="space-y-6">
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label for="fullName" class="block text-sm font-medium text-gray-700 mb-1">
                    Full Name
                  </label>
                  <input 
                    type="text" 
                    id="fullName" 
                    v-model="editedProfile.fullName" 
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-accent focus:border-accent"
                    required>
                </div>
                
                <div>
                  <label for="email" class="block text-sm font-medium text-gray-700 mb-1">
                    Email Address
                  </label>
                  <input 
                    type="email" 
                    id="email" 
                    v-model="editedProfile.email" 
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-accent focus:border-accent"
                    required>
                </div>
                
                <div>
                  <label for="phone" class="block text-sm font-medium text-gray-700 mb-1">
                    Phone Number
                  </label>
                  <input 
                    type="tel" 
                    id="phone" 
                    v-model="editedProfile.phone" 
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-accent focus:border-accent">
                </div>
                
                <div>
                  <label for="location" class="block text-sm font-medium text-gray-700 mb-1">
                    Location
                  </label>
                  <input 
                    type="text" 
                    id="location" 
                    v-model="editedProfile.location" 
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-accent focus:border-accent"
                    placeholder="City, Country">
                </div>
              </div>
              
              <div>
                <label for="bio" class="block text-sm font-medium text-gray-700 mb-1">
                  Bio
                </label>
                <textarea 
                  id="bio" 
                  v-model="editedProfile.bio" 
                  rows="4" 
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-accent focus:border-accent"
                  placeholder="Tell us about yourself..."></textarea>
                <p class="mt-1 text-sm text-gray-500">
                  {{ editedProfile.bio ? editedProfile.bio.length : 0 }}/200 characters
                </p>
              </div>
              
              <div>
                <label for="niches" class="block text-sm font-medium text-gray-700 mb-1">
                  Content Niches
                </label>
                <div class="flex flex-wrap gap-2 mb-2">
                  <span 
                    v-for="(niche, index) in editedProfile.niches" 
                    :key="index"
                    class="inline-flex items-center px-3 py-1 rounded-full text-sm bg-accent/10 text-accent">
                    {{ niche }}
                    <button 
                      type="button"
                      @click="removeNiche(index)"
                      class="ml-1 text-accent hover:text-accent/80 focus:outline-none">
                      <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                      </svg>
                    </button>
                  </span>
                </div>
                <div class="flex">
                  <input 
                    type="text" 
                    id="newNiche"
                    v-model="newNiche" 
                    class="flex-1 px-4 py-2 border border-gray-300 rounded-l-lg focus:ring-accent focus:border-accent"
                    placeholder="Add a niche (e.g. Beauty, Tech, Fitness)">
                  <button 
                    type="button"
                    @click="addNiche"
                    :disabled="!newNiche"
                    class="px-4 py-2 bg-accent text-white rounded-r-lg hover:bg-accent/80 focus:outline-none focus:ring-2 focus:ring-accent/30 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors duration-300">
                    Add
                  </button>
                </div>
              </div>
              
              <div class="flex justify-end">
                <button 
                  type="submit"
                  :disabled="isSaving"
                  class="px-6 py-2 bg-accent text-white rounded-lg hover:bg-accent/80 focus:outline-none focus:ring-2 focus:ring-accent/30 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors duration-300">
                  <span v-if="isSaving" class="flex items-center">
                    <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    Saving...
                  </span>
                  <span v-else>Save Changes</span>
                </button>
              </div>
            </form>
          </div>
          
          <!-- Social Media Tab -->
          <div v-if="activeTab === 'Social Media'" class="p-6">
            <form @submit.prevent="saveSocialMedia" class="space-y-6">
              <div class="grid grid-cols-1 gap-6">
                <div>
                  <label for="instagram" class="block text-sm font-medium text-gray-700 mb-1">
                    Instagram
                  </label>
                  <div class="mt-1 flex rounded-md shadow-sm">
                    <span class="inline-flex items-center px-3 rounded-l-md border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm">
                      instagram.com/
                    </span>
                    <input 
                      type="text" 
                      id="instagram" 
                      v-model="editedProfile.social.instagram" 
                      class="flex-1 min-w-0 block w-full px-3 py-2 rounded-none rounded-r-md border border-gray-300 focus:ring-accent focus:border-accent sm:text-sm"
                      placeholder="username">
                  </div>
                </div>
                
                <div>
                  <label for="twitter" class="block text-sm font-medium text-gray-700 mb-1">
                    Twitter
                  </label>
                  <div class="mt-1 flex rounded-md shadow-sm">
                    <span class="inline-flex items-center px-3 rounded-l-md border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm">
                      twitter.com/
                    </span>
                    <input 
                      type="text" 
                      id="twitter" 
                      v-model="editedProfile.social.twitter" 
                      class="flex-1 min-w-0 block w-full px-3 py-2 rounded-none rounded-r-md border border-gray-300 focus:ring-accent focus:border-accent sm:text-sm"
                      placeholder="username">
                  </div>
                </div>
                
                <div>
                  <label for="tiktok" class="block text-sm font-medium text-gray-700 mb-1">
                    TikTok
                  </label>
                  <div class="mt-1 flex rounded-md shadow-sm">
                    <span class="inline-flex items-center px-3 rounded-l-md border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm">
                      tiktok.com/@
                    </span>
                    <input 
                      type="text" 
                      id="tiktok" 
                      v-model="editedProfile.social.tiktok" 
                      class="flex-1 min-w-0 block w-full px-3 py-2 rounded-none rounded-r-md border border-gray-300 focus:ring-accent focus:border-accent sm:text-sm"
                      placeholder="username">
                  </div>
                </div>
                
                <div>
                  <label for="youtube" class="block text-sm font-medium text-gray-700 mb-1">
                    YouTube
                  </label>
                  <div class="mt-1 flex rounded-md shadow-sm">
                    <span class="inline-flex items-center px-3 rounded-l-md border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm">
                      youtube.com/c/
                    </span>
                    <input 
                      type="text" 
                      id="youtube" 
                      v-model="editedProfile.social.youtube" 
                      class="flex-1 min-w-0 block w-full px-3 py-2 rounded-none rounded-r-md border border-gray-300 focus:ring-accent focus:border-accent sm:text-sm"
                      placeholder="channel">
                  </div>
                </div>
                
                <div>
                  <label for="website" class="block text-sm font-medium text-gray-700 mb-1">
                    Website or Blog
                  </label>
                  <input 
                    type="url" 
                    id="website" 
                    v-model="editedProfile.social.website" 
                    class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-accent focus:border-accent"
                    placeholder="https://yourblog.com">
                </div>
              </div>
              
              <div class="flex justify-end">
                <button 
                  type="submit"
                  :disabled="isSaving"
                  class="px-6 py-2 bg-accent text-white rounded-lg hover:bg-accent/80 focus:outline-none focus:ring-2 focus:ring-accent/30 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors duration-300">
                  <span v-if="isSaving" class="flex items-center">
                    <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    Saving...
                  </span>
                  <span v-else>Save Changes</span>
                </button>
              </div>
            </form>
          </div>
          
          <!-- Password & Security Tab -->
          <div v-if="activeTab === 'Password & Security'" class="p-6">
            <form @submit.prevent="savePasswordChanges" class="space-y-6">
              <div>
                <label for="currentPassword" class="block text-sm font-medium text-gray-700 mb-1">
                  Current Password
                </label>
                <input 
                  type="password" 
                  id="currentPassword" 
                  v-model="passwordForm.currentPassword" 
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-accent focus:border-accent"
                  required>
              </div>
              
              <div>
                <label for="newPassword" class="block text-sm font-medium text-gray-700 mb-1">
                  New Password
                </label>
                <input 
                  type="password" 
                  id="newPassword" 
                  v-model="passwordForm.newPassword" 
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-accent focus:border-accent"
                  required>
                <div class="mt-2">
                  <div class="flex items-center mb-1">
                    <div class="text-xs mr-2">Password Strength:</div>
                    <div class="flex-1 bg-gray-200 rounded-full h-1">
                      <div 
                        :class="[
                          'h-1 rounded-full',
                          passwordStrength < 2 ? 'bg-red-500' : 
                          passwordStrength < 3 ? 'bg-yellow-500' : 
                          'bg-green-500'
                        ]"
                        :style="`width: ${passwordStrength * 25}%`"></div>
                    </div>
                  </div>
                  <ul class="text-xs text-gray-500 space-y-1">
                    <li :class="{ 'text-green-600': passwordForm.newPassword.length >= 8 }">
                      <span v-if="passwordForm.newPassword.length >= 8">✓</span>
                      <span v-else>·</span>
                      At least 8 characters
                    </li>
                    <li :class="{ 'text-green-600': /[A-Z]/.test(passwordForm.newPassword) }">
                      <span v-if="/[A-Z]/.test(passwordForm.newPassword)">✓</span>
                      <span v-else>·</span>
                      Contains uppercase letter
                    </li>
                    <li :class="{ 'text-green-600': /[0-9]/.test(passwordForm.newPassword) }">
                      <span v-if="/[0-9]/.test(passwordForm.newPassword)">✓</span>
                      <span v-else>·</span>
                      Contains number
                    </li>
                    <li :class="{ 'text-green-600': /[^A-Za-z0-9]/.test(passwordForm.newPassword) }">
                      <span v-if="/[^A-Za-z0-9]/.test(passwordForm.newPassword)">✓</span>
                      <span v-else>·</span>
                      Contains special character
                    </li>
                  </ul>
                </div>
              </div>
              
              <div>
                <label for="confirmPassword" class="block text-sm font-medium text-gray-700 mb-1">
                  Confirm New Password
                </label>
                <input 
                  type="password" 
                  id="confirmPassword" 
                  v-model="passwordForm.confirmPassword" 
                  class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-accent focus:border-accent"
                  required>
                <p 
                  v-if="passwordForm.newPassword && passwordForm.confirmPassword && passwordForm.newPassword !== passwordForm.confirmPassword" 
                  class="mt-1 text-sm text-red-600">
                  Passwords do not match
                </p>
              </div>
              
              <div class="pt-4 border-t border-gray-200">
                <h4 class="text-base font-medium text-gray-900 mb-4">Two-Factor Authentication</h4>
                
                <div class="flex items-center justify-between">
                  <div>
                    <p class="text-sm text-gray-900">Protect your account with 2FA</p>
                    <p class="text-xs text-gray-500 mt-1">
                      Add an extra layer of security by requiring a verification code in addition to your password.
                    </p>
                  </div>
                  <div class="flex items-center">
                    <button 
                      v-if="!profile.twoFactorEnabled"
                      type="button" 
                      @click="setupTwoFactor"
                      class="text-sm text-accent hover:text-accent/80 transition-colors duration-300 focus:outline-none">
                      Set up
                    </button>
                    <button 
                      v-else
                      type="button" 
                      @click="disableTwoFactor"
                      class="text-sm text-red-600 hover:text-red-800 transition-colors duration-300 focus:outline-none">
                      Disable
                    </button>
                    <div class="ml-3 relative inline-block w-10 align-middle select-none">
                      <input 
                        type="checkbox" 
                        id="toggle-2fa" 
                        name="toggle-2fa" 
                        :checked="profile.twoFactorEnabled"
                        @change="toggleTwoFactor"
                        class="sr-only peer">
                      <label 
                        for="toggle-2fa" 
                        class="block overflow-hidden h-6 rounded-full bg-gray-300 cursor-pointer peer-checked:bg-accent peer-focus:ring-2 peer-focus:ring-offset-2 peer-focus:ring-accent/30 transition-colors duration-300"></label>
                      <div class="absolute inset-y-0 left-0 w-4 h-4 m-1 rounded-full bg-white transform transition-transform duration-300 peer-checked:translate-x-4"></div>
                    </div>
                  </div>
                </div>
              </div>
              
              <div class="flex justify-end">
                <button 
                  type="submit"
                  :disabled="isSaving || !isPasswordFormValid"
                  class="px-6 py-2 bg-accent text-white rounded-lg hover:bg-accent/80 focus:outline-none focus:ring-2 focus:ring-accent/30 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors duration-300">
                  <span v-if="isSaving" class="flex items-center">
                    <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    Saving...
                  </span>
                  <span v-else>Update Password</span>
                </button>
              </div>
            </form>
          </div>
          
          <!-- Notifications Tab -->
          <div v-if="activeTab === 'Notifications'" class="p-6">
            <form @submit.prevent="saveNotificationSettings" class="space-y-6">
              <div>
                <h4 class="text-base font-medium text-gray-900 mb-4">Email Notifications</h4>
                
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm text-gray-900">New Referral</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Get notified when someone uses your referral code
                      </p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="editedProfile.notifications.newReferral" 
                          class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                      </label>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm text-gray-900">Booking Confirmed</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Get notified when a referral books a product
                      </p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="editedProfile.notifications.bookingConfirmed" 
                          class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                      </label>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm text-gray-900">Cashback Approved</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Get notified when a referral receives cashback
                      </p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="editedProfile.notifications.cashbackApproved" 
                          class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                      </label>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm text-gray-900">Payout Processed</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Get notified when your payout is processed
                      </p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="editedProfile.notifications.payoutProcessed" 
                          class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                      </label>
                    </div>
                  </div>
                </div>
              </div>
              
              <div class="pt-4 border-t border-gray-200">
                <h4 class="text-base font-medium text-gray-900 mb-4">Push Notifications</h4>
                
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm text-gray-900">Enable Push Notifications</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Receive real-time notifications on your device
                      </p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="editedProfile.notifications.enablePush" 
                          class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                      </label>
                    </div>
                  </div>
                </div>
              </div>
              
              <div class="flex justify-end">
                <button 
                  type="submit"
                  :disabled="isSaving"
                  class="px-6 py-2 bg-accent text-white rounded-lg hover:bg-accent/80 focus:outline-none focus:ring-2 focus:ring-accent/30 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors duration-300">
                  <span v-if="isSaving" class="flex items-center">
                    <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    Saving...
                  </span>
                  <span v-else>Save Changes</span>
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>

    <!-- Edit Referral Code Modal -->
    <div v-if="showEditCodeModal" class="fixed inset-0 z-50 overflow-y-auto" aria-labelledby="modal-title" role="dialog" aria-modal="true">
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <div 
          class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" 
          aria-hidden="true"
          v-motion
          :initial="{ opacity: 0 }"
          :enter="{ opacity: 1, transition: { duration: 300 } }"
          @click="showEditCodeModal = false"></div>

        <span class="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="true">&#8203;</span>

        <div 
          class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full"
          v-motion
          :initial="{ opacity: 0, scale: 0.9 }"
          :enter="{ opacity: 1, scale: 1, transition: { duration: 300 } }">
          <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <div class="sm:flex sm:items-start">
              <div class="mx-auto flex-shrink-0 flex items-center justify-center h-12 w-12 rounded-full bg-accent bg-opacity-20 sm:mx-0 sm:h-10 sm:w-10">
                <svg class="h-6 w-6 text-accent" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                </svg>
              </div>
              <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                <h3 class="text-lg leading-6 font-medium text-gray-900" id="modal-title">
                  Customize Referral Code
                </h3>
                <div class="mt-4">
                  <div class="mb-4">
                    <label for="customCode" class="block text-sm font-medium text-gray-700 mb-1">
                      Custom Referral Code
                    </label>
                    <div class="mt-1 relative rounded-md shadow-sm">
                      <input 
                        type="text" 
                        id="customCode" 
                        v-model="customReferralCode" 
                        maxlength="12"
                        class="focus:ring-accent focus:border-accent block w-full pr-10 sm:text-sm border-gray-300 rounded-md"
                        placeholder="e.g. INFLUENCER25"
                        @input="checkCodeAvailability">
                      <div class="absolute inset-y-0 right-0 flex items-center pr-3">
                        <svg 
                          v-if="isCheckingCode" 
                          class="animate-spin h-5 w-5 text-gray-400" 
                          xmlns="http://www.w3.org/2000/svg" 
                          fill="none" 
                          viewBox="0 0 24 24">
                          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                        </svg>
                        <svg 
                          v-else-if="customReferralCode && isCodeAvailable" 
                          class="h-5 w-5 text-green-500" 
                          fill="none" 
                          stroke="currentColor" 
                          viewBox="0 0 24 24" 
                          xmlns="http://www.w3.org/2000/svg">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                        </svg>
                        <svg 
                          v-else-if="customReferralCode && !isCodeAvailable" 
                          class="h-5 w-5 text-red-500" 
                          fill="none" 
                          stroke="currentColor" 
                          viewBox="0 0 24 24" 
                          xmlns="http://www.w3.org/2000/svg">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                        </svg>
                      </div>
                    </div>
                    <p class="mt-1 text-xs text-gray-500">
                      4-12 characters, letters and numbers only. No spaces.
                    </p>
                    <p v-if="customReferralCode && !isCodeAvailable && !isCheckingCode" class="mt-1 text-xs text-red-600">
                      This code is already taken. Please try another one.
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
            <button 
              type="button" 
              class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-accent text-base font-medium text-white hover:bg-accent/80 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-accent/30 sm:ml-3 sm:w-auto sm:text-sm transition-all duration-300"
              :disabled="!customReferralCode || !isCodeAvailable || isCheckingCode || isSaving"
              @click="updateReferralCode">
              <span v-if="isSaving" class="flex items-center">
                <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                Updating...
              </span>
              <span v-else>Update Code</span>
            </button>
            <button 
              type="button" 
              class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-accent/30 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm transition-all duration-300"
              @click="showEditCodeModal = false">
              Cancel
            </button>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Toast Notification -->
    <div 
      v-if="showToast"
      class="fixed top-4 right-4 flex items-center p-4 mb-4 text-sm rounded-lg shadow-lg max-w-xs"
      :class="[
        toastType === 'success' ? 'text-green-800 bg-green-50' : 
        toastType === 'error' ? 'text-red-800 bg-red-50' : 
        'text-blue-800 bg-blue-50'
      ]"
      role="alert"
      v-motion
      :initial="{ opacity: 0, x: 20 }"
      :enter="{ opacity: 1, x: 0, transition: { duration: 300 } }"
      :leave="{ opacity: 0, x: 20, transition: { duration: 300 } }">
      <div 
        class="inline-flex items-center justify-center flex-shrink-0 w-8 h-8 rounded-lg mr-2"
        :class="[
          toastType === 'success' ? 'text-green-500 bg-green-100' : 
          toastType === 'error' ? 'text-red-500 bg-red-100' : 
          'text-blue-500 bg-blue-100'
        ]">
        <svg 
          v-if="toastType === 'success'" 
          class="w-5 h-5" 
          fill="currentColor" 
          viewBox="0 0 20 20" 
          xmlns="http://www.w3.org/2000/svg">
          <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path>
        </svg>
        <svg 
          v-else-if="toastType === 'error'" 
          class="w-5 h-5" 
          fill="currentColor" 
          viewBox="0 0 20 20" 
          xmlns="http://www.w3.org/2000/svg">
          <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"></path>
        </svg>
        <svg 
          v-else 
          class="w-5 h-5" 
          fill="currentColor" 
          viewBox="0 0 20 20" 
          xmlns="http://www.w3.org/2000/svg">
          <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2h-1V9z" clip-rule="evenodd"></path>
        </svg>
      </div>
      <div>{{ toastMessage }}</div>
      <button 
        type="button" 
        class="ml-auto -mx-1.5 -my-1.5 rounded-lg p-1.5 inline-flex h-8 w-8 focus:outline-none"
        :class="[
          toastType === 'success' ? 'text-green-500 hover:bg-green-100' : 
          toastType === 'error' ? 'text-red-500 hover:bg-red-100' : 
          'text-blue-500 hover:bg-blue-100'
        ]"
        @click="showToast = false"
        aria-label="Close">
        <span class="sr-only">Close</span>
        <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
          <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd"></path>
        </svg>
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue';
import { useReferralStore } from '@/stores/referral';
import { useNotificationStore } from '@/stores/notification';

// Import VueUse motion for animations
import { useMotion } from '@vueuse/motion';

// Stores
const referralStore = useReferralStore();
const notificationStore = useNotificationStore();

// State
const fileInput = ref(null);
const activeTab = ref('Basic Information');
const isSaving = ref(false);
const showEditCodeModal = ref(false);
const customReferralCode = ref('');
const isCheckingCode = ref(false);
const isCodeAvailable = ref(false);
const newNiche = ref('');
const showToast = ref(false);
const toastMessage = ref('');
const toastType = ref('success');
const toastTimeout = ref(null);

// Password form
const passwordForm = ref({
  currentPassword: '',
  newPassword: '',
  confirmPassword: ''
});

// Mock profile data
const profile = ref({
  fullName: 'Sarah Johnson',
  email: 'sarah.johnson@example.com',
  phone: '+1 (555) 123-4567',
  location: 'Los Angeles, CA',
  bio: 'Beauty and lifestyle influencer with a passion for sustainable products and ethical brands.',
  joinDate: new Date('2024-05-15'),
  status: 'Active',
  avatar: null,
  referralCode: 'SARAHJ25',
  twoFactorEnabled: false,
  niches: ['Beauty', 'Lifestyle', 'Sustainability'],
  social: {
    instagram: 'sarahjbeauty',
    twitter: 'sarahj_beauty',
    tiktok: 'sarahjohnson',
    youtube: 'SarahJBeauty',
    website: 'https://sarahjohnson.com'
  },
  stats: {
    totalReferrals: 78,
    conversions: 45,
    pendingVerifications: 12
  },
  notifications: {
    newReferral: true,
    bookingConfirmed: true,
    cashbackApproved: true,
    payoutProcessed: true,
    enablePush: false
  }
});

// Create a copy for editing
const editedProfile = ref({...profile.value});

// Computed properties
const passwordStrength = computed(() => {
  let strength = 0;
  const password = passwordForm.value.newPassword;
  
  if (!password) return 0;
  
  // Length check
  if (password.length >= 8) strength++;
  
  // Uppercase check
  if (/[A-Z]/.test(password)) strength++;
  
  // Number check
  if (/[0-9]/.test(password)) strength++;
  
  // Special character check
  if (/[^A-Za-z0-9]/.test(password)) strength++;
  
  return strength;
});

const isPasswordFormValid = computed(() => {
  return (
    passwordForm.value.currentPassword && 
    passwordForm.value.newPassword && 
    passwordForm.value.confirmPassword && 
    passwordForm.value.newPassword === passwordForm.value.confirmPassword &&
    passwordStrength.value >= 3
  );
});

// Methods
const formatDate = (date) => {
  return new Intl.DateTimeFormat('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  }).format(date);
};

const getInitials = (name) => {
  if (!name) return '';
  return name
    .split(' ')
    .map(part => part.charAt(0))
    .join('')
    .toUpperCase()
    .substring(0, 2);
};

const triggerFileInput = () => {
  fileInput.value.click();
};

const handleFileUpload = (event) => {
  const file = event.target.files[0];
  if (!file) return;
  
  // Check if file is an image
  if (!file.type.startsWith('image/')) {
    showToastNotification('Please select an image file', 'error');
    return;
  }
  
  // Check file size (max 2MB)
  if (file.size > 2 * 1024 * 1024) {
    showToastNotification('Image size should be less than 2MB', 'error');
    return;
  }
  
  // Create a URL for the image
  const reader = new FileReader();
  reader.onload = (e) => {
    profile.value.avatar = e.target.result;
    editedProfile.value.avatar = e.target.result;
    showToastNotification('Profile picture updated successfully', 'success');
  };
  reader.readAsDataURL(file);
};

const saveBasicInfo = async () => {
  isSaving.value = true;
  
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1500));
    
    // Update profile
    profile.value.fullName = editedProfile.value.fullName;
    profile.value.email = editedProfile.value.email;
    profile.value.phone = editedProfile.value.phone;
    profile.value.location = editedProfile.value.location;
    profile.value.bio = editedProfile.value.bio;
    profile.value.niches = [...editedProfile.value.niches];
    
    showToastNotification('Profile information updated successfully', 'success');
  } catch (error) {
    showToastNotification('Failed to update profile information', 'error');
  } finally {
    isSaving.value = false;
  }
};

const saveSocialMedia = async () => {
  isSaving.value = true;
  
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1500));
    
    // Update profile
    profile.value.social = {
      instagram: editedProfile.value.social.instagram,
      twitter: editedProfile.value.social.twitter,
      tiktok: editedProfile.value.social.tiktok,
      youtube: editedProfile.value.social.youtube,
      website: editedProfile.value.social.website
    };
    
    showToastNotification('Social media profiles updated successfully', 'success');
  } catch (error) {
    showToastNotification('Failed to update social media profiles', 'error');
  } finally {
    isSaving.value = false;
  }
};

const savePasswordChanges = async () => {
  isSaving.value = true;
  
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1500));
    
    // Reset form
    passwordForm.value = {
      currentPassword: '',
      newPassword: '',
      confirmPassword: ''
    };
    
    showToastNotification('Password updated successfully', 'success');
  } catch (error) {
    showToastNotification('Failed to update password', 'error');
  } finally {
    isSaving.value = false;
  }
};

const saveNotificationSettings = async () => {
  isSaving.value = true;
  
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1500));
    
    // Update profile
    profile.value.notifications = {
      newReferral: editedProfile.value.notifications.newReferral,
      bookingConfirmed: editedProfile.value.notifications.bookingConfirmed,
      cashbackApproved: editedProfile.value.notifications.cashbackApproved,
      payoutProcessed: editedProfile.value.notifications.payoutProcessed,
      enablePush: editedProfile.value.notifications.enablePush
    };
    
    showToastNotification('Notification settings updated successfully', 'success');
  } catch (error) {
    showToastNotification('Failed to update notification settings', 'error');
  } finally {
    isSaving.value = false;
  }
};

const copyReferralCode = () => {
  navigator.clipboard.writeText(profile.value.referralCode)
    .then(() => {
      showToastNotification('Referral code copied to clipboard', 'success');
    })
    .catch(() => {
      showToastNotification('Failed to copy referral code', 'error');
    });
};

const checkCodeAvailability = () => {
  // Format and validate the code
  customReferralCode.value = customReferralCode.value.toUpperCase().replace(/[^A-Z0-9]/g, '');
  
  if (customReferralCode.value.length < 4) {
    isCodeAvailable.value = false;
    return;
  }
  
  isCheckingCode.value = true;
  
  // Simulate API call to check availability
  setTimeout(() => {
    // For demo, let's assume codes ending with even numbers are available
    const lastChar = customReferralCode.value.slice(-1);
    isCodeAvailable.value = !isNaN(lastChar) ? parseInt(lastChar) % 2 === 0 : true;
    isCheckingCode.value = false;
  }, 800);
};

const updateReferralCode = async () => {
  isSaving.value = true;
  
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1500));
    
    profile.value.referralCode = customReferralCode.value;
    showEditCodeModal.value = false;
    customReferralCode.value = '';
    
    showToastNotification('Referral code updated successfully', 'success');
  } catch (error) {
    showToastNotification('Failed to update referral code', 'error');
  } finally {
    isSaving.value = false;
  }
};

const regenerateReferralCode = async () => {
  isSaving.value = true;
  
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1500));
    
    // Generate a random code
    const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789';
    let result = '';
    for (let i = 0; i < 8; i++) {
      result += characters.charAt(Math.floor(Math.random() * characters.length));
    }
    
    profile.value.referralCode = result;
    
    showToastNotification('New referral code generated successfully', 'success');
  } catch (error) {
    showToastNotification('Failed to generate new referral code', 'error');
  } finally {
    isSaving.value = false;
  }
};

const addNiche = () => {
  if (!newNiche.value || editedProfile.value.niches.includes(newNiche.value)) return;
  
  editedProfile.value.niches.push(newNiche.value);
  newNiche.value = '';
};

const removeNiche = (index) => {
  editedProfile.value.niches.splice(index, 1);
};

const setupTwoFactor = () => {
  // Show 2FA setup modal (not implemented in this demo)
  showToastNotification('Two-factor authentication setup initiated', 'info');
};

const disableTwoFactor = () => {
  // Show 2FA disable confirmation (not implemented in this demo)
  showToastNotification('Two-factor authentication disabled', 'info');
  profile.value.twoFactorEnabled = false;
};

const toggleTwoFactor = () => {
  if (profile.value.twoFactorEnabled) {
    disableTwoFactor();
  } else {
    setupTwoFactor();
  }
};

const showToastNotification = (message, type = 'success') => {
  // Clear existing timeout
  if (toastTimeout.value) {
    clearTimeout(toastTimeout.value);
  }
  
  // Set toast data
  toastMessage.value = message;
  toastType.value = type;
  showToast.value = true;
  
  // Auto hide after 5 seconds
  toastTimeout.value = setTimeout(() => {
    showToast.value = false;
  }, 5000);
};

// Watchers
watch(() => profile.value, (newValue) => {
  // Update edited profile when profile changes
  editedProfile.value = JSON.parse(JSON.stringify(newValue));
}, { deep: true });

// Lifecycle hooks
onMounted(() => {
  // Fetch profile data (would normally be from API)
  // Simulate API call
  setTimeout(() => {
    // Data would be fetched from API here
  }, 800);
});
</script>

<style scoped>
.profile-page {
  @apply p-4 md:p-6 max-w-7xl mx-auto;
}
</style>