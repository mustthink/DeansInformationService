<template>
    <div>
        <div>
            <pre id="myNews" ref="myNews" ></pre>
        </div>
    </div>
</template>

<script>
    export default {
        mounted(){
            this.$refs.myNews.innerText = JSON.stringify(this.$store.state.news);
        },
  }
  
</script>

<style scoped>

</style>