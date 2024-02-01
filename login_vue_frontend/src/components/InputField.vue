<script setup>
import { defineProps, defineEmits } from 'vue';
defineProps(["inputType", "placeholder", "imagePath", "title"])
defineEmits(["update:data"])
</script>

<template>
    <div id="cont">
        <p>{{ title }}:</p>
        <div class="input">
            <div class="img-container">
                <img id="leading" :src="imagePath" alt="">
            </div>

            <input :type="inputType" :placeholder="placeholder" @input="$emit('update:data', $event.target.value)">
        </div>
    </div>
</template>

<style scoped>
div {
    width: 100%;
}

p {
    margin: 0;
    margin-bottom: .35rem;
    font-family: Arial, Helvetica, sans-serif;
    font-style: italic;
    font-size: 1.2rem;
    font-weight: 500;

}

.input {
    width: calc(100% - 20px);
    height: 45px;
    max-height: 45px;
    border-radius: 25px;
    border: none;
    background-color: #D9D9D9;
    /* box-shadow: inset 4px 15px 25px -6px rgba(0, 0, 0, 0.1); */
    padding-left: 5px;
    padding-right: 15px;
    display: grid;
    grid-template-columns: 45px auto;

    .img-container {
        height: 100%;
        width: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    input {
        background-color: transparent;
        border: none;
        font-size: 1.35rem;
        color: #5a5a5a;
        font-family: 'Franklin Gothic Medium', 'Arial Narrow', Arial, sans-serif;
        
    }

    input:focus {
        outline: none;
    }
}

#leading {
    width: 25px;
    filter: invert(36%) sepia(26%) saturate(0%) hue-rotate(183deg) brightness(98%) contrast(93%);
    transition-property: filter;
    transition-duration: .5s;
}

#cont:has(input:focus) #leading {
    filter: invert(59%) sepia(65%) saturate(350%) hue-rotate(138deg) brightness(96%) contrast(91%);
}
</style>