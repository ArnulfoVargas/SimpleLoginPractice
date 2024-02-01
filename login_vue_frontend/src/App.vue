<script setup>
import { RouterLink, RouterView } from 'vue-router'
import { onBeforeMount, ref, watch } from 'vue';
import { useUser, useUrlBeforeRedirect } from './global/user_auth'
import router from './router';

const currentPath = window.location.pathname

const isAuth = ref(false)

onBeforeMount(() => {
  const isAuthPath = () => {
    if (currentPath === "/login" || currentPath === "/signin") return true
    return false
  }

  const user = useUser()
  if (user.isAuth) {
    isAuth.value = true
    
    if (isAuthPath()){
      router.push("/")
    }
    return
  }
  
  if (isAuthPath()) return

  useUrlBeforeRedirect().value = currentPath
  router.push("/login")

  console.log(router)
})

watch(useUser(), () => {
  isAuth.value = useUser().isAuth
}, {
  deep: true
})
</script>

<template>

  <div class="page" v-if="isAuth">
    <div>
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
      </nav>
    </div>
  </div>
  
  <RouterView />
</template>

<style>
body{
  margin: 0;
}
</style>
