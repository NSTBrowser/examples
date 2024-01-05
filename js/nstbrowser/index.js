// JavaScript version of the Go code

export class Option {
    constructor({host, port, apiKey}) {
        this.host = host;
        this.port = port;
        this.apiKey = apiKey;
        this.launchConfig = undefined;
    }

    validate() {
        if (!this.host) {
            throw new Error("host is empty");
        }
        if (!this.apiKey) {
            throw new Error("apiKey is empty");
        }
    }

    parseQuery() {
        const params = new URLSearchParams();
        params.append('x-api-key', this.apiKey);
        params.append('config', JSON.stringify(this.launchConfig));
        return params.toString();
    }
}

export class LaunchConfig {
    constructor() {
        this.once = false;
        this.headless = false;
        this.autoClose = false;
        this.clearCacheOnClose = false;
        this.remoteDebuggingPort = 0;
        this.fingerprint = undefined;
    }

    withOnce() {
        this.once = true;
        return this;
    }

    withHeadless() {
        this.headless = "new";
        return this;
    }

    withAutoClose() {
        this.autoClose = true;
        return this;
    }

    withClearCacheOnClose() {
        this.clearCacheOnClose = true;
        return this;
    }

    withRemoteDebuggingPort(port) {
        this.remoteDebuggingPort = port;
        return this;
    }

    withFingerprint(fingerprint) {
        this.fingerprint = fingerprint;
        return this;
    }
}

export class Fingerprint {
    constructor({
                    name,
                    platform = "windows",
                    kernel = "chromium",
                    kernelMilestone = "120",
                    hardwareConcurrency = 4,
                    deviceMemory = 16,
                    proxy = ""
                }) {
        this.name = name
        this.platform = platform
        this.kernel = kernel
        this.kernelMilestone = kernelMilestone
        this.hardwareConcurrency = hardwareConcurrency
        this.deviceMemory = deviceMemory
        this.proxy = proxy
    }
}

// getConnectToLaunchedBrowserWS
// https://apifox.nstbrowser.io/doc-3527176#3-connect-to-launched-browser
export function getConnectToLaunchedBrowserWS(option, profileId) {
    option.validate();
    const queryString = option.parseQuery();
    return `ws://${option.host}:${option.port}/devtool/browser/${profileId}?${queryString}`;
}

// getLaunchAndConnectToBrowserWS
// https://apifox.nstbrowser.io/doc-3527176#2-launchandconnecttobrowser
export function getLaunchAndConnectToBrowserWS(option, profileId, launchConfig) {
    option.validate();
    option.launchConfig = launchConfig;
    const queryString = option.parseQuery();
    return `ws://${option.host}:${option.port}/devtool/launch/${profileId}?${queryString}`;
}

// getCreateAndConnectToBrowserURL
// https://apifox.nstbrowser.io/doc-3527176#3-createandconnecttobrowser
export function getCreateAndConnectToBrowserURL(option, launchConfig) {
    option.validate();
    option.launchConfig = launchConfig;

    const queryString = option.parseQuery();
    return `ws://${option.host}:${option.port}/devtool/launch?${queryString}`;
}