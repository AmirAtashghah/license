<template>
  <div class="q-pa-md" >

    <div class="q-py-md">
      <q-btn icon="add" :label="t('addCustomer')" unelevated
             @click="openCreateCustomerDialog" outline>
        <q-tooltip>{{ $t('addCustomer') }}</q-tooltip>
      </q-btn>
    </div>

    <q-table
        flat bordered
        :title="$t('customers')"
        :rows="rows"
        :columns="columns"
        :rows-per-page-label="$t('recordPerPage')"
        :no-data-label="$t('noRecordFound')"
        row-key="name"
        :filter="filter"
        :loading="false"

    >

      <template v-slot:body-cell-Action="props">
        <q-td :props="props">

          <q-btn
              flat
              round
              icon="system_update_alt"
              color="light-green-5"
              @click="openUpdateCustomerDialog(props.row.id)"
          >
            <q-tooltip>
              {{$t('update')}}
            </q-tooltip>
          </q-btn>

          <q-btn
              flat
              round
              icon="display_settings"
              color="primary"
              @click="openCustomerDetailDialog(props.row.id)"
          >
            <q-tooltip>
              {{$t('detail')}}
            </q-tooltip>
          </q-btn>

          <q-btn
              flat
              round
              icon="delete_outline"
              color="red"
              @click="openDeleteCustomerDialog(props.row.id,props.row.name)"
          >
            <q-tooltip>
              {{$t('delete')}}
            </q-tooltip>
          </q-btn>

        </q-td>
      </template>

    </q-table>
  </div>

  <q-dialog
      v-model="customerDetailDialog"
      transition-show="slide-up"
      transition-hide="slide-down"
      persistent

  >
    <q-card style="min-width: 300px">

      <q-bar>
        <q-btn dense flat rounded icon="close" v-close-popup>
          <q-tooltip class="bg-white text-primary">{{ t('close') }}</q-tooltip>
        </q-btn>
      </q-bar>

      <q-card-section>
        <div class="text-h6">{{t('detail')}}</div>
      </q-card-section>
      <q-card-section>
        <CustomerDetail :customer_id=customer_id_props />
      </q-card-section>

    </q-card>
  </q-dialog>

  <q-dialog
      v-model="updateCustomerDialog"
      transition-show="slide-up"
      transition-hide="slide-down"
      persistent

  >
    <q-card style="min-width: 300px">
      <q-bar>
        <q-btn dense flat rounded icon="close" v-close-popup>
          <q-tooltip class="bg-white text-primary">{{ t('close') }}</q-tooltip>
        </q-btn>
      </q-bar>
      <q-card-section>
        <div class="text-h6">{{t('update')}}</div>
      </q-card-section>
      <q-card-section>
        <UpdateCustomer :customer_id="customer_id_props" :reloadEvent="reloadCustomerList"/>
      </q-card-section>

    </q-card>
  </q-dialog>

  <q-dialog
      v-model="deleteCustomerDialog"

      transition-show="slide-up"
      transition-hide="slide-down"
      persistent
  >
    <q-card style="min-width: 400px">
      <q-bar>
        <q-btn dense flat rounded icon="close" v-close-popup>
          <q-tooltip class="bg-white text-primary">{{ t('close') }}</q-tooltip>
        </q-btn>
      </q-bar>
      <q-card-section>
        <div class="text">{{t('areYouSureDelete')}}</div>
      </q-card-section>

      <q-card-section>
        <DeleteCustomer :customer_id="customer_id_props" :customer_name="customer_name_props" :reloadEvent="reloadCustomerList"/>
      </q-card-section>

    </q-card>
  </q-dialog>

  <q-dialog
      v-model="createCustomerDialog"

      transition-show="slide-up"
      transition-hide="slide-down"
      persistent
  >
    <q-card style="min-width: 300px">
      <q-bar>
        <q-btn dense flat rounded icon="close" v-close-popup>
          <q-tooltip class="bg-white text-primary">{{ t('close') }}</q-tooltip>
        </q-btn>
      </q-bar>
      <q-card-section>
        <div class="text-h6">{{t('create')}}</div>
      </q-card-section>

      <q-card-section>
        <CreateCustomer :reloadEvent="reloadCustomerList" />
      </q-card-section>

    </q-card>
  </q-dialog>
</template>

<script setup>
import CreateCustomer from "components/customer/CreateCustomer.vue";
import UpdateCustomer from "components/customer/UpdateCustomer.vue";
import DeleteCustomer from "components/customer/DeleteCustomer.vue";
import CustomerDetail from "components/customer/CustomerDetail.vue";
import {computed, onMounted, ref} from 'vue'
import {useQuasar} from 'quasar'
import { useI18n } from 'vue-i18n'
import {useRouter} from "vue-router";

const {t} = useI18n()
const $q = useQuasar()
const rows = ref([])
const customer_id_props = ref()
const customer_name_props = ref()
const router = useRouter()

const columns = computed(()=>[
  // {name: 'ID', label: t('id'), align: 'left', field: 'id', sortable: true},
  {name: 'Name', label: t('customerName'), align: 'left', field: 'name', sortable: true},
  {name: 'Email', label: t('email'), align: 'center', field: 'email'},
  {name: 'Phone', label: t('phone'), align: 'center', field: 'phone'},
  {name: 'Action', label: '', field: 'action'},
])

let customerDetailDialog = ref(false)
let deleteCustomerDialog = ref(false)
let updateCustomerDialog = ref(false)
let createCustomerDialog = ref(false)



onMounted(async () => {
  await getCustomers()
});

async function reloadCustomerList(){
  await getCustomers()
}

async function getCustomers() {

  const response = await fetch(
      '/api/panel/customer/list',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({ // todo change filters
              "limit": 0,
              "offset": 0
            }
        )
      }
  );
  if (!response.ok) {

    rows.value = []
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
  rows.value = await response.json();

}

function openCreateCustomerDialog() {
  createCustomerDialog.value = !createCustomerDialog.value
}

function openUpdateCustomerDialog(customer_id) {
  updateCustomerDialog.value = !updateCustomerDialog.value
  customer_id_props.value = customer_id
}

function openCustomerDetailDialog(customer_id) {
  customerDetailDialog.value = !customerDetailDialog.value
  customer_id_props.value = customer_id
}

function openDeleteCustomerDialog(customer_id,customer_name) {
  deleteCustomerDialog.value = !deleteCustomerDialog.value
  customer_id_props.value = customer_id
  customer_name_props.value = customer_name
}

</script>
