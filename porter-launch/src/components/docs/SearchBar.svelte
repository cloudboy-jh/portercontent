<script lang="ts">
  let query = "";
</script>

<div class="search">
  <input type="search" bind:value={query} placeholder="Search docs" />
  <span>Pagefind</span>
</div>

<style>
  .search {
    display: flex;
    gap: 0.6rem;
    align-items: center;
    border: 1px solid #1e1f24;
    padding: 0.5rem 0.8rem;
    border-radius: 10px;
    background: #0d0f12;
  }

  input {
    flex: 1;
    background: transparent;
    border: none;
    color: var(--ink-100);
    outline: none;
  }

  span {
    font-size: 0.75rem;
    color: var(--ink-500);
  }
</style>
