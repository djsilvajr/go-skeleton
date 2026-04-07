@extends('layouts.app')

@section('content')

    <div class="domain-title">{{ $domainLabel }}</div>
    <div class="domain-subtitle">
        {{ count($routes) }} {{ Str::plural('route', count($routes)) }} documented
    </div>

    @forelse ($routes as $index => $route)
        <x-route-card :route="$route" :routeIndex="$index" />
    @empty
        <div class="ui warning message">
            <div class="header">No routes found</div>
            <p>There are no documented routes for this domain yet.</p>
        </div>
    @endforelse

@endsection
