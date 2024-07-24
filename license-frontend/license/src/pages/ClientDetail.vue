<template>
  <q-page padding>
    <q-card class="q-pa-md">
      <q-card-section>
        <div class="text-h5">Client Details</div>
      </q-card-section>

      <q-separator />

      <q-card-section>
        <q-list bordered padding class="rounded-borders">
          <q-item>
            <q-item-section>
              <q-item-label><strong>ID</strong></q-item-label>
              <q-item-label caption>{{ client.id }}</q-item-label>
            </q-item-section>
          </q-item>

          <q-item v-if="client.softwareName">
            <q-item-section>
              <q-item-label><strong>Software Name</strong></q-item-label>
              <q-item-label caption>{{ client.softwareName }}</q-item-label>
            </q-item-section>
          </q-item>

          <q-item v-if="client.softwareVersion">
            <q-item-section>
              <q-item-label><strong>Software Version</strong></q-item-label>
              <q-item-label caption>{{ client.softwareVersion }}</q-item-label>
            </q-item-section>
          </q-item>

          <q-item v-if="client.hardwareHash">
            <q-item-section>
              <q-item-label><strong>Hardware Hash</strong></q-item-label>
              <q-item-label caption>{{ client.hardwareHash }}</q-item-label>
            </q-item-section>
          </q-item>

          <q-item v-if="client.licenseType">
            <q-item-section>
              <q-item-label><strong>License Type</strong></q-item-label>
              <q-item-label caption>{{ client.licenseType }}</q-item-label>
            </q-item-section>
          </q-item>

          <q-item v-if="client.userMetadata">
            <q-item-section>
              <q-item-label><strong>User Metadata</strong></q-item-label>
              <q-item-label caption>{{ client.userMetadata }}</q-item-label>
            </q-item-section>
          </q-item>

          <q-item>
            <q-item-section>
              <q-item-label><strong>Is Active</strong></q-item-label>
              <q-item-label caption>{{ client.isActive ? 'Yes' : 'No' }}</q-item-label>
            </q-item-section>
          </q-item>

          <q-item>
            <q-item-section>
              <q-item-label><strong>Expires At</strong></q-item-label>
              <q-item-label caption>{{ formatDate(client.expiresAt) }}</q-item-label>
            </q-item-section>
          </q-item>

          <q-item>
            <q-item-section>
              <q-item-label><strong>Created At</strong></q-item-label>
              <q-item-label caption>{{ formatDate(client.createdAt) }}</q-item-label>
            </q-item-section>
          </q-item>

          <q-item>
            <q-item-section>
              <q-item-label><strong>Updated At</strong></q-item-label>
              <q-item-label caption>{{ formatDate(client.updatedAt) }}</q-item-label>
            </q-item-section>
          </q-item>

          <q-item v-if="client.deletedAt">
            <q-item-section>
              <q-item-label><strong>Deleted At</strong></q-item-label>
              <q-item-label caption>{{ formatDate(client.deletedAt) }}</q-item-label>
            </q-item-section>
          </q-item>
        </q-list>
      </q-card-section>
    </q-card>
  </q-page>

</template>


<script setup>

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

function formatDate(timestamp) {
  if (!timestamp) return 'N/A';
      const date = new Date(timestamp * 1000);
      return date.toLocaleString();
}

</script>

<style scoped>
.q-card {
  background: #ffffff;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.q-list {
  background: #f9f9f9;
  border-radius: 8px;
}

.q-item {
  padding: 16px;
  border-bottom: 1px solid #e0e0e0;
}

.q-item:last-child {
  border-bottom: none;
}

.text-h5 {
  font-size: 1.5rem;
  font-weight: bold;
  color: #333;
}

.q-item-label {
  color: #555;
}

.q-item-label strong {
  color: #333;
}
</style>
