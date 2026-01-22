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
    <div class="terminal">
      <slot name="terminal" />
    </div>
  </div>
</section>

<style>
  .hero {
    padding: 6rem 0 4rem;
  }

  .container {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
    gap: 3rem;
    align-items: center;
  }

  h1 {
    font-size: clamp(2.4rem, 4vw, 3.6rem);
    margin: 0.6rem 0;
  }

  .lead {
    color: var(--ink-300);
    font-size: 1.1rem;
  }

  .actions {
    display: flex;
    gap: 1rem;
    flex-wrap: wrap;
    margin-top: 2rem;
  }

  .terminal {
    background: #0f1114;
    border: 1px solid #222c36;
    border-radius: 16px;
    padding: 2rem;
    box-shadow: 0 20px 40px rgba(0, 0, 0, 0.4);
  }
</style>
