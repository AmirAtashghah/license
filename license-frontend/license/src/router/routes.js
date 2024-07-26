const routes = [
  {
    path: '/client',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '/dashboard', name:'dashboard',component: () => import('pages/Dashboard.vue') },
      { path: '/addNewClient', name:'addNewClient',component: () => import('pages/AddNewClient.vue') },
      { path: '/clientLog', name:'clientLog',component: () => import('components/ClientLog.vue') },
      { path: '/setting', name:'setting',component: () => import('pages/Setting.vue') },
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
