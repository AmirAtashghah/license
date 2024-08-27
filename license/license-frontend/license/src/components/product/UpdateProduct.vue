<template>

    <div class="row justify-center">
      <q-card class="col-12">

        <div class="row">
          <div class="q-pt-md q-px-md col-md-12 col-sm-12 col-xs-12">
            <q-input

                filled
                v-model="productTitle"
                :label="$t('productTitle')"
                lazy-rules
                :rules="[ val => val && val.length > 0 || $t('pleaseTypeSomething')]"

            />
          </div>

          <div class="q-pt-md q-px-md col-md-12 col-sm-12 col-xs-12">

            <q-input
                filled
                v-model="productName"
                :label="$t('productName')"
                lazy-rules
                :rules="[ val => val && val.length > 0 || $t('pleaseTypeSomething')]"

            />
          </div>

          <div class=" q-pt-md q-px-md col-md-12 col-sm-12 col-xs-12">

            <q-input
                filled
                v-model="productVersion"
                :label="$t('productVersion')"
                lazy-rules
                mask="v.##.##"
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
                @click="updateProduct"
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
import {computed, defineProps, onMounted, ref} from 'vue'
import {useQuasar} from "quasar";
import {useI18n} from "vue-i18n";
import {useRouter} from 'vue-router'

const router = useRouter()
const {t} = useI18n()
const loading = ref(false)
const $q = useQuasar()
const productID = ref()
const productName = ref();
const productVersion = ref();
const productTitle = ref();
const createdAt = ref()
const product = ref()
const props = defineProps(['product_id','reloadEvent'])

onMounted(async () => {
  await getProduct()
});

const isValid = computed(() => {
  return (
      productName.value &&
      productVersion.value &&
      productTitle.value
  );
});

async function getProduct() {

  const response = await fetch(
      '/api/panel/product/get',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              "id": props.product_id,
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

  product.value = await response.json();

  loadData(product.value)
}

function loadData(data) {

  productID.value = data.id;
  productName.value = data.name;
  productVersion.value = data.version;
  productTitle.value = data.title;
  createdAt.value = data.createdAt
}

async function updateProduct() {

  loading.value = true

  const response = await fetch(

      '/api/panel/product/update',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              id: productID.value,
              name: productName.value,
              title: productTitle.value,
              version: productVersion.value,
              createdAt: Math.floor(new Date(createdAt.value).getTime() / 1000),
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
