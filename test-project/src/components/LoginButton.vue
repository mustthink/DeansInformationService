<template>
    <div>
    <my-button class='auth' @click="showDialog"><span>Авторизироваться</span></my-button>
    <my-dialog v-model:show="dialogVisible">
        <h1 class="dialog__auth">Авторизация</h1>
        <my-input
        v-model="user.login"
        type="text" 
        placeholder="Введите логин *"
        />
        <my-input
        v-model="user.password"
        type="text" 
        placeholder="Введите пароль *"
        />
        <p>* означает, что поле обязательно</p>
        <div v-if="valid" class="btns">
        <my-button @click="$router.push('/')" style="margin-top:15px;">
        Войти
        </my-button>
        <my-button @click="$router.push('/'), CreateUser()" style="margin-top:15px;">
        Зарегистрироваться
        </my-button>
        </div>
    </my-dialog>
    </div>
</template>

<script>
import axios from 'axios';
export default {
       data(){
        return{
            user:{
                login:'',
                password:'',
            },
            dialogVisible: false,
            valid: true,
        }
    },
    methods:{
        showDialog(){
            this.dialogVisible = true;
        },
        async CreateUser(){
            if(this.user.login.length > 0 && this.user.password.length > 0){
                const newUser = {
                id: Date.now(),
                login: this.user.login,
                password:this.user.password,
            }
            console.log(newUser)
            // let result = await axios.post('', newUser);
            let result = 201;
            if (result == 201){
                this.$store.commit('setLogged');
            }
            }
            else{
                alert('Пожалуйста, заполните пустые поля')
            }
            

        },
        ValidateUser(){

        },
        
    },
    mounted(){
        if(this.$store.state.isLogged){
            this.$router.push('/')
        }
    },
}
</script>

<style scoped>
* {
    -webkit-box-sizing: border-box;    
    -moz-box-sizing: border-box;    
    box-sizing: border-box;
}

.buttons {
    margin: 10%;    
    text-align: center;
}

.auth {
    width: 200px;    
    font-size: 16px;    
    font-weight: 600;    
    color: #fff;    
    cursor: pointer;    
    margin: 20px;    
    height: 55px;    
    text-align:center;    
    border: none;    
    background-size: 300% 100%; 
    border-radius: 50px;    
    -moz-transition: all .4s ease-in-out;    
    -o-transition: all .4s ease-in-out;    
    -webkit-transition: all .4s ease-in-out;    
    transition: all .4s ease-in-out;
}

.auth:hover {
    background-position: 100% 0;    
    -moz-transition: all .4s ease-in-out;    
    -o-transition: all .4s ease-in-out;    
    -webkit-transition: all .4s ease-in-out;    
    transition: all .4s ease-in-out;
}

.auth:focus {
    outline: none;
}

.auth{
    display: flex;
    justify-content: center;
    align-items: center;
    /* background-image: linear-gradient(to right, #0ba360, #3cba92, #30dd8a, #2bb673);     */
    box-shadow: 0 4px 15px 0 rgba(23, 168, 108, 0.75);
}

span{
    font-family: 'Source Sans Pro', sans-serif;
    font-weight: 300;
    font-size: 1rem;
}

h1{
    font-family: 'Source Sans Pro', sans-serif;
    font-weight: 700;
    display: flex;
    align-items: center;
    justify-content: center;
}

.btns{
    display: flex;
    justify-content: space-evenly;
}

.dialog__auth{
    font-family: 'Source Sans Pro', sans-serif;
    font-weight: 700;
    font-size: 1.5rem;
}

p{
    font-family: 'Source Sans Pro', sans-serif;
    font-weight: 300;
    font-size: 0.7rem;
    margin-top: 5px;
    color: grey;
}

</style>