import axios, { type AxiosInstance, type AxiosResponse } from 'axios';
import { authStore } from '$lib/stores/auth';
import { get } from 'svelte/store';
import type { 
  AuthResponse, 
  User, 
  Course, 
  Assignment, 
  CreateCourseRequest, 
  CreateAssignmentRequest 
} from '$lib/types';

class ApiService {
  private api: AxiosInstance;

  constructor() {
    this.api = axios.create({
      baseURL: 'http://localhost:8080/api/v1',
      headers: {
        'Content-Type': 'application/json',
      },
    });

    // Request interceptor to add auth token
    this.api.interceptors.request.use((config) => {
      const auth = get(authStore);
      if (auth.token) {
        config.headers.Authorization = `Bearer ${auth.token}`;
      }
      return config;
    });

    // Response interceptor for error handling
    this.api.interceptors.response.use(
      (response) => response,
      (error) => {
        if (error.response?.status === 401) {
          authStore.logout();
        }
        return Promise.reject(error);
      }
    );
  }

  // Auth endpoints
  async register(data: {
    email: string;
    username: string;
    password: string;
    first_name: string;
    last_name: string;
    program: string;
    year: number;
    advisor: string;
  }): Promise<AuthResponse> {
    const response: AxiosResponse<AuthResponse> = await this.api.post('/auth/register', data);
    return response.data;
  }

  async login(email: string, password: string): Promise<AuthResponse> {
    const response: AxiosResponse<AuthResponse> = await this.api.post('/auth/login', {
      email,
      password,
    });
    return response.data;
  }

  async getProfile(): Promise<{ user: User }> {
    const response = await this.api.get('/users/profile');
    return response.data;
  }

  async updateProfile(data: Partial<User>): Promise<{ user: User; message: string }> {
    const response = await this.api.put('/users/profile', data);
    return response.data;
  }

  // Course endpoints
  async getCourses(): Promise<{ courses: Course[] }> {
    const response = await this.api.get('/courses');
    return response.data;
  }

  async createCourse(data: CreateCourseRequest): Promise<{ course: Course; message: string }> {
    const response = await this.api.post('/courses', data);
    return response.data;
  }

  async updateCourse(id: number, data: Partial<Course>): Promise<{ course: Course; message: string }> {
    const response = await this.api.put(`/courses/${id}`, data);
    return response.data;
  }

  async deleteCourse(id: number): Promise<{ message: string }> {
    const response = await this.api.delete(`/courses/${id}`);
    return response.data;
  }

  // Assignment endpoints
  async getAssignments(filters?: { status?: string; priority?: string }): Promise<{ assignments: Assignment[] }> {
    const params = new URLSearchParams();
    if (filters?.status) params.append('status', filters.status);
    if (filters?.priority) params.append('priority', filters.priority);
    
    const response = await this.api.get(`/assignments?${params.toString()}`);
    return response.data;
  }

  async createAssignment(data: CreateAssignmentRequest): Promise<{ assignment: Assignment; message: string }> {
    const response = await this.api.post('/assignments', data);
    return response.data;
  }

  async updateAssignment(id: number, data: Partial<Assignment>): Promise<{ assignment: Assignment; message: string }> {
    const response = await this.api.put(`/assignments/${id}`, data);
    return response.data;
  }

  async updateAssignmentStatus(id: number, status: string): Promise<{ assignment: Assignment; message: string }> {
    const response = await this.api.patch(`/assignments/${id}/status`, { status });
    return response.data;
  }

  async deleteAssignment(id: number): Promise<{ message: string }> {
    const response = await this.api.delete(`/assignments/${id}`);
    return response.data;
  }
}

export const apiService = new ApiService();