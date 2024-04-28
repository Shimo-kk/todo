import { defineStore } from 'pinia'
import { useRuntimeConfig } from '#app'
import type { RuntimeConfig } from 'nuxt/schema'

interface State {
  token: string
}

export const useCsrfStore = defineStore('csrf', {
  state: (): State => ({
    token: '',
  }),
  actions: {
    async fetch() {
      const runtimeConfig: RuntimeConfig = useRuntimeConfig()
      const url: string = runtimeConfig.public.api_url + '/csrf'

      const response: Response = await fetch(url)
      if (response.ok) {
        const data = await response.json()
        this.token = data.csrf
      }
    },
  },
})
