<template>
  <div class="q-pa-md" >

    <div class="q-py-md">
      <q-btn icon="add" :label="t('addProductToCustomer')" unelevated
             @click="openCreateCustomersProductDialog" outline>
        <q-tooltip>{{ $t('addProductToCustomer') }}</q-tooltip>
      </q-btn>
    </div>

    <q-table
        flat bordered
        :title="$t('customersProduct')"
        :rows="formattedData"
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
              @click="openUpdateCustomersProductDialog(props.row.id)"
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
              @click="openCustomersProductDetailDialog(props.row.id)"
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
              @click="openDeleteCustomersProductDialog(props.row.id)"
          >
            <q-tooltip>
              {{$t('delete')}}
            </q-tooltip>
          </q-btn>
        </q-td>
      </template>

      <template v-slot:body-cell-Status="props">
        <q-td :props="props">
          <q-icon :name="props.row.isActive?'radio_button_checked':'radio_button_checked'" size="md"
                  :color="props.row.isActive?'light-green-5':'red-5'">
            <q-tooltip>
              {{ props.row.isActive ? $t('active') : $t('inactive') }}
            </q-tooltip>
          </q-icon>
        </q-td>
      </template>

    </q-table>
  </div>

  <q-dialog
      v-model="customersProductDetailDialog"
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
        <CustomersProductDetail :customersProduct_id=customersProduct_id_props />
      </q-card-section>

    </q-card>
  </q-dialog>

  <q-dialog
      v-model="updateCustomersProductDialog"
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
        <div class="text-h6">{{t('DeleteCustomer')}}</div>
      </q-card-section>

      <q-card-section>
        <UpdateCustomersProduct :customersProduct_id="customersProduct_id_props" :reloadEvent="reloadCustomersProductList"/>
      </q-card-section>

    </q-card>
  </q-dialog>

  <q-dialog
      v-model="deleteCustomersProductDialog"
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
        <DeleteCustomersProduct :customers-product_id="customersProduct_id_props" :reloadEvent="getCustomersProducts"/>
      </q-card-section>

    </q-card>
  </q-dialog>

  <q-dialog
      v-model="createCustomersProductDialog"
      transition-show="slide-up"
      transition-hide="slide-down"
      persistent
  >
    <q-card style="min-width: 500px">
      <q-bar>
        <q-btn dense flat rounded icon="close" v-close-popup>
          <q-tooltip class="bg-white text-primary">{{ t('close') }}</q-tooltip>
        </q-btn>
      </q-bar>
      <q-card-section>
        <div class="text-h6">{{t('create')}}</div>
      </q-card-section>
      <q-card-section>
        <CreateCustomersProduct :reloadEvent="reloadCustomersProductList" />
      </q-card-section>

    </q-card>
  </q-dialog>

</template>

<script setup>

import {computed, onMounted, ref} from 'vue'
import {useQuasar} from 'quasar'
import { useI18n } from 'vue-i18n'
import {useRouter} from "vue-router";
import CustomersProductDetail from "components/customers_product/CustomersProductDetail.vue";
import UpdateCustomersProduct from "components/customers_product/UpdateCustomersProduct.vue";
import DeleteCustomersProduct from "components/customers_product/DeleteCustomersProduct.vue";
import CreateCustomersProduct from "components/customers_product/CreateCustomersProduct.vue";
import CreateCustomer from "components/customer/CreateCustomer.vue";
import dayJS from "boot/date";

const {locale,t} = useI18n()
const $q = useQuasar()
const rows = ref([])
const customerProducts = ref([])
const customersProduct_id_props = ref()
const router = useRouter()

const columns = computed(()=>[
  // {name: 'ID', label: t('id'), align: 'left', field: 'id', sortable: true},
  {name: 'CustomerID', label: t('customerID'), align: 'left', field: 'customerID'},
  {name: 'ProductID', label: t('productID'), align: 'center', field: 'productID'},
  {name: 'LicenseType', label: t('licenseType'), align: 'center', field: 'licenseType'},
  {name: 'ExpireAt', label: t('expireAt'), align: 'center', field: 'expireAt'},
  {name: 'Status', label: t('status'), align: 'center', field: 'isActive'},
  {name: 'Action', label: '', field: 'action'},
])

let customersProductDetailDialog = ref(false)
let deleteCustomersProductDialog = ref(false)
let updateCustomersProductDialog = ref(false)
let createCustomersProductDialog = ref(false)
let customersProductLogsDialog = ref(false)

onMounted(async () => {
  await getCustomersProducts()
});

function formatTimestamp(timestamp) {
  if (timestamp === -1) return 'N/A'; // Handle special case where timestamp is -1

  const date = new Date(timestamp);

  if (locale.value === 'fa-IR') {

    return dayJS.unix(timestamp).format('YYYY/MM/DD')
  }
  if (locale.value === 'en-US') {
    var year = date.toLocaleString("default", {year: "numeric"});
    var month = date.toLocaleString("default", {month: "2-digit"});
    var day = date.toLocaleString("default", {day: "2-digit"});

    return year + "/" + month + "/" + day
  }
}

const formattedData = computed(() => {
  return rows.value.map(item => ({
    ...item,
    expireAt: formatTimestamp(item.expireAt),
  }));
});

async function reloadCustomersProductList(){
  console.log("255 : ")
  let index = rows.value.length
  rows.value.splice(0,index)
  console.log("258 : ", rows.value)

  await getCustomersProducts()
}

async function getCustomersProducts() {

  const response = await fetch(
      '/api/panel/customer-product/list',
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
  customerProducts.value = await response.json();
  rows.value = await replaceCustomerIdsWithNames(customerProducts.value);
  rows.value = await replaceProductIdsWithNames(rows.value);

}

async function getProduct(productID) {
  const response = await fetch(
      '/api/panel/product/get',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              id: productID,
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

  return  await response.json();
}

async function getCustomer(customerID) {
  const response = await fetch(
      '/api/panel/customer/get',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              "id": customerID,
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

  return  await response.json();
}

async function replaceCustomerIdsWithNames(customerProducts) {
  const customerNameCache = new Map();

  async function getCustomerName(customerID) {
    if (customerNameCache.has(customerID)) {
      return customerNameCache.get(customerID);
    }

    const customerData = await getCustomer(customerID);
    const customerName = customerData.name; // Assuming the response JSON contains a name field
    customerNameCache.set(customerID, customerName);
    return customerName;
  }

  const updatedCustomerProducts = await Promise.all(customerProducts.map(async (product) => {
    const customerName = await getCustomerName(product.customerID);
    return {
      ...product,
      customerID: customerName,
    };
  }));

  return updatedCustomerProducts;
}

async function replaceProductIdsWithNames(customerProducts) {
  if (!Array.isArray(customerProducts)) {
    throw new TypeError('customerProducts is not an array');
  }

  const productNameCache = new Map();

  async function getProductName(productId) {
    if (productNameCache.has(productId)) {
      return productNameCache.get(productId);
    }

    const productData = await getProduct(productId);
    const productName = productData.name;
    productNameCache.set(productId, productName);
    return productName;
  }

  const updatedCustomerProducts = await Promise.all(customerProducts.map(async (product) => {
    const productName = await getProductName(product.productID);
    return {
      ...product,
      productID: productName,
    };
  }));

  return updatedCustomerProducts;
}

function openCreateCustomersProductDialog() {
  createCustomersProductDialog.value = !createCustomersProductDialog.value
}

function openUpdateCustomersProductDialog(customersProduct_id) {
  updateCustomersProductDialog.value = !updateCustomersProductDialog.value
  customersProduct_id_props.value = customersProduct_id
}

function openCustomersProductDetailDialog(customersProduct_id) {
  customersProductDetailDialog.value = !customersProductDetailDialog.value
  customersProduct_id_props.value = customersProduct_id
}

function openDeleteCustomersProductDialog(customersProduct_id) {
  deleteCustomersProductDialog.value = !deleteCustomersProductDialog.value
  customersProduct_id_props.value = customersProduct_id
}

</script>
