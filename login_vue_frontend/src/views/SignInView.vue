<script setup>
import AuthLeftSide from '@/components/AuthLeftSide.vue';
import WavesBackground from '@/components/WavesBackground.vue';
import InputsContainer from '@/components/InputsContainer.vue';
import InputField from '@/components/InputField.vue';
import { useUser, useUrlBeforeRedirect } from '@/global/user_auth';
import router from '@/router';
import { reactive } from 'vue';

const userData = reactive({
    username    : "",
    email       : "",
    password    : ""
})

const signin = () => {
    useUser().isAuth = true;
    const path = useUrlBeforeRedirect().value
    router.push(path ?? "/")
}

const clickBtn = () => {
    console.log(userData)
}
</script>

<template>
    <div class="container">
        <AuthLeftSide />

        <div class="right-side">
            <WavesBackground />

            <InputsContainer :title="'Create an Account'" :buttonText="'Sign In'" :url="'/login'" :text="'Already have an account?'"
                :urlText="'Login now!'" @clickBtn="clickBtn">
                <InputField :inputType="'text'" :placeholder="'Enter your name'" :imagePath="'/images/user.png'"
                    title="Username" v-model:data="userData.username"/>
                <InputField :inputType="'mail'" :placeholder="'Enter your email'" :imagePath="'/images/mail.png'"
                    title="Email" v-model:data="userData.email"/>
                <InputField :inputType="'password'" :placeholder="'Enter your password'" :imagePath="'/images/password.png'"
                    title="Password" v-model:data="userData.password"/>
            </InputsContainer>
        </div>
    </div>
</template>
  
<style scoped>
.container {
    width: 100dvw;
    height: 100dvh;
    display: grid;
    grid-template-columns: 1fr 1fr;
    background-color: #273134;
}</style>
  