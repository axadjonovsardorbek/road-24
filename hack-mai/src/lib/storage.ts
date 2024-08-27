export function loadState(key:string) {
    try {
      const serializedState = localStorage.getItem(key);
      if (!serializedState) return undefined;
      return JSON.parse(serializedState);
    } catch (e) {
      return undefined;
    }
  }
  
  export async function saveState(key:string, state:string) {
    try {
      const serializedState = JSON.stringify(state);
      localStorage.setItem(key, serializedState);
    } catch (e) {}
  }