import { AuthAPI } from "@/api";

export const AuthModule = {
    namespaced: true,

    state(){
        return{
            credentials: {
                token: null,
                userRole: null,
            }
        }
    },

    getters:{
        getUserRole: (state) => state.credentials.userRole,
    },

    mutations: {
        setToken(state, token){
            state.credentials.token = token;
        },
        setUserRole(state, userRole){
            state.credentials.userRole = userRole;
        },

    },

    actions:{
        onLogin({commit}, {login, password}){
            AuthAPI.login({login, password}).then(()=>{
                commit('setToken', res.data.token);
                commit('setUserRole', res.data.userRole);
            })
        },

        onLogout(){

        }
    }
}