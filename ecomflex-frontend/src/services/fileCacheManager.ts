// src/services/fileCacheManager.ts
// This service handles file caching for offline usage in the PWA

import { openDB, IDBPDatabase } from 'idb';

interface CachedFile {
  id: string;
  filename: string;
  type: string;
  data: Blob;
  url: string;
  size: number;
  timestamp: number;
  expiresAt: number | null;
  metadata?: any;
}

interface CacheStats {
  totalSize: number;
  fileCount: number;
  oldestFile: number;
  newestFile: number;
}

class FileCacheManager {
  private dbPromise: Promise<IDBPDatabase> | null = null;
  private readonly DB_NAME = 'ecomflex-file-cache';
  private readonly STORE_NAME = 'files';
  private readonly VERSION = 1;
  private maxCacheSize = 50 * 1024 * 1024; // 50MB default cache limit
  
  constructor() {
    this.initDB();
  }
  
  /**
   * Initialize the IndexedDB database
   */
  private initDB(): Promise<IDBPDatabase> {
    if (!this.dbPromise) {
      this.dbPromise = openDB(this.DB_NAME, this.VERSION, {
        upgrade(db) {
          // Create the file store if it doesn't exist
          if (!db.objectStoreNames.contains('files')) {
            const store = db.createObjectStore('files', { keyPath: 'id' });
            store.createIndex('timestamp', 'timestamp', { unique: false });
            store.createIndex('expiresAt', 'expiresAt', { unique: false });
            store.createIndex('url', 'url', { unique: true });
          }
        }
      });
    }
    return this.dbPromise;
  }
  
  /**
   * Set the maximum cache size in bytes
   * @param sizeInBytes Maximum cache size in bytes
   */
  public setMaxCacheSize(sizeInBytes: number): void {
    this.maxCacheSize = sizeInBytes;
    // Trigger cleanup if needed
    this.cleanupCache();
  }
  
  /**
   * Cache a file for offline use
   * @param file The file data to cache
   * @param ttl Time to live in milliseconds (optional)
   */
  public async cacheFile(file: Omit<CachedFile, 'timestamp' | 'expiresAt'>): Promise<string> {
    const db = await this.initDB();
    
    // Check if cache is full
    const stats = await this.getCacheStats();
    if (stats.totalSize + file.size > this.maxCacheSize) {
      await this.makeSpace(file.size);
    }
    
    const timestamp = Date.now();
    const expiresAt = null; // No expiration by default
    
    const cachedFile: CachedFile = {
      ...file,
      timestamp,
      expiresAt
    };
    
    try {
      await db.put(this.STORE_NAME, cachedFile);
      return file.id;
    } catch (error) {
      console.error('Error caching file:', error);
      throw error;
    }
  }
  
  /**
   * Cache a file with time-to-live
   * @param file The file data to cache
   * @param ttl Time to live in milliseconds
   */
  public async cacheFileWithTTL(file: Omit<CachedFile, 'timestamp' | 'expiresAt'>, ttl: number): Promise<string> {
    const db = await this.initDB();
    
    // Check if cache is full
    const stats = await this.getCacheStats();
    if (stats.totalSize + file.size > this.maxCacheSize) {
      await this.makeSpace(file.size);
    }
    
    const timestamp = Date.now();
    const expiresAt = timestamp + ttl;
    
    const cachedFile: CachedFile = {
      ...file,
      timestamp,
      expiresAt
    };
    
    try {
      await db.put(this.STORE_NAME, cachedFile);
      return file.id;
    } catch (error) {
      console.error('Error caching file:', error);
      throw error;
    }
  }
  
  /**
   * Retrieve a file from cache by ID
   * @param id The ID of the cached file
   */
  public async getFileById(id: string): Promise<CachedFile | null> {
    const db = await this.initDB();
    
    try {
      const file = await db.get(this.STORE_NAME, id);
      
      // Check if file exists and is not expired
      if (file && (file.expiresAt === null || file.expiresAt > Date.now())) {
        return file;
      }
      
      // File expired, delete it
      if (file && file.expiresAt !== null && file.expiresAt <= Date.now()) {
        await this.deleteFile(id);
      }
      
      return null;
    } catch (error) {
      console.error('Error retrieving file:', error);
      return null;
    }
  }
  
  /**
   * Retrieve a file from cache by URL
   * @param url The URL of the cached file
   */
  public async getFileByUrl(url: string): Promise<CachedFile | null> {
    const db = await this.initDB();
    
    try {
      const tx = db.transaction(this.STORE_NAME, 'readonly');
      const index = tx.store.index('url');
      const file = await index.get(url);
      
      // Check if file exists and is not expired
      if (file && (file.expiresAt === null || file.expiresAt > Date.now())) {
        return file;
      }
      
      // File expired, delete it
      if (file && file.expiresAt !== null && file.expiresAt <= Date.now()) {
        await this.deleteFile(file.id);
      }
      
      return null;
    } catch (error) {
      console.error('Error retrieving file by URL:', error);
      return null;
    }
  }
  
  /**
   * Delete a file from the cache
   * @param id The ID of the file to delete
   */
  public async deleteFile(id: string): Promise<boolean> {
    const db = await this.initDB();
    
    try {
      await db.delete(this.STORE_NAME, id);
      return true;
    } catch (error) {
      console.error('Error deleting file:', error);
      return false;
    }
  }
  
  /**
   * Clear all files from the cache
   */
  public async clearCache(): Promise<boolean> {
    const db = await this.initDB();
    
    try {
      await db.clear(this.STORE_NAME);
      return true;
    } catch (error) {
      console.error('Error clearing cache:', error);
      return false;
    }
  }
  
  /**
   * Get statistics about the file cache
   */
  public async getCacheStats(): Promise<CacheStats> {
    const db = await this.initDB();
    
    try {
      const tx = db.transaction(this.STORE_NAME, 'readonly');
      const store = tx.store;
      
      let totalSize = 0;
      let fileCount = 0;
      let oldestFile = Date.now();
      let newestFile = 0;
      
      // Fixed async iteration pattern
      let cursor = await store.openCursor();
      while (cursor) {
        const file = cursor.value as CachedFile;
        totalSize += file.size;
        fileCount++;
        
        if (file.timestamp < oldestFile) {
          oldestFile = file.timestamp;
        }
        
        if (file.timestamp > newestFile) {
          newestFile = file.timestamp;
        }
        
        cursor = await cursor.continue();
      }
      
      return {
        totalSize,
        fileCount,
        oldestFile: fileCount > 0 ? oldestFile : 0,
        newestFile
      };
    } catch (error) {
      console.error('Error getting cache stats:', error);
      return {
        totalSize: 0,
        fileCount: 0,
        oldestFile: 0,
        newestFile: 0
      };
    }
  }
  
  /**
   * Make space in the cache for a new file
   * @param neededBytes Number of bytes needed
   */
  private async makeSpace(neededBytes: number): Promise<void> {
    const db = await this.initDB();
    
    try {
      const tx = db.transaction(this.STORE_NAME, 'readwrite');
      const store = tx.store;
      const timestampIndex = store.index('timestamp');
      
      // First delete expired files
      const now = Date.now();
      const expiredFiles = await db.getAllFromIndex(this.STORE_NAME, 'expiresAt', IDBKeyRange.upperBound(now));
      
      for (const file of expiredFiles) {
        await store.delete(file.id);
      }
      
      // Check if we need to delete more files
      const stats = await this.getCacheStats();
      
      if (stats.totalSize + neededBytes <= this.maxCacheSize) {
        return; // We have enough space now
      }
      
      // Delete oldest files until we have enough space
      let spaceNeeded = stats.totalSize + neededBytes - this.maxCacheSize;
      
      // Get files ordered by timestamp (oldest first)
      let cursor = await timestampIndex.openCursor();
      
      while (cursor && spaceNeeded > 0) {
        const file = cursor.value as CachedFile;
        await store.delete(file.id);
        spaceNeeded -= file.size;
        cursor = await cursor.continue();
      }
    } catch (error) {
      console.error('Error making space in cache:', error);
      throw error;
    }
  }
  
  /**
   * Clean up the cache by removing expired files
   */
  public async cleanupCache(): Promise<void> {
    const db = await this.initDB();
    
    try {
      const now = Date.now();
      const tx = db.transaction(this.STORE_NAME, 'readwrite');
      const store = tx.store;
      
      // Get all expired files
      const expiredFiles = await db.getAllFromIndex(
        this.STORE_NAME, 
        'expiresAt', 
        IDBKeyRange.upperBound(now)
      );
      
      // Delete expired files
      for (const file of expiredFiles) {
        if (file.expiresAt !== null) {
          await store.delete(file.id);
        }
      }
      
      // Check if we need to delete more files to respect size limit
      const stats = await this.getCacheStats();
      
      if (stats.totalSize > this.maxCacheSize) {
        let spaceToFree = stats.totalSize - this.maxCacheSize;
        const timestampIndex = store.index('timestamp');
        let cursor = await timestampIndex.openCursor();
        
        while (cursor && spaceToFree > 0) {
          const file = cursor.value as CachedFile;
          await store.delete(file.id);
          spaceToFree -= file.size;
          cursor = await cursor.continue();
        }
      }
    } catch (error) {
      console.error('Error cleaning up cache:', error);
    }
  }
  
  /**
   * Cache an image from a URL
   * @param url The URL of the image
   * @param id Optional ID for the image
   * @param ttl Optional time-to-live in milliseconds
   */
  public async cacheImageFromUrl(url: string, id?: string, ttl?: number): Promise<string | null> {
    try {
      // Check if already cached
      const existingFile = await this.getFileByUrl(url);
      if (existingFile) {
        return existingFile.id;
      }
      
      // Fetch the image
      const response = await fetch(url, { cache: 'no-store' });
      
      if (!response.ok) {
        throw new Error(`Failed to fetch image: ${response.statusText}`);
      }
      
      const blob = await response.blob();
      const fileId = id || `img_${Date.now()}_${Math.random().toString(36).substr(2, 9)}`;
      
      const file = {
        id: fileId,
        filename: url.split('/').pop() || 'image',
        type: blob.type,
        data: blob,
        url: url,
        size: blob.size,
        metadata: {
          source: 'url',
          cached: Date.now()
        }
      };
      
      // Cache with or without TTL
      if (ttl) {
        await this.cacheFileWithTTL(file, ttl);
      } else {
        await this.cacheFile(file);
      }
      
      return fileId;
    } catch (error) {
      console.error('Error caching image from URL:', error);
      return null;
    }
  }
  
  /**
   * Get an object URL for a cached file
   * @param id The ID of the cached file
   */
  public async getObjectUrl(id: string): Promise<string | null> {
    const file = await this.getFileById(id);
    
    if (file) {
      return URL.createObjectURL(file.data);
    }
    
    return null;
  }
  
  /**
   * Get all cached files
   */
  public async getAllFiles(): Promise<CachedFile[]> {
    const db = await this.initDB();
    
    try {
      return await db.getAll(this.STORE_NAME);
    } catch (error) {
      console.error('Error getting all files:', error);
      return [];
    }
  }
  
  /**
   * Update the metadata for a cached file
   * @param id The ID of the cached file
   * @param metadata The new metadata
   */
  public async updateMetadata(id: string, metadata: any): Promise<boolean> {
    const db = await this.initDB();
    
    try {
      const file = await this.getFileById(id);
      
      if (!file) {
        return false;
      }
      
      file.metadata = { ...file.metadata, ...metadata };
      await db.put(this.STORE_NAME, file);
      return true;
    } catch (error) {
      console.error('Error updating metadata:', error);
      return false;
    }
  }
}

// Create and export a singleton instance
const fileCacheManager = new FileCacheManager();
export default fileCacheManager;