package org.example.money

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

/**
 * Money is a hand-written "complex" type (see `type: complex, complex: Money` fields in
 * inventory.emi.yml/billing.emi.yml, and their `complexes:` entries pointing at this
 * exact location) - a small value class the emi compiler never generates, only
 * references. Deliberately not a plain Double/Long: an amount stored as integer minor
 * units (cents) alongside its currency, the same reasoning Go/Swift's own hand-written
 * Money examples give for using a "complex" field instead of a plain numeric one.
 */
@Serializable
data class Money(
    @SerialName("currency") val currency: String = "USD",
    @SerialName("minorUnits") val minorUnits: Long = 0,
) {
    fun toDisplayString(): String {
        val whole = minorUnits / 100
        val fraction = (minorUnits % 100).let { if (it < 0) -it else it }
        return "$currency ${whole}.${fraction.toString().padStart(2, '0')}"
    }
}
