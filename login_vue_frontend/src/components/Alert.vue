<script setup>
import {computed, defineProps} from 'vue'
const props = defineProps(["errors"])
const canShow = computed(() => {
  if (props.errors.length > 0) return "show"
  return ""
})
</script>

<template>
    <div class="alert-container" :class="canShow">
        <p :class="canShow" v-for="(err, i ) in errors" :key="i">{{ err }}</p>
    </div>
</template>

<style scoped>
p{
    color: transparent;
    margin: 5px;
    font-size: 1rem;
    font-family: 'Franklin Gothic Medium', 'Arial Narrow', Arial, sans-serif;
    transition-property: color;
    transition-delay: 1s;
}
p.show {
    color: white;
}
.alert-container {
    background-color: red;
    width: 30dvw;
    height: 0;
    position: absolute;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    top: 1dvh;
    right: 10dvw;
    border-radius: 35px;
    transition: height .1s linear;
}

.alert-container.show {
    height: 4rem;
    transition: height .2s linear;
}
</style>