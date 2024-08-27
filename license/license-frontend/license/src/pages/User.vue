<template>
  <div class="q-pa-md" >

    <div class="q-py-md">
      <q-btn icon="add" :label="t('addAdmin')" unelevated
             @click="openCreateAdminDialog" outline>
        <q-tooltip>{{ $t('addAdmin') }}</q-tooltip>
      </q-btn>
    </div>

    <q-table
        flat bordered
        :title="$t('admins')"
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
              @click="openUpdateAdminDialog(props.row.username)"
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
              @click="openAdminDetailDialog(props.row.username)"
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
              @click="openDeleteAdminDialog(props.row.id,props.row.name)"
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
      v-model="adminDetailDialog"
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
        <UserDetail :admin_name="admin_name_props" />
      </q-card-section>

    </q-card>
  </q-dialog>

  <q-dialog
      v-model="updateAdminDialog"
      transition-show="slide-up"
      transition-hide="slide-down"
      persistent

  >
    <q-card  style="min-width: 300px">
      <q-bar>
        <q-btn dense flat rounded icon="close" v-close-popup>
          <q-tooltip class="bg-white text-primary">{{ t('close') }}</q-tooltip>
        </q-btn>
      </q-bar>
      <q-card-section>
        <div class="text-h6">{{t('update')}}</div>
      </q-card-section>

      <q-card-section>
        <UpdateUser :admin_name="admin_name_props" :reloadEvent="reloadUserList"/>
      </q-card-section>

    </q-card>
  </q-dialog>

  <q-dialog
      v-model="deleteAdminDialog"
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
        <div class="text">{{t('areYouSureDelete')}}</div>
      </q-card-section>
      <q-card-section>
      <DeleteUser :admin_name="admin_name_props" :admin_id="admin_id_props" :reloadEvent="reloadUserList" />
      </q-card-section>
    </q-card>
  </q-dialog>

  <q-dialog
      v-model="createAdminDialog"
      transition-show="slide-up"
      transition-hide="slide-down"
      persistent
  >
    <q-card>
      <q-bar>
        <q-btn dense flat rounded icon="close" v-close-popup>
          <q-tooltip class="bg-white text-primary">{{ t('close') }}</q-tooltip>
        </q-btn>
      </q-bar>
      <q-card-section>
        <div class="text-h6">{{t('create')}}</div>
      </q-card-section>
      <q-card-section>
        <CreateUser :reloadEvent="reloadUserList"/>
      </q-card-section>

    </q-card>
  </q-dialog>

</template>

<script setup>
import UserDetail from "components/user/UserDetail.vue"
import CreateUser from "components/user/CreateUser.vue"
import UpdateUser from "components/user/UpdateUser.vue"
import DeleteUser from "components/user/DeleteUser.vue"
import {computed, onMounted, ref} from 'vue'
import {useQuasar} from 'quasar'
import { useI18n } from 'vue-i18n'
import {useRouter} from "vue-router";

const {t} = useI18n()
const $q = useQuasar()
const rows = ref([])
const admin_id_props = ref()
const admin_name_props = ref()
const router = useRouter()

const columns =computed(()=> [
  // {name: 'ID', label: t('id'), align: 'left', field: 'id', sortable: true},
  {name: 'Name', label: t('name'), align: 'left', field: 'name'},
  {name: 'Username', label: t('username'), align: 'center', field: 'username'},
  {name: 'Role', label: t('role'), align: 'center', field: 'role'},
  {name: 'Action', label: '', field: 'action'},
])

let adminDetailDialog = ref(false)
let updateAdminDialog = ref(false)
let deleteAdminDialog = ref(false)
let createAdminDialog = ref(false)

onMounted(async () => {
  await getUsers()
});

async function reloadUserList(){
  await getUsers()
}

async function getUsers() {

  const response = await fetch(
      '/api/panel/admin/user/list',
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
    rows.value=[]
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

        $q.notify({
          color: 'red',
          textColor: 'white',
          message: t('notAccess'),
          position:"top"
        });

      await router.push({name: 'login'})
    }

    throw new Error(`HTTP error! status: ${response.status}`);
  }
  rows.value = await response.json();

}

function openAdminDetailDialog(admin_name) {
  adminDetailDialog.value = !adminDetailDialog.value
  admin_name_props.value = admin_name
}


function openUpdateAdminDialog(admin_name) {
  updateAdminDialog.value = !updateAdminDialog.value
  admin_name_props.value = admin_name
}

function openDeleteAdminDialog(admin_id,admin_name) {
  deleteAdminDialog.value = !deleteAdminDialog.value
  admin_id_props.value = admin_id
  admin_name_props.value = admin_name
}

function openCreateAdminDialog() {
  createAdminDialog.value = !createAdminDialog.value
}

</script>
