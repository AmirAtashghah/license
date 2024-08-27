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
                :label="$t('update')"
                type="submit"
                color="primary"
                @click="updateRestriction"
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
const restrictionID = ref()
const restrictionProductID = ref();
const restrictionKey = ref();
const restrictionValue = ref();
const createdAt = ref()
const restriction = ref()
const props = defineProps(['restriction_id','reloadEvent'])

onMounted(async () => {
  await getRestriction()
});

const isValid = computed(() => {
  return (
      restrictionKey.value
  );
});

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

  loadData(restriction.value)
}

function loadData(data) {

  restrictionID.value = data.id;
  restrictionProductID.value = data.product_id;
  restrictionKey.value = data.key;
  restrictionValue.value = data.value;
  createdAt.value = data.createdAt
}



async function updateRestriction() {

  loading.value = true

  const response = await fetch(

      '/api/panel/restriction/update',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              id: restrictionID.value,
              key: restrictionKey.value,
              createdAt:createdAt.value,
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
