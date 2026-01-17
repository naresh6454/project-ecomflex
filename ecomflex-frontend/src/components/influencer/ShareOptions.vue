<template>
  <div class="bg-white rounded-lg shadow overflow-hidden">
    <div class="p-5">
      <h3 class="text-lg leading-6 font-medium text-gray-900">Share Your Referral Link</h3>
      <div class="mt-3 flex flex-col space-y-4">
        <!-- Copy Link -->
        <div class="flex flex-col sm:flex-row items-center space-y-2 sm:space-y-0 sm:space-x-2">
          <div class="flex-grow w-full">
            <div class="relative rounded-md shadow-sm">
              <input 
                type="text" 
                class="focus:ring-blue-500 focus:border-blue-500 block w-full pr-10 sm:text-sm border-gray-300 rounded-md"
                :value="referralLink" 
                readonly 
                ref="linkInput" />
              <div class="absolute inset-y-0 right-0 pr-3 flex items-center">
                <span 
                  class="text-gray-400 hover:text-gray-500 cursor-pointer" 
                  @click="copyToClipboard"
                  v-if="!copied">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path d="M8 3a1 1 0 011-1h2a1 1 0 110 2H9a1 1 0 01-1-1z" />
                    <path d="M6 3a2 2 0 00-2 2v11a2 2 0 002 2h8a2 2 0 002-2V5a2 2 0 00-2-2 3 3 0 01-3 3H9a3 3 0 01-3-3z" />
                  </svg>
                </span>
                <span 
                  class="text-green-500" 
                  v-else>
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                  </svg>
                </span>
              </div>
            </div>
          </div>
          <button 
            type="button"
            class="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors"
            @click="copyToClipboard">
            {{ copied ? 'Copied!' : 'Copy Link' }}
          </button>
        </div>

        <!-- Share to Social Media -->
        <div>
          <p class="text-sm text-gray-500 mb-2">Share to social media</p>
          <div 
            class="flex flex-wrap gap-2"
            v-motion
            :initial="{ opacity: 0 }"
            :enter="{ opacity: 1, transition: { delay: 200 } }">
            <!-- WhatsApp -->
            <button 
              type="button"
              class="inline-flex items-center px-3.5 py-2 border border-transparent text-sm leading-4 font-medium rounded-md shadow-sm text-white bg-green-600 hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500 transition-colors"
              @click="shareToWhatsApp">
              <svg class="h-4 w-4 mr-1.5" fill="currentColor" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                <path d="M.057 24l1.687-6.163c-1.041-1.804-1.588-3.849-1.587-5.946.003-6.556 5.338-11.891 11.893-11.891 3.181.001 6.167 1.24 8.413 3.488 2.245 2.248 3.481 5.236 3.48 8.414-.003 6.557-5.338 11.892-11.893 11.892-1.99-.001-3.951-.5-5.688-1.448l-6.305 1.654zm6.597-3.807c1.676.995 3.276 1.591 5.392 1.592 5.448 0 9.886-4.434 9.889-9.885.002-5.462-4.415-9.89-9.881-9.892-5.452 0-9.887 4.434-9.889 9.884-.001 2.225.651 3.891 1.746 5.634l-.999 3.648 3.742-.981zm11.387-5.464c-.074-.124-.272-.198-.57-.347-.297-.149-1.758-.868-2.031-.967-.272-.099-.47-.149-.669.149-.198.297-.768.967-.941 1.165-.173.198-.347.223-.644.074-.297-.149-1.255-.462-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.297-.347.446-.521.151-.172.2-.296.3-.495.099-.198.05-.372-.025-.521-.075-.148-.669-1.611-.916-2.206-.242-.579-.487-.501-.669-.51l-.57-.01c-.198 0-.52.074-.792.372s-1.04 1.016-1.04 2.479 1.065 2.876 1.213 3.074c.149.198 2.095 3.2 5.076 4.487.709.306 1.263.489 1.694.626.712.226 1.36.194 1.872.118.571-.085 1.758-.719 2.006-1.413.248-.695.248-1.29.173-1.414z"/>
              </svg>
              WhatsApp
            </button>

            <!-- Twitter -->
            <button 
              type="button"
              class="inline-flex items-center px-3.5 py-2 border border-transparent text-sm leading-4 font-medium rounded-md shadow-sm text-white bg-blue-400 hover:bg-blue-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-400 transition-colors"
              @click="shareToTwitter">
              <svg class="h-4 w-4 mr-1.5" fill="currentColor" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                <path d="M23.953 4.57a10 10 0 01-2.825.775 4.958 4.958 0 002.163-2.723c-.951.555-2.005.959-3.127 1.184a4.92 4.92 0 00-8.384 4.482C7.69 8.095 4.067 6.13 1.64 3.162a4.822 4.822 0 00-.666 2.475c0 1.71.87 3.213 2.188 4.096a4.904 4.904 0 01-2.228-.616v.06a4.923 4.923 0 003.946 4.827 4.996 4.996 0 01-2.212.085 4.936 4.936 0 004.604 3.417 9.867 9.867 0 01-6.102 2.105c-.39 0-.779-.023-1.17-.067a13.995 13.995 0 007.557 2.209c9.053 0 13.998-7.496 13.998-13.985 0-.21 0-.42-.015-.63A9.935 9.935 0 0024 4.59z"/>
              </svg>
              Twitter
            </button>

            <!-- Facebook -->
            <button 
              type="button"
              class="inline-flex items-center px-3.5 py-2 border border-transparent text-sm leading-4 font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-600 transition-colors"
              @click="shareToFacebook">
              <svg class="h-4 w-4 mr-1.5" fill="currentColor" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                <path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/>
              </svg>
              Facebook
            </button>

            <!-- LinkedIn -->
            <button 
              type="button"
              class="inline-flex items-center px-3.5 py-2 border border-transparent text-sm leading-4 font-medium rounded-md shadow-sm text-white bg-blue-700 hover:bg-blue-800 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-700 transition-colors"
              @click="shareToLinkedIn">
              <svg class="h-4 w-4 mr-1.5" fill="currentColor" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                <path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"/>
              </svg>
              LinkedIn
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const props = defineProps({
  referralLink: {
    type: String,
    required: true
  }
})

const copied = ref(false)
const linkInput = ref<HTMLInputElement | null>(null)

// Copy to clipboard function
const copyToClipboard = () => {
  if (linkInput.value) {
    linkInput.value.select()
    document.execCommand('copy')
    copied.value = true
    
    // Reset copied state after 2 seconds
    setTimeout(() => {
      copied.value = false
    }, 2000)
  }
}

// Social media sharing functions
const shareToWhatsApp = () => {
  const message = encodeURIComponent(`Check out this product! Get 100% cashback using my referral link: ${props.referralLink}`)
  window.open(`https://wa.me/?text=${message}`, '_blank')
}

const shareToTwitter = () => {
  const message = encodeURIComponent(`Get 100% cashback on amazing products using my referral link: ${props.referralLink} #Ecomflex #Cashback`)
  window.open(`https://twitter.com/intent/tweet?text=${message}`, '_blank')
}

const shareToFacebook = () => {
  window.open(`https://www.facebook.com/sharer/sharer.php?u=${encodeURIComponent(props.referralLink)}`, '_blank')
}

const shareToLinkedIn = () => {
  const title = encodeURIComponent('Get 100% Cashback on Products')
  const summary = encodeURIComponent('Check out this amazing cashback offer! Use my referral link to get 100% cashback on your purchase.')
  window.open(`https://www.linkedin.com/shareArticle?mini=true&url=${encodeURIComponent(props.referralLink)}&title=${title}&summary=${summary}`, '_blank')
}
</script>