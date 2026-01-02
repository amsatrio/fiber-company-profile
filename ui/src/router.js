import { createRouter, createWebHistory } from 'vue-router';
import Home from './pages/Home.vue';
import Counter from './pages/Counter.vue';
import NotFound from './pages/errors/NotFound.vue';
import CompanyProfile1 from './pages/CompanyProfile1.vue';
import CompanyProfile2 from './pages/CompanyProfile2.vue';

const routes = [
  { path: '/', name: 'Home', component: Home },
  { path: '/counter', name: 'Counter', component: Counter },
  { path: '/company-profile-1', name: 'company-profile-1', component: CompanyProfile1 },
  { path: '/company-profile-2', name: 'company-profile-2', component: CompanyProfile2 },
  { path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound },
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;