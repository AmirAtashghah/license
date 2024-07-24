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
            <q-icon name="search" />
          </template>
        </q-input>
      </template>

      <template v-slot:body-cell-Action="props">
        <q-td :props="props">
          <q-btn flat round  icon="document_scanner" color="brown-5"  to="ClientLog">
            <q-tooltip>
            Log
          </q-tooltip>
          </q-btn>
          <q-btn flat round icon="system_update_alt" color="light-green-5" to="UpdateClient">
            <q-tooltip>
              Update client
            </q-tooltip>
          </q-btn>
          <q-btn flat round icon="display_settings" color="primary" to="ClientDetail">
            <q-tooltip>
              Detail
            </q-tooltip>
          </q-btn>
        </q-td>
      </template >

      <template v-slot:body-cell-Status="props">
        <q-td :props="props">
          <q-icon  :name="props.row.IsActive?'toggle_on':'toggle_off'" size="md" :color="props.row.IsActive?'light-green-5':'red-5'"  >
            <q-tooltip>
              {{props.row.IsActive?'active':'deactivate'}}
            </q-tooltip>
          </q-icon>
        </q-td>
      </template>
    </q-table>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router';
const router = useRouter();

const columns = [

  { name: 'ID', label: 'ID',align: 'left' , field: 'ID', sortable: true },
  { name: 'SoftwareName', align: 'center', label: 'Software Name', field: 'SoftwareName', sortable: true },
  { name: 'User', align: 'center' ,label: 'Owner', field: 'User' },
  { name: 'LicenseType',align: 'center', label: 'License Type', field: 'LicenseType' },
  { name: 'Status',align: 'center', label: 'Status', field: 'IsActive' },
  { name: 'Action', label: '', field: 'action' },


]

const rows = [
  {
    ID:"1000-0000-0000-0000",
    SoftwareName:"Citra",
    User:"University of Tehran",
    LicenseType:"Trial",
    IsActive:true
  },  {
    ID:"0000-0000-0000-0000",
    SoftwareName:"Citra",
    User:"University of Tehran",
    LicenseType:"Trial",
    IsActive:true

  },  {
    ID:"0000-0000-0000-0000",
    SoftwareName:"Citra",
    User:"University of Tehran",
    LicenseType:"Trial",
    IsActive:true

  },  {
    ID:"0000-0000-0000-0000",
    SoftwareName:"Citra",
    User:"University of Tehran",
    LicenseType:"Trial",
    IsActive:false

  },  {
    ID:"0000-0000-0000-0000",
    SoftwareName:"Citra",
    User:"University of Tehran",
    LicenseType:"Trial",
    IsActive:false

  },  {
    ID:"0000-0000-0000-0000",
    SoftwareName:"Citra",
    User:"University of Tehran",
    LicenseType:"Trial",
    IsActive:true

  },  {
    ID:"0000-0000-0000-0000",
    SoftwareName:"Citra",
    User:"University of Tehran",
    LicenseType:"Trial",
    IsActive:true

  },
]

const client= {
  id: "1",
  softwareName: "Example Software",
  softwareVersion: "1.0.0",
  hardwareHash: "abc123",
  licenseType: "premium",
  userMetadata: "Some metadata",
  isActive: true,
  expiresAt: 1728000000,
  createdAt: 1609459200,
  updatedAt: 1609459200,
  deletedAt: null
}


</script>
