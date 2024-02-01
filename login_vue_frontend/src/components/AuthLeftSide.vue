<template>
    <div class="left-side">
        <img class="bg" src="https://libertex.org/sites/default/files/inline-images/mercado-divisas-1.jpg" alt="bg">

        <div>
            <img class="logo" src="https://www.freeiconspng.com/thumbs/retail-store-icon/retail-store-icon-18.png"
                alt="logo">
            <h2 class="title">Test Shop</h2>
        </div>
    </div>
</template>

<style scoped>
.left-side {
    position: relative;
    object-fit: cover;
    overflow: hidden;

    .bg {
        position: relative;
        height: 150%;
        right: 0;
        top: 0;
        filter: blur(5px);
        opacity: 50%;
        animation-name: left-side-movement;
        animation-duration: 60s;
        animation-fill-mode: forwards;
        animation-timing-function: linear;
        animation-iteration-count: infinite;
    }

    div {
        position: absolute;
        z-index: 1;
        display: flex;
        flex-direction: column;
        top: 0;
        right: 0;
        height: 100%;
        width: 100%;
        justify-content: center;
        align-items: center;

        .title {
            color: white;
            font-family: 'Franklin Gothic Medium', 'Arial Narrow', Arial, sans-serif;
            transform: translate(0, -100%);
            font-size: 2rem;
        }
    }
}

@keyframes left-side-movement {
    0% {
        transform: translate(0%, -30%);
    }

    25% {
        transform: translate(-30%, -30%);
    }

    50% {
        transform: translate(-30%, 0%);
    }

    75% {
        transform: translate(0%, 0%);
    }

    100% {
        transform: translate(0%, -30%);
    }
}

.logo {
    width: 30%;
    filter: invert();
}</style>