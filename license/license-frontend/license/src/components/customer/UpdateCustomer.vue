<template>

    <div class="row ">
      <q-card class="col-12">

        <div class="row">
          <div class="q-pt-md q-px-md col-md-12 col-sm-12 col-xs-12">
            <q-input

                filled
                v-model="customerName"
                :label="$t('customerName')"
                lazy-rules
                :rules="[ val => val && val.length > 0 || $t('pleaseTypeSomething')]"

            />
          </div>

          <div class="q-pt-md q-px-md col-md-12 col-sm-12 col-xs-12">

            <q-input
                filled
                v-model="customerEmail"
                :label="$t('email')"
                lazy-rules
                type="email"
                :rules="[ val => val && val.length > 0 || $t('pleaseTypeSomething')]"

            />
          </div>

          <div class=" q-pt-md q-px-md col-md-12 col-sm-12 col-xs-12">

            <q-input
                filled
                v-model="customerPhone"
                :label="$t('phone')"
                lazy-rules
                type="number"
                :rules="[ val => val && val.length > 0 || $t('pleaseTypeSomething')]"

            />
          </div>
        </div>

        <div class="row justify-end">
          <div class="q-pa-md ">
            <q-btn
                :label="$t('update')"
                type="submit"
                color="primary"
                @click="updateCustomer"
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

const router = useRouter()
const {t} = useI18n()
const loading = ref(false)
const $q = useQuasar()
const customerID = ref()
const customerName = ref();
const customerPhone = ref();
const customerEmail = ref();
const createdAt = ref()
const customer = ref()
const props = defineProps(['customer_id','reloadEvent']);

onMounted(async () => {
  await getCustomer()
});

const isValid = computed(() => {
  return (
      customerName.value &&
      customerPhone.value &&
      customerEmail.value
  );
});


async function getCustomer() {

  const response = await fetch(
      '/api/panel/customer/get',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              "id": props.customer_id,
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

  customer.value = await response.json();

  loadData(customer.value)
}

function loadData(data) {

  customerID.value = data.id;
  customerName.value = data.name;
  customerEmail.value = data.email;
  customerPhone.value = data.phone;
  createdAt.value = data.createdAt
}

async function updateCustomer() {

  loading.value = true

  const response = await fetch(

      '/api/panel/customer/update',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              id: customerID.value,
              name: customerName.value,
              email: customerEmail.value,
              phone: customerPhone.value,
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
