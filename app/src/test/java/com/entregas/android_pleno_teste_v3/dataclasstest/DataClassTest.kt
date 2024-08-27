package com.entregas.android_pleno_teste_v3.dataclasstest

import com.entregas.android_pleno_teste_v3.domain.PhotoRequestDataClass
import com.entregas.android_pleno_teste_v3.domain.GPSRequestDataClass
import com.entregas.android_pleno_teste_v3.domain.GyroscopeRequestDataClass
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNotEquals
import org.junit.Test

class DataClassesTest {

    @Test
    fun `test FotoRequestDataClass creation`() {
        // Arrange
        val pictureBase64 = "exampleBase64String"
        val macAddress = "00:14:22:01:23:45"

        // Act
        val fotoRequest = PhotoRequestDataClass(pictureBase64, macAddress)

        // Assert
        assertEquals(pictureBase64, fotoRequest.pictureBase64)
        assertEquals(macAddress, fotoRequest.macAddress)
    }

    @Test
    fun `test FotoRequestDataClass toString`() {
        // Arrange
        val pictureBase64 = "exampleBase64String"
        val macAddress = "00:14:22:01:23:45"
        val fotoRequest = PhotoRequestDataClass(pictureBase64, macAddress)

        // Act
        val expectedToString = "FotoRequestDataClass(pictureBase64=$pictureBase64, macAddress=$macAddress)"

        // Assert
        assertEquals(expectedToString, fotoRequest.toString())
    }

    @Test
    fun `test FotoRequestDataClass equals and hashCode`() {
        // Arrange
        val fotoRequest1 = PhotoRequestDataClass("exampleBase64String", "00:14:22:01:23:45")
        val fotoRequest2 = PhotoRequestDataClass("exampleBase64String", "00:14:22:01:23:45")
        val fotoRequest3 = PhotoRequestDataClass("differentBase64String", "00:14:22:01:23:45")

        // Act & Assert
        assertEquals(fotoRequest1, fotoRequest2) // should be equal
        assertEquals(fotoRequest1.hashCode(), fotoRequest2.hashCode()) // should have same hashCode

        assertNotEquals(fotoRequest1, fotoRequest3) // should not be equal
    }

    @Test
    fun `test GPSRequestDataClass creation`() {
        // Arrange
        val latitude = 12.34
        val longitude = 56.78
        val macAddress = "00:14:22:01:23:45"

        // Act
        val gpsRequest = GPSRequestDataClass(latitude, longitude, macAddress)

        // Assert
        assertEquals(latitude, gpsRequest.latitude, 0.0)
        assertEquals(longitude, gpsRequest.longitude, 0.0)
        assertEquals(macAddress, gpsRequest.macAddress)
    }

    @Test
    fun `test GPSRequestDataClass toString`() {
        // Arrange
        val latitude = 12.34
        val longitude = 56.78
        val macAddress = "00:14:22:01:23:45"
        val gpsRequest = GPSRequestDataClass(latitude, longitude, macAddress)

        // Act
        val expectedToString = "GPSRequestDataClass(latitude=$latitude, longitude=$longitude, macAddress=$macAddress)"

        // Assert
        assertEquals(expectedToString, gpsRequest.toString())
    }

    @Test
    fun `test GPSRequestDataClass equals and hashCode`() {
        // Arrange
        val gpsRequest1 = GPSRequestDataClass(12.34, 56.78, "00:14:22:01:23:45")
        val gpsRequest2 = GPSRequestDataClass(12.34, 56.78, "00:14:22:01:23:45")
        val gpsRequest3 = GPSRequestDataClass(98.76, 54.32, "00:14:22:01:23:45")

        // Act & Assert
        assertEquals(gpsRequest1, gpsRequest2) // should be equal
        assertEquals(gpsRequest1.hashCode(), gpsRequest2.hashCode()) // should have same hashCode

        assertNotEquals(gpsRequest1, gpsRequest3) // should not be equal
    }

    @Test
    fun `test GyroscopeRequestDataClass creation`() {
        // Arrange
        val axlex = 0.01f
        val axley =0.02f
        val axlez = 0.03f
        val macAddress = "00:14:22:01:23:45"

        // Act
        val gyroscopeRequest = GyroscopeRequestDataClass(axlex, axley, axlez, macAddress)

        // Assert
        assertEquals(axlex, gyroscopeRequest.axlex)
        assertEquals(axley, gyroscopeRequest.axley)
        assertEquals(axlez, gyroscopeRequest.axlez)
        assertEquals(macAddress, gyroscopeRequest.macAddress)
    }

    @Test
    fun `test GyroscopeRequestDataClass toString`() {
        // Arrange
        val axlex = 0.01f
        val axley = 0.02f
        val axlez = 0.03f
        val macAddress = "00:14:22:01:23:45"
        val gyroscopeRequest = GyroscopeRequestDataClass(axlex, axley, axlez, macAddress)

        // Act
        val expectedToString = "GyroscopeRequestDataClass(axlex=$axlex, axley=$axley, axlez=$axlez, macAddress=$macAddress)"

        // Assert
        assertEquals(expectedToString, gyroscopeRequest.toString())
    }

    @Test
    fun `test GyroscopeRequestDataClass equals and hashCode`() {
        // Arrange
        val gyroscopeRequest1 = GyroscopeRequestDataClass(0.01f, 0.02f, 0.03f, "00:14:22:01:23:45")
        val gyroscopeRequest2 = GyroscopeRequestDataClass(0.01f, 0.02f, 0.03f, "00:14:22:01:23:45")
        val gyroscopeRequest3 = GyroscopeRequestDataClass(0.10f, 0.20f, 0.30f, "00:14:22:01:23:45")

        // Act & Assert
        assertEquals(gyroscopeRequest1, gyroscopeRequest2) // should be equal
        assertEquals(gyroscopeRequest1.hashCode(), gyroscopeRequest2.hashCode()) // should have same hashCode

        assertNotEquals(gyroscopeRequest1, gyroscopeRequest3) // should not be equal
    }
}
