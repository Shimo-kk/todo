<template>
  <div class="min-h-screen flex flex-col bg-gray-200 dark:bg-neutral-700">
    <LayoutHeader />
    <main class="flex-1">
      <slot />
    </main>
    <LayoutFooter />
  </div>
</template>
