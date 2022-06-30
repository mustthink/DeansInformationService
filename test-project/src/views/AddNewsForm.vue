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

    export default {

        data(){
            return {
                PeaceOfNews: {
                    title:'',
                    body:'',
                }
            }
        },
        methods:{
            CreateNews(){
            if(this.PeaceOfNews.title.length > 0 && this.PeaceOfNews.body.length > 0){
                const newNews = {
                    id: Date.now(),
                    title: this.PeaceOfNews.title,
                    body:this.PeaceOfNews.body,
                }
                this.$store.commit('setArticle',newNews.title);
                this.$store.commit('setBody',newNews.body);
                axios.post(`172.16.1.4:8081/news/create`, {
                    title: newNews.title,
                    body: newNews.body,
                });

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