<script lang="ts">
  export let code: string;
  export let language = "bash";
  let copied = false;

  async function copyCode() {
    await navigator.clipboard.writeText(code);
    copied = true;
    setTimeout(() => (copied = false), 2000);
  }
</script>

<div class="code-block">
  <div class="header">
    <span>{language}</span>
    <button type="button" on:click={copyCode}>{copied ? "Copied" : "Copy"}</button>
  </div>
  <pre><code>{code}</code></pre>
</div>

<style>
  .code-block {
    border-radius: 12px;
    border: 1px solid #1e2630;
    overflow: hidden;
    margin: 1.5rem 0;
  }

  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.5rem 0.8rem;
    background: #111821;
    font-size: 0.8rem;
  }

  button {
    background: transparent;
    border: 1px solid #2d3743;
    border-radius: 6px;
    color: var(--ink-300);
    padding: 0.25rem 0.6rem;
    cursor: pointer;
  }
</style>
