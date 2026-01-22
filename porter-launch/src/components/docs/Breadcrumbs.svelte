<script lang="ts">
  export let items: Array<{ label: string; href: string }> = [];
</script>

<nav class="breadcrumbs">
  {#each items as item, index}
    <a href={item.href}>{item.label}</a>
    {#if index < items.length - 1}
      <span>/</span>
    {/if}
  {/each}
</nav>

<style>
  .breadcrumbs {
    display: flex;
    gap: 0.5rem;
    color: var(--ink-500);
    font-size: 0.85rem;
  }

  a {
    text-decoration: none;
    color: inherit;
  }
</style>
