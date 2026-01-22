<script lang="ts">
  export let headings: Array<{ depth: number; slug: string; text: string }> = [];
</script>

<aside class="toc">
  <p class="eyebrow">On this page</p>
  <ul>
    {#each headings as heading}
      <li class={`depth-${heading.depth}`}>
        <a href={`#${heading.slug}`}>{heading.text}</a>
      </li>
    {/each}
  </ul>
</aside>

<style>
  .toc {
    position: sticky;
    top: 6rem;
    max-height: calc(100vh - 8rem);
  }

  ul {
    list-style: none;
    padding: 0;
  }

  li {
    margin: 0.4rem 0;
  }

  a {
    color: var(--ink-400);
    text-decoration: none;
    font-size: 0.85rem;
  }

  a:hover {
    color: var(--ink-50);
  }

  .depth-3 {
    margin-left: 1rem;
  }
</style>
