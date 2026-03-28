# Groupie Tracker 🎸

A high-performance web application built with **Go** that visualizes music artist data from a RESTful API — focusing on backend architecture, data transformation, and client-server communication patterns.

---

## 🚀 Features

- **Dynamic Data Visualization** — Transforms raw JSON from an external API into a fully interactive, themed UI across Artists, Members, Locations, Dates, and Relations.
- **Client-Server Search Engine** — A real request-response cycle: the browser dispatches a `fetch` to `/search`, Go processes and filters the data, and the client re-renders the grid dynamically.
- **In-Memory Cache** — The full artist and relations datasets are loaded once at startup into typed Go maps, eliminating all redundant external API calls at runtime.
- **Resilient Architecture** — Server fails fast at startup if the external API is unreachable, then runs stably with zero live dependencies.
- **Keyboard-first UI** — Full navigation without a mouse: `/` to search, `↑↓` to move between cards, `Esc` to clear, `B` to go back.
- **Responsive UI** — "Dark Gold" interface built with CSS Grid, Flexbox, custom animations, and Shneiderman's 8 Golden Rules of Interface Design.

---

## 🛠️ Tech Stack

| Layer | Choice |
|---|---|
| Backend | Go — standard library only (`net/http`, `encoding/json`, `html/template`) |
| Frontend | HTML5, CSS3, Vanilla JS |
| Data | Groupie Trackers REST API |
| Templating | Go `html/template` with server-side rendering |

---

## 🧠 Technical Highlights

### In-Memory Cache with O(1) Relation Lookup

The naive implementation would call the external API on every request — slow and rate-limited. Instead, `InitCache()` runs at startup and populates two global structures:

```go
var (
    ArtistCache   []Artist
    RelationCache map[int]RelationItem  // keyed by artist ID
)
```

Artists are stored as a slice for ordered rendering. Relations are stored as a `map[int]RelationItem` so that `ArtistHandler` can do `RelationCache[id]` in O(1) instead of scanning a slice every time.

### The Search Bridge

A complete client → server → client loop with 250ms debounce to avoid flooding the server on every keystroke:

```
User types → 250ms debounce → fetch /search?q=... → Go handler
  → case-insensitive match on Name + Members + CreationDate
  → JSON response → JS rebuilds grid with correct CSS classes and animations
```

The Go handler searches across three fields simultaneously — artist name, individual member names, and creation year — making it useful for queries like "1994" or a specific member's name.

### Template Safety at Startup

```go
func init() {
    HomeTemplate   = template.Must(template.ParseFiles("templates/index.html"))
    ArtistTemplate = template.Must(template.ParseFiles("templates/artist.html"))
}
```

`template.Must` panics at startup if a template is invalid, rather than silently failing on the first user request. This is the correct Go pattern — fail loud and early.

### Fail-Fast Server Startup

```go
if err := InitCache(); err != nil {
    log.Fatalf("Could not initialize data: %v", err)
}
log.Fatal(http.ListenAndServe(":8080", nil))
```

If the external API is down when the server starts, it exits immediately with a clear error instead of starting in a broken state. The `log.Fatal` on `ListenAndServe` ensures OS-level errors are also surfaced.

---

## ⚔️ Technical Challenges

### 1. The Global State Problem
The first working version had `getArtists()` called inside `HomeHandler` and `getRelations()` called inside `ArtistHandler` — one external API call per user request. This worked locally but would fail under any real load and hammer the upstream API. Moving to a startup cache required restructuring the data flow entirely: handlers stopped being data-fetchers and became data-presenters.

### 2. Slice vs Map for Relations
`getRelations()` returns a `Relations` struct with `Index []RelationItem`. The first approach was to scan that slice inside `ArtistHandler` with a `for` loop on every request — O(n) per page load. Converting it to `map[int]RelationItem` at cache-init time made every artist page load O(1). A small change with a meaningful architectural difference.

### 3. The Search Rebuild Bug
When the JS rebuilt the grid after a search response, it used different CSS class names (`artist-card`, `artist-img-wrapper`) from the Go template (`card`, `card-img-wrap`). The cards rendered but received none of the hover effects, animations, or focus styles. The fix was a `buildCard()` function that mirrors the Go template's HTML exactly — making the JS and the template a contract, not two independent implementations.

### 4. Dead Code After Refactoring
After moving artist lookup to the cache, `getArtist(id int)` remained in `api.go` — compiling cleanly, never called. Go's compiler doesn't warn on unused functions (only unused imports and variables). Code review caught it. Lesson: refactoring requires actively tracing call sites, not trusting the compiler to catch everything.

### 5. `nil` vs `[]` in JSON
Go encodes a `nil` slice as JSON `null`, not `[]`. When `SearchHandler` found no results, `json.NewEncoder(w).Encode(results)` sent `null` to the client. The JS had to defensively check `if (!data || data.length === 0)`. The correct fix is to initialize `results` as an empty slice so the contract is always an array:

```go
results := []Artist{}
```

---

## 📖 Medium Article *(In Progress)*

I'm writing a detailed breakdown of this project for Go developers, covering the decisions above in depth — including the mistakes made before arriving at the current architecture.

**Topics will include:**
- Why global cache beats per-request fetch, and when it doesn't
- How Go's `html/template` differs from other templating engines and why that matters for XSS
- Building a JS ↔ Go search bridge without a frontend framework
- What Shneiderman's 8 Golden Rules look like in actual code
- The dead code problem: what Go's compiler catches and what it doesn't

🔗 **[Link coming soon]** — follow me to be notified when it drops.

---

## 🏁 How to Run

```bash
git clone https://github.com/Anasmoner2022/groupie-tracker.git
cd groupie-tracker
go run .
```

Open `http://localhost:8080`

> **Requires Go 1.18+.** No external dependencies — standard library only.

---

## 📁 Project Structure

```
groupie-tracker/
├── main.go          # Server entry point, route registration, startup cache
├── api.go           # External API calls + cache initialization
├── handlers.go      # HTTP handlers (Home, Artist, Search)
├── models.go        # Data types matching the external API's JSON shape
├── templates/
│   ├── index.html   # Artist grid with search + sort
│   └── artist.html  # Artist detail with tour dates
└── static/          # CSS / JS assets (served at /static/)
```

---

*Built as part of a structured backend engineering curriculum at 01edu.*
