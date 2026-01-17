// src/vue-router.d.ts
declare module 'vue-router' {
  // Import and re-export all types from vue-router
  export interface RouteRecordRaw {
    path: string;
    redirect?: string | { name?: string; path?: string; params?: Record<string, any>; query?: Record<string, any> };
    name?: string;
    component?: any;
    components?: Record<string, any>;
    children?: RouteRecordRaw[];
    meta?: Record<string, any>;
    props?: boolean | Record<string, any> | ((to: any) => Record<string, any>);
    beforeEnter?: any;
    [key: string]: any;
  }

  export interface RouteLocationNormalized {
    path: string;
    fullPath: string;
    hash: string;
    query: Record<string, string | string[]>;
    params: Record<string, string>;
    name: string | null | undefined;
    matched: any[];
    meta: Record<string, any>;
    redirectedFrom?: RouteLocationNormalized;
    [key: string]: any;
  }
  
  export type NavigationGuardNext = (to?: any) => void;
  
  export function createRouter(options: {
    history: any;
    routes: RouteRecordRaw[];
    scrollBehavior?: any;
    [key: string]: any;
  }): any;
  
  export function createWebHistory(base?: string): any;
  export function createWebHashHistory(base?: string): any;
  export function createMemoryHistory(base?: string): any;
  
  export function useRouter(): any;
  export function useRoute(): RouteLocationNormalized;
  
  export const RouterLink: any;
  export const RouterView: any;
}