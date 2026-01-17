import api from './api'
import { Booking } from '@/stores/booking'

interface CreateBookingRequest {
  productId: string
  referralCode?: string
}

class BookingService {
  /**
   * Get all bookings for the current user
   */
  getUserBookings() {
    return api.get<Booking[]>('/bookings')
  }

  /**
   * Get booking by ID
   */
  getBookingById(bookingId: string) {
    return api.get<Booking>(`/bookings/${bookingId}`)
  }

  /**
   * Create a new booking
   */
  createBooking(bookingData: CreateBookingRequest) {
    return api.post<Booking>('/bookings', bookingData)
  }

  /**
   * Upload proof for a booking
   */
  uploadProof(bookingId: string, formData: FormData) {
    return api.post<Booking>(`/bookings/${bookingId}/proof`, formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
  }

  /**
   * Cancel a booking
   */
  cancelBooking(bookingId: string) {
    return api.delete(`/bookings/${bookingId}`)
  }

  /**
   * For admin: Get all bookings
   */
  getAllBookings(filters?: { 
    status?: string, 
    userId?: string, 
    page?: number, 
    limit?: number 
  }) {
    return api.get<{
      bookings: Booking[],
      total: number,
      page: number,
      limit: number
    }>('/admin/bookings', { params: filters })
  }

  /**
   * For admin: Approve a booking
   */
  approveBooking(bookingId: string) {
    return api.patch<Booking>(`/admin/bookings/${bookingId}/approve`, {})
  }

  /**
   * For admin: Reject a booking
   */
  rejectBooking(bookingId: string, reason: string) {
    return api.patch<Booking>(`/admin/bookings/${bookingId}/reject`, { reason })
  }

  /**
   * For admin: Mark cashback as paid
   */
  markCashbackPaid(bookingId: string) {
    return api.patch<Booking>(`/admin/bookings/${bookingId}/cashback`, { status: 'paid' })
  }
}

export default new BookingService()