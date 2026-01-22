<script lang="ts">
  export let cases: Array<{ title: string; description: string; example: string }>;
</script>

<section class="use-cases">
  <div class="container">
    <p class="eyebrow">Use cases</p>
    <h2>Real workflows, real issues</h2>
    <div class="grid">
      {#each cases as entry}
        <article class="card">
          <h3>{entry.title}</h3>
          <p>{entry.description}</p>
          <div class="example">{entry.example}</div>
        </article>
      {/each}
    </div>
  </div>
</section>

<style>
  .use-cases {
    padding: 4rem 0;
  }

  .grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
    gap: 1.5rem;
  }

  .card {
    background: #0d0f12;
    border: 1px solid #1e242d;
    border-radius: 16px;
    padding: 1.8rem;
  }

  .example {
    margin-top: 1rem;
    font-family: "JetBrains Mono", monospace;
    font-size: 0.85rem;
    color: var(--accent-500);
  }

  p {
    color: var(--ink-400);
  }
</style>
