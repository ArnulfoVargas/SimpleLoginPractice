import { reactive, ref } from "vue";
const user = reactive({
    userData    : {},
    isAuth      : false,
})

const urlBeforeRedirect = ref("")

export function useUser() {
    return user
}

export function useUrlBeforeRedirect() {
    return urlBeforeRedirect
}