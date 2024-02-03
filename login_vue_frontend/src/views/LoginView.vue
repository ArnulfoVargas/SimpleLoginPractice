<script setup>
import AuthLeftSide from '@/components/AuthLeftSide.vue';
import WavesBackground from '@/components/WavesBackground.vue';
import InputsContainer from '@/components/InputsContainer.vue';
import InputField from '@/components/InputField.vue';
import { useUser, useUrlBeforeRedirect, useStoreData } from '@/global/user_auth';
import router from '@/router';
import { reactive, ref, watch } from 'vue';
import Alert from '@/components/Alert.vue';

const userData = reactive({
    email       : "",
    password    : ""
})

const errors = ref([])

watch(userData, () => {
  errors.value = []
})

const login = () => {
  useUser().isAuth = true;
  const path = useUrlBeforeRedirect().value
  router.push(path !== "/login" || path !== "/signin" || path !== "" ? path : "/")
}

const validate = () => {
    if (userData.email.trim() === "" || userData.password.trim() === ""){
        errors.value = ["All fields are required"]
    }
}

const clickBtn = async() => {
  validate()

  if (errors.value.length > 0) return

  const res = await fetch("http://localhost:3000/login", {
    method: "POST",
    body: JSON.stringify(userData)
  })
  const data = await res.json()
  
  if (res.status === 200){
      useStoreData({...data})
      login()
  } else {
      const serverErrors = {...data.errors}
      errors.value = serverErrors.msg.split(".").filter(s => s !== "")
  }
}
</script>

<template>
  <div class="container">
    <AuthLeftSide />

    <div class="right-side">
      <WavesBackground />

      <Alert :errors="errors"/>

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