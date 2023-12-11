import Vue from 'vue'
import Router from 'vue-router'
/* Layout */
import Layout from '@/layout'

Vue.use(Router)

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [{
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
}, {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
}, {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/dashboard/index'),
        meta: {title: '首页', icon: 'dashboard'}
    }]
}, {
    path: '/basic',
    component: Layout,
    redirect: '/basic/collector',
    meta: {title: '基础数据', icon: 'el-icon-menu'},
    children: [
        {
            path: 'netWork',
            name: 'netWork',
            component: () => import('@/views/basic/net-work/index'),
            meta: {title: '网络接口', icon: 'el-icon-help'},
            hidden: true,
        },
        {
            path: 'collector',
            name: 'collector',
            component: () => import('@/views/basic/collector/index'),
            meta: {title: '采集接口', icon: 'el-icon-share'}
        }, {
            path: 'deviceType',
            name: 'deviceType',
            component: () => import('@/views/basic/device-type/index'),
            meta: {title: '设备类型', icon: 'el-icon-cpu'}
        }, {
            path: 'device',
            name: 'device',
            component: () => import('@/views/basic/device/index'),
            meta: {title: '设备列表', icon: 'el-icon-s-tools'}
        }]
}, {
    path: '/schedule',
    component: Layout,
    redirect: '/schedule/collect',
    meta: {title: '任务计划', icon: 'el-icon-stopwatch'},
    children: [{
        path: 'collect',
        name: 'collect',
        component: () => import('@/views/schedule/collect/index'),
        meta: {title: '采集任务', icon: 'el-icon-s-promotion'}
    }, {
        path: 'report',
        name: 'report',
        component: () => import('@/views/schedule/report/index'),
        meta: {title: '上报任务', icon: 'el-icon-upload'}
    }]
}, {
    path: '/debug',
    component: Layout,
    meta: {title: '运维服务', icon: 'form'},
    children: [{
        path: 'serial',
        name: 'Serial',
        component: () => import('@/views/debug/serial/index'),
        meta: {title: '串口调试', icon: 'el-icon-link'},
        hidden: true,
    }, {
        path: 'net',
        name: 'net',
        component: () => import('@/views/debug/net/index'),
        meta: {title: '系统升级', icon: 'el-icon-connection'}
    }, {
        path: 'operation',
        name: 'Operation',
        component: () => import('@/views/debug/operation/index'),
        meta: {title: '操作说明', icon: 'el-icon-data-line'},
        // hidden: true,
    }, {
        path: 'log',
        name: 'Log',
        component: () => import('@/views/debug/log/index'),
        meta: {title: '更新日志', icon: 'el-icon-document'},
        // hidden: true,
    }]
}, {
    path: '*',
    redirect: '/404',
    hidden: true
}]

const createRouter = () => new Router({
    // mode: 'history', // require service support
    scrollBehavior: () => ({y: 0}),
    routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
    const newRouter = createRouter()
    router.matcher = newRouter.matcher // reset router
}

export default router
