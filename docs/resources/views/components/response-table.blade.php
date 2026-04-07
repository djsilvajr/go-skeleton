{{--
    Component: response-table
    Props:
        $success — { status, body, fields[] }
--}}
@props(['success'])

<div class="section-title">Success Response</div>

<span class="status-badge status-2xx">{{ $success['status'] }}</span>

<div class="section-title" style="margin-top:14px">Example</div>
<pre class="code-block">{{ json_encode($success['body'], JSON_PRETTY_PRINT | JSON_UNESCAPED_SLASHES) }}</pre>

@if(!empty($success['fields']))
    <div class="section-title">Fields</div>
    <table class="ui very basic compact table">
        <thead>
            <tr>
                <th>Field</th>
                <th>Type</th>
                <th>Description</th>
            </tr>
        </thead>
        <tbody>
            @foreach ($success['fields'] as $f)
            <tr>
                <td><code>{{ $f['field'] }}</code></td>
                <td>{{ $f['type'] }}</td>
                <td>{{ $f['description'] }}</td>
            </tr>
            @endforeach
        </tbody>
    </table>
@endif
