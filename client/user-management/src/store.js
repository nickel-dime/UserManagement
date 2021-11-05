import { createStore } from "vuex";

export default createStore({
  state() {
    return {
      users: [
        // {
        //   id: "01",
        //   name: "Nikhil Goel",
        //   age: 18,
        //   date: "2021-07-15T19:13:32.159Z",
        // },
        // {
        //   id: "02",
        //   name: "Taylor Reinke",
        //   age: 20,
        //   date: "2021-07-14T19:13:32.159Z",
        // },
      ],
    };
  },
  mutations: {
    registerUser(state, payload) {
      state.users.push(payload);
    },
    setUsers(state, payload) {
      state.users = payload;
    },
    removeUser(state, payload) {
      var index = state.users.indexOf(payload);
      if (index > -1) {
        state.users.splice(index, 1);
      }
    },
    editUser(state, payload) {
      var user = state.users.find(
        (user) => user.id === payload.id
      )
      var index = state.users.indexOf(user);
      if (index > -1) {
        state.users[index] = payload;
      }
    }
  },
  getters: {
    users(state) {
      return state.users;
    },
    hasUsers(state) {
      return state.users && state.users.length > 0;
    },
    shouldUpdate(state) {
      const lastFetch = state.lastFetch;
      if (!lastFetch) {
        return true;
      }
      const currentTimeStamp = new Date().getTime();
      return (currentTimeStamp - lastFetch) / 1000 > 60;
    },
  },
  actions: {
    async registerUser(context, data) {
      const newUser = {
        // id: new Date().toISOString,
        name: data.name,
        age: Number(data.age),
        // date: new Date(),
      };

      const response = await fetch("http://localhost:8082/api/users", {
        method: "POST",
        body: JSON.stringify(newUser),
        headers: { "Content-Type": "application/json" },
      });
      const responseData = await response.json();

      if (!response.ok) {
        const error = new Error(responseData.message || "Failed to post!");
        throw error;
      }

      context.commit("registerUser", newUser);
    },
    async loadUsers(context, payload) {
      if (!payload.forceRefresh && !context.getters.shouldUpdate) {
        return;
      }

      const response = await fetch(`http://localhost:8082/api/users`);
      const responseData = await response.json();

      if (!response.ok) {
        const error = new Error(responseData.message || "Failed to fetch!");
        throw error;
      }
      const users = [];

      for (const user in responseData.data.users) {
        // const createdUser = {
        //   id: user.id,
        //   name: responseData[key].name,
        //   age: responseData[key].age,
        //   date: responseData[key].date,
        // };
        users.push(responseData.data.users[user]);
      }

      context.commit("setUsers", users);
    },
    async removeUser(context, payload) {
      await fetch(`http://localhost:8082/api/users/${payload.id}/`, {
        method: "DELETE",
      });
      context.commit("removeUser", payload);
    },
    async editUser(context, data) {
      const editedUser = {
        id: data.id,
        name: data.name,
        age: Number(data.age),
        date: data.registeredAt,
      };
      const response = await fetch(`http://localhost:8082/api/users/${editedUser.id}/`, {
        method: "POST",
        body: JSON.stringify(editedUser),
        headers: { "Content-Type": "application/json" },
      });
      const responseData = await response.json();

      if (!response.ok) {
        const error = new Error(responseData.message || "Failed to post!");
        throw error;
      }
      
      context.commit("editUser", editedUser);
    },
  },
});
