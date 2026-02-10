<script lang="ts">
  import { AppWindow, Key, ChatCircleText, GitPullRequest } from "phosphor-svelte";

  export let steps: Array<{ title: string; detail: string; outcome?: string }>;

  const icons = [AppWindow, Key, ChatCircleText, GitPullRequest];
</script>

<section class="how">
  <div class="container">
    <header class="section-header">
      <p class="eyebrow">How it works</p>
      <h2>Porter runs where your code already lives</h2>
    </header>
    <div class="steps">
      {#each steps as step, index}
        <div class="step">
          <div class="timeline">
            <span class="index">0{index + 1}</span>
            <span class="icon"><svelte:component this={icons[index] ?? AppWindow} size={16} weight="bold" /></span>
            {#if index < steps.length - 1}
              <div class="connector"></div>
            {/if}
          </div>
          <div class="content">
            <h3>{step.title}</h3>
            <p>{step.detail}</p>
            {#if step.outcome}
              <p class="outcome">{step.outcome}</p>
            {/if}
          </div>
        </div>
      {/each}
    </div>
  </div>
</section>

<style>
  .how {
    padding: 5rem 0;
    background: linear-gradient(to bottom, transparent, rgba(251, 146, 60, 0.02), transparent);
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

  .steps {
    display: grid;
    gap: 1rem;
  }

  .step {
    display: grid;
    grid-template-columns: auto 1fr;
    gap: 1.5rem;
    padding: 1.75rem;
    border-radius: 16px;
    background: var(--surface);
    border: 1px solid var(--border);
    transition: all 0.2s ease;
  }

  .step:hover {
    border-color: var(--primary-400);
    box-shadow: 0 0 30px rgba(251, 146, 60, 0.1);
  }

  .timeline {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 0.4rem;
  }

  .index {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 40px;
    height: 40px;
    border-radius: 50%;
    font-size: 0.875rem;
    font-weight: 700;
    color: var(--primary-400);
    background: rgba(251, 146, 60, 0.1);
    border: 1px solid rgba(251, 146, 60, 0.2);
    flex-shrink: 0;
  }

  .icon {
    display: inline-flex;
    width: 28px;
    height: 28px;
    align-items: center;
    justify-content: center;
    border-radius: 8px;
    color: var(--foreground-muted);
    border: 1px solid var(--border-subtle);
    background: var(--surface-elevated);
  }

  .connector {
    width: 2px;
    height: 100%;
    min-height: 24px;
    background: linear-gradient(to bottom, var(--primary-400), transparent);
    opacity: 0.3;
  }

  .content {
    display: flex;
    flex-direction: column;
    justify-content: center;
  }

  h3 {
    font-size: 1.125rem;
    font-weight: 600;
    margin: 0 0 0.375rem;
    color: var(--foreground);
  }

  p {
    color: var(--foreground-muted);
    font-size: 0.9375rem;
    line-height: 1.6;
    margin: 0;
  }

  .outcome {
    margin-top: 0.6rem;
    font-size: 0.84rem;
    color: var(--foreground-subtle);
  }

  @media (max-width: 640px) {
    .how {
      padding: 3.5rem 0;
    }

    .step {
      gap: 1rem;
      padding: 1.25rem;
    }

    .index {
      width: 36px;
      height: 36px;
      font-size: 0.8125rem;
    }
  }
</style>
