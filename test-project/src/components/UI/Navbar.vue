<template>
    <div class="navbar">
        <div class="navbar__text" @click="$router.push('/')">.dekkanat</div>
        <div class="navbar__btns">
            <div v-if="isStudent">
            <button class='navbar__btn' @click="$router.push('/login')" style="margin-left: 20px">Login Page</button>
            <button class='navbar__btn' @click="$router.push('/login')" style="margin-left: 20px">Получение справок</button>
            <button class='navbar__btn' @click="$router.push('/login')" style="margin-left: 20px">Преподаватели</button>
            </div>
            <div v-if="this.$store.state.isEducator">
            <button class='navbar__btn' @click="$router.push('/login')" style="margin-left: 20px">Login Page</button>
            <button class='navbar__btn' @click="$router.push('/login')" style="margin-left: 20px">Расписание</button>
            <button class='navbar__btn' @click="$router.push('/login')" style="margin-left: 20px">Студенты</button>
            </div>
            <div v-if="isDean">
            <button class='navbar__btn' @click="$router.push('/login')" style="margin-left: 20px">Login Page</button>
            <button class='navbar__btn' @click="$router.push('/login')" style="margin-left: 20px">Просмотр заявок</button>
            <button class='navbar__btn' @click="$router.push('/login')" style="margin-left: 20px">Преподаватели</button>
            <button class='navbar__btn' @click="$router.push('/login')" style="margin-left: 20px">Группы</button>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data(){
    return{

    } 
    }
}
</script>

<style scoped>

@import url('https://fonts.googleapis.com/css2?family=Roboto:ital,wght@1,700&display=swap');

.navbar{
    height: 70px;
    background-color:rgba(23, 168, 108, 0.75);;
    box-shadow: 0 4px 15px 0 rgba(23, 168, 108, 0.75);
    display: flex;
    align-items: center;
    padding: 0 15px;
    z-index: 999;
}

.navbar__btns{
    margin-left: auto;
   
}

.navbar__text{
    font-family: 'Roboto', sans-serif;
    font-weight: 700;
    font-size: 1.5rem;
    color: white;
    cursor: pointer;
}

.navbar__btn{
    padding: 10px 15px;
    color: teal;
    border: none;
    cursor: pointer;
    border-radius: 50px;
    background-color: none;
    box-shadow: 0 4px 15px 0 rgba(0, 128, 128, 0.75);
    color: teal;
}

.navbar__btn:hover{
    box-shadow: 0 4px 15px 0 rgba(23, 168, 108, 0.75);
}
</style>