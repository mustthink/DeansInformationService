import { createApp } from 'vue'
import App from './App.vue'
import router from '@/router'
import store from '@/store'
import components from '@/components/UI'
/*import Axios from 'axios'


Vue.prototype.$http = Axios;
const token = localStorage.getItem('token')
if (token) {
    Vue.prototype.$http.defaults.headers.common['Authorization'] = token
}
 */

const app = createApp(App)

components.forEach(
    component => {
        app.component(component.name, component)
    }
)

app.use(store).use(router).mount('#app')
