<script lang="ts">
  import { GithubLogo, Sun, Moon } from 'phosphor-svelte';
  import { onMount } from 'svelte';
  
  let isDark = true;
  
  onMount(() => {
    const savedTheme = localStorage.getItem('theme');
    if (savedTheme) {
      isDark = savedTheme === 'dark';
    } else if (document.documentElement.classList.contains('light')) {
      isDark = false;
    } else {
      isDark = !window.matchMedia('(prefers-color-scheme: light)').matches;
    }
    applyTheme();
  });
  
  function toggleTheme() {
    isDark = !isDark;
    localStorage.setItem('theme', isDark ? 'dark' : 'light');
    applyTheme();
  }
  
  function applyTheme() {
    if (isDark) {
      document.documentElement.classList.add('dark');
      document.documentElement.classList.remove('light');
    } else {
      document.documentElement.classList.add('light');
      document.documentElement.classList.remove('dark');
    }
  }
</script>

<header class="site-header">
  <div class="container">
    <div class="brand">
      <a href="/" class="logo-link">
        <img src="/images/porter-header.png" alt="Porter" class="logo-img" />
      </a>
    </div>
    <nav>
      <a href="https://porter.sh" target="_blank" rel="noopener noreferrer">Web App</a>
      <a href="/docs">Docs</a>
      <a href="https://github.com" target="_blank" rel="noopener noreferrer" class="github-link">
        <GithubLogo size={18} weight="fill" />
        <span>GitHub</span>
      </a>
      <button class="theme-toggle" on:click={toggleTheme} aria-label="Toggle theme">
        {#if isDark}
          <Moon size={18} weight="fill" />
        {:else}
          <Sun size={18} weight="fill" />
        {/if}
      </button>
    </nav>
  </div>
</header>

<style>
  .site-header {
    padding: 1rem 0;
    border-bottom: 1px solid var(--border);
    backdrop-filter: blur(20px);
    background: var(--surface);
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
    gap: 0.875rem;
  }

  .logo-link {
    display: flex;
    align-items: center;
  }

  .logo-img {
    height: 32px;
    width: auto;
    display: block;
  }

  nav {
    display: flex;
    align-items: center;
    gap: 2rem;
    font-size: 0.9rem;
    font-weight: 500;
  }

  nav a {
    text-decoration: none;
    color: var(--foreground-muted);
    transition: color 0.2s ease;
    position: relative;
  }

  nav a::after {
    content: "";
    position: absolute;
    bottom: -4px;
    left: 0;
    width: 0;
    height: 2px;
    background: var(--primary-400);
    transition: width 0.2s ease;
  }

  nav a:hover {
    color: var(--foreground);
  }

  nav a:hover::after {
    width: 100%;
  }

  .github-link {
    display: flex;
    align-items: center;
    gap: 0.375rem;
  }

  .theme-toggle {
    background: none;
    border: none;
    cursor: pointer;
    padding: 0.5rem;
    color: var(--foreground-muted);
    display: flex;
    align-items: center;
    transition: color 0.2s ease;
    border-radius: 6px;
  }

  .theme-toggle:hover {
    color: var(--foreground);
    background: var(--surface-elevated);
  }

  @media (max-width: 720px) {
    nav {
      display: none;
    }

    .logo-img {
      height: 28px;
    }
  }
</style>
