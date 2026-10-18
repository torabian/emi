// Browser client: microphone + camera -> the Go server, using the generated
// StreamMediaAction and the MediaStreamer helper shipped with emi's JS sdk.
import { StreamMediaAction, StreamMediaActionQueryParams } from "./sdk/StreamMediaAction";
import { MediaStreamer } from "./sdk/sdk/common/MediaStreamer";

const SAMPLE_RATE = 16000;
const $ = (id) => document.getElementById(id);

let socket;
let streamer;
let statsTimer;

function log(line) {
  $("log").textContent += line + "\n";
}

async function start() {
  const deviceId = $("device").value.trim() || "browser";

  // Same typed query the C/C++/Go sides use, generated from media.emi.yml.
  const qs = new StreamMediaActionQueryParams()
    .setSampleRate(SAMPLE_RATE)
    .setChannels(1)
    .setAudioCodec("pcm16")
    .setVideoCodec("mjpeg");

  const url = location.origin.replace(/^http/, "ws") + StreamMediaAction.NewUrl({ deviceId }, qs);
  socket = StreamMediaAction.Create(url, qs, { deviceId }); // binaryType = arraybuffer, from the yaml
  streamer = new MediaStreamer(socket, { sampleRate: SAMPLE_RATE, fps: 5, width: 320, height: 240 });
  window.__streamer = streamer;

  socket.addEventListener("open", async () => {
    log("connected: " + url);
    try {
      await streamer.start();
      log("capturing microphone + camera");
    } catch (e) {
      log("capture failed: " + e);
    }
  });
  socket.addEventListener("close", () => {
    log("closed");
    stop();
  });
  socket.addEventListener("error", () => log("socket error"));

  statsTimer = setInterval(() => {
    const s = streamer.stats;
    $("stats").textContent = `audio frames ${s.audioFrames} · video frames ${s.videoFrames} · skipped ${s.skippedVideoFrames}`;
  }, 250);
  $("start").disabled = true;
  $("stop").disabled = false;
}

function stop() {
  clearInterval(statsTimer);
  streamer?.stop();
  if (socket && socket.readyState <= WebSocket.OPEN) socket.close();
  $("start").disabled = false;
  $("stop").disabled = true;
}

$("start").onclick = start;
$("stop").onclick = stop;
$("preview").src = "";

// ?auto=1&seconds=4 starts immediately and stops itself (used by the headless test).
const params = new URLSearchParams(location.search);
if (params.get("auto")) {
  start();
  setTimeout(stop, Number(params.get("seconds") || 4) * 1000);
}
