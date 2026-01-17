import { defineStore } from 'pinia'
import bookingService from '@/services/booking.service'

export interface Booking {
  id: string
  productId: string
  userId: string
  status: 'initiated' | 'pending' | 'approved' | 'rejected'
  proofUrl?: string
  referralCode?: string
  cashbackStatus: 'not_paid' | 'paid'
  createdAt: string
  updatedAt: string
}

export interface BookingState {
  bookings: Booking[]
  currentBooking: Booking | null
  loading: boolean
  error: string | null
}

export const useBookingStore = defineStore('booking', {
  state: (): BookingState => ({
    bookings: [],
    currentBooking: null,
    loading: false,
    error: null
  }),
  
  getters: {
    // Get bookings filtered by status
    getBookingsByStatus: (state) => (status: Booking['status']) => {
      return state.bookings.filter(booking => booking.status === status)
    },
    
    // Get pending bookings (uploaded proof but not verified)
    pendingBookings: (state) => {
      return state.bookings.filter(booking => booking.status === 'pending')
    },
    
    // Get approved bookings
    approvedBookings: (state) => {
      return state.bookings.filter(booking => booking.status === 'approved')
    },
    
    // Get rejected bookings
    rejectedBookings: (state) => {
      return state.bookings.filter(booking => booking.status === 'rejected')
    },
    
    // Get bookings by cashback status
    getBookingsByCashbackStatus: (state) => (status: Booking['cashbackStatus']) => {
      return state.bookings.filter(booking => booking.cashbackStatus === status)
    }
  },
  
  actions: {
    // Get all bookings for the current user
    async fetchUserBookings() {
      this.loading = true
      this.error = null
      
      try {
        const response = await bookingService.getUserBookings()
        this.bookings = response.data
      } catch (error: any) {
        this.error = error.message || 'Failed to fetch bookings'
        console.error('Error fetching bookings:', error)
      } finally {
        this.loading = false
      }
    },
    
    // Get a specific booking by ID
    async fetchBookingById(bookingId: string) {
      this.loading = true
      this.error = null
      
      try {
        const response = await bookingService.getBookingById(bookingId)
        this.currentBooking = response.data
        return response.data
      } catch (error: any) {
        this.error = error.message || `Failed to fetch booking ${bookingId}`
        console.error(`Error fetching booking ${bookingId}:`, error)
        return null
      } finally {
        this.loading = false
      }
    },
    
    // Create a new booking
    async createBooking(productId: string, referralCode?: string) {
      this.loading = true
      this.error = null
      
      try {
        const response = await bookingService.createBooking({
          productId,
          referralCode
        })
        
        // Add the new booking to the bookings array
        this.bookings.push(response.data)
        this.currentBooking = response.data
        return response.data
      } catch (error: any) {
        this.error = error.message || 'Failed to create booking'
        console.error('Error creating booking:', error)
        return null
      } finally {
        this.loading = false
      }
    },
    
    // Upload proof for a booking
    async uploadProof(bookingId: string, proofFile: File) {
      this.loading = true
      this.error = null
      
      try {
        const formData = new FormData()
        formData.append('proof', proofFile)
        
        const response = await bookingService.uploadProof(bookingId, formData)
        
        // Update the booking in the bookings array
        const index = this.bookings.findIndex(b => b.id === bookingId)
        if (index !== -1) {
          this.bookings[index] = response.data
        }
        
        if (this.currentBooking?.id === bookingId) {
          this.currentBooking = response.data
        }
        
        return response.data
      } catch (error: any) {
        this.error = error.message || 'Failed to upload proof'
        console.error('Error uploading proof:', error)
        return null
      } finally {
        this.loading = false
      }
    },
    
    // Cancel a booking
    async cancelBooking(bookingId: string) {
      this.loading = true
      this.error = null
      
      try {
        await bookingService.cancelBooking(bookingId)
        
        // Remove the booking from the bookings array
        this.bookings = this.bookings.filter(b => b.id !== bookingId)
        
        if (this.currentBooking?.id === bookingId) {
          this.currentBooking = null
        }
        
        return true
      } catch (error: any) {
        this.error = error.message || 'Failed to cancel booking'
        console.error('Error canceling booking:', error)
        return false
      } finally {
        this.loading = false
      }
    },
    
    // Clear current booking
    clearCurrentBooking() {
      this.currentBooking = null
    }
  }
})