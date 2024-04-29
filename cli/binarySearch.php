<?php

/**
 * Searches the passed array for the key of the passed value. If input checks fail or there is no number in the array, returns -1.
 * @param array $data
 * @param int $number
 * @return int
 */
function binarySearch(array $data, int $number): int
{
    $key = -1;
    if (count($data) < PHP_INT_MAX && is_int($number)) {
        $start = 0;
        $end = count($data) - 1;
        while ($start <= $end) {
            $middle = round(($start + $end) / 2);
            $checkedValue = $data[$middle];
            if ($checkedValue == $number) {
                $key = $middle;
                if ($data[$middle - 1] == $number) {
                    $key = $middle - 1;
                }
            } else if ($checkedValue > $number) {
                $end = $middle  - 1;
            } else {
                $start = $middle  + 1;
            }
            if ($key > -1) {
                break;
            }
        }
    }
    return $key;
}

$data = [4, 8, 15, 15, 16, 23, 42];

print binarySearch($data, 15);
