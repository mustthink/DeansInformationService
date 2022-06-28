<template>
    <div>
    <h2 class="news_article">Последние новости:</h2>
    <NewsItem
    v-for="onenew in news" 
    :onenew="onenew"
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

.news_article{
    margin: 15px;
    margin-left: 14vw;
    margin-top: 30px;
    font-size: 1.5rem;
}
</style>