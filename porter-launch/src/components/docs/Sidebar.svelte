<script lang="ts">
  export let sections: Array<{ category: string; pages: Array<{ title: string; slug: string }> }>;
  export let currentSlug: string;
</script>

<aside class="sidebar">
  {#each sections as section}
    <div class="section">
      <h3>{section.category}</h3>
      <ul>
        {#each section.pages as page}
          <li>
            <a href={`/docs/${page.slug}`} class:active={page.slug === currentSlug}>
              {page.title}
            </a>
          </li>
        {/each}
      </ul>
    </div>
  {/each}
</aside>

<style>
  .sidebar {
    position: sticky;
    top: 6rem;
    max-height: calc(100vh - 8rem);
    overflow-y: auto;
    padding-right: 1.5rem;
    border-right: 1px solid #1f1f1f;
  }

  h3 {
    text-transform: uppercase;
    letter-spacing: 0.1em;
    font-size: 0.75rem;
    color: var(--ink-400);
    margin-bottom: 0.5rem;
  }

  ul {
    list-style: none;
    padding: 0;
    margin: 0 0 1.5rem;
  }

  a {
    display: block;
    padding: 0.4rem 0.6rem;
    border-radius: 8px;
    text-decoration: none;
    color: var(--ink-300);
  }

  a.active,
  a:hover {
    background: #141518;
    color: var(--ink-50);
  }
</style>
