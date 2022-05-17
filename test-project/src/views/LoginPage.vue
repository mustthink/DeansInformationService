<template>
    <div>
        
        <div class="hero">
            <!-- <pre class="ascii-art"></pre> -->
            <!-- <img class='yoba' src="@/assets/yoba.gif" alt="yoba"> -->
            <img class='nure' src="@/assets/nure.png" alt="yoba">
            <h1>Вы зашли на сайт деканата</h1>
            <h2>и вы знаете что нужно сделать ( ͡° ͜ʖ ͡°)</h2>
            <login-button class="login_btn"/>
        </div>
    </div>
</template>

<script>
import LoginButton from '@/components/LoginButton'
export default {
   components:{
       LoginButton
   }     
}
</script>

<style scoped>

*{
    margin: 0;
    padding: 0;
}

.hero{
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-top: 20vh;
}

h1{
    font-family: 'Source Sans Pro', sans-serif;
    font-weight: 900;
    margin-top: 40vh;
    font-size: 2.3rem;
    color: teal;
}

h2{
    font-family: 'Source Sans Pro', sans-serif;
    font-weight: 300;
    font-size: 1.2rem;
    margin-bottom: 15px;
}

.yoba{
    margin-bottom: -300px;
    margin-top: 5vh;
    width: 40vh;
}

.nure{
    width: 30vh;
    margin-bottom: -35vh;
}

.ascii-art {
    color: teal;
    font-family: monospace;
    white-space: pre;
}
</style>