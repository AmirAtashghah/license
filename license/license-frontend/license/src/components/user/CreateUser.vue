<template>
    <div class="row justify-center">
      <q-card class="col-12">

        <div class="row">
          <div class="q-pt-md q-px-md col-md-12 col-sm-12 col-xs-12">
            <q-input

                filled
                v-model="adminName"
                :label="$t('name')"
                lazy-rules
                :rules="[ val => val && val.length > 0 || $t('pleaseTypeSomething')]"

            />
          </div>

          <div class="q-pt-md q-px-md col-md-12 col-sm-12 col-xs-12">

            <q-input
                filled
                v-model="adminUsername"
                :label="$t('username')"
                lazy-rules
                :rules="[ val => val && val.length > 0 || $t('pleaseTypeSomething')]"

            />
          </div>
        </div>
        <div class="row">
          <div class=" q-pt-md q-px-md col-md-12 col-sm-12 col-xs-12">

            <q-input
                filled
                v-model="adminPassword"
                :label="$t('password')"
                lazy-rules
                type="password"
                :rules="[ val => val && val.length > 0 || $t('pleaseTypeSomething')]"

            />
          </div>
          <div class=" q-pt-md q-px-md col-md-12 col-sm-12 col-xs-12">

            <q-input
                filled
                v-model="adminRole"
                :label="$t('role')"
                lazy-rules
                disable
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
                @click="createAdmin"
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
import {computed, defineEmits, ref} from 'vue'
import {useQuasar} from "quasar";
import {useI18n} from "vue-i18n";
import {useRouter} from 'vue-router'

const router = useRouter()
const {t} = useI18n()
const loading = ref(false)
const $q = useQuasar()
const adminName = ref();
const adminUsername = ref();
const adminPassword = ref();
const adminRole = ref('admin');
const props = defineProps(['reloadEvent'])

const isValid = computed(() => {
  return (
      adminName.value &&
      adminUsername.value &&
      adminPassword.value &&
      adminRole.value
  );
});


async function createAdmin() {

  loading.value = true

  const response = await fetch(

      '/api/panel/admin/user/create',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              name: adminName.value,
              username: adminUsername.value,
              password: adminPassword.value,
              role: adminRole.value
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
