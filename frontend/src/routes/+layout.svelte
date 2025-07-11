<script lang="ts">
  import { onMount } from 'svelte';
  import { authStore } from '$lib/stores/auth';
  import '../app.css';
  import { Toaster } from 'svelte-french-toast';

  onMount(() => {
    authStore.initialize();
  });
</script>

<Toaster />

<main class="min-h-screen bg-gray-50">
  <slot />
</main>