<script lang="ts">
  let email = "";
  let github = "";
  let submitted = false;

  function handleSubmit() {
    submitted = true;
  }
</script>

<section class="signup" id="signup">
  <div class="container">
    <div class="surface">
      {#if !submitted}
        <div class="content">
          <div class="copy">
            <p class="eyebrow">Early access</p>
            <h2>Join the launch list</h2>
            <p>Launch day: Feb 24, 2026. Be the first to try Porter.</p>
          </div>
          <form on:submit|preventDefault={handleSubmit}>
            <div class="fields">
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
            </div>
            <button type="submit">Get early access</button>
          </form>
        </div>
      {:else}
        <div class="success">
          <div class="success-icon">✓</div>
          <h3>You're on the list</h3>
          <p>We'll email you when Porter opens for early access.</p>
        </div>
      {/if}
    </div>
  </div>
</section>

<style>
  .signup {
    padding: 4rem 0 6rem;
  }

  .surface {
    background: var(--surface);
    border: 1px solid var(--border);
    border-radius: 20px;
    padding: 2.5rem;
    box-shadow: 0 4px 32px rgba(0, 0, 0, 0.2);
  }

  .content {
    display: grid;
    gap: 2rem;
  }

  .copy {
    text-align: center;
  }

  .copy p:not(.eyebrow) {
    color: var(--foreground-muted);
    margin: 0.5rem 0 0;
  }

  h2 {
    font-size: clamp(1.5rem, 4vw, 1.875rem);
    font-weight: 700;
    letter-spacing: -0.02em;
    margin: 0.75rem 0 0;
    color: var(--foreground);
  }

  form {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    max-width: 480px;
    margin: 0 auto;
    width: 100%;
  }

  .fields {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 1rem;
  }

  input {
    padding: 0.875rem 1rem;
    border-radius: 10px;
    background: var(--background);
    border: 1px solid var(--border-subtle);
    color: var(--foreground);
    font-size: 0.9375rem;
    transition: all 0.2s ease;
  }

  input::placeholder {
    color: var(--foreground-subtle);
  }

  input:focus {
    outline: none;
    border-color: var(--primary-400);
    box-shadow: 0 0 0 3px rgba(251, 146, 60, 0.1);
  }

  button {
    padding: 0.875rem 1.5rem;
    border-radius: 999px;
    border: none;
    background: var(--primary-400);
    color: #1c1917;
    font-weight: 600;
    font-size: 0.9375rem;
    cursor: pointer;
    transition: all 0.2s ease;
  }

  button:hover {
    background: var(--primary-300);
    transform: translateY(-1px);
    box-shadow: 0 4px 16px var(--glow);
  }

  button:active {
    transform: translateY(0);
  }

  .success {
    text-align: center;
    padding: 2rem 0;
  }

  .success-icon {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 56px;
    height: 56px;
    border-radius: 50%;
    background: rgba(251, 146, 60, 0.1);
    border: 1px solid rgba(251, 146, 60, 0.2);
    color: var(--primary-400);
    font-size: 1.5rem;
    font-weight: 700;
    margin-bottom: 1rem;
  }

  .success h3 {
    font-size: 1.25rem;
    font-weight: 600;
    margin: 0 0 0.5rem;
    color: var(--foreground);
  }

  .success p {
    color: var(--foreground-muted);
    margin: 0;
  }

  @media (max-width: 640px) {
    .surface {
      padding: 1.75rem;
    }

    .fields {
      grid-template-columns: 1fr;
    }

    button {
      width: 100%;
    }
  }
</style>
