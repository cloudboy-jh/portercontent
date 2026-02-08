<script lang="ts">
  export let href = "#";
  export let variant: "primary" | "ghost" = "primary";
</script>

<a href={href} class={`btn ${variant}`}>
  <slot />
</a>

<style>
  .btn {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    padding: 0.85rem 1.6rem;
    border-radius: 999px;
    font-weight: 600;
    font-size: 0.95rem;
    text-decoration: none;
    transition: all 0.2s ease;
    letter-spacing: -0.01em;
  }

  .btn.primary {
    background: var(--primary-400);
    color: #1c1917;
    box-shadow: 0 4px 16px var(--glow);
    border: none;
  }

  .btn.primary:hover {
    background: var(--primary-300);
    box-shadow: 0 6px 24px var(--glow-strong);
    transform: translateY(-1px);
  }

  .btn.ghost {
    border: 1px solid var(--border-subtle);
    color: var(--foreground-muted);
    background: transparent;
  }

  .btn.ghost:hover {
    border-color: var(--primary-400);
    color: var(--primary-400);
    background: rgba(251, 146, 60, 0.05);
  }

  .btn:active {
    transform: translateY(0);
  }
</style>
