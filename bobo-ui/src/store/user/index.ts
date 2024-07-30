import { defineStore } from "pinia";
import {  Login, Profile } from "@/api/user";

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
      avatar: "",
    };
  },
  // 异步|逻辑的地方
  actions: {
    //登陆
    async userLogin(data: any) {
      let res = await Login(data);
      if (res.code == 200) {
        this.token = res.data;
        return res.message;
      } else {
        return res.message;
      }
    },
    // 用户详情
    async userProfile() {
      let res = await Profile();
      if (res.code == 200) {
        this.created_at = res.data.created_at;
        this.email = res.data.email;
        this.id = res.data.id;
        this.ip_address = res.data.ip_address;
        this.ip_source = res.data.ip_source;
        this.last_login_time = res.data.last_login_time;
        this.nickname = res.data.nickname;
        this.role = res.data.role;
        this.status = res.data.status;
        this.username = res.data.username;
        this.avatar = res.data.avatar;
        return res.message;
      } else {
        return res.message;
      }
    },
    // 退出登陆
    async userLogout() {
      this.token = "";
      this.created_at = null;
      this.email = null;
      this.id = null;
      this.ip_address = null;
      this.ip_source = null;
      this.last_login_time = null;
      this.nickname = null;
      this.role = null;
      this.status = null;
      this.username = null;
      this.avatar = null;
    },
  },
  getters: {},
  persist: true,
});

export default useUserStore;
