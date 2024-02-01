import { createApp } from 'vue'
import App from './App.vue'
import { useUser, useUrlBeforeRedirect } from './global/user_auth'
import router from './router'

const app = createApp(App)

useUser()
useUrlBeforeRedirect()

app.use(router)

app.mount('#app')
