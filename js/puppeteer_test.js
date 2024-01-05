import puppeteer from 'puppeteer-core';
import {
    Fingerprint,
    getConnectToLaunchedBrowserWS,
    getCreateAndConnectToBrowserURL,
    getLaunchAndConnectToBrowserWS,
    LaunchConfig,
    Option,
} from "./nstbrowser/index.js";

async function TestGetConnectToLaunchedBrowserURL() {
    const option = new Option({
        host: "192.168.50.64",
        port: 8838,
        apiKey: "896bf156-4f47-45ff-b79b-6c8597534eec",
    })

    const ws = getConnectToLaunchedBrowserWS(option, "b2dc29e1-3982-49e5-a2ba-0d80f27601fd");
    await screenshot(ws, "https://www.nstbrowser.io", "nstbrowser")
}

async function TestGetLaunchAndConnectToBrowserWS() {
    const option = new Option({
        host: "192.168.50.64",
        port: 8838,
        apiKey: "896bf156-4f47-45ff-b79b-6c8597534eec",
    })

    const ws = getLaunchAndConnectToBrowserWS(option, "b2dc29e1-3982-49e5-a2ba-0d80f27601fd",
        new LaunchConfig()
            .withHeadless());

    await screenshot(ws, "https://www.google.io", "google")
}

async function TestGetLaunchAndConnectToBrowserWSWithFingerprint() {
    const option = new Option({
        host: "192.168.50.64",
        port: 8838,
        apiKey: "896bf156-4f47-45ff-b79b-6c8597534eec",
    })

    const ws = getCreateAndConnectToBrowserURL(option, new LaunchConfig()
        .withOnce()
        .withHeadless()
        .withAutoClose()
        .withFingerprint(new Fingerprint({
            name: "Test_Once",
            platform: "windows",
            kernel: "chromium",
            kernelMilestone: "120",
            hardwareConcurrency: 4,
            deviceMemory: 8,
            proxy: ""
        }))
    );
    console.log(ws)
    await screenshot(ws, "https://chat.openai.com/", "openai")
}

async function screenshot(ws, pageURL, screenshotName) {
    try {
        const browser = await puppeteer.connect({
            browserWSEndpoint: ws,
            defaultViewport: null,
        });

        const page = await browser.newPage();
        await page.goto(pageURL);
        await page.screenshot({path: `${screenshotName}.png`});
        await page.close();
        await browser.disconnect();
    } catch (err) {
        console.error(err);
    }
}

// TestGetConnectToLaunchedBrowserURL().then()
// TestGetLaunchAndConnectToBrowserWS().then()
TestGetLaunchAndConnectToBrowserWSWithFingerprint().then()