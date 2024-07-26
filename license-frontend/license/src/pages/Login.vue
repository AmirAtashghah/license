<template>
  <q-layout>
    <q-page-container>
      <q-page class="flex flex-center">
        <q-card class="q-pa-md" style="min-width: 30%;">

          <q-form
              @submit="onSubmit"
              class="q-gutter-md"

          >
            <q-input
                filled
                v-model="username"
                label="Your Username *"
                lazy-rules
                :rules="[ val => val && val.length > 0 || 'Please type something']"
            />

            <q-input
                filled
                type="password"
                v-model="password"
                label="Password *"
                lazy-rules
                :rules="[ val => val && val.length > 0 || 'Please type something']"

            />

            <div>
              <q-btn label="Login" type="submit" color="primary" style="width: 100%;"/>
            </div>
          </q-form>

        </q-card>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script setup>
import {useQuasar} from 'quasar'
import {ref} from 'vue'

const $q = useQuasar()

const username = ref(null)
const password = ref(null)

// fill values call login api

async function login() {
  const response = await fetch(
      '/api/login',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({username: username.value, password: password.value})
      }
  );
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }
  const data = await response.json();

  $q.notify({
    message: JSON.stringify(data),

  })
}





</script>
