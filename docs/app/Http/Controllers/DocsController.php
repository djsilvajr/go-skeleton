<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class DocsController extends Controller
{
    /**
     * Available domains and their data files.
     */
    private array $domains = [
        'auth'  => 'Auth',
        'users' => 'Users',
    ];

    /**
     * Show the docs index — redirects to the first domain.
     */
    public function index()
    {
        $firstSlug = array_key_first($this->domains);
        return redirect()->route('docs.domain', ['domain' => $firstSlug]);
    }

    /**
     * Show all routes for a specific domain.
     */
    public function domain(string $domain)
    {
        abort_unless(array_key_exists($domain, $this->domains), 404);

        $routes = $this->loadRoutes($domain);

        return view('docs.domain', [
            'domains'       => $this->domains,
            'activeDomain'  => $domain,
            'domainLabel'   => $this->domains[$domain],
            'routes'        => $routes,
        ]);
    }

    /**
     * Load and decode route data from the JSON file.
     */
    private function loadRoutes(string $domain): array
    {
        $path = base_path("data/routes/{$domain}.json");

        if (! file_exists($path)) {
            return [];
        }

        return json_decode(file_get_contents($path), true) ?? [];
    }
}
