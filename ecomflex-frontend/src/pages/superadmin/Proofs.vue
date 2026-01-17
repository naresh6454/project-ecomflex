<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
    <!-- Page Header -->
    <div class="pb-5 border-b border-gray-200 sm:flex sm:items-center sm:justify-between">
      <h1 class="text-2xl font-bold leading-6 text-gray-900">Proof Verification</h1>
      <div class="mt-3 sm:mt-0 sm:ml-4 flex space-x-3">
        <button
          type="button"
          class="inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500"
        >
          <svg class="-ml-1 mr-2 h-5 w-5 text-gray-500" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M3 17a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm3.293-7.707a1 1 0 011.414 0L9 10.586V3a1 1 0 112 0v7.586l1.293-1.293a1 1 0 111.414 1.414l-3 3a1 1 0 01-1.414 0l-3-3a1 1 0 010-1.414z" clip-rule="evenodd" />
          </svg>
          Export
        </button>
        <button
          type="button"
          class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
        >
          <svg class="-ml-1 mr-2 h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
          </svg>
          Approve Selected
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
              class="focus:ring-brand-500 focus:border-brand-500 block w-full pl-10 sm:text-sm border-gray-300 rounded-md"
              placeholder="Search by product or user..."
            />
          </div>
        </div>

        <!-- Status Filter -->
        <select class="block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-brand-500 focus:border-brand-500 sm:text-sm rounded-md">
          <option>All Proofs</option>
          <option selected>Pending</option>
          <option>Approved</option>
          <option>Rejected</option>
        </select>

        <!-- Date Filter -->
        <select class="block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-brand-500 focus:border-brand-500 sm:text-sm rounded-md">
          <option>Any Date</option>
          <option>Today</option>
          <option>Last 7 Days</option>
          <option>Last 30 Days</option>
        </select>
      </div>
    </div>

    <!-- Proofs Grid -->
    <div class="mt-6 grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
      <div v-for="proof in proofs" :key="proof.id" class="bg-white overflow-hidden shadow rounded-lg">
        <!-- Proof Card Header -->
        <div class="px-4 py-5 sm:px-6 flex justify-between items-start">
          <div>
            <h3 class="text-lg font-medium text-gray-900 truncate">
              {{ proof.product.name }}
            </h3>
            <p class="mt-1 text-sm text-gray-500">
              ASIN: {{ proof.product.asin }}
            </p>
          </div>
          <input
            type="checkbox"
            class="h-4 w-4 text-brand-600 focus:ring-brand-500 border-gray-300 rounded"
          />
        </div>

        <!-- Proof Image/Preview -->
        <div class="border-t border-gray-200">
          <div class="h-48 bg-gray-100 flex items-center justify-center">
            <img
              v-if="proof.proofUrl"
              :src="proof.proofUrl"
              alt="Proof"
              class="h-full w-full object-contain"
            />
            <div v-else class="text-center px-4">
              <svg class="mx-auto h-12 w-12 text-gray-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
              </svg>
              <p class="mt-2 text-sm text-gray-500">Preview unavailable</p>
            </div>
          </div>
        </div>

        <!-- Proof Details -->
        <div class="px-4 py-4 sm:px-6 border-t border-gray-200">
          <dl class="grid grid-cols-1 gap-x-4 gap-y-2 sm:grid-cols-2">
            <div class="sm:col-span-1">
              <dt class="text-sm font-medium text-gray-500">
                User
              </dt>
              <dd class="mt-1 text-sm text-gray-900">
                {{ proof.user.name }}
              </dd>
            </div>
            <div class="sm:col-span-1">
              <dt class="text-sm font-medium text-gray-500">
                Submitted
              </dt>
              <dd class="mt-1 text-sm text-gray-900">
                {{ formatDate(proof.submittedAt) }}
              </dd>
            </div>
            <div class="sm:col-span-1">
              <dt class="text-sm font-medium text-gray-500">
                Referral Code
              </dt>
              <dd class="mt-1 text-sm text-gray-900">
                {{ proof.referralCode || 'None' }}
              </dd>
            </div>
            <div class="sm:col-span-1">
              <dt class="text-sm font-medium text-gray-500">
                Status
              </dt>
              <dd class="mt-1 text-sm">
                <span
                  class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full"
                  :class="{
                    'bg-yellow-100 text-yellow-800': proof.status === 'pending',
                    'bg-green-100 text-green-800': proof.status === 'approved',
                    'bg-red-100 text-red-800': proof.status === 'rejected'
                  }"
                >
                  {{ proof.status.charAt(0).toUpperCase() + proof.status.slice(1) }}
                </span>
              </dd>
            </div>
          </dl>
          <div class="mt-4 flex space-x-3">
            <button
              type="button"
              class="inline-flex items-center px-3 py-1.5 border border-transparent text-xs font-medium rounded shadow-sm text-white bg-brand-600 hover:bg-brand-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500"
              @click="openProofModal(proof)"
            >
              View Full
            </button>
            <button
              type="button"
              class="inline-flex items-center px-3 py-1.5 border border-transparent text-xs font-medium rounded shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
            >
              Approve
            </button>
            <button
              type="button"
              class="inline-flex items-center px-3 py-1.5 border border-transparent text-xs font-medium rounded shadow-sm text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
            >
              Reject
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Pagination -->
    <div class="mt-6 flex items-center justify-between">
      <div class="flex-1 flex justify-between sm:hidden">
        <a href="#" class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">
          Previous
        </a>
        <a href="#" class="ml-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">
          Next
        </a>
      </div>
      <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
        <div>
          <p class="text-sm text-gray-700">
            Showing <span class="font-medium">1</span> to <span class="font-medium">10</span> of <span class="font-medium">28</span> results
          </p>
        </div>
        <div>
          <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
            <a href="#" class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50">
              <span class="sr-only">Previous</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" />
              </svg>
            </a>
            <a href="#" class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium text-gray-700 hover:bg-gray-50">
              1
            </a>
            <a href="#" class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-brand-50 text-sm font-medium text-brand-600 hover:bg-brand-100">
              2
            </a>
            <a href="#" class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium text-gray-700 hover:bg-gray-50">
              3
            </a>
            <a href="#" class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50">
              <span class="sr-only">Next</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
              </svg>
            </a>
          </nav>
        </div>
      </div>
    </div>

    <!-- Proof Detail Modal -->
    <div v-if="selectedProof && showProofModal" class="fixed z-10 inset-0 overflow-y-auto" aria-labelledby="modal-title" role="dialog" aria-modal="true">
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <!-- Background overlay -->
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" aria-hidden="true"></div>

        <!-- Modal panel -->
        <div class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-2xl sm:w-full">
          <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <div class="sm:flex sm:items-start">
              <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left w-full">
                <h3 class="text-lg leading-6 font-medium text-gray-900" id="modal-title">
                  Proof Details
                </h3>
                <div class="mt-4">
                  <!-- Product and User Info -->
                  <div class="bg-gray-50 p-4 rounded-lg mb-4">
                    <dl class="grid grid-cols-1 gap-x-4 gap-y-4 sm:grid-cols-2">
                      <div class="sm:col-span-1">
                        <dt class="text-sm font-medium text-gray-500">
                          Product
                        </dt>
                        <dd class="mt-1 text-sm text-gray-900">
                          {{ selectedProof.product.name }}
                        </dd>
                      </div>
                      <div class="sm:col-span-1">
                        <dt class="text-sm font-medium text-gray-500">
                          ASIN
                        </dt>
                        <dd class="mt-1 text-sm text-gray-900">
                          {{ selectedProof.product.asin }}
                        </dd>
                      </div>
                      <div class="sm:col-span-1">
                        <dt class="text-sm font-medium text-gray-500">
                          User
                        </dt>
                        <dd class="mt-1 text-sm text-gray-900">
                          {{ selectedProof.user.name }}
                        </dd>
                      </div>
                      <div class="sm:col-span-1">
                        <dt class="text-sm font-medium text-gray-500">
                          Email
                        </dt>
                        <dd class="mt-1 text-sm text-gray-900">
                          {{ selectedProof.user.email }}
                        </dd>
                      </div>
                      <div class="sm:col-span-1">
                        <dt class="text-sm font-medium text-gray-500">
                          Submitted
                        </dt>
                        <dd class="mt-1 text-sm text-gray-900">
                          {{ formatDate(selectedProof.submittedAt) }}
                        </dd>
                      </div>
                      <div class="sm:col-span-1">
                        <dt class="text-sm font-medium text-gray-500">
                          Referral Code
                        </dt>
                        <dd class="mt-1 text-sm text-gray-900">
                          {{ selectedProof.referralCode || 'None' }}
                        </dd>
                      </div>
                    </dl>
                  </div>

                  <!-- Proof Image -->
                  <div class="mb-4">
                    <h4 class="text-sm font-medium text-gray-500 mb-2">Proof Image</h4>
                    <div class="h-80 bg-gray-100 flex items-center justify-center rounded-lg overflow-hidden">
                      <img
                        v-if="selectedProof.proofUrl"
                        :src="selectedProof.proofUrl"
                        alt="Proof"
                        class="max-h-full max-w-full object-contain"
                      />
                      <div v-else class="text-center px-4">
                        <svg class="mx-auto h-12 w-12 text-gray-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
                        </svg>
                        <p class="mt-2 text-sm text-gray-500">Preview unavailable</p>
                      </div>
                    </div>
                  </div>

                  <!-- Rejection Reason (only shown when rejecting) -->
                  <div v-if="isRejecting" class="mb-4">
                    <label for="rejection-reason" class="block text-sm font-medium text-gray-700">
                      Rejection Reason
                    </label>
                    <textarea
                      id="rejection-reason"
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
              v-if="!isRejecting"
              type="button"
              class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-green-600 text-base font-medium text-white hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 sm:ml-3 sm:w-auto sm:text-sm"
            >
              Approve
            </button>
            <button
              v-if="!isRejecting"
              type="button"
              @click="isRejecting = true"
              class="mt-3 w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-red-600 text-base font-medium text-white hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm"
            >
              Reject
            </button>
            <button
              v-if="isRejecting"
              type="button"
              class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-red-600 text-base font-medium text-white hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 sm:ml-3 sm:w-auto sm:text-sm"
            >
              Confirm Rejection
            </button>
            <button
              type="button"
              @click="closeProofModal"
              class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm"
            >
              {{ isRejecting ? 'Cancel' : 'Close' }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// Define TypeScript interfaces
interface Product {
  id: string;
  name: string;
  asin: string;
}

interface User {
  id: string;
  name: string;
  email: string;
}

interface Proof {
  id: string;
  product: Product;
  user: User;
  proofUrl: string | null;
  status: 'pending' | 'approved' | 'rejected';
  submittedAt: string;
  referralCode: string | null;
}

// State
const showProofModal = ref(false)
const selectedProof = ref<Proof | null>(null)
const isRejecting = ref(false)

// Methods
const formatDate = (dateString: string): string => {
  const date = new Date(dateString)
  return new Intl.DateTimeFormat('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  }).format(date)
}

const openProofModal = (proof: Proof) => {
  selectedProof.value = proof
  showProofModal.value = true
  isRejecting.value = false
}

const closeProofModal = () => {
  showProofModal.value = false
  isRejecting.value = false
}

// Mock proofs data
const proofs: Proof[] = [
  {
    id: '1',
    product: {
      id: '1',
      name: 'Smart Watch X2',
      asin: 'B07FZ4P7LZ'
    },
    user: {
      id: 'u1',
      name: 'John Doe',
      email: 'john.doe@example.com'
    },
    proofUrl: 'https://picsum.photos/seed/proof1/800/600',
    status: 'pending',
    submittedAt: '2025-09-10T10:30:00Z',
    referralCode: 'INFLU123'
  },
  {
    id: '2',
    product: {
      id: '2',
      name: 'Premium Bluetooth Headphones',
      asin: 'B08C7KG5LP'
    },
    user: {
      id: 'u2',
      name: 'Jane Smith',
      email: 'jane.smith@example.com'
    },
    proofUrl: 'https://picsum.photos/seed/proof2/800/600',
    status: 'pending',
    submittedAt: '2025-09-11T14:45:00Z',
    referralCode: null
  },
  {
    id: '3',
    product: {
      id: '3',
      name: 'Organic Skin Care Set',
      asin: 'B09D3V6HQT'
    },
    user: {
      id: 'u3',
      name: 'Michael Johnson',
      email: 'michael.johnson@example.com'
    },
    proofUrl: 'https://picsum.photos/seed/proof3/800/600',
    status: 'pending',
    submittedAt: '2025-09-12T09:15:00Z',
    referralCode: 'INFLU456'
  },
  {
    id: '4',
    product: {
      id: '4',
      name: 'Portable Blender',
      asin: 'B083FD83N7'
    },
    user: {
      id: 'u4',
      name: 'Sarah Wilson',
      email: 'sarah.wilson@example.com'
    },
    proofUrl: 'https://picsum.photos/seed/proof4/800/600',
    status: 'pending',
    submittedAt: '2025-09-12T11:20:00Z',
    referralCode: null
  },
  {
    id: '5',
    product: {
      id: '5',
      name: 'Yoga Mat Pro',
      asin: 'B075R7LMVR'
    },
    user: {
      id: 'u5',
      name: 'David Brown',
      email: 'david.brown@example.com'
    },
    proofUrl: 'https://picsum.photos/seed/proof5/800/600',
    status: 'pending',
    submittedAt: '2025-09-13T08:45:00Z',
    referralCode: 'INFLU789'
  },
  {
    id: '6',
    product: {
      id: '1',
      name: 'Smart Watch X2',
      asin: 'B07FZ4P7LZ'
    },
    user: {
      id: 'u6',
      name: 'Emily Davis',
      email: 'emily.davis@example.com'
    },
    proofUrl: 'https://picsum.photos/seed/proof6/800/600',
    status: 'pending',
    submittedAt: '2025-09-13T10:30:00Z',
    referralCode: 'INFLU123'
  }
]
</script>