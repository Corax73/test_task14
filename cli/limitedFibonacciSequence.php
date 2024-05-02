<?php

/**
 * If the argument is valid, returns the Fibonacci sequence string up to the value of the argument, otherwise -1.
 * @param int $limit
 * @return string
 */
function fiborow(int $limit): string
{
    $resp = '-1';
    if ($limit > 0 && $limit < PHP_INT_MAX) {
        $fib = [0, 1];
        while ($fib[count($fib) - 1] + $fib[count($fib) - 2] < $limit) {
            $fib[] = $fib[count($fib) - 1] + $fib[count($fib) - 2];
        }
        $resp = implode(' ', $fib);
    }
    return $resp;
}

print fiborow(10);
