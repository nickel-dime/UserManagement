<template>
  <Card style="width: 20rem; margin: 2em">
    <template #title>
      Register User
    </template>
    <template #content>
      <form class="p-mx-auto">
        <div class="p-field p-grid p-mx-auto">
          <span class="p-float-label">
            <InputText
              type="text"
              v-model="name.val"
              :class="{ 'p-invalid': !name.isValid }"
            />
            <label for="inputtext">Name</label>
          </span>
        </div>
        <div class="p-field p-grid p-mx-auto">
          <span class="p-float-label">
            <InputText
              type="number"
              v-model="age.val"
              :class="{ 'p-invalid': !age.isValid }"
            />
            <label for="inputtext">Age</label>
          </span>
        </div>
        <Button @click.prevent="submitForm">Submit Form</Button>
      </form>
    </template>
  </Card>
</template>

<script>
export default {
  data() {
    return {
      name: {
        val: "",
        isValid: true,
      },
      age: {
        val: null,
        isValid: true,
      },
      formIsValid: false,
    };
  },
  methods: {
    validateForm() {
      this.formIsValid = true;
      if (this.name.val === "") {
        this.formIsValid = false;
        this.name.isValid = false;
      }
      if (this.age.val <= 0) {
        this.formIsValid = false;
        this.age.isValid = false;
      }
    },
    async submitForm() {
      this.validateForm();

      if (!this.formIsValid) {
        return;
      }

      try {
        await this.$store.dispatch("registerUser", {
          name: this.name.val,
          age: this.age.val,
        });
      } catch (error) {
        this.error = error.message || "Something went wrong!";
      }
      this.$router.replace("/users");
    },
  },
};
</script>
