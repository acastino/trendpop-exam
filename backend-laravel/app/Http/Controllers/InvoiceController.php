<?php

namespace App\Http\Controllers;

use App\Http\Requests\StoreInvoiceRequest;
use Illuminate\Http\Request;
use App\Models\Invoice;

class InvoiceController extends Controller
{
    public function index(Request $request)
    {
        return Invoice::orderBy('id', 'desc')->simplePaginate();
    }

    public function store(StoreInvoiceRequest $request)
    {
        return Invoice::create($request->validated());
    }
}
