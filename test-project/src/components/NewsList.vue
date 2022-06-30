<template>
    <div>
    <my-button v-if="this.$store.state.isDean" class="btn" @click="$router.push('/addnews')">Создать пост</my-button>
    <h2 class="news_article">Последние новости:</h2>
    <NewsItem
    v-for="onenew in news" 
    :onenew="onenew"
    :key="onenew.id"
    />
  </div>
</template>

<script>
import NewsItem from "./NewsItem.vue";
export default {
    components: { NewsItem },
    props: {
        news: {
            type: Array,
            required: true
        }
  }
}
</script>

<style scoped>

.btn{
    margin: 15px;
    
    margin-top: 30px;
    font-size: 1rem;
}
.news_article{
    margin: 15px;
    margin-left: 14vw;
    margin-top: 15px;
    font-size: 1.5rem;
}
</style>