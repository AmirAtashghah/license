const routes = [
  {
    path: '/client',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '/Dashboard', name:'Dashboard',component: () => import('pages/Dashboard.vue') },
      { path: '/AddNewClient', name:'AddNewClient',component: () => import('pages/AddNewClient.vue') },
      { path: '/ClientDetail', name:'ClientDetail',component: () => import('pages/ClientDetail.vue') },
      { path: '/ClientLog', name:'ClientLog',component: () => import('pages/ClientLog.vue') },
      { path: '/Setting', name:'Setting',component: () => import('pages/Setting.vue') },
      { path: '/UpdateClient', name:'UpdateClient',component: () => import('pages/UpdateClient.vue') }
    ]
  },

  {
    path: '/Login',
    name: 'Login',
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
