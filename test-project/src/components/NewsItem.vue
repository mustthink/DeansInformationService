<template>
    <div class="news">
      <div class="title">{{ onenew.title }}</div>
      <div>{{ onenew.body }}</div>
    </div>
</template>

<script>
    export default {
        props: {
          onenew: {
            type: Object,
            required: true,
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

.title{
   font-weight: 100; 
   font-family: 'Roboto', sans-serif;
   /* font-family:system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif; */
   font-size:  1.2rem;
   color: teal
}
</style>