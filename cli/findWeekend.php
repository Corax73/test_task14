<?php

/**
 * If the arguments are correct, returns the number of days off in the passed time period, otherwise returns -1.
 * @param string $begin
 * @param string $end
 * @return int
 */
function weekend(string $begin, string $end): int
{
    $weekend = -1;
    $begin = strtotime($begin);
    $end = strtotime($end);
    if ($begin && $end && $begin < PHP_INT_MAX && $end < PHP_INT_MAX && $begin < $end) {
        $weekend = 0;
        while ($begin <= $end) {
            if (in_array(date('w', $begin), [5, 6])) {
                $weekend += 1;
                $begin += 86400;
            } elseif (date('w', $begin) < 5) {
                $begin += 86400 * (5 - date('w', $begin));
            }
        }
    }
    return $weekend;
}

$begin = '05.05.2024';
$end = '25.05.2024';
print weekend($begin, $end);
