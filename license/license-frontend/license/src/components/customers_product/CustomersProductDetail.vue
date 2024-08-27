<template>

    <q-card-section>
      <q-list bordered padding class="rounded-borders">
        <q-item>
          <q-item-section>
            <q-item-label><strong>{{$t('id')}}</strong> </q-item-label>
            <q-item-label  caption>{{ customersProduct?.id }}</q-item-label>
          </q-item-section>
        </q-item>
        <q-separator />

        <q-item >
          <q-item-section>
            <q-item-label><strong>{{$t('customerID')}}</strong></q-item-label>
            <q-item-label caption>{{ customersProduct?.customerID }}</q-item-label>
          </q-item-section>
        </q-item>
        <q-separator />

        <q-item >
          <q-item-section>
            <q-item-label><strong>{{$t('productID')}}</strong></q-item-label>
            <q-item-label caption>{{ customersProduct?.productID }}</q-item-label>
          </q-item-section>
        </q-item>
        <q-separator />

        <q-item >
          <q-item-section>
            <q-item-label><strong>{{ $t('hardwareHash') }}</strong></q-item-label>
            <q-item-label caption>{{ customersProduct?.hardwareHash }}</q-item-label>
          </q-item-section>
        </q-item>
        <q-separator />

        <q-item >
          <q-item-section>
            <q-item-label><strong>{{ $t('licenseType') }}</strong></q-item-label>
            <q-item-label caption>{{ customersProduct?.licenseType }}</q-item-label>
          </q-item-section>
        </q-item>
        <q-separator />

        <q-item>
          <q-item-section>
            <q-item-label><strong>{{ $t('status') }}</strong></q-item-label>
            <q-item-label caption>{{ customersProduct?.isActive ? $t('active') : $t('inactive') }}</q-item-label>
          </q-item-section>
        </q-item>
        <q-separator />

        <q-item>
          <q-item-section>
            <q-item-label><strong>{{ $t('expireAt') }}</strong></q-item-label>
            <q-item-label caption>{{ formatTimestamp(customersProduct?.expireAt) }}</q-item-label>
          </q-item-section>
        </q-item>
        <q-separator />

        <q-item>
          <q-item-section>
            <q-item-label><strong>{{ $t('firstConfirmedAt') }}</strong></q-item-label>
            <q-item-label caption>{{ formatDate(customersProduct?.firstConfirmedAt) }}</q-item-label>
          </q-item-section>
        </q-item>

        <q-separator />

        <q-item>
          <q-item-section>
            <q-item-label><strong>{{ $t('lastConfirmedAt') }}</strong></q-item-label>
            <q-item-label caption>{{ formatDate(customersProduct?.lastConfirmedAt) }}</q-item-label>
          </q-item-section>
        </q-item>
        <q-separator />

        <q-item>
          <q-item-section>
            <q-item-label><strong>{{ $t('createdAt') }}</strong></q-item-label>
            <q-item-label caption>{{ formatDate(customersProduct?.createdAt) }}</q-item-label>
          </q-item-section>
        </q-item>
        <q-separator />

        <q-item>
          <q-item-section>
            <q-item-label><strong>{{ $t('updatedAt') }}</strong></q-item-label>
            <q-item-label caption>{{ formatDate(customersProduct?.updatedAt) }}</q-item-label>
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
const customersProduct = ref()
const props = defineProps({
  customersProduct_id: {
    type: String,
    required: true
  }
});

onMounted(async () => {
  await getCustomersProduct()
});

function formatDate(timestamp) {
  if (timestamp === -1) return 'N/A'; // Handle special case where timestamp is -1

  const date = new Date(timestamp * 1000);

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

function formatTimestamp(timestamp) {
  if (timestamp === -1) return 'N/A'; // Handle special case where timestamp is -1

  const date = new Date(timestamp);

  if (locale.value === 'fa-IR') {

    return dayJS.unix(timestamp).format('YYYY/MM/DD')
  }
  if (locale.value === 'en-US') {
    var year = date.toLocaleString("default", {year: "numeric"});
    var month = date.toLocaleString("default", {month: "2-digit"});
    var day = date.toLocaleString("default", {day: "2-digit"});

    return year + "/" + month + "/" + day
  }
}

async function getCustomersProduct() {
  const response = await fetch(
      '/api/panel/customer-product/get',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              "id": props.customersProduct_id,
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

  customersProduct.value = await response.json();
}

</script>
