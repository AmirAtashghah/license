<template>
  <div class="q-pa-md" >

    <div class="q-py-md">
      <q-btn icon="add" :label="t('addRestriction')" unelevated
             @click="openCreateRestrictionDialog" outline>
        <q-tooltip>{{ $t('addRestriction') }}</q-tooltip>
      </q-btn>
    </div>

    <q-table
        flat bordered
        :title="$t('restrictions')"
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
              @click="openUpdateRestrictionDialog(props.row.id)"
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
              @click="openRestrictionDetailDialog(props.row.id)"
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
              @click="openDeleteRestrictionDialog(props.row.id,props.row.key)"
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
      v-model="restrictionDetailDialog"
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
        <RestrictionDetail :restriction_id=restriction_id_props.toString() />
      </q-card-section>

    </q-card>
  </q-dialog>

  <q-dialog
      v-model="updateRestrictionDialog"
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
        <UpdateRestriction :restriction_id="restriction_id_props.toString()" :reloadEvent="reloadRestrictionList"/>
      </q-card-section>

    </q-card>
  </q-dialog>

  <q-dialog
      v-model="deleteRestrictionDialog"
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
        <DeleteRestriction :restriction_id="restriction_id_props.toString()" :restriction_key="restriction_key_props" :reloadEvent="reloadRestrictionList"/>
      </q-card-section>

    </q-card>
  </q-dialog>

  <q-dialog
      v-model="createRestrictionDialog"
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
        <CreateRestriction  :reloadEvent="reloadRestrictionList" />
      </q-card-section>

    </q-card>
  </q-dialog>

</template>

<script setup>
import UpdateRestriction from "components/restriction/UpdateRestriction.vue"
import CreateRestriction from "components/restriction/CreateRestriction.vue"
import RestrictionDetail from "components/restriction/RestrictionDetail.vue"
import DeleteRestriction from "components/restriction/DeleteRestriction.vue";
import {computed, onMounted, ref} from 'vue'
import {useQuasar} from 'quasar'
import { useI18n } from 'vue-i18n'
import {useRouter} from "vue-router";

const {t} = useI18n()
const $q = useQuasar()
const rows = ref([])
const restriction_id_props = ref()
const restriction_key_props = ref()
const router = useRouter()
const productID = ref()

const columns = computed(()=> [
  // {name: 'ID', label: t('id'), align: 'left', field: 'id', sortable: true},
  // {name: 'ProductID', label: t('productID'), align: 'left', field: 'product_id'},
  {name: 'Key', label: t('key'), align: 'left', field: 'key'},
  // {name: 'Value', label: t('value'), align: 'center', field: 'value'},
  {name: 'Action', label: '', field: 'action'},
])

let restrictionDetailDialog = ref(false)
let updateRestrictionDialog = ref(false)
let deleteRestrictionDialog = ref(false)
let createRestrictionDialog = ref(false)

onMounted(async () => {
  await getRestriction()
});

async function reloadRestrictionList(){
  await getRestriction()
}

async function getRestriction() {

  const response = await fetch(
      '/api/panel/restriction/list',
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
      await router.push({name: 'login'})

    }

    throw new Error(`HTTP error! status: ${response.status}`);
  }
  rows.value = await response.json();
}

function openRestrictionDetailDialog(restriction_id) {
  restrictionDetailDialog.value = !restrictionDetailDialog.value
  restriction_id_props.value = restriction_id
}

function openUpdateRestrictionDialog(restriction_id) {
  updateRestrictionDialog.value = !updateRestrictionDialog.value
  restriction_id_props.value = restriction_id
}

function openDeleteRestrictionDialog(restriction_id,restriction_key) {
  deleteRestrictionDialog.value = !deleteRestrictionDialog.value
  restriction_id_props.value = restriction_id
  restriction_key_props.value = restriction_key
}

function openCreateRestrictionDialog() {
  createRestrictionDialog.value = !createRestrictionDialog.value
  productID.value = props.product_id
}

</script>
