<template>
  <div class="settings-page">
    <!-- Page Header -->
    <header class="mb-8">
      <h1 class="text-2xl md:text-3xl font-bold text-gray-900 mb-2">Account Settings</h1>
      <p class="text-gray-600">Manage your account preferences and payments</p>
    </header>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- Left Column - Settings Navigation -->
      <div 
        class="lg:col-span-1"
        v-motion
        :initial="{ opacity: 0, x: -20 }"
        :enter="{ opacity: 1, x: 0, transition: { duration: 500 } }">
        <div class="bg-white rounded-xl shadow-lg overflow-hidden sticky top-6">
          <div class="p-4 bg-gradient-to-r from-accent/30 to-accent/10">
            <h2 class="text-lg font-semibold text-gray-900">Settings</h2>
          </div>
          <nav class="p-2">
            <ul class="space-y-1">
              <li v-for="(section, index) in settingsSections" :key="index">
                <button 
              type="button" 
              class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-red-600 text-base font-medium text-white hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 sm:ml-3 sm:w-auto sm:text-sm transition-all duration-300"
              :disabled="deleteConfirmation !== 'DELETE' || isDeleting"
              @click="deleteAccount">
              <span v-if="isDeleting" class="flex items-center">
                <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                Deleting...
              </span>
              <span v-else>Delete Account</span>
            </button>    @click="activeSection = section.id"
                  class="w-full flex items-center px-4 py-3 rounded-lg text-sm transition-colors duration-300"
                  :class="activeSection === section.id ? 'bg-accent/10 text-accent font-medium' : 'text-gray-700 hover:bg-gray-100'">
                  <span class="mr-3">
                    <component :is="section.icon" class="w-5 h-5" />
                  </span>
                  <span>{{ section.name }}</span>
              </li>
            </ul>
          </nav>
        </div>
      </div>

      <!-- Right Column - Settings Content -->
      <div 
        class="lg:col-span-2"
        v-motion
        :initial="{ opacity: 0, x: 20 }"
        :enter="{ opacity: 1, x: 0, transition: { duration: 500 } }">
        <div class="bg-white rounded-xl shadow-lg overflow-hidden">
          
          <!-- Payment Methods Section -->
          <div v-if="activeSection === 'payment'" class="p-6">
            <h3 class="text-lg font-medium text-gray-900 mb-6">Payment Methods</h3>
            
            <!-- Payment Methods List -->
            <div class="space-y-4 mb-8">
              <div 
                v-for="(method, index) in paymentMethods" 
                :key="index"
                class="border border-gray-200 rounded-lg p-4 relative hover:border-accent/50 transition-colors duration-300"
                :class="{ 'border-accent': method.default }">
                <div class="absolute top-4 right-4 flex space-x-2">
                  <span 
                    v-if="method.default" 
                    class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-accent/10 text-accent">
                    Default
                  </span>
                  <button 
                    @click="editPaymentMethod(method)"
                    class="text-gray-400 hover:text-gray-600 focus:outline-none transition-colors duration-300">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z"></path>
                    </svg>
                  </button>
                  <button 
                    @click="removePaymentMethod(index)"
                    class="text-gray-400 hover:text-red-600 focus:outline-none transition-colors duration-300">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                    </svg>
                  </button>
                </div>
                
                <div class="flex items-center">
                  <div class="flex-shrink-0 mr-4">
                    <img 
                      v-if="method.type === 'paypal'" 
                      src="https://upload.wikimedia.org/wikipedia/commons/thumb/b/b5/PayPal.svg/124px-PayPal.svg.png" 
                      alt="PayPal" 
                      class="h-8">
                    <div 
                      v-else-if="method.type === 'bank'" 
                      class="w-8 h-8 rounded-full bg-blue-100 flex items-center justify-center">
                      <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 14v3m4-3v3m4-3v3M3 21h18M3 10h18M3 7l9-4 9 4M4 10h16v11H4V10z"></path>
                      </svg>
                    </div>
                    <div 
                      v-else-if="method.type === 'venmo'" 
                      class="w-8 h-8 rounded-full bg-blue-500 flex items-center justify-center">
                      <span class="text-white font-bold">V</span>
                    </div>
                    <div 
                      v-else-if="method.type === 'card'" 
                      class="w-10 h-8 flex items-center justify-center">
                      <img 
                        v-if="method.cardBrand === 'visa'" 
                        src="https://upload.wikimedia.org/wikipedia/commons/thumb/5/5e/Visa_Inc._logo.svg/200px-Visa_Inc._logo.svg.png" 
                        alt="Visa" 
                        class="h-5">
                      <img 
                        v-else-if="method.cardBrand === 'mastercard'" 
                        src="https://upload.wikimedia.org/wikipedia/commons/thumb/2/2a/Mastercard-logo.svg/200px-Mastercard-logo.svg.png" 
                        alt="Mastercard" 
                        class="h-6">
                      <img 
                        v-else-if="method.cardBrand === 'amex'" 
                        src="https://upload.wikimedia.org/wikipedia/commons/thumb/f/fa/American_Express_logo_%282018%29.svg/200px-American_Express_logo_%282018%29.svg.png" 
                        alt="American Express" 
                        class="h-5">
                    </div>
                  </div>
                  
                  <div>
                    <p class="text-sm font-medium text-gray-900">
                      {{ method.name }}
                      <button 
                        v-if="!method.default" 
                        @click="setDefaultPaymentMethod(index)"
                        class="ml-2 text-xs text-accent hover:text-accent/80 focus:outline-none transition-colors duration-300">
                        Set as default
                      </button>
                    </p>
                    <p class="text-xs text-gray-500">
                      <template v-if="method.type === 'paypal'">
                        {{ method.email }}
                      </template>
                      <template v-else-if="method.type === 'bank'">
                        •••• {{ method.accountLast4 }}
                      </template>
                      <template v-else-if="method.type === 'venmo'">
                        @{{ method.username }}
                      </template>
                      <template v-else-if="method.type === 'card'">
                        •••• •••• •••• {{ method.last4 }} | Expires {{ method.expMonth }}/{{ method.expYear }}
                      </template>
                    </p>
                  </div>
                </div>
              </div>
            </div>
            
            <!-- Add Payment Method -->
            <div>
              <button 
                @click="showAddPaymentModal = true"
                class="flex items-center px-4 py-2 border border-gray-300 rounded-lg text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-accent/30 transition-colors duration-300">
                <svg class="w-5 h-5 mr-2 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
                </svg>
                Add Payment Method
              </button>
            </div>
            
            <div class="mt-8 border-t border-gray-200 pt-6">
              <h3 class="text-lg font-medium text-gray-900 mb-4">Payout Preferences</h3>
              
              <div class="bg-gray-50 p-4 rounded-lg mb-6">
                <div class="flex items-center justify-between">
                  <div>
                    <p class="text-sm font-medium text-gray-900">Automatic Payouts</p>
                    <p class="text-xs text-gray-500 mt-1">Automatically receive payouts when your balance reaches the threshold</p>
                  </div>
                  <div>
                    <label class="relative inline-flex items-center cursor-pointer">
                      <input 
                        type="checkbox" 
                        v-model="settings.automaticPayouts" 
                        class="sr-only peer">
                      <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                    </label>
                  </div>
                </div>
              </div>
              
              <div class="space-y-4">
                <div>
                  <label for="payoutThreshold" class="block text-sm font-medium text-gray-700 mb-1">
                    Payout Threshold
                  </label>
                  <div class="mt-1 relative rounded-md shadow-sm w-full max-w-xs">
                    <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                      <span class="text-gray-500 sm:text-sm">$</span>
                    </div>
                    <input 
                      type="number" 
                      id="payoutThreshold" 
                      v-model="settings.payoutThreshold" 
                      min="50" 
                      step="10"
                      class="focus:ring-accent focus:border-accent block w-full pl-7 pr-12 sm:text-sm border-gray-300 rounded-md"
                      placeholder="100">
                    <div class="absolute inset-y-0 right-0 flex items-center">
                      <label for="currency" class="sr-only">Currency</label>
                      <select 
                        id="currency" 
                        v-model="settings.currency"
                        class="focus:ring-accent focus:border-accent h-full py-0 pl-2 pr-7 border-transparent bg-transparent text-gray-500 sm:text-sm rounded-r-md">
                        <option value="USD">USD</option>
                        <option value="EUR">EUR</option>
                        <option value="GBP">GBP</option>
                      </select>
                    </div>
                  </div>
                  <p class="mt-1 text-xs text-gray-500">Minimum $50.00</p>
                </div>
                
                <div>
                  <label for="payoutFrequency" class="block text-sm font-medium text-gray-700 mb-1">
                    Payout Frequency
                  </label>
                  <select 
                    id="payoutFrequency" 
                    v-model="settings.payoutFrequency"
                    class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-accent focus:border-accent sm:text-sm rounded-md">
                    <option value="immediately">Immediately when threshold reached</option>
                    <option value="weekly">Weekly (Monday)</option>
                    <option value="biweekly">Bi-weekly (1st and 15th)</option>
                    <option value="monthly">Monthly (1st of month)</option>
                  </select>
                </div>
              </div>
              
              <div class="mt-6">
                <button 
                  @click="savePayoutPreferences"
                  :disabled="isSaving"
                  class="px-4 py-2 bg-accent text-white rounded-lg hover:bg-accent/80 focus:outline-none focus:ring-2 focus:ring-accent/30 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors duration-300">
                  <span v-if="isSaving" class="flex items-center">
                    <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    Saving...
                  </span>
                  <span v-else>Save Preferences</span>
                </button>
              </div>
            </div>
          </div>
          
          <!-- Appearance Section -->
          <div v-if="activeSection === 'appearance'" class="p-6">
            <h3 class="text-lg font-medium text-gray-900 mb-6">Appearance</h3>
            
            <div class="space-y-6">
              <div>
                <h4 class="text-base font-medium text-gray-900 mb-4">Theme</h4>
                
                <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
                  <div 
                    v-for="theme in ['light', 'dark', 'system']" 
                    :key="theme"
                    @click="settings.theme = theme"
                    class="relative border rounded-lg overflow-hidden cursor-pointer transition-all duration-300 hover:shadow-md"
                    :class="settings.theme === theme ? 'border-accent shadow-sm' : 'border-gray-200'">
                    <div 
                      class="h-24"
                      :class="{
                        'bg-white': theme === 'light',
                        'bg-gray-900': theme === 'dark',
                        'bg-gradient-to-r from-white to-gray-900': theme === 'system'
                      }">
                      <div 
                        v-if="theme === 'light' || theme === 'system'" 
                        class="h-4 w-24 rounded ml-2 mt-2"
                        :class="theme === 'system' ? 'bg-gray-200' : 'bg-gray-200'"></div>
                      <div 
                        v-if="theme === 'dark' || theme === 'system'" 
                        class="h-4 w-24 rounded ml-2 mt-2"
                        :class="theme === 'system' ? 'bg-gray-700 ml-auto mr-2' : 'bg-gray-700'"></div>
                    </div>
                    <div class="p-2 bg-white">
                      <div class="flex items-center">
                        <div 
                          v-if="settings.theme === theme" 
                          class="flex-shrink-0 h-4 w-4 rounded-full bg-accent"></div>
                        <div 
                          v-else
                          class="flex-shrink-0 h-4 w-4 rounded-full border border-gray-300"></div>
                        <span class="ml-2 text-sm font-medium text-gray-900 capitalize">{{ theme }}</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              
              <div>
                <h4 class="text-base font-medium text-gray-900 mb-4">Dashboard Layout</h4>
                
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <div 
                    v-for="layout in ['compact', 'comfortable']" 
                    :key="layout"
                    @click="settings.layout = layout"
                    class="relative border rounded-lg overflow-hidden cursor-pointer transition-all duration-300 hover:shadow-md"
                    :class="settings.layout === layout ? 'border-accent shadow-sm' : 'border-gray-200'">
                    <div class="h-24 bg-white p-2">
                      <div 
                        v-if="layout === 'compact'"
                        class="space-y-1">
                        <div class="h-2 w-full bg-gray-200 rounded"></div>
                        <div class="h-2 w-full bg-gray-200 rounded"></div>
                        <div class="h-2 w-full bg-gray-200 rounded"></div>
                        <div class="h-2 w-full bg-gray-200 rounded"></div>
                        <div class="h-2 w-full bg-gray-200 rounded"></div>
                        <div class="h-2 w-full bg-gray-200 rounded"></div>
                      </div>
                      <div 
                        v-else
                        class="space-y-2">
                        <div class="h-3 w-full bg-gray-200 rounded"></div>
                        <div class="h-3 w-full bg-gray-200 rounded"></div>
                        <div class="h-3 w-full bg-gray-200 rounded"></div>
                        <div class="h-3 w-full bg-gray-200 rounded"></div>
                      </div>
                    </div>
                    <div class="p-2 bg-white">
                      <div class="flex items-center">
                        <div 
                          v-if="settings.layout === layout" 
                          class="flex-shrink-0 h-4 w-4 rounded-full bg-accent"></div>
                        <div 
                          v-else
                          class="flex-shrink-0 h-4 w-4 rounded-full border border-gray-300"></div>
                        <span class="ml-2 text-sm font-medium text-gray-900 capitalize">{{ layout }}</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              
              <div>
                <h4 class="text-base font-medium text-gray-900 mb-4">Animation Settings</h4>
                
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">Enable Animations</p>
                      <p class="text-xs text-gray-500 mt-1">Page transitions and UI animations</p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="settings.enableAnimations" 
                          class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                      </label>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">Reduced Motion</p>
                      <p class="text-xs text-gray-500 mt-1">Minimize animations for accessibility</p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="settings.reducedMotion" 
                          class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                      </label>
                    </div>
                  </div>
                </div>
              </div>
              
              <div class="pt-4">
                <button 
                  @click="saveAppearanceSettings"
                  :disabled="isSaving"
                  class="px-4 py-2 bg-accent text-white rounded-lg hover:bg-accent/80 focus:outline-none focus:ring-2 focus:ring-accent/30 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors duration-300">
                  <span v-if="isSaving" class="flex items-center">
                    <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    Saving...
                  </span>
                  <span v-else>Save Appearance Settings</span>
                </button>
              </div>
            </div>
          </div>
          
          <!-- Notifications Section -->
          <div v-if="activeSection === 'notifications'" class="p-6">
            <h3 class="text-lg font-medium text-gray-900 mb-6">Notification Settings</h3>
            
            <div class="space-y-6">
              <div>
                <h4 class="text-base font-medium text-gray-900 mb-4">Email Notifications</h4>
                
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">New Referral</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Get notified when someone uses your referral code
                      </p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="settings.notifications.email.newReferral" 
                          class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                      </label>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">Booking Confirmed</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Get notified when a referral books a product
                      </p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="settings.notifications.email.bookingConfirmed" 
                          class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                      </label>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">Cashback Approved</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Get notified when a referral receives cashback
                      </p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="settings.notifications.email.cashbackApproved" 
                          class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                      </label>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">Payout Processed</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Get notified when your payout is processed
                      </p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="settings.notifications.email.payoutProcessed" 
                          class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                      </label>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">Weekly Summary</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Receive a weekly summary of your referrals and earnings
                      </p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="settings.notifications.email.weeklySummary" 
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
                      <p class="text-sm font-medium text-gray-900">Enable Push Notifications</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Receive real-time notifications on your device
                      </p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="settings.notifications.push.enabled" 
                          class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                      </label>
                    </div>
                  </div>
                  
                  <div v-if="settings.notifications.push.enabled" class="pl-4 border-l-2 border-gray-200 space-y-4">
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-sm font-medium text-gray-900">New Referrals</p>
                      </div>
                      <div>
                        <label class="relative inline-flex items-center cursor-pointer">
                          <input 
                            type="checkbox" 
                            v-model="settings.notifications.push.newReferral" 
                            class="sr-only peer">
                          <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                        </label>
                      </div>
                    </div>
                    
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-sm font-medium text-gray-900">Booking Confirmations</p>
                      </div>
                      <div>
                        <label class="relative inline-flex items-center cursor-pointer">
                          <input 
                            type="checkbox" 
                            v-model="settings.notifications.push.bookingConfirmed" 
                            class="sr-only peer">
                          <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                        </label>
                      </div>
                    </div>
                    
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-sm font-medium text-gray-900">Cashback Approvals</p>
                      </div>
                      <div>
                        <label class="relative inline-flex items-center cursor-pointer">
                          <input 
                            type="checkbox" 
                            v-model="settings.notifications.push.cashbackApproved" 
                            class="sr-only peer">
                          <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                        </label>
                      </div>
                    </div>
                    
                    <div class="flex items-center justify-between">
                      <div>
                        <p class="text-sm font-medium text-gray-900">Payout Confirmations</p>
                      </div>
                      <div>
                        <label class="relative inline-flex items-center cursor-pointer">
                          <input 
                            type="checkbox" 
                            v-model="settings.notifications.push.payoutProcessed" 
                            class="sr-only peer">
                          <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                        </label>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              
              <div class="pt-4 border-t border-gray-200">
                <h4 class="text-base font-medium text-gray-900 mb-4">In-App Notifications</h4>
                
                <div class="flex items-center justify-between">
                  <div>
                    <p class="text-sm font-medium text-gray-900">Notification Sound</p>
                    <p class="text-xs text-gray-500 mt-1">
                      Play a sound for new notifications
                    </p>
                  </div>
                  <div>
                    <label class="relative inline-flex items-center cursor-pointer">
                      <input 
                        type="checkbox" 
                        v-model="settings.notifications.inApp.sound" 
                        class="sr-only peer">
                      <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                    </label>
                  </div>
                </div>
              </div>
              
              <div class="pt-4">
                <button 
                  @click="saveNotificationSettings"
                  :disabled="isSaving"
                  class="px-4 py-2 bg-accent text-white rounded-lg hover:bg-accent/80 focus:outline-none focus:ring-2 focus:ring-accent/30 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors duration-300">
                  <span v-if="isSaving" class="flex items-center">
                    <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    Saving...
                  </span>
                  <span v-else>Save Notification Settings</span>
                </button>
              </div>
            </div>
          </div>
          
          <!-- Privacy Section -->
          <div v-if="activeSection === 'privacy'" class="p-6">
            <h3 class="text-lg font-medium text-gray-900 mb-6">Privacy Settings</h3>
            
            <div class="space-y-6">
              <div>
                <h4 class="text-base font-medium text-gray-900 mb-4">Profile Visibility</h4>
                
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">Public Profile</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Allow your profile to be visible to other users
                      </p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="settings.privacy.publicProfile" 
                          class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                      </label>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">Show Earnings</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Show your earnings statistics publicly
                      </p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="settings.privacy.showEarnings" 
                          :disabled="!settings.privacy.publicProfile"
                          class="sr-only peer">
                        <div 
                          class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"
                          :class="!settings.privacy.publicProfile ? 'opacity-50' : ''"></div>
                      </label>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">Show Referral Count</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Show your total referrals publicly
                      </p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="settings.privacy.showReferralCount" 
                          :disabled="!settings.privacy.publicProfile"
                          class="sr-only peer">
                        <div 
                          class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"
                          :class="!settings.privacy.publicProfile ? 'opacity-50' : ''"></div>
                      </label>
                    </div>
                  </div>
                </div>
              </div>
              
              <div class="pt-4 border-t border-gray-200">
                <h4 class="text-base font-medium text-gray-900 mb-4">Data Preferences</h4>
                
                <div class="space-y-4">
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">Activity Tracking</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Allow us to track your activity for better recommendations
                      </p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="settings.privacy.activityTracking" 
                          class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                      </label>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">Marketing Emails</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Receive promotional emails and offers
                      </p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="settings.privacy.marketingEmails" 
                          class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                      </label>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">Cookie Preferences</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Manage which cookies are stored on your device
                      </p>
                    </div>
                    <div>
                      <button 
                        @click="showCookiePreferences = true"
                        class="text-sm text-accent hover:text-accent/80 focus:outline-none transition-colors duration-300">
                        Manage Cookies
                      </button>
                    </div>
                  </div>
                </div>
              </div>
              
              <div class="pt-4 border-t border-gray-200">
                <h4 class="text-base font-medium text-gray-900 mb-4">Account Data</h4>
                
                <div class="space-y-4">
                  <div>
                    <button 
                      @click="requestDataExport"
                      :disabled="isExporting"
                      class="flex items-center text-sm text-gray-700 hover:text-accent focus:outline-none transition-colors duration-300">
                      <span v-if="isExporting" class="mr-2">
                        <svg class="animate-spin h-4 w-4 text-gray-700" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                        </svg>
                      </span>
                      <span v-else class="mr-2">
                        <svg class="w-4 h-4 text-gray-700" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"></path>
                        </svg>
                      </span>
                      Export Your Data
                    </button>
                    <p class="ml-6 text-xs text-gray-500 mt-1">
                      Download all your account data in JSON format
                    </p>
                  </div>
                  
                  <div>
                    <button 
                      @click="showDeleteAccountModal = true"
                      class="flex items-center text-sm text-red-600 hover:text-red-800 focus:outline-none transition-colors duration-300">
                      <span class="mr-2">
                        <svg class="w-4 h-4 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                        </svg>
                      </span>
                      Delete Account
                    </button>
                    <p class="ml-6 text-xs text-gray-500 mt-1">
                      Permanently delete your account and all associated data
                    </p>
                  </div>
                </div>
              </div>
              
              <div class="pt-4">
                <button 
                  @click="savePrivacySettings"
                  :disabled="isSaving"
                  class="px-4 py-2 bg-accent text-white rounded-lg hover:bg-accent/80 focus:outline-none focus:ring-2 focus:ring-accent/30 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors duration-300">
                  <span v-if="isSaving" class="flex items-center">
                    <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    Saving...
                  </span>
                  <span v-else>Save Privacy Settings</span>
                </button>
              </div>
            </div>
          </div>
          
          <!-- Account Section -->
          <div v-if="activeSection === 'account'" class="p-6">
            <h3 class="text-lg font-medium text-gray-900 mb-6">Account Settings</h3>
            
            <div class="space-y-6">
              <div>
                <h4 class="text-base font-medium text-gray-900 mb-4">Personal Information</h4>
                
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div>
                    <label for="fullName" class="block text-sm font-medium text-gray-700 mb-1">
                      Full Name
                    </label>
                    <input 
                      type="text" 
                      id="fullName" 
                      v-model="settings.account.fullName" 
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
                      v-model="settings.account.email" 
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
                      v-model="settings.account.phone" 
                      class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-accent focus:border-accent">
                  </div>
                  
                  <div>
                    <label for="location" class="block text-sm font-medium text-gray-700 mb-1">
                      Location
                    </label>
                    <input 
                      type="text" 
                      id="location" 
                      v-model="settings.account.location" 
                      class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-accent focus:border-accent"
                      placeholder="City, Country">
                  </div>
                </div>
              </div>
              
              <div class="pt-4 border-t border-gray-200">
                <h4 class="text-base font-medium text-gray-900 mb-4">Password</h4>
                
                <form @submit.prevent="changePassword" class="space-y-4 max-w-md">
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
                  
                  <div>
                    <button 
                      type="submit"
                      :disabled="!isPasswordFormValid || isSaving"
                      class="px-4 py-2 bg-accent text-white rounded-lg hover:bg-accent/80 focus:outline-none focus:ring-2 focus:ring-accent/30 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors duration-300">
                      <span v-if="isSaving" class="flex items-center">
                        <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                        </svg>
                        Updating...
                      </span>
                      <span v-else>Update Password</span>
                    </button>
                  </div>
                </form>
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
                      v-if="!settings.account.twoFactorEnabled"
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
                        :checked="settings.account.twoFactorEnabled"
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
              
              <div class="pt-4 border-t border-gray-200">
                <h4 class="text-base font-medium text-gray-900 mb-4">Language & Region</h4>
                
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div>
                    <label for="language" class="block text-sm font-medium text-gray-700 mb-1">
                      Language
                    </label>
                    <select 
                      id="language" 
                      v-model="settings.account.language"
                      class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-accent focus:border-accent sm:text-sm rounded-md">
                      <option value="en">English</option>
                      <option value="es">Español</option>
                      <option value="fr">Français</option>
                      <option value="de">Deutsch</option>
                      <option value="it">Italiano</option>
                      <option value="pt">Português</option>
                      <option value="ja">日本語</option>
                      <option value="zh">中文</option>
                    </select>
                  </div>
                  
                  <div>
                    <label for="timezone" class="block text-sm font-medium text-gray-700 mb-1">
                      Timezone
                    </label>
                    <select 
                      id="timezone" 
                      v-model="settings.account.timezone"
                      class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-accent focus:border-accent sm:text-sm rounded-md">
                      <option value="UTC-12:00">UTC-12:00</option>
                      <option value="UTC-11:00">UTC-11:00</option>
                      <option value="UTC-10:00">UTC-10:00</option>
                      <option value="UTC-09:00">UTC-09:00</option>
                      <option value="UTC-08:00">UTC-08:00</option>
                      <option value="UTC-07:00">UTC-07:00</option>
                      <option value="UTC-06:00">UTC-06:00</option>
                      <option value="UTC-05:00">UTC-05:00</option>
                      <option value="UTC-04:00">UTC-04:00</option>
                      <option value="UTC-03:00">UTC-03:00</option>
                      <option value="UTC-02:00">UTC-02:00</option>
                      <option value="UTC-01:00">UTC-01:00</option>
                      <option value="UTC+00:00">UTC+00:00</option>
                      <option value="UTC+01:00">UTC+01:00</option>
                      <option value="UTC+02:00">UTC+02:00</option>
                      <option value="UTC+03:00">UTC+03:00</option>
                      <option value="UTC+04:00">UTC+04:00</option>
                      <option value="UTC+05:00">UTC+05:00</option>
                      <option value="UTC+05:30">UTC+05:30</option>
                      <option value="UTC+06:00">UTC+06:00</option>
                      <option value="UTC+07:00">UTC+07:00</option>
                      <option value="UTC+08:00">UTC+08:00</option>
                      <option value="UTC+09:00">UTC+09:00</option>
                      <option value="UTC+10:00">UTC+10:00</option>
                      <option value="UTC+11:00">UTC+11:00</option>
                      <option value="UTC+12:00">UTC+12:00</option>
                    </select>
                  </div>
                </div>
              </div>
              
              <div class="pt-4">
                <button 
                  @click="saveAccountSettings"
                  :disabled="isSaving"
                  class="px-4 py-2 bg-accent text-white rounded-lg hover:bg-accent/80 focus:outline-none focus:ring-2 focus:ring-accent/30 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors duration-300">
                  <span v-if="isSaving" class="flex items-center">
                    <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    Saving...
                  </span>
                  <span v-else>Save Account Settings</span>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Add Payment Method Modal -->
    <div v-if="showAddPaymentModal" class="fixed inset-0 z-50 overflow-y-auto" aria-labelledby="modal-title" role="dialog" aria-modal="true">
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <div 
          class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" 
          aria-hidden="true"
          v-motion
          :initial="{ opacity: 0 }"
          :enter="{ opacity: 1, transition: { duration: 300 } }"
          @click="showAddPaymentModal = false"></div>

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
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"></path>
                </svg>
              </div>
              <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left w-full">
                <h3 class="text-lg leading-6 font-medium text-gray-900" id="modal-title">
                  Add Payment Method
                </h3>
                <div class="mt-4">
                  <div class="mb-4">
                    <label for="paymentType" class="block text-sm font-medium text-gray-700 mb-1">
                      Payment Method Type
                    </label>
                    <select 
                      id="paymentType" 
                      v-model="newPaymentMethod.type"
                      class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-accent focus:border-accent sm:text-sm rounded-md">
                      <option value="paypal">PayPal</option>
                      <option value="bank">Bank Account</option>
                      <option value="venmo">Venmo</option>
                      <option value="card">Credit/Debit Card</option>
                    </select>
                  </div>
                  
                  <!-- PayPal Form -->
                  <div v-if="newPaymentMethod.type === 'paypal'" class="space-y-4">
                    <div>
                      <label for="paypalEmail" class="block text-sm font-medium text-gray-700 mb-1">
                        PayPal Email
                      </label>
                      <input 
                        type="email" 
                        id="paypalEmail" 
                        v-model="newPaymentMethod.email" 
                        class="mt-1 focus:ring-accent focus:border-accent block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                        required>
                    </div>
                    <div>
                      <label for="paypalName" class="block text-sm font-medium text-gray-700 mb-1">
                        Account Name
                      </label>
                      <input 
                        type="text" 
                        id="paypalName" 
                        v-model="newPaymentMethod.name" 
                        class="mt-1 focus:ring-accent focus:border-accent block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                        required>
                    </div>
                  </div>
                  
                  <!-- Bank Account Form -->
                  <div v-if="newPaymentMethod.type === 'bank'" class="space-y-4">
                    <div>
                      <label for="bankName" class="block text-sm font-medium text-gray-700 mb-1">
                        Bank Name
                      </label>
                      <input 
                        type="text" 
                        id="bankName" 
                        v-model="newPaymentMethod.bankName" 
                        class="mt-1 focus:ring-accent focus:border-accent block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                        required>
                    </div>
                    <div>
                      <label for="accountName" class="block text-sm font-medium text-gray-700 mb-1">
                        Account Holder Name
                      </label>
                      <input 
                        type="text" 
                        id="accountName" 
                        v-model="newPaymentMethod.name" 
                        class="mt-1 focus:ring-accent focus:border-accent block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                        required>
                    </div>
                    <div>
                      <label for="accountType" class="block text-sm font-medium text-gray-700 mb-1">
                        Account Type
                      </label>
                      <select 
                        id="accountType" 
                        v-model="newPaymentMethod.accountType"
                        class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-accent focus:border-accent sm:text-sm rounded-md">
                        <option value="checking">Checking</option>
                        <option value="savings">Savings</option>
                      </select>
                    </div>
                    <div>
                      <label for="routingNumber" class="block text-sm font-medium text-gray-700 mb-1">
                        Routing Number
                      </label>
                      <input 
                        type="text" 
                        id="routingNumber" 
                        v-model="newPaymentMethod.routingNumber" 
                        class="mt-1 focus:ring-accent focus:border-accent block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                        required>
                    </div>
                    <div>
                      <label for="accountNumber" class="block text-sm font-medium text-gray-700 mb-1">
                        Account Number
                      </label>
                      <input 
                        type="text" 
                        id="accountNumber" 
                        v-model="newPaymentMethod.accountNumber" 
                        class="mt-1 focus:ring-accent focus:border-accent block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                        required>
                    </div>
                  </div>
                  
                  <!-- Venmo Form -->
                  <div v-if="newPaymentMethod.type === 'venmo'" class="space-y-4">
                    <div>
                      <label for="venmoUsername" class="block text-sm font-medium text-gray-700 mb-1">
                        Venmo Username
                      </label>
                      <div class="mt-1 flex rounded-md shadow-sm">
                        <span class="inline-flex items-center px-3 rounded-l-md border border-r-0 border-gray-300 bg-gray-50 text-gray-500 sm:text-sm">
                          @
                        </span>
                        <input 
                          type="text" 
                          id="venmoUsername" 
                          v-model="newPaymentMethod.username" 
                          class="focus:ring-accent focus:border-accent flex-1 block w-full rounded-none rounded-r-md sm:text-sm border-gray-300"
                          required>
                      </div>
                    </div>
                    <div>
                      <label for="venmoName" class="block text-sm font-medium text-gray-700 mb-1">
                        Account Name
                      </label>
                      <input 
                        type="text" 
                        id="venmoName" 
                        v-model="newPaymentMethod.name" 
                        class="mt-1 focus:ring-accent focus:border-accent block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                        required>
                    </div>
                  </div>
                  
                  <!-- Credit/Debit Card Form -->
                  <div v-if="newPaymentMethod.type === 'card'" class="space-y-4">
                    <div>
                      <label for="cardholderName" class="block text-sm font-medium text-gray-700 mb-1">
                        Cardholder Name
                      </label>
                      <input 
                        type="text" 
                        id="cardholderName" 
                        v-model="newPaymentMethod.name" 
                        class="mt-1 focus:ring-accent focus:border-accent block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                        required>
                    </div>
                    <div>
                      <label for="cardNumber" class="block text-sm font-medium text-gray-700 mb-1">
                        Card Number
                      </label>
                      <input 
                        type="text" 
                        id="cardNumber" 
                        v-model="newPaymentMethod.cardNumber" 
                        class="mt-1 focus:ring-accent focus:border-accent block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                        maxlength="19"
                        placeholder="•••• •••• •••• ••••"
                        @input="formatCardNumber"
                        required>
                    </div>
                    <div class="grid grid-cols-2 gap-4">
                      <div>
                        <label for="expiryDate" class="block text-sm font-medium text-gray-700 mb-1">
                          Expiry Date
                        </label>
                        <input 
                          type="text" 
                          id="expiryDate" 
                          v-model="newPaymentMethod.expiryDate" 
                          class="mt-1 focus:ring-accent focus:border-accent block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                          maxlength="5"
                          placeholder="MM/YY"
                          @input="formatExpiryDate"
                          required>
                      </div>
                      <div>
                        <label for="cvv" class="block text-sm font-medium text-gray-700 mb-1">
                          CVV
                        </label>
                        <input 
                          type="text" 
                          id="cvv" 
                          v-model="newPaymentMethod.cvv" 
                          class="mt-1 focus:ring-accent focus:border-accent block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                          maxlength="4"
                          placeholder="•••"
                          required>
                      </div>
                    </div>
                  </div>
                  
                  <div class="mt-4">
                    <div class="relative flex items-start">
                      <div class="flex items-center h-5">
                        <input 
                          id="setDefault" 
                          name="setDefault" 
                          type="checkbox" 
                          v-model="newPaymentMethod.setDefault"
                          class="focus:ring-accent h-4 w-4 text-accent border-gray-300 rounded">
                      </div>
                      <div class="ml-3 text-sm">
                        <label for="setDefault" class="font-medium text-gray-700">Set as default payment method</label>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
            <button 
              type="button" 
              class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-accent text-base font-medium text-white hover:bg-accent/80 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-accent/30 sm:ml-3 sm:w-auto sm:text-sm transition-all duration-300"
              :disabled="!isPaymentFormValid || isSaving"
              @click="addPaymentMethod">
              <span v-if="isSaving" class="flex items-center">
                <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                Adding...
              </span>
              <span v-else>Add Payment Method</span>
            </button>
            <button 
              type="button" 
              class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-accent/30 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm transition-all duration-300"
              @click="showAddPaymentModal = false">
              Cancel
            </button>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Delete Account Modal -->
    <div v-if="showDeleteAccountModal" class="fixed inset-0 z-50 overflow-y-auto" aria-labelledby="modal-title" role="dialog" aria-modal="true">
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <div 
          class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" 
          aria-hidden="true"
          v-motion
          :initial="{ opacity: 0 }"
          :enter="{ opacity: 1, transition: { duration: 300 } }"
          @click="showDeleteAccountModal = false"></div>

        <span class="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="true">&#8203;</span>

        <div 
          class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full"
          v-motion
          :initial="{ opacity: 0, scale: 0.9 }"
          :enter="{ opacity: 1, scale: 1, transition: { duration: 300 } }">
          <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <div class="sm:flex sm:items-start">
              <div class="mx-auto flex-shrink-0 flex items-center justify-center h-12 w-12 rounded-full bg-red-100 sm:mx-0 sm:h-10 sm:w-10">
                <svg class="h-6 w-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"></path>
                </svg>
              </div>
              <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                <h3 class="text-lg leading-6 font-medium text-gray-900" id="modal-title">
                  Delete Account
                </h3>
                <div class="mt-2">
                  <p class="text-sm text-gray-500">
                    Are you sure you want to delete your account? All of your data will be permanently removed. This action cannot be undone.
                  </p>
                </div>
                <div class="mt-4">
                  <div class="mb-4">
                    <label for="deleteConfirmation" class="block text-sm font-medium text-gray-700 mb-1">
                      To confirm, type "DELETE" below
                    </label>
                    <input 
                      type="text" 
                      id="deleteConfirmation" 
                      v-model="deleteConfirmation" 
                      class="mt-1 focus:ring-red-500 focus:border-red-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                      placeholder="DELETE">
                  </div>
                  <div class="mb-4">
                    <label for="deleteReason" class="block text-sm font-medium text-gray-700 mb-1">
                      Please tell us why you're leaving (optional)
                    </label>
                    <select 
                      id="deleteReason" 
                      v-model="deleteReason"
                      class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-red-500 focus:border-red-500 sm:text-sm rounded-md">
                      <option value="">Select a reason</option>
                      <option value="not-useful">The platform isn't useful for me</option>
                      <option value="too-complex">Too complex to use</option>
                      <option value="better-alternative">Found a better alternative</option>
                      <option value="not-earning">Not earning enough</option>
                      <option value="privacy-concerns">Privacy concerns</option>
                      <option value="other">Other reason</option>
                    </select>
                  </div>
                  <div v-if="deleteReason === 'other'" class="mb-4">
                    <label for="otherDeleteReason" class="block text-sm font-medium text-gray-700 mb-1">
                      Please specify
                    </label>
                    <textarea 
                      id="otherDeleteReason" 
                      v-model="otherDeleteReason" 
                      rows="3" 
                      class="mt-1 focus:ring-red-500 focus:border-red-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                      placeholder="Tell us more..."></textarea>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
           <button 
              type="button" 
              class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-accent/30 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm transition-all duration-300"
              @click="showDeleteAccountModal = false">
              Cancel
            </button>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Cookie Preferences Modal -->
    <div v-if="showCookiePreferences" class="fixed inset-0 z-50 overflow-y-auto" aria-labelledby="modal-title" role="dialog" aria-modal="true">
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <div 
          class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" 
          aria-hidden="true"
          v-motion
          :initial="{ opacity: 0 }"
          :enter="{ opacity: 1, transition: { duration: 300 } }"
          @click="showCookiePreferences = false"></div>

        <span class="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="true">&#8203;</span>

        <div 
          class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full"
          v-motion
          :initial="{ opacity: 0, scale: 0.9 }"
          :enter="{ opacity: 1, scale: 1, transition: { duration: 300 } }">
          <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <div class="sm:flex sm:items-start">
              <div class="mx-auto flex-shrink-0 flex items-center justify-center h-12 w-12 rounded-full bg-blue-100 sm:mx-0 sm:h-10 sm:w-10">
                <svg class="h-6 w-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
              <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                <h3 class="text-lg leading-6 font-medium text-gray-900" id="modal-title">
                  Cookie Preferences
                </h3>
                <div class="mt-2">
                  <p class="text-sm text-gray-500">
                    Manage your cookie preferences. Essential cookies are always enabled as they are necessary for the website to function properly.
                  </p>
                </div>
                <div class="mt-4 space-y-4">
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">Essential Cookies</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Necessary for the website to function properly
                      </p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-not-allowed">
                        <input 
                          type="checkbox" 
                          checked
                          disabled
                          class="sr-only peer">
                        <div class="w-11 h-6 bg-accent peer-focus:outline-none rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all"></div>
                      </label>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">Functional Cookies</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Enhance your experience and provide personalized features
                      </p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="settings.cookies.functional" 
                          class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                      </label>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">Analytics Cookies</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Help us understand how you use our site
                      </p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="settings.cookies.analytics" 
                          class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                      </label>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between">
                    <div>
                      <p class="text-sm font-medium text-gray-900">Marketing Cookies</p>
                      <p class="text-xs text-gray-500 mt-1">
                        Allow us to show personalized advertisements
                      </p>
                    </div>
                    <div>
                      <label class="relative inline-flex items-center cursor-pointer">
                        <input 
                          type="checkbox" 
                          v-model="settings.cookies.marketing" 
                          class="sr-only peer">
                        <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-accent/30 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                      </label>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
            <button 
              type="button" 
              class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-accent text-base font-medium text-white hover:bg-accent/80 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-accent/30 sm:ml-3 sm:w-auto sm:text-sm transition-all duration-300"
              @click="saveCookiePreferences">
              <span v-if="isSaving" class="flex items-center">
                <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                Saving...
              </span>
              <span v-else>Save Preferences</span>
            </button>
            <button 
              type="button" 
              class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-accent/30 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm transition-all duration-300"
              @click="showCookiePreferences = false">
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

// Import icons for navigation
import {
  CreditCardIcon,
  PaintBrushIcon,
  BellIcon,
  ShieldCheckIcon,
  UserCircleIcon
} from '@heroicons/vue/24/outline';

// Stores
const referralStore = useReferralStore();
const notificationStore = useNotificationStore();

// State
const activeSection = ref('payment');
const isSaving = ref(false);
const isDeleting = ref(false);
const isExporting = ref(false);
const showAddPaymentModal = ref(false);
const showDeleteAccountModal = ref(false);
const showCookiePreferences = ref(false);
const deleteConfirmation = ref('');
const deleteReason = ref('');
const otherDeleteReason = ref('');
const showToast = ref(false);
const toastMessage = ref('');
const toastType = ref('success');
const toastTimeout = ref(null);

// Navigation sections
const settingsSections = [
  {
    id: 'payment',
    name: 'Payment Methods',
    icon: CreditCardIcon
  },
  {
    id: 'appearance',
    name: 'Appearance',
    icon: PaintBrushIcon
  },
  {
    id: 'notifications',
    name: 'Notifications',
    icon: BellIcon
  },
  {
    id: 'privacy',
    name: 'Privacy',
    icon: ShieldCheckIcon
  },
  {
    id: 'account',
    name: 'Account',
    icon: UserCircleIcon
  }
];

// Payment methods
const paymentMethods = ref([
  {
    id: 1,
    type: 'paypal',
    name: 'PayPal Account',
    email: 'sarah.johnson@example.com',
    default: true
  },
  {
    id: 2,
    type: 'bank',
    name: 'Chase Checking',
    bankName: 'Chase Bank',
    accountType: 'checking',
    accountLast4: '4567',
    default: false
  },
  {
    id: 3,
    type: 'card',
    name: 'VISA ending in 8912',
    cardBrand: 'visa',
    last4: '8912',
    expMonth: '06',
    expYear: '28',
    default: false
  }
]);

// New payment method form
const newPaymentMethod = ref({
  type: 'paypal',
  name: '',
  email: '',
  bankName: '',
  accountType: 'checking',
  routingNumber: '',
  accountNumber: '',
  username: '',
  cardNumber: '',
  expiryDate: '',
  cvv: '',
  setDefault: false
});

// Password form
const passwordForm = ref({
  currentPassword: '',
  newPassword: '',
  confirmPassword: ''
});

// Settings state
const settings = ref({
  automaticPayouts: true,
  payoutThreshold: 100,
  currency: 'USD',
  payoutFrequency: 'biweekly',
  theme: 'light',
  layout: 'comfortable',
  enableAnimations: true,
  reducedMotion: false,
  notifications: {
    email: {
      newReferral: true,
      bookingConfirmed: true,
      cashbackApproved: true,
      payoutProcessed: true,
      weeklySummary: false
    },
    push: {
      enabled: false,
      newReferral: true,
      bookingConfirmed: true,
      cashbackApproved: true,
      payoutProcessed: true
    },
    inApp: {
      sound: true
    }
  },
  privacy: {
    publicProfile: true,
    showEarnings: false,
    showReferralCount: true,
    activityTracking: true,
    marketingEmails: true
  },
  cookies: {
    functional: true,
    analytics: true,
    marketing: false
  },
  account: {
    fullName: 'Sarah Johnson',
    email: 'sarah.johnson@example.com',
    phone: '+1 (555) 123-4567',
    location: 'Los Angeles, CA',
    twoFactorEnabled: false,
    language: 'en',
    timezone: 'UTC-08:00'
  }
});

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

const isPaymentFormValid = computed(() => {
  const method = newPaymentMethod.value;
  
  if (!method.name) return false;
  
  if (method.type === 'paypal' && !method.email) {
    return false;
  } else if (method.type === 'bank' && 
    (!method.bankName || 
     !method.routingNumber || 
     !method.accountNumber)) {
    return false;
  } else if (method.type === 'venmo' && !method.username) {
    return false;
  } else if (method.type === 'card' && 
    (!method.cardNumber || 
     !method.expiryDate || 
     !method.cvv)) {
    return false;
  }
  
  return true;
});

// Methods
const savePayoutPreferences = async () => {
  isSaving.value = true;
  
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1500));
    
    showToastNotification('Payout preferences saved successfully', 'success');
  } catch (error) {
    showToastNotification('Failed to save payout preferences', 'error');
  } finally {
    isSaving.value = false;
  }
};

const saveAppearanceSettings = async () => {
  isSaving.value = true;
  
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1500));
    
    showToastNotification('Appearance settings saved successfully', 'success');
  } catch (error) {
    showToastNotification('Failed to save appearance settings', 'error');
  } finally {
    isSaving.value = false;
  }
};

const saveNotificationSettings = async () => {
  isSaving.value = true;
  
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1500));
    
    showToastNotification('Notification settings saved successfully', 'success');
  } catch (error) {
    showToastNotification('Failed to save notification settings', 'error');
  } finally {
    isSaving.value = false;
  }
};

const savePrivacySettings = async () => {
  isSaving.value = true;
  
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1500));
    
    showToastNotification('Privacy settings saved successfully', 'success');
  } catch (error) {
    showToastNotification('Failed to save privacy settings', 'error');
  } finally {
    isSaving.value = false;
  }
};

const saveAccountSettings = async () => {
  isSaving.value = true;
  
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1500));
    
    showToastNotification('Account settings saved successfully', 'success');
  } catch (error) {
    showToastNotification('Failed to save account settings', 'error');
  } finally {
    isSaving.value = false;
  }
};

const saveCookiePreferences = async () => {
  isSaving.value = true;
  
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1500));
    
    showCookiePreferences.value = false;
    showToastNotification('Cookie preferences saved successfully', 'success');
  } catch (error) {
    showToastNotification('Failed to save cookie preferences', 'error');
  } finally {
    isSaving.value = false;
  }
};

const addPaymentMethod = async () => {
  isSaving.value = true;
  
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1500));
    
    const method = { ...newPaymentMethod.value };
    const newId = paymentMethods.value.length + 1;
    
    // Process by type
    if (method.type === 'paypal') {
      paymentMethods.value.push({
        id: newId,
        type: 'paypal',
        name: method.name,
        email: method.email,
        default: method.setDefault
      });
    } else if (method.type === 'bank') {
      // Extract last 4 digits of account number
      const accountLast4 = method.accountNumber.slice(-4);
      
      paymentMethods.value.push({
        id: newId,
        type: 'bank',
        name: method.name,
        bankName: method.bankName,
        accountType: method.accountType,
        accountLast4,
        default: method.setDefault
      });
    } else if (method.type === 'venmo') {
      paymentMethods.value.push({
        id: newId,
        type: 'venmo',
        name: method.name,
        username: method.username,
        default: method.setDefault
      });
    } else if (method.type === 'card') {
      // Extract card data
      const cardNumber = method.cardNumber.replace(/\s/g, '');
      const last4 = cardNumber.slice(-4);
      
      // Determine card brand based on first digit
      let cardBrand = 'visa';
      const firstDigit = cardNumber.charAt(0);
      if (firstDigit === '4') {
        cardBrand = 'visa';
      } else if (firstDigit === '5') {
        cardBrand = 'mastercard';
      } else if (firstDigit === '3') {
        cardBrand = 'amex';
      }
      
      // Parse expiry date
      const [expMonth, expYear] = method.expiryDate.split('/');
      
      paymentMethods.value.push({
        id: newId,
        type: 'card',
        name: method.name,
        cardBrand,
        last4,
        expMonth,
        expYear,
        default: method.setDefault
      });
    }
    
    // If this method is set as default, update other methods
    if (method.setDefault) {
      paymentMethods.value.forEach(m => {
        if (m.id !== newId) {
          m.default = false;
        }
      });
    }
    
    // Reset form and close modal
    resetPaymentForm();
    showAddPaymentModal.value = false;
    
    showToastNotification('Payment method added successfully', 'success');
  } catch (error) {
    showToastNotification('Failed to add payment method', 'error');
  } finally {
    isSaving.value = false;
  }
};

const resetPaymentForm = () => {
  newPaymentMethod.value = {
    type: 'paypal',
    name: '',
    email: '',
    bankName: '',
    accountType: 'checking',
    routingNumber: '',
    accountNumber: '',
    username: '',
    cardNumber: '',
    expiryDate: '',
    cvv: '',
    setDefault: false
  };
};

const editPaymentMethod = (method) => {
  // Implementation would be similar to addPaymentMethod
  // but would prefill the form with existing data
  showToastNotification('Edit payment method functionality coming soon', 'info');
};

const removePaymentMethod = async (index) => {
  const method = paymentMethods.value[index];
  
  // Don't allow removing the default method
  if (method.default) {
    showToastNotification('Cannot remove default payment method', 'error');
    return;
  }
  
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 800));
    
    // Remove the method
    paymentMethods.value.splice(index, 1);
    
    showToastNotification('Payment method removed successfully', 'success');
  } catch (error) {
    showToastNotification('Failed to remove payment method', 'error');
  }
};

const setDefaultPaymentMethod = async (index) => {
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 800));
    
    // Update default status
    paymentMethods.value.forEach((method, i) => {
      method.default = i === index;
    });
    
    showToastNotification('Default payment method updated', 'success');
  } catch (error) {
    showToastNotification('Failed to update default payment method', 'error');
  }
};

const changePassword = async () => {
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
    
    showToastNotification('Password changed successfully', 'success');
  } catch (error) {
    showToastNotification('Failed to change password', 'error');
  } finally {
    isSaving.value = false;
  }
};

const setupTwoFactor = () => {
  // Show 2FA setup modal (not implemented in this demo)
  showToastNotification('Two-factor authentication setup initiated', 'info');
  settings.value.account.twoFactorEnabled = true;
};

const disableTwoFactor = () => {
  // Show 2FA disable confirmation (not implemented in this demo)
  showToastNotification('Two-factor authentication disabled', 'info');
  settings.value.account.twoFactorEnabled = false;
};

const toggleTwoFactor = () => {
  if (settings.value.account.twoFactorEnabled) {
    disableTwoFactor();
  } else {
    setupTwoFactor();
  }
};

const formatCardNumber = () => {
  // Format card number with spaces after every 4 digits
  let value = newPaymentMethod.value.cardNumber.replace(/\s/g, '');
  let formattedValue = '';
  
  for (let i = 0; i < value.length; i++) {
    if (i > 0 && i % 4 === 0) {
      formattedValue += ' ';
    }
    formattedValue += value[i];
  }
  
  newPaymentMethod.value.cardNumber = formattedValue;
};

const formatExpiryDate = () => {
  // Format expiry date as MM/YY
  let value = newPaymentMethod.value.expiryDate.replace(/\D/g, '');
  
  if (value.length > 2) {
    value = value.substring(0, 2) + '/' + value.substring(2, 4);
  }
  
  newPaymentMethod.value.expiryDate = value;
};

const requestDataExport = async () => {
  isExporting.value = true;
  
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 2000));
    
    showToastNotification('Data export request submitted. You will receive an email with your data shortly.', 'success');
  } catch (error) {
    showToastNotification('Failed to request data export', 'error');
  } finally {
    isExporting.value = false;
  }
};

const deleteAccount = async () => {
  isDeleting.value = true;
  
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 2000));
    
    // Would normally redirect to logout or confirmation page
    showDeleteAccountModal.value = false;
    showToastNotification('Account deletion process initiated. You will receive a confirmation email.', 'info');
    
    // Reset form
    deleteConfirmation.value = '';
    deleteReason.value = '';
    otherDeleteReason.value = '';
  } catch (error) {
    showToastNotification('Failed to delete account', 'error');
  } finally {
    isDeleting.value = false;
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

// Lifecycle hooks
onMounted(() => {
  // Fetch settings data (would normally be from API)
  // Simulate API call
  setTimeout(() => {
    // Data would be fetched from API here
  }, 800);
});
</script>

<style scoped>
.settings-page {
  @apply p-4 md:p-6 max-w-7xl mx-auto;
}
</style>