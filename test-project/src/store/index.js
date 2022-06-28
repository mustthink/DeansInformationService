import { createStore } from "vuex"

export default createStore({
   
    state:{
        isLogged: false, //если стоит true, то потом обяpательно поменять обратно на false
        //(изменил временно чтобы не ебаться каждый раз)
        isStudent: true,
        isEducator: false,
        isDean: false,
        news:{
            title: ``,
            body: ``,
        }
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
        setArticle(state, title){
            state.news.title = title;
        },
        setBody(state, body){
            state.news.body = body;
        }
    },
    actions:{

    },
    modules:{

    },
})