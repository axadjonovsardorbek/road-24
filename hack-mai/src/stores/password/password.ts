import create from 'zustand';

interface UserPasswordState {
  passwords: { [username: string]: string };
  setPassword: (username: string, password: string) => void;
  getPassword: (username: string) => string | undefined;
}

export const useUserPasswordStore = create<UserPasswordState>((set, get) => ({
  passwords: {},
  setPassword: (username, password) =>
    set((state) => ({
      passwords: { ...state.passwords, [username]: password },
    })),
  getPassword: (username) => get().passwords[username],
}));
