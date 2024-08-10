// src/services/authService.ts

const mockToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkFkYW0iLCJpYXQiOjE1MTYyMzkwMjJ9.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c";

export const login = (email: string, password: string): Promise<string> => {
  return new Promise((resolve, reject) => {
    // Simulate API call
    setTimeout(() => {
      if (email === "test@test.com" && password === "test") {
        localStorage.setItem("token", mockToken);
        resolve(mockToken);
      } else {
        reject(new Error("Invalid credentials"));
      }
    }, 1000);
  });
};

export const logout = (): void => {
  localStorage.removeItem("token");
};

export const isAuthenticated = (): boolean => {
  return !!localStorage.getItem("token");
};