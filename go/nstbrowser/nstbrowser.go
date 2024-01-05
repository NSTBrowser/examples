package nstbrowser

import (
  "encoding/json"
  "fmt"
  "net/url"
)

type Option struct {
  Host         string
  Port         int
  ApiKey       string
  launchConfig *LaunchConfig
}

func (opt *Option) validate() error {
  if opt.Host == "" {
    return fmt.Errorf("host is empty")
  }
  if opt.ApiKey == "" {
    return fmt.Errorf("apiKey is empty")
  }
  return nil
}

func (opt *Option) parseQuery() string {
  values := url.Values{
    "x-api-key": {opt.ApiKey},
  }
  if opt.launchConfig != nil {
    marshal, _ := json.Marshal(opt.launchConfig)
    queryEscape := url.QueryEscape(string(marshal))
    values.Add("config", queryEscape)
  }
  return values.Encode()
}

// LaunchConfig https://apifox.nstbrowser.io/doc-3527176#4-config-parameters-notes
type LaunchConfig struct {
  Once                bool         `json:"once"`
  Headless            any          `json:"headless"`
  AutoClose           bool         `json:"autoClose"`
  ClearCacheOnClose   bool         `json:"clearCacheOnClose"`
  RemoteDebuggingPort int          `json:"remoteDebuggingPort"`
  Fingerprint         *Fingerprint `json:"fingerprint"`
}

type Fingerprint struct {
  Name                string          `json:"name"`
  Platform            Platform        `json:"platform"`
  Kernel              Kernel          `json:"kernel"`
  KernelMilestone     KernelMilestone `json:"kernelMilestone"`
  HardwareConcurrency int             `json:"hardwareConcurrency"`
  DeviceMemory        int             `json:"deviceMemory"`
  Proxy               string          `json:"proxy"`
}

// GetConnectToLaunchedBrowserWS
// https://apifox.nstbrowser.io/doc-3527176#3-connect-to-launched-browser
func GetConnectToLaunchedBrowserWS(option Option, profileId string) (string, error) {
  if err := option.validate(); err != nil {
    return "", err
  }
  return fmt.Sprintf("ws://%s:%d/devtool/browser/%s?%s", option.Host, option.Port, profileId, option.parseQuery()), nil
}

// GetLaunchAndConnectToBrowserWS
// https://apifox.nstbrowser.io/doc-3527176#2-launchandconnecttobrowser
func GetLaunchAndConnectToBrowserWS(option Option, profileId string, launchOptions ...LaunchOption) (string, error) {
  if err := option.validate(); err != nil {
    return "", err
  }

  if option.launchConfig == nil {
    option.launchConfig = &LaunchConfig{}
  }
  for _, launchOption := range launchOptions {
    launchOption(option.launchConfig)
  }
  return fmt.Sprintf("ws://%s:%d/devtool/launch/%s?%s", option.Host, option.Port, profileId, option.parseQuery()), nil
}

// GetCreateAndConnectToBrowserURL
// https://apifox.nstbrowser.io/doc-3527176#3-createandconnecttobrowser
func GetCreateAndConnectToBrowserURL(option Option, launchOptions ...LaunchOption) (string, error) {
  if err := option.validate(); err != nil {
    return "", err
  }
  if option.launchConfig == nil {
    option.launchConfig = &LaunchConfig{}
  }
  for _, launchOption := range launchOptions {
    launchOption(option.launchConfig)
  }
  if option.launchConfig.Fingerprint == nil {
    return "", fmt.Errorf("fingerprint is empty")
  }
  return fmt.Sprintf("ws://%s:%d/devtool/launch?%s", option.Host, option.Port, option.parseQuery()), nil
}

type LaunchOption = func(config *LaunchConfig)

func WithOnce() LaunchOption {
  return func(config *LaunchConfig) {
    config.Once = true
  }
}

func WithHeadless() LaunchOption {
  return func(config *LaunchConfig) {
    config.Headless = "new"
  }
}

func WithAutoClose() LaunchOption {
  return func(config *LaunchConfig) {
    config.AutoClose = true
  }
}

func WithClearCacheOnClose() LaunchOption {
  return func(config *LaunchConfig) {
    config.ClearCacheOnClose = true
  }
}

func WithRemoteDebuggingPort(port int) LaunchOption {
  return func(config *LaunchConfig) {
    config.RemoteDebuggingPort = port
  }
}

func WithFingerprint(fingerprint *Fingerprint) LaunchOption {
  return func(config *LaunchConfig) {
    config.Fingerprint = fingerprint
  }
}

type Platform = string

var (
  PlatformWindows = "windows"
  PlatformMac     = "mac"
  PlatformLinux   = "linux"
)

type Kernel = string

var (
  KernelChromium = "chromium"
)

type KernelMilestone = string

var (
  KernelMilestone113 = "113"
  KernelMilestone115 = "115"
  KernelMilestone118 = "118"
  KernelMilestone120 = "120"
)
