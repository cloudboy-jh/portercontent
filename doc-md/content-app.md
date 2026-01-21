# Porter Content Application
**Tech Stack:** Astro + Svelte + Tailwind + Go CLI

---

## Architecture Overview

```
porter-launch/
├── astro/                    # Astro monorepo (landing + docs)
│   ├── src/
│   │   ├── pages/
│   │   ├── components/
│   │   ├── content/
│   │   └── layouts/
│   └── astro.config.mjs
└── social-gen/               # Go CLI for social content
    ├── main.go
    ├── templates/
    └── context.json
```

**Key Decision:** One Astro repo for both landing and docs
- Shared components, styles, and tooling
- Single build process
- Deploy to two CF Pages projects (or one with routing)

---

## Astro Monorepo

### Setup

```bash
npm create astro@latest porter-launch
cd porter-launch
npx astro add svelte tailwind
npm install pagefind @astrojs/sitemap
```

### Directory Structure

```
porter-launch/
├── src/
│   ├── pages/
│   │   ├── index.astro                    # Landing page root
│   │   └── docs/
│   │       ├── index.astro                # Docs home
│   │       └── [...slug].astro            # Dynamic doc routes
│   │
│   ├── components/
│   │   ├── shared/
│   │   │   ├── Header.svelte              # Shared nav
│   │   │   ├── Footer.svelte              # Shared footer
│   │   │   └── Button.svelte              # Reusable button
│   │   │
│   │   ├── landing/
│   │   │   ├── Hero.svelte                # Animated hero section
│   │   │   ├── Features.svelte            # 4-card feature grid
│   │   │   ├── HowItWorks.svelte          # Step-by-step flow
│   │   │   ├── UseCases.svelte            # Real scenario cards
│   │   │   ├── SignupForm.svelte          # Early access form
│   │   │   └── TerminalDemo.svelte        # Animated code example
│   │   │
│   │   └── docs/
│   │       ├── Sidebar.svelte             # Doc navigation
│   │       ├── TableOfContents.svelte     # Right sidebar TOC
│   │       ├── CodeBlock.svelte           # Syntax highlighted code
│   │       ├── SearchBar.svelte           # Pagefind integration
│   │       └── Breadcrumbs.svelte         # Navigation breadcrumbs
│   │
│   ├── content/
│   │   ├── config.ts                      # Content collections schema
│   │   │
│   │   ├── docs/                          # Markdown docs
│   │   │   ├── getting-started/
│   │   │   │   ├── quickstart.md
│   │   │   │   ├── installation.md
│   │   │   │   └── first-run.md
│   │   │   │
│   │   │   ├── concepts/
│   │   │   │   ├── architecture.md
│   │   │   │   ├── github-state.md
│   │   │   │   └── local-execution.md
│   │   │   │
│   │   │   ├── configuration/
│   │   │   │   ├── porter-yml.md
│   │   │   │   ├── agent-setup.md
│   │   │   │   └── authentication.md
│   │   │   │
│   │   │   ├── api/
│   │   │   │   ├── commands.md
│   │   │   │   ├── labels.md
│   │   │   │   └── webhooks.md
│   │   │   │
│   │   │   ├── examples/
│   │   │   │   ├── ci-debugging.md
│   │   │   │   ├── code-review.md
│   │   │   │   ├── feature-implementation.md
│   │   │   │   └── database-migration.md
│   │   │   │
│   │   │   └── guides/
│   │   │       ├── team-setup.md
│   │   │       ├── security.md
│   │   │       └── troubleshooting.md
│   │   │
│   │   └── marketing/                     # Landing page data
│   │       ├── hero.json                  # Hero section copy
│   │       ├── features.json              # Feature cards
│   │       └── use-cases.json             # Use case scenarios
│   │
│   ├── layouts/
│   │   ├── LandingLayout.astro            # Landing page wrapper
│   │   ├── DocsLayout.astro               # Docs page wrapper
│   │   └── BaseLayout.astro               # Shared base (SEO, fonts, etc.)
│   │
│   └── styles/
│       ├── global.css                     # Global styles
│       └── syntax-theme.css               # Code highlighting theme
│
├── public/
│   ├── fonts/                             # Self-hosted fonts
│   ├── images/                            # Static images
│   └── favicon.ico
│
├── astro.config.mjs                       # Astro configuration
├── tailwind.config.mjs                    # Tailwind configuration
├── tsconfig.json                          # TypeScript config
└── package.json
```

### Configuration Files

**astro.config.mjs:**
```javascript
import { defineConfig } from 'astro/config';
import svelte from '@astrojs/svelte';
import tailwind from '@astrojs/tailwind';
import sitemap from '@astrojs/sitemap';

export default defineConfig({
  site: 'https://porter.dev',
  integrations: [
    svelte(),
    tailwind(),
    sitemap(),
  ],
  output: 'static',
  build: {
    inlineStylesheets: 'auto',
  },
  vite: {
    build: {
      rollupOptions: {
        output: {
          manualChunks: {
            'landing': ['./src/components/landing'],
            'docs': ['./src/components/docs'],
          },
        },
      },
    },
  },
});
```

**content/config.ts:**
```typescript
import { defineCollection, z } from 'astro:content';

const docsCollection = defineCollection({
  type: 'content',
  schema: z.object({
    title: z.string(),
    description: z.string(),
    category: z.enum([
      'getting-started',
      'concepts',
      'configuration',
      'api',
      'examples',
      'guides'
    ]),
    order: z.number().optional(),
    updated: z.date().optional(),
  }),
});

const marketingCollection = defineCollection({
  type: 'data',
  schema: z.object({
    hero: z.object({
      headline: z.string(),
      subhead: z.string(),
      cta: z.string(),
    }).optional(),
    features: z.array(z.object({
      title: z.string(),
      description: z.string(),
      icon: z.string(),
    })).optional(),
    useCases: z.array(z.object({
      title: z.string(),
      description: z.string(),
      example: z.string(),
    })).optional(),
  }),
});

export const collections = {
  'docs': docsCollection,
  'marketing': marketingCollection,
};
```

---

## Landing Page Components

### Hero.svelte
```svelte
<script lang="ts">
  import { onMount } from 'svelte';
  import { fade, fly } from 'svelte/transition';
  
  let terminalLines = [
    '$ porter init',
    'Connected to github.com/yourorg/project',
    '$ # Comment on any issue with @porter',
    '',
    '> @porter investigate-ci-failure',
    '',
    '🤖 Porter: Analyzing CI logs...',
    '✓ Found: Flaky test in auth.test.ts',
    '✓ Fix committed: PR #423',
  ];
  
  let visibleLines = 0;
  
  onMount(() => {
    const interval = setInterval(() => {
      if (visibleLines < terminalLines.length) {
        visibleLines++;
      } else {
        clearInterval(interval);
      }
    }, 400);
    
    return () => clearInterval(interval);
  });
</script>

<section class="hero">
  <div class="container">
    <h1 in:fade>GitHub Issues as Your Agent Orchestrator</h1>
    <p in:fade={{ delay: 200 }}>
      Mention @porter in any issue. Agents run locally. 
      Results commit back.
    </p>
    
    <div class="terminal" in:fly={{ y: 20, delay: 400 }}>
      {#each terminalLines.slice(0, visibleLines) as line}
        <div class="line">{line}</div>
      {/each}
    </div>
    
    <a href="#signup" class="cta" in:fade={{ delay: 800 }}>
      Join Early Access
    </a>
  </div>
</section>

<style>
  .hero {
    min-height: 80vh;
    display: flex;
    align-items: center;
    background: linear-gradient(180deg, #0a0a0a 0%, #1a1a1a 100%);
  }
  
  .terminal {
    background: #0d1117;
    border: 1px solid #30363d;
    border-radius: 8px;
    padding: 1.5rem;
    font-family: 'JetBrains Mono', monospace;
    font-size: 0.9rem;
    color: #c9d1d9;
    max-width: 600px;
    margin: 2rem auto;
  }
  
  .line {
    margin: 0.5rem 0;
  }
  
  /* Additional styling... */
</style>
```

### Features.svelte
```svelte
<script lang="ts">
  export let features: Array<{
    title: string;
    description: string;
    icon: string;
  }>;
</script>

<section class="features">
  <div class="container">
    <h2>Why Porter?</h2>
    
    <div class="grid">
      {#each features as feature}
        <div class="card">
          <div class="icon">{feature.icon}</div>
          <h3>{feature.title}</h3>
          <p>{feature.description}</p>
        </div>
      {/each}
    </div>
  </div>
</section>

<style>
  .grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 2rem;
    margin-top: 3rem;
  }
  
  .card {
    background: #161b22;
    border: 1px solid #30363d;
    border-radius: 8px;
    padding: 2rem;
    transition: transform 0.2s;
  }
  
  .card:hover {
    transform: translateY(-4px);
  }
  
  /* Additional styling... */
</style>
```

### SignupForm.svelte
```svelte
<script lang="ts">
  let email = '';
  let github = '';
  let loading = false;
  let submitted = false;
  
  async function handleSubmit() {
    loading = true;
    
    try {
      const response = await fetch('/api/signup', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, github }),
      });
      
      if (response.ok) {
        submitted = true;
      }
    } catch (error) {
      console.error('Signup failed:', error);
    } finally {
      loading = false;
    }
  }
</script>

<section class="signup" id="signup">
  <div class="container">
    {#if !submitted}
      <h2>Join Early Access</h2>
      <p>Launch: February 24, 2026</p>
      
      <form on:submit|preventDefault={handleSubmit}>
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
        <button type="submit" disabled={loading}>
          {loading ? 'Submitting...' : 'Get Early Access'}
        </button>
      </form>
    {:else}
      <div class="success">
        <h3>✓ You're on the list!</h3>
        <p>We'll email you on launch day.</p>
      </div>
    {/if}
  </div>
</section>

<style>
  form {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    max-width: 400px;
    margin: 2rem auto;
  }
  
  input {
    padding: 1rem;
    background: #0d1117;
    border: 1px solid #30363d;
    border-radius: 6px;
    color: #c9d1d9;
    font-size: 1rem;
  }
  
  button {
    padding: 1rem 2rem;
    background: #238636;
    border: none;
    border-radius: 6px;
    color: white;
    font-weight: 600;
    cursor: pointer;
    transition: background 0.2s;
  }
  
  button:hover:not(:disabled) {
    background: #2ea043;
  }
  
  button:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }
  
  /* Additional styling... */
</style>
```

---

## Docs Components

### Sidebar.svelte
```svelte
<script lang="ts">
  export let sections: Array<{
    category: string;
    pages: Array<{ title: string; slug: string }>;
  }>;
  
  export let currentSlug: string;
</script>

<aside class="sidebar">
  <nav>
    {#each sections as section}
      <div class="section">
        <h3>{section.category}</h3>
        <ul>
          {#each section.pages as page}
            <li>
              <a
                href={`/docs/${page.slug}`}
                class:active={currentSlug === page.slug}
              >
                {page.title}
              </a>
            </li>
          {/each}
        </ul>
      </div>
    {/each}
  </nav>
</aside>

<style>
  .sidebar {
    position: sticky;
    top: 2rem;
    width: 250px;
    height: calc(100vh - 4rem);
    overflow-y: auto;
    border-right: 1px solid #30363d;
    padding-right: 2rem;
  }
  
  .section {
    margin-bottom: 2rem;
  }
  
  h3 {
    font-size: 0.75rem;
    text-transform: uppercase;
    letter-spacing: 0.05em;
    color: #8b949e;
    margin-bottom: 0.5rem;
  }
  
  ul {
    list-style: none;
    padding: 0;
  }
  
  a {
    display: block;
    padding: 0.5rem 0.75rem;
    color: #c9d1d9;
    text-decoration: none;
    border-radius: 6px;
    transition: background 0.2s;
  }
  
  a:hover {
    background: #161b22;
  }
  
  a.active {
    background: #161b22;
    color: #58a6ff;
  }
</style>
```

### CodeBlock.svelte
```svelte
<script lang="ts">
  export let code: string;
  export let language: string = 'bash';
  
  let copied = false;
  
  async function copyCode() {
    await navigator.clipboard.writeText(code);
    copied = true;
    setTimeout(() => copied = false, 2000);
  }
</script>

<div class="code-block">
  <div class="header">
    <span class="language">{language}</span>
    <button on:click={copyCode} class="copy">
      {copied ? '✓ Copied' : 'Copy'}
    </button>
  </div>
  <pre><code>{code}</code></pre>
</div>

<style>
  .code-block {
    background: #0d1117;
    border: 1px solid #30363d;
    border-radius: 6px;
    margin: 1.5rem 0;
    overflow: hidden;
  }
  
  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.75rem 1rem;
    background: #161b22;
    border-bottom: 1px solid #30363d;
  }
  
  .language {
    font-size: 0.75rem;
    color: #8b949e;
    text-transform: uppercase;
  }
  
  .copy {
    background: transparent;
    border: 1px solid #30363d;
    color: #c9d1d9;
    padding: 0.25rem 0.75rem;
    border-radius: 4px;
    font-size: 0.75rem;
    cursor: pointer;
  }
  
  pre {
    padding: 1rem;
    margin: 0;
    overflow-x: auto;
  }
  
  code {
    font-family: 'JetBrains Mono', monospace;
    font-size: 0.875rem;
    color: #c9d1d9;
  }
</style>
```

---

## Social Content Generator (Go CLI)

### Directory Structure

```
social-gen/
├── main.go
├── cmd/
│   ├── generate.go
│   └── publish.go
├── templates/
│   ├── problem.md
│   ├── approach.md
│   ├── demo.md
│   └── launch.md
├── context.json
├── output/
│   └── posts.json
└── go.mod
```

### main.go
```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Context struct {
	ProductName string   `json:"product_name"`
	LaunchDate  string   `json:"launch_date"`
	KeyFeatures []string `json:"key_features"`
	PainPoints  []string `json:"pain_points"`
}

type Post struct {
	Content       string    `json:"content"`
	SuggestedDate string    `json:"suggested_date"`
	CharCount     int       `json:"char_count"`
	Template      string    `json:"template"`
	Type          string    `json:"type"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: porter-social [generate|publish]")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "generate":
		generatePosts()
	case "publish":
		publishPosts()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}

func generatePosts() {
	// Load context
	context, err := loadContext("context.json")
	if err != nil {
		fmt.Printf("Error loading context: %v\n", err)
		os.Exit(1)
	}

	// Load templates
	templates := []string{"problem.md", "approach.md", "demo.md", "launch.md"}
	posts := []Post{}

	launchDate, _ := time.Parse("2006-01-02", context.LaunchDate)
	now := time.Now()
	daysUntilLaunch := int(launchDate.Sub(now).Hours() / 24)

	for i, tmpl := range templates {
		content, err := loadTemplate(tmpl)
		if err != nil {
			fmt.Printf("Error loading template %s: %v\n", tmpl, err)
			continue
		}

		// Fill in template variables
		content = fillTemplate(content, context)

		// Calculate suggested date
		var suggestedDate time.Time
		switch i {
		case 0: // problem - early
			suggestedDate = now.AddDate(0, 0, daysUntilLaunch/6)
		case 1: // approach - mid
			suggestedDate = now.AddDate(0, 0, daysUntilLaunch/2)
		case 2: // demo - late
			suggestedDate = now.AddDate(0, 0, int(float64(daysUntilLaunch)*0.75))
		case 3: // launch
			suggestedDate = launchDate
		}

		post := Post{
			Content:       content,
			SuggestedDate: suggestedDate.Format("2006-01-02"),
			CharCount:     len(content),
			Template:      tmpl,
			Type:          strings.TrimSuffix(tmpl, ".md"),
		}

		posts = append(posts, post)
	}

	// Write output
	outputPath := "output/posts.json"
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	data, err := json.MarshalIndent(posts, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling posts: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(outputPath, data, 0644); err != nil {
		fmt.Printf("Error writing output: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Generated %d posts\n", len(posts))
	fmt.Printf("✓ Output written to %s\n", outputPath)
	
	// Print summary
	for _, post := range posts {
		fmt.Printf("\n%s (%s, %d chars):\n%s\n", 
			post.Type, post.SuggestedDate, post.CharCount, post.Content)
	}
}

func loadContext(path string) (Context, error) {
	var ctx Context
	data, err := os.ReadFile(path)
	if err != nil {
		return ctx, err
	}
	err = json.Unmarshal(data, &ctx)
	return ctx, err
}

func loadTemplate(name string) (string, error) {
	path := filepath.Join("templates", name)
	data, err := os.ReadFile(path)
	return string(data), err
}

func fillTemplate(template string, ctx Context) string {
	replacements := map[string]string{
		"{{.ProductName}}":  ctx.ProductName,
		"{{.LaunchDate}}":   ctx.LaunchDate,
		"{{.PainPoint}}":    ctx.PainPoints[0],
		"{{.KeyFeature}}":   ctx.KeyFeatures[0],
		"{{.ExampleMention}}": "@porter investigate-ci-failure",
	}

	result := template
	for placeholder, value := range replacements {
		result = strings.ReplaceAll(result, placeholder, value)
	}

	return result
}

func publishPosts() {
	// Optional: integrate with Typefully API
	fmt.Println("Publish functionality not yet implemented")
	fmt.Println("Manually copy posts from output/posts.json")
}
```

### context.json
```json
{
  "product_name": "Porter",
  "launch_date": "2026-02-24",
  "key_features": [
    "GitHub-native orchestration",
    "Local execution model",
    "Universal agent compatibility",
    "Simple @mention interface"
  ],
  "pain_points": [
    "Managing multiple AI agent platforms",
    "Scattered state across tools",
    "Vendor lock-in concerns",
    "Complex workflow setup"
  ]
}
```

### templates/problem.md
```markdown
{{.PainPoint}}.

What if you could just {{.ExampleMention}} and have it done?

Building that. {{.KeyFeature}}. Launch {{.LaunchDate}}.
```

### templates/approach.md
```markdown
{{.ProductName}} uses GitHub Issues as its state store.

No new platform. No new UI.
Just @porter in issue comments.

Agent runs locally, reports back natively.
```

### templates/demo.md
```markdown
[Screenshot: GitHub issue with @porter mention → terminal execution → PR opened]

This is {{.ProductName}}. Available {{.LaunchDate}}.
```

### templates/launch.md
```markdown
{{.ProductName}} is live.

GitHub-native agent orchestrator.
Local execution. Universal compatibility.

Docs: [link]
Download: [link]

Let's see what you build.
```

---

## Build & Deploy

### Build Process

```bash
# Install dependencies
npm install

# Development
npm run dev                    # Start dev server (localhost:4321)

# Build for production
npm run build                  # Outputs to dist/

# Preview production build
npm run preview
```

### Cloudflare Pages Deployment

**Option 1: Two Projects (Recommended)**

```bash
# Deploy landing page
wrangler pages deploy dist --project-name=porter-web

# Deploy docs
wrangler pages deploy dist/docs --project-name=porter-docs
```

**Configure custom domains:**
- Landing: porter.dev
- Docs: docs.porter.dev

**Option 2: Single Project with Routing**

```bash
# Deploy entire dist/
wrangler pages deploy dist --project-name=porter

# Configure routes:
# / → landing page
# /docs/* → documentation
```

### CI/CD (GitHub Actions)

```yaml
# .github/workflows/deploy.yml
name: Deploy to Cloudflare Pages

on:
  push:
    branches: [main]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Setup Node
        uses: actions/setup-node@v3
        with:
          node-version: 18
      
      - name: Install dependencies
        run: npm ci
      
      - name: Build
        run: npm run build
      
      - name: Deploy Landing
        uses: cloudflare/pages-action@v1
        with:
          apiToken: ${{ secrets.CF_API_TOKEN }}
          accountId: ${{ secrets.CF_ACCOUNT_ID }}
          projectName: porter-web
          directory: dist
      
      - name: Deploy Docs
        uses: cloudflare/pages-action@v1
        with:
          apiToken: ${{ secrets.CF_API_TOKEN }}
          accountId: ${{ secrets.CF_ACCOUNT_ID }}
          projectName: porter-docs
          directory: dist/docs
```

---

## Performance Optimization

### Bundle Size Targets
- Landing page: < 50KB JS (compressed)
- Docs: < 30KB JS + Pagefind index
- Shared CSS: < 20KB (compressed)
- Images: WebP format, lazy loaded

### Techniques
1. **Code splitting**: Separate landing/docs chunks
2. **Island architecture**: Only hydrate interactive components
3. **Image optimization**: `astro:assets` for automatic optimization
4. **Font subsetting**: Only include used glyphs
5. **Preloading**: Critical resources in `<head>`

### Lighthouse Targets
- Performance: 95+
- Accessibility: 100
- Best Practices: 100
- SEO: 100

---

## Content Update Workflow

### Adding New Docs

1. Create markdown file in appropriate category:
   ```bash
   touch src/content/docs/examples/new-workflow.md
   ```

2. Add frontmatter:
   ```markdown
   ---
   title: "New Workflow Example"
   description: "How to use Porter for X"
   category: "examples"
   order: 5
   updated: 2026-02-15
   ---
   
   Content here...
   ```

3. Build regenerates routes automatically

### Updating Landing Page

1. Edit JSON in `src/content/marketing/`:
   ```json
   {
     "hero": {
       "headline": "Updated headline",
       "subhead": "New subhead",
       "cta": "Get Started"
     }
   }
   ```

2. Or edit components directly in `src/components/landing/`

### Generating Social Posts

```bash
cd social-gen
go run main.go generate

# Review output/posts.json
# Copy content to Twitter manually
```

---

## Testing Checklist

### Pre-Launch Testing

**Landing Page:**
- [ ] All sections render correctly
- [ ] Signup form submits (test with real email)
- [ ] Animations smooth on mobile
- [ ] All links functional
- [ ] Load time < 2s
- [ ] Mobile responsive (test on real device)

**Docs:**
- [ ] All pages accessible via navigation
- [ ] Search works (Pagefind indexed correctly)
- [ ] Code blocks render with syntax highlighting
- [ ] Copy buttons work
- [ ] Mobile readable (code blocks scroll)
- [ ] Load time < 1s

**Social Generator:**
- [ ] Generates all post types
- [ ] Character counts accurate
- [ ] Template variables filled correctly
- [ ] Dates calculated properly

**Cross-Browser:**
- [ ] Chrome (desktop + mobile)
- [ ] Firefox
- [ ] Safari (desktop + mobile)
- [ ] Edge

---

## Post-Launch Maintenance

### Weekly Tasks
- [ ] Monitor signup form submissions
- [ ] Review user questions → add to docs FAQ
- [ ] Generate weekly social post (if relevant)
- [ ] Check analytics (traffic, bounce rate)

### Monthly Tasks
- [ ] Review docs accuracy (outdated examples?)
- [ ] Update changelog
- [ ] Optimize slow pages
- [ ] Add new examples based on user requests

### As Needed
- [ ] Fix broken links
- [ ] Update screenshots/demos
- [ ] Add new features to landing page
- [ ] Expand API reference

---

## Monitoring & Analytics

**Cloudflare Web Analytics** (Recommended)
- Privacy-friendly (no cookies)
- Built into CF Pages
- Tracks: page views, visitors, referrers, devices

**Optional:**
- Plausible (lightweight, privacy-focused)
- PostHog (open-source, self-hostable)

**Metrics to Track:**
- Landing page conversion rate (visits → signups)
- Docs most-viewed pages (what's popular?)
- Search queries (what are users looking for?)
- Bounce rate (are users finding what they need?)

---

## Open Questions

1. **API Endpoint for Signup Form:**
   - Cloudflare Workers function?
   - Supabase Edge Function?
   - Direct Supabase REST API?

2. **Search Implementation:**
   - Pagefind (static, no backend)?
   - Algolia (hosted, requires account)?
   - Custom (Cloudflare KV + Workers)?

3. **Analytics:**
   - CF Web Analytics (free, simple)?
   - Plausible (paid, privacy-focused)?
   - None (truly minimal)?

4. **Social CLI Publishing:**
   - Manual posting only?
   - Typefully API integration?
   - Buffer/Hypefury?

---

## Next Steps

1. **Week 1:**
   - Set up Astro repo with Svelte
   - Create basic landing page structure
   - Draft core docs (Quickstart, Concepts)

2. **Week 2:**
   - Build landing page components
   - Write docs examples
   - Set up social generator CLI

3. **Week 3:**
   - Polish animations
   - Integrate search (Pagefind)
   - Test signup form

4. **Week 4:**
   - Deploy to CF Pages staging
   - Full testing (mobile, cross-browser)
   - Generate social posts

5. **Launch Week:**
   - Deploy production
   - Post launch announcement
   - Monitor feedback
