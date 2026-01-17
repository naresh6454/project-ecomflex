// sw.ts - Service Worker for Ecomflex PWA
/// <reference lib="webworker" />

import { precacheAndRoute, createHandlerBoundToURL } from 'workbox-precaching';
import { registerRoute, NavigationRoute } from 'workbox-routing';
import { NetworkFirst, StaleWhileRevalidate, CacheFirst } from 'workbox-strategies';
import { ExpirationPlugin } from 'workbox-expiration';
import { CacheableResponsePlugin } from 'workbox-cacheable-response';

declare const self: ServiceWorkerGlobalScope;

// Use with precache injection
precacheAndRoute(self.__WB_MANIFEST);

// Cache the Google Fonts stylesheets with a stale-while-revalidate strategy
registerRoute(
  ({ url }) => url.origin === 'https://fonts.googleapis.com',
  new StaleWhileRevalidate({
    cacheName: 'google-fonts-stylesheets',
  })
);

// Cache the underlying font files with a cache-first strategy for 1 year
registerRoute(
  ({ url }) => url.origin === 'https://fonts.gstatic.com',
  new CacheFirst({
    cacheName: 'google-fonts-webfonts',
    plugins: [
      new CacheableResponsePlugin({
        statuses: [0, 200],
      }),
      new ExpirationPlugin({
        maxAgeSeconds: 60 * 60 * 24 * 365, // 1 year
        maxEntries: 30,
      }),
    ],
  })
);

// Cache images
registerRoute(
  ({ request }) => request.destination === 'image',
  new CacheFirst({
    cacheName: 'images',
    plugins: [
      new ExpirationPlugin({
        maxEntries: 60,
        maxAgeSeconds: 30 * 24 * 60 * 60, // 30 days
      }),
    ],
  })
);

// Cache API responses
registerRoute(
  ({ url }) => url.pathname.startsWith('/api/'),
  new NetworkFirst({
    cacheName: 'api-cache',
    plugins: [
      new ExpirationPlugin({
        maxEntries: 100,
        maxAgeSeconds: 5 * 60, // 5 minutes
      }),
    ],
  })
);

// Static assets (JS, CSS)
registerRoute(
  ({ request }) => 
    request.destination === 'script' || 
    request.destination === 'style',
  new StaleWhileRevalidate({
    cacheName: 'static-resources',
  })
);

// Single Page Application Navigation
const handler = createHandlerBoundToURL('/index.html');
const navigationRoute = new NavigationRoute(handler, {
  // Don't serve the handler for API requests or other specific files
  denylist: [
    /^\/_/,
    /\/[^/?]+\.[^/]+$/,
    /\/api\//,
  ],
});
registerRoute(navigationRoute);

// Offline fallback page
const FALLBACK_HTML_URL = '/offline.html';
const FALLBACK_IMAGE_URL = '/images/offline-image.svg';

// Register a specific route for the offline page
registerRoute(
  ({ request }) => request.mode === 'navigate',
  async ({ event }) => {
    try {
      // Try to fetch the page normally
      return await new NetworkFirst({
        cacheName: 'pages',
        plugins: [
          new ExpirationPlugin({
            maxEntries: 25,
            maxAgeSeconds: 7 * 24 * 60 * 60, // 1 week
          }),
        ],
      }).handle({ event } as any);
    } catch (error) {
      // If offline and page not in cache, show the offline page
      const cache = await caches.open('offline');
      const cachedResponse = await cache.match(FALLBACK_HTML_URL);
      
      // Make sure we always return a Response
      if (cachedResponse) {
        return cachedResponse;
      } else {
        // If the offline page is not in cache, return a basic HTML response
        return new Response(
          '<html><body><h1>Offline</h1><p>You are currently offline.</p></body></html>',
          {
            headers: { 'Content-Type': 'text/html' }
          }
        );
      }
    }
  }
);

// Handle offline images
registerRoute(
  ({ request }) => request.destination === 'image',
  async ({ event }) => {
    try {
      return await new CacheFirst({
        cacheName: 'images',
        plugins: [
          new ExpirationPlugin({
            maxEntries: 60,
            maxAgeSeconds: 30 * 24 * 60 * 60, // 30 days
          }),
        ],
      }).handle({ event } as any);
    } catch (error) {
      const cache = await caches.open('offline');
      const cachedResponse = await cache.match(FALLBACK_IMAGE_URL);
      
      // Make sure we always return a Response
      if (cachedResponse) {
        return cachedResponse;
      } else {
        // If the fallback image is not in cache, return a transparent 1x1 pixel GIF
        // Data URI for a 1x1 transparent GIF
        const TRANSPARENT_GIF = 'data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7';
        
        // Create a Response with the data URI
        const response = await fetch(TRANSPARENT_GIF);
        return response;
      }
    }
  }
);

// Install and activate events
self.addEventListener('install', (event) => {
  event.waitUntil(
    caches.open('offline').then((cache) => {
      return cache.addAll([
        FALLBACK_HTML_URL,
        FALLBACK_IMAGE_URL,
        '/styles/offline.css',
      ]);
    })
  );
});

self.addEventListener('activate', (event) => {
  // Clean up old caches
  event.waitUntil(
    caches.keys().then((cacheNames) => {
      return Promise.all(
        cacheNames
          .filter((cacheName) => {
            // Delete old versioned caches
            return cacheName.startsWith('workbox-') && cacheName !== 'workbox-precache';
          })
          .map((cacheName) => {
            return caches.delete(cacheName);
          })
      );
    })
  );
});

// Message handling for updating PWA
self.addEventListener('message', (event) => {
  if (event.data && event.data.type === 'SKIP_WAITING') {
    self.skipWaiting();
  }
});