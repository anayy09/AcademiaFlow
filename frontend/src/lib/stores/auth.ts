import { writable } from 'svelte/store';
import { browser } from '$app/environment';
import type { User } from '$lib/types';

interface AuthState {
  user: User | null;
  token: string | null;
  isAuthenticated: boolean;
  isLoading: boolean;
}

const initialState: AuthState = {
  user: null,
  token: null,
  isAuthenticated: false,
  isLoading: true
};

function createAuthStore() {
  const { subscribe, set, update } = writable<AuthState>(initialState);

  return {
    subscribe,
    login: (user: User, token: string) => {
      if (browser) {
        localStorage.setItem('auth_token', token);
        localStorage.setItem('user', JSON.stringify(user));
      }
      set({
        user,
        token,
        isAuthenticated: true,
        isLoading: false
      });
    },
    logout: () => {
      if (browser) {
        localStorage.removeItem('auth_token');
        localStorage.removeItem('user');
      }
      set({
        user: null,
        token: null,
        isAuthenticated: false,
        isLoading: false
      });
    },
    initialize: () => {
      if (browser) {
        const token = localStorage.getItem('auth_token');
        const userStr = localStorage.getItem('user');
        
        if (token && userStr) {
          try {
            const user = JSON.parse(userStr);
            set({
              user,
              token,
              isAuthenticated: true,
              isLoading: false
            });
          } catch (error) {
            console.error('Error parsing stored user data:', error);
            set({ ...initialState, isLoading: false });
          }
        } else {
          set({ ...initialState, isLoading: false });
        }
      }
    },
    setLoading: (loading: boolean) => {
      update(state => ({ ...state, isLoading: loading }));
    }
  };
}

export const authStore = createAuthStore();