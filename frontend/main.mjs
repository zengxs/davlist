import Vue from 'vue';
import {BootstrapVue} from 'bootstrap-vue';
import App from './App.vue';
import 'bootstrap-vue/dist/bootstrap-vue.min.css';
import 'admin-lte/dist/css/adminlte.min.css';

Vue.use(BootstrapVue);

new Vue({
  el: '#app',
  render: h => h(App)
});
