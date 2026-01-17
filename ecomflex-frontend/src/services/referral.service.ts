import api from './api'

class ReferralService {
  /**
   * Get influencer's referral code
   */
  async getReferralCode() {
    try {
      const response = await api.get('/influencer/referral-code')
      return response.data
    } catch (error) {
      console.error('Error fetching referral code:', error)
      throw error
    }
  }

  /**
   * Get all referrals for the logged-in influencer
   */
  async getReferrals() {
    try {
      const response = await api.get('/influencer/referrals')
      return response.data
    } catch (error) {
      console.error('Error fetching referrals:', error)
      throw error
    }
  }

  /**
   * Get referral statistics for the logged-in influencer
   */
  async getReferralStats() {
    try {
      const response = await api.get('/influencer/stats')
      return response.data
    } catch (error) {
      console.error('Error fetching referral statistics:', error)
      throw error
    }
  }

  /**
   * Share referral link to social media
   */
  async shareReferralLink(platform: string, url: string, message: string) {
    try {
      const response = await api.post('/referrals/share', {
        platform,
        url,
        message
      })
      return response.data
    } catch (error) {
      console.error('Error sharing referral link:', error)
      throw error
    }
  }

  /**
   * Get referrals for a specific product
   */
  async getReferralsByProduct(productId: string) {
    try {
      const response = await api.get(`/referrals/product/${productId}`)
      return response.data
    } catch (error) {
      console.error('Error fetching referrals by product:', error)
      throw error
    }
  }

  /**
   * Get referrals for a specific time period
   */
  async getReferralsByPeriod(startDate: string, endDate: string) {
    try {
      const response = await api.get(`/referrals/period?startDate=${startDate}&endDate=${endDate}`)
      return response.data
    } catch (error) {
      console.error('Error fetching referrals by period:', error)
      throw error
    }
  }
}

export default new ReferralService()