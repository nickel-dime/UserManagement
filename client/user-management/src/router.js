import { createRouter, createWebHistory } from "vue-router";

import UserDetail from "./pages/UserDetail.vue";
import UserList from "./pages/UserList.vue";
import UserRegistration from "./pages/UserRegistration.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", redirect: "/users" },
    { path: "/users", component: UserList },
    {
      path: "/users/:id",
      name: "users",
      component: UserDetail,
      props: true
    },
    { path: "/register", component: UserRegistration },
    { path: "/:notFound(.*)", component: UserList },
  ],
});

export default router;
