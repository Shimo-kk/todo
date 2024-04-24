<template>
  <header
    class="flex flex-wrap sm:justify-start sm:flex-nowrap w-full bg-white text-sm py-4 dark:bg-neutral-800"
  >
    <nav
      class="max-w-[85rem] w-full mx-auto px-4 sm:flex sm:items-center sm:justify-between"
      aria-label="Global"
    >
      <div class="flex items-center justify-between">
        <NuxtLink
          class="flex-none text-xl font-semibold dark:text-white"
          href="/"
          >To Do</NuxtLink
        >
      </div>
      <div
        class="flex flex-col gap-5 mt-5 sm:flex-row sm:items-center sm:justify-end sm:mt-0 sm:ps-5"
      >
        <NuxtLink
          class="font-medium text-blue-500"
          href="/auth/signup"
          aria-current="page"
          >サインアップ</NuxtLink
        >
        <NuxtLink
          class="font-medium text-gray-600 hover:text-gray-400 dark:text-neutral-400 dark:hover:text-neutral-500"
          href="/auth/signin"
          >サインイン</NuxtLink
        >
      </div>
    </nav>
  </header>
</template>
