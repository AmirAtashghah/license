<template>
    <div class="row justify-center">
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
                :error-message="emailError"
                :rules="[validateEmail]"

            />
          </div>

          <div class=" q-pt-md q-px-md col-md-12 col-sm-12 col-xs-12">

            <q-input
                filled
                v-model="customerPhone"
                :label="$t('phone')"
                lazy-rules
                type="number"
                :error-message="phoneError"
                :rules="[validatePhone]"

            />
          </div>
        </div>

        <div class="row justify-end">
          <div class="q-pa-md ">
            <q-btn
                :label="$t('create')"
                type="submit"
                color="primary"
                @click="createCustomer"
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
import {defineEmits,computed, ref} from 'vue'
import {useQuasar} from "quasar";
import {useI18n} from "vue-i18n";
import {useRouter} from 'vue-router'

const router = useRouter()
const {t} = useI18n()
const loading = ref(false)
const $q = useQuasar()
const customerName = ref();
const customerEmail = ref();
const customerPhone = ref();
const props = defineProps(['reloadEvent'])
const emailError = ref('');
const phoneError = ref('');

const isValid = computed(() => {
  return (
      customerName.value &&

      validateEmail(customerEmail.value) &&
      validatePhone(customerPhone.value)
  );
});

const validateEmail = (val) => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (!val) {
    emailError.value = t('pleaseTypeSomething');
    return false;
  } else if (!emailRegex.test(val)) {
    emailError.value = t('inValidEmail');
    return false;
  }
  emailError.value = '';
  return true;
};

const validatePhone = (val) => {
  if (!val) {
    phoneError.value = t('pleaseTypeSomething');
    return false;
  } else if (val.length !== 11) {
    phoneError.value = t('inValidPhone');
    return false;
  }
  phoneError.value = '';
  return true;
};

async function createCustomer() {

  loading.value = true

  const response = await fetch(

      '/api/panel/customer/create',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              name: customerName.value,
              email:customerEmail.value,
              phone: customerPhone.value,
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
