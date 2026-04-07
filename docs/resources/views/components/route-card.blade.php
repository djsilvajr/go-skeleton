{{--
    Component: route-card
    Props:
        $route       — full route array from JSON
        $routeIndex  — unique integer for DOM scoping
--}}
@props(['route', 'routeIndex'])

@php
    $method   = $route['method'];
    $endpoint = $route['endpoint'];
    $name     = $route['name'];
    $desc     = $route['description'] ?? '';
    $request  = $route['request'] ?? [];
    $success  = $route['response']['success'] ?? null;
    $errors   = $route['errors'] ?? [];
    $example  = $route['example_request'] ?? null;
    $hasBody  = !empty($request['body']);
@endphp

<div class="route-card" id="route-{{ $routeIndex }}">

    {{-- Header (clickable) --}}
    <div class="route-card-header">
        <span class="method-badge method-{{ $method }}">{{ $method }}</span>
        <span class="endpoint-path">{{ $endpoint }}</span>
        <span class="route-name" style="margin-left:12px">{{ $name }}</span>
        <i class="chevron down icon open"></i>
    </div>

    {{-- Body --}}
    <div class="route-card-body">
        <p class="route-card-description">{{ $desc }}</p>

        {{-- 1. Request --}}
        <x-request-table
            :headers="$request['headers'] ?? []"
            :query="$request['query'] ?? []"
            :body="$request['body'] ?? null"
        />

        {{-- 2. Example Request --}}
        @if ($example)
            <div class="section-title">Example Request</div>

            @if ($example['json'])
                <div style="margin-bottom:4px;font-size:.78rem;color:#888">JSON</div>
                <pre class="code-block">{{ json_encode($example['json'], JSON_PRETTY_PRINT | JSON_UNESCAPED_SLASHES) }}</pre>
            @endif

            <div style="margin-top:8px;margin-bottom:4px;font-size:.78rem;color:#888">cURL</div>
            <pre class="code-block">{{ $example['curl'] }}</pre>
        @endif

        {{-- 3. Response --}}
        @if ($success)
            <x-response-table :success="$success" />
        @endif

        {{-- 4. Errors --}}
        <x-error-table :errors="$errors" />

        {{-- 5. Try It --}}
        <x-try-it-panel
            :routeIndex="$routeIndex"
            :method="$method"
            :endpoint="$endpoint"
            :hasBody="$hasBody"
            :body="$request['body'] ?? null"
        />
    </div>
</div>
