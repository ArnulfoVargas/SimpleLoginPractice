<script setup>
 const logout = () => {
  if (confirm("Are you secure?")){
      localStorage.clear()
      location.reload()
    }
 }
</script>

<template>
  <div class="home">
    <h1>This is a home page</h1>
    <button @click="logout">Log out</button>
  </div>
</template>

<style>
</style>
