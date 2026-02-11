<script lang="ts">
  export let cases: Array<{ title: string; issue: string; command: string; pr: string }>;
</script>

<section class="use-cases">
  <div class="container">
    <header class="section-header">
      <p class="eyebrow">Use cases</p>
      <h2>What it looks like in practice</h2>
    </header>
    <div class="threads">
      {#each cases as entry}
        <article class="thread-wrap">
          <p class="case-label">{entry.title}</p>
          <div class="thread">
            <header class="thread-header">
              <span class="issue-dot" aria-hidden="true"></span>
              <p class="issue-title">Issue: {entry.issue}</p>
            </header>

            <div class="comment user">
              <p class="meta">@you · 2 minutes ago</p>
              <p class="body">{entry.command}</p>
            </div>

            <div class="comment bot">
              <p class="meta">@porter-bot · just now</p>
              <ul class="status-list" aria-label="Porter status updates">
                <li>● Spawning Fly Machine...</li>
                <li>● Cloning repo, running agent</li>
                <li class="success">✓ {entry.pr}</li>
              </ul>
            </div>
          </div>
        </article>
      {/each}
    </div>
  </div>
</section>

<style>
  .use-cases {
    padding: 5rem 0;
  }

  .section-header {
    margin-bottom: 3rem;
  }

  h2 {
    font-size: clamp(1.75rem, 4vw, 2.25rem);
    font-weight: 700;
    letter-spacing: -0.02em;
    margin: 0.75rem 0 0;
    color: var(--foreground);
  }

  .threads {
    display: grid;
    gap: 1rem;
  }

  .thread-wrap {
    display: grid;
    gap: 0.5rem;
  }

  .case-label {
    margin: 0;
    font-size: 0.75rem;
    text-transform: uppercase;
    letter-spacing: 0.08em;
    color: var(--text-secondary, var(--foreground-subtle));
  }

  .thread {
    border: 1px solid var(--border);
    border-radius: 2px;
    background: var(--bg-surface, var(--surface));
  }

  .thread-header {
    display: flex;
    align-items: center;
    gap: 0.625rem;
    padding: 0.875rem 1rem;
    border-bottom: 1px solid var(--border);
  }

  .issue-dot {
    width: 10px;
    height: 10px;
    border-radius: 50%;
    background: #ef4444;
    flex-shrink: 0;
  }

  .issue-title {
    margin: 0;
    color: var(--text-primary, var(--foreground));
    font-size: 0.875rem;
    font-weight: 600;
  }

  .comment {
    padding: 0.875rem 1rem;
  }

  .bot {
    border-top: 1px solid var(--border);
  }

  .meta {
    margin: 0;
    color: var(--text-secondary, var(--foreground-subtle));
    font-size: 0.75rem;
  }

  .body {
    margin: 0.5rem 0 0;
    color: var(--text-primary, var(--foreground));
    font-size: 0.8125rem;
    line-height: 1.5;
    word-break: break-word;
  }

  .status-list {
    margin: 0.5rem 0 0;
    padding: 0;
    list-style: none;
    color: var(--text-secondary, var(--foreground-muted));
    font-size: 0.8125rem;
    line-height: 1.6;
  }

  .success {
    color: var(--success, #34d399);
  }

  @media (max-width: 640px) {
    .use-cases {
      padding: 3.5rem 0;
    }

    .thread-header,
    .comment {
      padding: 0.75rem;
    }
  }
</style>
