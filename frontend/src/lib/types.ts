export interface User {
  id: number;
  email: string;
  username: string;
  first_name: string;
  last_name: string;
  program: string;
  year: number;
  advisor: string;
}

export interface AuthResponse {
  message: string;
  token: string;
  user: User;
}

export interface Course {
  id: number;
  user_id: number;
  course_name: string;
  course_code: string;
  instructor: string;
  credits: number;
  semester: string;
  grade: string;
  status: string;
  created_at: string;
  updated_at: string;
}

export interface Assignment {
  id: number;
  user_id: number;
  course_id?: number;
  course?: Course;
  title: string;
  description: string;
  due_date: string;
  priority: 'low' | 'medium' | 'high';
  status: 'pending' | 'in_progress' | 'completed';
  estimated_hours: number;
  actual_hours: number;
  created_at: string;
  updated_at: string;
}

export interface CreateCourseRequest {
  course_name: string;
  course_code: string;
  instructor: string;
  credits: number;
  semester: string;
  status?: string;
}

export interface CreateAssignmentRequest {
  course_id?: number;
  title: string;
  description: string;
  due_date: string;
  priority: string;
  estimated_hours: number;
}