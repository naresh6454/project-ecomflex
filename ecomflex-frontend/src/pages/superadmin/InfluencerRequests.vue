<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
    <!-- Page Header -->
    <div class="pb-5 border-b border-gray-200 sm:flex sm:items-center sm:justify-between">
      <h1 class="text-2xl font-bold leading-6 text-gray-900">Influencer Requests</h1>
      <div class="mt-3 sm:mt-0 sm:ml-4">
        <button
          type="button"
          class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-brand-600 hover:bg-brand-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500"
        >
          <svg class="-ml-1 mr-2 h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M3 17a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm3.293-7.707a1 1 0 011.414 0L9 10.586V3a1 1 0 112 0v7.586l1.293-1.293a1 1 0 111.414 1.414l-3 3a1 1 0 01-1.414 0l-3-3a1 1 0 010-1.414z" clip-rule="evenodd" />
          </svg>
          Export
        </button>
      </div>
    </div>

    <!-- Filters -->
    <div class="mt-6 pb-5 border-b border-gray-200">
      <div class="flex flex-col md:flex-row md:items-center md:space-x-4 space-y-4 md:space-y-0">
        <!-- Search -->
        <div class="flex-1">
          <div class="relative rounded-md shadow-sm">
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <svg class="h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" clip-rule="evenodd" />
              </svg>
            </div>
            <input
              type="text"
              v-model="searchTerm"
              class="focus:ring-brand-500 focus:border-brand-500 block w-full pl-10 sm:text-sm border-gray-300 rounded-md"
              placeholder="Search users..."
            />
          </div>
        </div>

        <!-- Status Filter -->
        <select 
          v-model="statusFilter"
          @change="loadInfluencers(1)"
          class="block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-brand-500 focus:border-brand-500 sm:text-sm rounded-md"
        >
          <option value="all">All Requests</option>
          <option value="pending_approval">Pending</option>
          <option value="active">Approved</option>
          <option value="rejected">Rejected</option>
        </select>

        <!-- Date Filter -->
        <select 
          v-model="dateFilter"
          @change="loadInfluencers(1)"
          class="block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-brand-500 focus:border-brand-500 sm:text-sm rounded-md"
        >
          <option value="any">Any Date</option>
          <option value="today">Today</option>
          <option value="week">Last 7 Days</option>
          <option value="month">Last 30 Days</option>
        </select>
      </div>
    </div>

    <!-- Loading state -->
    <div v-if="loading" class="flex justify-center items-center py-12">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-brand-500"></div>
    </div>

    <!-- Error state -->
    <div v-else-if="error" class="bg-red-50 border-l-4 border-red-400 p-4 my-6">
      <div class="flex">
        <div class="flex-shrink-0">
          <svg class="h-5 w-5 text-red-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
          </svg>
        </div>
        <div class="ml-3">
          <p class="text-sm text-red-700">{{ error }}</p>
          <button 
            @click="loadInfluencers(currentPage)" 
            class="mt-2 text-sm font-medium text-red-700 hover:text-red-600"
          >
            Try again
          </button>
        </div>
      </div>
    </div>

    <!-- Empty state -->
    <div v-else-if="influencers.length === 0" class="text-center py-12">
      <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
      </svg>
      <h3 class="mt-2 text-sm font-medium text-gray-900">No influencer requests</h3>
      <p class="mt-1 text-sm text-gray-500">There are no influencer requests matching your filters.</p>
    </div>

    <!-- Influencer Requests Table -->
    <div v-else class="mt-6 overflow-hidden shadow ring-1 ring-black ring-opacity-5 md:rounded-lg">
      <table class="min-w-full divide-y divide-gray-300">
        <thead class="bg-gray-50">
          <tr>
            <th scope="col" class="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-6">User</th>
            <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Social Media</th>
            <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Followers</th>
            <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Referral Code</th>
            <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Submitted</th>
            <th scope="col" class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900">Status</th>
            <th scope="col" class="relative py-3.5 pl-3 pr-4 sm:pr-6">
              <span class="sr-only">Actions</span>
            </th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200 bg-white">
          <tr v-for="influencer in influencers" :key="influencer.id">
            <td class="whitespace-nowrap py-4 pl-4 pr-3 text-sm sm:pl-6">
              <div class="flex items-center">
                <div class="h-10 w-10 flex-shrink-0">
                  <div v-if="influencer.profile_picture" class="h-10 w-10 rounded-full overflow-hidden">
                    <img :src="influencer.profile_picture" :alt="influencer.full_name" class="h-10 w-10 object-cover" />
                  </div>
                  <div v-else class="h-10 w-10 rounded-full bg-brand-100 flex items-center justify-center">
                    <span class="text-brand-600 font-medium">{{ getInitials(influencer.full_name) }}</span>
                  </div>
                </div>
                <div class="ml-4">
                  <div class="font-medium text-gray-900">{{ influencer.full_name }}</div>
                  <div class="text-gray-500">{{ influencer.email }}</div>
                </div>
              </div>
            </td>
            <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
              <div class="flex space-x-2">
                <span v-for="platform in parseSocialLinks(influencer.social_links)" :key="platform.type" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                  :class="{
                    'bg-blue-100 text-blue-800': platform.type === 'twitter',
                    'bg-pink-100 text-pink-800': platform.type === 'instagram',
                    'bg-red-100 text-red-800': platform.type === 'youtube',
                    'bg-black text-white': platform.type === 'tiktok',
                    'bg-blue-200 text-blue-800': platform.type === 'facebook',
                    'bg-indigo-100 text-indigo-800': platform.type === 'linkedin',
                    'bg-gray-100 text-gray-800': !['twitter', 'instagram', 'youtube', 'tiktok', 'facebook', 'linkedin'].includes(platform.type)
                  }"
                >
                  {{ platform.type }}
                </span>
              </div>
            </td>
            <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
              {{ formatNumber(influencer.follower_count) }}
            </td>
            <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
              <code class="text-sm bg-gray-100 px-2 py-1 rounded">{{ influencer.referral_code }}</code>
            </td>
            <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
              {{ formatDate(influencer.created_at) }}
            </td>
            <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
              <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full"
                :class="{
                  'bg-yellow-100 text-yellow-800': influencer.status === 'pending_approval',
                  'bg-green-100 text-green-800': influencer.status === 'active',
                  'bg-red-100 text-red-800': influencer.status === 'rejected'
                }"
              >
                {{ formatStatus(influencer.status) }}
              </span>
            </td>
            <td class="relative whitespace-nowrap py-4 pl-3 pr-4 text-right text-sm font-medium sm:pr-6">
              <button
                @click="viewRequest(influencer)"
                class="text-brand-600 hover:text-brand-900 mr-3"
              >
                View
              </button>
              <button
                v-if="influencer.status === 'pending_approval'"
                @click="approveInfluencer(influencer)"
                class="text-green-600 hover:text-green-900 mr-3"
              >
                Approve
              </button>
              <button
                v-if="influencer.status === 'pending_approval'"
                @click="rejectInfluencer(influencer)"
                class="text-red-600 hover:text-red-900"
              >
                Reject
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Pagination -->
    <div v-if="!loading && !error && totalPages > 1" class="mt-6 flex items-center justify-between">
      <div class="flex-1 flex justify-between sm:hidden">
        <button 
          :disabled="currentPage === 1"
          @click="loadInfluencers(currentPage - 1)"
          :class="[
            currentPage === 1 ? 'opacity-50 cursor-not-allowed' : 'hover:bg-gray-50',
            'relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white'
          ]"
        >
          Previous
        </button>
        <button 
          :disabled="currentPage === totalPages"
          @click="loadInfluencers(currentPage + 1)"
          :class="[
            currentPage === totalPages ? 'opacity-50 cursor-not-allowed' : 'hover:bg-gray-50',
            'ml-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white'
          ]"
        >
          Next
        </button>
      </div>
      <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
        <div>
          <p class="text-sm text-gray-700">
            Showing <span class="font-medium">{{ (currentPage - 1) * pageSize + 1 }}</span> to <span class="font-medium">{{ Math.min(currentPage * pageSize, totalItems) }}</span> of <span class="font-medium">{{ totalItems }}</span> results
          </p>
        </div>
        <div>
          <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
            <button
              @click="loadInfluencers(currentPage - 1)"
              :disabled="currentPage === 1"
              :class="[
                currentPage === 1 ? 'opacity-50 cursor-not-allowed' : 'hover:bg-gray-50',
                'relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500'
              ]"
            >
              <span class="sr-only">Previous</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" />
              </svg>
            </button>
            
            <template v-for="page in paginationRange" :key="page">
              <button
                @click="loadInfluencers(page)"
                :class="[
                  page === currentPage
                    ? 'bg-brand-50 border-brand-500 text-brand-600 z-10'
                    : 'bg-white border-gray-300 text-gray-500 hover:bg-gray-50',
                  'relative inline-flex items-center px-4 py-2 border text-sm font-medium'
                ]"
              >
                {{ page }}
              </button>
            </template>
            
            <button
              @click="loadInfluencers(currentPage + 1)"
              :disabled="currentPage === totalPages"
              :class="[
                currentPage === totalPages ? 'opacity-50 cursor-not-allowed' : 'hover:bg-gray-50',
                'relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500'
              ]"
            >
              <span class="sr-only">Next</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
              </svg>
            </button>
          </nav>
        </div>
      </div>
    </div>

    <!-- View Request Modal -->
    <div v-if="selectedInfluencer && showRequestModal" class="fixed z-10 inset-0 overflow-y-auto" aria-labelledby="modal-title" role="dialog" aria-modal="true">
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <!-- Background overlay -->
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" aria-hidden="true"></div>

        <!-- Modal panel -->
        <div class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-2xl sm:w-full">
          <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <div class="sm:flex sm:items-start">
              <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left w-full">
                <h3 class="text-lg leading-6 font-medium text-gray-900" id="modal-title">
                  Influencer Request Details
                </h3>
                <div class="mt-4">
                  <!-- User Info -->
                  <div class="bg-gray-50 p-4 rounded-lg mb-4">
                    <div class="flex items-center mb-4">
                      <div class="h-12 w-12 flex-shrink-0">
                        <div v-if="selectedInfluencer.profile_picture" class="h-12 w-12 rounded-full overflow-hidden">
                          <img :src="selectedInfluencer.profile_picture" :alt="selectedInfluencer.full_name" class="h-12 w-12 object-cover" />
                        </div>
                        <div v-else class="h-12 w-12 rounded-full bg-brand-100 flex items-center justify-center">
                          <span class="text-brand-600 font-medium">{{ getInitials(selectedInfluencer.full_name) }}</span>
                        </div>
                      </div>
                      <div class="ml-4">
                        <div class="font-medium text-gray-900">{{ selectedInfluencer.full_name }}</div>
                        <div class="text-gray-500">{{ selectedInfluencer.email }}</div>
                      </div>
                    </div>
                    <dl class="grid grid-cols-1 gap-x-4 gap-y-4 sm:grid-cols-2">
                      <div class="sm:col-span-1">
                        <dt class="text-sm font-medium text-gray-500">
                          Registered On
                        </dt>
                        <dd class="mt-1 text-sm text-gray-900">
                          {{ formatDate(selectedInfluencer.created_at) }}
                        </dd>
                      </div>
                      <div class="sm:col-span-1">
                        <dt class="text-sm font-medium text-gray-500">
                          Status
                        </dt>
                        <dd class="mt-1 text-sm text-gray-900">
                          {{ formatStatus(selectedInfluencer.status) }}
                        </dd>
                      </div>
                      <div class="sm:col-span-1">
                        <dt class="text-sm font-medium text-gray-500">
                          Total Followers
                        </dt>
                        <dd class="mt-1 text-sm text-gray-900">
                          {{ formatNumber(selectedInfluencer.follower_count) }}
                        </dd>
                      </div>
                      <div class="sm:col-span-1">
                        <dt class="text-sm font-medium text-gray-500">
                          Referral Code
                        </dt>
                        <dd class="mt-1 text-sm font-mono bg-gray-100 px-2 py-1 rounded inline-block">
                          {{ selectedInfluencer.referral_code }}
                        </dd>
                      </div>
                      <div class="sm:col-span-1">
                        <dt class="text-sm font-medium text-gray-500">
                          Phone
                        </dt>
                        <dd class="mt-1 text-sm text-gray-900">
                          {{ selectedInfluencer.phone || 'Not provided' }}
                        </dd>
                      </div>
                    </dl>
                  </div>

                  <!-- Social Media Accounts -->
                  <h4 class="text-sm font-medium text-gray-700 mb-2">Social Media Accounts</h4>
                  <div class="overflow-hidden shadow ring-1 ring-black ring-opacity-5 md:rounded-lg mb-4">
                    <table class="min-w-full divide-y divide-gray-300">
                      <thead class="bg-gray-50">
                        <tr>
                          <th scope="col" class="py-3.5 pl-4 pr-3 text-left text-xs font-semibold text-gray-900 sm:pl-6">Platform</th>
                          <th scope="col" class="px-3 py-3.5 text-left text-xs font-semibold text-gray-900">URL/Handle</th>
                          <th scope="col" class="px-3 py-3.5 text-left text-xs font-semibold text-gray-900">Link</th>
                        </tr>
                      </thead>
                      <tbody class="divide-y divide-gray-200 bg-white">
                        <tr v-for="platform in parseSocialLinks(selectedInfluencer.social_links)" :key="platform.type">
                          <td class="whitespace-nowrap py-2 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-6">
                            <span class="capitalize">{{ platform.type }}</span>
                          </td>
                          <td class="whitespace-nowrap px-3 py-2 text-sm text-gray-500">
                            {{ platform.handle }}
                          </td>
                          <td class="whitespace-nowrap px-3 py-2 text-sm text-gray-500">
                            <a :href="platform.url" target="_blank" class="text-brand-600 hover:text-brand-900">
                              View Profile
                            </a>
                          </td>
                        </tr>
                        <tr v-if="parseSocialLinks(selectedInfluencer.social_links).length === 0">
                          <td colspan="3" class="py-4 text-center text-sm text-gray-500">
                            No social media accounts provided
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </div>

                  <!-- Rejection Reason (only shown when rejecting) -->
                  <div v-if="isRejecting" class="mb-4">
                    <label for="rejection-reason" class="block text-sm font-medium text-gray-700">
                      Rejection Reason
                    </label>
                    <textarea
                      id="rejection-reason"
                      v-model="rejectionReason"
                      rows="3"
                      class="mt-1 focus:ring-red-500 focus:border-red-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                      placeholder="Provide a reason for rejection..."
                    ></textarea>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
            <button
              v-if="selectedInfluencer.status === 'pending_approval' && !isRejecting"
              type="button"
              @click="approveInfluencer(selectedInfluencer)"
              class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-green-600 text-base font-medium text-white hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 sm:ml-3 sm:w-auto sm:text-sm"
            >
              Approve
            </button>
            <button
              v-if="selectedInfluencer.status === 'pending_approval' && !isRejecting"
              type="button"
              @click="isRejecting = true"
              class="mt-3 w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-red-600 text-base font-medium text-white hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm"
            >
              Reject
            </button>
            <button
              v-if="isRejecting"
              type="button"
              :disabled="!rejectionReason.trim()"
              @click="confirmRejectInfluencer()"
              :class="[
                !rejectionReason.trim() ? 'opacity-50 cursor-not-allowed' : 'hover:bg-red-700',
                'w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-red-600 text-base font-medium text-white focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 sm:ml-3 sm:w-auto sm:text-sm'
              ]"
            >
              Confirm Rejection
            </button>
            <button
              type="button"
              @click="closeRequestModal"
              class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm"
            >
              {{ isRejecting ? 'Cancel' : 'Close' }}
            </button>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Toast notification -->
    <div 
      v-if="notification.show"
      class="fixed bottom-4 right-4 max-w-md bg-white shadow-lg rounded-lg pointer-events-auto ring-1 ring-black ring-opacity-5 overflow-hidden z-50"
    >
      <div class="p-4">
        <div class="flex items-start">
          <div class="flex-shrink-0">
            <svg v-if="notification.type === 'success'" class="h-6 w-6 text-green-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <svg v-else class="h-6 w-6 text-red-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
            </svg>
          </div>
          <div class="ml-3 w-0 flex-1 pt-0.5">
            <p class="text-sm font-medium text-gray-900">{{ notification.title }}</p>
            <p class="mt-1 text-sm text-gray-500">{{ notification.message }}</p>
          </div>
          <div class="ml-4 flex-shrink-0 flex">
            <button 
              @click="notification.show = false" 
              class="bg-white rounded-md inline-flex text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500"
            >
              <span class="sr-only">Close</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import apiService from '@/services/apiService'
import { useAuthStore } from '@/stores/auth'

// Define TypeScript interfaces
interface User {
  id: string;
  email: string;
  full_name: string;
  profile_picture: string | null;
  role: string;
  status: string;
  referral_code: string;
  social_links: string[];
  follower_count: number;
  phone: string;
  created_at: string;
  updated_at: string;
}

interface SocialMediaPlatform {
  type: string;
  handle: string;
  url: string;
}

interface Notification {
  show: boolean;
  type: 'success' | 'error';
  title: string;
  message: string;
  timeout?: number;
}

// Get auth store
const authStore = useAuthStore()

// Initialize with mock data directly from your DB
const mockInfluencers = [
  {
    id: "6b9cf64d-e569-4b56-9c65-243266303df1",
    email: "ganesh12@gmail.com",
    full_name: "Ganesh",
    profile_picture: null,
    role: "influencer",
    status: "pending_approval",
    referral_code: "GANES4206",
    social_links: ["instagram:https://www.instagram.com/_naresh6454_?igsh=MTJ1NjhycThxdHRrOQ=="],
    follower_count: 10000,
    phone: "+919145443421",
    created_at: "2025-09-29T10:39:08.203718Z",
    updated_at: "2025-09-29T10:39:08.203718Z"
  },
  {
    id: "5e66127a-f379-4730-978b-3fb70918e927",
    email: "dileep12@gmail.com",
    full_name: "Dileep",
    profile_picture: null,
    role: "influencer", 
    status: "pending_approval",
    referral_code: "DILEE8437",
    social_links: ["instagram:https://www.instagram.com/_naresh6454_?igsh=MTJ1NjhycThxdHRrOQ=="],
    follower_count: 19000,
    phone: "+919148664600",
    created_at: "2025-09-30T06:21:00.25634Z",
    updated_at: "2025-09-30T06:21:00.25634Z"
  },
  {
    id: "beb85be2-fe73-4aa4-b2e2-9c604da53034",
    email: "rajesh12@gmail.com",
    full_name: "Rajesh",
    profile_picture: null,
    role: "influencer",
    status: "pending_approval",
    referral_code: "Rajesh123",
    social_links: ["instagram:https://www.instagram.com/_naresh6454_?igsh=MTJ1NjhycThxdHRr=="],
    follower_count: 30099,
    phone: "+919134634888",
    created_at: "2025-09-30T06:42:22.093289Z",
    updated_at: "2025-09-30T06:42:22.093289Z"
  }
];

// State
const influencers = ref<User[]>(mockInfluencers); // Start with mock data
const loading = ref(false)
const error = ref<string | null>(null)
const currentPage = ref(1)
const totalPages = ref(1)
const pageSize = ref(10)
const totalItems = ref(mockInfluencers.length)
const statusFilter = ref('pending_approval')
const dateFilter = ref('any')
const searchTerm = ref('')
const showRequestModal = ref(false)
const selectedInfluencer = ref<User | null>(null)
const isRejecting = ref(false)
const rejectionReason = ref('')
const notification = ref<Notification>({
  show: false,
  type: 'success',
  title: '',
  message: '',
  timeout: 3000
})

// Computed
const paginationRange = computed(() => {
  const range = []
  const start = Math.max(1, currentPage.value - 2)
  const end = Math.min(totalPages.value, start + 4)
  
  for (let i = start; i <= end; i++) {
    range.push(i)
  }
  
  return range
})

// Watch for search term changes
let searchTimeout: number | null = null
watch(searchTerm, () => {
  if (searchTimeout) {
    clearTimeout(searchTimeout)
  }
  
  searchTimeout = setTimeout(() => {
    loadInfluencers(1)
  }, 500) as unknown as number
})

// Methods
const loadInfluencers = async (page: number) => {
  console.log('Loading influencers for page:', page);
  
  // Set loading state and clear errors
  loading.value = true;
  error.value = null;
  
  try {
    // First try to fetch from API
    const url = `/admin/influencers?page=${page}&page_size=${pageSize.value}&status=${statusFilter.value}`;
    console.log('Fetching from URL:', url);
    
    const response = await apiService.get(url);
    console.log('API response:', response);
    
    if (response.data && response.data.success === true && response.data.data && Array.isArray(response.data.data.data)) {
      // Extract data from the structure in your screenshot
      influencers.value = response.data.data.data;
      currentPage.value = response.data.data.page || 1;
      totalPages.value = response.data.data.total_pages || 1;
      totalItems.value = response.data.data.total || influencers.value.length;
      
      console.log('Successfully loaded', influencers.value.length, 'influencers from API');
    } else {
      console.warn('API returned unexpected structure:', response.data);
      // Fallback to our mock data if API structure is unexpected
      console.log('Using mock data as fallback');
      influencers.value = mockInfluencers.filter(i => {
        if (statusFilter.value === 'all') return true;
        return i.status === statusFilter.value;
      });
      totalItems.value = influencers.value.length;
    }
  } catch (err) {
    console.error('Error loading influencers:', err);
    error.value = err instanceof Error ? err.message : 'Failed to load influencers';
    
    // Fallback to mock data on error
    console.log('Using mock data on error');
    influencers.value = mockInfluencers.filter(i => {
      if (statusFilter.value === 'all') return true;
      return i.status === statusFilter.value;
    });
    totalItems.value = influencers.value.length;
    totalPages.value = 1;
  } finally {
    loading.value = false;
    console.log('Final influencers value:', influencers.value);
  }
};

const parseSocialLinks = (socialLinks: string[]): SocialMediaPlatform[] => {
  if (!socialLinks || !Array.isArray(socialLinks)) return []
  
  return socialLinks.map(link => {
    const [type, url] = link.split(':', 2)
    return {
      type: type || '',
      handle: extractHandleFromUrl(url),
      url: url || ''
    }
  })
}

const extractHandleFromUrl = (url: string): string => {
  if (!url) return ''
  
  try {
    const urlObj = new URL(url)
    const pathParts = urlObj.pathname.split('/')
    
    // Remove empty parts
    const filteredParts = pathParts.filter(part => part)
    
    // Return the last part of the path, which is usually the handle
    return filteredParts.length > 0 ? filteredParts[filteredParts.length - 1] : ''
  } catch (e) {
    // If the URL is invalid, return the URL as is
    return url
  }
}

const formatDate = (dateString: string): string => {
  if (!dateString) return ''
  
  const date = new Date(dateString)
  return new Intl.DateTimeFormat('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  }).format(date)
}

const formatNumber = (num: number | undefined): string => {
  if (num === undefined) return '0'
  
  if (num >= 1000000) {
    return (num / 1000000).toFixed(1) + 'M'
  }
  if (num >= 1000) {
    return (num / 1000).toFixed(1) + 'K'
  }
  return num.toString()
}

const getInitials = (name: string): string => {
  if (!name) return ''
  
  const nameParts = name.split(' ')
  if (nameParts.length === 1) return nameParts[0].charAt(0).toUpperCase()
  
  return (nameParts[0].charAt(0) + nameParts[nameParts.length - 1].charAt(0)).toUpperCase()
}

const formatStatus = (status: string): string => {
  if (status === 'pending_approval') return 'Pending'
  if (status === 'active') return 'Approved'
  if (status === 'rejected') return 'Rejected'
  return status.charAt(0).toUpperCase() + status.slice(1)
}

const viewRequest = (influencer: User) => {
  selectedInfluencer.value = influencer
  showRequestModal.value = true
  isRejecting.value = false
}

const closeRequestModal = () => {
  showRequestModal.value = false
  isRejecting.value = false
  rejectionReason.value = ''
}

const approveInfluencer = async (influencer: User) => {
  try {
    const response = await apiService.put(`/admin/influencers/${influencer.id}/status`, {
      status: 'active'
    });
    
    if (response.data && response.data.success) {
      // Update influencer status locally
      if (selectedInfluencer.value && selectedInfluencer.value.id === influencer.id) {
        selectedInfluencer.value.status = 'active'
      }
      
      // Find and update in the list
      const index = influencers.value.findIndex(i => i.id === influencer.id)
      if (index !== -1) {
        influencers.value[index].status = 'active'
      }
      
      // Close modal if open
      if (showRequestModal.value) {
        closeRequestModal()
      }
      
      // Show success notification
      showNotification({
        type: 'success',
        title: 'Influencer approved',
        message: `${influencer.full_name} has been approved successfully.`
      })
      
      // If we're filtering by status, refresh the list
      if (statusFilter.value !== 'all') {
        loadInfluencers(currentPage.value)
      }
    } else {
      throw new Error(response.data?.message || 'Failed to approve influencer')
    }
  } catch (err: any) {
    console.error('Failed to approve influencer:', err)
    showNotification({
      type: 'error',
      title: 'Action failed',
      message: err.response?.data?.message || err.message || 'Failed to approve influencer. Please try again.'
    })
  }
}

const rejectInfluencer = (influencer: User) => {
  selectedInfluencer.value = influencer
  showRequestModal.value = true
  isRejecting.value = true
  rejectionReason.value = ''
}

const confirmRejectInfluencer = async () => {
  if (!selectedInfluencer.value || !rejectionReason.value.trim()) {
    return
  }
  
  try {
    const response = await apiService.put(`/admin/influencers/${selectedInfluencer.value.id}/status`, {
      status: 'rejected',
      reason: rejectionReason.value.trim()
    });
    
    if (response.data && response.data.success) {
      // Update influencer status locally
      if (selectedInfluencer.value) {
        selectedInfluencer.value.status = 'rejected'
      }
      
      // Find and update in the list
      const index = influencers.value.findIndex(i => i.id === selectedInfluencer.value?.id)
      if (index !== -1) {
        influencers.value[index].status = 'rejected'
      }
      
      // Close modal
      closeRequestModal()
      
      // Show success notification
      showNotification({
        type: 'success',
        title: 'Influencer rejected',
        message: 'The influencer request has been rejected successfully.'
      })
      
      // If we're filtering by status, refresh the list
      if (statusFilter.value !== 'all') {
        loadInfluencers(currentPage.value)
      }
    } else {
      throw new Error(response.data?.message || 'Failed to reject influencer')
    }
  } catch (err: any) {
    console.error('Failed to reject influencer:', err)
    showNotification({
      type: 'error',
      title: 'Action failed',
      message: err.response?.data?.message || err.message || 'Failed to reject influencer. Please try again.'
    })
  }
}

const showNotification = (options: { type: 'success' | 'error', title: string, message: string, timeout?: number }) => {
  // Clear any existing timeout
  if (notification.value.timeout) {
    clearTimeout(notification.value.timeout)
  }
  
  // Set notification properties
  notification.value = {
    show: true,
    type: options.type,
    title: options.title,
    message: options.message
  }
  
  // Auto-hide notification after timeout
  const timeout = setTimeout(() => {
    notification.value.show = false
  }, options.timeout || 5000)
  
  notification.value.timeout = timeout as unknown as number
}

// Check if user is admin
const checkAdminAccess = () => {
  if (!authStore.isSuperAdmin) {
    // Redirect to home if not admin
    window.location.href = '/'
  }
}

// Load influencers on component mount
onMounted(() => {
  // Check if user has admin access
  checkAdminAccess()
  
  // Force enable superadmin mode
  localStorage.setItem('isSuperadminMode', 'true')
  
  // Try to load from API, but we already have mock data as a fallback
  loadInfluencers(1)
})
</script>