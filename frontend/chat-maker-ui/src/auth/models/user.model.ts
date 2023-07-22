export interface User {
  name: string;
  email: string;
  picture: string;
  created_at: Date;
  id: string;
}
export interface UserWithLicenseValidation {
  user: User;
  is_licence_valid: boolean;
}

export interface UserCreated {
  is_created: boolean;
  user_id: string | null;
}
