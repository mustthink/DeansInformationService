<template>
<div>
    <navbar></navbar>
    <h1 style="margin:50px; font-size: 1.5rem;">Сайт работает :)</h1>
</div>
</template>

<script>
import Navbar from "@/components/UI/Navbar.vue";
export default {
     components: { Navbar },
    mounted() {
        if (this.$store.state.isLogged == false) {
            this.$router.push("/login");
        }
    },
   
}
</script>

<style scoped>

</style>