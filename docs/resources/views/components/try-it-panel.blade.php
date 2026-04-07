{{--
    Component: try-it-panel
    Props:
        $routeIndex — unique integer to avoid ID collisions
        $method     — HTTP method string
        $endpoint   — endpoint path (e.g. /api/v1/users/{id})
        $hasBody    — boolean
        $body       — associative array of field => type, or null
--}}
@props(['routeIndex', 'method', 'endpoint', 'hasBody' => false, 'body' => null])

@php
    // Extract path parameters like {id}
    preg_match_all('/\{(\w+)\}/', $endpoint, $pathParams);
    $pathParams = $pathParams[1];
@endphp

<div class="section-title">Try It</div>
<div class="try-it-panel">

    {{-- Path parameters --}}
    @foreach ($pathParams as $param)
        <div class="field">
            <label>Path: <code>{{ '{' . $param . '}' }}</code></label>
            <input type="text"
                   class="path-param"
                   data-param="{{ $param }}"
                   placeholder="{{ $param }}">
        </div>
    @endforeach

    {{-- Body --}}
    @if ($hasBody && $body)
        <div class="field">
            <label>Request Body (JSON)</label>
            <textarea class="body-input" placeholder="{{ json_encode($body, JSON_PRETTY_PRINT) }}"></textarea>
        </div>
    @endif

    <button class="ui blue button btn-send"
            data-method="{{ $method }}"
            data-endpoint="{{ $endpoint }}">
        <i class="paper plane icon"></i> Send Request
    </button>

    {{-- Response viewer --}}
    <div class="try-response" style="display:none">
        <div class="try-response-meta">
            <span>Status: <strong><span class="res-status status-badge"></span></strong></span>
            <span>Time: <strong><span class="res-time"></span></strong></span>
        </div>
        <pre class="code-block res-body" style="max-height:320px;overflow-y:auto"></pre>
    </div>
</div>
