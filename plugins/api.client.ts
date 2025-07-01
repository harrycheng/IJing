import { createRouter } from 'vue-router';
import { defineNuxtPlugin } from '#app';
import axios from 'axios';

export default defineNuxtPlugin((nuxtApp) => {
  const apiClient = axios.create({
    baseURL: process.env.API_BASE_URL || '/api',
    timeout: 5000,
  });

  nuxtApp.provide('api', apiClient);
});