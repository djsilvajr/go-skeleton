<?php

use App\Http\Controllers\DocsController;
use Illuminate\Support\Facades\Route;

Route::get('/', [DocsController::class, 'index'])->name('docs.index');
Route::get('/{domain}', [DocsController::class, 'domain'])->name('docs.domain');
