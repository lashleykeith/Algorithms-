// O(nlog(n)) time | O(1) space 
package main

import (
    "sort"
)

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool{
    sort.Slice(redShirtHeights, func(i, j int) bool{
        return redShirtHeights[i] > redShirtHeights[j]
    })

    sort.Slice(blueShirtHeights, func(i, j int) bool{
        return blueShirtHeights[i] > blueShirtHeights[j]
    })

    shirtColorInFirstRow := "BLUE"
    if redShirtHeights[0] < blueShirtHeights[0]{
        shirtColorInFirstRow = "RED"
    }

   for idx := range redShirtHeights{
    redShirtHeight := redShirtHeights[idx]
    blueShirtHeight := blueShirtHeights[idx]

    if shirtColorInFirstRow == "RED"{
        if redShirtHeight >= blueShirtHeight{
            return false
        }
    } else {
        if blueShirtHeight >= redShirtHeight{
            return false
        }
    }
   }

   return true
}

/*
Assuming these are the redShirtHeights and blueShirtHeights can you edit the code so that the new order of redShirtHeights and blueShirtHeights is printed out.

These are the arrays:

{
  "redShirtHeights": [5, 8, 1, 3, 4],
  "blueShirtHeights": [6, 9, 2, 4, 5]
}

This is the code:

package main

import (
    "sort"
)

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool{
    sort.Slice(redShirtHeights, func(i, j int) bool{
        return redShirtHeights[i] > redShirtHeights[j]
    })

    sort.Slice(blueShirtHeights, func(i, j int) bool{
        return blueShirtHeights[i] > blueShirtHeights[j]
    })

    shirtColorInFirstRow := "BLUE"
    if redShirtHeights[0] < blueShirtHeights[0]{
        shirtColorInFirstRow = "RED"
    }

   for idx := range redShirtHeights{
    redShirtHeight := redShirtHeights[idx]
    blueShirtHeight := blueShirtHeights[idx]

    if shirtColorInFirstRow == "RED"{
        if redShirtHeight >= blueShirtHeight{
            return false
        }
    } else {
        if blueShirtHeight >= redShirtHeight{
            return false
        }
    }
   }

   return true
}
*/

