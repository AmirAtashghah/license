<template>
  <div class="q-pa-md" >

    <div class="q-py-md">
    <q-btn icon="add" :label="t('addProduct')" unelevated
           @click="openCreateProductDialog" outline>
      <q-tooltip>{{ $t('addProduct') }}</q-tooltip>
    </q-btn>
    </div>

    <q-table
        flat bordered
        :title="$t('products')"
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
              @click="openUpdateProductDialog(props.row.id)"
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
              @click="openProductDetailDialog(props.row.id)"
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
              @click="openDeleteProductDialog(props.row.id,props.row.name)"
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
      v-model="productDetailDialog"
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
        <ProductDetail :product_id=product_id_props />
      </q-card-section>

    </q-card>
  </q-dialog>

  <q-dialog
      v-model="restrictionDialog"
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
        <Restriction :product_id=product_id_props />
      </q-card-section>

    </q-card>
  </q-dialog>

  <q-dialog
      v-model="updateProductDialog"
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
        <UpdateProduct :product_id="product_id_props" :reloadEvent="reloadProductList"/>
      </q-card-section>

    </q-card>
  </q-dialog>

  <q-dialog
      v-model="deleteProductDialog"
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
        <DeleteProduct :product_id="product_id_props" :product_name="product_name_props" :reloadEvent="reloadProductList"/>
      </q-card-section>

    </q-card>
  </q-dialog>

  <q-dialog
      v-model="createProductDialog"
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
        <CreateProduct :reloadEvent="reloadProductList"/>
      </q-card-section>

    </q-card>
  </q-dialog>

</template>

<script setup>
import ProductDetail from "components/product/ProductDetail.vue"
import CreateProduct from "components/product/CreateProduct.vue"
import Restriction from "pages/Restriction.vue"
import UpdateProduct from "components/product/UpdateProduct.vue"
import DeleteProduct from "components/product/DeleteProduct.vue"
import {computed, onMounted, ref} from 'vue'
import {useQuasar} from 'quasar'
import { useI18n } from 'vue-i18n'
import {useRouter} from "vue-router";

const {t} = useI18n()
const $q = useQuasar()
const rows = ref([])
const product_id_props = ref()
const product_name_props = ref()
const router = useRouter()

const columns = computed(()=> [
  // {name: 'ID', label: t('id'), align: 'left', field: 'id', sortable: true},
  {name: 'Name', label: t('productName'), align: 'left', field: 'name', sortable: true},
  {name: 'Title', label: t('title'), align: 'center', field: 'title'},
  {name: 'Version', label: t('version'), align: 'center', field: 'version'},
  {name: 'Action', label: '', field: 'action'},
])

let productDetailDialog = ref(false)
let updateProductDialog = ref(false)
let deleteProductDialog = ref(false)
let createProductDialog = ref(false)
let restrictionDialog = ref(false)


onMounted(async () => {
  await getProducts()
});

async function reloadProductList(){
  await getProducts()
}

async function getProducts() {

  const response = await fetch(
      '/api/panel/product/list',
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

function openProductDetailDialog(product_id) {
  productDetailDialog.value = !productDetailDialog.value
  product_id_props.value = product_id
}

function openUpdateProductDialog(product_id) {
  updateProductDialog.value = !updateProductDialog.value
  product_id_props.value = product_id
}

function openDeleteProductDialog(product_id,product_name) {
  deleteProductDialog.value = !deleteProductDialog.value
  product_id_props.value = product_id
  product_name_props.value = product_name
}

function openCreateProductDialog() {
  createProductDialog.value = !createProductDialog.value
}

</script>
