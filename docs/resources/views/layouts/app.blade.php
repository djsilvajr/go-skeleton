<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ config('app.name', 'API') }} Docs — {{ $domainLabel ?? 'Documentation' }}</title>

    <!-- Semantic UI -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/semantic-ui/2.5.0/semantic.min.css">

    <style>
        body { background: #f4f6f9; }

        /* ── Top Bar ── */
        #topbar {
            position: fixed; top: 0; left: 0; right: 0; z-index: 999;
            height: 56px;
            background: #1b1c1d;
            display: flex; align-items: center; gap: 12px;
            padding: 0 20px;
            box-shadow: 0 2px 6px rgba(0,0,0,.4);
        }
        #topbar .api-name {
            color: #fff; font-weight: 700; font-size: 1rem;
            white-space: nowrap; margin-right: 8px;
        }
        #topbar .env-badge {
            background: #21ba45; color: #fff;
            padding: 2px 10px; border-radius: 12px;
            font-size: .75rem; font-weight: 600; text-transform: uppercase;
        }
        #topbar .env-badge.staging { background: #f2711c; }
        #topbar .env-badge.production { background: #db2828; }
        #topbar .spacer { flex: 1; }
        #topbar .token-wrap { display: flex; align-items: center; gap: 6px; }
        #topbar .token-wrap label { color: #ccc; font-size: .8rem; white-space: nowrap; }
        #topbar input#global-token {
            background: #2d2e2f; border: 1px solid #444; color: #fff;
            border-radius: 4px; padding: 4px 10px; font-size: .82rem;
            width: 340px;
        }
        #topbar input#global-token::placeholder { color: #777; }
        #topbar .base-url-wrap { display: flex; align-items: center; gap: 6px; }
        #topbar .base-url-wrap label { color: #ccc; font-size: .8rem; white-space: nowrap; }
        #topbar input#global-base-url {
            background: #2d2e2f; border: 1px solid #444; color: #fff;
            border-radius: 4px; padding: 4px 10px; font-size: .82rem;
            width: 220px;
        }

        /* ── Sidebar ── */
        #sidebar {
            position: fixed; top: 56px; left: 0; bottom: 0;
            width: 220px; background: #23262b;
            overflow-y: auto; padding: 16px 0;
        }
        #sidebar .section-label {
            color: #888; font-size: .7rem; font-weight: 700;
            text-transform: uppercase; letter-spacing: .08em;
            padding: 8px 20px 4px;
        }
        #sidebar a {
            display: block; padding: 9px 20px;
            color: #ccc; font-size: .9rem;
            text-decoration: none;
            border-left: 3px solid transparent;
            transition: background .15s, color .15s;
        }
        #sidebar a:hover  { background: #2e3136; color: #fff; }
        #sidebar a.active { background: #2e3136; color: #fff; border-left-color: #2185d0; }

        /* ── Main content ── */
        #main {
            margin-left: 220px;
            margin-top: 56px;
            padding: 28px 32px;
            min-height: calc(100vh - 56px);
        }

        /* ── Route Card ── */
        .route-card {
            background: #fff; border-radius: 8px;
            box-shadow: 0 1px 4px rgba(0,0,0,.08);
            margin-bottom: 28px; overflow: hidden;
        }
        .route-card-header {
            padding: 16px 20px; display: flex; align-items: center; gap: 12px;
            border-bottom: 1px solid #eee; cursor: pointer;
            user-select: none;
        }
        .route-card-header:hover { background: #fafafa; }
        .route-card-body { padding: 20px; }
        .route-card-description { color: #555; margin-bottom: 18px; font-size: .93rem; }

        /* HTTP method badges */
        .method-badge {
            display: inline-block; padding: 3px 10px; border-radius: 4px;
            font-weight: 700; font-size: .78rem; text-transform: uppercase;
            min-width: 62px; text-align: center;
        }
        .method-GET    { background: #e8f5e9; color: #2e7d32; }
        .method-POST   { background: #e3f2fd; color: #1565c0; }
        .method-PUT    { background: #fff8e1; color: #e65100; }
        .method-PATCH  { background: #fff3e0; color: #bf360c; }
        .method-DELETE { background: #fce4ec; color: #c62828; }

        .endpoint-path { font-family: 'Courier New', monospace; font-size: .9rem; color: #333; }
        .route-name    { font-weight: 600; color: #222; font-size: .95rem; }
        .chevron { margin-left: auto; color: #aaa; transition: transform .2s; }
        .chevron.open { transform: rotate(180deg); }

        /* Section titles inside card */
        .section-title {
            font-size: .75rem; font-weight: 700; text-transform: uppercase;
            letter-spacing: .06em; color: #888; margin: 18px 0 8px;
        }
        .section-title:first-child { margin-top: 0; }

        /* Code blocks */
        pre.code-block {
            background: #282c34; color: #abb2bf;
            border-radius: 6px; padding: 14px 16px;
            font-size: .82rem; overflow-x: auto;
            margin: 8px 0;
        }

        /* Status badge */
        .status-badge {
            display: inline-block; padding: 2px 8px;
            border-radius: 10px; font-size: .75rem; font-weight: 700;
        }
        .status-2xx { background: #e8f5e9; color: #2e7d32; }
        .status-4xx { background: #fce4ec; color: #c62828; }
        .status-5xx { background: #fafafa;  color: #555; }

        /* No body notice */
        .no-body { color: #999; font-style: italic; font-size: .88rem; padding: 6px 0; }

        /* Try It panel */
        .try-it-panel { background: #f9f9f9; border: 1px solid #e8e8e8; border-radius: 6px; padding: 16px; margin-top: 10px; }
        .try-it-panel .field { margin-bottom: 10px; }
        .try-it-panel label { font-size: .82rem; font-weight: 600; color: #555; display: block; margin-bottom: 4px; }
        .try-it-panel input, .try-it-panel textarea {
            width: 100%; border: 1px solid #ddd; border-radius: 4px;
            padding: 7px 10px; font-size: .84rem; font-family: 'Courier New', monospace;
            background: #fff;
        }
        .try-it-panel textarea { min-height: 90px; resize: vertical; }
        .try-response { margin-top: 12px; }
        .try-response-meta { display: flex; gap: 16px; margin-bottom: 6px; font-size: .82rem; color: #555; }
        .try-response-meta span strong { color: #222; }
        .domain-title { font-size: 1.4rem; font-weight: 700; color: #1b1c1d; margin-bottom: 6px; }
        .domain-subtitle { color: #888; font-size: .9rem; margin-bottom: 24px; }
    </style>
</head>
<body>

{{-- ── Top Bar ── --}}
<div id="topbar">
    <span class="api-name">{{ config('app.name', 'API') }}</span>

    @php $env = app()->environment(); @endphp
    <span class="env-badge {{ $env }}">{{ $env }}</span>

    <span class="spacer"></span>

    <div class="base-url-wrap">
        <label for="global-base-url">Base URL</label>
        <input id="global-base-url" type="text" placeholder="{{ config('app.api_base_url') }}" value="{{ config('app.api_base_url') }}">
    </div>

    <div class="token-wrap">
        <label for="global-token">Bearer Token</label>
        <input id="global-token" type="text" placeholder="Paste your JWT here…">
    </div>
</div>

{{-- ── Sidebar ── --}}
<nav id="sidebar">
    <div class="section-label">Domains</div>
    @foreach ($domains as $slug => $label)
        <a href="{{ route('docs.domain', $slug) }}"
           class="{{ ($activeDomain ?? '') === $slug ? 'active' : '' }}">
            {{ $label }}
        </a>
    @endforeach
</nav>

{{-- ── Main Content ── --}}
<main id="main">
    @yield('content')
</main>

<!-- Semantic UI JS + jQuery -->
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.7.1/jquery.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/semantic-ui/2.5.0/semantic.min.js"></script>

<script>
    // ── Collapsible route cards ──
    document.querySelectorAll('.route-card-header').forEach(header => {
        header.addEventListener('click', () => {
            const body    = header.nextElementSibling;
            const chevron = header.querySelector('.chevron');
            const open    = body.style.display !== 'none';
            body.style.display = open ? 'none' : 'block';
            chevron.classList.toggle('open', !open);
        });
    });

    // ── Try It: send request ──
    document.querySelectorAll('.btn-send').forEach(btn => {
        btn.addEventListener('click', async () => {
            const panel    = btn.closest('.try-it-panel');
            const method   = btn.dataset.method;
            const rawUrl   = btn.dataset.endpoint;
            const baseUrl  = document.getElementById('global-base-url').value.replace(/\/$/, '');
            const token    = document.getElementById('global-token').value.trim();

            // Replace path params from input fields
            let url = baseUrl + rawUrl;
            panel.querySelectorAll('.path-param').forEach(input => {
                url = url.replace(`{${input.dataset.param}}`, encodeURIComponent(input.value || input.placeholder));
            });

            const headers = { 'Content-Type': 'application/json' };
            if (token) headers['Authorization'] = `Bearer ${token}`;

            const bodyInput = panel.querySelector('.body-input');
            const options   = { method, headers };
            if (bodyInput && bodyInput.value.trim()) {
                try { options.body = JSON.stringify(JSON.parse(bodyInput.value)); }
                catch { options.body = bodyInput.value; }
            }

            const statusEl  = panel.querySelector('.res-status');
            const timeEl    = panel.querySelector('.res-time');
            const bodyEl    = panel.querySelector('.res-body');
            const metaWrap  = panel.querySelector('.try-response');

            btn.classList.add('loading');
            metaWrap.style.display = 'none';

            const t0 = Date.now();
            try {
                const res  = await fetch(url, options);
                const ms   = Date.now() - t0;
                const text = await res.text();
                let pretty = text;
                try { pretty = JSON.stringify(JSON.parse(text), null, 2); } catch {}

                statusEl.textContent  = res.status + ' ' + res.statusText;
                statusEl.className    = 'status-badge ' + (res.ok ? 'status-2xx' : 'status-4xx');
                timeEl.textContent    = ms + ' ms';
                bodyEl.textContent    = pretty;
                metaWrap.style.display = 'block';
            } catch (err) {
                statusEl.textContent  = 'Network error';
                statusEl.className    = 'status-badge status-5xx';
                timeEl.textContent    = '-';
                bodyEl.textContent    = err.message;
                metaWrap.style.display = 'block';
            } finally {
                btn.classList.remove('loading');
            }
        });
    });
</script>

</body>
</html>
