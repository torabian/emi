package com.example.flightstick

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.compose.foundation.Image
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material3.Slider

import androidx.compose.runtime.Composable
import androidx.compose.runtime.mutableStateOf
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.TransformOrigin
import androidx.compose.ui.graphics.graphicsLayer
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.tooling.preview.Preview
import com.example.flightstick.ui.theme.FlightstickTheme
import androidx.compose.runtime.remember
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.getValue
import androidx.compose.runtime.setValue
import androidx.compose.ui.unit.dp
import kotlin.math.roundToInt

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContent {
            FlightstickTheme {
                Scaffold(modifier = Modifier.fillMaxSize()) { innerPadding ->
                    Column(
                        modifier = Modifier
                            .fillMaxSize()
                            .padding(innerPadding),
                        verticalArrangement = Arrangement.SpaceBetween, // split top & bottom
                        horizontalAlignment = Alignment.CenterHorizontally
                    ) {


                        var calibrationOffset by remember { mutableStateOf(0f) }

                        // Top: Attitude Indicator, takes most of the space
                        val (yaw, pitch, roll) = useGyroscopeStatus(calibrationOffset)

                        val status = StreamYokeAction("http://10.0.2.2:8080", Triple(yaw, pitch, roll))

                        // Bottom: Calibration slider
                        Column(horizontalAlignment = Alignment.CenterHorizontally) {
                            Slider(
                                value = calibrationOffset,
                                onValueChange = { calibrationOffset = it },
                                valueRange = -50f..50f,
                                modifier = Modifier.width(200.dp)
                            )
                            Text("Calibration: ${calibrationOffset.roundToInt()}")
                            Text("Socket: ${status}")

                        }

                        AttitudeIndicator(
                            pitch = pitch ,
                            roll = roll + calibrationOffset,
                            modifier = Modifier
                                .fillMaxWidth()
                                .weight(1f) // takes remaining space
                        )
                    }
                }

            }
        }
    }
}


@Composable
fun AttitudeIndicator(
    pitch: Float,
    roll: Float,
    modifier: Modifier = Modifier
) {
    Box(
        modifier = modifier
    ) {
        Image(
            painter = painterResource(id = R.drawable.attitude_indicator_deep_background),
            contentDescription = null
        )


        Box(
            modifier = Modifier
                .clip(CircleShape)
                .fillMaxSize()

        ) {
            Image(
                painter = painterResource(id = R.drawable.attitude_indicator_moving_part),
                contentDescription = null,
                modifier = Modifier  .graphicsLayer {
                    rotationZ = -pitch      // rotation in degrees
                    translationY = roll     // move vertically
                    transformOrigin = TransformOrigin(0.5f, 0.5f) // pivot at center
                }
            )
        }


        Image(
            painter = painterResource(id = R.drawable.attitude_indicator_cover),
            contentDescription = null
        )
    }
}

@Composable
fun Greeting(name: String, modifier: Modifier = Modifier) {
    Text(
        text = "New",
        modifier = modifier
    )
}

@Preview(showBackground = true)
@Composable
fun GreetingPreview() {
    FlightstickTheme {
        Greeting("Android")
    }
}


