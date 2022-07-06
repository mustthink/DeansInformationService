<template>
<div>
    <navbar></navbar>
    <!-- <my-button style='margin-top: 30px; margin-left: 30px;' @click="$router.push('/addnews')">Создать пост</my-button> -->
    <my-button class="testBtns" @click="setDean()">Set Dean</my-button>
    <news-list 
    :news="news"
    @remove="removeNews"
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
        async fetchNews(){
            try {
                const response = await axios.get('http://localhost:8081/news/list');
                this.news = response.data;
                console.log(response)
            } catch (error) {
                alert('Произошла чудовищная ошибка: ', error)
            }
        },
        removeNews(onenew){
            // axios.delete(`http://localhost:8081/news/delete?id=${onenew.ID}`)
            //альтернативный вид запроса (сверху)
            axios.delete(`http://localhost:8081/news/delete`,{ 
                params: { 
                    id: onenew.ID 
                } 
            }) 
            setTimeout(fetchNews(), 3000);
        },
        setDean(){
            if(this.$store.state.isDean == false){
                this.$store.commit('setDean');
                console.log("Dean's mode is done")
            }
            else{
                this.$store.commit('setNotDean');
                console.log("Dean's mode is undone")
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