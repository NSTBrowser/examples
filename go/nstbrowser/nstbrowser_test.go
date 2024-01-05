package nstbrowser

import (
  "context"
  "fmt"
  "github.com/chromedp/chromedp"
  "log"
  "os"
  "testing"
)

func TestGetConnectToLaunchedBrowserURL(t *testing.T) {
  option := Option{
    Host:   "192.168.50.64",
    Port:   8838,
    ApiKey: "896bf156-4f47-45ff-b79b-6c8597534eec",
  }
  ws, err := GetConnectToLaunchedBrowserWS(option, "b2dc29e1-3982-49e5-a2ba-0d80f27601fd")
  if err != nil {
    t.Fatal(err)
  }
  t.Log(ws)

  Screenshot(ws, "https://www.nstbrowser.io", "nstbrowser")
}

func TestGetLaunchAndConnectToBrowserURL(t *testing.T) {
  option := Option{
    Host:   "192.168.50.64",
    Port:   8838,
    ApiKey: "896bf156-4f47-45ff-b79b-6c8597534eec",
  }
  ws, err := GetLaunchAndConnectToBrowserWS(option, "b2dc29e1-3982-49e5-a2ba-0d80f27601fd",
    WithHeadless(),
  )
  if err != nil {
    t.Fatal(err)
  }
  t.Log(ws)
  Screenshot(ws, "https://www.google.com", "google")
}

func TestGetCreateAndConnectToBrowserWS(t *testing.T) {
  option := Option{
    Host:   "192.168.50.64",
    Port:   8838,
    ApiKey: "896bf156-4f47-45ff-b79b-6c8597534eec",
  }

  fingerprint := &Fingerprint{
    Name:                "Test_Once",
    Platform:            PlatformWindows,
    Kernel:              KernelChromium,
    KernelMilestone:     KernelMilestone113,
    HardwareConcurrency: 4,
    DeviceMemory:        8,
    Proxy:               "",
  }

  ws, err := GetCreateAndConnectToBrowserURL(option,
    WithOnce(),
    WithHeadless(),
    WithAutoClose(),
    WithFingerprint(fingerprint))
  if err != nil {
    t.Fatal(err)
  }
  t.Log(ws)

  Screenshot(ws, "https://chat.openai.com/", "openai")
}

func Screenshot(ws, pageURL, screenshotName string) {
  remoteAllocator, cancelFunc := chromedp.NewRemoteAllocator(context.Background(), ws, chromedp.NoModifyURL)
  defer cancelFunc()

  // create context
  ctx, cancel := chromedp.NewContext(remoteAllocator)
  defer cancel()

  var buf []byte
  if err := chromedp.Run(ctx, fullScreenshot(pageURL, 90, &buf)); err != nil {
    log.Fatal(err)
  }
  if err := os.WriteFile(fmt.Sprintf("%s.png", screenshotName), buf, 0o644); err != nil {
    log.Fatal(err)
  }
}

func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
  return chromedp.Tasks{
    chromedp.Navigate(urlstr),
    chromedp.FullScreenshot(res, quality),
  }
}
