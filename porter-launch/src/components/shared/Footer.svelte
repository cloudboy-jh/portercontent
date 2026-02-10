<footer class="site-footer">
  <div class="container">
    <div class="brand">
      <div class="logo">
        <img src="/images/porter-header.png" alt="Porter" class="logo-img" />
      </div>
      <p>Cloud-native AI agent orchestrator for GitHub Issues.</p>
    </div>
    <nav class="links">
      <a href="/docs">Docs</a>
      <a href="https://github.com">GitHub</a>
    </nav>
  </div>
  <div class="container bottom">
    <p class="copyright">© 2026 Porter. All rights reserved.</p>
    <div class="meta">
      <a href="/security">Security</a>
      <a href="/privacy">Privacy</a>
      <a href="/status">Status</a>
    </div>
  </div>
</footer>

<style>
  .site-footer {
    padding: 4rem 0 2rem;
    border-top: 1px solid var(--border);
    background: var(--background);
  }

  .container {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    gap: 2rem;
    flex-wrap: wrap;
  }

  .brand {
    max-width: 280px;
  }

  .logo {
    display: flex;
    align-items: center;
    gap: 0.75rem;
  }

  .logo-img {
    height: 28px;
    width: auto;
    display: block;
  }

  .brand p {
    margin: 0.5rem 0 0;
    color: var(--foreground-muted);
    font-size: 0.9375rem;
    line-height: 1.5;
  }

  .links {
    display: flex;
    gap: 2rem;
  }

  .links a {
    color: var(--foreground-muted);
    font-size: 0.9375rem;
    font-weight: 500;
    transition: color 0.2s ease;
  }

  .links a:hover {
    color: var(--primary-400);
  }

  .bottom {
    margin-top: 3rem;
    padding-top: 1.5rem;
    border-top: 1px solid var(--border);
    align-items: center;
  }

  .copyright {
    margin: 0;
    color: var(--foreground-subtle);
    font-size: 0.875rem;
  }

  .meta {
    display: flex;
    gap: 1.5rem;
  }

  .meta a {
    color: var(--foreground-subtle);
    font-size: 0.875rem;
    transition: color 0.2s ease;
  }

  .meta a:hover {
    color: var(--foreground-muted);
  }

  @media (max-width: 640px) {
    .container {
      flex-direction: column;
      align-items: center;
      text-align: center;
    }

    .brand {
      max-width: 100%;
    }

    .logo {
      justify-content: center;
    }

    .links {
      gap: 1.5rem;
    }

    .bottom {
      flex-direction: column-reverse;
      gap: 1rem;
    }
  }
</style>
