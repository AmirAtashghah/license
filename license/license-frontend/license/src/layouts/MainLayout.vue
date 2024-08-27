<template>

  <q-layout view="lHh Lpr rff" >
    <q-header elevated style="background-color: black" >
      <q-toolbar>
        <q-btn flat @click="drawer = !drawer" round dense icon="menu" />
        <q-toolbar-title></q-toolbar-title>

        <q-btn icon="language" flat dense rounded size="md"
               fab-mini>
          <q-tooltip>{{ $t('language') }}</q-tooltip>
          <q-menu auto-close>
            <q-list>
              <q-item clickable v-for="lang in localeOptions" :key="lang.value" v-close-popup
                      @click="toggleLanguage(lang.value)">
                <q-item-section class="q-mx-sm text-uppercase" :class="currentLang.value===lang.value&&'text-bold'">
                  {{ lang.label }}
                </q-item-section>
              </q-item>
            </q-list>
          </q-menu>
        </q-btn>

        <q-btn id="toggle_theme" v-if="$q.dark.isActive" flat round icon="dark_mode"  @click="changeThemeMod"
               size="md">
          <q-tooltip style="width: 7em">{{ $t('lightMode') }}</q-tooltip>
        </q-btn>
        <q-btn id="toggle_theme" v-else flat round  icon="brightness_7" @click="changeThemeMod">
          <q-tooltip>{{ $t('darkMode') }}</q-tooltip>
        </q-btn>
      </q-toolbar>


    </q-header>

    <q-drawer
      v-model="drawer"
      show-if-above
      :width="210"
      breakpoint="500"
      bordered
    >
      <q-scroll-area style="height: calc(100% - 150px); margin-top: 150px; border-right: 1px solid #ddd">

        <q-list >

          <template v-for="(menuItem, index) in menuList" :key="index">
            <q-item  clickable :to="menuItem.link"  v-ripple>
              <q-item-section avatar>
                <q-icon :name="menuItem.icon" />
              </q-item-section>
              <q-item-section>
                {{ $t(menuItem.label) }}
              </q-item-section>
            </q-item>
            <q-separator :key="'sep' + index"  v-if="menuItem.separator" />
          </template>

        </q-list>
      </q-scroll-area>

      <q-img class="absolute-top"  style="height: 150px;background-color: black">
        <div class="absolute-bottom bg-transparent">
          <q-avatar size="56px" class="q-mb-sm">
            <img src="../assets/user.png">
          </q-avatar>
          <div class="text-weight-bold">{{username}}</div>
          <div>{{username}}</div>
        </div>
      </q-img>

    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>

</template>

<script setup>

import {computed, onMounted, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useQuasar} from "quasar";

const $q = useQuasar()
const drawer = ref(false)
const themeMod = ref(false)
const username = ref()
const {locale,t} = useI18n()
const currentLang = ref()
const isFarsi = ref(true);
const localeOptions = [
  {
    label: t('fa'),
    value: 'fa-IR'
  },
  {
    label: t('en'),
    value: 'en-US'
  },
]
const menuList = [
  {
    icon: 'dashboard',
    label: 'dashboard',
    separator: false,
    link: 'dashboard',
  },
  {
    icon: 'inventory_2',
    label: 'products',
    separator: false,
    link: 'product',

  },
  {
    icon: 'groups',
    label: 'customers',
    separator: false,
    link: 'customer',

  },
  {
    icon: 'enhanced_encryption',
    label: 'restriction',
    separator: false,
    link: 'restriction'
  },
  {
    icon: 'admin_panel_settings',
    label: 'admins',
    separator: false,
    link: 'user',
  },
  {
    icon: 'dns',
    label: 'customersProduct',
    separator: false,
    link: 'customersProduct',
  },
  {
    icon: 'document_scanner',
    label: 'logs',
    separator: false,
    link: 'logs',

  },
  {
    icon: 'logout',
    label: 'logout',
    separator: false,
    link: 'logout',


  },
]

username.value = sessionStorage.getItem('username')
currentLang.value = locale.value

onMounted(async () => {
  locale.value = 'fa-IR';
  $q.lang.set({ rtl: true });
});

function changeThemeMod(){
  themeMod.value = !themeMod.value
  $q.dark.set(themeMod.value)
}

const toggleLanguage = (lang) => {
  if (lang==='en-US') {
    locale.value = 'en-US';
    $q.lang.set({ rtl: false });
    sessionStorage.setItem("lang",'fa')
  } else {
    locale.value = 'fa-IR';
    $q.lang.set({ rtl: true });
    sessionStorage.setItem("lang",'en')

  }
  isFarsi.value = !isFarsi.value;
};

</script>
