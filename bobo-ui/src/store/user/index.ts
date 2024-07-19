import { defineStore } from "pinia";

let useUserStore = defineStore("User", {
  state: (): any => {
    return {
      token: "",
      created_at: "",
      email: "",
      id: null,
      ip_address: "",
      ip_source: "",
      last_login_time: null,
      nickname: null,
      role: null,
      status: null,
      username: "",
      avatar:""
    };
  },
  // 异步|逻辑的地方
  actions: {},
  getters: {},
  persist: true,
});

export default useUserStore;
