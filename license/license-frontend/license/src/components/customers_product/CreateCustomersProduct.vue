<template>
  <div class="row justify-center">
    <q-card class="col-12">

      <div class="row ">

        <div class="q-pt-md q-px-md col-md-12 col-sm-12 col-xs-12">

          <q-select
              filled
              v-model="productID"
              :options="productOptions"
              :label="$t('product')"
              emit-value
              map-options
          />
        </div>

        <div class="q-pt-md q-px-md col-md-12 col-sm-12 col-xs-12">
          <q-select
              filled
              v-model="customerID"
              :options="customerOptions"
              :label="$t('customer')"
              emit-value
              map-options
          >

          </q-select>
        </div>

      </div>

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
      </div>

      <div class="row">
        <div class="q-pt-md q-px-md col-md-12 col-sm-12 col-xs-12">
          <q-input filled v-model="ExpiresAt" :label="$t('expireAt')" mask="date" :rules="['date']">
            <template v-slot:append>
              <q-icon name="event" class="cursor-pointer">
                <q-popup-proxy cover transition-show="scale" transition-hide="scale">
                  <q-date v-model="ExpiresAt" calendar="persian" first-day-of-week="6" :locale="myLocale" >
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

          <div v-for="(item,index) in options" class="row items-center full-width ">

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
              :label="$t('create')"
              type="submit"
              color="primary"
              @click="createCustomersProduct"
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
import {computed, onMounted, ref, watch} from 'vue'
import {useQuasar} from "quasar";
import {useI18n} from "vue-i18n";
import {useRouter} from 'vue-router'

const router = useRouter()
const {t} = useI18n()
const loading = ref(false)
const $q = useQuasar()
const productID = ref();
const customerID = ref()
const licenseType = ref();
const isActive = ref();
const ExpiresAt = ref();
const products = ref([])
const customers = ref([])
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
const props = defineProps(['reloadEvent'])
const myLocale = {

  days: 'یکشنبه_دوشنبه_سه‌شنبه_چهارشنبه_پنجشنبه_جمعه_شنبه'.split('_'),
  daysShort: 'یک_دو_سه_چهار_پنج_جمعه_شنبه'.split('_'),
  months: 'فروردین_اردیبهشت_خرداد_تیر_مرداد_شهریور_مهر_آبان_آذر_دی_بهمن_اسفند'.split('_'),
  monthsShort: 'فروردین_اردیبهشت_خرداد_تیر_مرداد_شهریور_مهر_آبان_آذر_دی_بهمن_اسفند'.split('_'),
  firstDayOfWeek: 1, // 0-6,
  format24h: true,
  pluralDay: 'روزها'
}
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
const options = ref([])

let restrictionOptions = []
let restrictionStr
let productOptions = ref([])
let customerOptions = ref([])

onMounted(async () => {
  await getCustomers()
  await getProducts()
  await getRestrictions()
});

function addoption() {
  options.value.push({
    selectedValue: null,
    inputValue: null,
    opts: restrictionOptions
  })
}

function removeoption(index){
  options.value.splice(index,1)
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

const isValid = computed(() => {
  return (
      customerID.value &&
      productID.value &&
      ExpiresAt.value
  );
});

async function getCustomers() {

  const response = await fetch(
      '/api/panel/customer/list',
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
  customers.value = await response.json();

  customerOptions.value = customers.value.map(customer => ({
    label: customer.name,
    value: customer.id
  }));

}

async function getProducts() {

  const response = await fetch(
      '/api/panel/product/list',
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

  products.value = await response.json();
  productOptions.value = products.value.map(product => ({
    label: product.name,
    value: product.id
  }));
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
}

async function createCustomersProduct() {

  loading.value = true

  const response = await fetch(
      '/api/panel/customer-product/create',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              customerID: customerID.value,
              productID: productID.value,
              restrictions: JSON.stringify(restrictionStr),
              licenseType: licenseType.value,
              isActive: isActive.value,
              expireAt: Math.floor(new Date(ExpiresAt.value).getTime() / 1000)
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
      // It's a single error
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

watch(options, function (newvalue) {
     restrictionStr = newvalue.map(restriction => ({
      id: restriction.selectedValue,
      value: restriction.inputValue,
    }));
},{deep:true})

</script>
