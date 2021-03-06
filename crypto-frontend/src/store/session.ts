import { Route } from "vue-router";
import { Module } from "vuex";
import { User } from "@/models/models";
import { GenericAPIResponse, fetchAuthenticated } from "@/lib/http/api";
import { AppRootState } from "@/store";
import { defineAbilitiesFor } from "@/lib/user/abilities";

export interface SessionState {
  user: User | null;
  loading: Boolean;
  error: any;
  userRequested: Boolean;
  savedRoute: Route | null;
}

const sessionModule: Module<SessionState, AppRootState> = {
  namespaced: true,
  state: () => ({
    user: null,
    loading: false,
    error: null,
    userRequested: false,
    savedRoute: null
  }),
  getters: {
    user: s => s.user,
    loading: s => s.loading,
    error: s => s.error,
    savedRoute: s => s.savedRoute,
    userRequested: s => s.userRequested
  },
  actions: {
    retrieveUser: ({ commit }) => {
      commit("setLoading", true);
      commit("setError", null);
      fetchAuthenticated(`auth/current`)
        .then(response => {
          if (response.ok) {
            return response
              .json()
              .then((responseJSON: GenericAPIResponse<User>) => {
                if (responseJSON.data) {
                  const userInSession = responseJSON.data;
                  userInSession.ability = defineAbilitiesFor(userInSession);
                  commit("setUser", userInSession);
                }
              });
          }
          throw response;
        })
        .catch((err: Response) => {
          commit("setError", err);
        })
        .finally(() => {
          commit("setUserRequested", true);
          commit("setLoading", false);
        });
    },
    setUser: ({ commit }, payload: User) => {
      commit("setUser", payload);
    },
    clearSavedRoute: ({ commit }) => {
      commit("setSavedRoute", null);
    },
    saveRoute: ({ commit }, payload: Route) => {
      commit("setSavedRoute", payload);
    }
  },
  mutations: {
    setUser: (s, payload: User) => {
      s.user = payload;
    },
    setUserRequested: (s, payload: Boolean) => {
      s.userRequested = payload;
    },
    setSavedRoute: (s, payload: Route | null) => {
      s.savedRoute = payload;
    },
    setLoading: (s, payload: Boolean) => {
      s.loading = payload;
    },
    setError: (s, payload: any) => {
      s.error = payload;
    }
  }
};

export default sessionModule;
