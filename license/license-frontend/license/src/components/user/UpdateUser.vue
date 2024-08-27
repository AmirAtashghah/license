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
                :label="$t('update')"
                type="submit"
                color="primary"
                @click="updateAdmin"
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
const adminID = ref()
const adminName = ref();
const adminUsername = ref();
const adminPassword = ref('');
const adminRole = ref('admin');
const createdAt = ref()
const admin = ref()

const props = defineProps(['admin_name','reloadEvent'])

const isValid = computed(() => {
  return (
      adminName.value &&
      adminUsername.value &&
      adminRole.value
  );
});

onMounted(async () => {
  await getAdmin()
});

async function getAdmin() {

  const response = await fetch(
      '/api/panel/admin/user/get',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              username: props.admin_name,
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

  admin.value = await response.json();

  loadData(admin.value)

}

function loadData(data) {

  adminID.value = data.id;
  adminName.value = data.name;
  adminUsername.value = data.username;
  adminRole.value = data.role;
  createdAt.value = data.created_at
}

async function updateAdmin() {

  loading.value = true

  const response = await fetch(

      '/api/panel/admin/user/update',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              id: adminID.value,
              name: adminName.value,
              username: adminUsername.value,
              password: adminPassword.value,
              role: adminRole.value,
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
