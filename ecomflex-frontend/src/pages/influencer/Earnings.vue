<template>
  <div class="earnings-page">
    <!-- Page Header -->
    <header class="mb-8">
      <h1 class="text-2xl md:text-3xl font-bold text-gray-900 mb-2">Earnings & Payouts</h1>
      <p class="text-gray-600">Track your referral earnings and manage payouts</p>
    </header>

    <!-- Stats Cards Row -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-8">
      <!-- Total Earnings Card -->
      <div class="bg-white rounded-xl shadow-lg p-5 border-l-4 border-accent">
        <div class="flex items-center justify-between">
          <div>
            <h3 class="text-gray-500 text-sm font-medium">Total Earnings</h3>
            <div class="mt-1 flex items-baseline">
              <p class="text-2xl font-semibold text-gray-900">
                <span 
                  class="animated-counter"
                  v-motion
                  :initial="{ opacity: 0, y: 20 }"
                  :enter="{ opacity: 1, y: 0, transition: { delay: 100, duration: 800 } }">
                  ${{ totalEarnings }}
                </span>
              </p>
            </div>
            <p class="text-xs text-green-600 mt-2" v-if="earningsChange > 0">
              <span class="flex items-center">
                <svg class="w-3 h-3 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 10l7-7m0 0l7 7m-7-7v18"></path>
                </svg>
                {{ earningsChange }}% from last month
              </span>
            </p>
          </div>
          <div class="rounded-full bg-accent/10 p-3">
            <svg class="w-6 h-6 text-accent" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- Available Balance Card -->
      <div class="bg-white rounded-xl shadow-lg p-5 border-l-4 border-green-500">
        <div class="flex items-center justify-between">
          <div>
            <h3 class="text-gray-500 text-sm font-medium">Available for Payout</h3>
            <div class="mt-1 flex items-baseline">
              <p class="text-2xl font-semibold text-gray-900">
                <span
                  class="animated-counter"
                  v-motion
                  :initial="{ opacity: 0, y: 20 }"
                  :enter="{ opacity: 1, y: 0, transition: { delay: 200, duration: 800 } }">
                  ${{ availableBalance }}
                </span>
              </p>
            </div>
            <button 
              class="text-xs bg-accent hover:bg-accent/80 text-white px-3 py-1 rounded-full mt-2 transition-all duration-300 focus:ring-2 focus:ring-accent/30"
              @click="showRequestPayoutModal = true"
              :disabled="availableBalance < 50">
              Request Payout
            </button>
          </div>
          <div class="rounded-full bg-green-500/10 p-3">
            <svg class="w-6 h-6 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 9V7a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2m2 4h10a2 2 0 002-2v-6a2 2 0 00-2-2H9a2 2 0 00-2 2v6a2 2 0 002 2zm7-5a2 2 0 11-4 0 2 2 0 014 0z"></path>
            </svg>
          </div>
        </div>
      </div>

      <!-- Pending Balance Card -->
      <div class="bg-white rounded-xl shadow-lg p-5 border-l-4 border-amber-500">
        <div class="flex items-center justify-between">
          <div>
            <h3 class="text-gray-500 text-sm font-medium">Pending Earnings</h3>
            <div class="mt-1 flex items-baseline">
              <p class="text-2xl font-semibold text-gray-900">
                <span
                  class="animated-counter"
                  v-motion
                  :initial="{ opacity: 0, y: 20 }"
                  :enter="{ opacity: 1, y: 0, transition: { delay: 300, duration: 800 } }">
                  ${{ pendingBalance }}
                </span>
              </p>
            </div>
            <p class="text-xs text-gray-500 mt-2">From {{ pendingBookings }} pending bookings</p>
          </div>
          <div class="rounded-full bg-amber-500/10 p-3">
            <svg class="w-6 h-6 text-amber-500" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
        </div>
      </div>
    </div>

    <!-- Earnings Chart -->
    <div class="bg-white rounded-xl shadow-lg p-5 mb-8">
      <div class="flex justify-between items-center mb-4">
        <h2 class="text-lg font-semibold text-gray-900">Earnings History</h2>
        <div class="flex space-x-2">
          <button 
            v-for="period in ['Week', 'Month', 'Year']" 
            :key="period"
            @click="selectedPeriod = period.toLowerCase()"
            :class="[
              'px-3 py-1 text-sm rounded-full transition-all duration-300',
              selectedPeriod === period.toLowerCase() 
                ? 'bg-accent text-white' 
                : 'bg-gray-100 text-gray-600 hover:bg-gray-200'
            ]">
            {{ period }}
          </button>
        </div>
      </div>
      
      <div class="h-64 relative">
        <canvas 
          ref="earningsChart"
          v-motion
          :initial="{ opacity: 0 }"
          :enter="{ opacity: 1, transition: { delay: 400, duration: 1000 } }">
        </canvas>
        <div 
          v-if="isLoading" 
          class="absolute inset-0 flex items-center justify-center bg-white bg-opacity-75">
          <div class="w-10 h-10 border-4 border-accent border-t-transparent rounded-full animate-spin"></div>
        </div>
      </div>
    </div>

    <!-- Earnings Table & Payout History Tabs -->
    <div class="bg-white rounded-xl shadow-lg overflow-hidden">
      <div class="border-b border-gray-200">
        <nav class="flex -mb-px">
          <button 
            v-for="tab in ['Earnings Details', 'Payout History']"
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

      <!-- Earnings Details Tab Content -->
      <div v-if="activeTab === 'Earnings Details'" class="p-4">
        <div class="flex justify-between items-center mb-4">
          <div class="relative w-64">
            <input 
              type="text" 
              placeholder="Search by product or user..." 
              class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-accent focus:border-accent"
              v-model="searchQuery">
            <div class="absolute left-3 top-2.5 text-gray-400">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
              </svg>
            </div>
          </div>
          <div>
            <select 
              v-model="selectedFilter"
              class="border border-gray-300 rounded-lg px-4 py-2 focus:ring-accent focus:border-accent">
              <option value="all">All Earnings</option>
              <option value="approved">Approved Only</option>
              <option value="pending">Pending Only</option>
            </select>
          </div>
        </div>

        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Date
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Product
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Buyer
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Amount
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Status
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr 
                v-for="(earning, index) in filteredEarnings" 
                :key="earning.id"
                v-motion
                :initial="{ opacity: 0, y: 10 }"
                :enter="{ opacity: 1, y: 0, transition: { delay: 100 * index, duration: 300 } }">
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ formatDate(earning.date) }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="h-10 w-10 flex-shrink-0">
                      <img class="h-10 w-10 rounded-md object-cover" :src="earning.productImage" alt="" />
                    </div>
                    <div class="ml-4">
                      <div class="text-sm font-medium text-gray-900">{{ earning.productName }}</div>
                      <div class="text-xs text-gray-500">ASIN: {{ earning.asin }}</div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="text-sm text-gray-900">{{ earning.buyerName }}</div>
                  <div class="text-xs text-gray-500">{{ earning.buyerEmail }}</div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 font-medium">
                  ${{ earning.amount.toFixed(2) }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span :class="[
                    'px-2 py-1 inline-flex text-xs leading-5 font-semibold rounded-full',
                    earning.status === 'approved' ? 'bg-green-100 text-green-800' : 
                    earning.status === 'pending' ? 'bg-yellow-100 text-yellow-800' : 
                    'bg-gray-100 text-gray-800'
                  ]">
                    {{ earning.status.charAt(0).toUpperCase() + earning.status.slice(1) }}
                  </span>
                </td>
              </tr>
              <tr v-if="filteredEarnings.length === 0">
                <td colspan="5" class="px-6 py-4 text-center text-sm text-gray-500">
                  No earnings found matching your criteria
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Pagination -->
        <div class="flex items-center justify-between px-4 py-3 bg-white border-t border-gray-200 sm:px-6">
          <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
            <div>
              <p class="text-sm text-gray-700">
                Showing
                <span class="font-medium">{{ paginationStart }}</span>
                to
                <span class="font-medium">{{ paginationEnd }}</span>
                of
                <span class="font-medium">{{ totalEarningsCount }}</span>
                results
              </p>
            </div>
            <div>
              <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
                <a 
                  href="#" 
                  @click.prevent="currentPage > 1 && (currentPage--)"
                  :class="[
                    'relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50',
                    currentPage === 1 ? 'cursor-not-allowed opacity-50' : ''
                  ]">
                  <span class="sr-only">Previous</span>
                  <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                    <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" />
                  </svg>
                </a>
                <a 
                  v-for="page in totalPages" 
                  :key="page" 
                  href="#" 
                  @click.prevent="currentPage = page"
                  :class="[
                    'relative inline-flex items-center px-4 py-2 border text-sm font-medium',
                    currentPage === page 
                      ? 'z-10 bg-accent border-accent text-white' 
                      : 'bg-white border-gray-300 text-gray-500 hover:bg-gray-50'
                  ]">
                  {{ page }}
                </a>
                <a 
                  href="#" 
                  @click.prevent="currentPage < totalPages && (currentPage++)"
                  :class="[
                    'relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50',
                    currentPage === totalPages ? 'cursor-not-allowed opacity-50' : ''
                  ]">
                  <span class="sr-only">Next</span>
                  <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                    <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
                  </svg>
                </a>
              </nav>
            </div>
          </div>
        </div>
      </div>

      <!-- Payout History Tab Content -->
      <div v-if="activeTab === 'Payout History'" class="p-4">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Date Requested
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Amount
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Method
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Status
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Date Processed
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr 
                v-for="(payout, index) in payouts" 
                :key="payout.id"
                v-motion
                :initial="{ opacity: 0, y: 10 }"
                :enter="{ opacity: 1, y: 0, transition: { delay: 100 * index, duration: 300 } }">
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ formatDate(payout.dateRequested) }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900 font-medium">
                  ${{ payout.amount.toFixed(2) }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ payout.method }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span :class="[
                    'px-2 py-1 inline-flex text-xs leading-5 font-semibold rounded-full',
                    payout.status === 'completed' ? 'bg-green-100 text-green-800' : 
                    payout.status === 'pending' ? 'bg-yellow-100 text-yellow-800' : 
                    payout.status === 'rejected' ? 'bg-red-100 text-red-800' :
                    'bg-gray-100 text-gray-800'
                  ]">
                    {{ payout.status.charAt(0).toUpperCase() + payout.status.slice(1) }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ payout.dateProcessed ? formatDate(payout.dateProcessed) : '—' }}
                </td>
              </tr>
              <tr v-if="payouts.length === 0">
                <td colspan="5" class="px-6 py-4 text-center text-sm text-gray-500">
                  No payout history found
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <!-- Request Payout Modal -->
    <div v-if="showRequestPayoutModal" class="fixed inset-0 z-50 overflow-y-auto" aria-labelledby="modal-title" role="dialog" aria-modal="true">
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <div 
          class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" 
          aria-hidden="true"
          v-motion
          :initial="{ opacity: 0 }"
          :enter="{ opacity: 1, transition: { duration: 300 } }"
          @click="showRequestPayoutModal = false"></div>

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
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
              <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                <h3 class="text-lg leading-6 font-medium text-gray-900" id="modal-title">
                  Request Payout
                </h3>
                <div class="mt-6">
                  <div class="mb-4">
                    <label for="amount" class="block text-sm font-medium text-gray-700">Amount</label>
                    <div class="mt-1 relative rounded-md shadow-sm">
                      <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                        <span class="text-gray-500 sm:text-sm">$</span>
                      </div>
                      <input 
                        type="number" 
                        name="amount" 
                        id="amount" 
                        class="focus:ring-accent focus:border-accent block w-full pl-7 pr-12 sm:text-sm border-gray-300 rounded-md"
                        placeholder="0.00"
                        v-model="payoutAmount"
                        :max="availableBalance"
                        :min="50">
                      <div class="absolute inset-y-0 right-0 flex items-center">
                        <label for="currency" class="sr-only">Currency</label>
                        <select 
                          id="currency" 
                          name="currency" 
                          class="focus:ring-accent focus:border-accent h-full py-0 pl-2 pr-7 border-transparent bg-transparent text-gray-500 sm:text-sm rounded-r-md">
                          <option>USD</option>
                        </select>
                      </div>
                    </div>
                    <p class="mt-1 text-xs text-gray-500">
                      Available: ${{ availableBalance }} (Min. $50.00)
                    </p>
                  </div>

                  <div class="mb-4">
                    <label for="paymentMethod" class="block text-sm font-medium text-gray-700">Payment Method</label>
                    <select 
                      id="paymentMethod" 
                      name="paymentMethod" 
                      class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-accent focus:border-accent sm:text-sm rounded-md"
                      v-model="payoutMethod">
                      <option value="paypal">PayPal</option>
                      <option value="bank">Bank Transfer</option>
                      <option value="venmo">Venmo</option>
                    </select>
                  </div>

                  <div v-if="payoutMethod === 'paypal'" class="mb-4">
                    <label for="paypalEmail" class="block text-sm font-medium text-gray-700">PayPal Email</label>
                    <input 
                      type="email" 
                      name="paypalEmail" 
                      id="paypalEmail" 
                      class="mt-1 focus:ring-accent focus:border-accent block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                      v-model="payoutDetails.email">
                  </div>

                  <div v-if="payoutMethod === 'bank'" class="space-y-4">
                    <div>
                      <label for="accountName" class="block text-sm font-medium text-gray-700">Account Name</label>
                      <input 
                        type="text" 
                        name="accountName" 
                        id="accountName" 
                        class="mt-1 focus:ring-accent focus:border-accent block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                        v-model="payoutDetails.accountName">
                    </div>
                    <div>
                      <label for="accountNumber" class="block text-sm font-medium text-gray-700">Account Number</label>
                      <input 
                        type="text" 
                        name="accountNumber" 
                        id="accountNumber" 
                        class="mt-1 focus:ring-accent focus:border-accent block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                        v-model="payoutDetails.accountNumber">
                    </div>
                    <div>
                      <label for="routingNumber" class="block text-sm font-medium text-gray-700">Routing Number</label>
                      <input 
                        type="text" 
                        name="routingNumber" 
                        id="routingNumber" 
                        class="mt-1 focus:ring-accent focus:border-accent block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                        v-model="payoutDetails.routingNumber">
                    </div>
                  </div>

                  <div v-if="payoutMethod === 'venmo'" class="mb-4">
                    <label for="venmoUsername" class="block text-sm font-medium text-gray-700">Venmo Username</label>
                    <div class="mt-1 flex rounded-md shadow-sm">
                      <span class="inline-flex items-center px-3 rounded-l-md border border-r-0 border-gray-300 bg-gray-50 text-gray-500 sm:text-sm">
                        @
                      </span>
                      <input 
                        type="text" 
                        name="venmoUsername" 
                        id="venmoUsername" 
                        class="focus:ring-accent focus:border-accent flex-1 block w-full rounded-none rounded-r-md sm:text-sm border-gray-300"
                        v-model="payoutDetails.username">
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
              :disabled="!isPayoutFormValid || isSubmitting"
              @click="submitPayoutRequest">
              <span v-if="isSubmitting" class="mr-2">
                <svg class="animate-spin h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
              </span>
              Request Payout
            </button>
            <button 
              type="button" 
              class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-accent/30 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm transition-all duration-300"
              @click="showRequestPayoutModal = false">
              Cancel
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue';
import { useReferralStore } from '@/stores/referral';
import { useNotificationStore } from '@/stores/notification';
import Chart from 'chart.js/auto';

// Import VueUse motion for animations
import { useMotion } from '@vueuse/motion';

// Stores
const referralStore = useReferralStore();
const notificationStore = useNotificationStore();

// State
const isLoading = ref(false);
const totalEarnings = ref(1248.75);
const availableBalance = ref(795.50);
const pendingBalance = ref(453.25);
const pendingBookings = ref(17);
const earningsChange = ref(12.5);
const selectedPeriod = ref('month');
const activeTab = ref('Earnings Details');
const searchQuery = ref('');
const selectedFilter = ref('all');
const currentPage = ref(1);
const itemsPerPage = 10;
const earningsChart = ref(null);
const chartInstance = ref(null);

// Payout modal state
const showRequestPayoutModal = ref(false);
const payoutAmount = ref(null);
const payoutMethod = ref('paypal');
const payoutDetails = ref({
  email: '',
  accountName: '',
  accountNumber: '',
  routingNumber: '',
  username: ''
});
const isSubmitting = ref(false);

// Mock data for earnings
const earnings = ref([
  {
    id: 1,
    date: new Date('2025-09-10'),
    productName: 'Wireless Earbuds Pro',
    productImage: 'https://placehold.co/200x200/FFD700/FFF?text=EP',
    asin: 'B07X123YZ9',
    buyerName: 'John Smith',
    buyerEmail: 'john.s@example.com',
    amount: 25.50,
    status: 'approved'
  },
  {
    id: 2,
    date: new Date('2025-09-08'),
    productName: 'Smart Water Bottle',
    productImage: 'https://placehold.co/200x200/FFD700/FFF?text=SWB',
    asin: 'B08Z456AB2',
    buyerName: 'Emma Johnson',
    buyerEmail: 'emma.j@example.com',
    amount: 18.75,
    status: 'approved'
  },
  {
    id: 3,
    date: new Date('2025-09-05'),
    productName: 'Portable Power Bank',
    productImage: 'https://placehold.co/200x200/FFD700/FFF?text=PB',
    asin: 'B09C789DE3',
    buyerName: 'Michael Brown',
    buyerEmail: 'michael.b@example.com',
    amount: 22.00,
    status: 'approved'
  },
  {
    id: 4,
    date: new Date('2025-09-01'),
    productName: 'Yoga Mat Premium',
    productImage: 'https://placehold.co/200x200/FFD700/FFF?text=YM',
    asin: 'B10F012GH4',
    buyerName: 'Sarah Wilson',
    buyerEmail: 'sarah.w@example.com',
    amount: 30.25,
    status: 'approved'
  },
  {
    id: 5,
    date: new Date('2025-08-28'),
    productName: 'Smart Home Speaker',
    productImage: 'https://placehold.co/200x200/FFD700/FFF?text=SHS',
    asin: 'B11I345JK5',
    buyerName: 'David Lee',
    buyerEmail: 'david.l@example.com',
    amount: 45.00,
    status: 'approved'
  },
  {
    id: 6,
    date: new Date('2025-08-25'),
    productName: 'Fitness Tracker Pro',
    productImage: 'https://placehold.co/200x200/FFD700/FFF?text=FT',
    asin: 'B12L678MN6',
    buyerName: 'Jennifer Garcia',
    buyerEmail: 'jennifer.g@example.com',
    amount: 35.50,
    status: 'pending'
  },
  {
    id: 7,
    date: new Date('2025-08-20'),
    productName: 'Coffee Maker Deluxe',
    productImage: 'https://placehold.co/200x200/FFD700/FFF?text=CM',
    asin: 'B13O901PQ7',
    buyerName: 'Robert Martinez',
    buyerEmail: 'robert.m@example.com',
    amount: 27.75,
    status: 'pending'
  },
  {
    id: 8,
    date: new Date('2025-08-15'),
    productName: 'Wireless Headphones',
    productImage: 'https://placehold.co/200x200/FFD700/FFF?text=WH',
    asin: 'B14R234ST8',
    buyerName: 'Lisa Thompson',
    buyerEmail: 'lisa.t@example.com',
    amount: 40.00,
    status: 'pending'
  },
  {
    id: 9,
    date: new Date('2025-08-10'),
    productName: 'Smart Watch Series 5',
    productImage: 'https://placehold.co/200x200/FFD700/FFF?text=SW',
    asin: 'B15U567VW9',
    buyerName: 'Kevin Anderson',
    buyerEmail: 'kevin.a@example.com',
    amount: 55.25,
    status: 'pending'
  },
  {
    id: 10,
    date: new Date('2025-08-05'),
    productName: 'Bluetooth Speaker',
    productImage: 'https://placehold.co/200x200/FFD700/FFF?text=BS',
    asin: 'B16X890YZ0',
    buyerName: 'Amanda White',
    buyerEmail: 'amanda.w@example.com',
    amount: 32.50,
    status: 'pending'
  },
]);

// Mock data for payouts
const payouts = ref([
  {
    id: 1,
    dateRequested: new Date('2025-08-01'),
    amount: 350.00,
    method: 'PayPal',
    status: 'completed',
    dateProcessed: new Date('2025-08-03')
  },
  {
    id: 2,
    dateRequested: new Date('2025-07-01'),
    amount: 275.50,
    method: 'Bank Transfer',
    status: 'completed',
    dateProcessed: new Date('2025-07-04')
  },
  {
    id: 3,
    dateRequested: new Date('2025-06-01'),
    amount: 420.75,
    method: 'PayPal',
    status: 'completed',
    dateProcessed: new Date('2025-06-02')
  },
  {
    id: 4,
    dateRequested: new Date('2025-09-15'),
    amount: 150.00,
    method: 'Venmo',
    status: 'pending',
    dateProcessed: null
  }
]);

// Computed properties
const filteredEarnings = computed(() => {
  let filtered = [...earnings.value];
  
  // Apply search filter
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase();
    filtered = filtered.filter(earning => 
      earning.productName.toLowerCase().includes(query) || 
      earning.buyerName.toLowerCase().includes(query) ||
      earning.buyerEmail.toLowerCase().includes(query) ||
      earning.asin.toLowerCase().includes(query)
    );
  }
  
  // Apply status filter
  if (selectedFilter.value !== 'all') {
    filtered = filtered.filter(earning => earning.status === selectedFilter.value);
  }
  
  // Apply pagination
  const start = (currentPage.value - 1) * itemsPerPage;
  const end = start + itemsPerPage;
  
  return filtered.slice(start, end);
});

const totalEarningsCount = computed(() => {
  let filtered = [...earnings.value];
  
  // Apply search filter
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase();
    filtered = filtered.filter(earning => 
      earning.productName.toLowerCase().includes(query) || 
      earning.buyerName.toLowerCase().includes(query) ||
      earning.buyerEmail.toLowerCase().includes(query) ||
      earning.asin.toLowerCase().includes(query)
    );
  }
  
  // Apply status filter
  if (selectedFilter.value !== 'all') {
    filtered = filtered.filter(earning => earning.status === selectedFilter.value);
  }
  
  return filtered.length;
});

const totalPages = computed(() => {
  return Math.ceil(totalEarningsCount.value / itemsPerPage);
});

const paginationStart = computed(() => {
  return totalEarningsCount.value === 0 ? 0 : (currentPage.value - 1) * itemsPerPage + 1;
});

const paginationEnd = computed(() => {
  return Math.min(currentPage.value * itemsPerPage, totalEarningsCount.value);
});

const isPayoutFormValid = computed(() => {
  if (!payoutAmount.value || payoutAmount.value < 50 || payoutAmount.value > availableBalance.value) {
    return false;
  }
  
  if (payoutMethod.value === 'paypal' && !payoutDetails.value.email) {
    return false;
  } else if (payoutMethod.value === 'bank' && 
    (!payoutDetails.value.accountName || 
     !payoutDetails.value.accountNumber || 
     !payoutDetails.value.routingNumber)) {
    return false;
  } else if (payoutMethod.value === 'venmo' && !payoutDetails.value.username) {
    return false;
  }
  
  return true;
});

// Methods
const formatDate = (date) => {
  return new Intl.DateTimeFormat('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  }).format(date);
};

const initChart = () => {
  isLoading.value = true;
  
  // Destroy existing chart if it exists
  if (chartInstance.value) {
    chartInstance.value.destroy();
  }
  
  setTimeout(() => {
    // Generate mock data based on selected period
    let labels = [];
    let data = [];
    
    if (selectedPeriod.value === 'week') {
      labels = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'];
      data = [45.25, 62.50, 30.00, 75.50, 48.75, 95.00, 58.25];
    } else if (selectedPeriod.value === 'month') {
      labels = ['Week 1', 'Week 2', 'Week 3', 'Week 4'];
      data = [210.50, 185.75, 325.00, 275.25];
    } else if (selectedPeriod.value === 'year') {
      labels = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
      data = [125.50, 145.25, 175.00, 195.75, 210.50, 230.25, 250.00, 275.75, 295.50, 0, 0, 0];
    }
    
    const ctx = earningsChart.value.getContext('2d');
    
    chartInstance.value = new Chart(ctx, {
      type: 'line',
      data: {
        labels: labels,
        datasets: [{
          label: 'Earnings',
          data: data,
          backgroundColor: 'rgba(255, 215, 0, 0.2)',
          borderColor: '#FFD700',
          borderWidth: 2,
          tension: 0.4,
          fill: true,
          pointBackgroundColor: '#FFD700',
          pointBorderColor: '#fff',
          pointBorderWidth: 2,
          pointRadius: 5,
          pointHoverRadius: 7
        }]
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: {
            display: false
          },
          tooltip: {
            backgroundColor: 'rgba(255, 255, 255, 0.8)',
            titleColor: '#333',
            bodyColor: '#666',
            borderColor: '#FFD700',
            borderWidth: 1,
            cornerRadius: 6,
            displayColors: false,
            callbacks: {
              label: function(context) {
                return `$${context.parsed.y.toFixed(2)}`;
              }
            }
          }
        },
        scales: {
          x: {
            grid: {
              display: false
            }
          },
          y: {
            beginAtZero: true,
            grid: {
              color: 'rgba(0, 0, 0, 0.05)'
            },
            ticks: {
              callback: function(value) {
                return '$' + value;
              }
            }
          }
        },
        animation: {
          duration: 1000,
          easing: 'easeOutQuart'
        }
      }
    });
    
    isLoading.value = false;
  }, 800);
};

const submitPayoutRequest = async () => {
  isSubmitting.value = true;
  
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1500));
    
    // Add the payout request to the list
    payouts.value.unshift({
      id: payouts.value.length + 1,
      dateRequested: new Date(),
      amount: parseFloat(payoutAmount.value),
      method: payoutMethod.value === 'paypal' ? 'PayPal' : 
              payoutMethod.value === 'bank' ? 'Bank Transfer' : 'Venmo',
      status: 'pending',
      dateProcessed: null
    });
    
    // Update available balance
    availableBalance.value -= parseFloat(payoutAmount.value);
    
    // Show success notification
    notificationStore.addNotification({
      type: 'success',
      message: `Payout request for $${payoutAmount.value} submitted successfully!`,
      duration: 5000
    });
    
    // Reset form and close modal
    payoutAmount.value = null;
    payoutMethod.value = 'paypal';
    payoutDetails.value = {
      email: '',
      accountName: '',
      accountNumber: '',
      routingNumber: '',
      username: ''
    };
    showRequestPayoutModal.value = false;
  } catch (error) {
    // Show error notification
    notificationStore.addNotification({
      type: 'error',
      message: 'Failed to submit payout request. Please try again.',
      duration: 5000
    });
  } finally {
    isSubmitting.value = false;
  }
};

// Watchers
watch(selectedPeriod, () => {
  initChart();
});

watch([searchQuery, selectedFilter], () => {
  currentPage.value = 1;
});

// Lifecycle hooks
onMounted(() => {
  initChart();
  
  // Fetch earnings data (would normally be from API)
  isLoading.value = true;
  
  // Simulate API call
  setTimeout(() => {
    // Data would be fetched from API here
    isLoading.value = false;
  }, 800);
});
</script>

<style scoped>
.earnings-page {
  @apply p-4 md:p-6 max-w-7xl mx-auto;
}

.animated-counter {
  @apply inline-block;
}
</style>