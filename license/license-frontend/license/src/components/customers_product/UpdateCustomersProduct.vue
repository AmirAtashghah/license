<template>
    <div class="row justify-center">
      <q-card class="col-12">

        <div class="row">
          <div class="q-pt-md q-px-md col-md-12 col-sm-12 col-xs-12">
            <q-select
                filled
                v-model="licenseType"
                :options="LicenseTypeOptions"
                :label="$t('licenseType')"
                emit-value
                map-options
            />
          </div>

          <div class="q-pt-md q-px-md col-md-12 col-sm-12 col-xs-12">
            <q-select
                filled
                v-model="isActive"
                :options="IsActiveOptions"
                :label="$t('status')"
                emit-value
                map-options
            />
          </div>

          <div class="q-pt-md q-px-md col-md-12 col-sm-12 col-xs-12">
            <q-input filled v-model="expireAt" :label="$t('expireAt')" mask="date" :rules="['date']">
              <template v-slot:append>
                <q-icon name="event" class="cursor-pointer">
                  <q-popup-proxy cover transition-show="scale" transition-hide="scale">
                    <q-date v-model="expireAt" calendar="persian" first-day-of-week="6" :locale="myLocale" >
                      <div class="row items-center justify-end">
                        <q-btn v-close-popup :label="$t('close')" color="primary" flat/>
                      </div>
                    </q-date>
                  </q-popup-proxy>
                </q-icon>
              </template>
            </q-input>
          </div>

        </div>

        <div class="row" >
          <div class="row q-pt-md q-px-md col-md-12 col-sm-12 col-xs-12" >
            <div>
              <div class="text-h8">{{t('restriction')}}</div>
            </div>
            <q-separator></q-separator>
        <q-btn flat icon="add" :label="$t('add')" @click="addoption" class="q-pa-md col-md-12 col-sm-12 col-xs-12" ></q-btn>

        <div v-for="(item,index) in options" class="row items-center full-width">

          <q-select
              filled
              v-model="item.selectedValue"
              :options="item.opts"
              emit-value
              map-options
              @update:model-value="onSelected(index,item.selectedValue)"
              :label="$t('restrictions')"
              class=" q-pa-md col-md-6 col-sm-12 col-xs-12"
          >
          </q-select>
          <q-input
              class=" q-pa-md col-md-5 col-sm-11 col-xs-11"
              filled
              v-model="item.inputValue"
              :label="$t('value')"
              lazy-rules
              :rules="[ val => val && val.length > 0 || $t('pleaseTypeSomething')]"/>

          <q-btn rounded fab-mini flat icon="delete" style="height: 50px" color="red" @click="removeoption(index)" class=" q-px-md col-md-1 col-sm-1 col-xs-1"></q-btn>
        </div>
          </div>
        </div>

        <div class="row justify-end">
          <div class="q-pa-md ">
            <q-btn
                :label="$t('update')"
                type="submit"
                color="primary"
                @click="updateCustomersProduct"
                :disable="!isValid"
                :loading="loading"
                class="col-md-6"
                v-close-popup
            />
          </div>
        </div>

      </q-card>
    </div>
</template>

<script setup>
import {computed, defineEmits, defineProps, onMounted, ref} from 'vue'
import {useQuasar} from "quasar";
import {useI18n} from "vue-i18n";
import {useRouter} from 'vue-router'
import jalaali from "jalaali-js";
import dayJS from '../../boot/date'

const router = useRouter()
const {locale,t} = useI18n()
const loading = ref(false)
const $q = useQuasar()
const customersProductID = ref()
const customerID = ref();
const productID = ref();
const hardwareHash = ref();
const firstConfirmedAt = ref()
const lastConfirmedAt = ref()
const createdAt = ref()
let customersProduct
const licenseType = ref();
const isActive = ref();
const expireAt = ref();
const myLocale = {

  days: 'یکشنبه_دوشنبه_سه‌شنبه_چهارشنبه_پنجشنبه_جمعه_شنبه'.split('_'),
  daysShort: 'یک_دو_سه_چهار_پنج_جمعه_شنبه'.split('_'),
  months: 'فروردین_اردیبهشت_خرداد_تیر_مرداد_شهریور_مهر_آبان_آذر_دی_بهمن_اسفند'.split('_'),
  monthsShort: 'فروردین_اردیبهشت_خرداد_تیر_مرداد_شهریور_مهر_آبان_آذر_دی_بهمن_اسفند'.split('_'),
  firstDayOfWeek: 1, // 0-6,
  format24h: true,
  pluralDay: 'روزها'
}
const IsActiveOptions = [
  {
    label: t('active'),
    value: true
  },
  {
    label: t('inactive'),
    value: false
  }
]
const LicenseTypeOptions = [
  {
    label: t('test'),
    value: 'test'
  },
  {
    label: t('primary'),
    value: 'primary'
  }
]
const props = defineProps(['customersProduct_id','reloadEvent'])

const isValid = computed(() => {
  return (
      expireAt.value &&
      licenseType.value
  );
});
const options = ref([])

let restrictionOptions = []

onMounted(async () => {
  await getCustomersProduct()
  await getRestrictions()
});

function addoption() {
  options.value.push({
    selectedValue: null,
    inputValue: null,
    opts: restrictionOptions
  })

}

function onSelected(index,selectedValue) {

  for (const i in options.value) {
    if (index !== parseInt(i)){

      for (const j in options.value[i].opts){
        if (selectedValue === options.value[index].opts[j].value){
          options.value[i].opts.splice(j, 1)
        }
      }
    }
  }
}

function removeoption(index){
  options.value.splice(index,1)
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

  customersProduct = await response.json();
  loadData(customersProduct)
}

function loadData(data) {

  customersProductID.value = data.id;
  customerID.value = data.customerID;
  productID.value = data.productID;
  hardwareHash.value = data.hardwareHash;
  licenseType.value = data.licenseType;
  isActive.value = data.isActive;
  expireAt.value = formatTimestamp(data.expireAt)
  firstConfirmedAt.value = data.firstConfirmedAt
  lastConfirmedAt.value = data.lastConfirmedAt
  createdAt.value = data.createdAt
}

function formatTimestamp(timestamp) {
  if (timestamp === -1) return 'N/A'; // Handle special case where timestamp is -1

  const date = new Date(timestamp * 1000);

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

async function getRestrictions() {

  const response = await fetch(
      '/api/panel/restriction/list',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({ // todo change filters
              "limit": 0,
              "offset": 0
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

  const t = await response.json();
  restrictionOptions = t.map(restriction => ({
    label: restriction.key,
    value: restriction.id,
  }));

  options.value = customersProduct.restrictions.map(item => ({
    selectedValue: item.restriction_id,
    inputValue: item.value,
    opts:restrictionOptions
  }))

}

async function updateCustomersProduct() {

  loading.value = true

  const response = await fetch(

      '/api/panel/customer-product/update',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              id: customersProductID.value,
              customerID: customerID.value,
              productId: productID.value,
              hardwareHash: hardwareHash.value,
              licenseType: licenseType.value,
              isActive: isActive.value,
              expireAt: Math.floor(new Date(expireAt.value).getTime() / 1000),
              firstConfirmedAt:  firstConfirmedAt.value,
              lastConfirmedAt:  lastConfirmedAt.value,
              createdAt: createdAt.value,
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

    loading.value = false

    throw new Error(`HTTP error! status: ${response.status}`);
  }

  loading.value = false

  props.reloadEvent()
}
</script>
