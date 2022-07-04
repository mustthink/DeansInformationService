<template>
    <div class="news">
      <div v-if="this.$store.state.isDean" class="id">Id: {{onenew.id}}</div>
      <div class="title">{{ onenew.Title }}</div>
      <div>{{ onenew.Content }}</div>
      <div v-if="this.$store.state.isDean" class="btn">
        <my-button @click="remove">Удалить</my-button>
      </div>
    </div>
</template>

<script>
    export default {
        props: {
          onenew: {
            type: Object,
            required: true,
          }
        },
        methods:{
          remove(){
            (async () => {
              // DELETE request using axios with async/await
              await axios.delete(`https://reqres.in/api/posts/${onenew.id}`);
            })();
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

.btn{
  margin-top: 6px;
}

/* .btn:hover{
  margin-top: 6px;
  color:aquamarine
} */

.title{
   font-weight: 100; 
   font-family: 'Roboto', sans-serif;
   /* font-family:system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif; */
   font-size:  1.2rem;
   color: teal
}
</style>