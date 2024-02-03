import router from "@/router";
import { reactive, ref } from "vue";
const user = reactive({
    userData    : {},
    isAuth      : false,
})

const urlBeforeRedirect = ref("")

export function useUser() {
    return user
}

export async function useValidate() {
    const token = localStorage.getItem("token")


    if (token) {
        const res = await fetch("http://localhost:3000/jwt", {
            method: "POST",
            headers: {"Authorization": "bearer " + token},
            body: JSON.stringify({"id": +localStorage.getItem("id")})
        })

        if (!res) return

        if (res.status === 401){
            user.isAuth = false
            localStorage.clear()
            return
        }

        const data = await res.json()
        if (res.status === 200){
            user.userData = {...data}
            user.isAuth = true
            router.push(urlBeforeRedirect.value)
            return
        } else {
            user.isAuth = false
        }
    }

    user.isAuth = false
    return
}

export function useUrlBeforeRedirect() {
    return urlBeforeRedirect
}

export function useStoreData(obj = {}){
  if (obj.token !== ""){
    localStorage.setItem("token", obj.token)
  }
  if (obj.user.id != 0){
    localStorage.setItem("id", obj.user.id)
  }
}