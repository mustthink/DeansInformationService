import { createStore } from "vuex"

export default createStore({
   
    state:{
        isLogged: true, //потом обяpательно поменять обратно на false
        //(изменил временно чтобы не ебаться каждый раз)
        isStudent: true,
        isEducator: false,
        isDean: false,
    },
    getters:{
        
    },
    mutations:{
        setLogged(state){
            state.isLogged = true;
        },
        setStudent(state){
            state.isStudent = true;
        },
        setEducator(state){
            state.isEducator = true;
        },
        setDean(state){
            state.isDean = true;
        },
    },
    actions:{

    },
    modules:{

    },
})