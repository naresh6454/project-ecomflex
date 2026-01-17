// URL Interceptor to fix S3 placeholder URLs
const originalFetch = window.fetch;

// Replace all placeholder S3 URLs with the real bucket
const fixS3Url = (url: string) => {
  if (typeof url === 'string' && url.includes('your-s3-bucket')) {
    console.log('Intercepted placeholder S3 URL:', url);
    
    // Extract the path and filename from the placeholder URL
    const urlParts = url.replace('https://your-s3-bucket.s3.amazonaws.com/', '').split('/');
    const filename = urlParts.pop() || '';
    const path = urlParts.length ? urlParts.join('/') + '/' : '';
    
    // Replace with the real bucket URL
    const fixedUrl = `https://ecomflex-uploads-dev.s3.eu-north-1.amazonaws.com/${path}${filename}`;
    console.log('Fixed to:', fixedUrl);
    return fixedUrl;
  }
  return url;
};

// Override fetch API
window.fetch = function(input, init) {
  if (typeof input === 'string') {
    input = fixS3Url(input);
  } else if (input instanceof Request) {
    input = new Request(fixS3Url(input.url), input);
  }
  return originalFetch.call(this, input, init);
};

// Override XMLHttpRequest
const originalXhrOpen = XMLHttpRequest.prototype.open;
XMLHttpRequest.prototype.open = function(method: string, url: string | URL, async: boolean = true, username?: string | null, password?: string | null): void {
  if (typeof url === 'string') {
    url = fixS3Url(url);
  }
  return originalXhrOpen.call(this, method, url, async, username || null, password || null);
};

// Override image src attributes
document.addEventListener('DOMContentLoaded', () => {
  // Fix image src attributes on page load
  const fixImageSrc = () => {
    const images = document.querySelectorAll('img[src*="your-s3-bucket"]');
    images.forEach(img => {
      const originalSrc = img.getAttribute('src');
      if (originalSrc) {
        img.setAttribute('src', fixS3Url(originalSrc));
      }
    });
  };

  // Run immediately
  fixImageSrc();
  
  // Also run a MutationObserver to catch dynamically added images
  const observer = new MutationObserver((mutations) => {
    mutations.forEach((mutation) => {
      if (mutation.type === 'attributes' && 
          mutation.attributeName === 'src' && 
          mutation.target instanceof HTMLImageElement) {
        const img = mutation.target;
        const src = img.getAttribute('src');
        if (src && src.includes('your-s3-bucket')) {
          img.setAttribute('src', fixS3Url(src));
        }
      } else if (mutation.addedNodes.length > 0) {
        fixImageSrc();
      }
    });
  });
  
  observer.observe(document.body, { 
    attributes: true, 
    attributeFilter: ['src'], 
    childList: true, 
    subtree: true 
  });
});

// Log that the interceptor is active
console.log('URL Interceptor loaded - S3 placeholder URLs will be automatically fixed');

export default { fixS3Url };