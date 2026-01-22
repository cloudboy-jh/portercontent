<script lang="ts">
  const lines = [
    "$ porter init",
    "Connected to github.com/yourorg/project",
    "",
    "> @porter investigate-ci-failure",
    "",
    "Porter: analyzing CI logs",
    "Found flaky test in auth.test.ts",
    "Fix committed: PR #423"
  ];
</script>

<div class="terminal-demo">
  {#each lines as line}
    <div class="line">{line}</div>
  {/each}
</div>

<style>
  .terminal-demo {
    font-family: "JetBrains Mono", monospace;
    font-size: 0.9rem;
    color: #c9d1d9;
    display: grid;
    gap: 0.5rem;
  }

  .line {
    white-space: pre-wrap;
  }
</style>
