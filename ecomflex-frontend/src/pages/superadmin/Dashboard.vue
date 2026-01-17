<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
    <!-- Page Header -->
    <div class="pb-5 border-b border-gray-200 sm:flex sm:items-center sm:justify-between">
      <h1 class="text-2xl font-bold leading-6 text-gray-900">Dashboard</h1>
      <div class="mt-3 sm:mt-0 sm:ml-4">
        <button
          type="button"
          class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-brand-600 hover:bg-brand-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500"
        >
          <svg class="-ml-1 mr-2 h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z" clip-rule="evenodd" />
          </svg>
          Generate Report
        </button>
      </div>
    </div>

    <!-- Stats Cards -->
    <div class="mt-6 grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-4">
      <!-- Total Users -->
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="px-4 py-5 sm:p-6">
          <div class="flex items-center">
            <div class="flex-shrink-0 bg-brand-100 rounded-md p-3">
              <svg class="h-6 w-6 text-brand-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
              </svg>
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">
                  Total Users
                </dt>
                <dd>
                  <div class="text-lg font-medium text-gray-900">
                    2,547
                  </div>
                </dd>
              </dl>
            </div>
          </div>
        </div>
        <div class="bg-gray-50 px-4 py-4 sm:px-6">
          <div class="text-sm">
            <a href="#" class="font-medium text-brand-600 hover:text-brand-500">
              View all<span class="sr-only"> stats</span>
            </a>
          </div>
        </div>
      </div>

      <!-- Influencers -->
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="px-4 py-5 sm:p-6">
          <div class="flex items-center">
            <div class="flex-shrink-0 bg-yellow-100 rounded-md p-3">
              <svg class="h-6 w-6 text-yellow-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
              </svg>
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">
                  Active Influencers
                </dt>
                <dd>
                  <div class="text-lg font-medium text-gray-900">
                    187
                  </div>
                </dd>
              </dl>
            </div>
          </div>
        </div>
        <div class="bg-gray-50 px-4 py-4 sm:px-6">
          <div class="text-sm">
            <a href="#" class="font-medium text-brand-600 hover:text-brand-500">
              View all<span class="sr-only"> stats</span>
            </a>
          </div>
        </div>
      </div>

      <!-- Products -->
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="px-4 py-5 sm:p-6">
          <div class="flex items-center">
            <div class="flex-shrink-0 bg-green-100 rounded-md p-3">
              <svg class="h-6 w-6 text-green-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10" />
              </svg>
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">
                  Active Products
                </dt>
                <dd>
                  <div class="text-lg font-medium text-gray-900">
                    42
                  </div>
                </dd>
              </dl>
            </div>
          </div>
        </div>
        <div class="bg-gray-50 px-4 py-4 sm:px-6">
          <div class="text-sm">
            <a href="#" class="font-medium text-brand-600 hover:text-brand-500">
              View all<span class="sr-only"> stats</span>
            </a>
          </div>
        </div>
      </div>

      <!-- Pending Proofs -->
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="px-4 py-5 sm:p-6">
          <div class="flex items-center">
            <div class="flex-shrink-0 bg-red-100 rounded-md p-3">
              <svg class="h-6 w-6 text-red-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
              </svg>
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">
                  Pending Proofs
                </dt>
                <dd>
                  <div class="text-lg font-medium text-gray-900">
                    28
                  </div>
                </dd>
              </dl>
            </div>
          </div>
        </div>
        <div class="bg-gray-50 px-4 py-4 sm:px-6">
          <div class="text-sm">
            <a href="#" class="font-medium text-brand-600 hover:text-brand-500">
              View all<span class="sr-only"> stats</span>
            </a>
          </div>
        </div>
      </div>
    </div>

    <!-- Charts Section -->
    <div class="mt-8 grid grid-cols-1 gap-8 lg:grid-cols-2">
      <!-- Bookings Over Time -->
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="px-4 py-5 sm:px-6 border-b border-gray-200">
          <h3 class="text-lg leading-6 font-medium text-gray-900">
            Bookings Over Time
          </h3>
        </div>
        <div class="px-4 py-5 sm:p-6">
          <div class="h-64 bg-gray-100 rounded-lg flex items-center justify-center">
            <p class="text-gray-500">Chart will be displayed here</p>
          </div>
        </div>
      </div>

      <!-- Top Influencers -->
      <div class="bg-white overflow-hidden shadow rounded-lg">
        <div class="px-4 py-5 sm:px-6 border-b border-gray-200">
          <h3 class="text-lg leading-6 font-medium text-gray-900">
            Top Influencers
          </h3>
        </div>
        <div class="px-4 py-5 sm:p-6">
          <div class="h-64 bg-gray-100 rounded-lg flex items-center justify-center">
            <p class="text-gray-500">Chart will be displayed here</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Recent Activity -->
    <div class="mt-8">
      <div class="bg-white shadow rounded-lg">
        <div class="px-4 py-5 sm:px-6 border-b border-gray-200">
          <h3 class="text-lg leading-6 font-medium text-gray-900">
            Recent Activity
          </h3>
        </div>
        <div class="bg-gray-50 px-4 py-5 sm:p-6">
          <ul class="divide-y divide-gray-200">
            <li class="py-4">
              <div class="flex space-x-3">
                <div class="flex-shrink-0">
                  <div class="h-10 w-10 rounded-full bg-brand-100 flex items-center justify-center">
                    <span class="text-brand-600 font-medium">JD</span>
                  </div>
                </div>
                <div class="min-w-0 flex-1">
                  <p class="text-sm font-medium text-gray-900">
                    John Doe uploaded a proof
                  </p>
                  <p class="text-sm text-gray-500">
                    <a href="#" class="hover:underline">Product #12345</a>
                  </p>
                  <p class="text-sm text-gray-500">
                    2 hours ago
                  </p>
                </div>
              </div>
            </li>
            <li class="py-4">
              <div class="flex space-x-3">
                <div class="flex-shrink-0">
                  <div class="h-10 w-10 rounded-full bg-brand-100 flex items-center justify-center">
                    <span class="text-brand-600 font-medium">AS</span>
                  </div>
                </div>
                <div class="min-w-0 flex-1">
                  <p class="text-sm font-medium text-gray-900">
                    Alice Smith applied to be an influencer
                  </p>
                  <p class="text-sm text-gray-500">
                    <a href="#" class="hover:underline">View application</a>
                  </p>
                  <p class="text-sm text-gray-500">
                    5 hours ago
                  </p>
                </div>
              </div>
            </li>
            <li class="py-4">
              <div class="flex space-x-3">
                <div class="flex-shrink-0">
                  <div class="h-10 w-10 rounded-full bg-brand-100 flex items-center justify-center">
                    <span class="text-brand-600 font-medium">BJ</span>
                  </div>
                </div>
                <div class="min-w-0 flex-1">
                  <p class="text-sm font-medium text-gray-900">
                    Bob Johnson's proof was approved
                  </p>
                  <p class="text-sm text-gray-500">
                    <a href="#" class="hover:underline">Product #67890</a>
                  </p>
                  <p class="text-sm text-gray-500">
                    8 hours ago
                  </p>
                </div>
              </div>
            </li>
          </ul>
        </div>
        <div class="px-4 py-4 sm:px-6 border-t border-gray-200">
          <div class="text-sm">
            <a href="#" class="font-medium text-brand-600 hover:text-brand-500">
              View all activity<span class="sr-only"> stats</span>
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Admin Dashboard component
</script>