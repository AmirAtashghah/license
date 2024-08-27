<template>
  <q-layout>
    <q-page-container>

      <q-page class="flex flex-center">
        <q-card class="q-gutter-md q-pa-md" style="min-width: 400px;">

          <q-card-section class="flex flex-center" >
                <img style="max-height: 200px" src="../assets/logo.png" alt="logo">
          </q-card-section>

          <q-card-section class="flex flex-center">
            <div class="text-h6">{{t('login')}}</div>
          </q-card-section>

          <q-card-section>
          <q-form
              @submit="login"
              class="q-gutter-md"

          >
            <q-input
                filled
                v-model="username"
                :label="$t('username')"
                icon="logout"
                lazy-rules

                :rules="[ val => val && val.length > 0 || $t('pleaseTypeSomething')]"
            >
              <template v-slot:prepend>
                <q-icon name="person" />
              </template>
            </q-input>

            <q-input
                filled
                type="password"
                v-model="password"
                :label="$t('password')"
                lazy-rules
                :rules="[ val => val && val.length > 0 || $t('pleaseTypeSomething')]"

            >
              <template v-slot:prepend>
                <q-icon name="password" />
              </template>
            </q-input>

            <div class="row justify-end">

              <q-btn :label="$t('login')"  type="submit" color="primary" class="end"/>

            </div>
          </q-form>
          </q-card-section>

          <q-card-section  >
          <q-btn flat round icon="language"  @click="toggleLanguage"
                 size="md">
            <q-tooltip style="width: 7em">{{ $t('language') }}</q-tooltip>
          </q-btn>
          </q-card-section>

        </q-card>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script setup>
import {useQuasar} from 'quasar'
import {onMounted, ref} from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'

const {locale,t} = useI18n()
const router = useRouter()
const $q = useQuasar()
const username = ref(null)
const password = ref(null)
const isFarsi = ref(true);

onMounted(async () => {
  locale.value = 'fa-IR';
  $q.lang.set({ rtl: true });
});

async function login() {

  const response = await fetch(
      '/api/login',
      {
        method: 'POST',
       headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({"username": username.value, "password": password.value})
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

    throw new Error(`HTTP error! status: ${response.status}`);
  }

  sessionStorage.setItem('username' , username.value)

  await router.push({ name: 'dashboard'})
}

const toggleLanguage = () => {
  if (isFarsi.value) {
    locale.value = 'en-US';
    $q.lang.set({ rtl: false });
  } else {
    locale.value = 'fa-IR';
    $q.lang.set({ rtl: true });
  }
  isFarsi.value = !isFarsi.value;
};

</script>
