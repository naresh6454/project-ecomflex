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
          Settings
        </h1>
        <p class="mt-1 text-sm text-gray-500">
          Manage your account settings and preferences.
        </p>
      </div>
    </div>

    <div class="grid grid-cols-1 gap-6 lg:grid-cols-4">
      <!-- Settings Navigation -->
      <div 
        class="bg-white shadow rounded-lg overflow-hidden"
        v-motion
        :initial="{ opacity: 0, x: -20 }"
        :enter="{ opacity: 1, x: 0 }">
        <div class="p-4">
          <h2 class="text-lg font-medium text-gray-900 mb-2">Settings</h2>
          <nav class="space-y-1">
            <button 
              v-for="tab in tabs" 
              :key="tab.id"
              @click="activeTab = tab.id"
              class="flex items-center px-3 py-2 text-sm font-medium rounded-md w-full text-left transition-colors"
              :class="activeTab === tab.id ? 'bg-blue-50 text-blue-700' : 'text-gray-600 hover:bg-gray-50 hover:text-gray-900'">
              <component 
                :is="tab.icon" 
                class="mr-3 flex-shrink-0 h-5 w-5" 
                :class="activeTab === tab.id ? 'text-blue-500' : 'text-gray-400'" 
                aria-hidden="true" />
              {{ tab.name }}
            </button>
          </nav>
        </div>
      </div>

      <!-- Settings Content -->
      <div 
        class="lg:col-span-3"
        v-motion
        :initial="{ opacity: 0, x: 20 }"
        :enter="{ opacity: 1, x: 0 }">
        <div class="bg-white shadow rounded-lg">
          <!-- Account Settings -->
          <div v-if="activeTab === 'account'" class="p-6">
            <h3 class="text-lg font-medium text-gray-900 mb-4">Account Settings</h3>
            
            <!-- Email Preferences -->
            <div class="mb-8">
              <h4 class="text-sm font-medium text-gray-900 mb-3">Email Preferences</h4>
              <div class="space-y-4">
                <div class="flex items-start">
                  <div class="flex items-center h-5">
                    <input 
                      id="email-marketing" 
                      type="checkbox" 
                      v-model="settings.emailPreferences.marketing"
                      class="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300 rounded"
                    />
                  </div>
                  <div class="ml-3 text-sm">
                    <label for="email-marketing" class="font-medium text-gray-700">Marketing Emails</label>
                    <p class="text-gray-500">Receive emails about new products, promotions, and platform news.</p>
                  </div>
                </div>
                
                <div class="flex items-start">
                  <div class="flex items-center h-5">
                    <input 
                      id="email-product" 
                      type="checkbox" 
                      v-model="settings.emailPreferences.productUpdates"
                      class="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300 rounded"
                    />
                  </div>
                  <div class="ml-3 text-sm">
                    <label for="email-product" class="font-medium text-gray-700">Product Updates</label>
                    <p class="text-gray-500">Get notified when new products are available for promotion.</p>
                  </div>
                </div>
                
                <div class="flex items-start">
                  <div class="flex items-center h-5">
                    <input 
                      id="email-referral" 
                      type="checkbox" 
                      v-model="settings.emailPreferences.referralUpdates"
                      class="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300 rounded"
                    />
                  </div>
                  <div class="ml-3 text-sm">
                    <label for="email-referral" class="font-medium text-gray-700">Referral Updates</label>
                    <p class="text-gray-500">Get notified when your referrals make bookings or are approved.</p>
                  </div>
                </div>
              </div>
            </div>
            
            <!-- Password -->
            <div class="mb-8">
              <h4 class="text-sm font-medium text-gray-900 mb-3">Password</h4>
              <div class="space-y-4">
                <div>
                  <label for="current-password" class="block text-sm font-medium text-gray-700">Current Password</label>
                  <div class="mt-1">
                    <input 
                      type="password" 
                      id="current-password" 
                      class="shadow-sm focus:ring-blue-500 focus:border-blue-500 block w-full sm:text-sm border-gray-300 rounded-md"
                      placeholder="••••••••"
                    />
                  </div>
                </div>
                
                <div>
                  <label for="new-password" class="block text-sm font-medium text-gray-700">New Password</label>
                  <div class="mt-1">
                    <input 
                      type="password" 
                      id="new-password" 
                      class="shadow-sm focus:ring-blue-500 focus:border-blue-500 block w-full sm:text-sm border-gray-300 rounded-md"
                      placeholder="••••••••"
                    />
                  </div>
                </div>
                
                <div>
                  <label for="confirm-password" class="block text-sm font-medium text-gray-700">Confirm New Password</label>
                  <div class="mt-1">
                    <input 
                      type="password" 
                      id="confirm-password" 
                      class="shadow-sm focus:ring-blue-500 focus:border-blue-500 block w-full sm:text-sm border-gray-300 rounded-md"
                      placeholder="••••••••"
                    />
                  </div>
                </div>
                
                <div>
                  <button 
                    type="button"
                    class="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors">
                    Update Password
                  </button>
                </div>
              </div>
            </div>
            
            <!-- Two-Factor Authentication -->
            <div class="mb-8">
              <div class="flex items-center justify-between mb-3">
                <h4 class="text-sm font-medium text-gray-900">Two-Factor Authentication</h4>
                <span 
                  class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                  :class="settings.twoFactorEnabled ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'">
                  {{ settings.twoFactorEnabled ? 'Enabled' : 'Disabled' }}
                </span>
              </div>
              <p class="text-sm text-gray-500 mb-4">
                Add an extra layer of security to your account by requiring more than just a password to sign in.
              </p>
              <button 
                type="button"
                class="inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors">
                {{ settings.twoFactorEnabled ? 'Disable' : 'Enable' }} Two-Factor Authentication
              </button>
            </div>
            
            <!-- Danger Zone -->
            <div>
              <h4 class="text-sm font-medium text-red-600 mb-3">Danger Zone</h4>
              <div class="rounded-md bg-red-50 p-4">
                <div class="flex">
                  <div class="flex-shrink-0">
                    <WarningIcon class="h-5 w-5 text-red-400" />
                  </div>
                  <div class="ml-3">
                    <h3 class="text-sm font-medium text-red-800">Deactivate Account</h3>
                    <div class="mt-2 text-sm text-red-700">
                      <p>
                        This will deactivate your account and remove your profile from the platform. All your data will be preserved, and you can reactivate at any time.
                      </p>
                    </div>
                    <div class="mt-4">
                      <button 
                        type="button"
                        class="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 transition-colors">
                        Deactivate Account
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          
          <!-- Notification Settings -->
          <div v-if="activeTab === 'notifications'" class="p-6">
            <h3 class="text-lg font-medium text-gray-900 mb-4">Notification Settings</h3>
            
            <div class="space-y-6">
              <!-- Email Notifications -->
              <div>
                <h4 class="text-sm font-medium text-gray-900 mb-3">Email Notifications</h4>
                <div class="space-y-4">
                  <div v-for="(setting, key) in settings.notifications.email" :key="`email-${key}`" class="flex items-start">
                    <div class="flex items-center h-5">
                      <input 
                        :id="`email-${key}`" 
                        type="checkbox" 
                        v-model="settings.notifications.email[key]"
                        class="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300 rounded"
                      />
                    </div>
                    <div class="ml-3 text-sm">
                      <label :for="`email-${key}`" class="font-medium text-gray-700">{{ formatLabel(key) }}</label>
                    </div>
                  </div>
                </div>
              </div>
              
              <!-- Push Notifications -->
              <div>
                <h4 class="text-sm font-medium text-gray-900 mb-3">Push Notifications</h4>
                <div class="space-y-4">
                  <div v-for="(setting, key) in settings.notifications.push" :key="`push-${key}`" class="flex items-start">
                    <div class="flex items-center h-5">
                      <input 
                        :id="`push-${key}`" 
                        type="checkbox" 
                        v-model="settings.notifications.push[key]"
                        class="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300 rounded"
                      />
                    </div>
                    <div class="ml-3 text-sm">
                      <label :for="`push-${key}`" class="font-medium text-gray-700">{{ formatLabel(key) }}</label>
                    </div>
                  </div>
                </div>
              </div>
              
              <!-- Notification Frequency -->
              <div>
                <h4 class="text-sm font-medium text-gray-900 mb-3">Notification Frequency</h4>
                <div>
                  <select 
                    id="frequency" 
                    v-model="settings.notifications.frequency"
                    class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm rounded-md">
                    <option value="realtime">Real-time</option>
                    <option value="daily">Daily Digest</option>
                    <option value="weekly">Weekly Digest</option>
                  </select>
                </div>
              </div>
            </div>
            
            <div class="mt-6">
              <button 
                type="button"
                class="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors">
                Save Notification Preferences
              </button>
            </div>
          </div>
          
          <!-- Privacy Settings -->
          <div v-if="activeTab === 'privacy'" class="p-6">
            <h3 class="text-lg font-medium text-gray-900 mb-4">Privacy Settings</h3>
            
            <div class="space-y-6">
              <!-- Profile Visibility -->
              <div>
                <h4 class="text-sm font-medium text-gray-900 mb-3">Profile Visibility</h4>
                <div class="space-y-4">
                  <div class="flex items-start">
                    <div class="flex items-center h-5">
                      <input 
                        id="visibility-public" 
                        type="checkbox" 
                        v-model="settings.privacy.profilePublic"
                        class="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300 rounded"
                      />
                    </div>
                    <div class="ml-3 text-sm">
                      <label for="visibility-public" class="font-medium text-gray-700">Public Profile</label>
                      <p class="text-gray-500">Make your profile visible to everyone, including non-members.</p>
                    </div>
                  </div>
                  
                  <div class="flex items-start">
                    <div class="flex items-center h-5">
                      <input 
                        id="visibility-searchable" 
                        type="checkbox" 
                        v-model="settings.privacy.searchable"
                        class="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300 rounded"
                      />
                    </div>
                    <div class="ml-3 text-sm">
                      <label for="visibility-searchable" class="font-medium text-gray-700">Searchable</label>
                      <p class="text-gray-500">Allow your profile to appear in search results.</p>
                    </div>
                  </div>
                </div>
              </div>
              
              <!-- Data Usage -->
              <div>
                <h4 class="text-sm font-medium text-gray-900 mb-3">Data Usage</h4>
                <div class="space-y-4">
                  <div class="flex items-start">
                    <div class="flex items-center h-5">
                      <input 
                        id="data-analytics" 
                        type="checkbox" 
                        v-model="settings.privacy.dataAnalytics"
                        class="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300 rounded"
                      />
                    </div>
                    <div class="ml-3 text-sm">
                      <label for="data-analytics" class="font-medium text-gray-700">Usage Analytics</label>
                      <p class="text-gray-500">Allow us to collect anonymous usage data to improve our platform.</p>
                    </div>
                  </div>
                  
                  <div class="flex items-start">
                    <div class="flex items-center h-5">
                      <input 
                        id="data-personalization" 
                        type="checkbox" 
                        v-model="settings.privacy.personalization"
                        class="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300 rounded"
                      />
                    </div>
                    <div class="ml-3 text-sm">
                      <label for="data-personalization" class="font-medium text-gray-700">Personalization</label>
                      <p class="text-gray-500">Allow us to use your data for personalized recommendations.</p>
                    </div>
                  </div>
                </div>
              </div>
              
              <!-- Connected Accounts -->
              <div>
                <h4 class="text-sm font-medium text-gray-900 mb-3">Connected Accounts</h4>
                <div class="space-y-4">
                  <div class="flex items-center justify-between py-3 border-b border-gray-200">
                    <div class="flex items-center">
                      <FacebookIcon class="h-6 w-6 text-blue-500" />
                      <div class="ml-3">
                        <h5 class="text-sm font-medium text-gray-900">Facebook</h5>
                        <p class="text-xs text-gray-500">Connected as Sarah Johnson</p>
                      </div>
                    </div>
                    <button 
                      type="button"
                      class="text-xs font-medium text-red-600 hover:text-red-900 transition-colors">
                      Disconnect
                    </button>
                  </div>
                  
                  <div class="flex items-center justify-between py-3 border-b border-gray-200">
                    <div class="flex items-center">
                      <TwitterIcon class="h-6 w-6 text-blue-400" />
                      <div class="ml-3">
                        <h5 class="text-sm font-medium text-gray-900">Twitter</h5>
                        <p class="text-xs text-gray-500">Connected as @sarahjohnson</p>
                      </div>
                    </div>
                    <button 
                      type="button"
                      class="text-xs font-medium text-red-600 hover:text-red-900 transition-colors">
                      Disconnect
                    </button>
                  </div>
                  
                  <div class="flex items-center justify-between py-3">
                    <div class="flex items-center">
                      <GoogleIcon class="h-6 w-6 text-gray-400" />
                      <div class="ml-3">
                        <h5 class="text-sm font-medium text-gray-900">Google</h5>
                        <p class="text-xs text-gray-500">Not connected</p>
                      </div>
                    </div>
                    <button 
                      type="button"
                      class="text-xs font-medium text-blue-600 hover:text-blue-900 transition-colors">
                      Connect
                    </button>
                  </div>
                </div>
              </div>
            </div>
            
            <div class="mt-6">
              <button 
                type="button"
                class="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors">
                Save Privacy Settings
              </button>
            </div>
          </div>
          
          <!-- Appearance Settings -->
          <div v-if="activeTab === 'appearance'" class="p-6">
            <h3 class="text-lg font-medium text-gray-900 mb-4">Appearance Settings</h3>
            
            <div class="space-y-6">
              <!-- Theme -->
              <div>
                <h4 class="text-sm font-medium text-gray-900 mb-3">Theme</h4>
                <div class="grid grid-cols-3 gap-4">
                  <button 
                    @click="settings.appearance.theme = 'light'"
                    class="relative rounded-lg border overflow-hidden p-2 flex flex-col items-center transition-colors"
                    :class="settings.appearance.theme === 'light' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'">
                    <div class="h-24 w-full bg-white rounded shadow-sm"></div>
                    <div class="mt-2 text-sm font-medium" :class="settings.appearance.theme === 'light' ? 'text-blue-600' : 'text-gray-900'">Light</div>
                    <div v-if="settings.appearance.theme === 'light'" class="absolute top-1 right-1 h-4 w-4 rounded-full bg-blue-500 flex items-center justify-center">
                      <CheckIcon class="h-3 w-3 text-white" />
                    </div>
                  </button>
                  
                  <button 
                    @click="settings.appearance.theme = 'dark'"
                    class="relative rounded-lg border overflow-hidden p-2 flex flex-col items-center transition-colors"
                    :class="settings.appearance.theme === 'dark' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'">
                    <div class="h-24 w-full bg-gray-900 rounded shadow-sm"></div>
                    <div class="mt-2 text-sm font-medium" :class="settings.appearance.theme === 'dark' ? 'text-blue-600' : 'text-gray-900'">Dark</div>
                    <div v-if="settings.appearance.theme === 'dark'" class="absolute top-1 right-1 h-4 w-4 rounded-full bg-blue-500 flex items-center justify-center">
                      <CheckIcon class="h-3 w-3 text-white" />
                    </div>
                  </button>
                  
                  <button 
                    @click="settings.appearance.theme = 'system'"
                    class="relative rounded-lg border overflow-hidden p-2 flex flex-col items-center transition-colors"
                    :class="settings.appearance.theme === 'system' ? 'border-blue-500 bg-blue-50' : 'border-gray-200 hover:border-gray-300'">
                    <div class="h-24 w-full bg-gradient-to-r from-white to-gray-900 rounded shadow-sm"></div>
                    <div class="mt-2 text-sm font-medium" :class="settings.appearance.theme === 'system' ? 'text-blue-600' : 'text-gray-900'">System</div>
                    <div v-if="settings.appearance.theme === 'system'" class="absolute top-1 right-1 h-4 w-4 rounded-full bg-blue-500 flex items-center justify-center">
                      <CheckIcon class="h-3 w-3 text-white" />
                    </div>
                  </button>
                </div>
              </div>
              
              <!-- Accent Color -->
              <div>
                <h4 class="text-sm font-medium text-gray-900 mb-3">Accent Color</h4>
                <div class="grid grid-cols-5 gap-3">
                  <button 
                    v-for="color in accentColors" 
                    :key="color.id"
                    @click="settings.appearance.accentColor = color.id"
                    class="relative rounded-full h-8 w-8 border-2 flex items-center justify-center transition-colors"
                    :class="[
                      color.class, 
                      settings.appearance.accentColor === color.id ? 'ring-2 ring-offset-2 ring-gray-300' : ''
                    ]">
                    <CheckIcon v-if="settings.appearance.accentColor === color.id" class="h-4 w-4 text-white" />
                  </button>
                </div>
              </div>
              
              <!-- Animations -->
              <div>
                <h4 class="text-sm font-medium text-gray-900 mb-3">Animations</h4>
                <div class="space-y-4">
                  <div class="flex items-start">
                    <div class="flex items-center h-5">
                      <input 
                        id="animations-enabled" 
                        type="checkbox" 
                        v-model="settings.appearance.animationsEnabled"
                        class="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300 rounded"
                      />
                    </div>
                    <div class="ml-3 text-sm">
                      <label for="animations-enabled" class="font-medium text-gray-700">Enable Animations</label>
                      <p class="text-gray-500">Show motion animations throughout the interface.</p>
                    </div>
                  </div>
                  
                  <div class="flex items-start">
                    <div class="flex items-center h-5">
                      <input 
                        id="reduce-motion" 
                        type="checkbox" 
                        v-model="settings.appearance.reduceMotion"
                        class="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300 rounded"
                      />
                    </div>
                    <div class="ml-3 text-sm">
                      <label for="reduce-motion" class="font-medium text-gray-700">Reduce Motion</label>
                      <p class="text-gray-500">Minimize animations for accessibility purposes.</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            
            <div class="mt-6">
              <button 
                type="button"
                class="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors">
                Save Appearance Settings
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, defineComponent, h } from 'vue'

// Icon components - converted to proper Vue render functions
const UserIcon = defineComponent({
  props: ['class'],
  render() {
    return h('svg', {
      xmlns: 'http://www.w3.org/2000/svg',
      fill: 'none',
      viewBox: '0 0 24 24',
      'stroke-width': '1.5',
      stroke: 'currentColor',
      class: this.class
    }, [
      h('path', {
        'stroke-linecap': 'round',
        'stroke-linejoin': 'round',
        d: 'M15.75 6a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0zM4.501 20.118a7.5 7.5 0 0114.998 0A17.933 17.933 0 0112 21.75c-2.676 0-5.216-.584-7.499-1.632z'
      })
    ])
  }
})

const BellIcon = defineComponent({
  props: ['class'],
  render() {
    return h('svg', {
      xmlns: 'http://www.w3.org/2000/svg',
      fill: 'none',
      viewBox: '0 0 24 24',
      'stroke-width': '1.5',
      stroke: 'currentColor',
      class: this.class
    }, [
      h('path', {
        'stroke-linecap': 'round',
        'stroke-linejoin': 'round',
        d: 'M14.857 17.082a23.848 23.848 0 005.454-1.31A8.967 8.967 0 0118 9.75v-.7V9A6 6 0 006 9v.75a8.967 8.967 0 01-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 01-5.714 0m5.714 0a3 3 0 11-5.714 0'
      })
    ])
  }
})

const LockIcon = defineComponent({
  props: ['class'],
  render() {
    return h('svg', {
      xmlns: 'http://www.w3.org/2000/svg',
      fill: 'none',
      viewBox: '0 0 24 24',
      'stroke-width': '1.5',
      stroke: 'currentColor',
      class: this.class
    }, [
      h('path', {
        'stroke-linecap': 'round',
        'stroke-linejoin': 'round',
        d: 'M16.5 10.5V6.75a4.5 4.5 0 10-9 0v3.75m-.75 11.25h10.5a2.25 2.25 0 002.25-2.25v-6.75a2.25 2.25 0 00-2.25-2.25H6.75a2.25 2.25 0 00-2.25 2.25v6.75a2.25 2.25 0 002.25 2.25z'
      })
    ])
  }
})

const PaletteIcon = defineComponent({
  props: ['class'],
  render() {
    return h('svg', {
      xmlns: 'http://www.w3.org/2000/svg',
      fill: 'none',
      viewBox: '0 0 24 24',
      'stroke-width': '1.5',
      stroke: 'currentColor',
      class: this.class
    }, [
      h('path', {
        'stroke-linecap': 'round',
        'stroke-linejoin': 'round',
        d: 'M9.53 16.122a3 3 0 00-5.78 1.128 2.25 2.25 0 01-2.4 2.245 4.5 4.5 0 008.4-2.245c0-.399-.078-.78-.22-1.128zm0 0a15.998 15.998 0 003.388-1.62m-5.043-.025a15.994 15.994 0 011.622-3.395m3.42 3.42a15.995 15.995 0 004.764-4.648l3.876-5.814a1.151 1.151 0 00-1.597-1.597L14.146 6.32a15.996 15.996 0 00-4.649 4.763m3.42 3.42a6.776 6.776 0 00-3.42-3.42'
      })
    ])
  }
})

const WarningIcon = defineComponent({
  props: ['class'],
  render() {
    return h('svg', {
      xmlns: 'http://www.w3.org/2000/svg',
      viewBox: '0 0 20 20',
      fill: 'currentColor',
      class: this.class
    }, [
      h('path', {
        'fill-rule': 'evenodd',
        d: 'M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z',
        'clip-rule': 'evenodd'
      })
    ])
  }
})

const CheckIcon = defineComponent({
  props: ['class'],
  render() {
    return h('svg', {
      xmlns: 'http://www.w3.org/2000/svg',
      viewBox: '0 0 20 20',
      fill: 'currentColor',
      class: this.class
    }, [
      h('path', {
        'fill-rule': 'evenodd',
        d: 'M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z',
        'clip-rule': 'evenodd'
      })
    ])
  }
})

const FacebookIcon = defineComponent({
  props: ['class'],
  render() {
    return h('svg', {
      class: this.class,
      fill: 'currentColor',
      viewBox: '0 0 24 24'
    }, [
      h('path', {
        d: 'M12 2.03998C6.5 2.03998 2 6.52998 2 12.06C2 17.06 5.66 21.21 10.44 21.96V14.96H7.9V12.06H10.44V9.84998C10.44 7.33998 11.93 5.95998 14.22 5.95998C15.31 5.95998 16.45 6.14998 16.45 6.14998V8.61998H15.19C13.95 8.61998 13.56 9.38998 13.56 10.18V12.06H16.34L15.89 14.96H13.56V21.96C15.9164 21.5878 18.0622 20.3855 19.6099 18.57C21.1576 16.7546 22.0054 14.4456 22 12.06C22 6.52998 17.5 2.03998 12 2.03998Z'
      })
    ])
  }
})

const TwitterIcon = defineComponent({
  props: ['class'],
  render() {
    return h('svg', {
      class: this.class,
      fill: 'currentColor',
      viewBox: '0 0 24 24'
    }, [
      h('path', {
        d: 'M22.46 6c-.77.35-1.6.58-2.46.69.88-.53 1.56-1.37 1.88-2.38-.83.5-1.75.85-2.72 1.05C18.37 4.5 17.26 4 16 4c-2.35 0-4.27 1.92-4.27 4.29 0 .34.04.67.11.98C8.28 9.09 5.11 7.38 3 4.79c-.37.63-.58 1.37-.58 2.15 0 1.49.75 2.81 1.91 3.56-.71 0-1.37-.2-1.95-.5v.03c0 2.08 1.48 3.82 3.44 4.21a4.22 4.22 0 0 1-1.93.07 4.28 4.28 0 0 0 4 2.98 8.521 8.521 0 0 1-5.33 1.84c-.34 0-.68-.02-1.02-.06C3.44 20.29 5.7 21 8.12 21 16 21 20.33 14.46 20.33 8.79c0-.19 0-.37-.01-.56.84-.6 1.56-1.36 2.14-2.23z'
      })
    ])
  }
})

const GoogleIcon = defineComponent({
  props: ['class'],
  render() {
    return h('svg', {
      class: this.class,
      fill: 'currentColor',
      viewBox: '0 0 24 24'
    }, [
      h('path', {
        d: 'M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 3c1.66 0 3 1.34 3 3s-1.34 3-3 3-3-1.34-3-3 1.34-3 3-3zm0 14.2c-2.5 0-4.71-1.28-6-3.22.03-1.99 4-3.08 6-3.08 1.99 0 5.97 1.09 6 3.08-1.29 1.94-3.5 3.22-6 3.22z'
      })
    ])
  }
})

// Settings tabs
const tabs = [
  { id: 'account', name: 'Account', icon: UserIcon },
  { id: 'notifications', name: 'Notifications', icon: BellIcon },
  { id: 'privacy', name: 'Privacy', icon: LockIcon },
  { id: 'appearance', name: 'Appearance', icon: PaletteIcon }
]

// Active tab
const activeTab = ref('account')

// Settings data
const settings = ref({
  emailPreferences: {
    marketing: true,
    productUpdates: true,
    referralUpdates: true
  },
  twoFactorEnabled: false,
  notifications: {
    email: {
      newReferrals: true,
      bookingConfirmations: true,
      proofApprovals: true,
      payouts: true,
      productUpdates: false,
      platformNews: false
    },
    push: {
      newReferrals: true,
      bookingConfirmations: true,
      proofApprovals: true,
      payouts: true,
      productUpdates: false,
      platformNews: false
    },
    frequency: 'realtime'
  },
  privacy: {
    profilePublic: true,
    searchable: true,
    dataAnalytics: true,
    personalization: true
  },
  appearance: {
    theme: 'light',
    accentColor: 'blue',
    animationsEnabled: true,
    reduceMotion: false
  }
})

// Accent colors
const accentColors = [
  { id: 'blue', class: 'bg-blue-600 border-blue-600' },
  { id: 'green', class: 'bg-green-600 border-green-600' },
  { id: 'purple', class: 'bg-purple-600 border-purple-600' },
  { id: 'red', class: 'bg-red-600 border-red-600' },
  { id: 'yellow', class: 'bg-yellow-500 border-yellow-500' }
]

// Format label function
const formatLabel = (key: string) => {
  return key
    .replace(/([A-Z])/g, ' $1')
    .replace(/^./, (str) => str.toUpperCase())
    .replace(/([a-z])([A-Z])/g, '$1 $2')
}
</script>