<template>
  <q-layout view="lHh Lpr lFf">

    <q-drawer
      v-model="drawer"
      show-if-above
      :width="200"
      :breakpoint="500"
      bordered
      :class="$q.lang.rtl ? 'bg-grey-9' : 'bg-grey-3'"
    >
      <q-scroll-area class="fit">
        <q-list>

          <template v-for="(menuItem, index) in menuList" :key="index">
            <q-item clickable :to="menuItem.link"  v-ripple>
              <q-item-section avatar>
                <q-icon :name="menuItem.icon" />
              </q-item-section>
              <q-item-section>
                {{ menuItem.label }}
              </q-item-section>
            </q-item>
            <q-separator :key="'sep' + index"  v-if="menuItem.separator" />
          </template>

        </q-list>
      </q-scroll-area>
    </q-drawer>


    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup>

import {ref} from "vue";

const drawer = ref(false)

const menuList = [
  {
    icon: 'add',
    label: 'Add new client',
    separator: false,
    link: 'AddNewClient'
  },
  {
    icon: 'dashboard',
    label: 'Dashboard',
    separator: false,
    link: 'Dashboard'
  },
  {
    icon: 'settings',
    label: 'Settings',
    separator: false,
    link: 'Setting'
  },
  {
    icon: 'logout',
    label: 'Logout',
    separator: false,
    link: 'Logout'

  },
  {
    icon: 'help',
    iconColor: 'primary',
    label: 'Help',
    separator: false
  }
]

</script>
