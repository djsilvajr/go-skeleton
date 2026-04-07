{{--
    Component: error-table
    Props:
        $errors — array of { code, message, description }
--}}
@props(['errors' => []])

<div class="section-title">Possible Errors</div>

@if(count($errors))
    <table class="ui very basic compact table">
        <thead>
            <tr>
                <th>Code</th>
                <th>Message</th>
                <th>Description</th>
            </tr>
        </thead>
        <tbody>
            @foreach ($errors as $e)
            <tr>
                <td>
                    <span class="status-badge {{ $e['code'] >= 500 ? 'status-5xx' : 'status-4xx' }}">
                        {{ $e['code'] }}
                    </span>
                </td>
                <td>{{ $e['message'] }}</td>
                <td>{{ $e['description'] }}</td>
            </tr>
            @endforeach
        </tbody>
    </table>
@else
    <p class="no-body">No documented errors</p>
@endif
