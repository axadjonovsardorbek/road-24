import create from 'zustand';

interface MenuStore {
  activeMenuKey: string | null;
  activeSubMenuKeys: string[];
  setActiveMenuKey: (key: string | null) => void;
  setActiveSubMenuKeys: (keys: string[]) => void;
}

export const useMenuStore = create<MenuStore>(set => ({
  activeMenuKey: null,
  activeSubMenuKeys: [],
  setActiveMenuKey: (key: string | null) => set({ activeMenuKey: key }),
  setActiveSubMenuKeys: (keys: string[]) => set({ activeSubMenuKeys: keys }),
}));
