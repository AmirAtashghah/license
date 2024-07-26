<template>
  <div class="q-pa-md ">
    <q-table
        :grid="$q.screen.xs"
        flat bordered
        title="Clients"
        :rows="rows"
        :columns="columns"
        row-key="name"
        :filter="filter"
        :loading="false"

    >
      <template v-slot:top-right>
        <q-input borderless dense debounce="300" v-model="filter" placeholder="Search">
          <template v-slot:append>
            <q-icon name="search"/>
          </template>
        </q-input>
      </template>

      <template v-slot:body-cell-Action="props">
        <q-td :props="props">
          <q-btn flat round icon="document_scanner" color="brown-5" @click="showClientLog">
            <q-tooltip>
              Log
            </q-tooltip>
          </q-btn>
          <q-btn flat round icon="system_update_alt" color="light-green-5" @click="showClientUpdate">
            <q-tooltip>
              Update client
            </q-tooltip>
          </q-btn>
          <q-btn flat round icon="display_settings" color="primary" @click="showClientDetail">
            <q-tooltip>
              Detail
            </q-tooltip>
          </q-btn>
        </q-td>
      </template>

      <template v-slot:body-cell-Status="props">
        <q-td :props="props">
          <q-icon :name="props.row.is_active?'toggle_on':'toggle_off'" size="md"
                  :color="props.row.is_active?'light-green-5':'red-5'">
            <q-tooltip>
              {{ props.row.IsActive ? 'active' : 'deactivate' }}
            </q-tooltip>
          </q-icon>
        </q-td>
      </template>
    </q-table>
  </div>

  <q-dialog
      v-model="clientDetail"
      persistent
      :maximized="maximizedToggle"
      transition-show="slide-up"
      transition-hide="slide-down"
  >
    <q-card>
      <q-bar style="background-color: white">
        <q-space/>

        <q-btn dense flat round icon="minimize" @click="maximizedToggle = false" :disable="!maximizedToggle">
          <q-tooltip v-if="maximizedToggle" class="bg-white text-primary">Minimize</q-tooltip>
        </q-btn>
        <q-btn dense flat round icon="crop_square" @click="maximizedToggle = true" :disable="maximizedToggle">
          <q-tooltip v-if="!maximizedToggle" class="bg-white text-primary">Maximize</q-tooltip>
        </q-btn>
        <q-btn dense flat round icon="close" v-close-popup>
          <q-tooltip class="bg-white text-primary">Close</q-tooltip>
        </q-btn>
      </q-bar>

      <q-card-section>
        <ClientDetail/>
      </q-card-section>

    </q-card>
  </q-dialog>

  <q-dialog
      v-model="clientLog"
      persistent
      :maximized="maximizedToggle"
      transition-show="slide-up"
      transition-hide="slide-down"
  >
    <q-card>
      <q-bar style="background-color: white">
        <q-space/>

        <q-btn dense flat round icon="minimize" @click="maximizedToggle = false" :disable="!maximizedToggle">
          <q-tooltip v-if="maximizedToggle" class="bg-white text-primary">Minimize</q-tooltip>
        </q-btn>
        <q-btn dense flat round icon="crop_square" @click="maximizedToggle = true" :disable="maximizedToggle">
          <q-tooltip v-if="!maximizedToggle" class="bg-white text-primary">Maximize</q-tooltip>
        </q-btn>
        <q-btn dense flat round icon="close" v-close-popup>
          <q-tooltip class="bg-white text-primary">Close</q-tooltip>
        </q-btn>
      </q-bar>

      <q-card-section>
        <ClientLog/>
      </q-card-section>

    </q-card>
  </q-dialog>

  <q-dialog
      v-model="clientUpdate"
      persistent
      :maximized="maximizedToggle"
      transition-show="slide-up"
      transition-hide="slide-down"
  >
    <q-card>
      <q-bar style="background-color: white">
        <q-space/>

        <q-btn dense flat round icon="minimize" @click="maximizedToggle = false" :disable="!maximizedToggle">
          <q-tooltip v-if="maximizedToggle" class="bg-white text-primary">Minimize</q-tooltip>
        </q-btn>
        <q-btn dense flat round icon="crop_square" @click="maximizedToggle = true" :disable="maximizedToggle">
          <q-tooltip v-if="!maximizedToggle" class="bg-white text-primary">Maximize</q-tooltip>
        </q-btn>
        <q-btn dense flat round icon="close" v-close-popup>
          <q-tooltip class="bg-white text-primary">Close</q-tooltip>
        </q-btn>
      </q-bar>

      <q-card-section>
        <ClientUpdate/>
      </q-card-section>

    </q-card>
  </q-dialog>

</template>

<script setup>
import ClientDetail from "components/ClientDetail.vue";
import ClientLog from "components/ClientLog.vue";
import ClientUpdate from "components/UpdateClient.vue"
import {onMounted, ref} from 'vue'
import {useRouter} from 'vue-router';
import {useQuasar} from 'quasar'

const $q = useQuasar()
const router = useRouter();
const columns = [
  {name: 'ID', label: 'ID', align: 'left', field: 'id', sortable: true},
  {name: 'SoftwareName', align: 'center', label: 'Software Name', field: 'software_name', sortable: true},
  {name: 'User', align: 'center', label: 'Owner', field: 'user_metadata'},
  {name: 'LicenseType', align: 'center', label: 'License Type', field: 'license_type'},
  {name: 'Status', align: 'center', label: 'Status', field: 'is_active'},
  {name: 'Action', label: '', field: 'action'},
]
const rows = ref([])


let maximizedToggle = ref(true)
let clientDetail = ref(false)
let clientLog = ref(false)
let clientUpdate = ref(false)

function showClientDetail() {
  clientDetail.value = !clientDetail.value
}

function showClientLog() {
  clientLog.value = !clientLog.value
}

function showClientUpdate() {
  clientUpdate.value = !clientUpdate.value
}

onMounted(() => {
  getClients()
});

async function getClients() {
  const response = await fetch(
      'http://127.0.0.1:3000/api/rows',
      {
        method: 'GET',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        // body: JSON.stringify({ // todo change filters
        //       "limit": null,
        //       "offset": null,
        //       "license_type": null,
        //       "user_metadata": null,
        //       "is_active": null
        //     }
        // )
      }
  );
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }
   rows.value = await response.json();

  // rows.value = data
  $q.notify({
    color: 'black',
    textColor: 'white',
    message: 'Load clients successfully'
  })
}

</script>
