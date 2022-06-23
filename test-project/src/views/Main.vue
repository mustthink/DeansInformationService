<template>
<div>
    <navbar></navbar>
    <!-- <h1 style="margin:50px; font-size: 1.5rem;">Сайт работает :)</h1> -->
    <!-- <button @click="fetchNews">Обновить список новостей</button> -->
    <news-list v-bind:news="news"/>
    <!-- <button @click="WriteJson">Записать и перейти</button> -->
</div>
</template>

<script>
import Navbar from "@/components/UI/Navbar.vue";
import NewsList from "@/components/NewsList.vue";
import axios from "axios";
// import { json } from "body-parser";
export default {
    components: { Navbar, NewsList},
    data(){
        return{
            news:[]
        }
    },
    methods:{
        WriteJson(){
            // const fs = require('fs');

            const user = {
            "id": 1,
            "name": "John Doe",
            "age": 22
            };

            const data = JSON.stringify(user);

            fs.writeFile('user.json', data, (err) => {
            if (err) {
                throw err;
            }
            console.log("JSON data is saved.");
            });
        },

        async fetchNews(){
            try {
                const response = await axios.get('https://jsonplaceholder.typicode.com/posts?_limit=5');
                this.news = response.data;
            } catch (error) {
                alert('Произошла чудовищная ошибка: ', error)
            }
        },


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

</style>