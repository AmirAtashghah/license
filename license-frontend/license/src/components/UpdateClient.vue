<template>
  <q-page >
    <div style="margin: 70px" >

      <q-field filled label="Filled" style="margin: 30px" stack-label>
        <template v-slot:control>
          <div class="self-center full-width no-outline" tabindex="0">ID</div>
        </template>
      </q-field>

      <q-input filled v-model="text" style="margin: 30px" label="Software Name" />

      <q-input filled v-model="text" style="margin: 30px" label="Software Version" />

      <q-field filled label="Filled" style="margin: 30px" stack-label>
        <template v-slot:control>
          <div class="self-center full-width no-outline" tabindex="0">Hardware Hash</div>
        </template>
      </q-field>

      <q-input filled v-model="text" style="margin: 30px" label="License Type" />

      <q-input filled v-model="text" style="margin: 30px" label="User Metadata" />


      <q-select filled v-model="model"  :options="options" style="margin: 30px" label="Activation Status" />

      <q-input filled v-model="text" style="margin: 30px" label="Expires Time" />

      <q-field filled label="Filled" style="margin: 30px" stack-label>
        <template v-slot:control>
          <div class="self-center full-width no-outline" tabindex="0">Created Time</div>
        </template>
      </q-field>

      <q-field filled label="Filled" style="margin: 30px" stack-label>
        <template v-slot:control>
          <div class="self-center full-width no-outline" tabindex="0">Updated Time</div>
        </template>
      </q-field>

      <q-field filled label="Filled" style="margin: 30px" stack-label>
        <template v-slot:control>
          <div class="self-center full-width no-outline" tabindex="0">Deleted Time</div>
        </template>
      </q-field>

      <q-btn :loading="loading[0]" style="width:100%;" color="light-green-5" @click="simulateProgress(0)" label="Update" />

    </div>
  </q-page>
</template>

<script>
import { ref } from 'vue'

export default {
  setup () {
    const loading = ref([false])

    const progress = ref(false)

    function simulateProgress (number) {
      // we set loading state
      loading.value[ number ] = true

      // simulate a delay
      setTimeout(() => {
        // we're done, we reset loading state
        loading.value[ number ] = false
      }, 3000)
    }
    return {
      loading,
      progress,
      simulateProgress,
      text: ref(''),
      model: ref(null),
      options: [
        'Enable', 'Disable'
      ]
    }
  }
}
</script>
