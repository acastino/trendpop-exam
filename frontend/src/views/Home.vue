<script lang="ts" setup>
  import { ref, reactive, onMounted } from 'vue';
  import InvoiceApi from '@/services/InvoiceApi'
  import InvoiceType from '@/models/InvoiceType';

  let showForm = ref(true);

  let inputs = reactive<InvoiceType>({
    workspace: 'Collab',
    subscription: 'Team',
    billing_amount: '$1,000,000',
    billing_period: 'Jan 11, 2023 - Feb 11, 2023',
    po_number: ''
  });

  let rules = {
    required: (value: any) => !!value
  };

  let form:any = ref(null);
  onMounted(() => {
    if (!form.value) return;
    form.value.validate();
  });

  let loading = ref(false);
  async function onSubmit() {
    const validationResult = await form.value.validate();
    if (!validationResult.valid) {
      return;
    }

    loading.value = true;
    const response = await InvoiceApi.create(inputs);
    loading.value = false;

    if (!response.ok) {
      console.error('API response: ', response);
      alert("Something went wrong.\nPlease check the Console logs for details.");
      return;
    }

    console.info('Saved into the Database successfully: ', response);
    showForm.value = false;
  }

  let rootPath = process.env.NODE_ENV === "production" ? "/trendpop-exam/" : "/"
</script>

<template>
  <v-form @submit.prevent="onSubmit" ref="form">
    <v-card :loading="loading" class="mx-auto my-8" min-width="555" max-width="555" min-height="650" elevation="8" >
      <v-card-text class="pa-16">

        <fieldset :disabled="loading" v-show="showForm">
          <div class="text-left text-h5 mb-11">Please provide your PO # for the invoice</div>
          <v-text-field label="Workspace name" variant="outlined" v-model="inputs.workspace" readonly></v-text-field>
          <v-text-field label="Subscription plan" variant="outlined" v-model="inputs.subscription" readonly></v-text-field>
          <v-text-field label="Billing amount" variant="outlined" v-model="inputs.billing_amount" readonly></v-text-field>
          <v-text-field label="Billing period" variant="outlined" v-model="inputs.billing_period" readonly></v-text-field>
          <v-text-field label="PO #" variant="outlined" v-model="inputs.po_number" :rules="[rules.required]"></v-text-field>
          <div class="mt-n5">
            <small class="text-left">This PO# will be attached to your invoice. Please check your PO# before submitting as this cannot be changed afterwards.</small>
          </div>
          <v-btn type="submit" variant="flat" block class="submit my-6 py-6">Submit PO #</v-btn>
        </fieldset>

        <div v-show="!showForm">
          <div class="text-left text-h5 mb-11">Thank you for submitting your PO#!</div>
        </div>

        <div class="text-left text-h6">Go back to <a :href="rootPath">Workspace Settings</a></div>
      </v-card-text>
    </v-card>
  </v-form>
</template>

<style scoped>
  fieldset {
    border: none;
  }

  .submit {
    background-color: #000;
    color: #fff;
  }
</style>