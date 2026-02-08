<script lang="ts">
  export let cases: Array<{ title: string; description: string; example: string }>;
</script>

<section class="use-cases">
  <div class="container">
    <header class="section-header">
      <p class="eyebrow">Use cases</p>
      <h2>Real workflows, real issues</h2>
    </header>
    <div class="grid">
      {#each cases as entry}
        <article class="card">
          <h3>{entry.title}</h3>
          <p>{entry.description}</p>
          <div class="example">
            <span class="prompt">$</span>
            <code>{entry.example}</code>
          </div>
        </article>
      {/each}
    </div>
  </div>
</section>

<style>
  .use-cases {
    padding: 5rem 0;
  }

  .section-header {
    margin-bottom: 3rem;
  }

  h2 {
    font-size: clamp(1.75rem, 4vw, 2.25rem);
    font-weight: 700;
    letter-spacing: -0.02em;
    margin: 0.75rem 0 0;
    color: var(--foreground);
  }

  .grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 1.5rem;
  }

  .card {
    background: var(--surface);
    border: 1px solid var(--border);
    border-radius: 16px;
    padding: 1.75rem;
    transition: all 0.2s ease;
  }

  .card:hover {
    border-color: var(--border-subtle);
    background: var(--surface-elevated);
    transform: translateY(-2px);
  }

  h3 {
    font-size: 1.125rem;
    font-weight: 600;
    margin: 0 0 0.5rem;
    color: var(--foreground);
  }

  p {
    color: var(--foreground-muted);
    font-size: 0.9375rem;
    line-height: 1.6;
    margin: 0 0 1.25rem;
  }

  .example {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.75rem 1rem;
    background: rgba(251, 146, 60, 0.05);
    border: 1px solid rgba(251, 146, 60, 0.15);
    border-radius: 8px;
    font-family: "JetBrains Mono", monospace;
    font-size: 0.8125rem;
  }

  .prompt {
    color: var(--primary-400);
    font-weight: 600;
  }

  code {
    color: var(--foreground-muted);
  }

  @media (max-width: 640px) {
    .use-cases {
      padding: 3.5rem 0;
    }

    .grid {
      grid-template-columns: 1fr;
    }
  }
</style>
