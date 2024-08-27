

const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '/dashboard', name:'dashboard',component: () => import('pages/Dashboard.vue') },
      { path: '/logs',
        name:'logs',
        component: () => import('pages/Logs.vue')},

      { path: '/logout',
        name:'logout',
        component: () => import('components/Logout.vue')},

      { path: '/product',
        name:'product',
        component: () => import('pages/Product.vue'),
      },

      { path: '/customer',
        name:'customer',
        component: () => import('pages/Customer.vue')},

      { path: '/customersProduct',
        name:'customersProduct',
        component: () => import('pages/CustomersProduct.vue')},

      { path: '/user',
        name:'user',
        component: () => import('pages/User.vue')},

      { path: '/restriction',
        name:'restriction',
        component: () => import('pages/Restriction.vue')},
    ]
  },

  {
    path: '/login',
    name: 'login',
    component: () => import('pages/Login.vue')
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue')
  }
]

export default routes
