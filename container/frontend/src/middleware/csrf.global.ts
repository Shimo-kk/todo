import { defineNuxtRouteMiddleware } from '#app'
import { useCsrfStore } from '@/store/csrfStore'

export default defineNuxtRouteMiddleware(async (to, from) => {
  const csrfStore = useCsrfStore()

  if (csrfStore.token === '') {
    await csrfStore.fetch()
  }
})
