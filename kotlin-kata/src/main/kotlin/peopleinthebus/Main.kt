package de.roskenet.peopleinthebus

import kotlin.test.assertEquals
import kotlin.test.Test

fun people(busStops: Array<Pair<Int, Int>>) : Int {
    //code here
    return 0
}

class TestExample {
    @Test
    fun `Basic Tests`() {
        assertEquals(17, people(arrayOf(3 to 0,9 to 1,4 to 10,12 to 2,6 to 1,7 to 10)))
        assertEquals(21, people(arrayOf(3 to 0,9 to 1,4 to 8,12 to 2,6 to 1,7 to 8)))
        assertEquals(5, people(arrayOf(10 to 0,3 to 5,5 to 8)))
    }
}