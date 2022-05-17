import { createStore } from "vuex"

export default createStore({
   
    state:{
        isLogged: false
    },
    getters:{
        
    },
    mutations:{
        setLogged(state){
            state.isLogged = true;
        },
    },
    actions:{

    },
    modules:{

    },
})