<template>
  <div class="q-pa-md">
    <q-table
        flat bordered
        title="Logs"
        :rows="rows"
        :columns="columns"
        row-key="name"
    >

      <template v-slot:header="props">
        <q-tr :props="props">
          <q-th auto-width />
          <q-th
              v-for="col in props.cols"
              :key="col.name"
              :props="props"
          >
            {{ col.label }}
          </q-th>
        </q-tr>
      </template>

      <template v-slot:body="props">
        <q-tr :props="props">
          <q-td auto-width>
            <q-btn size="sm" color="brown-5" round dense @click="props.expand = !props.expand" :icon="props.expand ? 'remove' : 'add'" />
          </q-td>
          <q-td
              v-for="col in props.cols"
              :key="col.name"
              :props="props"
          >
            {{ col.value }}
          </q-td>
        </q-tr>
        <q-tr v-show="props.expand" :props="props">
          <q-td colspan="100%">
            <q-list class="q-mb-md">
              <q-item><strong>ID:</strong>&nbsp;  {{ props.row.name }}</q-item>
              <q-separator />
              <q-item><strong>Client Hash:</strong>&nbsp;  {{ props.row.client_hash }}</q-item>
              <q-separator />
              <q-item><strong>Time:</strong>&nbsp;  {{ props.row.time }}</q-item>
              <q-separator />
              <q-item><strong>Level:</strong>&nbsp;  {{ props.row.level }}</q-item>
              <q-separator />
              <q-item><strong>Message:</strong>&nbsp; {{ props.row.message }}</q-item>
              <q-separator />
              <q-item><strong>Group:</strong>&nbsp;  {{ props.row.group }}</q-item>
              <q-separator />
              <q-item><strong>Path:</strong> &nbsp; {{ props.row.path }}</q-item>
              <q-separator />
              <q-item><strong>Line:</strong> &nbsp; {{ props.row.line }}</q-item>
              <q-separator />
              <q-item><strong>Function:</strong> &nbsp; {{ props.row.function }}</q-item>
              <q-separator />
              <q-item><strong>Code:</strong>&nbsp;  {{ props.row.code }}</q-item>
              <q-separator />
              <q-item><strong>Request Body: </strong></q-item>
                <q-list v-for="(value, key) in JSON.parse(props.row.request_body)" :key="key">
                  <q-th>&nbsp;&nbsp;&nbsp;&nbsp;{{ key }}: {{ value}}</q-th>
                </q-list>
            </q-list>
          </q-td>
        </q-tr>
      </template>
    </q-table>
  </div>
</template>

<script>
const columns = [
  {
    name: 'message',
    required: true,
    label: 'message',
    align: 'left',
    field: row => row.message,
    format: val => `${val}`,
    sortable: true
  }
]

const rows = [
  {
    name: '1',
    client_hash: 'abc123',
    time: '2023-07-25T12:34:56Z',
    level: 'info',
    message: 'User login successful',
    group: 'auth',
    path: '/api/login',
    line: 42,
    function: 'loginUser',
    request_body: '{"username":"test","password":"****","username1":"test","password1":"****","username2":"test","password2":"****","username3":"test","password3":"****"}',
    code: 200
  },
  {
    name: '2',
    client_hash: 'def456',
    time: '2023-07-25T13:45:56Z',
    level: 'error',
    message: 'Database connection failed',
    group: 'database',
    path: '/api/data',
    line: 88,
    function: 'fetchData',
    request_body: '{"query":"SELECT *"}',
    code: 500
  },
  {
    name: '3',
    client_hash: 'abc123',
    time: '2023-07-25T12:34:56Z',
    level: 'info',
    message: 'User login successful',
    group: 'auth',
    path: '/api/login',
    line: 42,
    function: 'loginUser',
    request_body: '{"username":"test","password":"****"}',
    code: 200
  },
  {
    name: '4',
    client_hash: 'def456',
    time: '2023-07-25T13:45:56Z',
    level: 'error',
    message: 'Database connection failed',
    group: 'database',
    path: '/api/data',
    line: 88,
    function: 'fetchData',
    request_body: '{"query":"SELECT *"}',
    code: 500
  },
  {
    name: '5',
    client_hash: 'abc123',
    time: '2023-07-25T12:34:56Z',
    level: 'info',
    message: 'User login successful',
    group: 'auth',
    path: '/api/login',
    line: 42,
    function: 'loginUser',
    request_body: '{"username":"test","password":"****"}',
    code: 200
  },
  {
    name: '6',
    client_hash: 'def456',
    time: '2023-07-25T13:45:56Z',
    level: 'error',
    message: 'Database connection failed',
    group: 'database',
    path: '/api/data',
    line: 88,
    function: 'fetchData',
    request_body: '{"query":"SELECT *"}',
    code: 500
  },
  {
    name: '7',
    client_hash: 'abc123',
    time: '2023-07-25T12:34:56Z',
    level: 'info',
    message: 'User login successful',
    group: 'auth',
    path: '/api/login',
    line: 42,
    function: 'loginUser',
    request_body: '{"username":"test","password":"****"}',
    code: 200
  },
  {
    name: '8',
    client_hash: 'def456',
    time: '2023-07-25T13:45:56Z',
    level: 'error',
    message: 'Database connection failed',
    group: 'database',
    path: '/api/data',
    line: 88,
    function: 'fetchData',
    request_body: '{"query":"SELECT *"}',
    code: 500
  }
]

export default {
  setup () {
    return {
      columns,
      rows
    }
  }
}
</script>
