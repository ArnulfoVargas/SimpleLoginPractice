<script setup>
import { defineProps, defineEmits } from 'vue';
import router from '../router';

const props = defineProps(["title", "buttonText", "url", "text", "urlText"])
defineEmits(["clickBtn"])
const onLinkClick = () => {
    router.push(props.url)
}
</script>

<template>
    <div class="input-container">
        <div class="input-bg">

            <div class="top">
                <h1>{{ title }}</h1>
            </div>

            <div class="inputs">
                <slot>
                </slot>
            </div>

            <div class="button">
                <button @click="$emit('clickBtn')">{{ buttonText }}</button>
            </div>

            <p>
               {{text}} <span @click="onLinkClick">{{urlText}}</span>
            </p>
        </div>
    </div>
</template>

<style scoped>
.input-container {
    position: absolute;
    top: 0;
    right: 0;
    width: 50dvw;
    height: 100dvh;

    display: flex;
    justify-content: center;
    align-items: center;
}

.input-bg {
    background-color: #F1F1F1;
    height: fit-content;
    width: calc(60% - 100px);
    border-radius: 35px;
    box-shadow: 50px 50px 100px #9d9d9d36,
        -50px -50px 100px #9d9d9d36;

    padding: 50px;
    padding-bottom: calc(50px - 1.5rem);

    display: flex;
    align-items: center;
    justify-content: space-between;
    flex-direction: column;
}

h1 {
    margin: 0;
    font-family: Arial, Helvetica, sans-serif;
    font-weight: 100;
    font-style: italic;
}

.button {
    width: 100%;
    height: 45px;
}

button {
    background-image: linear-gradient(to right, #4CB8C4 0%, #3CD3AD 51%, #4CB8C4 100%);
    text-align: center;
    transition: 0.5s;
    background-size: 200% auto;
    color: white;
    box-shadow: 0 0 20px #eee;
    border-radius: 22.5px;
    display: block;
    font-family: Arial, Helvetica, sans-serif;
    font-size: 1.75rem;
    font-style: italic;
    width: 100%;
    border: none;
    height: 100%;
    box-shadow: 0px 10px 15px -3px rgba(0, 0, 0, 0.1);
}

button:hover {
    background-position: right center;
    color: #fff;
    box-shadow: 0px 5px 7px 0px rgba(0, 0, 0, 0.1);
}

.inputs{
    width: 100%;
    display: flex;
    flex-direction: column;
    gap: 1rem;
    margin-bottom: 2rem;
}

p{
    margin-top: 1.5rem;
    font-size: 1rem;
    font-family: Arial, Helvetica, sans-serif;
    color: #5a5a5a;

    span{
        cursor: pointer;
        font-weight: bold;
        color: #4CB8C4;
    }
}
</style>