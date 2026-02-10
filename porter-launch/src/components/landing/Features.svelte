<script lang="ts">
  import { Cloud, Robot, Gear, CurrencyDollar } from 'phosphor-svelte';
  
  export let features: Array<{ title: string; description: string; icon: string }>;
  
  const iconMap: Record<string, any> = {
    Cloud,
    Robot,
    Gear,
    CurrencyDollar
  };
</script>

<section class="features">
  <div class="container">
    <header class="section-header">
      <p class="eyebrow">Why Porter</p>
      <h2>One orchestrator, zero lock-in</h2>
      <p class="subhead">Porter standardizes request, execution, and reporting flow across agents while keeping GitHub as your team interface.</p>
    </header>
    <div class="grid">
      {#each features as feature}
        <article class="card">
          <div class="icon-wrapper">
            <svelte:component this={iconMap[feature.icon]} size={24} weight="duotone" color="#fb923c" />
          </div>
          <h3>{feature.title}</h3>
          <p>{feature.description}</p>
        </article>
      {/each}
    </div>
  </div>
</section>

<style>
  .features {
    padding: 5rem 0;
  }

  .section-header {
    margin-bottom: 3rem;
  }

  .subhead {
    margin: 0.8rem 0 0;
    max-width: 68ch;
    color: var(--foreground-muted);
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
    grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
    gap: 1.5rem;
  }

  .card {
    padding: 1.75rem;
    border-radius: 16px;
    background: transparent;
    border: 1px solid var(--border);
    border-top-color: rgba(251, 146, 60, 0.4);
    transition: all 0.2s ease;
  }

  .card:hover {
    border-color: var(--border-subtle);
    background: rgba(255, 255, 255, 0.02);
    transform: translateY(-2px);
  }

  .icon-wrapper {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 48px;
    height: 48px;
    border-radius: 12px;
    background: var(--surface-elevated);
    border: 1px solid var(--border);
    margin-bottom: 1rem;
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
    margin: 0;
  }

  @media (max-width: 640px) {
    .features {
      padding: 3.5rem 0;
    }

    .grid {
      grid-template-columns: 1fr;
    }
  }
</style>
