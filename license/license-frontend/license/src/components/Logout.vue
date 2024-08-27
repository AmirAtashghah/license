<template>
</template>

<script setup>

import {onMounted} from "vue";
import { useRouter } from 'vue-router'
import {useQuasar} from "quasar";
import {useI18n} from "vue-i18n";

const router = useRouter()
const $q = useQuasar()
const {t} = useI18n()

onMounted(async () => {
  await logout()
});

async function logout() {

  const response = await fetch(

      '/api/panel/logout',
      {
        method: 'GET',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
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

    await router.push({ name: 'login'})

    throw new Error(`HTTP error! status: ${response.status}`);
  }

  await router.push({ name: 'login'})
}

</script>
