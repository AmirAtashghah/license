<template>


    <q-card-section>
      <q-list bordered padding class="rounded-borders">
        <q-item>
          <q-item-section>
            <q-item-label><strong>{{$t('id')}}</strong> </q-item-label>
            <q-item-label  caption>{{ admin?.id }}</q-item-label>
          </q-item-section>
        </q-item>
        <q-separator />

        <q-item >
          <q-item-section>
            <q-item-label><strong>{{$t('name')}}</strong></q-item-label>
            <q-item-label caption>{{ admin?.name }}</q-item-label>
          </q-item-section>
        </q-item>
        <q-separator />

        <q-item >
          <q-item-section>
            <q-item-label><strong>{{$t('username')}}</strong></q-item-label>
            <q-item-label caption>{{ admin?.username }}</q-item-label>
          </q-item-section>
        </q-item>
        <q-separator />

        <q-item >
          <q-item-section>
            <q-item-label><strong>{{ $t('role') }}</strong></q-item-label>
            <q-item-label caption>{{ admin?.role }}</q-item-label>
          </q-item-section>
        </q-item>
        <q-separator />

        <q-item>
          <q-item-section>
            <q-item-label><strong>{{ $t('createdAt') }}</strong></q-item-label>
            <q-item-label caption>{{ formatDate(admin?.createdAt) }}</q-item-label>
          </q-item-section>
        </q-item>
        <q-separator />

        <q-item>
          <q-item-section>
            <q-item-label><strong>{{ $t('updatedAt') }}</strong></q-item-label>
            <q-item-label caption>{{ formatDate(admin?.updatedAt) }}</q-item-label>
          </q-item-section>
        </q-item>


      </q-list>
    </q-card-section>



</template>


<script setup>
import {defineProps, onMounted,ref} from "vue";
import {useQuasar} from 'quasar'
import {useI18n} from "vue-i18n";
import {useRouter} from "vue-router";
import jalaali from "jalaali-js";

const router = useRouter()
const $q = useQuasar()
const {locale,t} = useI18n()
const admin = ref()
const props = defineProps({
  admin_name: {
    type: String,
    required: true
  }
});

onMounted(async () => {
  await getAdmin()
});

function formatDate(timestamp) {
  if (timestamp === -1) return 'N/A'; // Handle special case where timestamp is -1

  const date = new Date(timestamp * 1000);
  console.log(locale.value)

  if (locale.value === 'fa-IR') {

    const jDate = jalaali.toJalaali(date)

    return `${jDate.jy}/${jDate.jm}/${jDate.jd}`

  }
  if (locale.value === 'en-US') {
    var year = date.toLocaleString("default", {year: "numeric"});
    var month = date.toLocaleString("default", {month: "2-digit"});
    var day = date.toLocaleString("default", {day: "2-digit"});

    return year + "/" + month + "/" + day
  }
}

async function getAdmin() {
  const response = await fetch(
      '/api/panel/admin/user/get',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              username : props.admin_name,
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

  admin.value = await response.json();
}

</script>
