<template>
  <q-page class="q-pa-md" style="overflow-y: auto;height: calc(100vh - 50px)">
    <div class="row justify-center">

    <div  class="col-12 col-sm-6 col-md-4 col-lg-4 q-px-lg q-py-md">
      <q-card class="col-3 col-lg-3 q-pa-lg"  style="border-radius: 15px ">
        <div   class="card q-mb-none" style="width: 100%">
          <div class="flex justify-between q-mb-lg" >
            <div class="block text-bold font-medium q-mb-lg" style="text-decoration: none;user-select: none">
              {{t('numberOfCustomers')}}
            </div>

            <div
                style="width: 2.5rem; height: 2.5rem">
              <q-icon name="groups" class="text-grey-6"/>
            </div>
          </div>
          <small style="user-select: none;font-size: 14px" dir="ltr"
                 >{{ count?.customer }}</small>
        </div>
      </q-card>
    </div>
    <div  class="col-12 col-sm-6 col-md-4 col-lg-4 q-px-lg q-py-md">
      <q-card class="col-3 col-lg-3 q-pa-lg"  style="border-radius: 15px ">
        <div   class="card q-mb-none" style="width: 100%">
          <div class="flex justify-between q-mb-lg" >
            <div class="block text-bold font-medium q-mb-lg" style="text-decoration: none;user-select: none">
              {{t('numberOfProducts')}}
            </div>

            <div
                style="width: 2.5rem; height: 2.5rem">
              <q-icon name="inventory_2" class="text-grey-6"/>
            </div>
          </div>
          <small style="user-select: none;font-size: 14px" dir="ltr"
          >{{ count?.product }}</small>
        </div>
      </q-card>
    </div>
    <div  class="col-12 col-sm-6 col-md-4 col-lg-4 q-px-lg q-py-md">
      <q-card class="col-3 col-lg-3 q-pa-lg"  style="border-radius: 15px ">
        <div   class="card q-mb-none" style="width: 100%">
          <div class="flex justify-between q-mb-lg" >
            <div class="block text-bold font-medium q-mb-lg" style="text-decoration: none;user-select: none">
              {{t('numberOfAdmins')}}
            </div>

            <div
                style="width: 2.5rem; height: 2.5rem">
              <q-icon name="admin_panel_settings" class="text-grey-6"/>
            </div>
          </div>
          <small style="user-select: none;font-size: 14px" dir="ltr"
          >{{ count?.user }}</small>
        </div>
      </q-card>
    </div>

    <div class="col-12 col-sm-12 col-md-6 col-lg-6 q-px-lg q-py-md">
      <q-card class="col-3 col-lg-3 q-pa-lg"  style="border-radius: 15px ">
        <div   class="card q-mb-none" style="width: 100%">
          <div class="flex justify-between q-mb-lg" >
            <div class="block text-bold font-medium q-mb-lg" style="text-decoration: none;user-select: none">
              {{t('loginLog')}}
            </div>

            <div
                style="width: 2.5rem; height: 1rem">
              <q-icon name="login" class="text-grey-6"/>
            </div>

          </div>
          <div class="q-ma-md">
            <q-scroll-area style="height: 200px; width: 100%;" visible>
              <div v-for="log in loginLogs" :key="log.id" class="q-py-xs">
                {{ log.message }}
              </div>
            </q-scroll-area>
          </div>
        </div>
      </q-card>
    </div>

      <div class="col-12 col-sm-12 col-md-6 col-lg-6 q-px-lg q-py-md">
        <q-card class="col-3 col-lg-3 q-pa-lg"  style="border-radius: 15px ">
          <div   class="card q-mb-none" style="width: 100%">
            <div class="flex justify-between q-mb-lg" >
              <div class="block text-bold font-medium q-mb-lg" style="text-decoration: none;user-select: none">
                {{t('logoutLog')}}
              </div>

              <div
                  style="width: 2.5rem; height: 1rem">
                <q-icon name="logout" class="text-grey-6"/>
              </div>

            </div>
            <div class="q-ma-md">
              <q-scroll-area style="height: 200px; width: 100%;" visible>
                <div v-for="log in logoutLogs" :key="log.id" class="q-py-xs">
                  {{ log.message }}
                </div>
              </q-scroll-area>
            </div>
          </div>
        </q-card>
      </div>

    </div>

  </q-page>
</template>

<script setup>
import { onMounted, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useQuasar} from "quasar";

const $q = useQuasar()
const {locale,t} = useI18n()
const count = ref()
const logTitle = ref('create')
const logs = ref()
const loginLogs = ref()
const logoutLogs = ref()


onMounted(async () => {
  await getLoginLogs()
  await getLogoutLogs()
  await getCountRecord()
});

async function getLoginLogs() {
  let language
  if (locale.value === 'fa-IR') {
    language = 'fa'
  }

  if (locale.value === 'en-US') {
    language = 'en'
  }
  const response = await fetch(
      '/api/panel/logs/list-by-title',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              title: 'login',
              language: language
            }
        )
      }
  );

  if (!response.ok) {

    const data = await response.json();

    if (Array.isArray(data.error)) {
      data.error.forEach(err => {
        $q.notify({
          color: 'red',
          textColor: 'white',
          message: t(err.message),
          position:"top"
        });
      });
    } else if (typeof data.error === 'string') {

      $q.notify({
        color: 'red',
        textColor: 'white',
        message: t(data.error),
        position:"top"
      });
    }

    if (response.status === 403) {
      await router.push({name: 'login'})
    }

    throw new Error(`HTTP error! status: ${response.status}`);
  }

  loginLogs.value = await response.json();
}

async function getLogoutLogs() {
  let language
  if (locale.value === 'fa-IR') {
    language = 'fa'
  }

  if (locale.value === 'en-US') {
    language = 'en'
  }
  const response = await fetch(
      '/api/panel/logs/list-by-title',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              title: 'logout',
              language: language
            }
        )
      }
  );

  if (!response.ok) {

    const data = await response.json();

    if (Array.isArray(data.error)) {
      data.error.forEach(err => {
        $q.notify({
          color: 'red',
          textColor: 'white',
          message: t(err.message),
          position:"top"
        });
      });
    } else if (typeof data.error === 'string') {

      $q.notify({
        color: 'red',
        textColor: 'white',
        message: t(data.error),
        position:"top"
      });
    }

    if (response.status === 403) {
      await router.push({name: 'login'})
    }

    throw new Error(`HTTP error! status: ${response.status}`);
  }

  logoutLogs.value = await response.json();
}

async function getCountRecord() {

  const response = await fetch(
      '/api/panel/logs/count-records',
      {
        method: 'GET',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'}

      }
  );

  if (!response.ok) {

    const data = await response.json();

    if (Array.isArray(data.error)) {
      data.error.forEach(err => {
        $q.notify({
          color: 'red',
          textColor: 'white',
          message: t(err.message),
          position:"top"
        });
      });
    } else if (typeof data.error === 'string') {

      $q.notify({
        color: 'red',
        textColor: 'white',
        message: t(data.error),
        position:"top"
      });
    }

    if (response.status === 403) {
      await router.push({name: 'login'})
    }

    throw new Error(`HTTP error! status: ${response.status}`);
  }

  count.value = await response.json();
}

</script>
