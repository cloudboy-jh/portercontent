<script lang="ts">
  let email = "";
  let github = "";
  let submitted = false;

  function handleSubmit() {
    submitted = true;
  }
</script>

<section class="signup" id="signup">
  <div class="container surface">
    {#if !submitted}
      <div class="copy">
        <p class="eyebrow">Early access</p>
        <h2>Join the launch list</h2>
        <p>Launch day: Feb 24, 2026.</p>
      </div>
      <form on:submit|preventDefault={handleSubmit}>
        <input
          type="email"
          placeholder="your@email.com"
          bind:value={email}
          required
        />
        <input
          type="text"
          placeholder="GitHub username"
          bind:value={github}
          required
        />
        <button type="submit">Get early access</button>
      </form>
    {:else}
      <div class="success">
        <h3>You're on the list</h3>
        <p>We'll email you when Porter opens.</p>
      </div>
    {/if}
  </div>
</section>

<style>
  .signup {
    padding: 4rem 0 6rem;
  }

  .container {
    padding: 2.5rem;
    display: grid;
    gap: 2rem;
  }

  form {
    display: grid;
    gap: 1rem;
  }

  input {
    padding: 0.9rem 1rem;
    border-radius: 10px;
    background: #0b0c0f;
    border: 1px solid #1e1f24;
    color: var(--ink-100);
  }

  button {
    padding: 0.9rem 1rem;
    border-radius: 999px;
    border: none;
    background: var(--accent-500);
    color: #031017;
    font-weight: 600;
    cursor: pointer;
  }

  .success {
    text-align: center;
  }

  p {
    color: var(--ink-400);
  }
</style>
