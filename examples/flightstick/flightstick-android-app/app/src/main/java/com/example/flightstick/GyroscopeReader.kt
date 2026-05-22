package com.example.flightstick

import android.content.Context
import android.hardware.Sensor
import android.hardware.SensorEvent
import android.hardware.SensorEventListener
import android.hardware.SensorManager
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Text
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.unit.dp
import kotlin.math.PI
import kotlin.math.roundToInt

@Composable
fun useGyroscopeStatus(offset: Float = 0f): Triple<Float, Float, Float> {
    val context = LocalContext.current
    var azimuth by remember { mutableStateOf(0f) } // rotation around Z (yaw)
    var pitch by remember { mutableStateOf(0f) }   // rotation around X
    var roll by remember { mutableStateOf(0f) }    // rotation around Y

    // Inline helper for radians → degrees
    fun toDegrees(rad: Double): Float = (rad * 180.0 / PI).toFloat()

    DisposableEffect(Unit) {
        val sensorManager = context.getSystemService(Context.SENSOR_SERVICE) as SensorManager
        val rotationSensor = sensorManager.getDefaultSensor(Sensor.TYPE_ROTATION_VECTOR)
        if (rotationSensor == null) {
            azimuth = Float.NaN
            pitch = Float.NaN
            roll = Float.NaN
            return@DisposableEffect onDispose { }
        }

        val rotationMatrix = FloatArray(9)
        val orientationAngles = FloatArray(3)

        val listener = object : SensorEventListener {
            override fun onSensorChanged(event: SensorEvent) {
                SensorManager.getRotationMatrixFromVector(rotationMatrix, event.values)
                SensorManager.getOrientation(rotationMatrix, orientationAngles)

                azimuth = toDegrees(orientationAngles[0].toDouble())
                pitch = toDegrees(orientationAngles[1].toDouble())
                roll = toDegrees(orientationAngles[2].toDouble())
            }

            override fun onAccuracyChanged(sensor: Sensor?, accuracy: Int) {}
        }

        sensorManager.registerListener(listener, rotationSensor, SensorManager.SENSOR_DELAY_UI)

        onDispose {
            sensorManager.unregisterListener(listener)
        }
    }

    return Triple(azimuth, pitch, roll * 2 + offset)
}
