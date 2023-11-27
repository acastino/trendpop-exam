<?php

namespace App\Http\Middleware;

use Illuminate\Http\JsonResponse;

class AlwaysPrettyJson
{
    /**
     * Apply pretty print
     * 
     * @param $request
     * @param \Closure $next
     * @return mixed
     */
    public function handle($request, \Closure $next)
    {
        $response = $next($request);

        if ($response instanceof JsonResponse) {
            $response->setEncodingOptions(JSON_PRETTY_PRINT);
        }

        return $response;
    }
}
