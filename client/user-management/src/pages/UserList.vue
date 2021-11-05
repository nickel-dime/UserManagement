<template>
  <div>
    <section>
      <DataTable
        :value="users"
        :scrollable="true"
        editMode="row"
        v-model:selection="selectedUser"
        v-model:editingRows="editingUsers"
        @rowEditInit="onRowEditInit"
        @rowEditCancel="onRowEditCancel"
        @rowEditSave="onRowEditComplete"
        selectionMode="single"
      >
        <template #loading>
          Loading users, please wait...
        </template>
        <template #empty>
          No users found
        </template>
        <template #header>
          <div
            class="table-header p-d-flex p-flex-column p-flex-md-row p-jc-md-between"
          >
            <h3 class="p-mb-2 p-m-md-0 p-as-md-center">Manage Users</h3>
            <Button
              label="New"
              icon="pi pi-plus"
              class="p-button-success p-mr-2"
              @click="this.$router.push('register')"
            />
          </div>
        </template>
        <Column field="name" header="Name" class="p-col"
          ><template #editor="slotProps">
            <InputText
              v-model="slotProps.data[slotProps.column.props.field]"
              autofocus
            /> </template
        ></Column>
        <Column field="age" header="Age" class="p-col">
          <template #editor="slotProps">
            <InputText
              type="number"
              v-model="slotProps.data[slotProps.column.props.field]"
            />
          </template>
        </Column>
        <Column field="registeredAt" header="Registered At" class="p-col">
        </Column>
        <Column
          :rowEditor="true"
          bodyStyle="text-align:center"
          class="p-col"
          style="max-width:100px"
        >
        </Column>
        <Column>
          <template #body="item">
            <Button
              @click="deleteDeveloper(item.data)"
              class="pi pi-trash"
            ></Button>
          </template>
        </Column>
      </DataTable>
    </section>
  </div>
</template>

<script>
export default {
  data() {
    return {
      selectedUser: null,
      users: null,
      editingUsers: [],
      isLoading: false,
    };
  },
  originalRows: null,
  computed: {
    hasUsers() {
      return this.$store.getters.hasUsers;
    },
  },
  watch: {
    selectedUser(newUser) {
      this.$router.push({
        name: "users",
        params: {
          id: newUser.id,
        },
      });
    },
    $route(value) {
      console.log(value);
      this.loadUsers();
      console.log(this.users);
    },
  },
  methods: {
    onRowEditInit(event) {
      this.originalRows[event.index] = { ...this.users[event.index] };
    },
    onRowEditCancel(event) {
      this.users[event.index] = this.originalRows[event.index];
    },
    async onRowEditComplete(event) {
      if (!this.validateUser(event.data)) {
        this.users[event.index] = this.originalRows[event.index];
      } else {
        try {
          await this.$store.dispatch("editUser", event.data);
        } catch (error) {
          this.error = error.message || "Something went wrong!";
        }
        this.loadUsers();
      }
    },
    validateUser(user) {
      if (user.name && user.age) {
        if (user.name != "" && !Number.isNaN(user.age) && user.age > 0) {
          return true;
        }
      }
      return false;
    },
    async loadUsers(refresh = false) {
      this.isLoading = true;
      try {
        await this.$store.dispatch("loadUsers", {
          forceRefresh: refresh,
        });
      } catch (error) {
        this.error = error.message || "Something went wrong!";
      }
      this.isLoading = false;
      this.users = this.$store.getters.users;
    },
    deleteDeveloper(user) {
      this.$store.dispatch("removeUser", user);
      this.users = this.$store.getters.users;
    },
    beforeRouteUpdate(to, from, next) {
      console.log(to + " " + from + " " + next);
      this.loadUsers();
    },
  },
  // updated() {
  //   this.loadUsers();
  //   this.originalRows = {};
  // },
  created() {
    this.loadUsers();
    this.originalRows = {};
  },
};
</script>
