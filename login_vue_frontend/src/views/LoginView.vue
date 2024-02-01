<script setup>
import AuthLeftSide from '@/components/AuthLeftSide.vue';
import WavesBackground from '@/components/WavesBackground.vue';
import InputsContainer from '@/components/InputsContainer.vue';
import InputField from '@/components/InputField.vue';
import { useUser, useUrlBeforeRedirect } from '@/global/user_auth';
import router from '@/router';
import { reactive } from 'vue';

const userData = reactive({
    email       : "",
    password    : ""
})

const signin = () => {
  useUser().isAuth = true;
  const path = useUrlBeforeRedirect().value
  router.push(path ?? "/")
}

const clickBtn = async() => {
  const res = await fetch("http://localhost:3000/api/test", {
    method: "POST",
    body: JSON.stringify(userData)
  })
  const data = await res.json()
  console.log(data)
}
</script>

<template>
  <div class="container">
    <AuthLeftSide />

    <div class="right-side">
      <WavesBackground />

      <InputsContainer :title="'Login'" :buttonText="'Login'" :url="'/signin'"
        :text="'Dont have an account?'" :urlText="'Create one!'" @clickBtn="clickBtn">
        <InputField :inputType="'mail'" :placeholder="'Enter your email'" :imagePath="'/images/mail.png'" title="Email" 
          v-model:data="userData.email"/>
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
}
</style>