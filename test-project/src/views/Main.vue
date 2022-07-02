<template>
<div>
    <navbar></navbar>
    <!-- <my-button style='margin-top: 30px; margin-left: 30px;' @click="$router.push('/addnews')">Создать пост</my-button> -->
    <my-button class="testBtns" @click="setDean()">Set Dean</my-button>
    <news-list 
    :news="news"
    />
</div>
</template>

<script>
import Navbar from "@/components/UI/Navbar.vue";
import NewsList from "@/components/NewsList.vue";
import axios from "axios";
export default {
    components: { Navbar, NewsList},
    data(){
        return{
            news:[]
        }
    },
    methods:{
        // async 
        // fetchNews(){
        //     try {
        //         const response = axios.get('http://172.16.1.4:8081/news/list'); //await axios
        //         this.news = response.data;
        //     } catch (error) {
        //         alert('Произошла чудовищная ошибка: ', error)
        //     }
        // },
        async fetchNews(){
            try {
                const response = await fetch('http://172.16.1.4:8081/news/list');
                this.news = response.data;
            } catch (error) {
                alert('Произошла чудовищная ошибка: ', error)
            }
        },
        setDean(){
            if(this.$store.state.isDean == false){
                this.$store.commit('setDean');
                console.log('Worked')
            }
            else{
                this.$store.commit('setNotDean');
                console.log('Worked too')
            }
            
        }


    },
    mounted() {
        if (this.$store.state.isLogged == false) {
            this.$router.push("/login");
        }
        this.fetchNews();
    },
   
}
</script>

<style scoped>
.testBtns{
    position:absolute;
   top:90;
   right:0;
   background-color: white;
}
</style>