<template>
    <div>
    <h2 class="news_article">Последние новости:</h2>
    <div class="news" v-for="thing in news">
        <div>{{ thing.title }}</div>
        <div>{{ thing.body }}</div>
    </div>
  </div>
</template>

<script>
export default {
    props: {
        news: {
            type: Array,
            required: true
        }
  }
}
</script>

<style scoped>
.news{
    padding: 15px;
    border: 2px solid teal;
    margin-top: 15px;
    margin-left: 15px;
    display: grid;
    width: 40vw;
}

.news_article{
    margin: 15px;
    margin-left: 14vw;
    margin-top: 30px;
    font-size: 1.5rem;
}
</style>