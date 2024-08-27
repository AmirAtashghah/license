import {route} from 'quasar/wrappers'
import {createMemoryHistory, createRouter, createWebHashHistory, createWebHistory} from 'vue-router'
import routes from './routes'

/*
 * If not building with SSR mode, you can
 * directly export the Router instantiation;
 *
 * The function below can be async too; either use
 * async/await or return a Promise which resolves
 * with the Router instance.
 */

export default route(function (/* { store, ssrContext } */) {
    // consts createHistory = process.env.SERVER
    //     ? createMemoryHistory
    //     : (process.env.VUE_ROUTER_MODE === 'history' ? createWebHistory : createWebHashHistory)

    const Router = createRouter({

        scrollBehavior: () => ({left: 0, top: 0}),
        routes,

        // Leave this as is and make changes in quasar.conf.js instead!
        // quasar.conf.js -> build -> vueRouterMode
        // quasar.conf.js -> build -> publicPath
        history: createWebHistory()
    })

    function getCookie(name) {
        const value = `; ${document.cookie}`;
        const parts = value.split(`; ${name}=`);
        if (parts.length === 2) return parts.pop().split(';').shift();
    }


    // //  validation token
    //     async function validateToken(token) {
    //
    //         const response = await fetch(
    //             '/api/token',
    //             {
    //                 method: 'POST',
    //                 headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
    //                 body: JSON.stringify({token: token})
    //             }
    //         );
    //         return response.ok;
    //     }


    Router.beforeEach( async (to, from, next) => {

        if (to.path !== '/login') {
            const token = getCookie('token');
            if (!token) {
                next({name: 'login'})
                return

            }
        }
        next()
    })

    return Router
})

