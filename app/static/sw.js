const CACHE_NAME = 'plants-images-v1';
const API_CACHE_NAME = 'plants-api-v1';
const MAX_ENTRIES = 200;

self.addEventListener('install', (event) => {
  self.skipWaiting();
});

self.addEventListener('activate', (event) => {
  event.waitUntil(self.clients.claim());
});

self.addEventListener('fetch', (event) => {
  const req = event.request;
  const url = new URL(req.url);
  
  // Cache images with cache-first strategy
  const isImage = req.destination === 'image' || /\.(png|jpg|jpeg|webp|gif|svg)(\?.*)?$/.test(url.pathname);
  if (isImage && req.method === 'GET') {
    event.respondWith(
      caches.open(CACHE_NAME).then(async (cache) => {
        const cached = await cache.match(req);
        const fetchPromise = fetch(req)
          .then((res) => {
            if (res && res.ok) {
              cache.put(req, res.clone()).catch(() => {});
            }
            return res;
          })
          .catch(() => cached);
        return cached || fetchPromise;
      })
    );
    return;
  }

  // Cache API GET requests with network-first, fall back to cache on offline
  const isAPI = url.pathname.startsWith('/api/');
  if (isAPI && req.method === 'GET') {
    event.respondWith(
      caches.open(API_CACHE_NAME).then(async (cache) => {
        try {
          const res = await fetch(req);
          if (res && res.ok) {
            cache.put(req, res.clone()).catch(() => {});
          }
          return res;
        } catch (err) {
          const cached = await cache.match(req);
          if (cached) return cached;
          throw err;
        }
      })
    );
    return;
  }

  // Cache successful POST/PATCH responses (e.g., plant updates) to serve on offline
  if (isAPI && (req.method === 'POST' || req.method === 'PATCH')) {
    const reqClone = req.clone();
    event.respondWith(
      fetch(req).then((res) => {
        if (res && res.ok && res.status < 300) {
          // Cache the response for potential offline viewing
          caches.open(API_CACHE_NAME).then((cache) => {
            cache.put(reqClone, res.clone()).catch(() => {});
          });
        }
        return res.clone();
      }).catch(() => {
        // On offline, return a basic error response
        return new Response(JSON.stringify({ error: { message: 'Network unavailable' } }), {
          status: 503,
          headers: { 'Content-Type': 'application/json' }
        });
      })
    );
    return;
  }
});
