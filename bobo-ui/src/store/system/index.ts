import { defineStore } from "pinia";
import { darkTheme } from 'naive-ui'
let useSystemStore = defineStore("System", {
  state: (): any => {
    return {
      dart: null,
    };
  },
  // 异步|逻辑的地方
  actions: {
    //登陆
    async changeTheme() {
       if(this.dart==null){
        this.dart=darkTheme
       }else{
        this.dart=null
       }
    },
  },
  getters: {},
  persist: true,
});

export default useSystemStore;
