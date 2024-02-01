<script setup>

</script>

<template>
    <div>
        <img class="purple" src="/images/wave3.png">
        <img class="orange" src="/images/wave1.png">
        <img class="green" src="/images/wave2.png">
    </div>
</template>

<style scoped>
    div{
        position: relative;
        bottom: 0;
        height: 100%;
        width: 100%;
        overflow: hidden;
    }
    img{
        width: 150%;
        position: absolute;
        opacity: 60%;

        animation-fill-mode: forwards;
        animation-timing-function: ease-in-out;
        animation-iteration-count: infinite;
    }
    .orange{
        right: -50%;
        animation-name: orange-anim-x, orange-anim-y;
        animation-duration: 60s, 15s;
        bottom: -80px;
    }
    .green{
        right: 0%;
        animation-name: green-anim-x, green-anim-y;
        animation-duration: 50s, 20s;
        bottom: -100px;
    }

    .purple {
        animation-name: purple-anim-x, purple-anim-y;
        animation-duration: 80s, 100s;
        animation-timing-function: linear;
        bottom: -150px;
    }

    @keyframes orange-anim-x {
        0%{
            transform: translate(0%, 0%);
        }
        50%{
            transform: translate(-33%,0%);
        }
        100%{
            transform: translate(0%, 0%);
        }
    }

    @keyframes orange-anim-y {
        0%{
            bottom: -80px;
        }
        50%{
            bottom: -60px;
        }
        100%{
            bottom: -80px;
        }
    }

    @keyframes green-anim-x {
        0%{
            right: 0%;
        }
        50%{
            right: -50%;
        }
        100%{
            right: 0%;
        }
    }

    @keyframes green-anim-y {
        0%{
            bottom: -120px;
        }
        50%{
            bottom: -100px;
        }
        100%{
            bottom: -120px;
        }
    }

    @keyframes purple-anim-x {
        0%{
            transform: translate(-16%, 0%);
        }
        25%{
            transform: translate(-33%,0%);
        }
        50% {
            transform: translate(-16%, 0%);
        }
        66%{
            transform: translate(0%, 0%);
        }
        100%{
            transform: translate(-16%, 0%);
        }
    }

    @keyframes purple-anim-y {
        0%{
            bottom: -150px;
        }
        50%{
            bottom: -170px;
        }
        100%{
            bottom: -150px;
        }
    }
</style>