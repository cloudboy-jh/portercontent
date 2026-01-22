<script lang="ts">
  const nav = [
    { label: "Docs", href: "/docs" },
    { label: "Roadmap", href: "/docs/roadmap" },
    { label: "GitHub", href: "https://github.com" }
  ];
</script>

<header class="site-header">
  <div class="container">
    <div class="brand">
      <span class="logo">Porter</span>
      <span class="pill">Local-first agents</span>
    </div>
    <nav>
      {#each nav as item}
        <a href={item.href}>{item.label}</a>
      {/each}
    </nav>
  </div>
</header>

<style>
  .site-header {
    padding: 1.5rem 0;
    border-bottom: 1px solid #1f1f1f;
    backdrop-filter: blur(14px);
    background: rgba(5, 5, 5, 0.75);
    position: sticky;
    top: 0;
    z-index: 10;
  }

  .site-header .container {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .brand {
    display: flex;
    align-items: center;
    gap: 1rem;
  }

  .logo {
    font-size: 1.2rem;
    font-weight: 700;
    letter-spacing: 0.02em;
  }

  nav {
    display: flex;
    gap: 1.5rem;
    font-size: 0.95rem;
    color: var(--ink-300);
  }

  nav a {
    text-decoration: none;
    transition: color 0.2s ease;
  }

  nav a:hover {
    color: var(--ink-50);
  }

  @media (max-width: 720px) {
    nav {
      display: none;
    }
  }
</style>
