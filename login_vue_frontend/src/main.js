import { createApp } from 'vue'
import App from './App.vue'
import { useValidate, useUrlBeforeRedirect, useUser } from './global/user_auth'
import router from './router'

const app = createApp(App)

useUrlBeforeRedirect()
await useValidate()

app.use(router)

app.mount('#app')
