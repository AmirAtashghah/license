<template>

    <q-card-section>
      <q-list bordered padding class="rounded-borders">
        <q-item>
          <q-item-section>
            <q-item-label><strong>{{$t('id')}}</strong> </q-item-label>
            <q-item-label  caption>{{ restriction?.id }}</q-item-label>
          </q-item-section>
        </q-item>
        <q-separator />

        <q-item >
          <q-item-section>
            <q-item-label><strong>{{$t('key')}}</strong></q-item-label>
            <q-item-label caption>{{ restriction?.key }}</q-item-label>
          </q-item-section>
        </q-item>
        <q-separator />

        <q-item>
          <q-item-section>
            <q-item-label><strong>{{ $t('createdAt') }}</strong></q-item-label>
            <q-item-label caption>{{ formatDate(restriction?.createdAt) }}</q-item-label>
          </q-item-section>
        </q-item>
        <q-separator />

        <q-item>
          <q-item-section>
            <q-item-label><strong>{{ $t('updatedAt') }}</strong></q-item-label>
            <q-item-label caption>{{ formatDate(restriction?.updatedAt) }}</q-item-label>
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
import dayJS from "boot/date";

const router = useRouter()
const $q = useQuasar()
const {locale,t} = useI18n()
const restriction = ref()
const props = defineProps({
  restriction_id: {
    type: String,
    required: true
  }
});

onMounted(async () => {
  await getRestriction()
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

async function getRestriction() {
  const response = await fetch(
      '/api/panel/restriction/get',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              "id": props.restriction_id,
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

  restriction.value = await response.json();
}

</script>
