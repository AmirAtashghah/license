
<template>
  <div class="row">
    <div class="q-gutter-md col-lg-12 col-sm-12 col-md-12 col-xs-12 ">

      <q-btn
          color="red"
          :label="t('delete')"
          @click="deleteCustomersProduct"
          v-close-popup
      />
    </div>
  </div>
</template>

<script setup>
import {defineEmits, defineProps, ref} from 'vue';
import {useI18n} from "vue-i18n";
import {useQuasar} from "quasar";
import {useRouter} from "vue-router";

const router = useRouter()
const $q = useQuasar()
const {t} = useI18n()
const props = defineProps(['customersProduct_id','reloadEvent'])

async function deleteCustomersProduct() {

  const response = await fetch(

      '/api/panel/customer-product/delete',
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

  props.reloadEvent()
}

</script>

