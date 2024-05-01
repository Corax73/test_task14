<?php

/**
 * Converts a set of rgb color values ​​to an integer.
 * @param int $red
 * @param int $green
 * @param int $blue
 * @return int
 */
function rgbToInt(int $red, int $green, int $blue): int
{
    $result = -1;
    if ($red >= 0 && $red <= 255 && $green >= 0 && $green <= 255 && $blue >= 0 && $blue <= 255) {
        $str = '';
        foreach ([$red, $green, $blue] as $color) {
            $str .= str_pad(dechex($color), 2, '0', STR_PAD_LEFT);
        }
        $result = hexdec($str);
    }
    return $result;
}

print rgbToInt(255, 0, 255);
