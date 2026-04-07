{{--
    Component: request-table
    Props:
        $headers  — array of { name, value, required }
        $query    — array of { name, type, required, description }
        $body     — associative array of field => type, or null
--}}
@props(['headers' => [], 'query' => [], 'body' => null])

{{-- Headers --}}
@if(count($headers))
    <div class="section-title">Headers</div>
    <table class="ui very basic compact table">
        <thead>
            <tr>
                <th>Name</th>
                <th>Value</th>
                <th>Required</th>
            </tr>
        </thead>
        <tbody>
            @foreach ($headers as $h)
            <tr>
                <td><code>{{ $h['name'] }}</code></td>
                <td><code>{{ $h['value'] }}</code></td>
                <td>
                    @if($h['required'])
                        <span class="ui mini red label">Yes</span>
                    @else
                        <span class="ui mini grey label">No</span>
                    @endif
                </td>
            </tr>
            @endforeach
        </tbody>
    </table>
@endif

{{-- Query Params --}}
@if(count($query))
    <div class="section-title">Query Params</div>
    <table class="ui very basic compact table">
        <thead>
            <tr>
                <th>Name</th>
                <th>Type</th>
                <th>Required</th>
                <th>Description</th>
            </tr>
        </thead>
        <tbody>
            @foreach ($query as $q)
            <tr>
                <td><code>{{ $q['name'] }}</code></td>
                <td>{{ $q['type'] }}</td>
                <td>
                    @if($q['required'])
                        <span class="ui mini red label">Yes</span>
                    @else
                        <span class="ui mini grey label">No</span>
                    @endif
                </td>
                <td>{{ $q['description'] ?? '' }}</td>
            </tr>
            @endforeach
        </tbody>
    </table>
@endif

{{-- Body --}}
<div class="section-title">Body (JSON)</div>
@if($body)
    <table class="ui very basic compact table">
        <thead>
            <tr>
                <th>Field</th>
                <th>Type</th>
            </tr>
        </thead>
        <tbody>
            @foreach ($body as $field => $type)
            <tr>
                <td><code>{{ $field }}</code></td>
                <td>{{ $type }}</td>
            </tr>
            @endforeach
        </tbody>
    </table>
@else
    <p class="no-body">No body required</p>
@endif
