<script lang="ts">
  export let headline: string;
  export let subhead: string;
  export let cta: string;
</script>

<section class="hero">
  <div class="container">
    <div class="copy">
      <p class="eyebrow">GitHub-native agents</p>
      <h1>{headline}</h1>
      <p class="lead">{subhead}</p>
      <div class="actions">
        <slot name="cta" />
        <slot name="secondary" />
      </div>
    </div>
    <div class="terminal-wrapper">
      <div class="terminal">
        <slot name="terminal" />
      </div>
    </div>
  </div>
</section>

<style>
  .hero {
    padding: 5rem 0 4rem;
  }

  .container {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 4rem;
    align-items: center;
  }

  .copy {
    max-width: 540px;
  }

  h1 {
    font-size: clamp(2.25rem, 5vw, 3.5rem);
    font-weight: 700;
    line-height: 1.1;
    letter-spacing: -0.02em;
    margin: 1rem 0 1.25rem;
    color: var(--foreground);
  }

  .lead {
    color: var(--foreground-muted);
    font-size: 1.125rem;
    line-height: 1.6;
    margin: 0;
  }

  .actions {
    display: flex;
    gap: 1rem;
    flex-wrap: wrap;
    margin-top: 2.5rem;
  }

  .terminal-wrapper {
    position: relative;
  }

  .terminal-wrapper::before {
    content: "";
    position: absolute;
    inset: -2px;
    border-radius: 18px;
    background: linear-gradient(135deg, var(--primary-400), var(--primary-600));
    opacity: 0.15;
    filter: blur(20px);
  }

  .terminal {
    position: relative;
    background: var(--surface);
    border: 1px solid var(--border);
    border-radius: 16px;
    padding: 1.5rem;
  }

  /* Responsive */
  @media (max-width: 900px) {
    .container {
      grid-template-columns: 1fr;
      gap: 3rem;
    }

    .copy {
      max-width: 100%;
      text-align: center;
    }

    .actions {
      justify-content: center;
    }

    .terminal-wrapper {
      max-width: 600px;
      margin: 0 auto;
    }
  }

  @media (max-width: 480px) {
    .hero {
      padding: 3rem 0 2.5rem;
    }

    .actions {
      flex-direction: column;
      align-items: center;
    }

    .actions :global(.btn) {
      width: 100%;
      max-width: 280px;
    }
  }
</style>
