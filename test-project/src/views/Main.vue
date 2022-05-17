<template>
<div>
    <h1>Hello this is main page</h1>
</div>
</template>

<script>
export default {
mounted(){
    if(this.$store.state.isLogged == false){
        this.$router.push('/login')
    }
}      
}
</script>

<style scoped>

</style>