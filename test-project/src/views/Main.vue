<template>
<div>
    <navbar></navbar>
    <!-- <h1 style="margin:50px; font-size: 1.5rem;">Сайт работает :)</h1> -->
    <news-list v-bind:news="news"/>
</div>
</template>

<script>
import Navbar from "@/components/UI/Navbar.vue";
import NewsList from "@/components/NewsList.vue";
export default {
    components: { Navbar, NewsList},
    data(){
        return{
            news:[
                {id: 1, title: "Переменные", body: `Можно объявить при помощи:\n
                - let\n
                - const (константа, т.е. изменению не подлежит)\n
                - var (устаревший способ, подробности позже)\n
                Имя переменной может включать:\n
                Буквы и цифры, однако цифра не может быть первым символом.\n
                Символы $ и _ используются наряду с буквами.\n
                Иероглифы и символы нелатинского алфавита также допустимы, но обычно не используются.\n
                Переменные типизируются динамически. В них могут храниться любые значения:\n`},
                {id: 1, title: "Худшая неделя в жизни Илона Маска", body: `Его империи откровенно плохо, 
                а сам миллиардер переключается с одного скандала на другой. У Tesla медийно всё особенно плохо 
                — многолетнее восхищение стартапом резко меняется в критическую сторону. `},
                {id: 1, title: "Bloomberg считает", body: `что к 2024 году Tesla потеряет корону главного производителя 
                электромобилей. Их главный конкурент, Volkswagen, мощно наращивает продажи в Европе и Китае. 
                А Tesla застряла в США. `},
            ]
        }
    },
    mounted() {
        if (this.$store.state.isLogged == false) {
            this.$router.push("/login");
        }
    },
   
}
</script>

<style scoped>

</style>