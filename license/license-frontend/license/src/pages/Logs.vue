<template>
  <div class="q-pa-md">

    <q-table
        flat bordered
        :title="$t('logs')"
        :rows="formattedData"
        :columns="columns"
        :rows-per-page-label="$t('recordPerPage')"
        :no-data-label="$t('noRecordFound')"
        row-key="name"
        :filter="filter"
        :loading="false"
    >

      <template v-slot:top-right>

        <q-select

            dense
            v-model="logTitle"
            :options="logOptions"
            :label="$t('filter')"
            emit-value
            map-options
            rounded
            outlined
            debounce="500"
        >
          <template v-slot:append>
            <q-icon name="filter_alt"/>
          </template>
        </q-select>
      </template>

    </q-table>
  </div>

</template>

<script setup>
import {computed, onMounted, ref, watch} from 'vue'
import {useQuasar} from 'quasar'
import {useI18n} from 'vue-i18n'
import {useRouter} from "vue-router";
import jalaali from 'jalaali-js'

const {locale, t} = useI18n()
const $q = useQuasar()
const rows = ref([])
const router = useRouter()
const logTitle = ref('all')

const columns = computed(() => [
  // {name: 'ID', label: t('id'), align: 'left', field: 'id', sortable: true},
  {name: 'Title', label: t('title'), align: 'left', field: 'title'},
  {name: 'Message', label: t('message'), align: 'left', field: 'message'},
  {name: 'Time', label: t('time'), align: 'left', field: 'createdAt'},
  {name: 'Action', label: '', field: 'action'},
])

let logOptions = computed(() =>[
      {
        label: t('all'),
        value: 'all'
      },
      {
        label: t('createUser'),
        value: 'createUser'
      },
      {
        label: t('updateUser'),
        value: 'updateUser'
      },
      {
        label: t('getUser'),
        value: 'getUser'
      },
      {
        label: t('listUser'),
        value: 'listUser'
      },
      {
        label: t('deleteUser'),
        value: 'deleteUser'
      },
      {
        label: t('createProduct'),
        value: 'createProduct'
      },
      {
        label: t('updateProduct'),
        value: 'updateProduct'
      },
      {
        label: t('getProduct'),
        value: 'getProduct'
      },
      {
        label: t('listProduct'),
        value: 'listProduct'
      },
      {
        label: t('deleteProduct'),
        value: 'deleteProduct'
      },
      {
        label: t('createCustomer'),
        value: 'createCustomer'
      },
      {
        label: t('updateCustomer'),
        value: 'updateCustomer'
      },
      {
        label: t('getCustomer'),
        value: 'getCustomer'
      },
      {
        label: t('listCustomer'),
        value: 'listCustomer'
      },
      {
        label: t('deleteCustomer'),
        value: 'deleteCustomer'
      },
      {
        label: t('createCustomerProduct'),
        value: 'createCustomerProduct'
      },
      {
        label: t('updateCustomerProduct'),
        value: 'updateCustomerProduct'
      },
      {
        label: t('getCustomerProduct'),
        value: 'getCustomerProduct'
      },
      {
        label: t('listCustomerProduct'),
        value: 'listCustomerProduct'
      },
      {
        label: t('deleteCustomerProduct'),
        value: 'deleteCustomerProduct'
      },
      {
        label: t('createRestriction'),
        value: 'createRestriction'
      },
      {
        label: t('updateRestriction'),
        value: 'updateRestriction'
      },
      {
        label: t('getRestriction'),
        value: 'getRestriction'
      },
      {
        label: t('listRestriction'),
        value: 'listRestriction'
      },
      {
        label: t('deleteRestriction'),
        value: 'deleteRestriction'
      },
      {
        label: t('checkLicense'),
        value: 'checkLicense'
      }
    ]
)

onMounted(async () => {
  await getLogs();
});

watch(logTitle, async (newValue) => {
  if (newValue === 'all') {
    await getLogs();
  } else {
    await getLogsByTitle(newValue);
  }
});


async function getLogsByTitle(title) {
  let language
  if (locale.value === 'fa-IR') {
    language = 'fa'
  }

  if (locale.value === 'en-US') {
    language = 'en'
  }
  const response = await fetch(
      '/api/panel/logs/list-by-title',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              title: title,
              language:language
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
          position: "top"
        });
      });
    } else if (typeof data.error === 'string') {

      $q.notify({
        color: 'red',
        textColor: 'white',
        message: t(data.error),
        position: "top"
      });
    }

    if (response.status === 403) {
      await router.push({name: 'login'})
    }

    throw new Error(`HTTP error! status: ${response.status}`);
  }

  rows.value = await response.json();
}

async function getLogs() {

  let language
  if (locale.value === 'fa-IR') {
    language = 'fa'
  }

  if (locale.value === 'en-US') {
    language = 'en'
  }

  const response = await fetch(
      '/api/panel/logs/list',
      {
        method: 'POST',
        headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
        body: JSON.stringify({
              limit: 0,
              offset: 0,
              language:language
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
          position: "top"
        });
      });
    } else if (typeof data.error === 'string') {

      $q.notify({
        color: 'red',
        textColor: 'white',
        message: t(data.error),
        position: "top"
      });
    }

    if (response.status === 403) {
      await router.push({name: 'login'})

    }

    throw new Error(`HTTP error! status: ${response.status}`);
  }

  rows.value = await response.json();
}


function formatTimestamp(timestamp) {
  if (timestamp === -1) return 'N/A';

  const date = new Date(timestamp * 1000);

  if (locale.value === 'fa-IR') {
    const jDate = jalaali.toJalaali(date)

    return `${String(date.getHours()).padStart(2, '0')}:${String(date.getMinutes()).padStart(2, '0')}:${String(date.getSeconds()).padStart(2, '0')} - ${jDate.jy}/${jDate.jm}/${jDate.jd}`
  }
  if (locale.value === 'en-US') {
    var year = date.toLocaleString("default", {year: "numeric"});
    var month = date.toLocaleString("default", {month: "2-digit"});
    var day = date.toLocaleString("default", {day: "2-digit"});
    var hour = date.toLocaleString("default", {hour: "2-digit", hour12: false})
    var minute = date.toLocaleString("default", {minute: "numeric"})
    var sec = date.toLocaleString("default", {second: "numeric"})


    return year + "/" + month + "/" + day + "-" + hour + ":" + minute + ":" + sec
  }
}

const formattedData = computed(() => {
  return rows.value.map(item => ({
    ...item,
    title: t(item.title),
    createdAt: formatTimestamp(item.createdAt),
  }));
});

</script>
