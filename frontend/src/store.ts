import { defineStore } from 'pinia';
const useConfigStore = defineStore("config", {
  state: () => ({
    dark: false
  }),
  actions: {
    toggleDark() {
      document.documentElement.classList.toggle('dark')
      this.dark = document.documentElement.classList.contains('dark')
    },
    initPage() {
      if (this.dark) {
        document.documentElement.classList.add("dark");
      } else {
        document.documentElement.classList.remove("dark");
      }
    }
  },
  persist: true
})
export {useConfigStore}
