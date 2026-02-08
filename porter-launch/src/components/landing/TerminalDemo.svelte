<script lang="ts">
  const lines = [
    { prompt: ">", text: "GitHub issue #42: Fix auth redirect bug", highlight: false },
    { prompt: "", text: "", highlight: false },
    { prompt: "▼", text: "@porter opencode fix auth redirect", highlight: true },
    { prompt: "", text: "", highlight: false },
    { prompt: "●", text: "Spawning Fly Machine...", highlight: false },
    { prompt: "●", text: "Cloning repo...", highlight: false },
    { prompt: "●", text: "Running opencode agent...", highlight: false },
    { prompt: "✓", text: "Found bug in auth.ts:134", highlight: false },
    { prompt: "✓", text: "Created PR #423: Fix auth redirect", highlight: false },
    { prompt: "", text: "", highlight: false },
    { prompt: "●", text: "Machine destroyed ($0.03)", highlight: false },
  ];
</script>

<div class="terminal-demo">
  <div class="header">
    <div class="dots">
      <span class="dot red"></span>
      <span class="dot yellow"></span>
      <span class="dot green"></span>
    </div>
    <span class="title">porter — zsh</span>
  </div>
  <div class="content">
    {#each lines as line}
      <div class="line" class:highlight={line.highlight}>
        {#if line.prompt}
          <span class="prompt">{line.prompt}</span>
        {/if}
        <span class="text">{line.text}</span>
      </div>
    {/each}
    <div class="cursor-line">
      <span class="prompt">$</span>
      <span class="cursor"></span>
    </div>
  </div>
</div>

<style>
  .terminal-demo {
    font-family: "JetBrains Mono", monospace;
    font-size: 0.85rem;
    line-height: 1.6;
  }

  .header {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
    padding-bottom: 1rem;
    margin-bottom: 1rem;
    border-bottom: 1px solid var(--border);
    position: relative;
  }

  .dots {
    display: flex;
    gap: 0.375rem;
    position: absolute;
    left: 0;
  }

  .dot {
    width: 10px;
    height: 10px;
    border-radius: 50%;
  }

  .dot.red { background: #ff5f56; }
  .dot.yellow { background: #ffbd2e; }
  .dot.green { background: #27c93f; }

  .title {
    color: var(--foreground-muted);
    font-size: 0.75rem;
  }

  .content {
    display: flex;
    flex-direction: column;
    gap: 0.25rem;
  }

  .line {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.125rem 0;
  }

  .line.highlight {
    background: rgba(251, 146, 60, 0.1);
    border-radius: 4px;
    padding: 0.25rem 0.5rem;
    margin: 0 -0.5rem;
  }

  .prompt {
    color: var(--primary-400);
    font-weight: 500;
    min-width: 1rem;
    flex-shrink: 0;
  }

  .text {
    color: var(--foreground-muted);
  }

  .line.highlight .text {
    color: var(--foreground);
    font-weight: 500;
  }

  .cursor-line {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    margin-top: 0.5rem;
  }

  .cursor {
    display: inline-block;
    width: 8px;
    height: 1.1em;
    background: var(--primary-400);
    animation: blink 1s step-end infinite;
  }

  @keyframes blink {
    50% { opacity: 0; }
  }

  @media (prefers-reduced-motion: reduce) {
    .cursor {
      animation: none;
      opacity: 1;
    }
  }
</style>
