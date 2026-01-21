# Porter Content Strategy
**Launch:** February 24, 2026 | **Sprint:** 34 days

---

## Core Narrative

**The Problem:**
Developers are drowning in AI agent platforms. Every tool wants to be the orchestrator. State is scattered. Workflows are brittle. Vendor lock-in is real.

**The Porter Solution:**
Use what you already have: GitHub Issues. Agents run locally. You control everything. One @mention, that's it.

**The Positioning:**
Porter is the anti-platform platform. We're not another AI dev tool trying to own your workflow. We're infrastructure that disappears into your existing process.

---

## Target Audience

**Primary:** Individual developers who:
- Use GitHub daily
- Experiment with Claude Code, Cursor, or other agents
- Value local execution and control
- Prefer terminal-native tools
- Distrust vendor lock-in

**Secondary:** Small engineering teams (3-10 people) who:
- Already use GitHub Issues for project management
- Want to automate repetitive coding tasks
- Need audit trails and transparency
- Have limited budget for SaaS tools

**Not targeting (yet):**
- Enterprise (too complex, too slow)
- Non-technical users (GitHub-native requires git comfort)
- Teams without GitHub workflows

---

## Voice & Tone

**Voice Attributes:**
- **Technical but accessible** - No dumbing down, but no jargon for jargon's sake
- **Honest** - Admit limitations, don't overpromise
- **Confident but humble** - We built something useful, not revolutionary
- **Direct** - No marketing fluff, no buzzwords

**Example Good:**
"Porter uses GitHub Issues as its state store. No new platform to learn. Just @porter in your issue comments."

**Example Bad:**
"Revolutionize your development workflow with our cutting-edge AI orchestration paradigm!"

**Comparables:**
- Cloudflare's technical blog (smart, direct)
- Wails documentation (clear, no BS)
- Your own Pact README (terminal-native ethos)

---

## Content Pillars

### 1. Education (Docs Primary)
**Goal:** Make Porter immediately understandable and usable

**Content Types:**
- Quickstart guides (5 min to first success)
- Concept explanations (how/why it works)
- API reference (complete, accurate)
- Real-world examples (CI debugging, code review, etc.)

**Tone:** Patient teacher, not condescending

**Success Metric:** User goes from install → first successful @porter mention in <10 minutes

---

### 2. Demonstration (Landing Page + Social)
**Goal:** Show, don't tell

**Content Types:**
- Animated terminal examples (landing page hero)
- Screenshot workflows (social posts)
- Before/after comparisons (problem vs Porter solution)
- Real execution logs (not fake demos)

**Tone:** "Here's what it actually looks like"

**Success Metric:** Visitor understands value prop in <30 seconds

---

### 3. Credibility (Social + Docs)
**Goal:** Build trust through technical depth

**Content Types:**
- Architecture explanations (why local execution matters)
- Honest limitations (what Porter can't/won't do)
- Open development process (build in public vibes)
- Technical threads (deep dives on specific decisions)

**Tone:** Peer-to-peer, developer-to-developer

**Success Metric:** Target audience sees Porter as "built by someone who gets it"

---

### 4. Community (Post-Launch)
**Goal:** Amplify user wins, build momentum

**Content Types:**
- User success stories (with permission)
- Community-contributed examples
- Tips & tricks from early adopters
- Changelog highlights (what we shipped, why)

**Tone:** Celebration without hype

**Success Metric:** Users become Porter advocates organically

---

## Content Distribution

### Documentation (docs.porter.dev)
**Primary Goal:** Enable successful usage

**Structure:**
1. **Getting Started** - Install → First run → Success
2. **Concepts** - Architecture, design decisions, mental models
3. **Configuration** - .porter.yml, agent setup, auth
4. **API Reference** - @porter commands, labels, webhooks
5. **Examples** - Real workflows with copy-paste code
6. **Guides** - Team setup, security, troubleshooting

**Update Cadence:**
- Pre-launch: Daily iterations
- Post-launch: Weekly additions based on user questions
- Always: Immediate fixes for inaccuracies

**Quality Bar:**
- Every command has an example
- Every concept has a "why"
- Every guide is tested end-to-end
- No placeholder content at launch

---

### Landing Page (porter.dev)
**Primary Goal:** Convert visitor → early access signup

**Sections:**
1. **Hero** - Value prop + visual demo (15 seconds to "I get it")
2. **Problem/Solution** - Contrast current chaos with Porter simplicity
3. **Features** - 4 core differentiators (GitHub-native, local, universal, simple)
4. **How It Works** - 5-step flow with terminal examples
5. **Use Cases** - 4 real scenarios ("Debug CI overnight", etc.)
6. **Signup** - Email + GitHub username → waitlist

**Copy Principles:**
- Lead with developer pain, not Porter features
- Use code examples, not marketing speak
- Every claim has evidence (screenshot, code, or docs link)
- No superlatives ("revolutionary", "game-changing", etc.)

**Design Principles:**
- Dark theme (developer aesthetic)
- Code-forward (terminal visuals, syntax highlighting)
- Fast (<2s load, minimal JS)
- Modal.com vibes (clean, technical, confident)

---

### Social Media (Twitter Primary)
**Primary Goal:** Build awareness + credibility in AI dev community

**Pre-Launch (3-5 posts over 34 days):**

**Post 1 (Days 1-7): The Problem**
```
Debugging a CI failure at 11pm.

Could just @porter investigate-ci-failure and 
have it fixed by morning.

Building that. Local execution, GitHub as 
orchestrator. Launch Feb 24.
```
**Goal:** Establish pain point relatability

**Post 2 (Days 10-15): The Approach**
```
Porter uses GitHub Issues as its state store.

No new platform. No new UI.
Just @porter in issue comments.

Agent runs locally, reports back natively.

[Link to docs]
```
**Goal:** Explain technical approach, link to docs

**Post 3 (Days 18-22): Demo**
```
[Screenshot: GitHub issue with @porter mention → 
terminal execution → PR opened]

This is Porter. Available Feb 24.

[Link to landing page]
```
**Goal:** Visual proof, drive signups

**Post 4 (Launch Day: Feb 24)**
```
Porter is live.

GitHub-native agent orchestrator.
Local execution. Universal compatibility.

Docs: [link]
Download: [link]

Let's see what you build.
```
**Goal:** Announce, provide access, invite engagement

**Post 5 (Post-Launch: Weekly)**
```
First Porter user win:

@username used it to [specific achievement].

[Link to their issue/PR]

This is what local-first orchestration enables.
```
**Goal:** Amplify community, show real usage

**Posting Strategy:**
- Manual posting (not automated)
- Engage with every reply
- Retweet user wins
- No scheduling tools (stay present)

**What NOT to do:**
- Daily posting (too spammy)
- Hype language ("revolutionary", "game-changing")
- Comparison attacks (dunking on competitors)
- Engagement bait (polls, "thoughts?", etc.)

---

## Content Production Workflow

### Week 1-2: Foundation
1. Write core docs (Quickstart, Concepts, Config)
2. Draft landing page copy (all sections)
3. Generate social posts 1-2 (via CLI)
4. Review for voice consistency

### Week 3: Polish
1. Add docs examples (real workflows)
2. Landing page animations (terminal demo)
3. Generate social post 3 (demo)
4. Docs search integration

### Week 4: Final Push
1. Complete API reference
2. Landing page performance optimization
3. Social post 4 (launch)
4. Deploy staging sites for review

### Launch Week:
1. Final docs review (accuracy check)
2. Deploy production sites
3. Post launch announcement
4. Monitor feedback, rapid iteration

---

## Content Quality Checklist

**Documentation:**
- [ ] Every command/feature has example
- [ ] Every example is tested and works
- [ ] No placeholder or "coming soon" content
- [ ] Mobile-friendly (code blocks readable)
- [ ] Search works (all pages indexed)
- [ ] Load time <1s

**Landing Page:**
- [ ] Value prop clear in <15 seconds
- [ ] CTA visible above fold
- [ ] All links functional
- [ ] Forms submit correctly (test with real email)
- [ ] Animations smooth on mobile
- [ ] Load time <2s

**Social Posts:**
- [ ] Under 280 characters
- [ ] No typos/grammar errors
- [ ] Links tested and correct
- [ ] Screenshots/videos high quality
- [ ] Tone matches voice guide

---

## Post-Launch Content Roadmap

### Week 1-4 Post-Launch:
- **Social:** Weekly user wins + feature highlights
- **Docs:** Add FAQ section based on user questions
- **Blog (optional):** Technical deep-dive on architecture

### Month 2-3:
- **Social:** Community tips & tricks
- **Docs:** Advanced guides (plugin dev, custom agents)
- **Blog:** "Building Porter" series (if there's interest)

### Month 4+:
- **Social:** Case studies (with user permission)
- **Docs:** Integration guides (CI/CD, team workflows)
- **Blog:** Quarterly "State of Porter" updates

---

## Metrics & Success

**Pre-Launch:**
- Docs complete: 100% (all sections)
- Landing page ready: 100% (all sections functional)
- Social posts generated: 4-5 (reviewed and scheduled)

**Launch Day:**
- Early access signups: 50+ (baseline success)
- Social engagement: 10+ replies/retweets
- Docs traffic: 100+ unique visitors

**Week 1 Post-Launch:**
- Active users: 10+ (using Porter in real projects)
- Community feedback: 5+ GitHub issues/discussions
- Docs updates: 3+ (based on user questions)

**Month 1 Post-Launch:**
- Signups: 200+
- Active users: 50+
- User-generated content: 5+ (tweets, blog posts, etc.)

---

## Open Questions

1. **Blog?**
   - Add blog.porter.dev for technical deep-dives?
   - Or keep it simple (docs + social only)?
   - Lean: No blog until there's demand

2. **Video Content?**
   - Record demo videos for landing page?
   - Or just animated terminal GIFs?
   - Lean: Start with GIFs, add video if needed

3. **Community Platform?**
   - Discord server?
   - GitHub Discussions only?
   - Lean: GitHub Discussions (keeps everything in one place)

4. **Changelog?**
   - Separate changelog page?
   - Or just GitHub releases?
   - Lean: GitHub releases (no extra maintenance)

---

## Content Principles (Summary)

1. **Docs are source of truth** - Social and landing derive from docs
2. **Show, don't tell** - Code examples > marketing copy
3. **Technical credibility > hype** - Honest limitations, real workflows
4. **Manual > automated** - Especially during launch (stay engaged)
5. **Developer-first** - Write for people who build, not buy
