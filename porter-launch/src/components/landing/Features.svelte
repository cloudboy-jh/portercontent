<script lang="ts">
  export let features: Array<{ title: string; description: string; icon: string }>;
</script>

<section class="features">
  <div class="container">
    <p class="eyebrow">Why Porter</p>
    <h2>One orchestrator, zero lock-in</h2>
    <div class="grid">
      {#each features as feature}
        <article class="card">
          <div class="icon">{feature.icon}</div>
          <h3>{feature.title}</h3>
          <p>{feature.description}</p>
        </article>
      {/each}
    </div>
  </div>
</section>

<style>
  .features {
    padding: 4rem 0;
  }

  h2 {
    margin: 0.6rem 0 2rem;
  }

  .grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
    gap: 1.5rem;
  }

  .card {
    padding: 1.8rem;
    border-radius: 16px;
    background: #111214;
    border: 1px solid #1f1f1f;
  }

  .icon {
    font-size: 1.4rem;
    margin-bottom: 0.5rem;
  }

  p {
    color: var(--ink-400);
  }
</style>
