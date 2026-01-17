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
          Earnings & Payouts
        </h1>
        <p class="mt-1 text-sm text-gray-500">
          Track your earnings, view payout history, and manage payment methods.
        </p>
      </div>
      <div class="mt-4 flex md:mt-0 md:ml-4">
        <button 
          type="button" 
          class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors">
          <svg class="-ml-1 mr-2 h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
          </svg>
          Download Report
        </button>
      </div>
    </div>

    <!-- Earnings Overview Card -->
    <div 
      v-motion
      :initial="{ opacity: 0, scale: 0.95 }"
      :enter="{ opacity: 1, scale: 1 }"
      class="bg-gradient-to-r from-blue-500 to-blue-600 rounded-lg shadow-lg mb-8 text-white overflow-hidden">
      <div class="p-6">
        <div class="flex flex-col md:flex-row md:items-center md:justify-between">
          <div>
            <h2 class="text-xl font-semibold">Total Earnings</h2>
            <div class="mt-2 flex items-baseline">
              <p class="text-4xl font-bold">${{ totalEarnings }}</p>
              <p class="ml-2 text-sm opacity-80">USD</p>
            </div>
          </div>
          <div class="mt-4 md:mt-0">
            <div class="flex items-center">
              <span class="text-sm font-medium">Available for Payout:</span>
              <span class="ml-2 text-xl font-bold">${{ availableForPayout }}</span>
            </div>
            <div class="mt-2">
              <button 
                type="button"
                class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-blue-600 bg-white hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors">
                Request Payout
              </button>
            </div>
          </div>
        </div>

        <div class="mt-6 grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-4">
          <div class="bg-white bg-opacity-10 rounded-lg p-4">
            <div class="flex items-center">
              <div class="flex-shrink-0">
                <svg class="h-6 w-6 text-blue-100" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
                </svg>
              </div>
              <div class="ml-3">
                <p class="text-sm font-medium text-blue-100">Earnings This Month</p>
                <p class="text-lg font-semibold">${{ thisMonthEarnings }}</p>
              </div>
            </div>
          </div>

          <div class="bg-white bg-opacity-10 rounded-lg p-4">
            <div class="flex items-center">
              <div class="flex-shrink-0">
                <svg class="h-6 w-6 text-blue-100" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                </svg>
              </div>
              <div class="ml-3">
                <p class="text-sm font-medium text-blue-100">Earnings Last Month</p>
                <p class="text-lg font-semibold">${{ lastMonthEarnings }}</p>
              </div>
            </div>
          </div>

          <div class="bg-white bg-opacity-10 rounded-lg p-4">
            <div class="flex items-center">
              <div class="flex-shrink-0">
                <svg class="h-6 w-6 text-blue-100" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                </svg>
              </div>
              <div class="ml-3">
                <p class="text-sm font-medium text-blue-100">Referrals This Month</p>
                <p class="text-lg font-semibold">{{ thisMonthReferrals }}</p>
              </div>
            </div>
          </div>

          <div class="bg-white bg-opacity-10 rounded-lg p-4">
            <div class="flex items-center">
              <div class="flex-shrink-0">
                <svg class="h-6 w-6 text-blue-100" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z" />
                </svg>
              </div>
              <div class="ml-3">
                <p class="text-sm font-medium text-blue-100">Conversion Rate</p>
                <p class="text-lg font-semibold">{{ conversionRate }}%</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Earnings Chart -->
    <div 
      v-motion
      :initial="{ opacity: 0, y: 20 }"
      :enter="{ opacity: 1, y: 0, transition: { delay: 200 } }"
      class="bg-white rounded-lg shadow mb-8 overflow-hidden">
      <div class="px-6 py-5 border-b border-gray-200">
        <div class="flex items-center justify-between">
          <h3 class="text-lg font-medium leading-6 text-gray-900">Earnings Trends</h3>
          <div class="flex space-x-3">
            <button 
              @click="chartPeriod = 'weekly'"
              class="px-3 py-1.5 text-sm font-medium rounded-md transition-colors"
              :class="chartPeriod === 'weekly' ? 'bg-blue-100 text-blue-800' : 'text-gray-700 hover:bg-gray-100'">
              Weekly
            </button>
            <button 
              @click="chartPeriod = 'monthly'"
              class="px-3 py-1.5 text-sm font-medium rounded-md transition-colors"
              :class="chartPeriod === 'monthly' ? 'bg-blue-100 text-blue-800' : 'text-gray-700 hover:bg-gray-100'">
              Monthly
            </button>
            <button 
              @click="chartPeriod = 'yearly'"
              class="px-3 py-1.5 text-sm font-medium rounded-md transition-colors"
              :class="chartPeriod === 'yearly' ? 'bg-blue-100 text-blue-800' : 'text-gray-700 hover:bg-gray-100'">
              Yearly
            </button>
          </div>
        </div>
      </div>
      <div class="p-6">
        <!-- Placeholder for chart (in a real app, use a charting library like Chart.js) -->
        <div class="h-80 w-full flex items-end space-x-2">
          <div v-for="(item, index) in chartData" :key="index" class="flex-1 flex flex-col items-center">
            <div 
              class="w-full bg-blue-500 rounded-t"
              :style="{ height: `${item.value}%` }"
              v-motion
              :initial="{ opacity: 0, height: 0 }"
              :enter="{ opacity: 1, height: `${item.value}%`, transition: { delay: index * 100 } }">
            </div>
            <div class="text-xs text-gray-500 mt-2">{{ item.label }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Recent Payouts -->
    <div 
      v-motion
      :initial="{ opacity: 0, y: 20 }"
      :enter="{ opacity: 1, y: 0, transition: { delay: 300 } }"
      class="bg-white rounded-lg shadow mb-8">
      <div class="px-6 py-5 border-b border-gray-200">
        <h3 class="text-lg font-medium leading-6 text-gray-900">Recent Payouts</h3>
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Date
              </th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Amount
              </th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Payment Method
              </th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Status
              </th>
              <th scope="col" class="relative px-6 py-3">
                <span class="sr-only">Actions</span>
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="(payout, index) in recentPayouts" :key="payout.id"
              v-motion
              :initial="{ opacity: 0, y: 10 }"
              :enter="{ opacity: 1, y: 0, transition: { delay: index * 50 } }">
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {{ payout.date }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                ${{ payout.amount }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                <div class="flex items-center">
                  <div class="h-8 w-8 flex-shrink-0">
                    <img :src="payout.methodIcon" alt="" class="h-8 w-auto">
                  </div>
                  <div class="ml-3">
                    {{ payout.method }}
                    <div class="text-xs text-gray-400">{{ payout.methodDetails }}</div>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span 
                  class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full"
                  :class="{
                    'bg-green-100 text-green-800': payout.status === 'Completed',
                    'bg-yellow-100 text-yellow-800': payout.status === 'Processing',
                    'bg-gray-100 text-gray-800': payout.status === 'Scheduled'
                  }">
                  {{ payout.status }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                <button class="text-blue-600 hover:text-blue-900 transition-colors">Details</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="bg-gray-50 px-6 py-4 border-t border-gray-200">
        <div class="flex items-center justify-between">
          <span class="text-sm text-gray-700">
            Showing latest 5 payouts
          </span>
          <button class="text-sm font-medium text-blue-600 hover:text-blue-500 transition-colors">
            View All Payouts
          </button>
        </div>
      </div>
    </div>

    <!-- Payment Methods -->
    <div 
      v-motion
      :initial="{ opacity: 0, y: 20 }"
      :enter="{ opacity: 1, y: 0, transition: { delay: 400 } }"
      class="bg-white rounded-lg shadow">
      <div class="px-6 py-5 border-b border-gray-200 flex justify-between items-center">
        <h3 class="text-lg font-medium leading-6 text-gray-900">Payment Methods</h3>
        <button 
          type="button"
          class="inline-flex items-center px-3 py-1.5 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors">
          Add Method
        </button>
      </div>
      <div class="p-6">
        <div v-for="(method, index) in paymentMethods" :key="method.id" 
          class="flex items-center justify-between py-4 border-b border-gray-200 last:border-b-0"
          v-motion
          :initial="{ opacity: 0, x: -20 }"
          :enter="{ opacity: 1, x: 0, transition: { delay: index * 100 } }">
          <div class="flex items-center">
            <div class="h-10 w-10 flex-shrink-0">
              <img :src="method.icon" alt="" class="h-10 w-auto">
            </div>
            <div class="ml-4">
              <p class="text-sm font-medium text-gray-900">{{ method.name }}</p>
              <p class="text-sm text-gray-500">{{ method.details }}</p>
            </div>
          </div>
          <div class="flex items-center space-x-4">
            <span 
              v-if="method.isDefault" 
              class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
              Default
            </span>
            <button class="text-sm text-blue-600 hover:text-blue-900 transition-colors">Edit</button>
            <button class="text-sm text-red-600 hover:text-red-900 transition-colors">Remove</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useReferralStore } from '@/stores/referral'

// Referral store
const referralStore = useReferralStore()

// Stats data
const totalEarnings = ref(3240)
const availableForPayout = ref(840)
const thisMonthEarnings = ref(720)
const lastMonthEarnings = ref(580)
const thisMonthReferrals = ref(32)
const conversionRate = ref(64)

// Chart data
const chartPeriod = ref('monthly')

// Dynamic chart data based on selected period
const chartData = computed(() => {
  if (chartPeriod.value === 'weekly') {
    return [
      { label: 'Mon', value: 25 },
      { label: 'Tue', value: 40 },
      { label: 'Wed', value: 30 },
      { label: 'Thu', value: 70 },
      { label: 'Fri', value: 55 },
      { label: 'Sat', value: 60 },
      { label: 'Sun', value: 45 }
    ]
  } else if (chartPeriod.value === 'monthly') {
    return [
      { label: 'Jan', value: 30 },
      { label: 'Feb', value: 40 },
      { label: 'Mar', value: 45 },
      { label: 'Apr', value: 55 },
      { label: 'May', value: 60 },
      { label: 'Jun', value: 75 },
      { label: 'Jul', value: 70 },
      { label: 'Aug', value: 65 },
      { label: 'Sep', value: 80 },
      { label: 'Oct', value: 0 },
      { label: 'Nov', value: 0 },
      { label: 'Dec', value: 0 }
    ]
  } else {
    return [
      { label: '2021', value: 45 },
      { label: '2022', value: 60 },
      { label: '2023', value: 70 },
      { label: '2024', value: 85 },
      { label: '2025', value: 70 }
    ]
  }
})

// Recent payouts data
const recentPayouts = ref([
  {
    id: 1,
    date: 'Sept 1, 2025',
    amount: '720.00',
    method: 'PayPal',
    methodDetails: 'sarah@example.com',
    methodIcon: '/icons/paypal.svg',
    status: 'Completed'
  },
  {
    id: 2,
    date: 'Aug 1, 2025',
    amount: '580.00',
    method: 'Bank Transfer',
    methodDetails: '****4589',
    methodIcon: '/icons/bank.svg',
    status: 'Completed'
  },
  {
    id: 3,
    date: 'Jul 1, 2025',
    amount: '490.00',
    method: 'PayPal',
    methodDetails: 'sarah@example.com',
    methodIcon: '/icons/paypal.svg',
    status: 'Completed'
  },
  {
    id: 4,
    date: 'Jun 1, 2025',
    amount: '610.00',
    method: 'Bank Transfer',
    methodDetails: '****4589',
    methodIcon: '/icons/bank.svg',
    status: 'Completed'
  },
  {
    id: 5,
    date: 'May 1, 2025',
    amount: '540.00',
    method: 'PayPal',
    methodDetails: 'sarah@example.com',
    methodIcon: '/icons/paypal.svg',
    status: 'Completed'
  }
])

// Payment methods
const paymentMethods = ref([
  {
    id: 1,
    name: 'PayPal',
    details: 'sarah@example.com',
    icon: '/icons/paypal.svg',
    isDefault: true
  },
  {
    id: 2,
    name: 'Bank Account',
    details: 'CHASE BANK ****4589',
    icon: '/icons/bank.svg',
    isDefault: false
  }
])

// Fetch data on component mount
onMounted(async () => {
  // In a real app, fetch earnings data from API
  // await referralStore.fetchEarningsStats()
})
</script>