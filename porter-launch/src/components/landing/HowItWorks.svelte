<script lang="ts">
  export let steps: Array<{ title: string; detail: string }>;
</script>

<section class="how">
  <div class="container">
    <p class="eyebrow">Flow</p>
    <h2>Porter runs where your code already lives</h2>
    <div class="steps">
      {#each steps as step, index}
        <div class="step">
          <span class="index">0{index + 1}</span>
          <div>
            <h3>{step.title}</h3>
            <p>{step.detail}</p>
          </div>
        </div>
      {/each}
    </div>
  </div>
</section>

<style>
  .how {
    padding: 4rem 0;
  }

  .steps {
    display: grid;
    gap: 1.5rem;
  }

  .step {
    display: grid;
    grid-template-columns: auto 1fr;
    gap: 1rem;
    padding: 1.5rem;
    border-radius: 14px;
    background: #101216;
    border: 1px solid #1d2430;
  }

  .index {
    font-size: 1.2rem;
    font-weight: 600;
    color: var(--accent-500);
  }

  p {
    color: var(--ink-400);
  }
</style>
