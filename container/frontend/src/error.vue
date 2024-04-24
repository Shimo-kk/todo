<script setup lang="ts">
import { computed } from 'vue'

const error = useError()

const errorLabel = computed(() => {
  if (error?.value?.statusCode === 404) {
    return 'Page Not Found'
  } else if (error?.value?.statusCode === 500) {
    return 'Internal Server Error'
  }
  return ''
})

const errorMessage = computed(() => {
  if (error?.value?.statusCode === 404) {
    return 'お探しのページが見つかりません。'
  } else if (error?.value?.statusCode === 500) {
    return 'サーバーでエラーが発生しました。'
  }
  return ''
})
</script>

<template>
  <div class="min-h-screen flex flex-col bg-gray-200 dark:bg-neutral-700">
    <LayoutHeader />
    <main class="flex-1">
      <div class="py-6 sm:py-8 lg:py-12">
        <div class="px-4 md:px-8">
          <div class="bg-white dark:bg-neutral-800 py-6 sm:py-8 lg:py-12">
            <div class="mx-auto max-w-screen-2xl px-4 md:px-8">
              <div class="flex flex-col items-center">
                <p
                  class="mb-4 text-sm font-semibold text-blue-500 md:text-base"
                >
                  That’s a
                  <span>{{ error?.statusCode }}</span>
                </p>
                <h1
                  class="mb-2 text-center text-2xl font-bold text-gray-800 dark:text-white md:text-3xl"
                >
                  {{ errorLabel }}
                </h1>

                <p
                  class="mb-12 max-w-screen-md text-center text-gray-400 md:text-lg"
                >
                  {{ errorMessage }}
                </p>

                <NuxtLink
                  href="/"
                  class="inline-block rounded-lg bg-gray-200 px-8 py-3 text-center text-sm font-semibold text-gray-500 outline-none ring-indigo-300 transition duration-100 hover:bg-gray-300 focus-visible:ring active:text-gray-700 md:text-base"
                  >Go home</NuxtLink
                >
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
    <LayoutFooter />
  </div>
</template>
