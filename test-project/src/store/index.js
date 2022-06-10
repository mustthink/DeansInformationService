import { createStore } from "vuex"

export default createStore({
   
    state:{
        isLogged: false,
        isStudent: false,
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