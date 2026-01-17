// src/services/imageProxy.ts
const ECOMFLEX_S3_URL = 'https://ecomflex-uploads-dev.s3.eu-north-1.amazonaws.com';

export class ImageProxyService {
  // Log detailed info about image loading attempts
  static logImageAttempt(imageUrl: string, productId: string, productName: string) {
    console.log(`[ImageProxy] Loading image for product ${productId} (${productName})`);
    console.log(`[ImageProxy] Raw URL: ${imageUrl}`);
  }

  // Get a working image URL with a timestamp to bust cache
  static getProxyUrl(url: string | undefined): string {
    if (!url) return '';
    
    // Add cache-busting parameter
    const cacheBuster = `?t=${Date.now()}`;
    
    // Handle S3 URLs specifically
    if (url.includes('ecomflex-uploads-dev.s3') || url.includes('your-s3-bucket')) {
      // Get just the key part (after the domain and bucket)
      let parts = url.split('/products/');
      if (parts.length === 2) {
        const key = parts[1];
        return `${ECOMFLEX_S3_URL}/products/${key}${cacheBuster}`;
      }
    }
    
    // Just add cache buster for normal URLs
    return url + cacheBuster;
  }
  
  // Check if an image is accessible
  static checkImageExists(url: string): Promise<boolean> {
    return new Promise((resolve) => {
      const img = new Image();
      img.onload = () => resolve(true);
      img.onerror = () => resolve(false);
      img.src = url;
    });
  }
}

export default ImageProxyService;