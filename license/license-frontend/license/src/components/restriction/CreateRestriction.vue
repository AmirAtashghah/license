<template>

    <div class="row justify-center">
      <q-card class="col-12">

        <div class="row">
          <div class="q-pt-md q-px-md col-md-12 col-sm-12 col-xs-12">
            <q-input

                filled
                v-model="restrictionKey"
                :label="$t('key')"
                lazy-rules
                :rules="[ val => val && val.length > 0 || $t('pleaseTypeSomething')]"

            />
          </div>
        </div>

        <div class="row justify-end">
          <div class="q-pa-md ">
            <q-btn
                :label="$t('create')"
                type="submit"
                color="primary"
                @click="createRestriction"
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
import {computed,defineProps, ref} from 'vue'
import {useQuasar} from "quasar";
import {useI18n} from "vue-i18n";
import {useRouter} from 'vue-router'

const router = useRouter()
const {t} = useI18n()
const loading = ref(false)
const $q = useQuasar()
const restrictionKey = ref();
const props = defineProps(['reloadEvent'])

const isValid = computed(() => {
  return (
      restrictionKey.value );
});

async function createRestriction() {

  loading.value = true

  const response = await fetch(

      '/api/panel/restriction/create',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              key: restrictionKey.value,
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
