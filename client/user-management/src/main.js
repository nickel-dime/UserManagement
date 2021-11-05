import {createApp} from 'vue';
import App from './App.vue';
// import UserItem from './components/UserItem.vue'
import router from './router.js';
import store from './store.js';


import PrimeVue from 'primevue/config';
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Card from 'primevue/card';
import Avatar from 'primevue/avatar';

import 'primevue/resources/themes/saga-blue/theme.css'
import 'primevue/resources/primevue.min.css'
import 'primeicons/primeicons.css'
import 'primeflex/primeflex.css';

const app = createApp(App);
app.use(PrimeVue);
app.use(router);
app.use(store)

app.component('InputText', InputText);
app.component('Button', Button);
app.component('DataTable', DataTable);
app.component('Column', Column);
app.component('Card', Card);
app.component('Avatar', Avatar);


app.mount('#app');