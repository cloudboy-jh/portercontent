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
    text-decoration: none;
    transition: transform 0.2s ease, box-shadow 0.2s ease;
  }

  .btn.primary {
    background: linear-gradient(135deg, var(--accent-500), #7cf6ff);
    color: #041015;
    box-shadow: 0 12px 30px var(--glow);
  }

  .btn.ghost {
    border: 1px solid #2a2a2a;
    color: var(--ink-200);
  }

  .btn:hover {
    transform: translateY(-2px);
  }
</style>
