<template>
    <div>
        <form @submit.prevent class="form">
            <h4>Создание новой новости</h4>
            
            <input 
            v-model="PeaceOfNews.title"
            placeholder="Введите заглавие новости"
            class="input"
            type="text">

            <input 
            v-model="PeaceOfNews.body"
            placeholder="Введите тело новости"
            class="input"
            type="text">
            <div class="btn">
                <my-button style="margin-top:15px" @click="CreateNews()">Создать</my-button>
                <my-button style="margin-top:15px" @click="$router.push('/news')">Просмотреть</my-button>
            </div>
            

        </form>
    </div>
</template>

<script>
    import axios from "axios";
    export default {

        data(){
            return {
                PeaceOfNews: {
                    title:'',
                    body:'',
                    //test sync
                }
            }
        },
        methods:{
            CreateNews(){
            if(this.PeaceOfNews.title.length > 0 && this.PeaceOfNews.body.length > 0){
                const newNews = {
                    title: this.PeaceOfNews.title,
                    content:this.PeaceOfNews.body,
                }
                // this.$store.commit('setArticle',newNews.title);
                // this.$store.commit('setBody',newNews.body);              
                // Title: newNews.title,
                // Content: newNews.body,
                let data = JSON.stringify(newNews)
                axios.post(`http://localhost:8081/news/create`, data);

            }
        }
        }
    }
</script>

<style scoped>
.form{
    margin: 15px;
}
.btn{
    display: grid;
    width: 150px;
}
.input{
    width: 50vw;    
    border: 2px solid teal;
    padding: 10px 15px;
    margin-top: 15px;
}
</style>