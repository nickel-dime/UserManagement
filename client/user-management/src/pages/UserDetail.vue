<template>
  <Card style="width: 20rem; margin: 2em">
    <template #content>
      <form class="p-mx-auto">
        <div class="p-field p-grid p-mx-auto">
          <span class="p-float-label p-d-inline"> 
            <h1>{{ selectedUser.name }}</h1>
            <Avatar label="P" class="p-mr-2" size="medium" shape="circle" />
          </span>
        </div>
        <div class="p-field p-grid p-mx-auto">
          <span class="p-float-label">
            <h2>{{ selectedUser.age }}</h2>
          </span>
        </div>
      </form>
    </template>
  </Card>
</template>

<script>
export default {
  // TODO -> inline spans, passing props to user detail (refresh on detail doesnt stay) and fixing the update/vue issue
  // content header fixed my crons/network problme
  // async/await - good practice fixed my detail doens't stay
  // good practice to have only initilaize two varibales in frontend then rest in backend?
  //  
  props: ["id"],
  created() {
    this.selectedUser = this.$store.getters.users.find(
      (user) => user.id === this.id
    );
  },
  data() {
    return {
      selectedUser: null,
    };
  },
};
</script>
